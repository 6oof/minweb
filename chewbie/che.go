package chew

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/6oof/chewbie/app"
	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
)

type Chewbie struct {
	Router *chi.Mux
}

func ChewbieInit() *Chewbie {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	app.RegisterRoutes(r)

	Chewbie := &Chewbie{
		Router: r,
	}

	return Chewbie
}

func ChewbieServe(port string) {
	c := ChewbieInit()
	// The HTTP Server
	server := &http.Server{Addr: port, Handler: c.Router}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Create a custom logger without date and time

	art := "[̲̅$̲̅(̲̅ιο̲̅̅)̲̅$̲̅] [̲̅$̲̅(̲̅ιο̲̅̅)̲̅$̲̅] [̲̅$̲̅(̲̅ιο̲̅̅)̲̅$̲̅] [̲̅$̲̅(̲̅ιο̲̅̅)̲̅$̲̅] [̲̅$̲̅(̲̅ιο̲̅̅)̲̅$̲̅]"
	logger := log.New(os.Stdout, "", 0)

	// Use the custom logger for subsequent log messages

	logger.Print(art)
	logger.Print("")
	logger.Print("Server running on port ", port)
	logger.Print("")
	logger.Print(art)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
