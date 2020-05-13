package grpc

import (
	"github.com/Jurassic-Park/go-util/auth"
	"github.com/Jurassic-Park/go-util/recovery"
	"github.com/Jurassic-Park/go-util/zap"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func GetInterceptors() (grpc.ServerOption, grpc.ServerOption) {
	streamInterceptors := grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_opentracing.StreamServerInterceptor(),
		grpc_prometheus.StreamServerInterceptor,
		grpc_recovery.StreamServerInterceptor(),
		grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
		grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
		grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),
	))
	unaryInterceptors := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_opentracing.UnaryServerInterceptor(),
		grpc_prometheus.UnaryServerInterceptor,
		grpc_recovery.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
		grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
		grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
	))

	return streamInterceptors, unaryInterceptors
}
