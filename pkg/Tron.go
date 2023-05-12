package pkg

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"math/big"
	"time"
)

// ResourceCode
// Account Resource
type ResourceCode int32

const (
	ResourceCode_BANDWIDTH  ResourceCode = 0
	ResourceCode_ENERGY     ResourceCode = 1
	ResourceCode_TRON_POWER ResourceCode = 2
)

// Tron
// Contract struct
type Tron struct {
	Account     *Account
	client      *client.GrpcClient
	transaction *core.Transaction
}

// initClient
func (t *Tron) initClient() error {
	t.client = &client.GrpcClient{}
	t.client.SetTimeout(10 * time.Second)

	return t.client.Start(grpc.WithInsecure())
}

// createAddress
func (t *Tron) createAddress() (*Account, error) {

	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	private := common.Encode(crypto.FromECDSA(privateKeyECDSA))
	address := address.PubkeyToAddress(privateKeyECDSA.PublicKey)

	return &Account{
		PrivateKey:    private,
		PublicAddress: address.String(),
		Context:       context.Background(),
	}, nil
}

// trans
func (t *Tron) trans(number, feeLimit int64) (bool, []byte, error) {

	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	return true, nil, nil
}

// trans20
func (t *Tron) trans20(number, feeLimit int64) (bool, []byte, error) {

	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	if feeLimit == 0 {
		feeLimit = 5000000
	}

	res, err := t.client.TRC20Send(t.Account.PublicAddress, t.Account.ToAddress, t.Account.ContractAddress, big.NewInt(number), feeLimit)
	if err != nil {
		return false, nil, err
	}

	t.transaction = res.GetTransaction()

	bro, err := t.broadcast()
	if err != nil {
		return false, nil, err
	}

	return bro.Result, bro.Message, nil
}

// freeze
func (t *Tron) freeze(number int64, code ResourceCode) (bool, []byte, error) {

	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	res, err := t.client.FreezeBalanceV2(t.Account.PublicAddress, core.ResourceCode(code), number)
	if err != nil {
		return false, nil, err
	}

	t.transaction = res.GetTransaction()

	bro, err := t.broadcast()
	if err != nil {
		return false, nil, err
	}

	return bro.Result, bro.Message, nil
}

// unFreeze
func (t *Tron) unFreeze(number int64) (bool, []byte, error) {

	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	res, err := t.client.UnfreezeBalanceV2(t.Account.PublicAddress, core.ResourceCode_ENERGY, number)
	if err != nil {
		return false, nil, err
	}

	t.transaction = res.GetTransaction()

	bro, err := t.broadcast()
	if err != nil {
		return false, nil, err
	}

	return bro.Result, bro.Message, nil
}

// witness
func (t *Tron) witness(witnessMap map[string]int64) (bool, []byte, error) {

	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	res, err := t.client.VoteWitnessAccount(t.Account.PublicAddress, witnessMap)
	if err != nil {
		return false, nil, err
	}

	t.transaction = res.Transaction

	bro, err := t.broadcast()
	if err != nil {
		return false, nil, err
	}

	return bro.Result, bro.Message, nil
}

// witnessWithdraw
func (t *Tron) witnessWithdraw() (bool, []byte, error) {

	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	res, err := t.client.WithdrawBalance(t.Account.PublicAddress)
	if err != nil {
		return false, nil, err
	}

	t.transaction = res.GetTransaction()

	bro, err := t.broadcast()
	if err != nil {
		return false, nil, err
	}

	return bro.Result, bro.Message, nil
}

// delegateResource
func (t *Tron) delegateResource(number int64, code ResourceCode, lock bool) (bool, []byte, error) {
	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	res, err := t.client.DelegateResource(t.Account.PublicAddress, t.Account.ToAddress, core.ResourceCode(code), number, lock)
	if err != nil {
		if err != nil {
			return false, nil, err
		}
	}

	t.transaction = res.GetTransaction()
	err = t.sign()
	if err != nil {
		return false, nil, err
	}

	bro, err := t.broadcast()
	if err != nil {
		return false, nil, err
	}

	return bro.GetResult(), bro.GetMessage(), nil
}

// unDelegateResource
func (t *Tron) unDelegateResource(number int64, code ResourceCode, lock bool) (bool, []byte, error) {
	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	res, err := t.client.UnDelegateResource(t.Account.PublicAddress, t.Account.ToAddress, core.ResourceCode(code), number, lock)
	if err != nil {
		if err != nil {
			return false, nil, err
		}
	}

	t.transaction = res.GetTransaction()
	err = t.sign()
	if err != nil {
		return false, nil, err
	}

	bro, err := t.broadcast()
	if err != nil {
		return false, nil, err
	}

	return bro.GetResult(), bro.GetMessage(), nil
}

// getBalance
func (t *Tron) getBalance() (*Wallet, error) {
	err := t.initClient()
	if err != nil {
		return nil, err
	}
	defer t.client.Conn.Close()

	account, err := t.client.GetAccount(t.Account.PublicAddress)
	if err != nil {
		return nil, err
	}

	coin, err := t.client.TRC20ContractBalance(t.Account.PublicAddress, t.Account.ContractAddress)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		TRX:   account.GetBalance(),
		TRC20: coin.Int64(),
		ERC20: 0,
		OKTC:  0,
	}, nil

}

// getAccountResource
func (t *Tron) getAccountResource() (*AccountResource, error) {
	err := t.initClient()
	if err != nil {
		return nil, err
	}
	defer t.client.Conn.Close()

	res, err := t.client.GetAccountResource(t.Account.PublicAddress)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("t.client.GetAccount [ERROR] : %v", err))
	}

	return &AccountResource{
		FreeNetLimit:           res.GetFreeNetLimit(),
		TotalNetLimit:          res.GetNetLimit(),
		TotalNetWeight:         res.GetTotalNetWeight(),
		TronPowerUsed:          res.GetTronPowerUsed(),
		TronPowerLimit:         res.GetFreeNetLimit(),
		EnergyLimit:            res.GetEnergyLimit(),
		TotalEnergyLimit:       res.GetTotalEnergyLimit(),
		TotalEnergyWeightLimit: res.GetTotalEnergyWeight(),
		AssetNetUsed:           res.GetAssetNetUsed(),
		AssetNetLimit:          res.GetAssetNetLimit(),
	}, nil

}

// broadcast
func (t *Tron) broadcast() (*api.Return, error) {
	if t.transaction == nil {
		return nil, errors.New("transaction [ERROR] : transaction is nil")
	}

	err := t.sign()
	if err != nil {
		return nil, err
	}

	bro, err := t.client.Broadcast(t.transaction)
	if err != nil {
		return nil, err
	}
	return bro, nil
}

// sign
func (t *Tron) sign() error {

	privateDec, err := common.Decode(t.Account.PrivateKey)
	if err != nil {
		return errors.New(fmt.Sprintf("common.Decode [ERROR] : %v", err))
	}
	privateECDSA, err := crypto.ToECDSA(privateDec)
	if err != nil {
		return errors.New(fmt.Sprintf("crypto.ToECDSA [ERROR] : %v", err))
	}

	rawData, err := proto.Marshal(t.transaction.GetRawData())
	if err != nil {
		return errors.New(fmt.Sprintf("proto.Marshal [ERROR] : %v", err))
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	signature, err := crypto.Sign(hash, privateECDSA)
	if err != nil {
		return errors.New(fmt.Sprintf("crypto.Sign [ERROR] : %v", err))
	}

	t.transaction.Signature = append(t.transaction.Signature, signature)
	return nil
}
