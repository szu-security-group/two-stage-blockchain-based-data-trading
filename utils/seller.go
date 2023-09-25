package utils

import (
	"dataTrading/aes"
	"dataTrading/hmac"
	merkleTree "dataTrading/merkle"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Divide the file in 1MB blocks and encrypt each block with AES
// Add the key to the header of each block after encryption
func EncryptFileWithMasterKey(inFile, outFile, outkey string, masterKey []byte) error {
	// Open input file
	in, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer in.Close()

	// Processes 100kB of data at a time
	buf := make([]byte, 100*1024)
	i := 1
	var totalEncryptionTime time.Duration
	for {
		// Create output file
		out, err := os.Create(outFile + strconv.Itoa(i) + ".zip")
		if err != nil {
			return err
		}
		out2, err := os.Create(outkey + strconv.Itoa(i) + ".key")
		if err != nil {
			return err
		}
		defer out.Close()
		// Reads 1MB of data from an input file
		n, err := in.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		startTime := time.Now()
		subKeyi := hmac.HMACK(masterKey, IntToBytes(i))
		// Encrypted Data
		encrypt, err := aes.Encrypt(buf[:n], subKeyi)
		endTime := time.Now()
		encryptionTime := endTime.Sub(startTime)
		totalEncryptionTime += encryptionTime
		// Save the key as well
		// Write encrypted data to the output file
		if _, err := out.Write(encrypt); err != nil {
			return err
		}
		if _, err := out2.Write(subKeyi); err != nil {
			return err
		}
		i++
	}
	fmt.Println("EncryptFileWithMasterKey:", totalEncryptionTime.Seconds())
	return nil
}

func EncryptFileWithMasterKey2(k1Map map[string]string, filedir, outEncFileDir, outkey2Dir string, masterKey []byte) {
	var totalEncryptionTime time.Duration
	for key1name, key1Filepath := range k1Map {
		plaintext, _ := ioutil.ReadFile(key1Filepath)
		i, _ := strconv.Atoi(key1name)
		startTime := time.Now()
		subKeyi := hmac.HMACK(masterKey, common.LeftPadBytes(IntToBytes(i), 32))
		endTime := time.Now()
		totalEncryptionTime += endTime.Sub(startTime)
		out2, _ := os.Create(outkey2Dir + strconv.Itoa(i) + ".key")
		out2.Write(subKeyi)
		startTime = time.Now()
		encKey, _ := aes.Encrypt(plaintext, subKeyi)
		endTime = time.Now()
		totalEncryptionTime += endTime.Sub(startTime)
		newpath := outEncFileDir + key1name + ".key.enc"
		f, _ := os.Create(newpath)
		// Write the encrypted contents to the file next
		f.Write(encKey)
	}
	fmt.Println("EncryptFileWithMasterKey2:", totalEncryptionTime.Seconds())
}

func readDir(dir string) ([][]byte, error) {
	var data [][]byte
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			data = append(data, fileData)
		}
		return nil
	})
	return data, err
}

//Make a commitment to data
func CommitData(keyMap, fileMap map[string]string) *merkleTree.MerkleTree {
	var datas [][]byte
	var totalEncryptionTime time.Duration
	for fileName, filePath := range keyMap {
		// Read the first file
		data1, _ := ioutil.ReadFile(filePath)
		data2, _ := ioutil.ReadFile(fileMap[fileName])
		data := append(data1, data2...)
		datas = append(datas, data)
	}
	startTime := time.Now()
	tree := merkleTree.NewMerkleTree(datas)
	endTime := time.Now()
	totalEncryptionTime += endTime.Sub(startTime)
	fmt.Println("CommitData:", totalEncryptionTime.Seconds())
	return tree
}

func CommitSampleKeys(enckeysDir string, challenge int) *merkleTree.MerkleTree {
	var datas [][]byte
	var files []string
	err := filepath.Walk(enckeysDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		if i == challenge {
			break
		}
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
		}
		datas = append(datas, data)
	}
	start := time.Now()
	tree := merkleTree.NewMerkleTree(datas)
	end := time.Now()
	println("CommitSampleKeys:", end.Sub(start))
	return tree
}
func Createfile2Buyer(keyMap, fileMap map[string]string, outFile string) {
	i := 0
	for fileName, filePath := range keyMap {
		// Read the first file
		data1, _ := ioutil.ReadFile(filePath)
		data2, _ := ioutil.ReadFile(fileMap[fileName])
		data := append(data1, data2...)
		out, _ := os.Create(outFile + strconv.Itoa(i) + ".zip")
		out.Write(data)
		i++
	}
}
