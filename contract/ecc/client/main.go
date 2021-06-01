package main

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// constants
const (
	CONTRACTPATH = "fbi/contract/ecc"
	CHAINCODEID  = "chaincodeid"
)

func main() {

	h, _ := ioutil.ReadFile("hosts")
	hosts := strings.Split(string(h), "\n")

	keyfile, err := os.Open("privKey")
	d := json.NewDecoder(keyfile)

	var bPrivKey []byte
	check(d.Decode(&bPrivKey))

	privKey, err := x509.ParseECPrivateKey(bPrivKey)
	check(err)

	chaincodeID, _ := ioutil.ReadFile(CHAINCODEID)

	pubKeyStr := "ll_01_01_MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDQYipMdqIzb80asaBoRUgtzNOhhTgLu6dunz3NuhAkDQVwkcfv0iR6+HSBy8FmeiSklhsQpTdyeEum6uh5fPYw==_0"

	c := Client{
		privKey:     privKey,
		chaincodeID: string(chaincodeID),
		pubKeyStr:   pubKeyStr,
		hosts:       hosts,
	}

	switch os.Args[1] {
	case "deploy":
		{ // deploy
			c.Deploy(CONTRACTPATH, CHAINCODEID, hosts[0])
		}

	case "verify":
		{ // verify
			reading := readMeter(createMeterID())
			fmt.Printf("Reading: %+v\n", reading)

			hash := hashReading(reading)
			fmt.Printf("Hash: %s\n", hash)

			blockHeight, _ := getBlockHeight(hosts[0])
			fmt.Printf("Block Height: %d\n", blockHeight)

			hash = fmt.Sprintf("%s_%d", hash, blockHeight)

			r, s, err := signPayload(c.privKey, hash)
			check(err)
			fmt.Printf("Signature: \n R: %s \n S: %s\n", r, s)

			fmt.Println(c.Verify(
				hosts[0], c.pubKeyStr, hash, r, s,
			))
		}

	case "query":
		{ // query <hash>
			fmt.Println(c.Query(hosts[0], os.Args[2]))
		}

	case "start":
		{ // start <txtype> <nthreads> <nOutstandingRequests> <nRequestPerSecond>
			// <nMeters> <nServers> <duration>
			// ./client start log 10 100 10 5 1 20

			client := &Client{
				nThreads:             toInt(os.Args[3]),
				nOutstandingRequests: toInt(os.Args[4]),
				nRequestsPerSecond:   toInt(os.Args[5]),
				nMeters:              toInt(os.Args[6]),
				hosts:                hosts[:toInt(os.Args[7])],
				duration:             toInt(os.Args[8]),

				privKey:     privKey,
				chaincodeID: string(chaincodeID),
				pubKeyStr:   pubKeyStr,
			}

			client.Start(os.Args[2])

		}

	case "latency":
		{ // latency
			c.Latency()
		}

	case "addkey":
		{ // addkey
			proposalType, proposedKey, bhStr, hash, r, s :=
				c.genAddKeyProposal(
					c.hosts[0], 1,
				)

			fmt.Println("ProposerKey:", c.pubKeyStr)
			fmt.Println("ProposedKey:", proposedKey)
			fmt.Println("Block height:", bhStr)
			fmt.Println("Hash:", hash)

			txID := c.PutProposal(
				c.hosts[0], c.pubKeyStr, proposalType, proposedKey, bhStr, hash, r, s,
			)
			fmt.Println("txID", txID)

			startTime := time.Now()

			for {
				time.Sleep(time.Millisecond * 500)
				if findTx(c.hosts[0], txID) {
					break
				}
			}

			latency := int(time.Now().Unix() - startTime.Unix())

			fmt.Printf("Latency: %d\n", latency)
		}

	case "addwr":
		{
			ts, reading, r, s := c.genWaterReading(toInt(os.Args[2]))

			fmt.Println("timestamp", ts)
			fmt.Println("reading", reading)

			txID := c.AddWaterReading(c.hosts[0], c.pubKeyStr, ts, reading, r, s)
			fmt.Println("txID", txID)

			startTime := time.Now()

			for {
				time.Sleep(time.Millisecond * 500)
				if findTx(c.hosts[0], txID) {
					break
				}
			}

			latency := int(time.Now().Unix() - startTime.Unix())

			fmt.Printf("Latency: %d\n", latency)
		}
	}

	fmt.Printf("Done\n")
}
