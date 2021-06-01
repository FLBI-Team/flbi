//go:generate abigen --sol ecc.sol --pkg ecc --type ECC --out ecc_gen.go

package ecc

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewECCSession creates a new ECC contract session
func NewECCSession(
	client *ethclient.Client, auth *bind.TransactOpts, conAddr string,
) (*ECCSession, error) {

	hashStorage, err := NewECC(
		common.HexToAddress(conAddr),
		client,
	)
	if err != nil {
		return nil, err
	}

	return &ECCSession{
		Contract:     hashStorage,
		TransactOpts: *auth,
	}, nil
}
