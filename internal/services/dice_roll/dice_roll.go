package dice_roll

import (
	"context"
	"log/slog"
	"time"

	"dice_roll__v1_not_provablyfair/internal/models"
	"dice_roll__v1_not_provablyfair/internal/services/dice"
)

// Структура сервиса с бизнес-логикой игры
type DiceRollService struct {
	log  *slog.Logger
	game DiceRollGame
}

type DiceRollGame interface {
	SaveGame(ctx context.Context, log *slog.Logger, results models.GameResult) (models.GameResult, error)
}

func New(log *slog.Logger, game DiceRollGame) *DiceRollService {
	return &DiceRollService{
		log:  log,
		game: game,
	}
}

func (dr *DiceRollService) Play(ctx context.Context) (models.GameResult, error) {
	const op = "services.dice_roll.Play"
	plog := dr.log.With(slog.String("op", op))

	dice := dice.New(dr.log)
	server := dice.Roll()
	player := dice.Roll()
	created_at := time.Now().Format(time.RFC3339)
	roller := dice.CurrentRollerName()

	winner := "draft"
	if server > player {
		winner = "server"
	} else if server < player {
		winner = "player"
	}

	game_results := models.GameResult{
		CreatedAt:  created_at,
		ServerRoll: int32(server),
		PlayerRoll: int32(player),
		Winner:     winner,
		Roller:     roller,
	}

	_, err := dr.game.SaveGame(ctx, dr.log, game_results)
	if err != nil {
		plog.Error("Ошибка при попытке сохранения игры: ", slog.Any("err", err))
	}

	return game_results, nil
}
