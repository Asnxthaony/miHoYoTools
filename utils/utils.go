package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"strings"

	"golang.org/x/sys/windows/registry"
)

const (
	companyName = "miHoYo"
	productName = "Genshin Impact"
)

func GetAccountDataList() string {
	k, err := registry.OpenKey(registry.CURRENT_USER, fmt.Sprintf("SOFTWARE\\%s\\%s", companyName, productName), registry.QUERY_VALUE)

	if err != nil {
		log.Fatal("[ERR] Unable to open registry key: ", err)
	}
	defer k.Close()

	s, _, err := k.GetBinaryValue("MIHOYOSDK_ADL_PROD_OVERSEA_h1158948810")

	if err != nil {
		log.Fatal("[ERR] Unable to retrieve registry key value: ", err)
	}

	return string(s)[:len(s)-1]
}

func SetAccountDataList(accountDataList string) {
	k, err := registry.OpenKey(registry.CURRENT_USER, fmt.Sprintf("SOFTWARE\\%s\\%s", companyName, productName), registry.SET_VALUE)

	if err != nil {
		log.Fatal("[ERR] Unable to open registry key: ", err)
	}
	defer k.Close()

	err = k.SetBinaryValue("MIHOYOSDK_ADL_PROD_OVERSEA_h1158948810", append([]byte(accountDataList), 0))

	if err != nil {
		log.Fatal("[ERR] Unable to set registry key value: ", err)
	}
}

func DecodeString(decryptString string) string {
	data, err := base64.StdEncoding.DecodeString(decryptString)

	if err != nil {
		log.Fatal("[ERR] Failed to decode base64 string: ", err)
	}

	desKey := []byte(GetEncodeValue())
	desIv := []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}

	block, err := des.NewCipher(desKey)

	if err != nil {
		log.Fatal("[ERR] Failed to initialize decryption key: ", err)
	}

	decryptedData := make([]byte, len(data))
	cipher.NewCBCDecrypter(block, desIv).CryptBlocks(decryptedData, data)
	decryptedData = PKCS7UnPadding(decryptedData)

	return string(decryptedData)
}

func EncodeString(encryptString string) string {
	desKey := []byte(GetEncodeValue())
	desIv := []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}

	block, err := des.NewCipher(desKey)

	if err != nil {
		log.Fatal("[ERR] Failed to initialize decryption key: ", err)
	}

	padding := PKCS7Padding([]byte(encryptString), block.BlockSize())
	encryptedData := make([]byte, len(padding))
	cipher.NewCBCEncrypter(block, desIv).CryptBlocks(encryptedData, padding)

	data := base64.StdEncoding.EncodeToString(encryptedData)

	return string(data)
}

func GetEncodeValue() string {
	var macAddr = GetMacAddress()

	if len(macAddr) < 8 {
		log.Fatal("[ERR] Unable to get encode value")
	}

	return macAddr[0:8]
}

func GetMacAddress() string {
	adapters, err := net.Interfaces()

	if err != nil {
		log.Fatal("[ERR] Unable to get network interfaces: ", err)
	}

	for _, adapter := range adapters {
		macAddr := adapter.HardwareAddr.String()

		if macAddr == "" {
			continue
		}

		if strings.HasPrefix(macAddr, "00") {
			continue
		}

		macAddr = strings.ReplaceAll(macAddr, ":", "")
		macAddr = strings.ToUpper(macAddr)

		return macAddr
	}

	return ""
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
