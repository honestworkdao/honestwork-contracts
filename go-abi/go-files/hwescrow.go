// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hwescrow

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HWEscrowDeal is an auto generated low-level Go binding around an user-defined struct.
type HWEscrowDeal struct {
	Recruiter       common.Address
	Creator         common.Address
	PaymentToken    common.Address
	TotalPayment    *big.Int
	SuccessFee      *big.Int
	ClaimedAmount   *big.Int
	ClaimableAmount *big.Int
	Status          uint8
	RecruiterRating []*big.Int
	CreatorRating   []*big.Int
}

// HwescrowMetaData contains all meta data concerning the Hwescrow contract.
var HwescrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stableCoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recruiter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"}],\"name\":\"AdditionalPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newPaymentLimit\",\"type\":\"uint256\"}],\"name\":\"ExtraLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newSuccessFee\",\"type\":\"uint256\"}],\"name\":\"FeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"FeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_grossRevenue\",\"type\":\"uint256\"}],\"name\":\"GrossRevenueUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recruiter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_totalPayment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"}],\"name\":\"OfferCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_paymentReceived\",\"type\":\"uint256\"}],\"name\":\"PaymentClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recruiter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_unlockedAmount\",\"type\":\"uint256\"}],\"name\":\"PaymentUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumHWEscrow.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"PaymentWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_collector\",\"type\":\"address\"}],\"name\":\"TotalFeeClaimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_recruiterNFT\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_rating\",\"type\":\"uint128\"}],\"name\":\"additionalPayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"additionalPaymentLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"allowNativePayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_limit\",\"type\":\"uint64\"}],\"name\":\"changeExtraPaymentLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIHWRegistry\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"changeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_fee\",\"type\":\"uint128\"}],\"name\":\"changeSuccessFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_rating\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"_creatorNFT\",\"type\":\"uint256\"}],\"name\":\"claimPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"claimSuccessFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"claimTotalSuccessFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recruiter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_downPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_recruiterNFTId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"createDeal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recruiter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_downPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_recruiterNFTId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"createDealSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dealIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dealsMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recruiter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"successFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimableAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumHWEscrow.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraPaymentLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getAdditionalPaymentLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getAggregatedRating\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getAvgCreatorRating\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getAvgRecruiterRating\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getCreatorRating\",\"outputs\":[{\"internalType\":\"uint128[]\",\"name\":\"\",\"type\":\"uint128[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recruiter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"successFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimableAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumHWEscrow.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint128[]\",\"name\":\"recruiterRating\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"creatorRating\",\"type\":\"uint128[]\"}],\"internalType\":\"structHWEscrow.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getDealCompletionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getDealStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getDealSuccessFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getDealsOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getEthPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recruiter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_downPayment\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getNFTGrossRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getPaymentToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrecision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getRecruiter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getRecruiterRating\",\"outputs\":[{\"internalType\":\"uint128[]\",\"name\":\"\",\"type\":\"uint128[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getTotalPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSuccessFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"getclaimedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"honestWorkSuccessFee\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativePaymentAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIHWRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router01\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stableCoin\",\"type\":\"address\"}],\"name\":\"setStableCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableCoin\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCollectedSuccessFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_paymentAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_rating\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"_recruiterNFT\",\"type\":\"uint256\"}],\"name\":\"unlockPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealId\",\"type\":\"uint256\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HwescrowABI is the input ABI used to generate the binding from.
// Deprecated: Use HwescrowMetaData.ABI instead.
var HwescrowABI = HwescrowMetaData.ABI

// Hwescrow is an auto generated Go binding around an Ethereum contract.
type Hwescrow struct {
	HwescrowCaller     // Read-only binding to the contract
	HwescrowTransactor // Write-only binding to the contract
	HwescrowFilterer   // Log filterer for contract events
}

// HwescrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type HwescrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HwescrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HwescrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HwescrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HwescrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HwescrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HwescrowSession struct {
	Contract     *Hwescrow         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HwescrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HwescrowCallerSession struct {
	Contract *HwescrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HwescrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HwescrowTransactorSession struct {
	Contract     *HwescrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HwescrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type HwescrowRaw struct {
	Contract *Hwescrow // Generic contract binding to access the raw methods on
}

// HwescrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HwescrowCallerRaw struct {
	Contract *HwescrowCaller // Generic read-only contract binding to access the raw methods on
}

// HwescrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HwescrowTransactorRaw struct {
	Contract *HwescrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHwescrow creates a new instance of Hwescrow, bound to a specific deployed contract.
func NewHwescrow(address common.Address, backend bind.ContractBackend) (*Hwescrow, error) {
	contract, err := bindHwescrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hwescrow{HwescrowCaller: HwescrowCaller{contract: contract}, HwescrowTransactor: HwescrowTransactor{contract: contract}, HwescrowFilterer: HwescrowFilterer{contract: contract}}, nil
}

// NewHwescrowCaller creates a new read-only instance of Hwescrow, bound to a specific deployed contract.
func NewHwescrowCaller(address common.Address, caller bind.ContractCaller) (*HwescrowCaller, error) {
	contract, err := bindHwescrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HwescrowCaller{contract: contract}, nil
}

// NewHwescrowTransactor creates a new write-only instance of Hwescrow, bound to a specific deployed contract.
func NewHwescrowTransactor(address common.Address, transactor bind.ContractTransactor) (*HwescrowTransactor, error) {
	contract, err := bindHwescrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HwescrowTransactor{contract: contract}, nil
}

// NewHwescrowFilterer creates a new log filterer instance of Hwescrow, bound to a specific deployed contract.
func NewHwescrowFilterer(address common.Address, filterer bind.ContractFilterer) (*HwescrowFilterer, error) {
	contract, err := bindHwescrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HwescrowFilterer{contract: contract}, nil
}

// bindHwescrow binds a generic wrapper to an already deployed contract.
func bindHwescrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HwescrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hwescrow *HwescrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hwescrow.Contract.HwescrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hwescrow *HwescrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hwescrow.Contract.HwescrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hwescrow *HwescrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hwescrow.Contract.HwescrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hwescrow *HwescrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hwescrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hwescrow *HwescrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hwescrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hwescrow *HwescrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hwescrow.Contract.contract.Transact(opts, method, params...)
}

// AdditionalPaymentLimit is a free data retrieval call binding the contract method 0x18e335b1.
//
// Solidity: function additionalPaymentLimit(uint256 ) view returns(uint256)
func (_Hwescrow *HwescrowCaller) AdditionalPaymentLimit(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "additionalPaymentLimit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdditionalPaymentLimit is a free data retrieval call binding the contract method 0x18e335b1.
//
// Solidity: function additionalPaymentLimit(uint256 ) view returns(uint256)
func (_Hwescrow *HwescrowSession) AdditionalPaymentLimit(arg0 *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.AdditionalPaymentLimit(&_Hwescrow.CallOpts, arg0)
}

// AdditionalPaymentLimit is a free data retrieval call binding the contract method 0x18e335b1.
//
// Solidity: function additionalPaymentLimit(uint256 ) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) AdditionalPaymentLimit(arg0 *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.AdditionalPaymentLimit(&_Hwescrow.CallOpts, arg0)
}

// DealIds is a free data retrieval call binding the contract method 0x39e10379.
//
// Solidity: function dealIds() view returns(uint256 _value)
func (_Hwescrow *HwescrowCaller) DealIds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "dealIds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DealIds is a free data retrieval call binding the contract method 0x39e10379.
//
// Solidity: function dealIds() view returns(uint256 _value)
func (_Hwescrow *HwescrowSession) DealIds() (*big.Int, error) {
	return _Hwescrow.Contract.DealIds(&_Hwescrow.CallOpts)
}

// DealIds is a free data retrieval call binding the contract method 0x39e10379.
//
// Solidity: function dealIds() view returns(uint256 _value)
func (_Hwescrow *HwescrowCallerSession) DealIds() (*big.Int, error) {
	return _Hwescrow.Contract.DealIds(&_Hwescrow.CallOpts)
}

// DealsMapping is a free data retrieval call binding the contract method 0x79f6375f.
//
// Solidity: function dealsMapping(uint256 ) view returns(address recruiter, address creator, address paymentToken, uint256 totalPayment, uint256 successFee, uint256 claimedAmount, uint256 claimableAmount, uint8 status)
func (_Hwescrow *HwescrowCaller) DealsMapping(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Recruiter       common.Address
	Creator         common.Address
	PaymentToken    common.Address
	TotalPayment    *big.Int
	SuccessFee      *big.Int
	ClaimedAmount   *big.Int
	ClaimableAmount *big.Int
	Status          uint8
}, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "dealsMapping", arg0)

	outstruct := new(struct {
		Recruiter       common.Address
		Creator         common.Address
		PaymentToken    common.Address
		TotalPayment    *big.Int
		SuccessFee      *big.Int
		ClaimedAmount   *big.Int
		ClaimableAmount *big.Int
		Status          uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recruiter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PaymentToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TotalPayment = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SuccessFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ClaimedAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ClaimableAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// DealsMapping is a free data retrieval call binding the contract method 0x79f6375f.
//
// Solidity: function dealsMapping(uint256 ) view returns(address recruiter, address creator, address paymentToken, uint256 totalPayment, uint256 successFee, uint256 claimedAmount, uint256 claimableAmount, uint8 status)
func (_Hwescrow *HwescrowSession) DealsMapping(arg0 *big.Int) (struct {
	Recruiter       common.Address
	Creator         common.Address
	PaymentToken    common.Address
	TotalPayment    *big.Int
	SuccessFee      *big.Int
	ClaimedAmount   *big.Int
	ClaimableAmount *big.Int
	Status          uint8
}, error) {
	return _Hwescrow.Contract.DealsMapping(&_Hwescrow.CallOpts, arg0)
}

// DealsMapping is a free data retrieval call binding the contract method 0x79f6375f.
//
// Solidity: function dealsMapping(uint256 ) view returns(address recruiter, address creator, address paymentToken, uint256 totalPayment, uint256 successFee, uint256 claimedAmount, uint256 claimableAmount, uint8 status)
func (_Hwescrow *HwescrowCallerSession) DealsMapping(arg0 *big.Int) (struct {
	Recruiter       common.Address
	Creator         common.Address
	PaymentToken    common.Address
	TotalPayment    *big.Int
	SuccessFee      *big.Int
	ClaimedAmount   *big.Int
	ClaimableAmount *big.Int
	Status          uint8
}, error) {
	return _Hwescrow.Contract.DealsMapping(&_Hwescrow.CallOpts, arg0)
}

// ExtraPaymentLimit is a free data retrieval call binding the contract method 0x493d7e0c.
//
// Solidity: function extraPaymentLimit() view returns(uint64)
func (_Hwescrow *HwescrowCaller) ExtraPaymentLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "extraPaymentLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ExtraPaymentLimit is a free data retrieval call binding the contract method 0x493d7e0c.
//
// Solidity: function extraPaymentLimit() view returns(uint64)
func (_Hwescrow *HwescrowSession) ExtraPaymentLimit() (uint64, error) {
	return _Hwescrow.Contract.ExtraPaymentLimit(&_Hwescrow.CallOpts)
}

// ExtraPaymentLimit is a free data retrieval call binding the contract method 0x493d7e0c.
//
// Solidity: function extraPaymentLimit() view returns(uint64)
func (_Hwescrow *HwescrowCallerSession) ExtraPaymentLimit() (uint64, error) {
	return _Hwescrow.Contract.ExtraPaymentLimit(&_Hwescrow.CallOpts)
}

// GetAdditionalPaymentLimit is a free data retrieval call binding the contract method 0xfdfe0a3f.
//
// Solidity: function getAdditionalPaymentLimit(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetAdditionalPaymentLimit(opts *bind.CallOpts, _dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getAdditionalPaymentLimit", _dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAdditionalPaymentLimit is a free data retrieval call binding the contract method 0xfdfe0a3f.
//
// Solidity: function getAdditionalPaymentLimit(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetAdditionalPaymentLimit(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetAdditionalPaymentLimit(&_Hwescrow.CallOpts, _dealId)
}

// GetAdditionalPaymentLimit is a free data retrieval call binding the contract method 0xfdfe0a3f.
//
// Solidity: function getAdditionalPaymentLimit(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetAdditionalPaymentLimit(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetAdditionalPaymentLimit(&_Hwescrow.CallOpts, _dealId)
}

