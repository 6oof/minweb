package kernel

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/6oof/minweb/app/configs"
	"github.com/6oof/minweb/app/helpers"
	"github.com/go-chi/chi/v5"
)

// Serve starts the MiniWeb server with specified configurations.
func Serve(mux *chi.Mux) {

	appPort := fmt.Sprintf(":%s", helpers.Env("PORT", "3003"))
	// Create a server with timeouts

	server := configs.ServerConfig()
	server.Addr = appPort
	server.Handler = mux

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		// Trigger graceful shutdown
		err := server.Shutdown(ctx)
		if err != nil {
			log.Println(err)
		}
	}()

	// Start the server in a goroutine so that it doesn't block.
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	yellow := "\033[33m"
	// ANSI escape code to reset text color
	reset := "\033[0m"

	fmt.Println("")
	fmt.Printf("%s>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>%s\n", yellow, reset)
	fmt.Printf("%s /\\  /\\\\  /\\  /%s\n", yellow, reset)
	fmt.Printf("%s/  \\/  \\\\/  \\/%s\n", yellow, reset)
	fmt.Printf("%sServer running on port %s%s (http://localhost:%s)\n", yellow, appPort, reset, appPort)
	fmt.Println("Using docker-compose? visit http://localhost:8080)")
	fmt.Printf("%s<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<%s\n", yellow, reset)
	fmt.Println("")

	// Block until an interrupt signal is received.
	<-sig
	log.Println("Shutting down...")
}
