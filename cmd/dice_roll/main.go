package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	dice_roll_main_app "dice_roll__v1_not_provablyfair/internal/app"
	"dice_roll__v1_not_provablyfair/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()    // грузим конфиг
	log := setupLogger(cfg.Env) // грузим логгер

	main_app := dice_roll_main_app.New(
		log,
		cfg.GRPC.Port,
		// Немного отличается строка в докере TODO сделать нормально
		// Чтобы запусть не в докере, надо поменять на cfg.postgresConnStr
		cfg.PostgresConnStrForDocker,
	)
	// Создаем главный DiceRollMainApp и асинхронно отправляем его GRPCServer в MustRun()
	go func() {
		main_app.GRPCServer.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	// Ожидаем SIGINT (pkill -2) или SIGTERM
	<-stop
	// Инициировать graceful shutdown
	main_app.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}

// Функция выбирает логгер в зависимости от окружения
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
