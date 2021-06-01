package cmd

import (
	"fbi/eth/client/core"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// paddkeyCmd represents the paddkey command
var paddkeyCmd = &cobra.Command{
	Use:   "paddkey",
	Short: "Propose addkey",
	Run: func(cmd *cobra.Command, args []string) {

		client := newEthClient(0)
		session := newECCSession(client, 0)

		acc := common.BigToAddress(big.NewInt(rand.Int63()))
		fmt.Printf("Account: %s\n", acc.Hex())

		startTime := time.Now()

		tx, err := session.ProposeAddKey(acc, 2)
		check(err)
		success, err := core.CheckTx(client, tx.Hash())
		check(err)

		latency := time.Now().Unix() - startTime.Unix()
		printTxStatus(success)
		fmt.Printf("Latency: %ds\n", latency)
	},
}

func init() {
	rootCmd.AddCommand(paddkeyCmd)
}