// GetAggregatedRating is a free data retrieval call binding the contract method 0xb7ba315c.
//
// Solidity: function getAggregatedRating(address _address) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetAggregatedRating(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getAggregatedRating", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggregatedRating is a free data retrieval call binding the contract method 0xb7ba315c.
//
// Solidity: function getAggregatedRating(address _address) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetAggregatedRating(_address common.Address) (*big.Int, error) {
	return _Hwescrow.Contract.GetAggregatedRating(&_Hwescrow.CallOpts, _address)
}

// GetAggregatedRating is a free data retrieval call binding the contract method 0xb7ba315c.
//
// Solidity: function getAggregatedRating(address _address) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetAggregatedRating(_address common.Address) (*big.Int, error) {
	return _Hwescrow.Contract.GetAggregatedRating(&_Hwescrow.CallOpts, _address)
}

// GetAvgCreatorRating is a free data retrieval call binding the contract method 0x896adcbd.
//
// Solidity: function getAvgCreatorRating(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetAvgCreatorRating(opts *bind.CallOpts, _dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getAvgCreatorRating", _dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvgCreatorRating is a free data retrieval call binding the contract method 0x896adcbd.
//
// Solidity: function getAvgCreatorRating(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetAvgCreatorRating(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetAvgCreatorRating(&_Hwescrow.CallOpts, _dealId)
}

// GetAvgCreatorRating is a free data retrieval call binding the contract method 0x896adcbd.
//
// Solidity: function getAvgCreatorRating(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetAvgCreatorRating(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetAvgCreatorRating(&_Hwescrow.CallOpts, _dealId)
}

// GetAvgRecruiterRating is a free data retrieval call binding the contract method 0x714d00b1.
//
// Solidity: function getAvgRecruiterRating(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetAvgRecruiterRating(opts *bind.CallOpts, _dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getAvgRecruiterRating", _dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvgRecruiterRating is a free data retrieval call binding the contract method 0x714d00b1.
//
// Solidity: function getAvgRecruiterRating(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetAvgRecruiterRating(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetAvgRecruiterRating(&_Hwescrow.CallOpts, _dealId)
}

// GetAvgRecruiterRating is a free data retrieval call binding the contract method 0x714d00b1.
//
// Solidity: function getAvgRecruiterRating(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetAvgRecruiterRating(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetAvgRecruiterRating(&_Hwescrow.CallOpts, _dealId)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0x7d8ca242.
//
// Solidity: function getClaimableAmount(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetClaimableAmount(opts *bind.CallOpts, _dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getClaimableAmount", _dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableAmount is a free data retrieval call binding the contract method 0x7d8ca242.
//
// Solidity: function getClaimableAmount(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetClaimableAmount(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetClaimableAmount(&_Hwescrow.CallOpts, _dealId)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0x7d8ca242.
//
// Solidity: function getClaimableAmount(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetClaimableAmount(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetClaimableAmount(&_Hwescrow.CallOpts, _dealId)
}

// GetCreator is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(uint256 _dealId) view returns(address)
func (_Hwescrow *HwescrowCaller) GetCreator(opts *bind.CallOpts, _dealId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getCreator", _dealId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreator is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(uint256 _dealId) view returns(address)
func (_Hwescrow *HwescrowSession) GetCreator(_dealId *big.Int) (common.Address, error) {
	return _Hwescrow.Contract.GetCreator(&_Hwescrow.CallOpts, _dealId)
}

// GetCreator is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(uint256 _dealId) view returns(address)
func (_Hwescrow *HwescrowCallerSession) GetCreator(_dealId *big.Int) (common.Address, error) {
	return _Hwescrow.Contract.GetCreator(&_Hwescrow.CallOpts, _dealId)
}

// GetCreatorRating is a free data retrieval call binding the contract method 0x2d7f7b51.
//
// Solidity: function getCreatorRating(uint256 _dealId) view returns(uint128[])
func (_Hwescrow *HwescrowCaller) GetCreatorRating(opts *bind.CallOpts, _dealId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getCreatorRating", _dealId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCreatorRating is a free data retrieval call binding the contract method 0x2d7f7b51.
//
// Solidity: function getCreatorRating(uint256 _dealId) view returns(uint128[])
func (_Hwescrow *HwescrowSession) GetCreatorRating(_dealId *big.Int) ([]*big.Int, error) {
	return _Hwescrow.Contract.GetCreatorRating(&_Hwescrow.CallOpts, _dealId)
}

// GetCreatorRating is a free data retrieval call binding the contract method 0x2d7f7b51.
//
// Solidity: function getCreatorRating(uint256 _dealId) view returns(uint128[])
func (_Hwescrow *HwescrowCallerSession) GetCreatorRating(_dealId *big.Int) ([]*big.Int, error) {
	return _Hwescrow.Contract.GetCreatorRating(&_Hwescrow.CallOpts, _dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 _dealId) view returns((address,address,address,uint256,uint256,uint256,uint256,uint8,uint128[],uint128[]))
func (_Hwescrow *HwescrowCaller) GetDeal(opts *bind.CallOpts, _dealId *big.Int) (HWEscrowDeal, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getDeal", _dealId)

	if err != nil {
		return *new(HWEscrowDeal), err
	}

	out0 := *abi.ConvertType(out[0], new(HWEscrowDeal)).(*HWEscrowDeal)

	return out0, err

}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 _dealId) view returns((address,address,address,uint256,uint256,uint256,uint256,uint8,uint128[],uint128[]))
func (_Hwescrow *HwescrowSession) GetDeal(_dealId *big.Int) (HWEscrowDeal, error) {
	return _Hwescrow.Contract.GetDeal(&_Hwescrow.CallOpts, _dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 _dealId) view returns((address,address,address,uint256,uint256,uint256,uint256,uint8,uint128[],uint128[]))
func (_Hwescrow *HwescrowCallerSession) GetDeal(_dealId *big.Int) (HWEscrowDeal, error) {
	return _Hwescrow.Contract.GetDeal(&_Hwescrow.CallOpts, _dealId)
}

// GetDealCompletionRate is a free data retrieval call binding the contract method 0x64fe46d2.
//
// Solidity: function getDealCompletionRate(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetDealCompletionRate(opts *bind.CallOpts, _dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getDealCompletionRate", _dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDealCompletionRate is a free data retrieval call binding the contract method 0x64fe46d2.
//
// Solidity: function getDealCompletionRate(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetDealCompletionRate(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetDealCompletionRate(&_Hwescrow.CallOpts, _dealId)
}

// GetDealCompletionRate is a free data retrieval call binding the contract method 0x64fe46d2.
//
// Solidity: function getDealCompletionRate(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetDealCompletionRate(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetDealCompletionRate(&_Hwescrow.CallOpts, _dealId)
}

// GetDealStatus is a free data retrieval call binding the contract method 0xc698f53b.
//
// Solidity: function getDealStatus(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetDealStatus(opts *bind.CallOpts, _dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getDealStatus", _dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDealStatus is a free data retrieval call binding the contract method 0xc698f53b.
//
// Solidity: function getDealStatus(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetDealStatus(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetDealStatus(&_Hwescrow.CallOpts, _dealId)
}

// GetDealStatus is a free data retrieval call binding the contract method 0xc698f53b.
//
// Solidity: function getDealStatus(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetDealStatus(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetDealStatus(&_Hwescrow.CallOpts, _dealId)
}

// GetDealSuccessFee is a free data retrieval call binding the contract method 0x89170a02.
//
// Solidity: function getDealSuccessFee(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetDealSuccessFee(opts *bind.CallOpts, _dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getDealSuccessFee", _dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDealSuccessFee is a free data retrieval call binding the contract method 0x89170a02.
//
// Solidity: function getDealSuccessFee(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetDealSuccessFee(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetDealSuccessFee(&_Hwescrow.CallOpts, _dealId)
}

// GetDealSuccessFee is a free data retrieval call binding the contract method 0x89170a02.
//
// Solidity: function getDealSuccessFee(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetDealSuccessFee(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetDealSuccessFee(&_Hwescrow.CallOpts, _dealId)
}

// GetDealsOf is a free data retrieval call binding the contract method 0x530f29e4.
//
// Solidity: function getDealsOf(address _address) view returns(uint256[])
func (_Hwescrow *HwescrowCaller) GetDealsOf(opts *bind.CallOpts, _address common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getDealsOf", _address)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDealsOf is a free data retrieval call binding the contract method 0x530f29e4.
//
// Solidity: function getDealsOf(address _address) view returns(uint256[])
func (_Hwescrow *HwescrowSession) GetDealsOf(_address common.Address) ([]*big.Int, error) {
	return _Hwescrow.Contract.GetDealsOf(&_Hwescrow.CallOpts, _address)
}

// GetDealsOf is a free data retrieval call binding the contract method 0x530f29e4.
//
// Solidity: function getDealsOf(address _address) view returns(uint256[])
func (_Hwescrow *HwescrowCallerSession) GetDealsOf(_address common.Address) ([]*big.Int, error) {
	return _Hwescrow.Contract.GetDealsOf(&_Hwescrow.CallOpts, _address)
}

// GetEthPrice is a free data retrieval call binding the contract method 0x870d365d.
//
// Solidity: function getEthPrice(uint256 _amount) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetEthPrice(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getEthPrice", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthPrice is a free data retrieval call binding the contract method 0x870d365d.
//
// Solidity: function getEthPrice(uint256 _amount) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetEthPrice(_amount *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetEthPrice(&_Hwescrow.CallOpts, _amount)
}

