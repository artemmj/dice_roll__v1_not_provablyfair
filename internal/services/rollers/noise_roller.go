package rollers

import (
	"time"
)

// Генератор на основе шума окружения
type EnvNoiseRoller struct{}

func (e *EnvNoiseRoller) Roll() int {
	ns := time.Now().Nanosecond()
	return (ns % 6) + 1
}

func (m *EnvNoiseRoller) Name() string {
	return "EnvNoiseRoller"
}
