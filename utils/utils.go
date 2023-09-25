package utils

import (
	bytes2 "bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

// NewAuth Unified create authorized transactor function
func NewAuth(conn *ethclient.Client, address string, PrivateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	nonce, _ := conn.NonceAt(context.Background(), common.HexToAddress(address), nil)
	gasPrice, _ := conn.SuggestGasPrice(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth.GasLimit = uint64(300000)
	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.GasPrice = gasPrice
	return auth
}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes2.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func LoadConfig(path string) {
	log.Println("Loading config from ", path)
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func ConfigString(key string) (val string) {
	return viper.GetString(key)
}

func ConfigInt(key string) (val int) {
	return viper.GetInt(key)
}

func ConfigFloat(key string) (val float64) {
	return viper.GetFloat64(key)
}

func ConfigBool(key string) (val bool) {
	return viper.GetBool(key)
}

func ConfigSlice(key string) []string {
	return viper.GetStringSlice(key)
}
func ConfigStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}
func ScanDir(dirPath string) (map[string]string, error) {
	// 初始化 map 存储文件信息
	fileInfoMap := make(map[string]string)
	// 遍历目录
	err := filepath.Walk(dirPath, func(filePath string, file os.FileInfo, err error) error {
		// 如果当前路径是目录，则跳过该路径
		if file.IsDir() {
			return nil
		}

		// 获取文件名
		fileName := strings.Split(file.Name(), ".")[0]
		// 将文件名和路径存储在 map 中
		fileInfoMap[fileName] = filePath

		return nil
	})

	// 返回 map 和可能出现的错误信息
	return fileInfoMap, err
}
