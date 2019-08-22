package cmakeys

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
)

var passphrase = "Sri Jayewardenepura Kotte"
var key = []byte{0xA9, 0xFA, 0x5A, 0x62, 0x79, 0x9F, 0xCC, 0x4C, 0x72, 0x6B, 0x4E, 0x2C, 0xE3, 0x50, 0x6D, 0x38}

func GenerateKeyStr(aid string) string {
	derivedKey := append([]byte(aid), []byte(passphrase)...)
	sha := sha256.New()
	sha.Write(derivedKey)

	return hex.EncodeToString(decrypt(sha.Sum(nil), key))
}

func decrypt(cipherData []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	decrypted := make([]byte, len(cipherData))
	if err != nil {
		panic(err.Error())
	}

	te := NewECBDecrypter(cipher)
	te.CryptBlocks(decrypted, cipherData)

	return decrypted
}
