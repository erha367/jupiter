package library

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//解密courseId,lessonId
func DecryptDES(src string) (int64, error) {
	keyByte := doDesKey()
	data, err := hex.DecodeString(src)
	if err != nil {
		return 0, err
	}
	block, err := des.NewTripleDESCipher(keyByte)
	if err != nil {
		return 0, err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		err = errors.New("crypto/cipher: input not full blocks")
		return 0, err
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = pkcs5unPadding(out)
	last, err := strconv.ParseUint(string(out), 36, 64)
	if err != nil {
		return 0, err
	}
	return int64(last), nil
}

//加密courseId,lessonId
func EncryptDES(num64 int64) (string, error) {
	//int64 10 to 36
	oriData32 := strconv.FormatInt(num64, 36)
	data := []byte(oriData32)
	keyByte := doDesKey()
	block, err := des.NewTripleDESCipher(keyByte)
	if err != nil {
		return ``, err
	}
	bs := block.BlockSize()
	//对明文数据进行补码
	data = pkcs5Padding(data, bs)
	if len(data)%bs != 0 {
		err = errors.New("need a multiple of the blocksize")
		return ``, err
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%x", out), nil
}

/**
减码操作
*/
func pkcs5unPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/**
加码操作
*/
func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

/**
加盐key补全
*/
func doDesKey() []byte {
	var key = "eeocn"
	key = key + strings.Repeat("0", 19)
	keyByte := []byte(key)
	return keyByte
}
