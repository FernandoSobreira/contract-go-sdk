package pkg

import (
	"context"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

// Account
// Contact Account struct
type Account struct {
	PrivateKey      string
	PublicAddress   string
	ToAddress       string
	ContractAddress string
	Context         context.Context
	Transaction     *core.Transaction
}

// AccountResource
// Contact AccountResource struct
type AccountResource struct {
	FreeNetLimit           int64
	TotalNetLimit          int64
	TotalNetWeight         int64
	TronPowerUsed          int64
	TronPowerLimit         int64
	EnergyLimit            int64
	TotalEnergyLimit       int64
	TotalEnergyWeightLimit int64
	AssetNetUsed           map[string]int64
	AssetNetLimit          map[string]int64
}

// Wallet
// Contact Network Wallet Balance
type Wallet struct {
	TRX   int64 // trx
	TRC20 int64 // trc20
	ERC20 int64 // erc20
	OKTC  int64 // oktc
}
