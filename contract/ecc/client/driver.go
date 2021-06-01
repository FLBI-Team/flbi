package main

import (
	"crypto/ecdsa"
	"fmt"
	"time"
)

// Client ...
type Client struct {
	nThreads             int      // number of threads
	nOutstandingRequests int      // number of outstanding requests
	nRequestsPerSecond   int      // rate, invert to see the next request
	nMeters              int      // number of meters
	hosts                []string // all host to send requests
	duration             int      // how long to run, in seconds

	meterIDs []string

	privKey     *ecdsa.PrivateKey
	chaincodeID string

	pubKeyStr string
}

func (c *Client) worker(job chan int, closed chan bool, stats chan int) {
	// receive job, which is id for the next update
	hidx := 0
	nhosts := len(c.hosts)
	txIDs := make([]string, 0)
	for {
		select {
		case idx := <-job:
			{
				idx = idx % c.nMeters
				reading := readMeter(c.meterIDs[idx])
				hash := hashReading(reading)

				blockHeight, _ := getBlockHeight(c.hosts[hidx])
				hash = fmt.Sprintf("%s_%d", hash, blockHeight)

				r, s, err := signPayload(c.privKey, hash)
				check(err)

				txID := c.Verify(c.hosts[hidx], c.pubKeyStr, hash, r, s)
				txIDs = append(txIDs, txID)

				hidx = (hidx + 1) % nhosts
			}
		case <-closed:
			{
				// for _, t := range txIDs {
				// 	fmt.Println(t)
				// }
				stats <- len(txIDs)
				return
			}
		}
	}
}

func (c *Client) workerAddKey(job chan int, closed chan bool, stats chan int) {
	// receive job, which is id for the next update
	hidx := 0
	nhosts := len(c.hosts)
	txIDs := make([]string, 0)
	for {
		select {
		case idx := <-job:
			{
				proposalType, proposedKey, bhStr, hash, r, s :=
					c.genAddKeyProposal(
						c.hosts[hidx], idx,
					)

				txID := c.PutProposal(
					c.hosts[hidx],
					c.pubKeyStr, proposalType, proposedKey, bhStr, hash, r, s,
				)

				txIDs = append(txIDs, txID)

				hidx = (hidx + 1) % nhosts
			}
		case <-closed:
			{
				// for _, t := range txIDs {
				// 	fmt.Println(t)
				// }
				stats <- len(txIDs)
				return
			}
		}
	}
}

func (c *Client) workerAddWR(job chan int, closed chan bool, stats chan int) {
	// receive job, which is id for the next update
	hidx := 0
	nhosts := len(c.hosts)
	txIDs := make([]string, 0)
	for {
		select {
		case idx := <-job:
			{
				ts, reading, r, s := c.genWaterReading(idx)

				txID := c.AddWaterReading(c.hosts[0], c.pubKeyStr, ts, reading, r, s)

				txIDs = append(txIDs, txID)

				hidx = (hidx + 1) % nhosts
			}
		case <-closed:
			{
				// for _, t := range txIDs {
				// 	fmt.Println(t)
				// }
				stats <- len(txIDs)
				return
			}
		}
	}
}

// Start verifies meter readings for a period
func (c *Client) Start(txType string) {

	c.meterIDs = createMeters(c.nMeters)

	buffer := make(chan int, c.nOutstandingRequests)
	done := make(chan bool)
	stats := make(chan int) // see how many were sent
	for i := 0; i < c.nThreads; i++ {
		switch txType {
		case "addwr":
			go c.workerAddWR(buffer, done, stats)
		case "addkey":
			go c.workerAddKey(buffer, done, stats)
		default:
			go c.worker(buffer, done, stats)
		}
	}

	go measure(c.hosts[0], c.duration)

	// loop for duration
	t := time.Now()
	tEnd := t.Add(time.Duration(c.duration * 1000 * 1000 * 1000)) // in nanosecond
	idx := 0
	for time.Now().Before(tEnd) {
		if time.Now().Before(t) {
			continue
		}
		buffer <- idx
		idx++
		t = c.nextRequest()
	}

	// then close
	nSent := 0
	for i := 0; i < c.nThreads; i++ {
		done <- true
		nSent += <-stats
	}
	fmt.Printf("Total sent over %d(s): %v\n", c.duration, nSent)
}

func measure(endpoint string, duration int) {

	delay := duration * 10 / 100
	time.Sleep(time.Second * time.Duration(delay))

	offsetBlock, err := getBlockHeight(endpoint)
	check(err)
	offsetBlock--

	startTime := time.Now()

	for {

		time.Sleep(time.Second * 1)

		duration := int(time.Now().Unix() - startTime.Unix())
		fmt.Printf("\nDuration: %d\n", duration)

		currentBlock, _ := getBlockHeight(endpoint)
		currentBlock--

		blockCount := currentBlock - offsetBlock
		fmt.Printf("Block Count: %d\n", blockCount)

		txCount := getTotalTxCount(endpoint, offsetBlock, currentBlock)
		fmt.Printf("Transaction Count: %d\n", txCount)

		txRate := txCount / duration
		fmt.Printf("Transaction Rate: %d\n", txRate)

	}

}

// Latency ...
func (c *Client) Latency() {
	reading := readMeter(createMeterID())
	hash := hashReading(reading)

	endpoint := c.hosts[0]

	blockHeight, _ := getBlockHeight(endpoint)
	hash = fmt.Sprintf("%s_%d", hash, blockHeight)

	r, s, err := signPayload(c.privKey, hash)
	check(err)

	txID := c.Verify(endpoint, c.pubKeyStr, hash, r, s)

	startTime := time.Now()

	for {
		time.Sleep(time.Millisecond * 500)
		if findTx(endpoint, txID) {
			break
		}
	}

	latency := int(time.Now().Unix() - startTime.Unix())

	fmt.Printf("Latency: %d\n", latency)
}

func findTx(endpoint, txID string) bool {

	currentBlock, _ := getBlockHeight(endpoint)
	currentBlock--

	block, _ := getBlock(endpoint, currentBlock)
	for _, tx := range block.Transactions {
		if tx.TxID == txID {
			return true
		}
	}
	return false
}