// GetEthPrice is a free data retrieval call binding the contract method 0x870d365d.
//
// Solidity: function getEthPrice(uint256 _amount) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetEthPrice(_amount *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetEthPrice(&_Hwescrow.CallOpts, _amount)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Hwescrow *HwescrowCaller) GetEthSignedMessageHash(opts *bind.CallOpts, _messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getEthSignedMessageHash", _messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Hwescrow *HwescrowSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Hwescrow.Contract.GetEthSignedMessageHash(&_Hwescrow.CallOpts, _messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Hwescrow *HwescrowCallerSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Hwescrow.Contract.GetEthSignedMessageHash(&_Hwescrow.CallOpts, _messageHash)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xc30aaeb7.
//
// Solidity: function getMessageHash(address _recruiter, address _creator, address _paymentToken, uint256 _totalPayment, uint256 _downPayment) pure returns(bytes32)
func (_Hwescrow *HwescrowCaller) GetMessageHash(opts *bind.CallOpts, _recruiter common.Address, _creator common.Address, _paymentToken common.Address, _totalPayment *big.Int, _downPayment *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getMessageHash", _recruiter, _creator, _paymentToken, _totalPayment, _downPayment)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0xc30aaeb7.
//
// Solidity: function getMessageHash(address _recruiter, address _creator, address _paymentToken, uint256 _totalPayment, uint256 _downPayment) pure returns(bytes32)
func (_Hwescrow *HwescrowSession) GetMessageHash(_recruiter common.Address, _creator common.Address, _paymentToken common.Address, _totalPayment *big.Int, _downPayment *big.Int) ([32]byte, error) {
	return _Hwescrow.Contract.GetMessageHash(&_Hwescrow.CallOpts, _recruiter, _creator, _paymentToken, _totalPayment, _downPayment)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xc30aaeb7.
//
// Solidity: function getMessageHash(address _recruiter, address _creator, address _paymentToken, uint256 _totalPayment, uint256 _downPayment) pure returns(bytes32)
func (_Hwescrow *HwescrowCallerSession) GetMessageHash(_recruiter common.Address, _creator common.Address, _paymentToken common.Address, _totalPayment *big.Int, _downPayment *big.Int) ([32]byte, error) {
	return _Hwescrow.Contract.GetMessageHash(&_Hwescrow.CallOpts, _recruiter, _creator, _paymentToken, _totalPayment, _downPayment)
}

// GetNFTGrossRevenue is a free data retrieval call binding the contract method 0x13faf9df.
//
// Solidity: function getNFTGrossRevenue(uint256 _tokenId) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetNFTGrossRevenue(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getNFTGrossRevenue", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNFTGrossRevenue is a free data retrieval call binding the contract method 0x13faf9df.
//
// Solidity: function getNFTGrossRevenue(uint256 _tokenId) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetNFTGrossRevenue(_tokenId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetNFTGrossRevenue(&_Hwescrow.CallOpts, _tokenId)
}

// GetNFTGrossRevenue is a free data retrieval call binding the contract method 0x13faf9df.
//
// Solidity: function getNFTGrossRevenue(uint256 _tokenId) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetNFTGrossRevenue(_tokenId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetNFTGrossRevenue(&_Hwescrow.CallOpts, _tokenId)
}

// GetPaymentToken is a free data retrieval call binding the contract method 0xcdc2aebf.
//
// Solidity: function getPaymentToken(uint256 _dealId) view returns(address)
func (_Hwescrow *HwescrowCaller) GetPaymentToken(opts *bind.CallOpts, _dealId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getPaymentToken", _dealId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPaymentToken is a free data retrieval call binding the contract method 0xcdc2aebf.
//
// Solidity: function getPaymentToken(uint256 _dealId) view returns(address)
func (_Hwescrow *HwescrowSession) GetPaymentToken(_dealId *big.Int) (common.Address, error) {
	return _Hwescrow.Contract.GetPaymentToken(&_Hwescrow.CallOpts, _dealId)
}

// GetPaymentToken is a free data retrieval call binding the contract method 0xcdc2aebf.
//
// Solidity: function getPaymentToken(uint256 _dealId) view returns(address)
func (_Hwescrow *HwescrowCallerSession) GetPaymentToken(_dealId *big.Int) (common.Address, error) {
	return _Hwescrow.Contract.GetPaymentToken(&_Hwescrow.CallOpts, _dealId)
}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() pure returns(uint256)
func (_Hwescrow *HwescrowCaller) GetPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getPrecision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() pure returns(uint256)
func (_Hwescrow *HwescrowSession) GetPrecision() (*big.Int, error) {
	return _Hwescrow.Contract.GetPrecision(&_Hwescrow.CallOpts)
}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() pure returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetPrecision() (*big.Int, error) {
	return _Hwescrow.Contract.GetPrecision(&_Hwescrow.CallOpts)
}

// GetRecruiter is a free data retrieval call binding the contract method 0x2d0358c6.
//
// Solidity: function getRecruiter(uint256 _dealId) view returns(address)
func (_Hwescrow *HwescrowCaller) GetRecruiter(opts *bind.CallOpts, _dealId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getRecruiter", _dealId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecruiter is a free data retrieval call binding the contract method 0x2d0358c6.
//
// Solidity: function getRecruiter(uint256 _dealId) view returns(address)
func (_Hwescrow *HwescrowSession) GetRecruiter(_dealId *big.Int) (common.Address, error) {
	return _Hwescrow.Contract.GetRecruiter(&_Hwescrow.CallOpts, _dealId)
}

// GetRecruiter is a free data retrieval call binding the contract method 0x2d0358c6.
//
// Solidity: function getRecruiter(uint256 _dealId) view returns(address)
func (_Hwescrow *HwescrowCallerSession) GetRecruiter(_dealId *big.Int) (common.Address, error) {
	return _Hwescrow.Contract.GetRecruiter(&_Hwescrow.CallOpts, _dealId)
}

// GetRecruiterRating is a free data retrieval call binding the contract method 0xf6f6ed12.
//
// Solidity: function getRecruiterRating(uint256 _dealId) view returns(uint128[])
func (_Hwescrow *HwescrowCaller) GetRecruiterRating(opts *bind.CallOpts, _dealId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getRecruiterRating", _dealId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRecruiterRating is a free data retrieval call binding the contract method 0xf6f6ed12.
//
// Solidity: function getRecruiterRating(uint256 _dealId) view returns(uint128[])
func (_Hwescrow *HwescrowSession) GetRecruiterRating(_dealId *big.Int) ([]*big.Int, error) {
	return _Hwescrow.Contract.GetRecruiterRating(&_Hwescrow.CallOpts, _dealId)
}

// GetRecruiterRating is a free data retrieval call binding the contract method 0xf6f6ed12.
//
// Solidity: function getRecruiterRating(uint256 _dealId) view returns(uint128[])
func (_Hwescrow *HwescrowCallerSession) GetRecruiterRating(_dealId *big.Int) ([]*big.Int, error) {
	return _Hwescrow.Contract.GetRecruiterRating(&_Hwescrow.CallOpts, _dealId)
}

// GetTotalPayment is a free data retrieval call binding the contract method 0x3bb1b2ec.
//
// Solidity: function getTotalPayment(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetTotalPayment(opts *bind.CallOpts, _dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getTotalPayment", _dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPayment is a free data retrieval call binding the contract method 0x3bb1b2ec.
//
// Solidity: function getTotalPayment(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetTotalPayment(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetTotalPayment(&_Hwescrow.CallOpts, _dealId)
}

// GetTotalPayment is a free data retrieval call binding the contract method 0x3bb1b2ec.
//
// Solidity: function getTotalPayment(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetTotalPayment(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetTotalPayment(&_Hwescrow.CallOpts, _dealId)
}

// GetTotalSuccessFee is a free data retrieval call binding the contract method 0x45c60359.
//
// Solidity: function getTotalSuccessFee() view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetTotalSuccessFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getTotalSuccessFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSuccessFee is a free data retrieval call binding the contract method 0x45c60359.
//
// Solidity: function getTotalSuccessFee() view returns(uint256)
func (_Hwescrow *HwescrowSession) GetTotalSuccessFee() (*big.Int, error) {
	return _Hwescrow.Contract.GetTotalSuccessFee(&_Hwescrow.CallOpts)
}

// GetTotalSuccessFee is a free data retrieval call binding the contract method 0x45c60359.
//
// Solidity: function getTotalSuccessFee() view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetTotalSuccessFee() (*big.Int, error) {
	return _Hwescrow.Contract.GetTotalSuccessFee(&_Hwescrow.CallOpts)
}

// GetclaimedAmount is a free data retrieval call binding the contract method 0x709556df.
//
// Solidity: function getclaimedAmount(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCaller) GetclaimedAmount(opts *bind.CallOpts, _dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "getclaimedAmount", _dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetclaimedAmount is a free data retrieval call binding the contract method 0x709556df.
//
// Solidity: function getclaimedAmount(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowSession) GetclaimedAmount(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetclaimedAmount(&_Hwescrow.CallOpts, _dealId)
}

// GetclaimedAmount is a free data retrieval call binding the contract method 0x709556df.
//
// Solidity: function getclaimedAmount(uint256 _dealId) view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) GetclaimedAmount(_dealId *big.Int) (*big.Int, error) {
	return _Hwescrow.Contract.GetclaimedAmount(&_Hwescrow.CallOpts, _dealId)
}

// HonestWorkSuccessFee is a free data retrieval call binding the contract method 0x72a1ade2.
//
// Solidity: function honestWorkSuccessFee() view returns(uint128)
func (_Hwescrow *HwescrowCaller) HonestWorkSuccessFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "honestWorkSuccessFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HonestWorkSuccessFee is a free data retrieval call binding the contract method 0x72a1ade2.
//
// Solidity: function honestWorkSuccessFee() view returns(uint128)
func (_Hwescrow *HwescrowSession) HonestWorkSuccessFee() (*big.Int, error) {
	return _Hwescrow.Contract.HonestWorkSuccessFee(&_Hwescrow.CallOpts)
}

// HonestWorkSuccessFee is a free data retrieval call binding the contract method 0x72a1ade2.
//
// Solidity: function honestWorkSuccessFee() view returns(uint128)
func (_Hwescrow *HwescrowCallerSession) HonestWorkSuccessFee() (*big.Int, error) {
	return _Hwescrow.Contract.HonestWorkSuccessFee(&_Hwescrow.CallOpts)
}

// NativePaymentAllowed is a free data retrieval call binding the contract method 0x7709807b.
//
// Solidity: function nativePaymentAllowed() view returns(bool)
func (_Hwescrow *HwescrowCaller) NativePaymentAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "nativePaymentAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NativePaymentAllowed is a free data retrieval call binding the contract method 0x7709807b.
//
// Solidity: function nativePaymentAllowed() view returns(bool)
func (_Hwescrow *HwescrowSession) NativePaymentAllowed() (bool, error) {
	return _Hwescrow.Contract.NativePaymentAllowed(&_Hwescrow.CallOpts)
}

