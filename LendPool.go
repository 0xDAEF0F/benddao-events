// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
	_ = abi.ConvertType
)

// DataTypesNftConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNftConfigurationMap struct {
	Data *big.Int
}

// DataTypesNftData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNftData struct {
	Configuration DataTypesNftConfigurationMap
	BNftAddress   common.Address
	Id            uint8
	MaxSupply     *big.Int
	MaxTokenId    *big.Int
}

// DataTypesReserveConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveConfigurationMap struct {
	Data *big.Int
}

// DataTypesReserveData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveData struct {
	Configuration             DataTypesReserveConfigurationMap
	LiquidityIndex            *big.Int
	VariableBorrowIndex       *big.Int
	CurrentLiquidityRate      *big.Int
	CurrentVariableBorrowRate *big.Int
	LastUpdateTimestamp       *big.Int
	BTokenAddress             common.Address
	DebtTokenAddress          common.Address
	InterestRateAddress       common.Address
	Id                        uint8
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Auction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"durationTime\",\"type\":\"uint256\"}],\"name\":\"PausedTimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fineAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"ReserveDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"auction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"nftAssets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nftTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"batchBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nftAssets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nftTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchRepay\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceFromBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceToBefore\",\"type\":\"uint256\"}],\"name\":\"finalizeTransfer\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressesProvider\",\"outputs\":[{\"internalType\":\"contractILendPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxNumberOfNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxNumberOfReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftAuctionData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFine\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftAuctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemEndTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reserveAsset\",\"type\":\"address\"}],\"name\":\"getNftCollateralData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralInETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateralInReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsInETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsInReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getNftConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NftConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getNftData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NftConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"bNftAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NftData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftDebtData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reserveAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftLiquidatePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidatePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paybackAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNftsList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"bTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"internalType\":\"structDataTypes.ReserveData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedVariableDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bNftAddress\",\"type\":\"address\"}],\"name\":\"initNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateAddress\",\"type\":\"address\"}],\"name\":\"initReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILendPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFine\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setMaxNumberOfNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setMaxNumberOfReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"}],\"name\":\"setNftConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"}],\"name\":\"setNftMaxSupplyAndTokenId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationTime\",\"type\":\"uint256\"}],\"name\":\"setPausedTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"}],\"name\":\"setReserveConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rateAddress\",\"type\":\"address\"}],\"name\":\"setReserveInterestRateAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// FinalizeTransfer is a free data retrieval call binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) view returns()
func (_Main *MainCaller) FinalizeTransfer(opts *bind.CallOpts, asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) error {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "finalizeTransfer", asset, from, to, amount, balanceFromBefore, balanceToBefore)

	if err != nil {
		return err
	}

	return err

}

// FinalizeTransfer is a free data retrieval call binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) view returns()
func (_Main *MainSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) error {
	return _Main.Contract.FinalizeTransfer(&_Main.CallOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FinalizeTransfer is a free data retrieval call binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) view returns()
func (_Main *MainCallerSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) error {
	return _Main.Contract.FinalizeTransfer(&_Main.CallOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_Main *MainCaller) GetAddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAddressesProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_Main *MainSession) GetAddressesProvider() (common.Address, error) {
	return _Main.Contract.GetAddressesProvider(&_Main.CallOpts)
}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_Main *MainCallerSession) GetAddressesProvider() (common.Address, error) {
	return _Main.Contract.GetAddressesProvider(&_Main.CallOpts)
}

// GetMaxNumberOfNfts is a free data retrieval call binding the contract method 0xdd90ff38.
//
// Solidity: function getMaxNumberOfNfts() view returns(uint256)
func (_Main *MainCaller) GetMaxNumberOfNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getMaxNumberOfNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumberOfNfts is a free data retrieval call binding the contract method 0xdd90ff38.
//
// Solidity: function getMaxNumberOfNfts() view returns(uint256)
func (_Main *MainSession) GetMaxNumberOfNfts() (*big.Int, error) {
	return _Main.Contract.GetMaxNumberOfNfts(&_Main.CallOpts)
}

// GetMaxNumberOfNfts is a free data retrieval call binding the contract method 0xdd90ff38.
//
// Solidity: function getMaxNumberOfNfts() view returns(uint256)
func (_Main *MainCallerSession) GetMaxNumberOfNfts() (*big.Int, error) {
	return _Main.Contract.GetMaxNumberOfNfts(&_Main.CallOpts)
}

// GetMaxNumberOfReserves is a free data retrieval call binding the contract method 0x08ac08b9.
//
// Solidity: function getMaxNumberOfReserves() view returns(uint256)
func (_Main *MainCaller) GetMaxNumberOfReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getMaxNumberOfReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumberOfReserves is a free data retrieval call binding the contract method 0x08ac08b9.
//
// Solidity: function getMaxNumberOfReserves() view returns(uint256)
func (_Main *MainSession) GetMaxNumberOfReserves() (*big.Int, error) {
	return _Main.Contract.GetMaxNumberOfReserves(&_Main.CallOpts)
}

// GetMaxNumberOfReserves is a free data retrieval call binding the contract method 0x08ac08b9.
//
// Solidity: function getMaxNumberOfReserves() view returns(uint256)
func (_Main *MainCallerSession) GetMaxNumberOfReserves() (*big.Int, error) {
	return _Main.Contract.GetMaxNumberOfReserves(&_Main.CallOpts)
}

// GetNftAuctionData is a free data retrieval call binding the contract method 0xe5bceca5.
//
// Solidity: function getNftAuctionData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address bidderAddress, uint256 bidPrice, uint256 bidBorrowAmount, uint256 bidFine)
func (_Main *MainCaller) GetNftAuctionData(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId          *big.Int
	BidderAddress   common.Address
	BidPrice        *big.Int
	BidBorrowAmount *big.Int
	BidFine         *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getNftAuctionData", nftAsset, nftTokenId)

	outstruct := new(struct {
		LoanId          *big.Int
		BidderAddress   common.Address
		BidPrice        *big.Int
		BidBorrowAmount *big.Int
		BidFine         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LoanId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BidderAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BidPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BidBorrowAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BidFine = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftAuctionData is a free data retrieval call binding the contract method 0xe5bceca5.
//
// Solidity: function getNftAuctionData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address bidderAddress, uint256 bidPrice, uint256 bidBorrowAmount, uint256 bidFine)
func (_Main *MainSession) GetNftAuctionData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId          *big.Int
	BidderAddress   common.Address
	BidPrice        *big.Int
	BidBorrowAmount *big.Int
	BidFine         *big.Int
}, error) {
	return _Main.Contract.GetNftAuctionData(&_Main.CallOpts, nftAsset, nftTokenId)
}

// GetNftAuctionData is a free data retrieval call binding the contract method 0xe5bceca5.
//
// Solidity: function getNftAuctionData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address bidderAddress, uint256 bidPrice, uint256 bidBorrowAmount, uint256 bidFine)
func (_Main *MainCallerSession) GetNftAuctionData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId          *big.Int
	BidderAddress   common.Address
	BidPrice        *big.Int
	BidBorrowAmount *big.Int
	BidFine         *big.Int
}, error) {
	return _Main.Contract.GetNftAuctionData(&_Main.CallOpts, nftAsset, nftTokenId)
}

