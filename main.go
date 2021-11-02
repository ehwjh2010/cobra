package main

import (
	"context"
	"fmt"
	"ginLearn/middleware"
	"ginLearn/resource"
	"ginLearn/src/route"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	handler := gin.New()

	route.BindRoute(handler)

	middleware.UseMiddleWare(handler)

	addr := fmt.Sprintf(":%d", resource.Conf.ServerPort)

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			resourceErrs := resource.Close()
			log.Fatalf("Listen: %s\nresource: %#v", err, resourceErrs)
		}
	}()

	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		resource.Close()
		log.Fatal("Server forced to shutdown: ", err)
	}

	resource.Close()
	log.Println("Server exiting")
}
