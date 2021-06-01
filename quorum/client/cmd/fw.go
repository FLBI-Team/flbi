package cmd

import (
	"fbi/quorum/client/core"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

// fwCmd represenfw the fw command
var fwCmd = &cobra.Command{
	Use:   "fw",
	Short: "Update Firmware",
	Run: func(cmd *cobra.Command, args []string) {

		client := newEthClient(0)
		session := newECCSession(client, 0)

		hash := core.HashBytes([]byte("firmware" + time.Now().String()))

		// sign with manufacturer's key
		// In real scanerio, manufacturer's signature already exists
		// at the time of firmware update transaction
		pKey := readKey(keyFile(1), keyFilePasswd)
		mfSig, err := crypto.Sign(hash[:], pKey)
		check(err)

		fmt.Printf("Hash: %s\n", common.ToHex(hash[:]))
		fmt.Printf("MfSig: %s\n", common.ToHex(mfSig[:]))

		startTime := time.Now()

		tx, err := session.UpdateFirmware(hash, mfSig)
		check(err)
		success, err := core.CheckTx(client, tx.Hash())
		check(err)

		latency := time.Now().Unix() - startTime.Unix()
		printTxStatus(success)
		fmt.Printf("Latency: %ds\n", latency)
	},
}

func init() {
	rootCmd.AddCommand(fwCmd)
}
