package pkg

// CreateAddress
// Create Contract Address
func CreateAddress(cli Client) (*Account, error) {
	return cli.createAddress()
}

// Transaction20
// Trans hot coin (ps: trc20, erc20...)
// @param number => 1000000 == 1
// @param feeLimit => 50000000 == 50
func Transaction20(cli Client, number, feeLimit int64) (bool, []byte, error) {
	return cli.trans20(number, feeLimit)
}

// Freeze
// freeze coin
// @param number => 1000000 == 1
// @param code ResourceCode_ENERGY
func Freeze(cli Client, number int64, code ResourceCode) (bool, []byte, error) {
	return cli.freeze(number, code)
}

// UnFreeze
// unFreeze coin
// @param number => 1000000 == 1
func UnFreeze(cli Client, number int64) (bool, []byte, error) {
	return cli.unFreeze(number)
}

// Witness
// Vote for witnesses
// @param witnessMap map[string]int64{"Super address": 1}
func Witness(cli Client, witnessMap map[string]int64) (bool, []byte, error) {
	return cli.witness(witnessMap)
}

// WitnessWithdraw
// Super Representative or user withdraw rewards, usable every 24 hours
func WitnessWithdraw(cli Client) (bool, []byte, error) {
	return cli.witnessWithdraw()
}

// DelegateResource
// Proxy bandwidth or power resources to other accounts
func DelegateResource(cli Client, number int64, code ResourceCode, lock bool) (bool, []byte, error) {
	return cli.delegateResource(number, code, lock)
}

// UnDelegateResource
// Un broker bandwidth or power for the destination address
func UnDelegateResource(cli Client, number int64, code ResourceCode, lock bool) (bool, []byte, error) {
	return cli.unDelegateResource(number, code, lock)
}

// GetBalance
// Get Contract Balance
func GetBalance(cli Client) (*Wallet, error) {
	return cli.getBalance()
}

// GetAccountResource
// Get Contract Account Resource
func GetAccountResource(cli Client) (*AccountResource, error) {
	return cli.getAccountResource()
}
