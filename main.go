package main

import (
	"context"
	"flag"
	router "gta4roy/address_service/api"
	"gta4roy/address_service/log"
	"gta4roy/address_service/util"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
)

var (
	listenAddr string
)

func init() {
	// Parse log level from command line
	logLevel := util.GetProperty(util.LogLevel)
	// Calling the SetLogLevel with the command-line argument
	log.SetLogLevel(logLevel, "logs.txt")
	log.Trace.Println("Loging initialised")
	flag.StringVar(&listenAddr, "listen-addr", ":"+util.GetProperty(util.Port), "server listen address")
	flag.Parse()

}

func processRequestURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		log.Trace.Println("URL Path is ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func newWebServer() *http.Server {
	router := router.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Host", "Origin", "Connection", "Upgrade", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "X-Requested-With", "Content-type", "Authorisation", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "access-control-allow-origin", "access-control-allow-headers"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	_, err := os.Getwd()
	if err != nil {
		log.Error.Println(err)
	}
	return &http.Server{
		Addr:         listenAddr,
		Handler:      handlers.CORS(headers, methods, origins, credentials)(processRequestURL(router)),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
}

func graceFullShutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {
	sig := <-quit
	log.Trace.Println("Server is shutting down ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	router.CloseConnections()
	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Trace.Println("Could not gracefull shutdown server 5v\n", err)
	}
	close(done)
}

func main() {
	log.Trace.Println("Application started ")

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	server := newWebServer()

	go graceFullShutdown(server, quit, done)
	log.Trace.Println("Server is ready to handle request %s", listenAddr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error.Println("Could not listen on the %s : %v\n", listenAddr, err)
	}

	<-done

	log.Trace.Println("Server stopped....")

}