// NativePaymentAllowed is a free data retrieval call binding the contract method 0x7709807b.
//
// Solidity: function nativePaymentAllowed() view returns(bool)
func (_Hwescrow *HwescrowCallerSession) NativePaymentAllowed() (bool, error) {
	return _Hwescrow.Contract.NativePaymentAllowed(&_Hwescrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hwescrow *HwescrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hwescrow *HwescrowSession) Owner() (common.Address, error) {
	return _Hwescrow.Contract.Owner(&_Hwescrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hwescrow *HwescrowCallerSession) Owner() (common.Address, error) {
	return _Hwescrow.Contract.Owner(&_Hwescrow.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Hwescrow *HwescrowCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Hwescrow *HwescrowSession) Pool() (common.Address, error) {
	return _Hwescrow.Contract.Pool(&_Hwescrow.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Hwescrow *HwescrowCallerSession) Pool() (common.Address, error) {
	return _Hwescrow.Contract.Pool(&_Hwescrow.CallOpts)
}

// RecoverSigner is a free data retrieval call binding the contract method 0xd45167d0.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Hwescrow *HwescrowCaller) RecoverSigner(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "recoverSigner", _ethSignedMessageHash, v, r, s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0xd45167d0.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Hwescrow *HwescrowSession) RecoverSigner(_ethSignedMessageHash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Hwescrow.Contract.RecoverSigner(&_Hwescrow.CallOpts, _ethSignedMessageHash, v, r, s)
}

// RecoverSigner is a free data retrieval call binding the contract method 0xd45167d0.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Hwescrow *HwescrowCallerSession) RecoverSigner(_ethSignedMessageHash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Hwescrow.Contract.RecoverSigner(&_Hwescrow.CallOpts, _ethSignedMessageHash, v, r, s)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Hwescrow *HwescrowCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Hwescrow *HwescrowSession) Registry() (common.Address, error) {
	return _Hwescrow.Contract.Registry(&_Hwescrow.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Hwescrow *HwescrowCallerSession) Registry() (common.Address, error) {
	return _Hwescrow.Contract.Registry(&_Hwescrow.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Hwescrow *HwescrowCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Hwescrow *HwescrowSession) Router() (common.Address, error) {
	return _Hwescrow.Contract.Router(&_Hwescrow.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Hwescrow *HwescrowCallerSession) Router() (common.Address, error) {
	return _Hwescrow.Contract.Router(&_Hwescrow.CallOpts)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Hwescrow *HwescrowCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "splitSignature", sig)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
		V uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.V = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Hwescrow *HwescrowSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Hwescrow.Contract.SplitSignature(&_Hwescrow.CallOpts, sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Hwescrow *HwescrowCallerSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Hwescrow.Contract.SplitSignature(&_Hwescrow.CallOpts, sig)
}

// StableCoin is a free data retrieval call binding the contract method 0x992642e5.
//
// Solidity: function stableCoin() view returns(address)
func (_Hwescrow *HwescrowCaller) StableCoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "stableCoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StableCoin is a free data retrieval call binding the contract method 0x992642e5.
//
// Solidity: function stableCoin() view returns(address)
func (_Hwescrow *HwescrowSession) StableCoin() (common.Address, error) {
	return _Hwescrow.Contract.StableCoin(&_Hwescrow.CallOpts)
}

// StableCoin is a free data retrieval call binding the contract method 0x992642e5.
//
// Solidity: function stableCoin() view returns(address)
func (_Hwescrow *HwescrowCallerSession) StableCoin() (common.Address, error) {
	return _Hwescrow.Contract.StableCoin(&_Hwescrow.CallOpts)
}

// TotalCollectedSuccessFee is a free data retrieval call binding the contract method 0xa19b94f8.
//
// Solidity: function totalCollectedSuccessFee() view returns(uint256)
func (_Hwescrow *HwescrowCaller) TotalCollectedSuccessFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hwescrow.contract.Call(opts, &out, "totalCollectedSuccessFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollectedSuccessFee is a free data retrieval call binding the contract method 0xa19b94f8.
//
// Solidity: function totalCollectedSuccessFee() view returns(uint256)
func (_Hwescrow *HwescrowSession) TotalCollectedSuccessFee() (*big.Int, error) {
	return _Hwescrow.Contract.TotalCollectedSuccessFee(&_Hwescrow.CallOpts)
}

// TotalCollectedSuccessFee is a free data retrieval call binding the contract method 0xa19b94f8.
//
// Solidity: function totalCollectedSuccessFee() view returns(uint256)
func (_Hwescrow *HwescrowCallerSession) TotalCollectedSuccessFee() (*big.Int, error) {
	return _Hwescrow.Contract.TotalCollectedSuccessFee(&_Hwescrow.CallOpts)
}

// AdditionalPayment is a paid mutator transaction binding the contract method 0xac04b3b3.
//
// Solidity: function additionalPayment(uint256 _dealId, uint256 _payment, uint256 _recruiterNFT, uint128 _rating) payable returns()
func (_Hwescrow *HwescrowTransactor) AdditionalPayment(opts *bind.TransactOpts, _dealId *big.Int, _payment *big.Int, _recruiterNFT *big.Int, _rating *big.Int) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "additionalPayment", _dealId, _payment, _recruiterNFT, _rating)
}

// AdditionalPayment is a paid mutator transaction binding the contract method 0xac04b3b3.
//
// Solidity: function additionalPayment(uint256 _dealId, uint256 _payment, uint256 _recruiterNFT, uint128 _rating) payable returns()
func (_Hwescrow *HwescrowSession) AdditionalPayment(_dealId *big.Int, _payment *big.Int, _recruiterNFT *big.Int, _rating *big.Int) (*types.Transaction, error) {
	return _Hwescrow.Contract.AdditionalPayment(&_Hwescrow.TransactOpts, _dealId, _payment, _recruiterNFT, _rating)
}

// AdditionalPayment is a paid mutator transaction binding the contract method 0xac04b3b3.
//
// Solidity: function additionalPayment(uint256 _dealId, uint256 _payment, uint256 _recruiterNFT, uint128 _rating) payable returns()
func (_Hwescrow *HwescrowTransactorSession) AdditionalPayment(_dealId *big.Int, _payment *big.Int, _recruiterNFT *big.Int, _rating *big.Int) (*types.Transaction, error) {
	return _Hwescrow.Contract.AdditionalPayment(&_Hwescrow.TransactOpts, _dealId, _payment, _recruiterNFT, _rating)
}

// AllowNativePayment is a paid mutator transaction binding the contract method 0x90376302.
//
// Solidity: function allowNativePayment(bool _bool) returns()
func (_Hwescrow *HwescrowTransactor) AllowNativePayment(opts *bind.TransactOpts, _bool bool) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "allowNativePayment", _bool)
}

// AllowNativePayment is a paid mutator transaction binding the contract method 0x90376302.
//
// Solidity: function allowNativePayment(bool _bool) returns()
func (_Hwescrow *HwescrowSession) AllowNativePayment(_bool bool) (*types.Transaction, error) {
	return _Hwescrow.Contract.AllowNativePayment(&_Hwescrow.TransactOpts, _bool)
}

// AllowNativePayment is a paid mutator transaction binding the contract method 0x90376302.
//
// Solidity: function allowNativePayment(bool _bool) returns()
func (_Hwescrow *HwescrowTransactorSession) AllowNativePayment(_bool bool) (*types.Transaction, error) {
	return _Hwescrow.Contract.AllowNativePayment(&_Hwescrow.TransactOpts, _bool)
}

// ChangeExtraPaymentLimit is a paid mutator transaction binding the contract method 0x89c278d0.
//
// Solidity: function changeExtraPaymentLimit(uint64 _limit) returns()
func (_Hwescrow *HwescrowTransactor) ChangeExtraPaymentLimit(opts *bind.TransactOpts, _limit uint64) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "changeExtraPaymentLimit", _limit)
}

// ChangeExtraPaymentLimit is a paid mutator transaction binding the contract method 0x89c278d0.
//
// Solidity: function changeExtraPaymentLimit(uint64 _limit) returns()
func (_Hwescrow *HwescrowSession) ChangeExtraPaymentLimit(_limit uint64) (*types.Transaction, error) {
	return _Hwescrow.Contract.ChangeExtraPaymentLimit(&_Hwescrow.TransactOpts, _limit)
}

// ChangeExtraPaymentLimit is a paid mutator transaction binding the contract method 0x89c278d0.
//
// Solidity: function changeExtraPaymentLimit(uint64 _limit) returns()
func (_Hwescrow *HwescrowTransactorSession) ChangeExtraPaymentLimit(_limit uint64) (*types.Transaction, error) {
	return _Hwescrow.Contract.ChangeExtraPaymentLimit(&_Hwescrow.TransactOpts, _limit)
}

// ChangeRegistry is a paid mutator transaction binding the contract method 0x15554c55.
//
// Solidity: function changeRegistry(address _registry) returns()
func (_Hwescrow *HwescrowTransactor) ChangeRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "changeRegistry", _registry)
}

// ChangeRegistry is a paid mutator transaction binding the contract method 0x15554c55.
//
// Solidity: function changeRegistry(address _registry) returns()
func (_Hwescrow *HwescrowSession) ChangeRegistry(_registry common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.ChangeRegistry(&_Hwescrow.TransactOpts, _registry)
}

// ChangeRegistry is a paid mutator transaction binding the contract method 0x15554c55.
//
// Solidity: function changeRegistry(address _registry) returns()
func (_Hwescrow *HwescrowTransactorSession) ChangeRegistry(_registry common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.ChangeRegistry(&_Hwescrow.TransactOpts, _registry)
}

// ChangeSuccessFee is a paid mutator transaction binding the contract method 0x635b0af7.
//
// Solidity: function changeSuccessFee(uint128 _fee) returns()
func (_Hwescrow *HwescrowTransactor) ChangeSuccessFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "changeSuccessFee", _fee)
}

// ChangeSuccessFee is a paid mutator transaction binding the contract method 0x635b0af7.
//
// Solidity: function changeSuccessFee(uint128 _fee) returns()
func (_Hwescrow *HwescrowSession) ChangeSuccessFee(_fee *big.Int) (*types.Transaction, error) {
	return _Hwescrow.Contract.ChangeSuccessFee(&_Hwescrow.TransactOpts, _fee)
}

// ChangeSuccessFee is a paid mutator transaction binding the contract method 0x635b0af7.
//
// Solidity: function changeSuccessFee(uint128 _fee) returns()
func (_Hwescrow *HwescrowTransactorSession) ChangeSuccessFee(_fee *big.Int) (*types.Transaction, error) {
	return _Hwescrow.Contract.ChangeSuccessFee(&_Hwescrow.TransactOpts, _fee)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0x885e2091.
//
// Solidity: function claimPayment(uint256 _dealId, uint256 _withdrawAmount, uint128 _rating, uint256 _creatorNFT) returns()
func (_Hwescrow *HwescrowTransactor) ClaimPayment(opts *bind.TransactOpts, _dealId *big.Int, _withdrawAmount *big.Int, _rating *big.Int, _creatorNFT *big.Int) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "claimPayment", _dealId, _withdrawAmount, _rating, _creatorNFT)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0x885e2091.
//
// Solidity: function claimPayment(uint256 _dealId, uint256 _withdrawAmount, uint128 _rating, uint256 _creatorNFT) returns()
func (_Hwescrow *HwescrowSession) ClaimPayment(_dealId *big.Int, _withdrawAmount *big.Int, _rating *big.Int, _creatorNFT *big.Int) (*types.Transaction, error) {
	return _Hwescrow.Contract.ClaimPayment(&_Hwescrow.TransactOpts, _dealId, _withdrawAmount, _rating, _creatorNFT)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0x885e2091.
//
// Solidity: function claimPayment(uint256 _dealId, uint256 _withdrawAmount, uint128 _rating, uint256 _creatorNFT) returns()
func (_Hwescrow *HwescrowTransactorSession) ClaimPayment(_dealId *big.Int, _withdrawAmount *big.Int, _rating *big.Int, _creatorNFT *big.Int) (*types.Transaction, error) {
	return _Hwescrow.Contract.ClaimPayment(&_Hwescrow.TransactOpts, _dealId, _withdrawAmount, _rating, _creatorNFT)
}

// ClaimSuccessFee is a paid mutator transaction binding the contract method 0xa1d16c49.
//
// Solidity: function claimSuccessFee(uint256 _dealId, address _feeCollector) returns()
func (_Hwescrow *HwescrowTransactor) ClaimSuccessFee(opts *bind.TransactOpts, _dealId *big.Int, _feeCollector common.Address) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "claimSuccessFee", _dealId, _feeCollector)
}

