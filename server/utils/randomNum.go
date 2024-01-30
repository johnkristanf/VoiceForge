package utils

import (
	"crypto/rand"
	"math/big"
)

func SecureRandomNumber(length int) (int64, error) {
	max := big.NewInt(10)
	max.Exp(max, big.NewInt(int64(length)), nil)
	
	num, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}

	return num.Int64(), nil
}
