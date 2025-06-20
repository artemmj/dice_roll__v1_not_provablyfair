package grpcapp

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	drgrpc "dice_roll__v1_not_provablyfair/internal/grpc/dice_roll"
)

// Поле GRPCServer содержащаяся в DiceRollMainApp (главной структуре gRPC-приложения)
type GRPCApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func New(log *slog.Logger, gameApi drgrpc.DiceRollGameAPI, port int) *GRPCApp {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error("Recovered from panic", slog.Any("panic", p))
			return status.Errorf(codes.Internal, "internal error")
		}),
	}
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
	))
	drgrpc.Register(log, gRPCServer, gameApi)

	return &GRPCApp{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// MustRun запускает gRPC сервер и паникует если что-то идет не так
func (a *GRPCApp) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run запускает gRPC сервер
func (a *GRPCApp) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	a.log.Info("grpc server started", slog.String("addr", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// Stop отсанавливает gRPC server
func (a *GRPCApp) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))
	// Используем встроенный в gRPCServer механизм graceful shutdown
	a.gRPCServer.GracefulStop()
}
