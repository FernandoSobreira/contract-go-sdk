package pkg

// Client
// Make To Client Contract interface
type Client interface {
	createAddress() (*Account, error)                                                    // Create Contract Address
	trans(number, feeLimit int64) (bool, []byte, error)                                  // Trans basic coin (ps: trx...)
	trans20(number, feeLimit int64) (bool, []byte, error)                                // Trans hot coin (ps: trc20, erc20...)
	freeze(number int64, code ResourceCode) (bool, []byte, error)                        // freeze coin
	unFreeze(number int64) (bool, []byte, error)                                         // unFreeze coin
	witness(witnessMap map[string]int64) (bool, []byte, error)                           // witness
	witnessWithdraw() (bool, []byte, error)                                              // witnessWithdraw
	delegateResource(number int64, code ResourceCode, lock bool) (bool, []byte, error)   // delegateResource
	unDelegateResource(number int64, code ResourceCode, lock bool) (bool, []byte, error) // unDelegateResource
	getBalance() (*Wallet, error)                                                        // getBalance
	getAccountResource() (*AccountResource, error)                                       // getAccountResource
}
