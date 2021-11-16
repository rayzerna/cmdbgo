package control

import (
	"bytes"
	"cmdbgo/control/class"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	RegistryDate string `json:"registry_date"`
}

// Sighup
func Sighup(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var params map[string]string
	decoder.Decode(&params)
	result := Registry(params["name"], params["password"])
	rtn := class.RtnData{}
	if result {
		resp := rtn.OK()
		fmt.Fprintf(writer, string(resp.ToJson()))
	}
}

// Registry
func Registry(name string, password string) bool {
	u := User{Name: name}
	// Crypto
	key := "x3x2huP34AZ1eHMng44#@2I&0atrv$I6"
	encryptoPassword := AesEncrypt(u.Password, key)
	u.Password = encryptoPassword
	now := time.Now()
	u.RegistryDate = now.Format("2006-01-02 15:04:05")
	uJson, err := json.Marshal(u)
	class.CheckError(err)
	var uMap map[string]interface{}
	err = json.Unmarshal([]byte(uJson), &uMap)
	class.CheckError(err)
	result := CreateItem("users", uMap)
	if result {
		return true
	}
	return false
}

// Login
func Login(username string, password string) {

}

func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}
func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
