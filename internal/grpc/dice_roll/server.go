package server

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	diceroll_genv1 "dice_roll__v1_not_provablyfair/gen/go/dice_roll"
	"dice_roll__v1_not_provablyfair/internal/models"
)

// Структура которая будет реализовывать функционал API
type serverAPI struct {
	diceroll_genv1.UnimplementedDiceRollGameAPIServer // помогает обеспечить обратную совместимость
	log                                               *slog.Logger
	diceRollApi                                       DiceRollGameAPI
}

// Каркасы RPC-методов которые будем вызвать (пока только Play)
type DiceRollGameAPI interface {
	Play(ctx context.Context) (models.GameResult, error)
}

// Функция регистрирует эту serverAPI в gRPC-сервере
func Register(log *slog.Logger, gRPCServer *grpc.Server, diceRollApi DiceRollGameAPI) {
	// Сгенерированная функция регистрации
	diceroll_genv1.RegisterDiceRollGameAPIServer(gRPCServer, &serverAPI{diceRollApi: diceRollApi, log: log})
}

// Хэндлер (обработчик) запроса
func (s *serverAPI) Play(
	ctx context.Context,
	// Сгенерированы так же объекты запросов / ответов
	in *diceroll_genv1.PlayRequest,
) (*diceroll_genv1.PlayResponse, error) {
	const op = "grpc.server.Play"
	log := s.log.With(slog.String("op", op))
	log.Debug("IN")

	// Здесь валидация входных данных
	result, err := s.diceRollApi.Play(ctx)
	if err != nil {
		return nil, fmt.Errorf("ERROR IN GRPS SERVER; TRIED DiceRollGameAPI.Play(): %v", err)
	}
	// И формирование и возврат выходных данных
	return &diceroll_genv1.PlayResponse{
		CreatedAt:  result.CreatedAt,
		ServerRoll: result.ServerRoll,
		PlayerRoll: result.PlayerRoll,
		Winner:     result.Winner,
		Roller:     result.Roller,
	}, nil
}
