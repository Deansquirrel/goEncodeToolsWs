package object

import "github.com/Deansquirrel/goToolSecret"

type ObjSecret struct {
}

func (s *ObjSecret) Encrypt(plainText string, key string) (string, error) {
	return goToolSecret.EncryptToBase64Format(plainText, key)
}

func (s *ObjSecret) Decrypt(cipherText string, key string) (string, error) {
	return goToolSecret.DecryptFromBase64Format(cipherText, key)
}
