package rollers

import "math/rand"

// Генератор на math/rand
type MathRandRoller struct {
	r rand.Rand
}

func NewMathRandRoller(seed int64) *MathRandRoller {
	src := rand.NewSource(seed)
	return &MathRandRoller{r: *rand.New(src)}
}

func (m *MathRandRoller) Roll() int {
	return m.r.Intn(6) + 1
}

func (m *MathRandRoller) Name() string {
	return "MathRandRoller"
}
