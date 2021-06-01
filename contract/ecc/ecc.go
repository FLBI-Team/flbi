package main

import (
	"fmt"
	"math"

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// FBIChainCode godoc
type FBIChainCode struct{}

func main() {
	err := shim.Start(new(FBIChainCode))
	if err != nil {
		fmt.Printf("Error starting BlockchainID: %s", err)
	}
}

// key types
var (
	PubKeyTypeAdmin = "Admin"
	PubKeyTypeMeter = "Meter"
	PubKeyTypeMf    = "Manufacturer"

	KeyMJCount     = "MJCount"
	KeyProposal    = "MeterProposal"
	KeyMeasurement = "Measurement"
	KeyFirmware    = "Firmware"
	KeyADModels    = "ADModels"

	KeyMembershipState   = "MembershipState"
	KeyMembershipUpdates = "MembershipUpdates"

	KeyAdminKeyState   = "AdminKeyState"
	KeyAdminKeyUpdates = "AdminKeyUpdates"
)

var rootAdminKeyState = `
{
	"version": 0,
	"key": "A8RyOesZGkIA1sliE1tnPue4CE5OIeYoEo1Cs6HAZj3H"
}
`

// Init godoc
func (cc *FBIChainCode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// We could initialize.
	adminKeys := []string{
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEaMDIHXKFZyLgNzintGwRYoXBo6hQ7vyEpudHeB8iHQr9n1fffwkJP3nIBm5TD9XEUjk72rAZEbylkSJOekTW0A==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEA+y+Wg6a963cFTSbljspxyQ2XiTeLumSfBHpLfWk4YB0OoQP/0e406cvXlvB0kHwejQrOQ4+cjAzpTwbZATbew==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJaeRxAsru5DGm6QJIUR3TVSItwE5GY1Peklq3/7S8RO3JK7pCvbDfv4UpB/fjpN5g9Q5Xq4LreLf6RuGBxfxpw==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEdWCZO190zxDXf5+WnjdeIXNDFaoR7hMSIRN7jIN26n94vtHj+czclJ4GX3qsq/9fKENl5zAMudrfwSOnjBiTtw==",
	}

	for _, key := range adminKeys {
		stub.PutState(key, []byte(PubKeyTypeAdmin))
	}
	mjCount := uint32(3)
	mjCountB := make([]byte, 4)
	binary.LittleEndian.PutUint32(mjCountB, mjCount)
	stub.PutState(KeyMJCount, mjCountB)

	meterKeys := []string{
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDQYipMdqIzb80asaBoRUgtzNOhhTgLu6dunz3NuhAkDQVwkcfv0iR6+HSBy8FmeiSklhsQpTdyeEum6uh5fPYw==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEVi7BchwENEK/Y4e4kof1rORWQRpvoqQIoHYAAYH/MsAtNNYUWHMAiH96wkC/bAdFgleOnd34OFnL2CK2wnzYNA==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEgmSEvRzwLjlp6cBERDE0e//7qHZLw4Z2bkJSPiilj+GHAkT4loMn4EFHSaML8E4SiR26m8+lXnUNPObuABQtCA==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE6UtbXT0jBD18LDuHAODhHWSIy8EMal6Eq43ObMEHwGCepkxWolC15/HnODPisOgu/nRPYBu5CW5F0TQ+QqQCnw==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE0iKZAH22tb0uv6etUXKKE3iiwnJJze6yuetBb89mb9Totdm8WbQNcpbit7gUaErAfzVT9XuUy7mWVjYb4lIu7Q==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHdDkElFdg71Tf7/n+MTXojqpzvAhb4Sqllh6lE4KmqY8j2NOcLp4ax+22R16Q6aqXi13Z7vXLSZvWeJ3sPcUnw==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAElQj4o9lHRgufuA019DizlmQvKkV4VxMAfClB0KnicjrrhKo76a9Pw+7Me8ICkH81HZKIvWDblUuvAjpc2DfcBw==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7FLH5zuRbtKH104OFp2wChj7FYmL1lfnS2GKZKkBLmue8ssXnxm1MtN95+skaekDm8M6WA0pzYUfVM6JZWMNFg==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfsukuaFXGar7IoCei3P/85I9Mze3Ih3yeb0CJvFqYi+XB4ZAIWbOwVKgn1Xr7wvG9XlQ4uZ7rBKfyCSrAkC+GA==",
		"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEwhVv8pZFB/FOUboYpL+5S9kSHT4nGEsoXZxLMr38jKv41ssOUux5m6Qm/Iw4fZ3kyXFNgcJS1MIkEw/Kgb5CMw==",
	}

	for _, key := range meterKeys {
		stub.PutState(key, []byte(PubKeyTypeMeter))
	}

	stub.PutState(KeyMembershipState, []byte("initial membershiop json"))
	stub.PutState(KeyAdminKeyState, []byte(rootAdminKeyState))
	return nil, nil
}

// Invoke godoc
func (cc *FBIChainCode) Invoke(
	stub shim.ChaincodeStubInterface, function string, args []string,
) ([]byte, error) {

	switch function {

	case "proposePubKey":
		return cc.proposePubKey(stub, args)

	case "votePubKeyProposal":
		return cc.votePubKeyProposal(stub, args)

	case "putMeasurement":
		return cc.putMeasurement(stub, args)

	case "putFirmware":
		return cc.putFirmware(stub, args)

	case "vetoFirmware":
		return cc.vetoFirmware(stub, args)

	case "updateAdminKey":
		return cc.updateAdminKey(stub, args)

	case "updateMembership":
		return cc.updateMembership(stub, args)
	}
	return nil, fmt.Errorf("method not found")
}

// Query godoc
func (cc *FBIChainCode) Query(
	stub shim.ChaincodeStubInterface, function string, args []string,
) ([]byte, error) {

	switch function {

	case "pubKeyProposal":
		return stub.GetState(makeProposalKey(args[0]))

	case "pubKey":
		return stub.GetState(args[0])

	case "measurement":
		return stub.GetState(makeMeasurementKey(args[0]))

	case "detectAnomaly":
		return cc.detectAnomaly(stub, args)

	case "checkFirmware":
		return cc.checkFirmware(stub, args)

	case "checkFirmwareUpgrade":
		return cc.checkFirmwareUpgrade(stub, args)

	case "firmware":
		return cc.queryFirmware(stub, args)

	case "adminKeyState":
		return stub.GetState(KeyAdminKeyState)

	case "adminKeyUpdates":
		return stub.GetState(KeyAdminKeyUpdates)

	case "membershipState":
		return stub.GetState(KeyMembershipState)

	case "membershipUpdates":
		return stub.GetState(KeyMembershipUpdates)
	}

	return nil, fmt.Errorf("method not found")
}

// PubKeyProposal godoc
type PubKeyProposal struct {
	Type     string
	Add      bool
	Key      string
	Proposer string
	Voters   map[string]bool
	Approved bool
}

func (cc *FBIChainCode) proposePubKey(
	stub shim.ChaincodeStubInterface, args []string,
) ([]byte, error) {

	data := args[0]
	sender := args[1]
	r := args[2]
	s := args[3]

	if err := cc.verifySender(stub, sender, r, s, data, PubKeyTypeAdmin); err != nil {
		return nil, err
	}
	if p, err := cc.loadProposal(stub, data); err == nil {
		if !p.Approved {
			return nil, fmt.Errorf("proposal already pending")
		}
	}
	pubKey, kType, add := parseInputProposalKey(data)
	p := &PubKeyProposal{
		Type:     kType,
		Add:      add,
		Key:      pubKey,
		Proposer: sender,
		Voters:   map[string]bool{sender: true},
		Approved: false,
	}
	cc.putProposal(stub, data, p)
	return []byte("Proposed public key"), nil
}

func parseInputProposalKey(data string) (string, string, bool) {
	strs := strings.Split(data, "_")
	key := strs[0]
	kt := strs[1]
	add, _ := strconv.ParseBool(strs[2])
	return key, kt, add
}

func (cc *FBIChainCode) putProposal(
	stub shim.ChaincodeStubInterface, data string, p *PubKeyProposal,
) error {
	pB, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return stub.PutState(makeProposalKey(data), pB)
}

func (cc *FBIChainCode) loadProposal(
	stub shim.ChaincodeStubInterface, data string,
) (*PubKeyProposal, error) {
	pB, _ := stub.GetState(makeProposalKey(data))
	if pB == nil {
		return nil, fmt.Errorf("failed to load proposal")
	}
	p := new(PubKeyProposal)
	return p, json.Unmarshal(pB, p)
}

func makeProposalKey(data string) string {
	return fmt.Sprintf("%s_%s", KeyProposal, data)
}

func (cc *FBIChainCode) votePubKeyProposal(
	stub shim.ChaincodeStubInterface, args []string,
) ([]byte, error) {

	data := args[0]
	sender := args[1]
	r := args[2]
	s := args[3]

	if err := cc.verifySender(stub, sender, r, s, data, PubKeyTypeAdmin); err != nil {
		return nil, err
	}
	p, err := cc.loadProposal(stub, data)
	if err != nil {
		return nil, err
	}
	if p.Approved {
		return nil, fmt.Errorf("proposal already approved")
	}
	if p.Voters[sender] {
		return nil, fmt.Errorf("cannot vote twice by same voter")
	}
	p.Voters[sender] = true
	mjCount := cc.loadMJCount(stub)

	if len(p.Voters) >= int(mjCount) {
		p.Approved = true
		stub.PutState(p.Key, []byte(p.Type))
	}
	cc.putProposal(stub, data, p)
	return []byte("Voted Add Meter"), nil
}

func (cc *FBIChainCode) loadMJCount(stub shim.ChaincodeStubInterface) uint32 {
	b, _ := stub.GetState(KeyMJCount)
	return binary.LittleEndian.Uint32(b)
}

func (cc *FBIChainCode) putMeasurement(
	stub shim.ChaincodeStubInterface, args []string,
) ([]byte, error) {

	data := args[0]
	sender := args[1]
	r := args[2]
	s := args[3]

	if err := cc.verifySender(stub, sender, r, s, data, PubKeyTypeMeter); err != nil {
		return nil, err
	}
	hash, blkHeight := parseHashAndBlkHeight(data)
	value := fmt.Sprintf("%d_%s", blkHeight, sender)
	if outsideTimeBound(blkHeight) {
		value += "_OutsideTB"
	}

	stub.PutState(makeMeasurementKey(hash), []byte(value))
	return []byte("Put Measurement succeeded"), nil
}

func parseHashAndBlkHeight(data string) (hash string, blkHeight int) {
	strs := strings.Split(data, "_")
	hash = strs[0]
	blkHeight, _ = strconv.Atoi(strs[1])
	return hash, blkHeight
}

func makeMeasurementKey(hash string) string {
	return fmt.Sprintf("%s_%s", KeyMeasurement, hash)
}

// AnomalyDetectionModels godoc
type AnomalyDetectionModels struct {
	FrequencyLB float64
	FrequencyUB float64
	VoltageLB   float64
	VoltageUB   float64
	CurrentUB   float64
}

func (cc *FBIChainCode) detectAnomaly(
	stub shim.ChaincodeStubInterface, args []string,
) ([]byte, error) {

	frequency, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return nil, err
	}
	voltage, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return nil, err
	}
	current, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return nil, err
	}

	adModels := cc.loadADModels(stub)
	anomalyMsg := make([]string, 0, 3)

	if frequency < adModels.FrequencyLB || frequency > adModels.FrequencyUB {
		anomalyMsg = append(anomalyMsg, "Frequency Anomaly")
	}
	if voltage < adModels.VoltageLB || voltage > adModels.VoltageUB {
		anomalyMsg = append(anomalyMsg, "Voltage Anomaly")
	}
	if current > adModels.CurrentUB {
		anomalyMsg = append(anomalyMsg, "Current Anomaly")
	}
	return []byte(strings.Join(anomalyMsg, ", ")), nil
}

