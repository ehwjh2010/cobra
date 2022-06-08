package server

import (
	"context"
	"errors"
	"github.com/ehwjh2010/viper"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	wrapErrs "github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/ehwjh2010/viper/log"
	"github.com/ehwjh2010/viper/verror"
)

var (
	InvalidGrpcServer = errors.New("invalid grpc server")
	InvalidGrpcConf   = errors.New("invalid grpc config")
)

// GraceGrpcServer 优雅启动grpc服务
func GraceGrpcServer(graceGrpc *GraceGrpc) error {
	log.Info(viper.SIGN + "\n" + "Viper Version: " + viper.VERSION)

	if graceGrpc == nil {
		return InvalidGrpcConf
	}

	if graceGrpc.Server == nil {
		return InvalidGrpcServer
	}

	if graceGrpc.RegisterReflect {
		// 注册 grpcurl 所需的 reflection 服务
		reflection.Register(graceGrpc.Server)
	}

	log.Debug("execute on startup functions")
	if err := graceGrpc.ExecuteStartUp(); err != nil {
		return wrapErrs.Wrap(err, "on start function occur err")
	}

	defer func() {
		log.Debug("execute on shutdown functions")
		if closeErrs := graceGrpc.ExecuteShutDown(); closeErrs != nil {
			log.E(closeErrs)
		}
	}()

	lis, err := net.Listen("tcp", graceGrpc.Addr)
	if err != nil {
		//return errors.
		return wrapErrs.Wrap(err, "listen addr err")
	}

	stopChan := getStopChan()
	errChan := getErrChan()

	go func() {
		if err := graceGrpc.Server.Serve(lis); err != nil {
			errChan <- err
		}
	}()

	var (
		gatewayServer *http.Server
		gatewayFlag   bool
	)
	if graceGrpc.EnableGateway {
		go func() {
			ct := context.Background()
			ct, cancel := context.WithCancel(ct)
			defer cancel()
			mux := runtime.NewServeMux()
			options := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
			errs := verror.MultiErr{}
			for _, register := range graceGrpc.HttpHandlers {
				errs.AddErr(register(ct, mux, graceGrpc.Addr, options))
			}
			if err := errs.AsStdErr(); err != nil {
				errChan <- err
				return
			}

			gatewayServer = &http.Server{Addr: graceGrpc.GatewayAddr, Handler: mux}
			gatewayFlag = true
			log.Debug("start gateway server")
			errChan <- gatewayServer.ListenAndServe()
		}()
	}

	select {
	case <-stopChan:
		var err error
		log.Info("start shutting down gracefully")
		if gatewayFlag {
			ctx, cancel := context.WithTimeout(context.Background(), graceGrpc.GatewayWaitTime)
			defer cancel()
			err = gatewayServer.Shutdown(ctx)
			if err == nil {
				log.Debug("shutting down gateway success")
			} else {
				log.Debug("shutting down gateway failed", zap.Error(err))
			}
		}
		graceGrpc.Server.GracefulStop()
		return err
	// TODO 未区分gateway和grpc错误, 直接返回错误
	case e := <-errChan:
		return e
	}

}
