package cmd

import (
	"fbi/quorum/client/core"
	"fbi/quorum/client/ecc"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deployeccCmd represents the deployecc command
var deployeccCmd = &cobra.Command{
	Use:   "deployecc",
	Short: "deploy ECC smart contract",
	Run: func(cmd *cobra.Command, args []string) {
		client := newEthClient(0)
		conAddr, tx, _, err := ecc.DeployECC(newTransactor(0), client)
		check(err)

		fmt.Printf("Contract Address: %s\n", conAddr.Hex())

		success, err := core.CheckTx(client, tx.Hash())
		printTxStatus(success)
		if !success {
			return
		}

		f, err := os.Create(conAddrFile)
		check(err)
		f.WriteString(conAddr.Hex())
	},
}

func init() {
	rootCmd.AddCommand(deployeccCmd)
}
