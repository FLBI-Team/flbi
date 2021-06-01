package cmd

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// querytsCmd represents the queryts command
var querytsCmd = &cobra.Command{
	Use:   "queryts",
	Short: "Query the previously timestamped hash",
	Run: func(cmd *cobra.Command, args []string) {

		client := newEthClient(0)
		session := newECCSession(client, 0)

		hash, err := cmd.Flags().GetString("hash")
		check(err)

		ts, err := session.Timestamps(common.HexToHash(hash))
		check(err)

		fmt.Printf("Time: %s\n", time.Unix(ts.UnixTime, 0).Format(timeFormat))
		fmt.Printf("Client: %s\n", ts.Device.Hex())
	},
}

func init() {
	rootCmd.AddCommand(querytsCmd)
	querytsCmd.Flags().StringP("hash", "x", "", "Hash to query")
	querytsCmd.MarkFlagRequired("hash")
}
