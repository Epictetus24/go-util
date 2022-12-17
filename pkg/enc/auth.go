package enc

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

/*
Some crypto authentication helpers, from https://leanpub.com/gocrypto/read#leanpub-auto-aes-gcm

*/

var (
	ErrEncrypt = errors.New("secret: encryption failed")
	ErrDecrypt = errors.New("secret: decryption failed")
)

const (
	keySize = 32
)

func Sum256(msg []byte) ([]byte, error) {

	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		return nil, err
	}

	msgHashSum := msgHash.Sum(nil)
	return msgHashSum, nil

}

func GenKey(Size int) ([]byte, error) {
	key := make([]byte, Size)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func GenNonce(Size int) ([]byte, error) {

	nonce := make([]byte, Size)
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return nonce, nil
}
