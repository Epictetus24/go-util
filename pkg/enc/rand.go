package enc

import (
	"crypto/rand"
	"math/big"
)

func RandInt(min int, max int) (int, error) {
	max64 := big.NewInt(int64(max))
	outbig, err := rand.Int(rand.Reader, max64)
	if err != nil {
		return 0, err

	}

	out64 := outbig.Int64()
	return int(out64), nil
}
