package rpc

import (
	"context"
	"github.com/coschain/contentos-go/common/logging"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"strings"
)

func streamRecoveryLoggingInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {

	defer func() {
		if r := recover(); r != nil {
			logging.CLog().Errorf("stream recovery interceptor err: [%v]", r)
			err = ErrPanicResp
		}
	}()

	logging.CLog().WithFields(logrus.Fields{
		"method": info.FullMethod,
	}).Info("Rpc request.")

	return handler(srv, ss)
}

func unaryRecoveryLoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	defer func() {
		if r := recover(); r != nil {
			logging.CLog().Errorf("unary recovery interceptor err: [%v]", r)
			err = ErrPanicResp
		}
	}()

	if strings.Contains(info.FullMethod, "ApiService") {
		logging.CLog().WithFields(logrus.Fields{
			"method": info.FullMethod,
			"params": req,
		}).Info("Rpc request.")
	} else {
		logging.CLog().WithFields(logrus.Fields{
			"method": info.FullMethod,
		}).Info("Rpc request.")
	}

	return handler(ctx, req)
}
