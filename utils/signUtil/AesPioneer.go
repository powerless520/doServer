package signUtil

import (
	"crypto/aes"
	"encoding/base64"
)

func Pioneer_AesDecryptECB(hexbase64ed_key, encrypted_str string) (data string, err error) {
	key, err := base64.StdEncoding.DecodeString(hexbase64ed_key)
	if err != nil {
		return "", err
	}
	encrypted, err := base64.StdEncoding.DecodeString(encrypted_str)
	if err != nil {
		return "", err
	}

	cipher, _ := aes.NewCipher(pioneer_generateKey(key))
	decrypted := make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return string(decrypted[:trim]), nil
}

func pioneer_generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func Pioneer_AesEncryptECB(origData_str, hexbase64ed_key string) (encrypted_str string, err error) {
	key, err := base64.StdEncoding.DecodeString(hexbase64ed_key)
	if err != nil {
		return "", err
	}
	cipher, _ := aes.NewCipher(pioneer_generateKey(key))
	origData := []byte(origData_str)
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}
