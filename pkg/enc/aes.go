package enc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

// Crypto padding
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// Encrypt and Decrypt
func AESEncrypt(key, plainbytes []byte) ([]byte, error) {

	if len(plainbytes)%aes.BlockSize != 0 {
		plainbytes = PKCS5Padding(plainbytes, aes.BlockSize)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	cipherbytes := make([]byte, aes.BlockSize+len(plainbytes))
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		return nil, err

	}

	copy(cipherbytes[:aes.BlockSize], iv)

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(cipherbytes[aes.BlockSize:], plainbytes)

	return cipherbytes, nil
}

func AESDecrypt(encKey, encBuf []byte) ([]byte, error) {

	var block cipher.Block

	block, err := aes.NewCipher(encKey)
	if err != nil {
		return nil, err
	}

	if len(encBuf) < aes.BlockSize {

		return nil, errors.New("problem with encrypted test blocksize")
	}
	iv := encBuf[:aes.BlockSize]
	encBuf = encBuf[aes.BlockSize:]

	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(encBuf, encBuf)
	decBuf := PKCS5Trimming(encBuf)

	return decBuf, nil
}
