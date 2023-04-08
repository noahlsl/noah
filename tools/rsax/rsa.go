package rsax

import "github.com/golang-module/dongle"

// Encode Rsa加密,公钥私钥自适应
func Encode(key, in string) string {

	return dongle.Encrypt.FromString(in).ByRsa(key).ToRawString()
}

// Decode Rsa解密,公钥私钥自适应
func Decode(key, in string) string {

	return dongle.Decrypt.FromRawString(in).ByRsa(key).ToString() // hello world
}