type AdminKeyState struct {
	Version uint32 `json:"version"`
	Key     []byte `json:"key"`
}

func (ks *AdminKeyState) Hash() []byte {
	h := sha256.New()
	binary.Write(h, binary.LittleEndian, ks.Version)
	h.Write(ks.Key)
	return h.Sum(nil)
}

type AdminKeyUpdate struct {
	AdminKeyState
	Signature []byte `json:"signature"`
}

func (cc *FBIChainCode) updateAdminKey(
	stub shim.ChaincodeStubInterface, args []string,
) ([]byte, error) {

	input := new(AdminKeyUpdate)
	err := json.Unmarshal([]byte(args[0]), input)
	if err != nil {
		return nil, err
	}

	b, err := stub.GetState(KeyAdminKeyState)
	if err != nil {
		return nil, err
	}
	state := new(AdminKeyState)
	err = json.Unmarshal(b, state)
	if err != nil {
		return nil, err
	}

	if input.Version != state.Version+1 {
		return nil, fmt.Errorf("invlaid key version")
	}

	if !verifySig(state.Key, input.Signature, input.AdminKeyState.Hash()) {
		return nil, fmt.Errorf("invalid signatures")
	}

	b, _ = stub.GetState(KeyAdminKeyUpdates)
	var updates []*AdminKeyUpdate
	json.Unmarshal(b, &updates)
	updates = append(updates, input)

	b, _ = json.Marshal(updates)
	stub.PutState(KeyAdminKeyUpdates, b)

	b, err = json.Marshal(input.AdminKeyState)
	if err != nil {
		return nil, err
	}
	return nil, stub.PutState(KeyAdminKeyState, b)
}

