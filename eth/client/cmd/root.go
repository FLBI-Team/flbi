package cmd

import (
	"crypto/ecdsa"
	"fbi/eth/client/core"
	"fbi/eth/client/ecc"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "./client",
	Short: "ECC client for Ethereum",
}

const (
	hostFile      = "hosts"
	conAddrFile   = "conAddr.txt"
	keyFilePasswd = "password"

	timeFormat = "02 Jan 2006 03:04:05 PM"
)

var hosts []string

func keyFile(idx int) string {
	return fmt.Sprintf("keys/clientkey%03d", idx)
}

// Execute ...
func Execute() error {
	hoststr, err := ioutil.ReadFile(hostFile)
	check(err)
	hosts = strings.Split(string(hoststr), "\n")
	return rootCmd.Execute()
}

func newEthClient(hostIdx int) *ethclient.Client {
	c, err := ethclient.Dial(fmt.Sprintf("http://%s:8545", hosts[hostIdx]))
	check(err)
	return c
}

func newTransactor(keyIdx int) *bind.TransactOpts {
	auth, err := core.NewTransactor(keyFile(keyIdx), keyFilePasswd)
	check(err)
	auth.GasLimit = 1000000000000
	return auth
}

func newECCSession(
	client *ethclient.Client, keyIndex int,
) *ecc.ECCSession {
	conAddr, err := ioutil.ReadFile(conAddrFile)
	check(err)
	s, err := ecc.NewECCSession(client, newTransactor(keyIndex), string(conAddr))
	check(err)
	return s
}

func printTxStatus(success bool) {
	if success {
		fmt.Println("Transaction successful")
	} else {
		fmt.Println("Transaction failed")
	}
}

func readKey(kFile, passphrase string) *ecdsa.PrivateKey {
	json, err := ioutil.ReadFile(kFile)
	check(err)
	key, err := keystore.DecryptKey(json, passphrase)
	check(err)
	return key.PrivateKey
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
