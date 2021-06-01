package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"math"
	mRand "math/rand"
	"os"
)

// Deploy ...
func (c *Client) Deploy(contractPath, chaincodeIDFile, endpoint string) {

	data := makeDeployRequest("deploy", contractPath, "Init", "")
	res := c.postRequest(data, endpoint)
	f, _ := os.Create(chaincodeIDFile)
	defer f.Close()
	f.WriteString(res.Result.Message)

}

// Verify ...
func (c *Client) Verify(
	endpoint, pubKey, hash, r, s string,
) string {

	data := makeRequest(
		"invoke",
		c.chaincodeID,
		"verifymeter_sc",
		"null",
		pubKey, hash, r, s,
	)

	// fmt.Println("size", len(data))
	res := c.postRequest(data, endpoint)

	return res.Result.Message
}

// Query ...
func (c *Client) Query(
	endpoint, hash string,
) string {

	data := makeRequest(
		"query",
		c.chaincodeID,
		"none",
		hash,
		"querymeter",
	)
	res := c.postRequest(data, endpoint)
	return res.Result.Message
}

// PutProposal ...
func (c *Client) PutProposal(
	endpoint, pubKey, proposalType, proposedKey, ts, hash, r, s string,
) string {
	data := makeRequest(
		"invoke",
		c.chaincodeID,
		"putproposal_sc",
		proposalType, pubKey, proposedKey, ts, hash, r, s,
	)

	// fmt.Println("size", len(data))
	res := c.postRequest(data, endpoint)

	return res.Result.Message
}

func (c *Client) genAddKeyProposal(endpoint string, index int) (
	proposalType, proposedKey, bhStr, hash, r, s string,
) {

	proposalType = "addkey"
	blockHeight, _ := getBlockHeight(endpoint)
	bhStr = fmt.Sprintf("%d", blockHeight)

	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pubKey, _ := x509.MarshalPKIXPublicKey(privKey.Public())

	keystr := base64.StdEncoding.EncodeToString(pubKey)

	proposedKey = fmt.Sprintf("ll_%d_01_%s_%s", index, keystr, bhStr)
	hash = hashString(proposedKey)

	r, s, _ = signPayload(c.privKey, hash)

	return proposalType, proposedKey, bhStr, hash, r, s
}

// AddWaterReading ...
func (c *Client) AddWaterReading(
	endpoint, pubKey, ts, reading, r, s string,
) string {

	data := makeRequest(
		"invoke",
		c.chaincodeID,
		"addwaterreading_sc",
		"addwaterreading",
		pubKey, ts, "", reading, r, s,
	)

	// fmt.Println("size", len(data))
	res := c.postRequest(data, endpoint)

	return res.Result.Message
}

func (c *Client) genWaterReading(minutes int) (ts, reading, r, s string) {
	min := minutes % 60
	hour := int(math.Floor(float64(minutes) / 60))
	ts = fmt.Sprintf("%d_%02d", hour, min)

	f1 := mRand.Float64() * 1000
	f2 := mRand.Float64() * 2000
	f3 := mRand.Float64() * 1500

	reading = fmt.Sprintf("%s,20,%f;%s,40,%f;%s,50,%f", ts, f1, ts, f2, ts, f3)
	r, s, _ = signPayload(c.privKey, reading)

	return ts, reading, r, s
}