type MembershipState struct {
	Version uint32   `json:"version"`
	Members [][]byte `json:"members"`
}

func (ms *MembershipState) Hash() []byte {
	h := sha256.New()
	binary.Write(h, binary.LittleEndian, ms.Version)
	for _, m := range ms.Members {
		h.Write(m)
	}
	return h.Sum(nil)
}

type MembershipUpdate struct {
	NextState  MembershipState `json:"nextState"`
	SigIndexes []int           `json:"sigIndexes"`
	SigValues  [][]byte        `json:"sigValues"`
}

func hasInvalidSigs(state *MembershipState, input *MembershipUpdate) bool {
	hash := input.NextState.Hash()
	for i, idx := range input.SigIndexes {
		if idx >= len(state.Members) {
			return true
		}
		if !verifySig(state.Members[idx], input.SigValues[i], hash) {
			return true // found invalid sig
		}
	}
	return false
}

func verifySig(pub, sig, msg []byte) bool {
	curve := elliptic.P256()
	x, y := elliptic.UnmarshalCompressed(curve, pub)
	r, s := sig[:32], sig[32:]
	pubKey := &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}
	return ecdsa.Verify(pubKey, msg, big.NewInt(0).SetBytes(r), big.NewInt(0).SetBytes(s))
}

func (cc *FBIChainCode) updateMembership(
	stub shim.ChaincodeStubInterface, args []string,
) ([]byte, error) {

	input := new(MembershipUpdate)
	err := json.Unmarshal([]byte(args[0]), input)
	if err != nil {
		return nil, err
	}

	b, err := stub.GetState(KeyMembershipState)
	if err != nil {
		return nil, err
	}
	state := new(MembershipState)
	err = json.Unmarshal(b, state)
	if err != nil {
		return nil, err
	}

	if input.NextState.Version != state.Version+1 {
		return nil, fmt.Errorf("invlaid membership version")
	}

	if len(input.SigIndexes) < majorityCount(len(state.Members)) {
		return nil, fmt.Errorf("not enough signatures")
	}

	if hasDuplicates(input.SigIndexes) {
		return nil, fmt.Errorf("input contains duplicate signatures")
	}

	if hasInvalidSigs(state, input) {
		return nil, fmt.Errorf("input contains invalid signatures")
	}

	b, _ = stub.GetState(KeyMembershipUpdates)
	var updates []*MembershipUpdate
	json.Unmarshal(b, &updates)
	updates = append(updates, input)

	b, _ = json.Marshal(updates)
	stub.PutState(KeyMembershipUpdates, b)

	b, err = json.Marshal(input.NextState)
	if err != nil {
		return nil, err
	}
	return nil, stub.PutState(KeyMembershipState, b)
}

