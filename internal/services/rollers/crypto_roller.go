package rollers

import (
	crypto_rand "crypto/rand"
	"math/big"
)

// Криптографический генератор
type CryptoRandRoller struct{}

func (c *CryptoRandRoller) Roll() int {
	n, _ := crypto_rand.Int(crypto_rand.Reader, big.NewInt(6))
	return int(n.Int64()) + 1
}

func (m *CryptoRandRoller) Name() string {
	return "CryptoRandRoller"
}
