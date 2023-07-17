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

	// Processes 1MB of data at a time
	buf := make([]byte, 1024*1024)
	i := 1
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
		subKeyi := hmac.HMACK(masterKey, IntToBytes(i))
		// Encrypted Data
		encrypt, err := aes.Encrypt(buf[:n], subKeyi)
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

	return nil
}
func EncryptFileWithMasterKey2(k1Map map[string]string, filedir, outEncFileDir, outkey2Dir string, masterKey []byte) {

	for key1name, key1Filepath := range k1Map {
		plaintext, _ := ioutil.ReadFile(key1Filepath)
		i, _ := strconv.Atoi(key1name)
		subKeyi := hmac.HMACK(masterKey, common.LeftPadBytes(IntToBytes(i), 32))
		out2, _ := os.Create(outkey2Dir + strconv.Itoa(i) + ".key")
		out2.Write(subKeyi)
		encKey, _ := aes.Encrypt(plaintext, subKeyi)
		newpath := outEncFileDir + key1name + ".key.enc"
		f, _ := os.Create(newpath)
		// Write the encrypted contents to the file next
		f.Write(encKey)
	}
	//// Directory containing the files to be encrypted
	//
	//// Loop through all files in the directory
	//i := 1
	//err := filepath.Walk(filedir, func(path string, info os.FileInfo, err error) error {
	//	// Only process regular files (not directories or special files)
	//	if !info.Mode().IsRegular() {
	//		return nil
	//	}
	//
	//	// Read the file contents into memory
	//	plaintext, err := ioutil.ReadFile(path)
	//	if err != nil {
	//		return err
	//	}
	//	//Requires both inputs to be 32-bit
	//	subKeyi := hmac.HMACK(masterKey, common.LeftPadBytes(IntToBytes(i), 32))
	//	out2, err := os.Create(outkey2Dir + strconv.Itoa(i) + ".key")
	//	if _, err := out2.Write(subKeyi); err != nil {
	//		return err
	//	}
	//	// Encrypted Data
	//	encrypt, err := aes.Encrypt(plaintext, subKeyi)
	//	// Write the encrypted contents to a new file with a .enc extension
	//	_, fileName := filepath.Split(path)
	//	newpath := outEncFileDir + fileName + ".enc"
	//	f, err := os.Create(newpath)
	//	if err != nil {
	//		return err
	//	}
	//	defer f.Close()
	//	// Write the encrypted contents to the file next
	//	_, err = f.Write(encrypt)
	//	if err != nil {
	//		return err
	//	}
	//	i++
	//	return nil
	//})
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(1)
	//}
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

	for fileName, filePath := range keyMap {
		// Read the first file
		data1, _ := ioutil.ReadFile(filePath)
		data2, _ := ioutil.ReadFile(fileMap[fileName])
		data := append(data1, data2...)
		datas = append(datas, data)
	}
	//filepath.Walk(enckeysDir, func(path string, info os.FileInfo, err error) error {
	//	if err != nil {
	//		fmt.Println(err)
	//		return err
	//	}
	//	// If the current file is a directory, skip
	//	if info.IsDir() {
	//		return nil
	//	}
	//
	//	// Get file name
	//	filename := info.Name()
	//	// Read the first file
	//	data1, err := ioutil.ReadFile(path)
	//	if err != nil {
	//		return err
	//	}
	//
	//	filename = strings.Split(filename, ".")[0]
	//	// Find files with the same name in the B folder
	//	filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
	//		if err != nil {
	//			fmt.Println(err)
	//			return err
	//		}
	//
	//		// If the current file is a directory, skip
	//		if info.IsDir() {
	//			return nil
	//		}
	//
	//		// If a file with the same name is found, the two files are concatenated
	//		if strings.Split(info.Name(), ".")[0] == filename {
	//
	//			// Read the second file
	//			data2, err := ioutil.ReadFile(path)
	//			if err != nil {
	//				return err
	//			}
	//			// Splice the contents of two files
	//			data := append(data1, data2...)
	//			datas = append(datas, data)
	//		}
	//		return nil
	//	})
	//	return nil
	//})

	tree := merkleTree.NewMerkleTree(datas)
	//tree.PrintTree()
	//MekleTree commit
	return tree
}
func CommitSampleKeys(enckeysDir string) *merkleTree.MerkleTree {
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
		if i == 950 {
			break
		}
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
		}
		datas = append(datas, data)
	}
	return merkleTree.NewMerkleTree(datas)
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