// majorityCount returns 2f + 1 members
func majorityCount(total int) int {
	// n=3f+1 -> f=floor((n-1)3) -> m=n-f -> m=ceil((2n+1)/3)
	return int(math.Ceil(float64(2*total+1) / 3))
}

func hasDuplicates(list []int) bool {
	dupMap := make(map[int]struct{}, len(list))
	for _, v := range list {
		if _, found := dupMap[v]; found {
			return true // found duplicate
		}
		dupMap[v] = struct{}{}
	}
	return false
}

type rawFirmwareInfo struct {
	BlkHeight    int    `json:"blkHeight"`
	Manufacturer string `json:"manufacturer"`
	Vetoed       bool   `json:"vetoed"`
	VetoedBy     string `json:"vetoedBy"`
}

// FirmwareInfo godoc
type FirmwareInfo struct {
	Status       string `json:"status"` // "Not Trusted | Pending | Granted | Vetoed"
	BlockHeight  int    `json:"blockHeight"`
	Manufacturer string `json:"manufacturer"`
	VetoedBy     string `json:"vetoedBy"`
}

func (cc *FBIChainCode) putFirmware(
	stub shim.ChaincodeStubInterface, args []string,
) ([]byte, error) {

	data := args[0]
	sender := args[1]
	r := args[2]
	s := args[3]

	if err := cc.verifySender(stub, sender, r, s, data, PubKeyTypeMf); err != nil {
		return nil, err
	}
	hash, blkHeight := parseHashAndBlkHeight(data)
	if _, err := cc.loadRawFirmware(stub, hash); err == nil {
		return nil, fmt.Errorf("firmware already added") // firmware found
	}
	fwi := &rawFirmwareInfo{
		BlkHeight:    blkHeight,
		Manufacturer: sender,
	}
	cc.storeFirmware(stub, hash, fwi)
	return []byte("Put Firmware succeeded"), nil
}

