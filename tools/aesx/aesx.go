package aesx

import "github.com/golang-module/dongle"

type Cfg struct {
	Key string // 秘钥 必须 16/24/32个字符
	IV  string // 偏移量 必须16 个字符
}

func (c *Cfg) NewCipher() *dongle.Cipher {

	switch len(c.Key) {
	case 16, 24, 32:
	default:
		panic("key must be 16, 24 or 32 bytes")
	}

	switch len(c.IV) {
	case 16:
	default:
		panic("iv must be 16 bytes (ECB mode doesn't require setting iv)")
	}

	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)      // CBC、CFB、OFB、CTR、ECB
	cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
	cipher.SetKey(c.Key)            // key must be 16, 24 or 32 bytes
	cipher.SetIV(c.IV)              // iv must be 16 bytes (ECB mode doesn't require setting iv)
	return cipher
}

// EncodeToStr 从字符串中通过 aes 加密并输出原始字符串
func EncodeToStr(cipher *dongle.Cipher, data string) string {
	// 从字符串中通过 aes 加密并输出原始字符串
	return dongle.Encrypt.FromString(data).ByAes(cipher).ToRawString()
}

// EncodeToBytes 从字符串中通过 aes 加密并输出原始bytes字符
func EncodeToBytes(cipher *dongle.Cipher, data string) []byte {

	return dongle.Encrypt.FromString(data).ByAes(cipher).ToRawBytes()
}

// DecodeToStr 从原始字符串和输出字符串中通过 aes 解密
func DecodeToStr(cipher *dongle.Cipher, data string) string {

	return dongle.Decrypt.FromRawString(data).ByAes(cipher).ToString()
}

// DecodeToBytes 从原始字符串和输出bytes 通过 aes 解密
func DecodeToBytes(cipher *dongle.Cipher, data string) []byte {

	return dongle.Decrypt.FromRawString(data).ByAes(cipher).ToBytes()
}
