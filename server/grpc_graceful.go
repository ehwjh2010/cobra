package server

import (
	"errors"
	cliServer "github.com/ehwjh2010/viper/client/server"
	"github.com/ehwjh2010/viper/log"
	wrapErrs "github.com/pkg/errors"
	"google.golang.org/grpc/reflection"
	"net"
)

var (
	InvalidGrpcServer = errors.New("invalid grpc server")
	InvalidGrpcConf   = errors.New("invalid grpc config")
)

// GraceGrpcServer 优雅启动grpc服务
func GraceGrpcServer(graceGrpc *cliServer.GraceGrpc) error {
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

	select {
	case <-stopChan:
		log.Info("Shutting down gracefully")
		graceGrpc.Server.GracefulStop()
		return nil

	case e := <-errChan:
		return e
	}

}