func (cc *FBIChainCode) storeFirmware(
	stub shim.ChaincodeStubInterface, hash string, fwi *rawFirmwareInfo,
) error {
	fwB, err := json.Marshal(fwi)
	if err != nil {
		return err
	}
	return stub.PutState(makeFirmwareKey(hash), fwB)
}

func (cc *FBIChainCode) queryFirmware(
	stub shim.ChaincodeStubInterface, args []string,
) ([]byte, error) {
	hash := args[0]
	fw, err := cc.loadFirmware(stub, hash)
	if err != nil {
		return nil, err
	}
	return json.Marshal(fw)
}

func (cc *FBIChainCode) loadFirmware(
	stub shim.ChaincodeStubInterface, hash string,
) (*FirmwareInfo, error) {
	raw, err := cc.loadRawFirmware(stub, hash)
	fw := new(FirmwareInfo)
	if err != nil {
		fw.Status = "Not Trusted"
		return fw, nil
	}
	fw.BlockHeight = raw.BlkHeight
	fw.Manufacturer = raw.Manufacturer
	if raw.Vetoed {
		fw.Status = "Vetoed"
		fw.VetoedBy = raw.VetoedBy
		return fw, nil
	}
	if !isCooldownTimePassed(raw.BlkHeight) {
		fw.Status = "Pending"
		return fw, nil
	}
	fw.Status = "Granted"
	return fw, nil
}

func (cc *FBIChainCode) loadRawFirmware(
	stub shim.ChaincodeStubInterface, hash string,
) (*rawFirmwareInfo, error) {
	fwB, _ := stub.GetState(makeFirmwareKey(hash))
	if fwB == nil {
		return nil, fmt.Errorf("failed to load firmware")
	}
	fwi := new(rawFirmwareInfo)
	return fwi, json.Unmarshal(fwB, fwi)
}

func makeFirmwareKey(hash string) string {
	return fmt.Sprintf("%s_%s", KeyFirmware, hash)
}

func (cc *FBIChainCode) checkFirmware(
	stub shim.ChaincodeStubInterface, args []string,
) ([]byte, error) {
	hash := args[0]
	fw, err := cc.loadRawFirmware(stub, hash)
	if err != nil {
		return nil, fmt.Errorf("firmware not found")
	}
	if err := cc.verifyFwGranted(fw); err != nil {
		return nil, err
	}
	return []byte("firmware GRANTED"), nil
}

func (cc *FBIChainCode) verifyFwGranted(fw *rawFirmwareInfo) error {
	if fw.Vetoed {
		return fmt.Errorf("vetoed")
	}
	if !isCooldownTimePassed(fw.BlkHeight) {
		return fmt.Errorf("firmware pending")
	}
	return nil
}

