package chaoxing

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
)

func EncryptDesECB(data, key []byte) (string, error) {
	if len(key) > 8 {
		key = key[:8]
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", errors.New(err.Error())
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return "", errors.New("EncryptDesECB need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func PKCS5Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