// ClaimSuccessFee is a paid mutator transaction binding the contract method 0xa1d16c49.
//
// Solidity: function claimSuccessFee(uint256 _dealId, address _feeCollector) returns()
func (_Hwescrow *HwescrowSession) ClaimSuccessFee(_dealId *big.Int, _feeCollector common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.ClaimSuccessFee(&_Hwescrow.TransactOpts, _dealId, _feeCollector)
}

// ClaimSuccessFee is a paid mutator transaction binding the contract method 0xa1d16c49.
//
// Solidity: function claimSuccessFee(uint256 _dealId, address _feeCollector) returns()
func (_Hwescrow *HwescrowTransactorSession) ClaimSuccessFee(_dealId *big.Int, _feeCollector common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.ClaimSuccessFee(&_Hwescrow.TransactOpts, _dealId, _feeCollector)
}

// ClaimTotalSuccessFee is a paid mutator transaction binding the contract method 0x5da9c898.
//
// Solidity: function claimTotalSuccessFee(address _feeCollector) returns()
func (_Hwescrow *HwescrowTransactor) ClaimTotalSuccessFee(opts *bind.TransactOpts, _feeCollector common.Address) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "claimTotalSuccessFee", _feeCollector)
}

// ClaimTotalSuccessFee is a paid mutator transaction binding the contract method 0x5da9c898.
//
// Solidity: function claimTotalSuccessFee(address _feeCollector) returns()
func (_Hwescrow *HwescrowSession) ClaimTotalSuccessFee(_feeCollector common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.ClaimTotalSuccessFee(&_Hwescrow.TransactOpts, _feeCollector)
}

// ClaimTotalSuccessFee is a paid mutator transaction binding the contract method 0x5da9c898.
//
// Solidity: function claimTotalSuccessFee(address _feeCollector) returns()
func (_Hwescrow *HwescrowTransactorSession) ClaimTotalSuccessFee(_feeCollector common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.ClaimTotalSuccessFee(&_Hwescrow.TransactOpts, _feeCollector)
}

// CreateDeal is a paid mutator transaction binding the contract method 0x69e4f998.
//
// Solidity: function createDeal(address _recruiter, address _creator, address _paymentToken, uint256 _totalPayment, uint256 _downPayment, uint256 _recruiterNFTId, uint8 v, bytes32 r, bytes32 s) payable returns(uint256)
func (_Hwescrow *HwescrowTransactor) CreateDeal(opts *bind.TransactOpts, _recruiter common.Address, _creator common.Address, _paymentToken common.Address, _totalPayment *big.Int, _downPayment *big.Int, _recruiterNFTId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "createDeal", _recruiter, _creator, _paymentToken, _totalPayment, _downPayment, _recruiterNFTId, v, r, s)
}

// CreateDeal is a paid mutator transaction binding the contract method 0x69e4f998.
//
// Solidity: function createDeal(address _recruiter, address _creator, address _paymentToken, uint256 _totalPayment, uint256 _downPayment, uint256 _recruiterNFTId, uint8 v, bytes32 r, bytes32 s) payable returns(uint256)
func (_Hwescrow *HwescrowSession) CreateDeal(_recruiter common.Address, _creator common.Address, _paymentToken common.Address, _totalPayment *big.Int, _downPayment *big.Int, _recruiterNFTId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Hwescrow.Contract.CreateDeal(&_Hwescrow.TransactOpts, _recruiter, _creator, _paymentToken, _totalPayment, _downPayment, _recruiterNFTId, v, r, s)
}

// CreateDeal is a paid mutator transaction binding the contract method 0x69e4f998.
//
// Solidity: function createDeal(address _recruiter, address _creator, address _paymentToken, uint256 _totalPayment, uint256 _downPayment, uint256 _recruiterNFTId, uint8 v, bytes32 r, bytes32 s) payable returns(uint256)
func (_Hwescrow *HwescrowTransactorSession) CreateDeal(_recruiter common.Address, _creator common.Address, _paymentToken common.Address, _totalPayment *big.Int, _downPayment *big.Int, _recruiterNFTId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Hwescrow.Contract.CreateDeal(&_Hwescrow.TransactOpts, _recruiter, _creator, _paymentToken, _totalPayment, _downPayment, _recruiterNFTId, v, r, s)
}

// CreateDealSignature is a paid mutator transaction binding the contract method 0x2cb7ad6f.
//
// Solidity: function createDealSignature(address _recruiter, address _creator, address _paymentToken, uint256 _totalPayment, uint256 _downPayment, uint256 _recruiterNFTId, bytes _signature) payable returns(uint256)
func (_Hwescrow *HwescrowTransactor) CreateDealSignature(opts *bind.TransactOpts, _recruiter common.Address, _creator common.Address, _paymentToken common.Address, _totalPayment *big.Int, _downPayment *big.Int, _recruiterNFTId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "createDealSignature", _recruiter, _creator, _paymentToken, _totalPayment, _downPayment, _recruiterNFTId, _signature)
}

// CreateDealSignature is a paid mutator transaction binding the contract method 0x2cb7ad6f.
//
// Solidity: function createDealSignature(address _recruiter, address _creator, address _paymentToken, uint256 _totalPayment, uint256 _downPayment, uint256 _recruiterNFTId, bytes _signature) payable returns(uint256)
func (_Hwescrow *HwescrowSession) CreateDealSignature(_recruiter common.Address, _creator common.Address, _paymentToken common.Address, _totalPayment *big.Int, _downPayment *big.Int, _recruiterNFTId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Hwescrow.Contract.CreateDealSignature(&_Hwescrow.TransactOpts, _recruiter, _creator, _paymentToken, _totalPayment, _downPayment, _recruiterNFTId, _signature)
}

// CreateDealSignature is a paid mutator transaction binding the contract method 0x2cb7ad6f.
//
// Solidity: function createDealSignature(address _recruiter, address _creator, address _paymentToken, uint256 _totalPayment, uint256 _downPayment, uint256 _recruiterNFTId, bytes _signature) payable returns(uint256)
func (_Hwescrow *HwescrowTransactorSession) CreateDealSignature(_recruiter common.Address, _creator common.Address, _paymentToken common.Address, _totalPayment *big.Int, _downPayment *big.Int, _recruiterNFTId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Hwescrow.Contract.CreateDealSignature(&_Hwescrow.TransactOpts, _recruiter, _creator, _paymentToken, _totalPayment, _downPayment, _recruiterNFTId, _signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hwescrow *HwescrowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hwescrow *HwescrowSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hwescrow.Contract.RenounceOwnership(&_Hwescrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hwescrow *HwescrowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hwescrow.Contract.RenounceOwnership(&_Hwescrow.TransactOpts)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _pool) returns()
func (_Hwescrow *HwescrowTransactor) SetPool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "setPool", _pool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _pool) returns()
func (_Hwescrow *HwescrowSession) SetPool(_pool common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.SetPool(&_Hwescrow.TransactOpts, _pool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _pool) returns()
func (_Hwescrow *HwescrowTransactorSession) SetPool(_pool common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.SetPool(&_Hwescrow.TransactOpts, _pool)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_Hwescrow *HwescrowTransactor) SetRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "setRouter", _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_Hwescrow *HwescrowSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.SetRouter(&_Hwescrow.TransactOpts, _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_Hwescrow *HwescrowTransactorSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.SetRouter(&_Hwescrow.TransactOpts, _router)
}

// SetStableCoin is a paid mutator transaction binding the contract method 0x23af4e17.
//
// Solidity: function setStableCoin(address _stableCoin) returns()
func (_Hwescrow *HwescrowTransactor) SetStableCoin(opts *bind.TransactOpts, _stableCoin common.Address) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "setStableCoin", _stableCoin)
}

// SetStableCoin is a paid mutator transaction binding the contract method 0x23af4e17.
//
// Solidity: function setStableCoin(address _stableCoin) returns()
func (_Hwescrow *HwescrowSession) SetStableCoin(_stableCoin common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.SetStableCoin(&_Hwescrow.TransactOpts, _stableCoin)
}

// SetStableCoin is a paid mutator transaction binding the contract method 0x23af4e17.
//
// Solidity: function setStableCoin(address _stableCoin) returns()
func (_Hwescrow *HwescrowTransactorSession) SetStableCoin(_stableCoin common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.SetStableCoin(&_Hwescrow.TransactOpts, _stableCoin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hwescrow *HwescrowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hwescrow *HwescrowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.TransferOwnership(&_Hwescrow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hwescrow *HwescrowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hwescrow.Contract.TransferOwnership(&_Hwescrow.TransactOpts, newOwner)
}

// UnlockPayment is a paid mutator transaction binding the contract method 0x6f2a1562.
//
// Solidity: function unlockPayment(uint256 _dealId, uint256 _paymentAmount, uint128 _rating, uint256 _recruiterNFT) returns()
func (_Hwescrow *HwescrowTransactor) UnlockPayment(opts *bind.TransactOpts, _dealId *big.Int, _paymentAmount *big.Int, _rating *big.Int, _recruiterNFT *big.Int) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "unlockPayment", _dealId, _paymentAmount, _rating, _recruiterNFT)
}

// UnlockPayment is a paid mutator transaction binding the contract method 0x6f2a1562.
//
// Solidity: function unlockPayment(uint256 _dealId, uint256 _paymentAmount, uint128 _rating, uint256 _recruiterNFT) returns()
func (_Hwescrow *HwescrowSession) UnlockPayment(_dealId *big.Int, _paymentAmount *big.Int, _rating *big.Int, _recruiterNFT *big.Int) (*types.Transaction, error) {
	return _Hwescrow.Contract.UnlockPayment(&_Hwescrow.TransactOpts, _dealId, _paymentAmount, _rating, _recruiterNFT)
}

