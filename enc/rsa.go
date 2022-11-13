package enc

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

// Generate New RSA Private key (which includes a corresponding public key.)
// Make a copy with &privkey to save the value to something like a struct, avoids Garbage Collection issues in long running apps.
func NewRSAKey() (*rsa.PrivateKey, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// Encrypt []byte with RSA publickey
func RSAEncrypt(plainMsg []byte, pubKey *rsa.PublicKey) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, plainMsg, nil)
	if err != nil {
		return nil, err
	}
	return encryptedBytes, nil

}

// Decrypt []byte with RSA privatekey
func RSADecrypt(encBytes []byte, privKey *rsa.PrivateKey) ([]byte, error) {
	decryptedBytes, err := privKey.Decrypt(nil, encBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return nil, err
	}

	return decryptedBytes, nil

}

// Sign a byte array with a privkey
func RSASign256(msg []byte, privkey *rsa.PrivateKey) ([]byte, error) {

	msgSum, err := Sum256(msg)
	if err != nil {
		return nil, err
	}

	signature, err := rsa.SignPSS(rand.Reader, privkey, crypto.SHA256, msgSum, nil)
	if err != nil {
		return nil, err
	}

	return signature, nil

}

// Verify signature ([]byte) of message ([]byte) returns true if signature is valid.
func RSAVerify256(msg, msgSig []byte, pubKey *rsa.PublicKey) (bool, error) {
	msgSum, err := Sum256(msg)
	if err != nil {
		return false, err
	}

	err = rsa.VerifyPSS(pubKey, crypto.SHA256, msgSum, msgSig, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Pem Encode Private Key file
func PemEncPrivKey(privKey *rsa.PrivateKey) []byte {

	var pemPrivateBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	pemEnc := pem.EncodeToMemory(pemPrivateBlock)

	return pemEnc

}

// Decode Private Key file
func PemDecPriv(pemEnc []byte) (*rsa.PrivateKey, error) {

	data, _ := pem.Decode(pemEnc)
	privKey, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		return nil, err
	}

	return privKey, nil

}
