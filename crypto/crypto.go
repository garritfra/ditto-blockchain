package crypto

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
)

func CalculateHash(obj interface{}) string {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(obj); err != nil {
		panic(err)
	}
	hasher := sha256.New()
	bytes := buffer.Bytes()
	hasher.Write(bytes)

	sum := hex.EncodeToString(hasher.Sum(nil))

	return string(sum)
}