func (cc *FBIChainCode) checkFirmwareUpgrade(
	stub shim.ChaincodeStubInterface, args []string,
) ([]byte, error) {
	hashNew := args[0]
	hashOld := args[1]

	fwNew, err := cc.loadRawFirmware(stub, hashNew)
	if err != nil {
		return nil, fmt.Errorf("firmware not found")
	}
	if err := cc.verifyFwGranted(fwNew); err != nil {
		return nil, err
	}

	fwOld, err := cc.loadRawFirmware(stub, hashOld)
	if err == nil {
		if fwNew.BlkHeight <= fwOld.BlkHeight {
			return nil, fmt.Errorf("old firmware, cannot upgrade")
		}
	}
	return []byte("firmware upgrade GRANTED"), nil
}

func (cc *FBIChainCode) vetoFirmware(
	stub shim.ChaincodeStubInterface, args []string,
) ([]byte, error) {
	hash := args[0]
	sender := args[1]
	r := args[2]
	s := args[3]

	if err := cc.verifySender(stub, sender, r, s, hash, PubKeyTypeAdmin); err != nil {
		return nil, err
	}
	fw, err := cc.loadRawFirmware(stub, hash)
	if err != nil {
		return nil, err
	}
	if fw.Vetoed {
		return nil, fmt.Errorf("cannot veto twice")
	}
	if isCooldownTimePassed(fw.BlkHeight) {
		return nil, fmt.Errorf("cannot veto after cooldown time passed")
	}
	fw.Vetoed = true
	fw.VetoedBy = sender
	cc.storeFirmware(stub, hash, fw)
	return []byte("Vetoed firmware"), nil
}

func isCooldownTimePassed(startBlk int) bool {
	endpoint := "127.0.0.1"
	cooldowntime := 5 //blocks
	blocknow := findheight(endpoint).Height

	return blocknow > startBlk+cooldowntime
}

func (cc *FBIChainCode) verifySender(
	stub shim.ChaincodeStubInterface, sender, r, s, data, keyType string,
) error {
	kt, _ := stub.GetState(sender)
	if string(kt) != keyType {
		return fmt.Errorf("key not found on Blockchain %s", keyType)
	}
	if !verifySigOld(sender, r, s, data) {
		return fmt.Errorf("signature verification failed")
	}
	return nil
}

func verifySigOld(keyStr, rStr, sStr, hash string) bool {
	keyb, err := base64.StdEncoding.DecodeString(keyStr)
	if err != nil {
		return false
	}
	rb, err := base64.StdEncoding.DecodeString(rStr)
	if err != nil {
		return false
	}
	sb, err := base64.StdEncoding.DecodeString(sStr)
	if err != nil {
		return false
	}
	key, err := x509.ParsePKIXPublicKey(keyb)
	if err != nil {
		return false
	}
	r, s := big.NewInt(0), big.NewInt(0)
	err = r.UnmarshalText(rb)
	if err != nil {
		return false
	}
	err = s.UnmarshalText(sb)
	if err != nil {
		return false
	}
	return ecdsa.Verify(key.(*ecdsa.PublicKey), []byte(hash), r, s)
}

func outsideTimeBound(clientblock int) bool {
	endpoint := "127.0.0.1"
	allowedskew := 5 //blocks for testing. actual must be lesser.
	blocknow := findheight(endpoint).Height
	plusskew := blocknow + allowedskew
	minusskew := blocknow - allowedskew

	return clientblock < minusskew || clientblock > plusskew
}

func findheight(endpoint string) getResponse {
	return getRequest(endpoint)
}

func getRequest(endpoint string) getResponse {
	var rs getResponse
	res, err := http.Get("http://" + endpoint + ":7050/chain")
	if err != nil {
		return rs
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&rs)
	return rs
}

type getResponse struct {
	Height            int    `json:"height,omitempty"`
	CurrentBlockHash  string `json:"currentBlockHash,omitempty"`
	PreviousBlockHash string `json:"previousBlockHash,omitempty"`
}

func (cc *FBIChainCode) loadADModels(stub shim.ChaincodeStubInterface) *AnomalyDetectionModels {
	return &AnomalyDetectionModels{
		FrequencyLB: 40,
		FrequencyUB: 60,
		VoltageLB:   200,
		VoltageUB:   250,
		CurrentUB:   4,
	}
}
