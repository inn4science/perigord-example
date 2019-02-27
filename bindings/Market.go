// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// MarketABI is the input ABI used to generate the binding from.
const MarketABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deals\",\"outputs\":[{\"name\":\"senderAddress\",\"type\":\"address\"},{\"name\":\"receiverAddress\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getDealReceiver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"returnDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getDealSender\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getReceiver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"setDealSender\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"setSenderReceiverPair\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getDealDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"setDealReceiver\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_investor\",\"type\":\"address\"}],\"name\":\"getDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"getTokenSender\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDealId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// MarketBin is the compiled bytecode used for deploying new contracts.
const MarketBin = `60806040819052600180556729a2241af62c000060025560008054600160a060020a0319163317808255600160a060020a0316917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3610927806100676000396000f3006080604052600436106100f05763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166303988f84811461018857806307f54239146101ca57806318b683fb146101fe5780631f766ed61461022a5780633f036ce6146102425780635c990dfb14610263578063715018a6146102875780637e07e3111461029e5780638da5cb5b146102c55780638f32d59b146102da57806393d73a9e146102ef578063b47c9d5614610319578063e1254fba1461033d578063ef78c91b1461035e578063f2fde38b1461037f578063fa238137146103a0578063fc7e286d146103b5575b600080600034111561017f57600254341461010a57600080fd5b610113336103d6565b9150600160a060020a038216151561012a57600080fd5b6101353383346103f4565b600160a060020a038316600090815260066020526040902054909150610161903463ffffffff61048b16565b600160a060020a038316600090815260066020526040902055610184565b600080fd5b5050005b34801561019457600080fd5b506101a06004356104a1565b60408051600160a060020a0394851681529290931660208301528183015290519081900360600190f35b3480156101d657600080fd5b506101e26004356104cf565b60408051600160a060020a039092168252519081900360200190f35b34801561020a57600080fd5b506102166004356104ed565b604080519115158252519081900360200190f35b34801561023657600080fd5b506101e26004356105b4565b34801561024e57600080fd5b506101e2600160a060020a03600435166103d6565b34801561026f57600080fd5b50610216600160a060020a03600435166024356105cf565b34801561029357600080fd5b5061029c61062d565b005b3480156102aa57600080fd5b5061029c600160a060020a036004358116906024351661068a565b3480156102d157600080fd5b506101e2610783565b3480156102e657600080fd5b50610216610792565b3480156102fb57600080fd5b506103076004356107a3565b60408051918252519081900360200190f35b34801561032557600080fd5b50610216600160a060020a03600435166024356107b8565b34801561034957600080fd5b50610307600160a060020a036004351661081b565b34801561036a57600080fd5b506101e2600160a060020a0360043516610836565b34801561038b57600080fd5b5061029c600160a060020a0360043516610854565b3480156103ac57600080fd5b50610307610873565b3480156103c157600080fd5b50610307600160a060020a0360043516610879565b600160a060020a039081166000908152600460205260409020541690565b600180546000818152600360205260408082208054600160a060020a03808a16600160a060020a03199283161780845583880180548b841690851617815560028086018b81558a548952968820805486169385169390931783559054828a01805490951693169290921790925592549201919091558354919390929161047f9163ffffffff61048b16565b60015595945050505050565b60008282018381101561049a57fe5b9392505050565b600360205260009081526040902080546001820154600290920154600160a060020a03918216929091169083565b600090815260036020526040902060010154600160a060020a031690565b60008060006104fa610792565b151561050557600080fd5b5050600082815260036020526040808220600281018054908490556001820154925191939092600160a060020a03169183156108fc0291849190818181858888f1935050505015801561055c573d6000803e3d6000fd5b50600084815260036020526040902082548154600160a060020a0319908116600160a060020a039283161783556001808601548185018054909316931692909217905560028085015492019190915592505050919050565b600090815260036020526040902054600160a060020a031690565b6000806105da610792565b15156105e557600080fd5b505060009081526003602052604090208054600160a060020a03928316600160a060020a03199182161780821690841617825560019182018054918216919093161790915590565b610635610792565b151561064057600080fd5b60008054604051600160a060020a03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a360008054600160a060020a0319169055565b600080600080610698610792565b15156106a357600080fd5b600160a060020a038616158015906106c35750600160a060020a03851615155b15156106ce57600080fd5b505050600160a060020a0380841660009081526004602081815260408084205460058352818520548887168652828620549490935293205492841615945083161592908116159116158380156107215750825b801561072a5750805b80156107335750815b1561017f57600160a060020a0380871660008181526004602090815260408083208054958b16600160a060020a031996871681179091558352600590915290208054909216179055505050505050565b600054600160a060020a031690565b600054600160a060020a0316331490565b60009081526003602052604090206002015490565b6000806107c3610792565b15156107ce57600080fd5b50506000818152600360205260409020600180820180548354600160a060020a0319818116600160a060020a03928316178655818816928116929092179182169116179055905092915050565b600160a060020a031660009081526006602052604090205490565b600160a060020a039081166000908152600560205260409020541690565b61085c610792565b151561086757600080fd5b6108708161088b565b50565b60015481565b60066020526000908152604090205481565b600160a060020a03811615156108a057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a360008054600160a060020a031916600160a060020a03929092169190911790555600a165627a7a723058208d9d703509425951c319ec5d4ef3bc16b2e6de72191556fbff5b1d545e1441fe0029`

