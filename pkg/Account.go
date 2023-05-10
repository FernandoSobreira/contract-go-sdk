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