// GetNftAuctionEndTime is a free data retrieval call binding the contract method 0x17c8a456.
//
// Solidity: function getNftAuctionEndTime(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, uint256 bidStartTimestamp, uint256 bidEndTimestamp, uint256 redeemEndTimestamp)
func (_Main *MainCaller) GetNftAuctionEndTime(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId             *big.Int
	BidStartTimestamp  *big.Int
	BidEndTimestamp    *big.Int
	RedeemEndTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getNftAuctionEndTime", nftAsset, nftTokenId)

	outstruct := new(struct {
		LoanId             *big.Int
		BidStartTimestamp  *big.Int
		BidEndTimestamp    *big.Int
		RedeemEndTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LoanId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BidStartTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BidEndTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RedeemEndTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftAuctionEndTime is a free data retrieval call binding the contract method 0x17c8a456.
//
// Solidity: function getNftAuctionEndTime(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, uint256 bidStartTimestamp, uint256 bidEndTimestamp, uint256 redeemEndTimestamp)
func (_Main *MainSession) GetNftAuctionEndTime(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId             *big.Int
	BidStartTimestamp  *big.Int
	BidEndTimestamp    *big.Int
	RedeemEndTimestamp *big.Int
}, error) {
	return _Main.Contract.GetNftAuctionEndTime(&_Main.CallOpts, nftAsset, nftTokenId)
}

// GetNftAuctionEndTime is a free data retrieval call binding the contract method 0x17c8a456.
//
// Solidity: function getNftAuctionEndTime(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, uint256 bidStartTimestamp, uint256 bidEndTimestamp, uint256 redeemEndTimestamp)
func (_Main *MainCallerSession) GetNftAuctionEndTime(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId             *big.Int
	BidStartTimestamp  *big.Int
	BidEndTimestamp    *big.Int
	RedeemEndTimestamp *big.Int
}, error) {
	return _Main.Contract.GetNftAuctionEndTime(&_Main.CallOpts, nftAsset, nftTokenId)
}

// GetNftCollateralData is a free data retrieval call binding the contract method 0xcc8ccdf2.
//
// Solidity: function getNftCollateralData(address nftAsset, address reserveAsset) view returns(uint256 totalCollateralInETH, uint256 totalCollateralInReserve, uint256 availableBorrowsInETH, uint256 availableBorrowsInReserve, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus)
func (_Main *MainCaller) GetNftCollateralData(opts *bind.CallOpts, nftAsset common.Address, reserveAsset common.Address) (struct {
	TotalCollateralInETH      *big.Int
	TotalCollateralInReserve  *big.Int
	AvailableBorrowsInETH     *big.Int
	AvailableBorrowsInReserve *big.Int
	Ltv                       *big.Int
	LiquidationThreshold      *big.Int
	LiquidationBonus          *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getNftCollateralData", nftAsset, reserveAsset)

	outstruct := new(struct {
		TotalCollateralInETH      *big.Int
		TotalCollateralInReserve  *big.Int
		AvailableBorrowsInETH     *big.Int
		AvailableBorrowsInReserve *big.Int
		Ltv                       *big.Int
		LiquidationThreshold      *big.Int
		LiquidationBonus          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCollateralInETH = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalCollateralInReserve = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsInETH = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsInReserve = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LiquidationBonus = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftCollateralData is a free data retrieval call binding the contract method 0xcc8ccdf2.
//
// Solidity: function getNftCollateralData(address nftAsset, address reserveAsset) view returns(uint256 totalCollateralInETH, uint256 totalCollateralInReserve, uint256 availableBorrowsInETH, uint256 availableBorrowsInReserve, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus)
func (_Main *MainSession) GetNftCollateralData(nftAsset common.Address, reserveAsset common.Address) (struct {
	TotalCollateralInETH      *big.Int
	TotalCollateralInReserve  *big.Int
	AvailableBorrowsInETH     *big.Int
	AvailableBorrowsInReserve *big.Int
	Ltv                       *big.Int
	LiquidationThreshold      *big.Int
	LiquidationBonus          *big.Int
}, error) {
	return _Main.Contract.GetNftCollateralData(&_Main.CallOpts, nftAsset, reserveAsset)
}

// GetNftCollateralData is a free data retrieval call binding the contract method 0xcc8ccdf2.
//
// Solidity: function getNftCollateralData(address nftAsset, address reserveAsset) view returns(uint256 totalCollateralInETH, uint256 totalCollateralInReserve, uint256 availableBorrowsInETH, uint256 availableBorrowsInReserve, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus)
func (_Main *MainCallerSession) GetNftCollateralData(nftAsset common.Address, reserveAsset common.Address) (struct {
	TotalCollateralInETH      *big.Int
	TotalCollateralInReserve  *big.Int
	AvailableBorrowsInETH     *big.Int
	AvailableBorrowsInReserve *big.Int
	Ltv                       *big.Int
	LiquidationThreshold      *big.Int
	LiquidationBonus          *big.Int
}, error) {
	return _Main.Contract.GetNftCollateralData(&_Main.CallOpts, nftAsset, reserveAsset)
}

// GetNftConfiguration is a free data retrieval call binding the contract method 0x87c32dec.
//
// Solidity: function getNftConfiguration(address asset) view returns((uint256))
func (_Main *MainCaller) GetNftConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesNftConfigurationMap, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getNftConfiguration", asset)

	if err != nil {
		return *new(DataTypesNftConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNftConfigurationMap)).(*DataTypesNftConfigurationMap)

	return out0, err

}

// GetNftConfiguration is a free data retrieval call binding the contract method 0x87c32dec.
//
// Solidity: function getNftConfiguration(address asset) view returns((uint256))
func (_Main *MainSession) GetNftConfiguration(asset common.Address) (DataTypesNftConfigurationMap, error) {
	return _Main.Contract.GetNftConfiguration(&_Main.CallOpts, asset)
}

// GetNftConfiguration is a free data retrieval call binding the contract method 0x87c32dec.
//
// Solidity: function getNftConfiguration(address asset) view returns((uint256))
func (_Main *MainCallerSession) GetNftConfiguration(asset common.Address) (DataTypesNftConfigurationMap, error) {
	return _Main.Contract.GetNftConfiguration(&_Main.CallOpts, asset)
}

// GetNftData is a free data retrieval call binding the contract method 0x77bdc0c3.
//
// Solidity: function getNftData(address asset) view returns(((uint256),address,uint8,uint256,uint256))
func (_Main *MainCaller) GetNftData(opts *bind.CallOpts, asset common.Address) (DataTypesNftData, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getNftData", asset)

	if err != nil {
		return *new(DataTypesNftData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNftData)).(*DataTypesNftData)

	return out0, err

}

// GetNftData is a free data retrieval call binding the contract method 0x77bdc0c3.
//
// Solidity: function getNftData(address asset) view returns(((uint256),address,uint8,uint256,uint256))
func (_Main *MainSession) GetNftData(asset common.Address) (DataTypesNftData, error) {
	return _Main.Contract.GetNftData(&_Main.CallOpts, asset)
}

// GetNftData is a free data retrieval call binding the contract method 0x77bdc0c3.
//
// Solidity: function getNftData(address asset) view returns(((uint256),address,uint8,uint256,uint256))
func (_Main *MainCallerSession) GetNftData(asset common.Address) (DataTypesNftData, error) {
	return _Main.Contract.GetNftData(&_Main.CallOpts, asset)
}

// GetNftDebtData is a free data retrieval call binding the contract method 0xec765d3d.
//
// Solidity: function getNftDebtData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address reserveAsset, uint256 totalCollateral, uint256 totalDebt, uint256 availableBorrows, uint256 healthFactor)
func (_Main *MainCaller) GetNftDebtData(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId           *big.Int
	ReserveAsset     common.Address
	TotalCollateral  *big.Int
	TotalDebt        *big.Int
	AvailableBorrows *big.Int
	HealthFactor     *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getNftDebtData", nftAsset, nftTokenId)

	outstruct := new(struct {
		LoanId           *big.Int
		ReserveAsset     common.Address
		TotalCollateral  *big.Int
		TotalDebt        *big.Int
		AvailableBorrows *big.Int
		HealthFactor     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LoanId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveAsset = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TotalCollateral = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrows = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HealthFactor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftDebtData is a free data retrieval call binding the contract method 0xec765d3d.
//
// Solidity: function getNftDebtData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address reserveAsset, uint256 totalCollateral, uint256 totalDebt, uint256 availableBorrows, uint256 healthFactor)
func (_Main *MainSession) GetNftDebtData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId           *big.Int
	ReserveAsset     common.Address
	TotalCollateral  *big.Int
	TotalDebt        *big.Int
	AvailableBorrows *big.Int
	HealthFactor     *big.Int
}, error) {
	return _Main.Contract.GetNftDebtData(&_Main.CallOpts, nftAsset, nftTokenId)
}

// GetNftDebtData is a free data retrieval call binding the contract method 0xec765d3d.
//
// Solidity: function getNftDebtData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address reserveAsset, uint256 totalCollateral, uint256 totalDebt, uint256 availableBorrows, uint256 healthFactor)
func (_Main *MainCallerSession) GetNftDebtData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId           *big.Int
	ReserveAsset     common.Address
	TotalCollateral  *big.Int
	TotalDebt        *big.Int
	AvailableBorrows *big.Int
	HealthFactor     *big.Int
}, error) {
	return _Main.Contract.GetNftDebtData(&_Main.CallOpts, nftAsset, nftTokenId)
}

// GetNftLiquidatePrice is a free data retrieval call binding the contract method 0x798b9e3d.
//
// Solidity: function getNftLiquidatePrice(address nftAsset, uint256 nftTokenId) view returns(uint256 liquidatePrice, uint256 paybackAmount)
func (_Main *MainCaller) GetNftLiquidatePrice(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LiquidatePrice *big.Int
	PaybackAmount  *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getNftLiquidatePrice", nftAsset, nftTokenId)

	outstruct := new(struct {
		LiquidatePrice *big.Int
		PaybackAmount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidatePrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PaybackAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftLiquidatePrice is a free data retrieval call binding the contract method 0x798b9e3d.
//
// Solidity: function getNftLiquidatePrice(address nftAsset, uint256 nftTokenId) view returns(uint256 liquidatePrice, uint256 paybackAmount)
func (_Main *MainSession) GetNftLiquidatePrice(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LiquidatePrice *big.Int
	PaybackAmount  *big.Int
}, error) {
	return _Main.Contract.GetNftLiquidatePrice(&_Main.CallOpts, nftAsset, nftTokenId)
}

// GetNftLiquidatePrice is a free data retrieval call binding the contract method 0x798b9e3d.
//
// Solidity: function getNftLiquidatePrice(address nftAsset, uint256 nftTokenId) view returns(uint256 liquidatePrice, uint256 paybackAmount)
func (_Main *MainCallerSession) GetNftLiquidatePrice(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LiquidatePrice *big.Int
	PaybackAmount  *big.Int
}, error) {
	return _Main.Contract.GetNftLiquidatePrice(&_Main.CallOpts, nftAsset, nftTokenId)
}

// GetNftsList is a free data retrieval call binding the contract method 0x6b25c835.
//
// Solidity: function getNftsList() view returns(address[])
func (_Main *MainCaller) GetNftsList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getNftsList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetNftsList is a free data retrieval call binding the contract method 0x6b25c835.
//
// Solidity: function getNftsList() view returns(address[])
func (_Main *MainSession) GetNftsList() ([]common.Address, error) {
	return _Main.Contract.GetNftsList(&_Main.CallOpts)
}

// GetNftsList is a free data retrieval call binding the contract method 0x6b25c835.
//
// Solidity: function getNftsList() view returns(address[])
func (_Main *MainCallerSession) GetNftsList() ([]common.Address, error) {
	return _Main.Contract.GetNftsList(&_Main.CallOpts)
}

// GetPausedTime is a free data retrieval call binding the contract method 0x8fc42188.
//
// Solidity: function getPausedTime() view returns(uint256, uint256)
func (_Main *MainCaller) GetPausedTime(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getPausedTime")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPausedTime is a free data retrieval call binding the contract method 0x8fc42188.
//
// Solidity: function getPausedTime() view returns(uint256, uint256)
func (_Main *MainSession) GetPausedTime() (*big.Int, *big.Int, error) {
	return _Main.Contract.GetPausedTime(&_Main.CallOpts)
}

// GetPausedTime is a free data retrieval call binding the contract method 0x8fc42188.
//
// Solidity: function getPausedTime() view returns(uint256, uint256)
func (_Main *MainCallerSession) GetPausedTime() (*big.Int, *big.Int, error) {
	return _Main.Contract.GetPausedTime(&_Main.CallOpts)
}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address asset) view returns((uint256))
func (_Main *MainCaller) GetReserveConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesReserveConfigurationMap, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getReserveConfiguration", asset)

	if err != nil {
		return *new(DataTypesReserveConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveConfigurationMap)).(*DataTypesReserveConfigurationMap)

	return out0, err

}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address asset) view returns((uint256))
func (_Main *MainSession) GetReserveConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _Main.Contract.GetReserveConfiguration(&_Main.CallOpts, asset)
}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address asset) view returns((uint256))
func (_Main *MainCallerSession) GetReserveConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _Main.Contract.GetReserveConfiguration(&_Main.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint40,address,address,address,uint8))
func (_Main *MainCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (DataTypesReserveData, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getReserveData", asset)

	if err != nil {
		return *new(DataTypesReserveData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveData)).(*DataTypesReserveData)

	return out0, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint40,address,address,address,uint8))
