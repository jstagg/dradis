package main

// Example input:
// http://machinename:8888/apiname?input=yourinputhere
// Response will be:
// "value"
// "No record found"

// Thanks to:
// Rajeev Singh (https://www.callicoder.com/about)
// THWD (https://stackoverflow.com/a/16512668/14076872)

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	//"github.com/go-redis/redis"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome! Please hit your API like `http://localhost:8888/alias?input=SPR1011` to activate."))
}

func customerHandler(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := r.FormValue("input")
		val, err := client.Get(input).Result()
		if err == redis.Nil {
			log.Println("No record found for ", input)
			if err != nil {
				fmt.Fprintf(w, "No record found for %s", input)
				return
			}
		} else {
			log.Println("Found ", input)
			w.Write([]byte(val))
		}
	}
}

func aliasHandler(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := r.FormValue("input")
		val, err := client.Get(input).Result()
		if err == redis.Nil {
			log.Println("No record found for ", input)
			if err != nil {
				fmt.Fprintf(w, "No record found for %s", input)
				return
			}
		} else {
			log.Println("Found ", input)
			w.Write([]byte(val))
		}
	}
}

func main() {
	// Create Redis Client
	var (
		host     = getEnv("REDIS_HOST", "dradis-back")
		port     = string(getEnv("REDIS_PORT", "6379"))
		password = getEnv("REDIS_PASSWORD", "")
		//db       = getEnv("REDIS_DB", "10")
	)

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/customer", customerHandler(client))
	r.HandleFunc("/alias", aliasHandler(client))

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8888",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
