package cmd

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// queryfwCmd represenfw the queryfw command
var queryfwCmd = &cobra.Command{
	Use:   "queryfw",
	Short: "Query the firmware by hash",
	Run: func(cmd *cobra.Command, args []string) {

		client := newEthClient(0)
		session := newECCSession(client, 0)

		hash, err := cmd.Flags().GetString("hash")
		check(err)

		fw, err := session.Firmwares(common.HexToHash(hash))
		check(err)

		fmt.Printf("Time: %s\n", time.Unix(fw.UnixTime, 0).Format(timeFormat))
		fmt.Printf("Admin: %s\n", fw.Admin.Hex())
		fmt.Printf("Manufacturer: %s\n", fw.Mf.Hex())
		fmt.Printf("Manufacturer Signature: %s\n", common.ToHex(fw.MfSig))
	},
}

func init() {
	rootCmd.AddCommand(queryfwCmd)
	queryfwCmd.Flags().StringP("hash", "x", "", "Hash to query")
	queryfwCmd.MarkFlagRequired("hash")
}
