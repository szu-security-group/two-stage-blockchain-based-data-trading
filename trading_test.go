package newversion

import (
	"context"
	"crypto/ecdsa"
	"dataTrading/aes"
	merkleTree "dataTrading/merkle"
	"dataTrading/mycontract"
	"dataTrading/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"testing"
	"time"
)

var (
	sh                                                    *shell.Shell
	nodeUrl, contractAddress, sellerAddress, buyerAddress string
	sellerPrivateKey                                      *ecdsa.PrivateKey
	buyerPrivateKey                                       *ecdsa.PrivateKey
	conn                                                  *ethclient.Client
	contract                                              *mycontract.Trading
	//Data waiting to be sold
	inFile = "./encTest/10GBfile"
	//Location of processed data
	outFile = "./encTest/output/"
	//the keys dir which encrypt raw data,these key creat by masterK1
	outKey = "./encTest/keys/"
	//the enckeys dir which encrypted by key2
	outencKey = "./encTest/enckeys/"
	//the keys dir which encrypt raw keys,these key creat by masterK2
	key2sdir = "./encTest/key2s/"
	//The location where the simulated file to be sent to the DB is stored.
	file2Buyer = "./encTest/file2Buyer/"
	masterK1   = crypto.Keccak256([]byte("123"))
	masterK2   = crypto.Keccak256([]byte("321"))
	//merkle tree for data
	dataTree *merkleTree.MerkleTree
	//merkle tree for sample
	sampleTree *merkleTree.MerkleTree
)

func TestMain(m *testing.M) {
	utils.LoadConfig("./config.yaml")
	nodeUrl = utils.ConfigString("node.url")
	contractAddress = utils.ConfigString("contract.address")
	sellerAddress = utils.ConfigString("seller.address")
	sellerPrivateKey, _ = crypto.HexToECDSA(utils.ConfigString("seller.privatekey"))
	buyerAddress = utils.ConfigString("buyer.address")
	buyerPrivateKey, _ = crypto.HexToECDSA(utils.ConfigString("buyer.privatekey"))
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	dail, _ := ethclient.Dial(nodeUrl)
	conn = dail
	contract, _ = mycontract.NewTrading(common.HexToAddress(contractAddress), conn)
	// exit
	exitCode := m.Run()
	os.Exit(exitCode)
}

/**
Preparation phase, encrypting the file with the master key and generating the Merkle tree
*/
func TestPreparation(t *testing.T) {
	err := utils.EncryptFileWithMasterKey(inFile, outFile, outKey, masterK1)
	if err != nil {
		fmt.Println(err)
	}
	k1Map, _ := utils.ScanDir(outKey)
	utils.EncryptFileWithMasterKey2(k1Map, outKey, outencKey, key2sdir, masterK2)
	sampleTree = utils.CommitSampleKeys(key2sdir)
	sampleTree.PrintTree()
}

func BenchmarkEncryptFileWithMasterKey(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := utils.EncryptFileWithMasterKey(inFile, outFile, outKey, masterK1)
		if err != nil {
			fmt.Println(err)
		}
	}
}
func BenchmarkEncryptFileWithMasterKey2(b *testing.B) {
	k1Map, _ := utils.ScanDir(outKey)
	for n := 0; n < b.N; n++ {
		utils.EncryptFileWithMasterKey2(k1Map, outKey, outencKey, key2sdir, masterK2)
	}
}

/**
Send id and delta_1 to the chain
*/
func TestCommitData(t *testing.T) {
	outencKeyMap, _ := utils.ScanDir(outencKey)
	outFileMap, _ := utils.ScanDir(outFile)
	dataTree = utils.CommitData(outencKeyMap, outFileMap)
	auth := utils.NewAuth(conn, sellerAddress, sellerPrivateKey)
	var hash32 [32]byte
	copy(hash32[:], dataTree.RootNode.Hash)
	tx, err := contract.CommitData(auth, big.NewInt(1), hash32)
	if err != nil {
		log.Fatalf("Failed to update value: %v", err)
	}
	fmt.Printf("Update pending: 0x%x\n", tx.Hash())
	startTime := time.Now()
	mined, err := bind.WaitMined(context.Background(), conn, tx)
	endTime := time.Now()
	fmt.Printf("%d", mined.GasUsed)
	fmt.Printf("CommitData took %v seconds\n", endTime.Sub(startTime).Seconds())
}

