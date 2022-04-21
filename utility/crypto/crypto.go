package crypto

import (
	"github.com/gogf/gf/v2/crypto/gaes"
)

var (
	aesCfbkey = []byte("f84b0fd36ba929699454a49df63a573b")
)

// AesEncrypt CBC加密，密钥必须为16/24/32位长度
func AesEncrypt(text []byte) (encryptCFB []byte, err error) {
	encryptCFB, err = gaes.EncryptCBC(text, aesCfbkey)
	if err != nil {
		return
	}
	return
}

// AesDecrypt CBC解密
func AesDecrypt(cipherText []byte) (decryptCFB []byte, err error) {
	decryptCFB, err = gaes.DecryptCBC(cipherText, aesCfbkey)
	if err != nil {
		return
	}
	return
}
