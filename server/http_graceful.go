package server

import (
	"context"
	"errors"
	"github.com/ehwjh2010/viper"
	"github.com/ehwjh2010/viper/constant"
	"net/http"
	"os"
	"os/signal"
	"time"

	wrapErrs "github.com/pkg/errors"

	"github.com/ehwjh2010/viper/log"
)

var (
	InvalidHttpConf   = errors.New("invalid http config")
	InvalidHttpEngine = errors.New("invalid http engine")
)

func getStopChan() chan os.Signal {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, constant.ListenSignals...)
	return stopChan
}

func getErrChan() chan error {
	errChan := make(chan error)
	return errChan
}

func GraceHttpServer(graceHttp *GraceHttp) error {
	log.Info(viper.SIGN + "\n" + "Viper Version: " + viper.VERSION)

	if graceHttp == nil {
		return InvalidHttpConf
	}

	if graceHttp.Engine == nil {
		return InvalidHttpEngine
	}

	//Invoke OnStartUp
	log.Debug("execute on startup functions")
	if err := graceHttp.ExecuteStartUp(); err != nil {
		return wrapErrs.Wrap(err, "on start function occur err")
	}

	defer func() {
		log.Debug("execute on shutdown functions")
		if closeErrs := graceHttp.ExecuteStartUp(); closeErrs != nil {
			log.E(closeErrs)
		}
	}()

	srv := &http.Server{
		Addr:    graceHttp.Addr,
		Handler: graceHttp.Engine,
	}

	stopChan := getStopChan()
	errChan := getErrChan()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	select {
	case <-stopChan:
		log.Info("Shutting down gracefully")
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(graceHttp.WaitSecond)*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			return wrapErrs.Wrap(err, "stop server failed!!!")
		}
		return nil
	case e := <-errChan:
		return e
	}
}
