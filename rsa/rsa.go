package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// GenerateRSAKeyPair 生成 RSA 秘钥对文件
func GenerateRSAKeyPair(keyLength int, privateKeyPath string, publicKeyPath string) error {
	// 生成密钥对，大小为 keyLength 比特
	privateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return err
	}

	// 把私钥编码成 PEM 格式字符串并输出到文件中
	privateKeyFile, err := os.Create(privateKeyPath)
	if err != nil {
		return err
	}
	defer func(privateKeyFile *os.File) {
		err := privateKeyFile.Close()
		if err != nil {

		}
	}(privateKeyFile)
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// 打印私钥
	fmt.Println("生成的 RSA 私钥如下：")
	fmt.Println(string(pem.EncodeToMemory(privateKeyPEM)))

	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		return err
	}

	// 把公钥编码成 PEM 格式字符串并输出到文件中
	publicKey := &privateKey.PublicKey
	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		return err
	}
	defer func(publicKeyFile *os.File) {
		err := publicKeyFile.Close()
		if err != nil {

		}
	}(publicKeyFile)
	publicKeyPEM, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicKeyPEMBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyPEM,
	}

	// 打印公钥
	fmt.Println("生成的 RSA 公钥如下：")
	fmt.Println(string(pem.EncodeToMemory(publicKeyPEMBlock)))

	err = pem.Encode(publicKeyFile, publicKeyPEMBlock)
	if err != nil {
		return err
	}

	fmt.Println("RSA 密钥对已生成成功！")
	return nil
}

//func main() {
//	keyLength := 2048                   // 设置默认的秘钥长度
//	privateKeyPath := "private_key.pem" // 设置默认的私钥文件保存路径
//	publicKeyPath := "public_key.pem"   // 设置默认的公钥文件保存路径
//
//	err := GenerateRSAKeyPair(keyLength, privateKeyPath, publicKeyPath)
//	if err != nil {
//		panic(err)
//	}
//}
