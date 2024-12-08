package app

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"menu/internal/data"
	"menu/internal/loger"
	"menu/internal/proto/pb"
	"menu/internal/server"
	"net"
	"net/http"
)

type App struct {
	lis    net.Listener
	gRPC   *grpc.Server
	Logger loger.Loger
	server server.Server
	DB     data.DB
}

func Init() App {
	logger := loger.New()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Logger.Error().Err(err)
	}
	db := data.New()

	return App{
		Logger: logger,
		server: server.Server{
			Logger: logger,
			DB:     db,
		},
		lis:  lis,
		DB:   db,
		gRPC: grpc.NewServer(),
	}
}

func (a App) Run() {
	reflection.Register(a.gRPC)
	pb.RegisterBookingServer(a.gRPC, a.server)

	go func() {
		a.Logger.Logger.Info().Msg("Starting gRPC Server on 0.0.0.0:50052")
		if err := a.gRPC.Serve(a.lis); err != nil {
			a.Logger.Logger.Error().Err(err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50052",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		a.Logger.Logger.Error().Err(err)
	}

	mx := runtime.NewServeMux()
	err = pb.RegisterBookingHandler(context.Background(), mx, conn)
	if err != nil {
		a.Logger.Logger.Error().Err(err)
	}

	s := &http.Server{
		Addr:    "0.0.0.0:50052",
		Handler: mx,
	}

	if err := s.ListenAndServe(); err != nil {
		a.Logger.Logger.Error().Err(err)
	}
}
