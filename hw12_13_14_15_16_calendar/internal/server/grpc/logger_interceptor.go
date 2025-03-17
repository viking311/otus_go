package internalgrpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	"google.golang.org/grpc"
)

type LoggerInterceptor struct {
	loggerr app.Logger
}

func (li *LoggerInterceptor) GetInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		end := time.Since(start)

		ip := "unknown"
		peerInfo, ok := peer.FromContext(ctx)
		if ok {
			ip = peerInfo.Addr.String()
		}

		userAgent := "unknown"
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			userAgent = md.Get("user-agent")[0]
		}

		statusCode := codes.Unknown
		if st, ok := status.FromError(err); ok {
			statusCode = st.Code()
		}

		msg := fmt.Sprintf("%s %s %d %s %s",
			ip,
			info.FullMethod,
			statusCode,
			end.String(),
			userAgent)
		li.loggerr.Info(msg)

		if err != nil {
			li.loggerr.Error(err.Error())
		}

		return resp, err
	}
}

func NewLoggerInterceptor(loggerr app.Logger) *LoggerInterceptor {
	return &LoggerInterceptor{
		loggerr: loggerr,
	}
}
