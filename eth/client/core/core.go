package core

import (
	"context"
	"errors"
	"log"
	"math/big"
	"os"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewTransactor creates bind.TransactOpts to sign tx
func NewTransactor(keyFile, keyFilePasswd string) (*bind.TransactOpts, error) {
	f, err := os.Open(keyFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return bind.NewTransactor(f, keyFilePasswd)
}

// CheckTx checks the status of tx
// it loops until it found the tx on the blockchain
func CheckTx(client *ethclient.Client, hash common.Hash) (bool, error) {
	for {
		r, err := client.TransactionReceipt(context.Background(), hash)
		if err == ethereum.NotFound {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if err != nil {
			return false, err
		}
		return r.Status == types.ReceiptStatusSuccessful, nil
	}
}

// SendTxFunc is used to send tx by Benchmark
type SendTxFunc func(idx int) (*types.Transaction, error)

// Benchmark is used to benchmark a specific transaction
type Benchmark struct {
	client *ethclient.Client

	sendTx SendTxFunc

	nThread   int
	reqPerSec int
	duration  int // duration in second

	bufLen   int
	interval time.Duration
}

// NewBenchmark gives a new Benchmark instanse
func NewBenchmark(
	sendTx SendTxFunc, nThread, reqPerSec, duration int,
	client *ethclient.Client,
) *Benchmark {

	return &Benchmark{
		client: client,

		sendTx:    sendTx,
		nThread:   nThread,
		reqPerSec: reqPerSec,
		duration:  duration,

		bufLen:   2 * nThread,
		interval: time.Second / time.Duration(reqPerSec),
	}
}

// Run benchmark and gives nSent and measurement
func (b *Benchmark) Run() (int, int, *Measurement) {
	buffer := make(chan int, b.bufLen)
	done := make(chan bool)
	stats := make(chan int)
	for i := 0; i < b.nThread; i++ {
		go b.worker(buffer, done, stats)
	}

	msmRes := make(chan *Measurement)
	go b.measure(msmRes)

	// loop for duration
	t := time.Now()
	tEnd := t.Add(time.Duration(b.duration) * time.Second)
	idx := 0
	for time.Now().Before(tEnd) {
		if time.Now().Before(t) {
			time.Sleep(time.Microsecond * 500)
			continue
		}
		buffer <- idx
		idx++
		t = t.Add(b.interval)
	}

	nSent := 0
	for i := 0; i < b.nThread; i++ {
		done <- true
		nSent += <-stats
	}

	msm := <-msmRes

	rps := nSent / b.duration
	return nSent, rps, msm
}

func (b *Benchmark) worker(job chan int, closed chan bool, stats chan int) {
	count := 0
	for {
		select {
		case idx := <-job:
			{
				_, err := b.sendTx(idx)
				if err == nil {
					count++
				} else {
					log.Println(err)
				}
			}
		case <-closed:
			{
				stats <- count
				return
			}
		}
	}
}

// BlockSummary represents summary for single block
type BlockSummary struct {
	Number  int            `csv:"blk_no"`
	Sealer  common.Address `csv:"sealer"`
	Time    int            `csv:"time"`
	TxCount int            `csv:"tx_count"`
}

// Measurement represents total txCount and other matrices
type Measurement struct {
	Blocks []BlockSummary

	TxCount  int
	TxPerSec int
	Duration int
}

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header) (common.Address, error) {

	// Retrieve the signature from the header extra-data
	if len(header.Extra) < crypto.SignatureLength {
		return common.Address{}, errors.New("missing sig")
	}
	signature := header.Extra[len(header.Extra)-crypto.SignatureLength:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(clique.SealHash(header).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	return signer, nil
}

func (b *Benchmark) measure(
	result chan *Measurement,
) {

	warmup := b.duration * 20 / 100 // 10 percent of duration
	time.Sleep(time.Second * time.Duration(warmup))

	// headers := make(chan *types.Header)
	// stop := make(chan bool)
	// defer func() {
	// 	stop <- true
	// }()
	// SubscribeNewHead(b.client, headers, stop)

	var msm Measurement
	msm.Duration = b.duration - warmup

	// allocating slice assuming 1 block per second
	msm.Blocks = make([]BlockSummary, 0, msm.Duration)

	startBlk, err := b.client.HeaderByNumber(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	time.Sleep(time.Duration(msm.Duration) * time.Second)

	endBlk, err := b.client.HeaderByNumber(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	for i := startBlk.Number.Int64() + 1; i < endBlk.Number.Int64(); i++ {
		header, _ := b.client.HeaderByNumber(context.Background(), big.NewInt(i))

		txCount, err := b.client.TransactionCount(
			context.Background(), header.Hash(),
		)
		if err != nil {
			panic(err)
		}

		// sealer, _ := ecrecover(header)
		bs := BlockSummary{
			Number: int(header.Number.Int64()),
			// Sealer:  sealer,
			// Time:    int(header.Time),
			TxCount: int(txCount),
		}
		msm.Blocks = append(msm.Blocks, bs)
		msm.TxCount += bs.TxCount
	}

	msm.TxPerSec = msm.TxCount / msm.Duration
	result <- &msm

	// // tEnd := time.Now().Add(time.Duration(b.duration) * time.Second)

	// for {
	// 	select {
	// 	case header := <-headers:
	// 		txCount, err := b.client.TransactionCount(
	// 			context.Background(), header.Hash(),
	// 		)
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		sealer, _ := ecrecover(header)
	// 		bs := BlockSummary{
	// 			Number:  int(header.Number.Int64()),
	// 			Sealer:  sealer,
	// 			Time:    int(header.Time),
	// 			TxCount: int(txCount),
	// 		}
	// 		msm.Blocks = append(msm.Blocks, bs)
	// 		msm.TxCount += bs.TxCount

	// 	default:
	// 		if time.Now().After(tEnd) {
	// 			msm.TxPerSec = msm.TxCount / msm.Duration
	// 			result <- &msm
	// 			return
	// 		}
	// 		time.Sleep(time.Millisecond * 500)
	// 	}
	// }
}

// SubscribeNewHead subscribes to new blocks
func SubscribeNewHead(
	client *ethclient.Client, headers chan *types.Header, stop chan bool,
) {
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err == nil {
		<-stop
		sub.Unsubscribe()
		return
	}

	go func() {
		lastBlk, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			panic(err)
		}

		for {
			select {
			case <-stop:
				return

			default:
				currentBlk, _ := client.HeaderByNumber(context.Background(), nil)
				for i := lastBlk.Number.Int64() + 1; i < currentBlk.Number.Int64(); i++ {
					hdr, _ := client.HeaderByNumber(context.Background(), big.NewInt(i))
					headers <- hdr
				}
				if currentBlk.Number.Int64() > lastBlk.Number.Int64() {
					headers <- currentBlk
				}
				lastBlk = currentBlk
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
}
