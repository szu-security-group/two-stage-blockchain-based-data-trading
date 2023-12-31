package utils

import (
	"dataTrading/aes"
	merkleTree "dataTrading/merkle"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

func CheckSample(key2Map, encKeyMap, encDataMap map[string]string, count int) {
	var totalEncryptionTime time.Duration
	for key2Name, filePath := range key2Map {
		key2, _ := aes.ReadKey(filePath)
		enckey1, _ := ioutil.ReadFile(encKeyMap[key2Name])
		startTime := time.Now()
		key1, _ := aes.Decrypt(enckey1, key2)
		endTime := time.Now()
		totalEncryptionTime += endTime.Sub(startTime)
		encdata, _ := ioutil.ReadFile(encDataMap[key2Name])
		startTime = time.Now()
		aes.Decrypt(encdata, key1)
		endTime = time.Now()
		totalEncryptionTime += endTime.Sub(startTime)
		//println(data)
		count--
		if count == 0 {
			fmt.Println("CheckSample:", totalEncryptionTime.Seconds())
			break
		}
	}
}
func CheckDelta1(directory string) *merkleTree.MerkleTree {
	var datas [][]byte
	files, err := ioutil.ReadDir(directory) // 读取目录中的所有文件
	if err != nil {
		fmt.Println("Error:", err)
	}
	// 遍历文件列表，读取每个文件的内容
	for _, file := range files {
		if file.IsDir() {
			continue // 忽略子目录
		}
		filePath := filepath.Join(directory, file.Name()) // 获取文件路径
		fileData, _ := ioutil.ReadFile(filePath)          // 读取文件内容
		datas = append(datas, fileData)
	}
	startTime := time.Now()
	tree := merkleTree.NewMerkleTree(datas)
	endTime := time.Now()
	fmt.Println("CheckDelta1:", endTime.Sub(startTime).Seconds())
	return tree
}

func FinalstepDecryptAllKey(encKeyMap, key2Map map[string]string) {
	var totalEncryptionTime time.Duration
	for key2Name, filePath := range key2Map {
		key2, _ := aes.ReadKey(filePath)
		enckey1, _ := ioutil.ReadFile(encKeyMap[key2Name])
		startTime := time.Now()
		aes.Decrypt(enckey1, key2)
		endTime := time.Now()
		totalEncryptionTime += endTime.Sub(startTime)
	}
	fmt.Println("FinalstepDecryptAllKey:", totalEncryptionTime.Seconds())
}
func FinalstepDecryptAllData(key2Map, encKeyMap, encDataMap map[string]string) {
	var totalEncryptionTime time.Duration
	for key2Name, filePath := range key2Map {
		key2, _ := aes.ReadKey(filePath)
		enckey1, _ := ioutil.ReadFile(encKeyMap[key2Name])
		startTime := time.Now()
		key1, _ := aes.Decrypt(enckey1, key2)
		endTime := time.Now()
		totalEncryptionTime += endTime.Sub(startTime)
		encdata, _ := ioutil.ReadFile(encDataMap[key2Name])
		startTime = time.Now()
		aes.Decrypt(encdata, key1)
		endTime = time.Now()
		totalEncryptionTime += endTime.Sub(startTime)
	}
	fmt.Println("FinalstepDecryptAllData:", totalEncryptionTime.Seconds())
}
