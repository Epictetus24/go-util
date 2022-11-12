package enc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
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
func AESCBCEncrypt(key, plainbytes []byte) ([]byte, error) {

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

// Decrypts using CBC Cipher
func AESCBCDecrypt(encKey, encBuf []byte) ([]byte, error) {

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

// GCM Methods

// EncryptWithID secures a message and prepends a sender ID.
// ID Length is calculated from byte array size. So use something consistent in length
func AESGCMEncID(key, message []byte, senderid []byte) ([]byte, error) {

	idlen := len(senderid)
	fmt.Println(idlen)

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce, err := GenNonce(12)
	if err != nil {
		return nil, err
	}

	senderid = append(senderid, nonce[:]...)
	return gcm.Seal(senderid, nonce[:], message, senderid[:idlen]), nil

}

// Decrypts the message using the key for x id.
// Will strip the ID out of the message, but you must provide the id length so it cuts in the right place.
func AESGCMDecID(senderkey, message []byte, idlen int) ([]byte, error) {

	if len(message) <= 12+idlen {
		return nil, ErrDecrypt
	}

	c, err := aes.NewCipher(senderkey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, 12)
	copy(nonce, message[idlen:])

	// Decrypt the message, using the sender ID as the additional
	// data requiring authentication.
	out, err := gcm.Open(nil, nonce, message[idlen+12:], message[:idlen])
	if err != nil {
		return nil, err
	}
	return out, nil
}
