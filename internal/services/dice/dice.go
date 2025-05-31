package dice

import (
	"log/slog"
	"math/rand"
	"time"

	"dice_roll__v1_not_provablyfair/internal/services/rollers"
)

// Интерфейс для генерации случайных чисел
type DiceRoller interface {
	Roll() int
	Name() string
}

// Игровой кубик с возможностью выбора роллера
type Dice struct {
	log           *slog.Logger
	rollers       []DiceRoller
	currentRoller DiceRoller
}

func New(log *slog.Logger) *Dice {
	d := &Dice{
		log: log,
		rollers: []DiceRoller{
			rollers.NewMathRandRoller(time.Now().UnixNano()),
			&rollers.CryptoRandRoller{},
			&rollers.EnvNoiseRoller{},
		},
	}
	d.SelectRandomRoller() // Выбрать случайный роллер при создании
	return d
}

// Выбирает случайный роллер из доступных
func (d *Dice) SelectRandomRoller() {
	if len(d.rollers) == 0 {
		panic("no rollers available")
	}
	// Используем системный роллер для выбора
	idx := rand.Intn(len(d.rollers))
	d.currentRoller = d.rollers[idx]
}

// Выполняет бросок кубика текущим роллером
func (d *Dice) Roll() int {
	return d.currentRoller.Roll()
}

// Возвращает имя текущего роллера
func (d *Dice) CurrentRollerName() string {
	return d.currentRoller.Name()
}