// DeployMarket deploys a new Ethereum contract, binding an instance of Market to it.
func DeployMarket(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Market, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// Deals is a free data retrieval call binding the contract method 0x03988f84.
//
// Solidity: function deals( uint256) constant returns(senderAddress address, receiverAddress address, deposit uint256)
func (_Market *MarketCaller) Deals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SenderAddress   common.Address
	ReceiverAddress common.Address
	Deposit         *big.Int
}, error) {
	ret := new(struct {
		SenderAddress   common.Address
		ReceiverAddress common.Address
		Deposit         *big.Int
	})
	out := ret
	err := _Market.contract.Call(opts, out, "deals", arg0)
	return *ret, err
}

// Deals is a free data retrieval call binding the contract method 0x03988f84.
//
// Solidity: function deals( uint256) constant returns(senderAddress address, receiverAddress address, deposit uint256)
func (_Market *MarketSession) Deals(arg0 *big.Int) (struct {
	SenderAddress   common.Address
	ReceiverAddress common.Address
	Deposit         *big.Int
}, error) {
	return _Market.Contract.Deals(&_Market.CallOpts, arg0)
}

// Deals is a free data retrieval call binding the contract method 0x03988f84.
//
// Solidity: function deals( uint256) constant returns(senderAddress address, receiverAddress address, deposit uint256)
func (_Market *MarketCallerSession) Deals(arg0 *big.Int) (struct {
	SenderAddress   common.Address
	ReceiverAddress common.Address
	Deposit         *big.Int
}, error) {
	return _Market.Contract.Deals(&_Market.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits( address) constant returns(uint256)
func (_Market *MarketCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "deposits", arg0)
	return *ret0, err
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits( address) constant returns(uint256)
func (_Market *MarketSession) Deposits(arg0 common.Address) (*big.Int, error) {
	return _Market.Contract.Deposits(&_Market.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits( address) constant returns(uint256)
func (_Market *MarketCallerSession) Deposits(arg0 common.Address) (*big.Int, error) {
	return _Market.Contract.Deposits(&_Market.CallOpts, arg0)
}

// GetDealDeposit is a free data retrieval call binding the contract method 0x93d73a9e.
//
// Solidity: function getDealDeposit(id uint256) constant returns(uint256)
func (_Market *MarketCaller) GetDealDeposit(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getDealDeposit", id)
	return *ret0, err
}

// GetDealDeposit is a free data retrieval call binding the contract method 0x93d73a9e.
//
// Solidity: function getDealDeposit(id uint256) constant returns(uint256)
func (_Market *MarketSession) GetDealDeposit(id *big.Int) (*big.Int, error) {
	return _Market.Contract.GetDealDeposit(&_Market.CallOpts, id)
}

// GetDealDeposit is a free data retrieval call binding the contract method 0x93d73a9e.
//
// Solidity: function getDealDeposit(id uint256) constant returns(uint256)
func (_Market *MarketCallerSession) GetDealDeposit(id *big.Int) (*big.Int, error) {
	return _Market.Contract.GetDealDeposit(&_Market.CallOpts, id)
}

// GetDealReceiver is a free data retrieval call binding the contract method 0x07f54239.
//
// Solidity: function getDealReceiver(id uint256) constant returns(address)
func (_Market *MarketCaller) GetDealReceiver(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getDealReceiver", id)
	return *ret0, err
}

// GetDealReceiver is a free data retrieval call binding the contract method 0x07f54239.
//
// Solidity: function getDealReceiver(id uint256) constant returns(address)
func (_Market *MarketSession) GetDealReceiver(id *big.Int) (common.Address, error) {
	return _Market.Contract.GetDealReceiver(&_Market.CallOpts, id)
}

// GetDealReceiver is a free data retrieval call binding the contract method 0x07f54239.
//
// Solidity: function getDealReceiver(id uint256) constant returns(address)
func (_Market *MarketCallerSession) GetDealReceiver(id *big.Int) (common.Address, error) {
	return _Market.Contract.GetDealReceiver(&_Market.CallOpts, id)
}

// GetDealSender is a free data retrieval call binding the contract method 0x1f766ed6.
//
// Solidity: function getDealSender(id uint256) constant returns(address)
func (_Market *MarketCaller) GetDealSender(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getDealSender", id)
	return *ret0, err
}

// GetDealSender is a free data retrieval call binding the contract method 0x1f766ed6.
//
// Solidity: function getDealSender(id uint256) constant returns(address)
func (_Market *MarketSession) GetDealSender(id *big.Int) (common.Address, error) {
	return _Market.Contract.GetDealSender(&_Market.CallOpts, id)
}

// GetDealSender is a free data retrieval call binding the contract method 0x1f766ed6.
//
// Solidity: function getDealSender(id uint256) constant returns(address)
func (_Market *MarketCallerSession) GetDealSender(id *big.Int) (common.Address, error) {
	return _Market.Contract.GetDealSender(&_Market.CallOpts, id)
}

// GetDeposit is a free data retrieval call binding the contract method 0xe1254fba.
//
// Solidity: function getDeposit(_investor address) constant returns(uint256)
func (_Market *MarketCaller) GetDeposit(opts *bind.CallOpts, _investor common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getDeposit", _investor)
	return *ret0, err
}

// GetDeposit is a free data retrieval call binding the contract method 0xe1254fba.
//
// Solidity: function getDeposit(_investor address) constant returns(uint256)
func (_Market *MarketSession) GetDeposit(_investor common.Address) (*big.Int, error) {
	return _Market.Contract.GetDeposit(&_Market.CallOpts, _investor)
}

// GetDeposit is a free data retrieval call binding the contract method 0xe1254fba.
//
// Solidity: function getDeposit(_investor address) constant returns(uint256)
func (_Market *MarketCallerSession) GetDeposit(_investor common.Address) (*big.Int, error) {
	return _Market.Contract.GetDeposit(&_Market.CallOpts, _investor)
}

// GetReceiver is a free data retrieval call binding the contract method 0x3f036ce6.
//
// Solidity: function getReceiver(_sender address) constant returns(address)
func (_Market *MarketCaller) GetReceiver(opts *bind.CallOpts, _sender common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getReceiver", _sender)
	return *ret0, err
}

// GetReceiver is a free data retrieval call binding the contract method 0x3f036ce6.
//
// Solidity: function getReceiver(_sender address) constant returns(address)
func (_Market *MarketSession) GetReceiver(_sender common.Address) (common.Address, error) {
	return _Market.Contract.GetReceiver(&_Market.CallOpts, _sender)
}

// GetReceiver is a free data retrieval call binding the contract method 0x3f036ce6.
//
// Solidity: function getReceiver(_sender address) constant returns(address)
func (_Market *MarketCallerSession) GetReceiver(_sender common.Address) (common.Address, error) {
	return _Market.Contract.GetReceiver(&_Market.CallOpts, _sender)
}

// GetTokenSender is a free data retrieval call binding the contract method 0xef78c91b.
//
// Solidity: function getTokenSender(_receiver address) constant returns(address)
func (_Market *MarketCaller) GetTokenSender(opts *bind.CallOpts, _receiver common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getTokenSender", _receiver)
	return *ret0, err
}

// GetTokenSender is a free data retrieval call binding the contract method 0xef78c91b.
//
// Solidity: function getTokenSender(_receiver address) constant returns(address)
func (_Market *MarketSession) GetTokenSender(_receiver common.Address) (common.Address, error) {
	return _Market.Contract.GetTokenSender(&_Market.CallOpts, _receiver)
}

// GetTokenSender is a free data retrieval call binding the contract method 0xef78c91b.
//
// Solidity: function getTokenSender(_receiver address) constant returns(address)
func (_Market *MarketCallerSession) GetTokenSender(_receiver common.Address) (common.Address, error) {
	return _Market.Contract.GetTokenSender(&_Market.CallOpts, _receiver)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Market *MarketCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Market *MarketSession) IsOwner() (bool, error) {
	return _Market.Contract.IsOwner(&_Market.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Market *MarketCallerSession) IsOwner() (bool, error) {
	return _Market.Contract.IsOwner(&_Market.CallOpts)
}

// LastDealId is a free data retrieval call binding the contract method 0xfa238137.
//
// Solidity: function lastDealId() constant returns(uint256)
func (_Market *MarketCaller) LastDealId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "lastDealId")
	return *ret0, err
}

// LastDealId is a free data retrieval call binding the contract method 0xfa238137.
//
// Solidity: function lastDealId() constant returns(uint256)
func (_Market *MarketSession) LastDealId() (*big.Int, error) {
	return _Market.Contract.LastDealId(&_Market.CallOpts)
}

// LastDealId is a free data retrieval call binding the contract method 0xfa238137.
//
// Solidity: function lastDealId() constant returns(uint256)
func (_Market *MarketCallerSession) LastDealId() (*big.Int, error) {
	return _Market.Contract.LastDealId(&_Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Market *MarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Market *MarketSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Market *MarketCallerSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _Market.Contract.RenounceOwnership(&_Market.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Market.Contract.RenounceOwnership(&_Market.TransactOpts)
}

// ReturnDeposit is a paid mutator transaction binding the contract method 0x18b683fb.
//
// Solidity: function returnDeposit(id uint256) returns(bool)
func (_Market *MarketTransactor) ReturnDeposit(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "returnDeposit", id)
}

// ReturnDeposit is a paid mutator transaction binding the contract method 0x18b683fb.
//
// Solidity: function returnDeposit(id uint256) returns(bool)
func (_Market *MarketSession) ReturnDeposit(id *big.Int) (*types.Transaction, error) {
	return _Market.Contract.ReturnDeposit(&_Market.TransactOpts, id)
}

// ReturnDeposit is a paid mutator transaction binding the contract method 0x18b683fb.
//
// Solidity: function returnDeposit(id uint256) returns(bool)
func (_Market *MarketTransactorSession) ReturnDeposit(id *big.Int) (*types.Transaction, error) {
	return _Market.Contract.ReturnDeposit(&_Market.TransactOpts, id)
}

// SetDealReceiver is a paid mutator transaction binding the contract method 0xb47c9d56.
//
// Solidity: function setDealReceiver(receiver address, id uint256) returns(bool)
func (_Market *MarketTransactor) SetDealReceiver(opts *bind.TransactOpts, receiver common.Address, id *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setDealReceiver", receiver, id)
}

// SetDealReceiver is a paid mutator transaction binding the contract method 0xb47c9d56.
//
// Solidity: function setDealReceiver(receiver address, id uint256) returns(bool)
func (_Market *MarketSession) SetDealReceiver(receiver common.Address, id *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetDealReceiver(&_Market.TransactOpts, receiver, id)
}

// SetDealReceiver is a paid mutator transaction binding the contract method 0xb47c9d56.
//
// Solidity: function setDealReceiver(receiver address, id uint256) returns(bool)
func (_Market *MarketTransactorSession) SetDealReceiver(receiver common.Address, id *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetDealReceiver(&_Market.TransactOpts, receiver, id)
}

// SetDealSender is a paid mutator transaction binding the contract method 0x5c990dfb.
//
// Solidity: function setDealSender(sender address, id uint256) returns(bool)
func (_Market *MarketTransactor) SetDealSender(opts *bind.TransactOpts, sender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setDealSender", sender, id)
}

// SetDealSender is a paid mutator transaction binding the contract method 0x5c990dfb.
//
// Solidity: function setDealSender(sender address, id uint256) returns(bool)
func (_Market *MarketSession) SetDealSender(sender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetDealSender(&_Market.TransactOpts, sender, id)
}

// SetDealSender is a paid mutator transaction binding the contract method 0x5c990dfb.
//
// Solidity: function setDealSender(sender address, id uint256) returns(bool)
func (_Market *MarketTransactorSession) SetDealSender(sender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetDealSender(&_Market.TransactOpts, sender, id)
}

// SetSenderReceiverPair is a paid mutator transaction binding the contract method 0x7e07e311.
//
// Solidity: function setSenderReceiverPair(_sender address, _receiver address) returns()
func (_Market *MarketTransactor) SetSenderReceiverPair(opts *bind.TransactOpts, _sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setSenderReceiverPair", _sender, _receiver)
}

// SetSenderReceiverPair is a paid mutator transaction binding the contract method 0x7e07e311.
//
// Solidity: function setSenderReceiverPair(_sender address, _receiver address) returns()
func (_Market *MarketSession) SetSenderReceiverPair(_sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Market.Contract.SetSenderReceiverPair(&_Market.TransactOpts, _sender, _receiver)
}

// SetSenderReceiverPair is a paid mutator transaction binding the contract method 0x7e07e311.
//
// Solidity: function setSenderReceiverPair(_sender address, _receiver address) returns()
func (_Market *MarketTransactorSession) SetSenderReceiverPair(_sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Market.Contract.SetSenderReceiverPair(&_Market.TransactOpts, _sender, _receiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Market *MarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Market *MarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Market.Contract.TransferOwnership(&_Market.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Market *MarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Market.Contract.TransferOwnership(&_Market.TransactOpts, newOwner)
}

// MarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Market contract.
type MarketOwnershipTransferredIterator struct {
	Event *MarketOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketOwnershipTransferred represents a OwnershipTransferred event raised by the Market contract.
type MarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Market *MarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketOwnershipTransferredIterator{contract: _Market.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Market *MarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketOwnershipTransferred)
				if err := _Market.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
