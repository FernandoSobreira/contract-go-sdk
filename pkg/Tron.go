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

// Torn
// Contract struct
type Torn struct {
	Account     *Account
	client      *client.GrpcClient
	transaction *core.Transaction
}

// initClient
func (t *Torn) initClient() error {
	t.client = &client.GrpcClient{}
	t.client.SetTimeout(10 * time.Second)

	return t.client.Start(grpc.WithInsecure())
}

// createAddress
func (t *Torn) createAddress() (*Account, error) {

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
func (t *Torn) trans(number, feeLimit int64) (bool, []byte, error) {

	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	return true, nil, nil
}

// trans20
func (t *Torn) trans20(number, feeLimit int64) (bool, []byte, error) {

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
		return false, nil, errors.New(fmt.Sprintf("t.client.TRC20Send [ERROR] : %v", err))
	}

	t.transaction = res.GetTransaction()

	bro, err := t.broadcast()
	if err != nil {
		return false, nil, err
	}

	return bro.Result, bro.Message, nil
}

// freeze
func (t *Torn) freeze(number int64) (bool, []byte, error) {

	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	res, err := t.client.FreezeBalanceV2(t.Account.PublicAddress, core.ResourceCode_ENERGY, number)
	if err != nil {
		return false, nil, errors.New(fmt.Sprintf("t.client.FreezeBalanceV2 [ERROR] : %v", err))
	}

	t.transaction = res.GetTransaction()

	bro, err := t.broadcast()
	if err != nil {
		return false, nil, err
	}

	return bro.Result, bro.Message, nil
}

// witness
func (t *Torn) witness(witnessMap map[string]int64) (bool, []byte, error) {

	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	res, err := t.client.VoteWitnessAccount(t.Account.PublicAddress, witnessMap)
	if err != nil {
		return false, nil, errors.New(fmt.Sprintf("t.client.VoteWitnessAccount [ERROR]: %v", err))
	}

	t.transaction = res.Transaction

	bro, err := t.broadcast()
	if err != nil {
		return false, nil, err
	}

	return bro.Result, bro.Message, nil
}

// witnessWithdraw
func (t *Torn) witnessWithdraw() (bool, []byte, error) {

	err := t.initClient()
	if err != nil {
		return false, nil, err
	}
	defer t.client.Conn.Close()

	res, err := t.client.WithdrawBalance(t.Account.PublicAddress)
	if err != nil {
		return false, nil, errors.New(fmt.Sprintf("t.client.WithdrawBalance [ERROR] : %v", err))
	}

	t.transaction = res.GetTransaction()

	bro, err := t.broadcast()
	if err != nil {
		return false, nil, err
	}

	return bro.Result, bro.Message, nil
}

// broadcast
func (t *Torn) broadcast() (*api.Return, error) {
	if t.transaction == nil {
		return nil, errors.New("t.transaction [ERROR] : transaction is nil")
	}

	err := t.sign()
	if err != nil {
		return nil, err
	}

	bro, err := t.client.Broadcast(t.transaction)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("t.client.Broadcast [ERROR] : %v", err))
	}
	return bro, nil
}

// sign
func (t *Torn) sign() error {

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
