package app

import (
	"log/slog"

	grpcapp "dice_roll__v1_not_provablyfair/internal/app/grpc"
	"dice_roll__v1_not_provablyfair/internal/services/dice_roll"
	"dice_roll__v1_not_provablyfair/internal/storage/postgres"
)

// Главная структура gRPC-приложения, объявляется в main
type DiceRollMainApp struct {
	GRPCServer *grpcapp.GRPCApp
}

// Запуск приложения GRPCServer, возврат в DiceRollMainApp
func New(log *slog.Logger, grpcPort int, storageConnStr string) *DiceRollMainApp {
	// Подключается storage
	storage, err := postgres.New(storageConnStr)
	if err != nil {
		panic(err)
	}
	// Создаем сервис DiceRollService с бизнес-логикой
	diceRollService := dice_roll.New(log, storage)
	// Создаем GRPCApp главной структуры приложения, передав сервис бл
	grpcApp := grpcapp.New(log, diceRollService, grpcPort)
	// Возвращаем главную структуру в main
	return &DiceRollMainApp{
		GRPCServer: grpcApp,
	}
}
