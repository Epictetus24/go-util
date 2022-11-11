package enc

import "crypto/sha256"

func Sum256(msg []byte) ([]byte, error) {

	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		return nil, err
	}

	msgHashSum := msgHash.Sum(nil)
	return msgHashSum, nil

}
