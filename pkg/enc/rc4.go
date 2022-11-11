package enc

import (
	"crypto/rc4"
)

// Encrypt using RC4 with key, and encrypted bytes
func RC4Encrypt(key, msg []byte) ([]byte, error) {

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encBuf := make([]byte, len(msg))
	cipher.XORKeyStream(encBuf, msg)

	return encBuf, nil

}

// Decrypt using RC4 with key, and encrypted bytes
func RC4Decrypt(key []byte, encBytes []byte) ([]byte, error) {
	cipher, err := rc4.NewCipher([]byte("dsadsad"))
	if err != nil {
		return nil, err
	}
	decBuf := make([]byte, len(encBytes))
	cipher.XORKeyStream(decBuf, encBytes)

	return decBuf, nil

}