func BenchmarkCommitData(b *testing.B) {
	outencKeyMap, _ := utils.ScanDir(outencKey)
	outFileMap, _ := utils.ScanDir(outFile)
	for n := 0; n < b.N; n++ {
		dataTree = utils.CommitData(outencKeyMap, outFileMap)
	}
}
func BenchmarkCommitDataContractInteraction(b *testing.B) {
	outencKeyMap, _ := utils.ScanDir(outencKey)
	outFileMap, _ := utils.ScanDir(outFile)
	dataTree = utils.CommitData(outencKeyMap, outFileMap)
	auth := utils.NewAuth(conn, sellerAddress, sellerPrivateKey)
	var hash32 [32]byte
	copy(hash32[:], dataTree.RootNode.Hash)
	for n := 0; n < b.N; n++ {
		tx, _ := contract.CommitData(auth, big.NewInt(1), hash32)
		bind.WaitMined(context.Background(), conn, tx)
	}
}

func TestPreChallenge(t *testing.T) {
	//simulation file transfer
	outencKeyMap, _ := utils.ScanDir(outencKey)
	outFileMap, _ := utils.ScanDir(outFile)
	utils.Createfile2Buyer(outencKeyMap, outFileMap, file2Buyer)
}

/**
Send kc,p,p1 to the chain
challenge,and make earnest payment
*/
func TestChallenge(t *testing.T) {
	//simulation check delta1
	dataTree = utils.CheckDelta1(file2Buyer)
	Kc := crypto.Keccak256([]byte("random"))
	auth := utils.NewAuth(conn, buyerAddress, buyerPrivateKey)
	var Kc32 [32]byte
	copy(Kc32[:], Kc)
	//earnest payment
	auth.Value = big.NewInt(1000)
	tx, err := contract.Challenge(auth, Kc32, big.NewInt(1000), big.NewInt(100))
	if err != nil {
		log.Fatalf("Failed to update value: %v", err)
	}
	fmt.Printf("Update pending: 0x%x\n", tx.Hash())
	startTime := time.Now()
	mined, err := bind.WaitMined(context.Background(), conn, tx)
	endTime := time.Now()
	fmt.Printf("%d", mined.GasUsed)
	fmt.Printf("CommitData took %v seconds\n", endTime.Sub(startTime).Seconds())
}

func BenchmarkChallengeCheckDelta1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dataTree = utils.CheckDelta1(file2Buyer)
	}
}

func BenchmarkChallenge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Kc := crypto.Keccak256([]byte("random"))
		var Kc32 [32]byte
		copy(Kc32[:], Kc)
	}
}

/**
send delta_2 to the chain
*/
func TestResponse(t *testing.T) {
	sampleTree = utils.CommitSampleKeys(outencKey)
	auth := utils.NewAuth(conn, sellerAddress, sellerPrivateKey)
	var hash32 [32]byte
	copy(hash32[:], sampleTree.RootNode.Hash)
	tx, err := contract.Response(auth, hash32, big.NewInt(10001))
	if err != nil {
		log.Fatalf("Failed to update value: %v", err)
	}
	fmt.Printf("Update pending: 0x%x\n", tx.Hash())
	startTime := time.Now()
	mined, err := bind.WaitMined(context.Background(), conn, tx)
	endTime := time.Now()
	fmt.Printf("%d", mined.GasUsed)
	fmt.Printf("CommitData took %v seconds\n", endTime.Sub(startTime).Seconds())
}
func BenchmarkCommitSampleKeys(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sampleTree = utils.CommitSampleKeys(outencKey)
	}
}
func BenchmarkResponseContractInteraction(b *testing.B) {
	sampleTree = utils.CommitSampleKeys(outencKey)
	auth := utils.NewAuth(conn, sellerAddress, sellerPrivateKey)
	var hash32 [32]byte
	copy(hash32[:], sampleTree.RootNode.Hash)
	tx, _ := contract.Response(auth, hash32, big.NewInt(10001))
	for n := 0; n < b.N; n++ {
		bind.WaitMined(context.Background(), conn, tx)
	}
}