func (_Main *MainSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _Main.Contract.GetReserveData(&_Main.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint40,address,address,address,uint8))
func (_Main *MainCallerSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _Main.Contract.GetReserveData(&_Main.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_Main *MainCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getReserveNormalizedIncome", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_Main *MainSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _Main.Contract.GetReserveNormalizedIncome(&_Main.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_Main *MainCallerSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _Main.Contract.GetReserveNormalizedIncome(&_Main.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_Main *MainCaller) GetReserveNormalizedVariableDebt(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getReserveNormalizedVariableDebt", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_Main *MainSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _Main.Contract.GetReserveNormalizedVariableDebt(&_Main.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_Main *MainCallerSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _Main.Contract.GetReserveNormalizedVariableDebt(&_Main.CallOpts, asset)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_Main *MainCaller) GetReservesList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getReservesList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_Main *MainSession) GetReservesList() ([]common.Address, error) {
	return _Main.Contract.GetReservesList(&_Main.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_Main *MainCallerSession) GetReservesList() ([]common.Address, error) {
	return _Main.Contract.GetReservesList(&_Main.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_Main *MainCaller) OnERC721Received(opts *bind.CallOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "onERC721Received", operator, from, tokenId, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_Main *MainSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _Main.Contract.OnERC721Received(&_Main.CallOpts, operator, from, tokenId, data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_Main *MainCallerSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _Main.Contract.OnERC721Received(&_Main.CallOpts, operator, from, tokenId, data)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Main *MainCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Main *MainSession) Paused() (bool, error) {
	return _Main.Contract.Paused(&_Main.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Main *MainCallerSession) Paused() (bool, error) {
	return _Main.Contract.Paused(&_Main.CallOpts)
}

// Auction is a paid mutator transaction binding the contract method 0xa4c0166b.
//
// Solidity: function auction(address nftAsset, uint256 nftTokenId, uint256 bidPrice, address onBehalfOf) returns()
func (_Main *MainTransactor) Auction(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, bidPrice *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "auction", nftAsset, nftTokenId, bidPrice, onBehalfOf)
}

// Auction is a paid mutator transaction binding the contract method 0xa4c0166b.
//
// Solidity: function auction(address nftAsset, uint256 nftTokenId, uint256 bidPrice, address onBehalfOf) returns()
func (_Main *MainSession) Auction(nftAsset common.Address, nftTokenId *big.Int, bidPrice *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Main.Contract.Auction(&_Main.TransactOpts, nftAsset, nftTokenId, bidPrice, onBehalfOf)
}

// Auction is a paid mutator transaction binding the contract method 0xa4c0166b.
//
// Solidity: function auction(address nftAsset, uint256 nftTokenId, uint256 bidPrice, address onBehalfOf) returns()
func (_Main *MainTransactorSession) Auction(nftAsset common.Address, nftTokenId *big.Int, bidPrice *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Main.Contract.Auction(&_Main.TransactOpts, nftAsset, nftTokenId, bidPrice, onBehalfOf)
}

// BatchBorrow is a paid mutator transaction binding the contract method 0xc9fef2fe.
//
// Solidity: function batchBorrow(address[] assets, uint256[] amounts, address[] nftAssets, uint256[] nftTokenIds, address onBehalfOf, uint16 referralCode) returns()
func (_Main *MainTransactor) BatchBorrow(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int, nftAssets []common.Address, nftTokenIds []*big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "batchBorrow", assets, amounts, nftAssets, nftTokenIds, onBehalfOf, referralCode)
}

// BatchBorrow is a paid mutator transaction binding the contract method 0xc9fef2fe.
//
// Solidity: function batchBorrow(address[] assets, uint256[] amounts, address[] nftAssets, uint256[] nftTokenIds, address onBehalfOf, uint16 referralCode) returns()
func (_Main *MainSession) BatchBorrow(assets []common.Address, amounts []*big.Int, nftAssets []common.Address, nftTokenIds []*big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Main.Contract.BatchBorrow(&_Main.TransactOpts, assets, amounts, nftAssets, nftTokenIds, onBehalfOf, referralCode)
}

// BatchBorrow is a paid mutator transaction binding the contract method 0xc9fef2fe.
//
// Solidity: function batchBorrow(address[] assets, uint256[] amounts, address[] nftAssets, uint256[] nftTokenIds, address onBehalfOf, uint16 referralCode) returns()
func (_Main *MainTransactorSession) BatchBorrow(assets []common.Address, amounts []*big.Int, nftAssets []common.Address, nftTokenIds []*big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Main.Contract.BatchBorrow(&_Main.TransactOpts, assets, amounts, nftAssets, nftTokenIds, onBehalfOf, referralCode)
}

// BatchRepay is a paid mutator transaction binding the contract method 0x5bc5bbf1.
//
// Solidity: function batchRepay(address[] nftAssets, uint256[] nftTokenIds, uint256[] amounts) returns(uint256[], bool[])
func (_Main *MainTransactor) BatchRepay(opts *bind.TransactOpts, nftAssets []common.Address, nftTokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "batchRepay", nftAssets, nftTokenIds, amounts)
}

// BatchRepay is a paid mutator transaction binding the contract method 0x5bc5bbf1.
//
// Solidity: function batchRepay(address[] nftAssets, uint256[] nftTokenIds, uint256[] amounts) returns(uint256[], bool[])
func (_Main *MainSession) BatchRepay(nftAssets []common.Address, nftTokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.BatchRepay(&_Main.TransactOpts, nftAssets, nftTokenIds, amounts)
}

// BatchRepay is a paid mutator transaction binding the contract method 0x5bc5bbf1.
//
// Solidity: function batchRepay(address[] nftAssets, uint256[] nftTokenIds, uint256[] amounts) returns(uint256[], bool[])
func (_Main *MainTransactorSession) BatchRepay(nftAssets []common.Address, nftTokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.BatchRepay(&_Main.TransactOpts, nftAssets, nftTokenIds, amounts)
}

// Borrow is a paid mutator transaction binding the contract method 0xb6529aee.
//
// Solidity: function borrow(address asset, uint256 amount, address nftAsset, uint256 nftTokenId, address onBehalfOf, uint16 referralCode) returns()
func (_Main *MainTransactor) Borrow(opts *bind.TransactOpts, asset common.Address, amount *big.Int, nftAsset common.Address, nftTokenId *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "borrow", asset, amount, nftAsset, nftTokenId, onBehalfOf, referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xb6529aee.
//
// Solidity: function borrow(address asset, uint256 amount, address nftAsset, uint256 nftTokenId, address onBehalfOf, uint16 referralCode) returns()
func (_Main *MainSession) Borrow(asset common.Address, amount *big.Int, nftAsset common.Address, nftTokenId *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Main.Contract.Borrow(&_Main.TransactOpts, asset, amount, nftAsset, nftTokenId, onBehalfOf, referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xb6529aee.
//
// Solidity: function borrow(address asset, uint256 amount, address nftAsset, uint256 nftTokenId, address onBehalfOf, uint16 referralCode) returns()
func (_Main *MainTransactorSession) Borrow(asset common.Address, amount *big.Int, nftAsset common.Address, nftTokenId *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Main.Contract.Borrow(&_Main.TransactOpts, asset, amount, nftAsset, nftTokenId, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_Main *MainTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "deposit", asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_Main *MainSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Main.Contract.Deposit(&_Main.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_Main *MainTransactorSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Main.Contract.Deposit(&_Main.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// InitNft is a paid mutator transaction binding the contract method 0x873e4dab.
//
// Solidity: function initNft(address asset, address bNftAddress) returns()
func (_Main *MainTransactor) InitNft(opts *bind.TransactOpts, asset common.Address, bNftAddress common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "initNft", asset, bNftAddress)
}

// InitNft is a paid mutator transaction binding the contract method 0x873e4dab.
//
// Solidity: function initNft(address asset, address bNftAddress) returns()
func (_Main *MainSession) InitNft(asset common.Address, bNftAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.InitNft(&_Main.TransactOpts, asset, bNftAddress)
}

// InitNft is a paid mutator transaction binding the contract method 0x873e4dab.
//
// Solidity: function initNft(address asset, address bNftAddress) returns()
func (_Main *MainTransactorSession) InitNft(asset common.Address, bNftAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.InitNft(&_Main.TransactOpts, asset, bNftAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x8bd25677.
//
// Solidity: function initReserve(address asset, address bTokenAddress, address debtTokenAddress, address interestRateAddress) returns()
func (_Main *MainTransactor) InitReserve(opts *bind.TransactOpts, asset common.Address, bTokenAddress common.Address, debtTokenAddress common.Address, interestRateAddress common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "initReserve", asset, bTokenAddress, debtTokenAddress, interestRateAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x8bd25677.
//
// Solidity: function initReserve(address asset, address bTokenAddress, address debtTokenAddress, address interestRateAddress) returns()
func (_Main *MainSession) InitReserve(asset common.Address, bTokenAddress common.Address, debtTokenAddress common.Address, interestRateAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.InitReserve(&_Main.TransactOpts, asset, bTokenAddress, debtTokenAddress, interestRateAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x8bd25677.
//
// Solidity: function initReserve(address asset, address bTokenAddress, address debtTokenAddress, address interestRateAddress) returns()
func (_Main *MainTransactorSession) InitReserve(asset common.Address, bTokenAddress common.Address, debtTokenAddress common.Address, interestRateAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.InitReserve(&_Main.TransactOpts, asset, bTokenAddress, debtTokenAddress, interestRateAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_Main *MainTransactor) Initialize(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "initialize", provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_Main *MainSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _Main.Contract.Initialize(&_Main.TransactOpts, provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_Main *MainTransactorSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _Main.Contract.Initialize(&_Main.TransactOpts, provider)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256)
func (_Main *MainTransactor) Liquidate(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "liquidate", nftAsset, nftTokenId, amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256)
func (_Main *MainSession) Liquidate(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Liquidate(&_Main.TransactOpts, nftAsset, nftTokenId, amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256)
func (_Main *MainTransactorSession) Liquidate(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Liquidate(&_Main.TransactOpts, nftAsset, nftTokenId, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xea2092f3.
//
// Solidity: function redeem(address nftAsset, uint256 nftTokenId, uint256 amount, uint256 bidFine) returns(uint256)
func (_Main *MainTransactor) Redeem(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, amount *big.Int, bidFine *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "redeem", nftAsset, nftTokenId, amount, bidFine)
}

// Redeem is a paid mutator transaction binding the contract method 0xea2092f3.
//
// Solidity: function redeem(address nftAsset, uint256 nftTokenId, uint256 amount, uint256 bidFine) returns(uint256)
func (_Main *MainSession) Redeem(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int, bidFine *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Redeem(&_Main.TransactOpts, nftAsset, nftTokenId, amount, bidFine)
}

// Redeem is a paid mutator transaction binding the contract method 0xea2092f3.
//
// Solidity: function redeem(address nftAsset, uint256 nftTokenId, uint256 amount, uint256 bidFine) returns(uint256)
func (_Main *MainTransactorSession) Redeem(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int, bidFine *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Redeem(&_Main.TransactOpts, nftAsset, nftTokenId, amount, bidFine)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256, bool)
func (_Main *MainTransactor) Repay(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "repay", nftAsset, nftTokenId, amount)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256, bool)
func (_Main *MainSession) Repay(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Repay(&_Main.TransactOpts, nftAsset, nftTokenId, amount)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256, bool)
func (_Main *MainTransactorSession) Repay(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Repay(&_Main.TransactOpts, nftAsset, nftTokenId, amount)
}

// SetMaxNumberOfNfts is a paid mutator transaction binding the contract method 0xfdff6f26.
//
// Solidity: function setMaxNumberOfNfts(uint256 val) returns()
func (_Main *MainTransactor) SetMaxNumberOfNfts(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setMaxNumberOfNfts", val)
}

// SetMaxNumberOfNfts is a paid mutator transaction binding the contract method 0xfdff6f26.
//
// Solidity: function setMaxNumberOfNfts(uint256 val) returns()
func (_Main *MainSession) SetMaxNumberOfNfts(val *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetMaxNumberOfNfts(&_Main.TransactOpts, val)
}

// SetMaxNumberOfNfts is a paid mutator transaction binding the contract method 0xfdff6f26.
//
// Solidity: function setMaxNumberOfNfts(uint256 val) returns()
func (_Main *MainTransactorSession) SetMaxNumberOfNfts(val *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetMaxNumberOfNfts(&_Main.TransactOpts, val)
}

// SetMaxNumberOfReserves is a paid mutator transaction binding the contract method 0x746c35a2.
//
// Solidity: function setMaxNumberOfReserves(uint256 val) returns()
func (_Main *MainTransactor) SetMaxNumberOfReserves(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setMaxNumberOfReserves", val)
}

// SetMaxNumberOfReserves is a paid mutator transaction binding the contract method 0x746c35a2.
//
// Solidity: function setMaxNumberOfReserves(uint256 val) returns()
func (_Main *MainSession) SetMaxNumberOfReserves(val *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetMaxNumberOfReserves(&_Main.TransactOpts, val)
}

// SetMaxNumberOfReserves is a paid mutator transaction binding the contract method 0x746c35a2.
//
// Solidity: function setMaxNumberOfReserves(uint256 val) returns()
func (_Main *MainTransactorSession) SetMaxNumberOfReserves(val *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetMaxNumberOfReserves(&_Main.TransactOpts, val)
}

// SetNftConfiguration is a paid mutator transaction binding the contract method 0x83c8afd7.
//
// Solidity: function setNftConfiguration(address asset, uint256 configuration) returns()
func (_Main *MainTransactor) SetNftConfiguration(opts *bind.TransactOpts, asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setNftConfiguration", asset, configuration)
}

// SetNftConfiguration is a paid mutator transaction binding the contract method 0x83c8afd7.
//
// Solidity: function setNftConfiguration(address asset, uint256 configuration) returns()
func (_Main *MainSession) SetNftConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetNftConfiguration(&_Main.TransactOpts, asset, configuration)
}

// SetNftConfiguration is a paid mutator transaction binding the contract method 0x83c8afd7.
//
// Solidity: function setNftConfiguration(address asset, uint256 configuration) returns()
func (_Main *MainTransactorSession) SetNftConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetNftConfiguration(&_Main.TransactOpts, asset, configuration)
}

// SetNftMaxSupplyAndTokenId is a paid mutator transaction binding the contract method 0xdb78f216.
//
// Solidity: function setNftMaxSupplyAndTokenId(address asset, uint256 maxSupply, uint256 maxTokenId) returns()
func (_Main *MainTransactor) SetNftMaxSupplyAndTokenId(opts *bind.TransactOpts, asset common.Address, maxSupply *big.Int, maxTokenId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setNftMaxSupplyAndTokenId", asset, maxSupply, maxTokenId)
}

// SetNftMaxSupplyAndTokenId is a paid mutator transaction binding the contract method 0xdb78f216.
//
// Solidity: function setNftMaxSupplyAndTokenId(address asset, uint256 maxSupply, uint256 maxTokenId) returns()
func (_Main *MainSession) SetNftMaxSupplyAndTokenId(asset common.Address, maxSupply *big.Int, maxTokenId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetNftMaxSupplyAndTokenId(&_Main.TransactOpts, asset, maxSupply, maxTokenId)
}

// SetNftMaxSupplyAndTokenId is a paid mutator transaction binding the contract method 0xdb78f216.
//
// Solidity: function setNftMaxSupplyAndTokenId(address asset, uint256 maxSupply, uint256 maxTokenId) returns()
func (_Main *MainTransactorSession) SetNftMaxSupplyAndTokenId(asset common.Address, maxSupply *big.Int, maxTokenId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetNftMaxSupplyAndTokenId(&_Main.TransactOpts, asset, maxSupply, maxTokenId)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_Main *MainTransactor) SetPause(opts *bind.TransactOpts, val bool) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setPause", val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_Main *MainSession) SetPause(val bool) (*types.Transaction, error) {
	return _Main.Contract.SetPause(&_Main.TransactOpts, val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_Main *MainTransactorSession) SetPause(val bool) (*types.Transaction, error) {
	return _Main.Contract.SetPause(&_Main.TransactOpts, val)
}

// SetPausedTime is a paid mutator transaction binding the contract method 0x2f923ff7.
//
// Solidity: function setPausedTime(uint256 startTime, uint256 durationTime) returns()
func (_Main *MainTransactor) SetPausedTime(opts *bind.TransactOpts, startTime *big.Int, durationTime *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setPausedTime", startTime, durationTime)
}

// SetPausedTime is a paid mutator transaction binding the contract method 0x2f923ff7.
//
// Solidity: function setPausedTime(uint256 startTime, uint256 durationTime) returns()
func (_Main *MainSession) SetPausedTime(startTime *big.Int, durationTime *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetPausedTime(&_Main.TransactOpts, startTime, durationTime)
}

// SetPausedTime is a paid mutator transaction binding the contract method 0x2f923ff7.
//
// Solidity: function setPausedTime(uint256 startTime, uint256 durationTime) returns()
func (_Main *MainTransactorSession) SetPausedTime(startTime *big.Int, durationTime *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetPausedTime(&_Main.TransactOpts, startTime, durationTime)
}

// SetReserveConfiguration is a paid mutator transaction binding the contract method 0x43f0f733.
//
// Solidity: function setReserveConfiguration(address asset, uint256 configuration) returns()
func (_Main *MainTransactor) SetReserveConfiguration(opts *bind.TransactOpts, asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setReserveConfiguration", asset, configuration)
}

// SetReserveConfiguration is a paid mutator transaction binding the contract method 0x43f0f733.
//
// Solidity: function setReserveConfiguration(address asset, uint256 configuration) returns()
func (_Main *MainSession) SetReserveConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetReserveConfiguration(&_Main.TransactOpts, asset, configuration)
}

// SetReserveConfiguration is a paid mutator transaction binding the contract method 0x43f0f733.
//
// Solidity: function setReserveConfiguration(address asset, uint256 configuration) returns()
func (_Main *MainTransactorSession) SetReserveConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetReserveConfiguration(&_Main.TransactOpts, asset, configuration)
}

// SetReserveInterestRateAddress is a paid mutator transaction binding the contract method 0x83b1555f.
//
// Solidity: function setReserveInterestRateAddress(address asset, address rateAddress) returns()
func (_Main *MainTransactor) SetReserveInterestRateAddress(opts *bind.TransactOpts, asset common.Address, rateAddress common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setReserveInterestRateAddress", asset, rateAddress)
}

// SetReserveInterestRateAddress is a paid mutator transaction binding the contract method 0x83b1555f.
//
// Solidity: function setReserveInterestRateAddress(address asset, address rateAddress) returns()
func (_Main *MainSession) SetReserveInterestRateAddress(asset common.Address, rateAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetReserveInterestRateAddress(&_Main.TransactOpts, asset, rateAddress)
}

// SetReserveInterestRateAddress is a paid mutator transaction binding the contract method 0x83b1555f.
//
// Solidity: function setReserveInterestRateAddress(address asset, address rateAddress) returns()
func (_Main *MainTransactorSession) SetReserveInterestRateAddress(asset common.Address, rateAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetReserveInterestRateAddress(&_Main.TransactOpts, asset, rateAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_Main *MainTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "withdraw", asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_Main *MainSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Main.Contract.Withdraw(&_Main.TransactOpts, asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_Main *MainTransactorSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Main.Contract.Withdraw(&_Main.TransactOpts, asset, amount, to)
}

// MainAuctionIterator is returned from FilterAuction and is used to iterate over the raw logs and unpacked data for Auction events raised by the Main contract.
type MainAuctionIterator struct {
	Event *MainAuction // Event containing the contract specifics and raw log

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
func (it *MainAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAuction)
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
		it.Event = new(MainAuction)
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
func (it *MainAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAuction represents a Auction event raised by the Main contract.
type MainAuction struct {
	User       common.Address
	Reserve    common.Address
	BidPrice   *big.Int
	NftAsset   common.Address
	NftTokenId *big.Int
	OnBehalfOf common.Address
	Borrower   common.Address
	LoanId     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuction is a free log retrieval operation binding the contract event 0xd4c7449fa0ea241233dd0a9e78a940879918f95e0caa34e0399a7d2813c8efba.
//
// Solidity: event Auction(address user, address indexed reserve, uint256 bidPrice, address indexed nftAsset, uint256 nftTokenId, address onBehalfOf, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) FilterAuction(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*MainAuctionIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Auction", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &MainAuctionIterator{contract: _Main.contract, event: "Auction", logs: logs, sub: sub}, nil
}

// WatchAuction is a free log subscription operation binding the contract event 0xd4c7449fa0ea241233dd0a9e78a940879918f95e0caa34e0399a7d2813c8efba.
//
// Solidity: event Auction(address user, address indexed reserve, uint256 bidPrice, address indexed nftAsset, uint256 nftTokenId, address onBehalfOf, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) WatchAuction(opts *bind.WatchOpts, sink chan<- *MainAuction, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Auction", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAuction)
				if err := _Main.contract.UnpackLog(event, "Auction", log); err != nil {
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

// ParseAuction is a log parse operation binding the contract event 0xd4c7449fa0ea241233dd0a9e78a940879918f95e0caa34e0399a7d2813c8efba.
//
// Solidity: event Auction(address user, address indexed reserve, uint256 bidPrice, address indexed nftAsset, uint256 nftTokenId, address onBehalfOf, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) ParseAuction(log types.Log) (*MainAuction, error) {
	event := new(MainAuction)
	if err := _Main.contract.UnpackLog(event, "Auction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Main contract.
type MainBorrowIterator struct {
	Event *MainBorrow // Event containing the contract specifics and raw log

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
func (it *MainBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainBorrow)
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
		it.Event = new(MainBorrow)
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
func (it *MainBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainBorrow represents a Borrow event raised by the Main contract.
type MainBorrow struct {
	User       common.Address
	Reserve    common.Address
	Amount     *big.Int
	NftAsset   common.Address
	NftTokenId *big.Int
	OnBehalfOf common.Address
	BorrowRate *big.Int
	LoanId     *big.Int
	Referral   uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xcfb3a669117d9dc90f0d3521228dc9fe67c5102dde205ef16fe9a1f81be698d5.
//
// Solidity: event Borrow(address user, address indexed reserve, uint256 amount, address nftAsset, uint256 nftTokenId, address indexed onBehalfOf, uint256 borrowRate, uint256 loanId, uint16 indexed referral)
func (_Main *MainFilterer) FilterBorrow(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (*MainBorrowIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return &MainBorrowIterator{contract: _Main.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xcfb3a669117d9dc90f0d3521228dc9fe67c5102dde205ef16fe9a1f81be698d5.
//
// Solidity: event Borrow(address user, address indexed reserve, uint256 amount, address nftAsset, uint256 nftTokenId, address indexed onBehalfOf, uint256 borrowRate, uint256 loanId, uint16 indexed referral)
func (_Main *MainFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *MainBorrow, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainBorrow)
				if err := _Main.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xcfb3a669117d9dc90f0d3521228dc9fe67c5102dde205ef16fe9a1f81be698d5.
//
// Solidity: event Borrow(address user, address indexed reserve, uint256 amount, address nftAsset, uint256 nftTokenId, address indexed onBehalfOf, uint256 borrowRate, uint256 loanId, uint16 indexed referral)
func (_Main *MainFilterer) ParseBorrow(log types.Log) (*MainBorrow, error) {
	event := new(MainBorrow)
	if err := _Main.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Main contract.
type MainDepositIterator struct {
	Event *MainDeposit // Event containing the contract specifics and raw log

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
func (it *MainDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainDeposit)
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
		it.Event = new(MainDeposit)
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
func (it *MainDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainDeposit represents a Deposit event raised by the Main contract.
type MainDeposit struct {
	User       common.Address
	Reserve    common.Address
	Amount     *big.Int
	OnBehalfOf common.Address
	Referral   uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x443ff2d25883a4800d36062db52ca3dd7ced05bd8627c8a6a37f8699715b5431.
//
// Solidity: event Deposit(address user, address indexed reserve, uint256 amount, address indexed onBehalfOf, uint16 indexed referral)
func (_Main *MainFilterer) FilterDeposit(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (*MainDepositIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return &MainDepositIterator{contract: _Main.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x443ff2d25883a4800d36062db52ca3dd7ced05bd8627c8a6a37f8699715b5431.
//
// Solidity: event Deposit(address user, address indexed reserve, uint256 amount, address indexed onBehalfOf, uint16 indexed referral)
func (_Main *MainFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MainDeposit, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainDeposit)
				if err := _Main.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x443ff2d25883a4800d36062db52ca3dd7ced05bd8627c8a6a37f8699715b5431.
//
// Solidity: event Deposit(address user, address indexed reserve, uint256 amount, address indexed onBehalfOf, uint16 indexed referral)
func (_Main *MainFilterer) ParseDeposit(log types.Log) (*MainDeposit, error) {
	event := new(MainDeposit)
	if err := _Main.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the Main contract.
type MainLiquidateIterator struct {
	Event *MainLiquidate // Event containing the contract specifics and raw log

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
func (it *MainLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainLiquidate)
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
		it.Event = new(MainLiquidate)
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
func (it *MainLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainLiquidate represents a Liquidate event raised by the Main contract.
type MainLiquidate struct {
	User         common.Address
	Reserve      common.Address
	RepayAmount  *big.Int
	RemainAmount *big.Int
	NftAsset     common.Address
	NftTokenId   *big.Int
	Borrower     common.Address
	LoanId       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0xf028795898a18c6fc88094dc5671c6a79d5dc3458c44015e9299fbc6c6268cf8.
//
// Solidity: event Liquidate(address user, address indexed reserve, uint256 repayAmount, uint256 remainAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) FilterLiquidate(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*MainLiquidateIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Liquidate", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &MainLiquidateIterator{contract: _Main.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0xf028795898a18c6fc88094dc5671c6a79d5dc3458c44015e9299fbc6c6268cf8.
//
// Solidity: event Liquidate(address user, address indexed reserve, uint256 repayAmount, uint256 remainAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *MainLiquidate, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Liquidate", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainLiquidate)
				if err := _Main.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0xf028795898a18c6fc88094dc5671c6a79d5dc3458c44015e9299fbc6c6268cf8.
//
// Solidity: event Liquidate(address user, address indexed reserve, uint256 repayAmount, uint256 remainAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) ParseLiquidate(log types.Log) (*MainLiquidate, error) {
	event := new(MainLiquidate)
	if err := _Main.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Main contract.
type MainPausedIterator struct {
	Event *MainPaused // Event containing the contract specifics and raw log

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
func (it *MainPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainPaused)
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
		it.Event = new(MainPaused)
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
func (it *MainPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainPaused represents a Paused event raised by the Main contract.
type MainPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Main *MainFilterer) FilterPaused(opts *bind.FilterOpts) (*MainPausedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MainPausedIterator{contract: _Main.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Main *MainFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MainPaused) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainPaused)
				if err := _Main.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Main *MainFilterer) ParsePaused(log types.Log) (*MainPaused, error) {
	event := new(MainPaused)
	if err := _Main.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainPausedTimeUpdatedIterator is returned from FilterPausedTimeUpdated and is used to iterate over the raw logs and unpacked data for PausedTimeUpdated events raised by the Main contract.
type MainPausedTimeUpdatedIterator struct {
	Event *MainPausedTimeUpdated // Event containing the contract specifics and raw log

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
func (it *MainPausedTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainPausedTimeUpdated)
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
		it.Event = new(MainPausedTimeUpdated)
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
func (it *MainPausedTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainPausedTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainPausedTimeUpdated represents a PausedTimeUpdated event raised by the Main contract.
type MainPausedTimeUpdated struct {
	StartTime    *big.Int
	DurationTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPausedTimeUpdated is a free log retrieval operation binding the contract event 0xd897a722b1c0a957941f99a13c0ea24d7d4ffafe0953658f68f49e13ccba5c5a.
//
// Solidity: event PausedTimeUpdated(uint256 startTime, uint256 durationTime)
func (_Main *MainFilterer) FilterPausedTimeUpdated(opts *bind.FilterOpts) (*MainPausedTimeUpdatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "PausedTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &MainPausedTimeUpdatedIterator{contract: _Main.contract, event: "PausedTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchPausedTimeUpdated is a free log subscription operation binding the contract event 0xd897a722b1c0a957941f99a13c0ea24d7d4ffafe0953658f68f49e13ccba5c5a.
//
// Solidity: event PausedTimeUpdated(uint256 startTime, uint256 durationTime)
func (_Main *MainFilterer) WatchPausedTimeUpdated(opts *bind.WatchOpts, sink chan<- *MainPausedTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "PausedTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainPausedTimeUpdated)
				if err := _Main.contract.UnpackLog(event, "PausedTimeUpdated", log); err != nil {
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

// ParsePausedTimeUpdated is a log parse operation binding the contract event 0xd897a722b1c0a957941f99a13c0ea24d7d4ffafe0953658f68f49e13ccba5c5a.
//
// Solidity: event PausedTimeUpdated(uint256 startTime, uint256 durationTime)
func (_Main *MainFilterer) ParsePausedTimeUpdated(log types.Log) (*MainPausedTimeUpdated, error) {
	event := new(MainPausedTimeUpdated)
	if err := _Main.contract.UnpackLog(event, "PausedTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Main contract.
type MainRedeemIterator struct {
	Event *MainRedeem // Event containing the contract specifics and raw log

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
func (it *MainRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRedeem)
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
		it.Event = new(MainRedeem)
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
func (it *MainRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRedeem represents a Redeem event raised by the Main contract.
type MainRedeem struct {
	User         common.Address
	Reserve      common.Address
	BorrowAmount *big.Int
	FineAmount   *big.Int
	NftAsset     common.Address
	NftTokenId   *big.Int
	Borrower     common.Address
	LoanId       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x0fcfe1a3f2afab13e32fa3c091795159ed5dfe66dc078e21c7f521f42e163afc.
//
// Solidity: event Redeem(address user, address indexed reserve, uint256 borrowAmount, uint256 fineAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) FilterRedeem(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*MainRedeemIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Redeem", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &MainRedeemIterator{contract: _Main.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x0fcfe1a3f2afab13e32fa3c091795159ed5dfe66dc078e21c7f521f42e163afc.
//
// Solidity: event Redeem(address user, address indexed reserve, uint256 borrowAmount, uint256 fineAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *MainRedeem, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Redeem", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRedeem)
				if err := _Main.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0x0fcfe1a3f2afab13e32fa3c091795159ed5dfe66dc078e21c7f521f42e163afc.
//
// Solidity: event Redeem(address user, address indexed reserve, uint256 borrowAmount, uint256 fineAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) ParseRedeem(log types.Log) (*MainRedeem, error) {
	event := new(MainRedeem)
	if err := _Main.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the Main contract.
type MainRepayIterator struct {
	Event *MainRepay // Event containing the contract specifics and raw log

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
func (it *MainRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRepay)
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
		it.Event = new(MainRepay)
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
func (it *MainRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRepay represents a Repay event raised by the Main contract.
type MainRepay struct {
	User       common.Address
	Reserve    common.Address
	Amount     *big.Int
	NftAsset   common.Address
	NftTokenId *big.Int
	Borrower   common.Address
	LoanId     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x50e03867c1178391f204f7bf0eb2f52d5167dc65a99a9650a95abe55d17be17e.
//
// Solidity: event Repay(address user, address indexed reserve, uint256 amount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) FilterRepay(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*MainRepayIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Repay", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &MainRepayIterator{contract: _Main.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x50e03867c1178391f204f7bf0eb2f52d5167dc65a99a9650a95abe55d17be17e.
//
// Solidity: event Repay(address user, address indexed reserve, uint256 amount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *MainRepay, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Repay", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRepay)
				if err := _Main.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x50e03867c1178391f204f7bf0eb2f52d5167dc65a99a9650a95abe55d17be17e.
//
// Solidity: event Repay(address user, address indexed reserve, uint256 amount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_Main *MainFilterer) ParseRepay(log types.Log) (*MainRepay, error) {
	event := new(MainRepay)
	if err := _Main.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainReserveDataUpdatedIterator is returned from FilterReserveDataUpdated and is used to iterate over the raw logs and unpacked data for ReserveDataUpdated events raised by the Main contract.
type MainReserveDataUpdatedIterator struct {
	Event *MainReserveDataUpdated // Event containing the contract specifics and raw log

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
func (it *MainReserveDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainReserveDataUpdated)
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
		it.Event = new(MainReserveDataUpdated)
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
func (it *MainReserveDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainReserveDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainReserveDataUpdated represents a ReserveDataUpdated event raised by the Main contract.
type MainReserveDataUpdated struct {
	Reserve             common.Address
	LiquidityRate       *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReserveDataUpdated is a free log retrieval operation binding the contract event 0x4063a2df84b66bb796eb32622851d833e57b2c4292900c18f963af8808b13e35.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_Main *MainFilterer) FilterReserveDataUpdated(opts *bind.FilterOpts, reserve []common.Address) (*MainReserveDataUpdatedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return &MainReserveDataUpdatedIterator{contract: _Main.contract, event: "ReserveDataUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveDataUpdated is a free log subscription operation binding the contract event 0x4063a2df84b66bb796eb32622851d833e57b2c4292900c18f963af8808b13e35.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_Main *MainFilterer) WatchReserveDataUpdated(opts *bind.WatchOpts, sink chan<- *MainReserveDataUpdated, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainReserveDataUpdated)
				if err := _Main.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
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

// ParseReserveDataUpdated is a log parse operation binding the contract event 0x4063a2df84b66bb796eb32622851d833e57b2c4292900c18f963af8808b13e35.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_Main *MainFilterer) ParseReserveDataUpdated(log types.Log) (*MainReserveDataUpdated, error) {
	event := new(MainReserveDataUpdated)
	if err := _Main.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Main contract.
type MainUnpausedIterator struct {
	Event *MainUnpaused // Event containing the contract specifics and raw log

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
func (it *MainUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainUnpaused)
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
		it.Event = new(MainUnpaused)
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
func (it *MainUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainUnpaused represents a Unpaused event raised by the Main contract.
type MainUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Main *MainFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MainUnpausedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MainUnpausedIterator{contract: _Main.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Main *MainFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MainUnpaused) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainUnpaused)
				if err := _Main.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Main *MainFilterer) ParseUnpaused(log types.Log) (*MainUnpaused, error) {
	event := new(MainUnpaused)
	if err := _Main.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Main contract.
type MainWithdrawIterator struct {
	Event *MainWithdraw // Event containing the contract specifics and raw log

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
func (it *MainWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainWithdraw)
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
		it.Event = new(MainWithdraw)
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
func (it *MainWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainWithdraw represents a Withdraw event raised by the Main contract.
type MainWithdraw struct {
	User    common.Address
	Reserve common.Address
	Amount  *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3ed4ee04a905a278b050a856bbe7ddaaf327a30514373e65aa6103beeae488c3.
//
// Solidity: event Withdraw(address indexed user, address indexed reserve, uint256 amount, address indexed to)
func (_Main *MainFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, reserve []common.Address, to []common.Address) (*MainWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Withdraw", userRule, reserveRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MainWithdrawIterator{contract: _Main.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3ed4ee04a905a278b050a856bbe7ddaaf327a30514373e65aa6103beeae488c3.
//
// Solidity: event Withdraw(address indexed user, address indexed reserve, uint256 amount, address indexed to)
func (_Main *MainFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MainWithdraw, user []common.Address, reserve []common.Address, to []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Withdraw", userRule, reserveRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainWithdraw)
				if err := _Main.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x3ed4ee04a905a278b050a856bbe7ddaaf327a30514373e65aa6103beeae488c3.
//
// Solidity: event Withdraw(address indexed user, address indexed reserve, uint256 amount, address indexed to)
func (_Main *MainFilterer) ParseWithdraw(log types.Log) (*MainWithdraw, error) {
	event := new(MainWithdraw)
	if err := _Main.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
