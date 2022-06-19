package server

import (
	"context"
	"errors"
	"github.com/ehwjh2010/viper"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	wrapErrs "github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"

	"github.com/ehwjh2010/viper/log"
	"github.com/ehwjh2010/viper/verror"
)

var (
	InvalidGrpcServer = errors.New("invalid grpc server")
	InvalidGrpcConf   = errors.New("invalid grpc config")
)

// graceGrpcServer 优雅启动grpc服务
func graceGrpcServer(graceGrpc *GraceGrpc, errChan chan<- error) {
	log.Info(viper.SIGN + "\n" + "Viper Version: " + viper.VERSION)

	if graceGrpc == nil {
		panic(InvalidGrpcConf)
	}

	if graceGrpc.Server == nil {
		panic(InvalidGrpcServer)
	}

	if graceGrpc.RegisterReflect {
		// 注册 grpcurl 所需的 reflection 服务
		reflection.Register(graceGrpc.Server)
	}

	log.Debug("execute grpc on startup functions")
	if err := graceGrpc.ExecuteStartUp(); err != nil {
		panic(wrapErrs.Wrap(err, "on start function occur err"))
	}

	defer func() {
		log.Debug("execute grpc on shutdown functions")
		if closeErrs := graceGrpc.ExecuteShutDown(); closeErrs != nil {
			log.E(closeErrs)
		}
	}()

	lis, err := net.Listen("tcp", graceGrpc.Addr)
	if err != nil {
		panic(wrapErrs.Wrap(err, "listen addr err"))
	}

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

}

// GraceGrpcServer 优雅启动grpc服务
func GraceGrpcServer(graceGrpc *GraceGrpc) error {
	stopChan := getStopChan()
	errChan := getErrChan()

	graceGrpcServer(graceGrpc, errChan)

	select {
	case <-stopChan:
		graceGrpc.Server.GracefulStop()
		return nil
	// TODO 未区分gateway和grpc错误, 直接返回错误
	case e := <-errChan:
		return e
	}

}