// UnlockPayment is a paid mutator transaction binding the contract method 0x6f2a1562.
//
// Solidity: function unlockPayment(uint256 _dealId, uint256 _paymentAmount, uint128 _rating, uint256 _recruiterNFT) returns()
func (_Hwescrow *HwescrowTransactorSession) UnlockPayment(_dealId *big.Int, _paymentAmount *big.Int, _rating *big.Int, _recruiterNFT *big.Int) (*types.Transaction, error) {
	return _Hwescrow.Contract.UnlockPayment(&_Hwescrow.TransactOpts, _dealId, _paymentAmount, _rating, _recruiterNFT)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0xdbd3cd62.
//
// Solidity: function withdrawPayment(uint256 _dealId) returns()
func (_Hwescrow *HwescrowTransactor) WithdrawPayment(opts *bind.TransactOpts, _dealId *big.Int) (*types.Transaction, error) {
	return _Hwescrow.contract.Transact(opts, "withdrawPayment", _dealId)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0xdbd3cd62.
//
// Solidity: function withdrawPayment(uint256 _dealId) returns()
func (_Hwescrow *HwescrowSession) WithdrawPayment(_dealId *big.Int) (*types.Transaction, error) {
	return _Hwescrow.Contract.WithdrawPayment(&_Hwescrow.TransactOpts, _dealId)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0xdbd3cd62.
//
// Solidity: function withdrawPayment(uint256 _dealId) returns()
func (_Hwescrow *HwescrowTransactorSession) WithdrawPayment(_dealId *big.Int) (*types.Transaction, error) {
	return _Hwescrow.Contract.WithdrawPayment(&_Hwescrow.TransactOpts, _dealId)
}

// HwescrowAdditionalPaymentIterator is returned from FilterAdditionalPayment and is used to iterate over the raw logs and unpacked data for AdditionalPayment events raised by the Hwescrow contract.
type HwescrowAdditionalPaymentIterator struct {
	Event *HwescrowAdditionalPayment // Event containing the contract specifics and raw log

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
func (it *HwescrowAdditionalPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwescrowAdditionalPayment)
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
		it.Event = new(HwescrowAdditionalPayment)
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
func (it *HwescrowAdditionalPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwescrowAdditionalPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwescrowAdditionalPayment represents a AdditionalPayment event raised by the Hwescrow contract.
type HwescrowAdditionalPayment struct {
	DealId    *big.Int
	Recruiter common.Address
	Payment   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAdditionalPayment is a free log retrieval operation binding the contract event 0x19c66b05a0a4d2be41b123b9d0daa0670bb8f1317b3745200243a81751b07f43.
//
// Solidity: event AdditionalPayment(uint256 indexed _dealId, address indexed _recruiter, uint256 indexed _payment)
func (_Hwescrow *HwescrowFilterer) FilterAdditionalPayment(opts *bind.FilterOpts, _dealId []*big.Int, _recruiter []common.Address, _payment []*big.Int) (*HwescrowAdditionalPaymentIterator, error) {

	var _dealIdRule []interface{}
	for _, _dealIdItem := range _dealId {
		_dealIdRule = append(_dealIdRule, _dealIdItem)
	}
	var _recruiterRule []interface{}
	for _, _recruiterItem := range _recruiter {
		_recruiterRule = append(_recruiterRule, _recruiterItem)
	}
	var _paymentRule []interface{}
	for _, _paymentItem := range _payment {
		_paymentRule = append(_paymentRule, _paymentItem)
	}

	logs, sub, err := _Hwescrow.contract.FilterLogs(opts, "AdditionalPayment", _dealIdRule, _recruiterRule, _paymentRule)
	if err != nil {
		return nil, err
	}
	return &HwescrowAdditionalPaymentIterator{contract: _Hwescrow.contract, event: "AdditionalPayment", logs: logs, sub: sub}, nil
}

// WatchAdditionalPayment is a free log subscription operation binding the contract event 0x19c66b05a0a4d2be41b123b9d0daa0670bb8f1317b3745200243a81751b07f43.
//
// Solidity: event AdditionalPayment(uint256 indexed _dealId, address indexed _recruiter, uint256 indexed _payment)
func (_Hwescrow *HwescrowFilterer) WatchAdditionalPayment(opts *bind.WatchOpts, sink chan<- *HwescrowAdditionalPayment, _dealId []*big.Int, _recruiter []common.Address, _payment []*big.Int) (event.Subscription, error) {

	var _dealIdRule []interface{}
	for _, _dealIdItem := range _dealId {
		_dealIdRule = append(_dealIdRule, _dealIdItem)
	}
	var _recruiterRule []interface{}
	for _, _recruiterItem := range _recruiter {
		_recruiterRule = append(_recruiterRule, _recruiterItem)
	}
	var _paymentRule []interface{}
	for _, _paymentItem := range _payment {
		_paymentRule = append(_paymentRule, _paymentItem)
	}

	logs, sub, err := _Hwescrow.contract.WatchLogs(opts, "AdditionalPayment", _dealIdRule, _recruiterRule, _paymentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwescrowAdditionalPayment)
				if err := _Hwescrow.contract.UnpackLog(event, "AdditionalPayment", log); err != nil {
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

// ParseAdditionalPayment is a log parse operation binding the contract event 0x19c66b05a0a4d2be41b123b9d0daa0670bb8f1317b3745200243a81751b07f43.
//
// Solidity: event AdditionalPayment(uint256 indexed _dealId, address indexed _recruiter, uint256 indexed _payment)
func (_Hwescrow *HwescrowFilterer) ParseAdditionalPayment(log types.Log) (*HwescrowAdditionalPayment, error) {
	event := new(HwescrowAdditionalPayment)
	if err := _Hwescrow.contract.UnpackLog(event, "AdditionalPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwescrowExtraLimitChangedIterator is returned from FilterExtraLimitChanged and is used to iterate over the raw logs and unpacked data for ExtraLimitChanged events raised by the Hwescrow contract.
type HwescrowExtraLimitChangedIterator struct {
	Event *HwescrowExtraLimitChanged // Event containing the contract specifics and raw log

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
func (it *HwescrowExtraLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwescrowExtraLimitChanged)
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
		it.Event = new(HwescrowExtraLimitChanged)
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
func (it *HwescrowExtraLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwescrowExtraLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwescrowExtraLimitChanged represents a ExtraLimitChanged event raised by the Hwescrow contract.
type HwescrowExtraLimitChanged struct {
	NewPaymentLimit *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExtraLimitChanged is a free log retrieval operation binding the contract event 0x0527b3df1c258a1e5e5df0f62af25e3b367b14525e6bbf79abb69664f6fbdd7b.
//
// Solidity: event ExtraLimitChanged(uint256 _newPaymentLimit)
func (_Hwescrow *HwescrowFilterer) FilterExtraLimitChanged(opts *bind.FilterOpts) (*HwescrowExtraLimitChangedIterator, error) {

	logs, sub, err := _Hwescrow.contract.FilterLogs(opts, "ExtraLimitChanged")
	if err != nil {
		return nil, err
	}
	return &HwescrowExtraLimitChangedIterator{contract: _Hwescrow.contract, event: "ExtraLimitChanged", logs: logs, sub: sub}, nil
}

// WatchExtraLimitChanged is a free log subscription operation binding the contract event 0x0527b3df1c258a1e5e5df0f62af25e3b367b14525e6bbf79abb69664f6fbdd7b.
//
// Solidity: event ExtraLimitChanged(uint256 _newPaymentLimit)
func (_Hwescrow *HwescrowFilterer) WatchExtraLimitChanged(opts *bind.WatchOpts, sink chan<- *HwescrowExtraLimitChanged) (event.Subscription, error) {

	logs, sub, err := _Hwescrow.contract.WatchLogs(opts, "ExtraLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwescrowExtraLimitChanged)
				if err := _Hwescrow.contract.UnpackLog(event, "ExtraLimitChanged", log); err != nil {
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

// ParseExtraLimitChanged is a log parse operation binding the contract event 0x0527b3df1c258a1e5e5df0f62af25e3b367b14525e6bbf79abb69664f6fbdd7b.
//
// Solidity: event ExtraLimitChanged(uint256 _newPaymentLimit)
func (_Hwescrow *HwescrowFilterer) ParseExtraLimitChanged(log types.Log) (*HwescrowExtraLimitChanged, error) {
	event := new(HwescrowExtraLimitChanged)
	if err := _Hwescrow.contract.UnpackLog(event, "ExtraLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwescrowFeeChangedIterator is returned from FilterFeeChanged and is used to iterate over the raw logs and unpacked data for FeeChanged events raised by the Hwescrow contract.
type HwescrowFeeChangedIterator struct {
	Event *HwescrowFeeChanged // Event containing the contract specifics and raw log

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
func (it *HwescrowFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwescrowFeeChanged)
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
		it.Event = new(HwescrowFeeChanged)
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
func (it *HwescrowFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwescrowFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwescrowFeeChanged represents a FeeChanged event raised by the Hwescrow contract.
type HwescrowFeeChanged struct {
	NewSuccessFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeChanged is a free log retrieval operation binding the contract event 0x6bbc57480a46553fa4d156ce702beef5f3ad66303b0ed1a5d4cb44966c6584c3.
//
// Solidity: event FeeChanged(uint256 _newSuccessFee)
func (_Hwescrow *HwescrowFilterer) FilterFeeChanged(opts *bind.FilterOpts) (*HwescrowFeeChangedIterator, error) {

	logs, sub, err := _Hwescrow.contract.FilterLogs(opts, "FeeChanged")
	if err != nil {
		return nil, err
	}
	return &HwescrowFeeChangedIterator{contract: _Hwescrow.contract, event: "FeeChanged", logs: logs, sub: sub}, nil
}

// WatchFeeChanged is a free log subscription operation binding the contract event 0x6bbc57480a46553fa4d156ce702beef5f3ad66303b0ed1a5d4cb44966c6584c3.
//
// Solidity: event FeeChanged(uint256 _newSuccessFee)
func (_Hwescrow *HwescrowFilterer) WatchFeeChanged(opts *bind.WatchOpts, sink chan<- *HwescrowFeeChanged) (event.Subscription, error) {

	logs, sub, err := _Hwescrow.contract.WatchLogs(opts, "FeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwescrowFeeChanged)
				if err := _Hwescrow.contract.UnpackLog(event, "FeeChanged", log); err != nil {
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

// ParseFeeChanged is a log parse operation binding the contract event 0x6bbc57480a46553fa4d156ce702beef5f3ad66303b0ed1a5d4cb44966c6584c3.
//
// Solidity: event FeeChanged(uint256 _newSuccessFee)
func (_Hwescrow *HwescrowFilterer) ParseFeeChanged(log types.Log) (*HwescrowFeeChanged, error) {
	event := new(HwescrowFeeChanged)
	if err := _Hwescrow.contract.UnpackLog(event, "FeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwescrowFeeClaimedIterator is returned from FilterFeeClaimed and is used to iterate over the raw logs and unpacked data for FeeClaimed events raised by the Hwescrow contract.
type HwescrowFeeClaimedIterator struct {
	Event *HwescrowFeeClaimed // Event containing the contract specifics and raw log

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
func (it *HwescrowFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwescrowFeeClaimed)
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
		it.Event = new(HwescrowFeeClaimed)
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
func (it *HwescrowFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwescrowFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwescrowFeeClaimed represents a FeeClaimed event raised by the Hwescrow contract.
type HwescrowFeeClaimed struct {
	DealId *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeClaimed is a free log retrieval operation binding the contract event 0x8f50c68aa31577462ee7e15eb3417fa4442a7167ecb21ddba35eb88926d72689.
//
// Solidity: event FeeClaimed(uint256 indexed _dealId, uint256 _amount)
func (_Hwescrow *HwescrowFilterer) FilterFeeClaimed(opts *bind.FilterOpts, _dealId []*big.Int) (*HwescrowFeeClaimedIterator, error) {

	var _dealIdRule []interface{}
	for _, _dealIdItem := range _dealId {
		_dealIdRule = append(_dealIdRule, _dealIdItem)
	}

	logs, sub, err := _Hwescrow.contract.FilterLogs(opts, "FeeClaimed", _dealIdRule)
	if err != nil {
		return nil, err
	}
	return &HwescrowFeeClaimedIterator{contract: _Hwescrow.contract, event: "FeeClaimed", logs: logs, sub: sub}, nil
}

// WatchFeeClaimed is a free log subscription operation binding the contract event 0x8f50c68aa31577462ee7e15eb3417fa4442a7167ecb21ddba35eb88926d72689.
//
// Solidity: event FeeClaimed(uint256 indexed _dealId, uint256 _amount)
func (_Hwescrow *HwescrowFilterer) WatchFeeClaimed(opts *bind.WatchOpts, sink chan<- *HwescrowFeeClaimed, _dealId []*big.Int) (event.Subscription, error) {

	var _dealIdRule []interface{}
	for _, _dealIdItem := range _dealId {
		_dealIdRule = append(_dealIdRule, _dealIdItem)
	}

	logs, sub, err := _Hwescrow.contract.WatchLogs(opts, "FeeClaimed", _dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwescrowFeeClaimed)
				if err := _Hwescrow.contract.UnpackLog(event, "FeeClaimed", log); err != nil {
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

// ParseFeeClaimed is a log parse operation binding the contract event 0x8f50c68aa31577462ee7e15eb3417fa4442a7167ecb21ddba35eb88926d72689.
//
// Solidity: event FeeClaimed(uint256 indexed _dealId, uint256 _amount)
func (_Hwescrow *HwescrowFilterer) ParseFeeClaimed(log types.Log) (*HwescrowFeeClaimed, error) {
	event := new(HwescrowFeeClaimed)
	if err := _Hwescrow.contract.UnpackLog(event, "FeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwescrowGrossRevenueUpdatedIterator is returned from FilterGrossRevenueUpdated and is used to iterate over the raw logs and unpacked data for GrossRevenueUpdated events raised by the Hwescrow contract.
type HwescrowGrossRevenueUpdatedIterator struct {
	Event *HwescrowGrossRevenueUpdated // Event containing the contract specifics and raw log

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
func (it *HwescrowGrossRevenueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwescrowGrossRevenueUpdated)
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
		it.Event = new(HwescrowGrossRevenueUpdated)
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
func (it *HwescrowGrossRevenueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwescrowGrossRevenueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwescrowGrossRevenueUpdated represents a GrossRevenueUpdated event raised by the Hwescrow contract.
type HwescrowGrossRevenueUpdated struct {
	TokenId      *big.Int
	GrossRevenue *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGrossRevenueUpdated is a free log retrieval operation binding the contract event 0x810094c20e68a3de910ca01b2ec16e0f7ba3df4060e865e10a93ddb2bb974102.
//
// Solidity: event GrossRevenueUpdated(uint256 indexed _tokenId, uint256 _grossRevenue)
func (_Hwescrow *HwescrowFilterer) FilterGrossRevenueUpdated(opts *bind.FilterOpts, _tokenId []*big.Int) (*HwescrowGrossRevenueUpdatedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Hwescrow.contract.FilterLogs(opts, "GrossRevenueUpdated", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HwescrowGrossRevenueUpdatedIterator{contract: _Hwescrow.contract, event: "GrossRevenueUpdated", logs: logs, sub: sub}, nil
}

// WatchGrossRevenueUpdated is a free log subscription operation binding the contract event 0x810094c20e68a3de910ca01b2ec16e0f7ba3df4060e865e10a93ddb2bb974102.
//
// Solidity: event GrossRevenueUpdated(uint256 indexed _tokenId, uint256 _grossRevenue)
func (_Hwescrow *HwescrowFilterer) WatchGrossRevenueUpdated(opts *bind.WatchOpts, sink chan<- *HwescrowGrossRevenueUpdated, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Hwescrow.contract.WatchLogs(opts, "GrossRevenueUpdated", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwescrowGrossRevenueUpdated)
				if err := _Hwescrow.contract.UnpackLog(event, "GrossRevenueUpdated", log); err != nil {
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

// ParseGrossRevenueUpdated is a log parse operation binding the contract event 0x810094c20e68a3de910ca01b2ec16e0f7ba3df4060e865e10a93ddb2bb974102.
//
// Solidity: event GrossRevenueUpdated(uint256 indexed _tokenId, uint256 _grossRevenue)
func (_Hwescrow *HwescrowFilterer) ParseGrossRevenueUpdated(log types.Log) (*HwescrowGrossRevenueUpdated, error) {
	event := new(HwescrowGrossRevenueUpdated)
	if err := _Hwescrow.contract.UnpackLog(event, "GrossRevenueUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwescrowOfferCreatedIterator is returned from FilterOfferCreated and is used to iterate over the raw logs and unpacked data for OfferCreated events raised by the Hwescrow contract.
type HwescrowOfferCreatedIterator struct {
	Event *HwescrowOfferCreated // Event containing the contract specifics and raw log

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
func (it *HwescrowOfferCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwescrowOfferCreated)
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
		it.Event = new(HwescrowOfferCreated)
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
func (it *HwescrowOfferCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwescrowOfferCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwescrowOfferCreated represents a OfferCreated event raised by the Hwescrow contract.
type HwescrowOfferCreated struct {
	Recruiter    common.Address
	Creator      common.Address
	TotalPayment *big.Int
	PaymentToken common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOfferCreated is a free log retrieval operation binding the contract event 0x4f05be72ad7f57c27e555ace8452f56d8b1e82c9e6e1cd4fd282f34518b7729a.
//
// Solidity: event OfferCreated(address indexed _recruiter, address indexed _creator, uint256 indexed _totalPayment, address _paymentToken)
func (_Hwescrow *HwescrowFilterer) FilterOfferCreated(opts *bind.FilterOpts, _recruiter []common.Address, _creator []common.Address, _totalPayment []*big.Int) (*HwescrowOfferCreatedIterator, error) {

	var _recruiterRule []interface{}
	for _, _recruiterItem := range _recruiter {
		_recruiterRule = append(_recruiterRule, _recruiterItem)
	}
	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}
	var _totalPaymentRule []interface{}
	for _, _totalPaymentItem := range _totalPayment {
		_totalPaymentRule = append(_totalPaymentRule, _totalPaymentItem)
	}

	logs, sub, err := _Hwescrow.contract.FilterLogs(opts, "OfferCreated", _recruiterRule, _creatorRule, _totalPaymentRule)
	if err != nil {
		return nil, err
	}
	return &HwescrowOfferCreatedIterator{contract: _Hwescrow.contract, event: "OfferCreated", logs: logs, sub: sub}, nil
}

// WatchOfferCreated is a free log subscription operation binding the contract event 0x4f05be72ad7f57c27e555ace8452f56d8b1e82c9e6e1cd4fd282f34518b7729a.
//
// Solidity: event OfferCreated(address indexed _recruiter, address indexed _creator, uint256 indexed _totalPayment, address _paymentToken)
func (_Hwescrow *HwescrowFilterer) WatchOfferCreated(opts *bind.WatchOpts, sink chan<- *HwescrowOfferCreated, _recruiter []common.Address, _creator []common.Address, _totalPayment []*big.Int) (event.Subscription, error) {

	var _recruiterRule []interface{}
	for _, _recruiterItem := range _recruiter {
		_recruiterRule = append(_recruiterRule, _recruiterItem)
	}
	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}
	var _totalPaymentRule []interface{}
	for _, _totalPaymentItem := range _totalPayment {
		_totalPaymentRule = append(_totalPaymentRule, _totalPaymentItem)
	}

	logs, sub, err := _Hwescrow.contract.WatchLogs(opts, "OfferCreated", _recruiterRule, _creatorRule, _totalPaymentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwescrowOfferCreated)
				if err := _Hwescrow.contract.UnpackLog(event, "OfferCreated", log); err != nil {
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

// ParseOfferCreated is a log parse operation binding the contract event 0x4f05be72ad7f57c27e555ace8452f56d8b1e82c9e6e1cd4fd282f34518b7729a.
//
// Solidity: event OfferCreated(address indexed _recruiter, address indexed _creator, uint256 indexed _totalPayment, address _paymentToken)
func (_Hwescrow *HwescrowFilterer) ParseOfferCreated(log types.Log) (*HwescrowOfferCreated, error) {
	event := new(HwescrowOfferCreated)
	if err := _Hwescrow.contract.UnpackLog(event, "OfferCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwescrowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Hwescrow contract.
type HwescrowOwnershipTransferredIterator struct {
	Event *HwescrowOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HwescrowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwescrowOwnershipTransferred)
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
		it.Event = new(HwescrowOwnershipTransferred)
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
func (it *HwescrowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwescrowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwescrowOwnershipTransferred represents a OwnershipTransferred event raised by the Hwescrow contract.
type HwescrowOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hwescrow *HwescrowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HwescrowOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hwescrow.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HwescrowOwnershipTransferredIterator{contract: _Hwescrow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hwescrow *HwescrowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HwescrowOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hwescrow.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwescrowOwnershipTransferred)
				if err := _Hwescrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hwescrow *HwescrowFilterer) ParseOwnershipTransferred(log types.Log) (*HwescrowOwnershipTransferred, error) {
	event := new(HwescrowOwnershipTransferred)
	if err := _Hwescrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwescrowPaymentClaimedIterator is returned from FilterPaymentClaimed and is used to iterate over the raw logs and unpacked data for PaymentClaimed events raised by the Hwescrow contract.
type HwescrowPaymentClaimedIterator struct {
	Event *HwescrowPaymentClaimed // Event containing the contract specifics and raw log

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
func (it *HwescrowPaymentClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwescrowPaymentClaimed)
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
		it.Event = new(HwescrowPaymentClaimed)
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
func (it *HwescrowPaymentClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwescrowPaymentClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwescrowPaymentClaimed represents a PaymentClaimed event raised by the Hwescrow contract.
type HwescrowPaymentClaimed struct {
	DealId          *big.Int
	Creator         common.Address
	PaymentReceived *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaymentClaimed is a free log retrieval operation binding the contract event 0xdd74eb9cba06f514be98342b0ed0290146aee9feaf27e8a7ddcb151bb43dac86.
//
// Solidity: event PaymentClaimed(uint256 indexed _dealId, address indexed _creator, uint256 indexed _paymentReceived)
func (_Hwescrow *HwescrowFilterer) FilterPaymentClaimed(opts *bind.FilterOpts, _dealId []*big.Int, _creator []common.Address, _paymentReceived []*big.Int) (*HwescrowPaymentClaimedIterator, error) {

	var _dealIdRule []interface{}
	for _, _dealIdItem := range _dealId {
		_dealIdRule = append(_dealIdRule, _dealIdItem)
	}
	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}
	var _paymentReceivedRule []interface{}
	for _, _paymentReceivedItem := range _paymentReceived {
		_paymentReceivedRule = append(_paymentReceivedRule, _paymentReceivedItem)
	}

	logs, sub, err := _Hwescrow.contract.FilterLogs(opts, "PaymentClaimed", _dealIdRule, _creatorRule, _paymentReceivedRule)
	if err != nil {
		return nil, err
	}
	return &HwescrowPaymentClaimedIterator{contract: _Hwescrow.contract, event: "PaymentClaimed", logs: logs, sub: sub}, nil
}

// WatchPaymentClaimed is a free log subscription operation binding the contract event 0xdd74eb9cba06f514be98342b0ed0290146aee9feaf27e8a7ddcb151bb43dac86.
//
// Solidity: event PaymentClaimed(uint256 indexed _dealId, address indexed _creator, uint256 indexed _paymentReceived)
func (_Hwescrow *HwescrowFilterer) WatchPaymentClaimed(opts *bind.WatchOpts, sink chan<- *HwescrowPaymentClaimed, _dealId []*big.Int, _creator []common.Address, _paymentReceived []*big.Int) (event.Subscription, error) {

	var _dealIdRule []interface{}
	for _, _dealIdItem := range _dealId {
		_dealIdRule = append(_dealIdRule, _dealIdItem)
	}
	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}
	var _paymentReceivedRule []interface{}
	for _, _paymentReceivedItem := range _paymentReceived {
		_paymentReceivedRule = append(_paymentReceivedRule, _paymentReceivedItem)
	}

	logs, sub, err := _Hwescrow.contract.WatchLogs(opts, "PaymentClaimed", _dealIdRule, _creatorRule, _paymentReceivedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwescrowPaymentClaimed)
				if err := _Hwescrow.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
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

// ParsePaymentClaimed is a log parse operation binding the contract event 0xdd74eb9cba06f514be98342b0ed0290146aee9feaf27e8a7ddcb151bb43dac86.
//
// Solidity: event PaymentClaimed(uint256 indexed _dealId, address indexed _creator, uint256 indexed _paymentReceived)
func (_Hwescrow *HwescrowFilterer) ParsePaymentClaimed(log types.Log) (*HwescrowPaymentClaimed, error) {
	event := new(HwescrowPaymentClaimed)
	if err := _Hwescrow.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwescrowPaymentUnlockedIterator is returned from FilterPaymentUnlocked and is used to iterate over the raw logs and unpacked data for PaymentUnlocked events raised by the Hwescrow contract.
type HwescrowPaymentUnlockedIterator struct {
	Event *HwescrowPaymentUnlocked // Event containing the contract specifics and raw log

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
func (it *HwescrowPaymentUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwescrowPaymentUnlocked)
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
		it.Event = new(HwescrowPaymentUnlocked)
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
func (it *HwescrowPaymentUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwescrowPaymentUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwescrowPaymentUnlocked represents a PaymentUnlocked event raised by the Hwescrow contract.
type HwescrowPaymentUnlocked struct {
	DealId         *big.Int
	Recruiter      common.Address
	UnlockedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPaymentUnlocked is a free log retrieval operation binding the contract event 0x4b20f285c1b78452a0587f8653b5b7d15737605123dafe28d7b687d04176ebbc.
//
// Solidity: event PaymentUnlocked(uint256 _dealId, address indexed _recruiter, uint256 indexed _unlockedAmount)
func (_Hwescrow *HwescrowFilterer) FilterPaymentUnlocked(opts *bind.FilterOpts, _recruiter []common.Address, _unlockedAmount []*big.Int) (*HwescrowPaymentUnlockedIterator, error) {

	var _recruiterRule []interface{}
	for _, _recruiterItem := range _recruiter {
		_recruiterRule = append(_recruiterRule, _recruiterItem)
	}
	var _unlockedAmountRule []interface{}
	for _, _unlockedAmountItem := range _unlockedAmount {
		_unlockedAmountRule = append(_unlockedAmountRule, _unlockedAmountItem)
	}

	logs, sub, err := _Hwescrow.contract.FilterLogs(opts, "PaymentUnlocked", _recruiterRule, _unlockedAmountRule)
	if err != nil {
		return nil, err
	}
	return &HwescrowPaymentUnlockedIterator{contract: _Hwescrow.contract, event: "PaymentUnlocked", logs: logs, sub: sub}, nil
}

// WatchPaymentUnlocked is a free log subscription operation binding the contract event 0x4b20f285c1b78452a0587f8653b5b7d15737605123dafe28d7b687d04176ebbc.
//
// Solidity: event PaymentUnlocked(uint256 _dealId, address indexed _recruiter, uint256 indexed _unlockedAmount)
func (_Hwescrow *HwescrowFilterer) WatchPaymentUnlocked(opts *bind.WatchOpts, sink chan<- *HwescrowPaymentUnlocked, _recruiter []common.Address, _unlockedAmount []*big.Int) (event.Subscription, error) {

	var _recruiterRule []interface{}
	for _, _recruiterItem := range _recruiter {
		_recruiterRule = append(_recruiterRule, _recruiterItem)
	}
	var _unlockedAmountRule []interface{}
	for _, _unlockedAmountItem := range _unlockedAmount {
		_unlockedAmountRule = append(_unlockedAmountRule, _unlockedAmountItem)
	}

	logs, sub, err := _Hwescrow.contract.WatchLogs(opts, "PaymentUnlocked", _recruiterRule, _unlockedAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwescrowPaymentUnlocked)
				if err := _Hwescrow.contract.UnpackLog(event, "PaymentUnlocked", log); err != nil {
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

// ParsePaymentUnlocked is a log parse operation binding the contract event 0x4b20f285c1b78452a0587f8653b5b7d15737605123dafe28d7b687d04176ebbc.
//
// Solidity: event PaymentUnlocked(uint256 _dealId, address indexed _recruiter, uint256 indexed _unlockedAmount)
func (_Hwescrow *HwescrowFilterer) ParsePaymentUnlocked(log types.Log) (*HwescrowPaymentUnlocked, error) {
	event := new(HwescrowPaymentUnlocked)
	if err := _Hwescrow.contract.UnpackLog(event, "PaymentUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwescrowPaymentWithdrawnIterator is returned from FilterPaymentWithdrawn and is used to iterate over the raw logs and unpacked data for PaymentWithdrawn events raised by the Hwescrow contract.
type HwescrowPaymentWithdrawnIterator struct {
	Event *HwescrowPaymentWithdrawn // Event containing the contract specifics and raw log

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
func (it *HwescrowPaymentWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwescrowPaymentWithdrawn)
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
		it.Event = new(HwescrowPaymentWithdrawn)
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
func (it *HwescrowPaymentWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwescrowPaymentWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwescrowPaymentWithdrawn represents a PaymentWithdrawn event raised by the Hwescrow contract.
type HwescrowPaymentWithdrawn struct {
	DealId *big.Int
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentWithdrawn is a free log retrieval operation binding the contract event 0x8d59dd7e829f52869194dc5eebe828bdda1ee85be0b51f6946653a90e936aca9.
//
// Solidity: event PaymentWithdrawn(uint256 indexed _dealId, uint8 status)
func (_Hwescrow *HwescrowFilterer) FilterPaymentWithdrawn(opts *bind.FilterOpts, _dealId []*big.Int) (*HwescrowPaymentWithdrawnIterator, error) {

	var _dealIdRule []interface{}
	for _, _dealIdItem := range _dealId {
		_dealIdRule = append(_dealIdRule, _dealIdItem)
	}

	logs, sub, err := _Hwescrow.contract.FilterLogs(opts, "PaymentWithdrawn", _dealIdRule)
	if err != nil {
		return nil, err
	}
	return &HwescrowPaymentWithdrawnIterator{contract: _Hwescrow.contract, event: "PaymentWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPaymentWithdrawn is a free log subscription operation binding the contract event 0x8d59dd7e829f52869194dc5eebe828bdda1ee85be0b51f6946653a90e936aca9.
//
// Solidity: event PaymentWithdrawn(uint256 indexed _dealId, uint8 status)
func (_Hwescrow *HwescrowFilterer) WatchPaymentWithdrawn(opts *bind.WatchOpts, sink chan<- *HwescrowPaymentWithdrawn, _dealId []*big.Int) (event.Subscription, error) {

	var _dealIdRule []interface{}
	for _, _dealIdItem := range _dealId {
		_dealIdRule = append(_dealIdRule, _dealIdItem)
	}

	logs, sub, err := _Hwescrow.contract.WatchLogs(opts, "PaymentWithdrawn", _dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwescrowPaymentWithdrawn)
				if err := _Hwescrow.contract.UnpackLog(event, "PaymentWithdrawn", log); err != nil {
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

// ParsePaymentWithdrawn is a log parse operation binding the contract event 0x8d59dd7e829f52869194dc5eebe828bdda1ee85be0b51f6946653a90e936aca9.
//
// Solidity: event PaymentWithdrawn(uint256 indexed _dealId, uint8 status)
func (_Hwescrow *HwescrowFilterer) ParsePaymentWithdrawn(log types.Log) (*HwescrowPaymentWithdrawn, error) {
	event := new(HwescrowPaymentWithdrawn)
	if err := _Hwescrow.contract.UnpackLog(event, "PaymentWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HwescrowTotalFeeClaimedIterator is returned from FilterTotalFeeClaimed and is used to iterate over the raw logs and unpacked data for TotalFeeClaimed events raised by the Hwescrow contract.
type HwescrowTotalFeeClaimedIterator struct {
	Event *HwescrowTotalFeeClaimed // Event containing the contract specifics and raw log

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
func (it *HwescrowTotalFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HwescrowTotalFeeClaimed)
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
		it.Event = new(HwescrowTotalFeeClaimed)
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
func (it *HwescrowTotalFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HwescrowTotalFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HwescrowTotalFeeClaimed represents a TotalFeeClaimed event raised by the Hwescrow contract.
type HwescrowTotalFeeClaimed struct {
	Collector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTotalFeeClaimed is a free log retrieval operation binding the contract event 0xfea685e9aaf14c8da90ce1e747e388a191e2062a563acfee68847c3fbd5c512c.
//
// Solidity: event TotalFeeClaimed(address _collector)
func (_Hwescrow *HwescrowFilterer) FilterTotalFeeClaimed(opts *bind.FilterOpts) (*HwescrowTotalFeeClaimedIterator, error) {

	logs, sub, err := _Hwescrow.contract.FilterLogs(opts, "TotalFeeClaimed")
	if err != nil {
		return nil, err
	}
	return &HwescrowTotalFeeClaimedIterator{contract: _Hwescrow.contract, event: "TotalFeeClaimed", logs: logs, sub: sub}, nil
}

// WatchTotalFeeClaimed is a free log subscription operation binding the contract event 0xfea685e9aaf14c8da90ce1e747e388a191e2062a563acfee68847c3fbd5c512c.
//
// Solidity: event TotalFeeClaimed(address _collector)
func (_Hwescrow *HwescrowFilterer) WatchTotalFeeClaimed(opts *bind.WatchOpts, sink chan<- *HwescrowTotalFeeClaimed) (event.Subscription, error) {

	logs, sub, err := _Hwescrow.contract.WatchLogs(opts, "TotalFeeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HwescrowTotalFeeClaimed)
				if err := _Hwescrow.contract.UnpackLog(event, "TotalFeeClaimed", log); err != nil {
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

// ParseTotalFeeClaimed is a log parse operation binding the contract event 0xfea685e9aaf14c8da90ce1e747e388a191e2062a563acfee68847c3fbd5c512c.
//
// Solidity: event TotalFeeClaimed(address _collector)
func (_Hwescrow *HwescrowFilterer) ParseTotalFeeClaimed(log types.Log) (*HwescrowTotalFeeClaimed, error) {
	event := new(HwescrowTotalFeeClaimed)
	if err := _Hwescrow.contract.UnpackLog(event, "TotalFeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