/**
check sample data and make remainding payment
*/
func TestPayleft(t *testing.T) {
	//key2
	key2Map, _ := utils.ScanDir(key2sdir)
	//enckey1
	encKeyMap, _ := utils.ScanDir(outencKey)
	//encdata
	encDataMap, _ := utils.ScanDir(outFile)
	utils.CheckSample(key2Map, encKeyMap, encDataMap, 1530)
	auth := utils.NewAuth(conn, buyerAddress, buyerPrivateKey)
	auth.Value = big.NewInt(900)
	tx, err := contract.Payleft(auth)
	if err != nil {
		log.Fatalf("Failed to update value: %v", err)
	}
	fmt.Printf("Update pending: 0x%x\n", tx.Hash())
	startTime := time.Now()
	mined, err := bind.WaitMined(context.Background(), conn, tx)
	endTime := time.Now()
	fmt.Printf("%d", mined.GasUsed)
	fmt.Printf("CommitData took %v seconds\n", endTime.Sub(startTime).Seconds())
}

func BenchmarkCheckSample(b *testing.B) {
	//key2
	key2Map, _ := utils.ScanDir(key2sdir)
	//enckey1
	encKeyMap, _ := utils.ScanDir(outencKey)
	//encdata
	encDataMap, _ := utils.ScanDir(outFile)
	for n := 0; n < b.N; n++ {
		utils.CheckSample(key2Map, encKeyMap, encDataMap, 950)
	}
}

/**
Send id and delta_1 to the chain
*/
func TestReleaseMasterkey(t *testing.T) {
	auth := utils.NewAuth(conn, sellerAddress, sellerPrivateKey)
	var hash32 [32]byte
	copy(hash32[:], masterK2)
	tx, err := contract.ReleaseMasterkey(auth, hash32)
	if err != nil {
		log.Fatalf("Failed to update value: %v", err)
	}
	fmt.Printf("Update pending: 0x%x\n", tx.Hash())
	startTime := time.Now()
	mined, err := bind.WaitMined(context.Background(), conn, tx)
	endTime := time.Now()
	fmt.Printf("%d", mined.GasUsed)
	fmt.Printf("CommitData took %v seconds\n", endTime.Sub(startTime).Seconds())
}

func BenchmarkFinalstepDecryptAllKey(b *testing.B) {
	//key2
	key2Map, _ := utils.ScanDir(key2sdir)
	//enckey1
	encKeyMap, _ := utils.ScanDir(outencKey)
	for n := 0; n < b.N; n++ {
		for key2Name, filePath := range key2Map {
			key2, _ := aes.ReadKey(filePath)
			enckey1, _ := ioutil.ReadFile(encKeyMap[key2Name])
			aes.Decrypt(enckey1, key2)
		}
	}
}
func BenchmarkFinalstepDecryptAllData(b *testing.B) {
	//key2
	key2Map, _ := utils.ScanDir(key2sdir)
	//enckey1
	encKeyMap, _ := utils.ScanDir(outencKey)
	//encdata
	encDataMap, _ := utils.ScanDir(outFile)
	for n := 0; n < b.N; n++ {
		for key2Name, filePath := range key2Map {
			key2, _ := aes.ReadKey(filePath)
			enckey1, _ := ioutil.ReadFile(encKeyMap[key2Name])
			key1, _ := aes.Decrypt(enckey1, key2)
			encdata, _ := ioutil.ReadFile(encDataMap[key2Name])
			aes.Decrypt(encdata, key1)
		}
	}
}
