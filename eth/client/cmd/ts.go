package cmd

import (
	"fbi/eth/client/core"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// tsCmd represents the ts command
var tsCmd = &cobra.Command{
	Use:   "ts",
	Short: "Timestamp hash of the meter reading",
	Run: func(cmd *cobra.Command, args []string) {

		client := newEthClient(0)
		session := newECCSession(client, 0)

		hash := core.HashReading(core.ReadMeter(core.CreateMeterID()))
		fmt.Printf("Hash: %s\n", common.Bytes2Hex(hash[:]))

		startTime := time.Now()

		tx, err := session.TimestampHash(hash)
		check(err)
		success, err := core.CheckTx(client, tx.Hash())
		check(err)

		latency := time.Now().Unix() - startTime.Unix()
		printTxStatus(success)
		fmt.Printf("Latency: %ds\n", latency)
	},
}

func init() {
	rootCmd.AddCommand(tsCmd)
}
