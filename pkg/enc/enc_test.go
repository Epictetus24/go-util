package enc

import (
	"testing"
)

func TestKeyGen(t *testing.T) {

	_, got := NewRSAKey()
	if got != nil {
		t.Errorf("got %q, wanted nil", got)

	}

}

func TestRSAE2E(t *testing.T) {
	message := []byte("Really super secret message containing all sorts of crazy bits and bytes and things about your mum's pc.")
	want := string(message)

	privkey, err := NewRSAKey()
	if err != nil {
		t.Errorf("got %q, wanted nil", err)
	}

	encBytes, err := RSAEncrypt(message, &privkey.PublicKey)
	if err != nil {
		t.Errorf("got %q, wanted nil", err)
	}

	decMsg, err := RSADecrypt(encBytes, privkey)
	if err != nil {
		t.Errorf("got %q, wanted nil", err)
	}
	if string(decMsg) != want {
		t.Errorf("got %q, wanted %s", decMsg, want)
	}

}

func TestRSASig(t *testing.T) {
	privkey, err := NewRSAKey()
	if err != nil {
		t.Errorf("got %q, wanted nil", err)
	}

	message := []byte("Really super secret message containing all sorts of crazy bits and bytes and things about your mum's pc.")
	sig, err := RSASign256(message, privkey)
	if err != nil {
		t.Errorf("got %q, wanted nil", err)
	}
	RSAVerify256(message, sig, &privkey.PublicKey)

}

func RC4E2E(t *testing.T) {

	want := []byte("Super secret message with stuff that you wouldn't want on your facebook feed.")
	key := []byte("password123")

	encrypted, err := RC4Encrypt(key, want)
	if err != nil {
		t.Errorf("got %q, wanted nil", err)
	}

	got, err := RC4Decrypt(key, encrypted)
	if err != nil {
		t.Errorf("got %q, wanted nil", err)
	}

	if string(got) != string(want) {
		t.Errorf("got %q, wanted %s", got, want)

	}

}

func AESE2E(t *testing.T) {

	want := []byte("Your browsing history")
	key := []byte("spring2020")

	encrypted, err := AESEncrypt(key, want)
	if err != nil {
		t.Errorf("got %q, wanted nil", err)
	}

	got, err := AESDecrypt(key, encrypted)
	if err != nil {
		t.Errorf("got %q, wanted nil", err)
	}

	if string(got) != string(want) {
		t.Errorf("got %q, wanted %s", got, want)

	}

}
