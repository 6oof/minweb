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
	"github.com/6oof/minweb/app/middleware"
	"github.com/6oof/minweb/app/router"
	"github.com/go-chi/chi/v5"
)

// MiniWeb represents the main structure for the MiniWeb application, including its router.
type MiniWeb struct {
	Router *chi.Mux
}

// MbinInit initializes the MiniWeb application, sets up middleware, and registers routes.
func MbinInit() *MiniWeb {
	r := chi.NewRouter()

	// Middleware setup
	r.Use(middleware.Logger())
	r.Use(middleware.Recoverer())
	r.Use(middleware.DevPanicPrint)
	r.Use(middleware.Cors())

	// Static file serving
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Register custom routes
	router.RegisterWebRoutes(r)
	router.RegisterApiRoutes(r)
	router.RegisterFragmentRoutes(r)

	// Apply CSRF protection to all routes
	miniWeb := &MiniWeb{
		Router: r,
	}

	return miniWeb
}

// MbinServe starts the MiniWeb server with specified configurations.
func MbinServe(port string) {
	c := MbinInit()

	// Create a server with timeouts
	server := configs.ServerConfig()
	server.Addr = port
	server.Handler = c.Router

	appPort := helpers.Env("PORT", "8080")

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
