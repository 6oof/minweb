package stubs

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/6oof/chewbie/app"
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

type MiniWeb struct {
	Router *mux.Router
}

func MbinInit() *MiniWeb {
	r := mux.NewRouter()

	r.Use(panicRecoverMiddleware)
	r.Use(loggingMiddleware)
	r.Use(compressResponseMiddleware)

	// CSRF protection
	csrfKey := securecookie.GenerateRandomKey(32)
	csrfMiddleware := csrf.Protect(csrfKey, csrf.Secure(false), csrf.CookieName("ccsrf")) // Set Secure to true in production
	r.Use(csrfMiddleware)

	// Add your routes as needed
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	app.RegisterRoutes(r)
	// Apply CSRF protection to all routes
	Chewbie := &MiniWeb{
		Router: r,
	}

	return Chewbie
}

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}

func panicRecoverMiddleware(next http.Handler) http.Handler {
	return handlers.RecoveryHandler()(next)
}

func compressResponseMiddleware(next http.Handler) http.Handler {
	return handlers.CompressHandler(next)
}

func MbinServe(port string) {
	c := MbinInit()

	// Create a server with timeouts
	server := &http.Server{
		Addr:         port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      c.Router,
	}

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

	log.Printf("Server running on port %s\n", port)

	// Block until an interrupt signal is received.
	<-sig
	log.Println("Shutting down...")

	// Optionally, you could run server.Shutdown in a goroutine
	// and block on <-ctx.Done() if your application should wait
	// for other services to finalize based on context cancellation.
	log.Println("Server gracefully stopped.")
}
