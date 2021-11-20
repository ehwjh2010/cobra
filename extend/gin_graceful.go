package extend

import (
	"context"
	"fmt"
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func invokeFunc(functions []func() error) *types.MultiErr {
	if functions == nil {
		return nil
	}

	var multiErr types.MultiErr

	for _, function := range functions {
		if err := function(); err != nil {
			multiErr.AddErr(err)
		}
	}

	return &multiErr
}

func GraceServer(engine *gin.Engine, serverConfig client.Server, onStartUp []func() error, onShutDown []func() error) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	host, port, timeout := serverConfig.Host, serverConfig.Port, serverConfig.ShutDownTimeout

	addr := fmt.Sprintf("%s:%d", host, port)

	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	//Invoke OnStartUp
	if multiErr := invokeFunc(onStartUp); multiErr.IsNotEmpty() {
		log.Fatalf("Invoke start function failed!, %v", multiErr.Error())
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			multiErr := invokeFunc(onShutDown)

			if multiErr.IsNotEmpty() {
				log.Fatalf("Listen: %s, resource: %s", err, multiErr.Error())
			} else {
				log.Fatalf("Listen: %s", err)
			}

		}
	}()

	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Debug("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		multiErr := invokeFunc(onShutDown)
		if multiErr.IsNotEmpty() {
			log.Fatalf("Server forced to shutdown: err: %v, resource: %s\n", err, multiErr.Error())
		} else {
			log.Fatal("Server forced to shutdown: ", err)
		}
	}

	multiErr := invokeFunc(onShutDown)
	if multiErr.IsNotEmpty() {
		log.Errorf("Server exiting, resource: %s", multiErr.Error())
	} else {
		log.Debug("Server exiting")
	}
}
