package cmd

import (
	"fbi/eth/client/core"
	"fbi/eth/client/ecc"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// benchtsCmd represents the benchts command
var benchtsCmd = &cobra.Command{
	Use:   "benchts",
	Short: "Benchmark timestamping hash",
	Run: func(cmd *cobra.Command, args []string) {

		nThread, err := cmd.Flags().GetInt("nThread")
		check(err)
		fmt.Printf("nThread %d\n", nThread)

		reqPerSec, err := cmd.Flags().GetInt("reqPerSec")
		check(err)
		fmt.Printf("reqPerSec %d\n", reqPerSec)

		duration, err := cmd.Flags().GetInt("duration")
		check(err)
		fmt.Printf("duration %d\n", duration)

		keyCount, err := cmd.Flags().GetInt("keyCount")
		check(err)
		fmt.Printf("keyCount %d\n", keyCount)

		nodeCount, err := cmd.Flags().GetInt("nodeCount")
		check(err)
		fmt.Printf("nodeCount %d\n", nodeCount)

		clients := make([]*ethclient.Client, nodeCount)

		for i := 0; i < nodeCount; i++ {
			clients[i] = newEthClient(i)
		}

		slist := make([]*ecc.ECCSession, keyCount)
		for i := 0; i < keyCount; i++ {
			slist[i] = newECCSession(clients[i%nodeCount], i)
		}

		b := core.NewBenchmark(
			func(idx int) (*types.Transaction, error) {

				session := slist[idx%keyCount]
				hash := core.HashReading(core.ReadMeter(core.CreateMeterID()))
				return session.TimestampHash(hash)

			},
			nThread, reqPerSec, duration, clients[0],
		)

		nSent, rps, msm := b.Run()
		fmt.Printf("Total sent: %d\n", nSent)
		fmt.Printf("Actual Requests per second: %d\n", rps)
		fmt.Printf("Measurement Duration: %ds\n", msm.Duration)
		fmt.Printf("Transaction Count: %d\n", msm.TxCount)
		fmt.Printf("Throughput: %d tx/s\n", msm.TxPerSec)
		// csvStr, _ := gocsv.MarshalString(msm.Blocks)
		// fmt.Println(csvStr)
	},
}

func init() {
	rootCmd.AddCommand(benchtsCmd)
	benchtsCmd.Flags().IntP("nThread", "t", 50, "Number of threads")
	benchtsCmd.Flags().IntP("reqPerSec", "r", 50, "Request per second")
	benchtsCmd.Flags().IntP("duration", "d", 120, "Duration in seconds")
	benchtsCmd.Flags().IntP("keyCount", "k", 50, "Number of client key files")
	benchtsCmd.Flags().IntP("nodeCount", "n", 1, "Number of blockchain nodes")
}
