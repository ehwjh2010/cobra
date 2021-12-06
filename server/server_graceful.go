package server

import (
	"context"
	"fmt"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/types"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func invokeFunc(functions []func() error) error {
	if functions == nil {
		return nil
	}

	var multiErr types.MultiErr

	for _, function := range functions {
		if function == nil {
			continue
		}

		multiErr.AddErr(function())
	}

	if multiErr.IsEmpty() {
		return nil
	}

	return &multiErr
}

func GraceServer(engine http.Handler, host string, port, timeout int, onStartUp []func() error, onShutDown []func() error) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	addr := fmt.Sprintf("%s:%d", host, port)

	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	//Invoke OnStartUp
	if err := invokeFunc(onStartUp); err != nil {
		log.Fatal("Invoke start function failed!!!", zap.Error(err))
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			multiErr := invokeFunc(onShutDown)

			if multiErr != nil {
				log.Fatal("Listen!!!", zap.Error(err), zap.String("multiErr", multiErr.Error()))
			} else {
				log.Fatal("Listen err!!!", zap.Error(err))
			}

		}
	}()

	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Info("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		multiErr := invokeFunc(onShutDown)
		if multiErr != nil {
			log.Fatal("Server forced to shutdown!!!", zap.String("serverErr", err.Error()),
				zap.String("mulErr", multiErr.Error()))
		} else {
			log.Fatal("Server forced to shutdown: ", zap.Error(err))
		}
	}

	multiErr := invokeFunc(onShutDown)
	if multiErr != nil {
		log.Error("Server exiting", zap.Error(multiErr))
	} else {
		log.Info("Server exiting")
	}
}
