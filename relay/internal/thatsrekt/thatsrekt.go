// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package thatsrekt

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

// ThatsRektMetaData contains all meta data concerning the ThatsRekt contract.
var ThatsRektMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_ADDRESSES_PER_POST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_TITLE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_VIEW_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activePostsBefore\",\"inputs\":[{\"name\":\"beforeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ids\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addAttackers\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newAttackers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addVictims\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newVictims\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addWhitelisted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"amendNote\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newNote\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"amendTitle\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newTitle\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"attackerAppearances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attackerReport\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"score\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"appearances\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attackerScore\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirm\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enumThatsRekt.ConfirmDirection\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"confirmationOf\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumThatsRekt.ConfirmDirection\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfirmerCount\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfirmers\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDisconfirmerCount\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDisconfirmers\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPost\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"poster\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"attackedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"confirmations\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disconfirmations\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"removed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"attackers_\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"victims_\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"lastUpdatedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"headPostId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialWhitelistAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialWhitelistRemover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPurgeAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPurgeRemover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialWhitelisters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPurged\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isVictim\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isWhitelisted\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextPostId\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"peekNextPostId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"post\",\"inputs\":[{\"name\":\"expectedPostId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"title\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attackers_\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"victims_\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"note\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attackedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"postCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postTitle\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prevPostId\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"purgeAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"purgePost\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"purgeRemover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recentActivePosts\",\"inputs\":[{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ids\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeWhitelisted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"retract\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokePurgeAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeWhitelistAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPurgeAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPurgeRemover\",\"inputs\":[{\"name\":\"newRemover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWhitelistAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWhitelistRemover\",\"inputs\":[{\"name\":\"newRemover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tailPostId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unconfirm\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"whitelistAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistRemover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AttackersAdded\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAttackers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Confirmed\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"confirmer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldDirection\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumThatsRekt.ConfirmDirection\"},{\"name\":\"newDirection\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumThatsRekt.ConfirmDirection\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PostCreated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"poster\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"attackedAt\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"title\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"attackers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"victims\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"note\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PostNoteAmended\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newNote\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PostPurged\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PostRemoved\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumThatsRekt.RemovalReason\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PostTitleAmended\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newTitle\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PurgeAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PurgeRemoverTransferred\",\"inputs\":[{\"name\":\"previousRemover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRemover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VictimsAdded\",\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newVictims\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WhitelistAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WhitelistRemoverTransferred\",\"inputs\":[{\"name\":\"previousRemover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRemover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WhitelistUpdated\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AlreadyPurged\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAdditions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyPost\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAttackedAt\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfirmDirection\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoConfirmationChange\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPoster\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPurgeAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPurgeRemover\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistRemover\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NothingToUnconfirm\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PostIdMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PostIsPurged\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostIsRemoved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PosterCannotConfirm\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TitleEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TitleTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// ThatsRektABI is the input ABI used to generate the binding from.
// Deprecated: Use ThatsRektMetaData.ABI instead.
var ThatsRektABI = ThatsRektMetaData.ABI

// ThatsRekt is an auto generated Go binding around an Ethereum contract.
type ThatsRekt struct {
	ThatsRektCaller     // Read-only binding to the contract
	ThatsRektTransactor // Write-only binding to the contract
	ThatsRektFilterer   // Log filterer for contract events
}

// ThatsRektCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThatsRektCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThatsRektTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThatsRektTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThatsRektFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThatsRektFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThatsRektSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThatsRektSession struct {
	Contract     *ThatsRekt        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThatsRektCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThatsRektCallerSession struct {
	Contract *ThatsRektCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ThatsRektTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThatsRektTransactorSession struct {
	Contract     *ThatsRektTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ThatsRektRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThatsRektRaw struct {
	Contract *ThatsRekt // Generic contract binding to access the raw methods on
}

// ThatsRektCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThatsRektCallerRaw struct {
	Contract *ThatsRektCaller // Generic read-only contract binding to access the raw methods on
}

// ThatsRektTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThatsRektTransactorRaw struct {
	Contract *ThatsRektTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThatsRekt creates a new instance of ThatsRekt, bound to a specific deployed contract.
func NewThatsRekt(address common.Address, backend bind.ContractBackend) (*ThatsRekt, error) {
	contract, err := bindThatsRekt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ThatsRekt{ThatsRektCaller: ThatsRektCaller{contract: contract}, ThatsRektTransactor: ThatsRektTransactor{contract: contract}, ThatsRektFilterer: ThatsRektFilterer{contract: contract}}, nil
}

// NewThatsRektCaller creates a new read-only instance of ThatsRekt, bound to a specific deployed contract.
func NewThatsRektCaller(address common.Address, caller bind.ContractCaller) (*ThatsRektCaller, error) {
	contract, err := bindThatsRekt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThatsRektCaller{contract: contract}, nil
}

// NewThatsRektTransactor creates a new write-only instance of ThatsRekt, bound to a specific deployed contract.
func NewThatsRektTransactor(address common.Address, transactor bind.ContractTransactor) (*ThatsRektTransactor, error) {
	contract, err := bindThatsRekt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThatsRektTransactor{contract: contract}, nil
}

// NewThatsRektFilterer creates a new log filterer instance of ThatsRekt, bound to a specific deployed contract.
func NewThatsRektFilterer(address common.Address, filterer bind.ContractFilterer) (*ThatsRektFilterer, error) {
	contract, err := bindThatsRekt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThatsRektFilterer{contract: contract}, nil
}

// bindThatsRekt binds a generic wrapper to an already deployed contract.
func bindThatsRekt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ThatsRektMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThatsRekt *ThatsRektRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThatsRekt.Contract.ThatsRektCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThatsRekt *ThatsRektRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThatsRekt.Contract.ThatsRektTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThatsRekt *ThatsRektRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThatsRekt.Contract.ThatsRektTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThatsRekt *ThatsRektCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThatsRekt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThatsRekt *ThatsRektTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThatsRekt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThatsRekt *ThatsRektTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThatsRekt.Contract.contract.Transact(opts, method, params...)
}

// MAXADDRESSESPERPOST is a free data retrieval call binding the contract method 0x22d4164d.
//
// Solidity: function MAX_ADDRESSES_PER_POST() view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) MAXADDRESSESPERPOST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "MAX_ADDRESSES_PER_POST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXADDRESSESPERPOST is a free data retrieval call binding the contract method 0x22d4164d.
//
// Solidity: function MAX_ADDRESSES_PER_POST() view returns(uint256)
func (_ThatsRekt *ThatsRektSession) MAXADDRESSESPERPOST() (*big.Int, error) {
	return _ThatsRekt.Contract.MAXADDRESSESPERPOST(&_ThatsRekt.CallOpts)
}

// MAXADDRESSESPERPOST is a free data retrieval call binding the contract method 0x22d4164d.
//
// Solidity: function MAX_ADDRESSES_PER_POST() view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) MAXADDRESSESPERPOST() (*big.Int, error) {
	return _ThatsRekt.Contract.MAXADDRESSESPERPOST(&_ThatsRekt.CallOpts)
}

// MAXTITLELENGTH is a free data retrieval call binding the contract method 0x2ef9a160.
//
// Solidity: function MAX_TITLE_LENGTH() view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) MAXTITLELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "MAX_TITLE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTITLELENGTH is a free data retrieval call binding the contract method 0x2ef9a160.
//
// Solidity: function MAX_TITLE_LENGTH() view returns(uint256)
func (_ThatsRekt *ThatsRektSession) MAXTITLELENGTH() (*big.Int, error) {
	return _ThatsRekt.Contract.MAXTITLELENGTH(&_ThatsRekt.CallOpts)
}

// MAXTITLELENGTH is a free data retrieval call binding the contract method 0x2ef9a160.
//
// Solidity: function MAX_TITLE_LENGTH() view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) MAXTITLELENGTH() (*big.Int, error) {
	return _ThatsRekt.Contract.MAXTITLELENGTH(&_ThatsRekt.CallOpts)
}

// MAXVIEWLIMIT is a free data retrieval call binding the contract method 0x34fe4650.
//
// Solidity: function MAX_VIEW_LIMIT() view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) MAXVIEWLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "MAX_VIEW_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVIEWLIMIT is a free data retrieval call binding the contract method 0x34fe4650.
//
// Solidity: function MAX_VIEW_LIMIT() view returns(uint256)
func (_ThatsRekt *ThatsRektSession) MAXVIEWLIMIT() (*big.Int, error) {
	return _ThatsRekt.Contract.MAXVIEWLIMIT(&_ThatsRekt.CallOpts)
}

// MAXVIEWLIMIT is a free data retrieval call binding the contract method 0x34fe4650.
//
// Solidity: function MAX_VIEW_LIMIT() view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) MAXVIEWLIMIT() (*big.Int, error) {
	return _ThatsRekt.Contract.MAXVIEWLIMIT(&_ThatsRekt.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ThatsRekt *ThatsRektCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ThatsRekt *ThatsRektSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ThatsRekt.Contract.UPGRADEINTERFACEVERSION(&_ThatsRekt.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ThatsRekt *ThatsRektCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ThatsRekt.Contract.UPGRADEINTERFACEVERSION(&_ThatsRekt.CallOpts)
}

// ActivePostsBefore is a free data retrieval call binding the contract method 0x17a1e54d.
//
// Solidity: function activePostsBefore(uint256 beforeId, uint256 limit) view returns(uint256[] ids)
func (_ThatsRekt *ThatsRektCaller) ActivePostsBefore(opts *bind.CallOpts, beforeId *big.Int, limit *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "activePostsBefore", beforeId, limit)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ActivePostsBefore is a free data retrieval call binding the contract method 0x17a1e54d.
//
// Solidity: function activePostsBefore(uint256 beforeId, uint256 limit) view returns(uint256[] ids)
func (_ThatsRekt *ThatsRektSession) ActivePostsBefore(beforeId *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _ThatsRekt.Contract.ActivePostsBefore(&_ThatsRekt.CallOpts, beforeId, limit)
}

// ActivePostsBefore is a free data retrieval call binding the contract method 0x17a1e54d.
//
// Solidity: function activePostsBefore(uint256 beforeId, uint256 limit) view returns(uint256[] ids)
func (_ThatsRekt *ThatsRektCallerSession) ActivePostsBefore(beforeId *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _ThatsRekt.Contract.ActivePostsBefore(&_ThatsRekt.CallOpts, beforeId, limit)
}

// AttackerAppearances is a free data retrieval call binding the contract method 0x640c6395.
//
// Solidity: function attackerAppearances(address ) view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) AttackerAppearances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "attackerAppearances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttackerAppearances is a free data retrieval call binding the contract method 0x640c6395.
//
// Solidity: function attackerAppearances(address ) view returns(uint256)
func (_ThatsRekt *ThatsRektSession) AttackerAppearances(arg0 common.Address) (*big.Int, error) {
	return _ThatsRekt.Contract.AttackerAppearances(&_ThatsRekt.CallOpts, arg0)
}

// AttackerAppearances is a free data retrieval call binding the contract method 0x640c6395.
//
// Solidity: function attackerAppearances(address ) view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) AttackerAppearances(arg0 common.Address) (*big.Int, error) {
	return _ThatsRekt.Contract.AttackerAppearances(&_ThatsRekt.CallOpts, arg0)
}

// AttackerReport is a free data retrieval call binding the contract method 0x07363ce8.
//
// Solidity: function attackerReport(address a) view returns(int256 score, uint256 appearances)
func (_ThatsRekt *ThatsRektCaller) AttackerReport(opts *bind.CallOpts, a common.Address) (struct {
	Score       *big.Int
	Appearances *big.Int
}, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "attackerReport", a)

	outstruct := new(struct {
		Score       *big.Int
		Appearances *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Score = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Appearances = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AttackerReport is a free data retrieval call binding the contract method 0x07363ce8.
//
// Solidity: function attackerReport(address a) view returns(int256 score, uint256 appearances)
func (_ThatsRekt *ThatsRektSession) AttackerReport(a common.Address) (struct {
	Score       *big.Int
	Appearances *big.Int
}, error) {
	return _ThatsRekt.Contract.AttackerReport(&_ThatsRekt.CallOpts, a)
}

// AttackerReport is a free data retrieval call binding the contract method 0x07363ce8.
//
// Solidity: function attackerReport(address a) view returns(int256 score, uint256 appearances)
func (_ThatsRekt *ThatsRektCallerSession) AttackerReport(a common.Address) (struct {
	Score       *big.Int
	Appearances *big.Int
}, error) {
	return _ThatsRekt.Contract.AttackerReport(&_ThatsRekt.CallOpts, a)
}

// AttackerScore is a free data retrieval call binding the contract method 0x6559e955.
//
// Solidity: function attackerScore(address ) view returns(int256)
func (_ThatsRekt *ThatsRektCaller) AttackerScore(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "attackerScore", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttackerScore is a free data retrieval call binding the contract method 0x6559e955.
//
// Solidity: function attackerScore(address ) view returns(int256)
func (_ThatsRekt *ThatsRektSession) AttackerScore(arg0 common.Address) (*big.Int, error) {
	return _ThatsRekt.Contract.AttackerScore(&_ThatsRekt.CallOpts, arg0)
}

// AttackerScore is a free data retrieval call binding the contract method 0x6559e955.
//
// Solidity: function attackerScore(address ) view returns(int256)
func (_ThatsRekt *ThatsRektCallerSession) AttackerScore(arg0 common.Address) (*big.Int, error) {
	return _ThatsRekt.Contract.AttackerScore(&_ThatsRekt.CallOpts, arg0)
}

// ConfirmationOf is a free data retrieval call binding the contract method 0xb1fd994b.
//
// Solidity: function confirmationOf(uint256 , address ) view returns(uint8)
func (_ThatsRekt *ThatsRektCaller) ConfirmationOf(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (uint8, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "confirmationOf", arg0, arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ConfirmationOf is a free data retrieval call binding the contract method 0xb1fd994b.
//
// Solidity: function confirmationOf(uint256 , address ) view returns(uint8)
func (_ThatsRekt *ThatsRektSession) ConfirmationOf(arg0 *big.Int, arg1 common.Address) (uint8, error) {
	return _ThatsRekt.Contract.ConfirmationOf(&_ThatsRekt.CallOpts, arg0, arg1)
}

// ConfirmationOf is a free data retrieval call binding the contract method 0xb1fd994b.
//
// Solidity: function confirmationOf(uint256 , address ) view returns(uint8)
func (_ThatsRekt *ThatsRektCallerSession) ConfirmationOf(arg0 *big.Int, arg1 common.Address) (uint8, error) {
	return _ThatsRekt.Contract.ConfirmationOf(&_ThatsRekt.CallOpts, arg0, arg1)
}

// GetConfirmerCount is a free data retrieval call binding the contract method 0xa8beb1da.
//
// Solidity: function getConfirmerCount(uint256 postId) view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) GetConfirmerCount(opts *bind.CallOpts, postId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "getConfirmerCount", postId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfirmerCount is a free data retrieval call binding the contract method 0xa8beb1da.
//
// Solidity: function getConfirmerCount(uint256 postId) view returns(uint256)
func (_ThatsRekt *ThatsRektSession) GetConfirmerCount(postId *big.Int) (*big.Int, error) {
	return _ThatsRekt.Contract.GetConfirmerCount(&_ThatsRekt.CallOpts, postId)
}

// GetConfirmerCount is a free data retrieval call binding the contract method 0xa8beb1da.
//
// Solidity: function getConfirmerCount(uint256 postId) view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) GetConfirmerCount(postId *big.Int) (*big.Int, error) {
	return _ThatsRekt.Contract.GetConfirmerCount(&_ThatsRekt.CallOpts, postId)
}

// GetConfirmers is a free data retrieval call binding the contract method 0xb779e1c0.
//
// Solidity: function getConfirmers(uint256 postId) view returns(address[])
func (_ThatsRekt *ThatsRektCaller) GetConfirmers(opts *bind.CallOpts, postId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "getConfirmers", postId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConfirmers is a free data retrieval call binding the contract method 0xb779e1c0.
//
// Solidity: function getConfirmers(uint256 postId) view returns(address[])
func (_ThatsRekt *ThatsRektSession) GetConfirmers(postId *big.Int) ([]common.Address, error) {
	return _ThatsRekt.Contract.GetConfirmers(&_ThatsRekt.CallOpts, postId)
}

// GetConfirmers is a free data retrieval call binding the contract method 0xb779e1c0.
//
// Solidity: function getConfirmers(uint256 postId) view returns(address[])
func (_ThatsRekt *ThatsRektCallerSession) GetConfirmers(postId *big.Int) ([]common.Address, error) {
	return _ThatsRekt.Contract.GetConfirmers(&_ThatsRekt.CallOpts, postId)
}

// GetDisconfirmerCount is a free data retrieval call binding the contract method 0x59dafd3c.
//
// Solidity: function getDisconfirmerCount(uint256 postId) view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) GetDisconfirmerCount(opts *bind.CallOpts, postId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "getDisconfirmerCount", postId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisconfirmerCount is a free data retrieval call binding the contract method 0x59dafd3c.
//
// Solidity: function getDisconfirmerCount(uint256 postId) view returns(uint256)
func (_ThatsRekt *ThatsRektSession) GetDisconfirmerCount(postId *big.Int) (*big.Int, error) {
	return _ThatsRekt.Contract.GetDisconfirmerCount(&_ThatsRekt.CallOpts, postId)
}

// GetDisconfirmerCount is a free data retrieval call binding the contract method 0x59dafd3c.
//
// Solidity: function getDisconfirmerCount(uint256 postId) view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) GetDisconfirmerCount(postId *big.Int) (*big.Int, error) {
	return _ThatsRekt.Contract.GetDisconfirmerCount(&_ThatsRekt.CallOpts, postId)
}

// GetDisconfirmers is a free data retrieval call binding the contract method 0x7e7d5daf.
//
// Solidity: function getDisconfirmers(uint256 postId) view returns(address[])
func (_ThatsRekt *ThatsRektCaller) GetDisconfirmers(opts *bind.CallOpts, postId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "getDisconfirmers", postId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDisconfirmers is a free data retrieval call binding the contract method 0x7e7d5daf.
//
// Solidity: function getDisconfirmers(uint256 postId) view returns(address[])
func (_ThatsRekt *ThatsRektSession) GetDisconfirmers(postId *big.Int) ([]common.Address, error) {
	return _ThatsRekt.Contract.GetDisconfirmers(&_ThatsRekt.CallOpts, postId)
}

// GetDisconfirmers is a free data retrieval call binding the contract method 0x7e7d5daf.
//
// Solidity: function getDisconfirmers(uint256 postId) view returns(address[])
func (_ThatsRekt *ThatsRektCallerSession) GetDisconfirmers(postId *big.Int) ([]common.Address, error) {
	return _ThatsRekt.Contract.GetDisconfirmers(&_ThatsRekt.CallOpts, postId)
}

// GetPost is a free data retrieval call binding the contract method 0x40731c24.
//
// Solidity: function getPost(uint256 id) view returns(address poster, uint64 attackedAt, uint32 confirmations, uint32 disconfirmations, bool removed, address[] attackers_, address[] victims_, uint64 lastUpdatedAt)
func (_ThatsRekt *ThatsRektCaller) GetPost(opts *bind.CallOpts, id *big.Int) (struct {
	Poster           common.Address
	AttackedAt       uint64
	Confirmations    uint32
	Disconfirmations uint32
	Removed          bool
	Attackers        []common.Address
	Victims          []common.Address
	LastUpdatedAt    uint64
}, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "getPost", id)

	outstruct := new(struct {
		Poster           common.Address
		AttackedAt       uint64
		Confirmations    uint32
		Disconfirmations uint32
		Removed          bool
		Attackers        []common.Address
		Victims          []common.Address
		LastUpdatedAt    uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Poster = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AttackedAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Confirmations = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Disconfirmations = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Removed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Attackers = *abi.ConvertType(out[5], new([]common.Address)).(*[]common.Address)
	outstruct.Victims = *abi.ConvertType(out[6], new([]common.Address)).(*[]common.Address)
	outstruct.LastUpdatedAt = *abi.ConvertType(out[7], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetPost is a free data retrieval call binding the contract method 0x40731c24.
//
// Solidity: function getPost(uint256 id) view returns(address poster, uint64 attackedAt, uint32 confirmations, uint32 disconfirmations, bool removed, address[] attackers_, address[] victims_, uint64 lastUpdatedAt)
func (_ThatsRekt *ThatsRektSession) GetPost(id *big.Int) (struct {
	Poster           common.Address
	AttackedAt       uint64
	Confirmations    uint32
	Disconfirmations uint32
	Removed          bool
	Attackers        []common.Address
	Victims          []common.Address
	LastUpdatedAt    uint64
}, error) {
	return _ThatsRekt.Contract.GetPost(&_ThatsRekt.CallOpts, id)
}

// GetPost is a free data retrieval call binding the contract method 0x40731c24.
//
// Solidity: function getPost(uint256 id) view returns(address poster, uint64 attackedAt, uint32 confirmations, uint32 disconfirmations, bool removed, address[] attackers_, address[] victims_, uint64 lastUpdatedAt)
func (_ThatsRekt *ThatsRektCallerSession) GetPost(id *big.Int) (struct {
	Poster           common.Address
	AttackedAt       uint64
	Confirmations    uint32
	Disconfirmations uint32
	Removed          bool
	Attackers        []common.Address
	Victims          []common.Address
	LastUpdatedAt    uint64
}, error) {
	return _ThatsRekt.Contract.GetPost(&_ThatsRekt.CallOpts, id)
}

// HeadPostId is a free data retrieval call binding the contract method 0xf93b72e1.
//
// Solidity: function headPostId() view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) HeadPostId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "headPostId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HeadPostId is a free data retrieval call binding the contract method 0xf93b72e1.
//
// Solidity: function headPostId() view returns(uint256)
func (_ThatsRekt *ThatsRektSession) HeadPostId() (*big.Int, error) {
	return _ThatsRekt.Contract.HeadPostId(&_ThatsRekt.CallOpts)
}

// HeadPostId is a free data retrieval call binding the contract method 0xf93b72e1.
//
// Solidity: function headPostId() view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) HeadPostId() (*big.Int, error) {
	return _ThatsRekt.Contract.HeadPostId(&_ThatsRekt.CallOpts)
}

// IsPurged is a free data retrieval call binding the contract method 0x327990bb.
//
// Solidity: function isPurged(uint256 postId) view returns(bool)
func (_ThatsRekt *ThatsRektCaller) IsPurged(opts *bind.CallOpts, postId *big.Int) (bool, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "isPurged", postId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPurged is a free data retrieval call binding the contract method 0x327990bb.
//
// Solidity: function isPurged(uint256 postId) view returns(bool)
func (_ThatsRekt *ThatsRektSession) IsPurged(postId *big.Int) (bool, error) {
	return _ThatsRekt.Contract.IsPurged(&_ThatsRekt.CallOpts, postId)
}

// IsPurged is a free data retrieval call binding the contract method 0x327990bb.
//
// Solidity: function isPurged(uint256 postId) view returns(bool)
func (_ThatsRekt *ThatsRektCallerSession) IsPurged(postId *big.Int) (bool, error) {
	return _ThatsRekt.Contract.IsPurged(&_ThatsRekt.CallOpts, postId)
}

// IsVictim is a free data retrieval call binding the contract method 0x2d10cc3d.
//
// Solidity: function isVictim(address ) view returns(bool)
func (_ThatsRekt *ThatsRektCaller) IsVictim(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "isVictim", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVictim is a free data retrieval call binding the contract method 0x2d10cc3d.
//
// Solidity: function isVictim(address ) view returns(bool)
func (_ThatsRekt *ThatsRektSession) IsVictim(arg0 common.Address) (bool, error) {
	return _ThatsRekt.Contract.IsVictim(&_ThatsRekt.CallOpts, arg0)
}

// IsVictim is a free data retrieval call binding the contract method 0x2d10cc3d.
//
// Solidity: function isVictim(address ) view returns(bool)
func (_ThatsRekt *ThatsRektCallerSession) IsVictim(arg0 common.Address) (bool, error) {
	return _ThatsRekt.Contract.IsVictim(&_ThatsRekt.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_ThatsRekt *ThatsRektCaller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "isWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_ThatsRekt *ThatsRektSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _ThatsRekt.Contract.IsWhitelisted(&_ThatsRekt.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_ThatsRekt *ThatsRektCallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _ThatsRekt.Contract.IsWhitelisted(&_ThatsRekt.CallOpts, arg0)
}

// NextPostId is a free data retrieval call binding the contract method 0x932375e9.
//
// Solidity: function nextPostId(uint256 ) view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) NextPostId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "nextPostId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPostId is a free data retrieval call binding the contract method 0x932375e9.
//
// Solidity: function nextPostId(uint256 ) view returns(uint256)
func (_ThatsRekt *ThatsRektSession) NextPostId(arg0 *big.Int) (*big.Int, error) {
	return _ThatsRekt.Contract.NextPostId(&_ThatsRekt.CallOpts, arg0)
}

// NextPostId is a free data retrieval call binding the contract method 0x932375e9.
//
// Solidity: function nextPostId(uint256 ) view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) NextPostId(arg0 *big.Int) (*big.Int, error) {
	return _ThatsRekt.Contract.NextPostId(&_ThatsRekt.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ThatsRekt *ThatsRektCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ThatsRekt *ThatsRektSession) Owner() (common.Address, error) {
	return _ThatsRekt.Contract.Owner(&_ThatsRekt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ThatsRekt *ThatsRektCallerSession) Owner() (common.Address, error) {
	return _ThatsRekt.Contract.Owner(&_ThatsRekt.CallOpts)
}

// PeekNextPostId is a free data retrieval call binding the contract method 0x8b46945a.
//
// Solidity: function peekNextPostId() view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) PeekNextPostId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "peekNextPostId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeekNextPostId is a free data retrieval call binding the contract method 0x8b46945a.
//
// Solidity: function peekNextPostId() view returns(uint256)
func (_ThatsRekt *ThatsRektSession) PeekNextPostId() (*big.Int, error) {
	return _ThatsRekt.Contract.PeekNextPostId(&_ThatsRekt.CallOpts)
}

// PeekNextPostId is a free data retrieval call binding the contract method 0x8b46945a.
//
// Solidity: function peekNextPostId() view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) PeekNextPostId() (*big.Int, error) {
	return _ThatsRekt.Contract.PeekNextPostId(&_ThatsRekt.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ThatsRekt *ThatsRektCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ThatsRekt *ThatsRektSession) PendingOwner() (common.Address, error) {
	return _ThatsRekt.Contract.PendingOwner(&_ThatsRekt.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ThatsRekt *ThatsRektCallerSession) PendingOwner() (common.Address, error) {
	return _ThatsRekt.Contract.PendingOwner(&_ThatsRekt.CallOpts)
}

// PostCount is a free data retrieval call binding the contract method 0x17906c2e.
//
// Solidity: function postCount() view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) PostCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "postCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostCount is a free data retrieval call binding the contract method 0x17906c2e.
//
// Solidity: function postCount() view returns(uint256)
func (_ThatsRekt *ThatsRektSession) PostCount() (*big.Int, error) {
	return _ThatsRekt.Contract.PostCount(&_ThatsRekt.CallOpts)
}

// PostCount is a free data retrieval call binding the contract method 0x17906c2e.
//
// Solidity: function postCount() view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) PostCount() (*big.Int, error) {
	return _ThatsRekt.Contract.PostCount(&_ThatsRekt.CallOpts)
}

// PostTitle is a free data retrieval call binding the contract method 0xc3765c32.
//
// Solidity: function postTitle(uint256 ) view returns(string)
func (_ThatsRekt *ThatsRektCaller) PostTitle(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "postTitle", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PostTitle is a free data retrieval call binding the contract method 0xc3765c32.
//
// Solidity: function postTitle(uint256 ) view returns(string)
func (_ThatsRekt *ThatsRektSession) PostTitle(arg0 *big.Int) (string, error) {
	return _ThatsRekt.Contract.PostTitle(&_ThatsRekt.CallOpts, arg0)
}

// PostTitle is a free data retrieval call binding the contract method 0xc3765c32.
//
// Solidity: function postTitle(uint256 ) view returns(string)
func (_ThatsRekt *ThatsRektCallerSession) PostTitle(arg0 *big.Int) (string, error) {
	return _ThatsRekt.Contract.PostTitle(&_ThatsRekt.CallOpts, arg0)
}

// PrevPostId is a free data retrieval call binding the contract method 0xe3c91286.
//
// Solidity: function prevPostId(uint256 ) view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) PrevPostId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "prevPostId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrevPostId is a free data retrieval call binding the contract method 0xe3c91286.
//
// Solidity: function prevPostId(uint256 ) view returns(uint256)
func (_ThatsRekt *ThatsRektSession) PrevPostId(arg0 *big.Int) (*big.Int, error) {
	return _ThatsRekt.Contract.PrevPostId(&_ThatsRekt.CallOpts, arg0)
}

// PrevPostId is a free data retrieval call binding the contract method 0xe3c91286.
//
// Solidity: function prevPostId(uint256 ) view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) PrevPostId(arg0 *big.Int) (*big.Int, error) {
	return _ThatsRekt.Contract.PrevPostId(&_ThatsRekt.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ThatsRekt *ThatsRektCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ThatsRekt *ThatsRektSession) ProxiableUUID() ([32]byte, error) {
	return _ThatsRekt.Contract.ProxiableUUID(&_ThatsRekt.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ThatsRekt *ThatsRektCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ThatsRekt.Contract.ProxiableUUID(&_ThatsRekt.CallOpts)
}

// PurgeAdmin is a free data retrieval call binding the contract method 0x06aae84c.
//
// Solidity: function purgeAdmin() view returns(address)
func (_ThatsRekt *ThatsRektCaller) PurgeAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "purgeAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PurgeAdmin is a free data retrieval call binding the contract method 0x06aae84c.
//
// Solidity: function purgeAdmin() view returns(address)
func (_ThatsRekt *ThatsRektSession) PurgeAdmin() (common.Address, error) {
	return _ThatsRekt.Contract.PurgeAdmin(&_ThatsRekt.CallOpts)
}

// PurgeAdmin is a free data retrieval call binding the contract method 0x06aae84c.
//
// Solidity: function purgeAdmin() view returns(address)
func (_ThatsRekt *ThatsRektCallerSession) PurgeAdmin() (common.Address, error) {
	return _ThatsRekt.Contract.PurgeAdmin(&_ThatsRekt.CallOpts)
}

// PurgeRemover is a free data retrieval call binding the contract method 0xdf52dc43.
//
// Solidity: function purgeRemover() view returns(address)
func (_ThatsRekt *ThatsRektCaller) PurgeRemover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "purgeRemover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PurgeRemover is a free data retrieval call binding the contract method 0xdf52dc43.
//
// Solidity: function purgeRemover() view returns(address)
func (_ThatsRekt *ThatsRektSession) PurgeRemover() (common.Address, error) {
	return _ThatsRekt.Contract.PurgeRemover(&_ThatsRekt.CallOpts)
}

// PurgeRemover is a free data retrieval call binding the contract method 0xdf52dc43.
//
// Solidity: function purgeRemover() view returns(address)
func (_ThatsRekt *ThatsRektCallerSession) PurgeRemover() (common.Address, error) {
	return _ThatsRekt.Contract.PurgeRemover(&_ThatsRekt.CallOpts)
}

// RecentActivePosts is a free data retrieval call binding the contract method 0xccde89ce.
//
// Solidity: function recentActivePosts(uint256 limit) view returns(uint256[] ids)
func (_ThatsRekt *ThatsRektCaller) RecentActivePosts(opts *bind.CallOpts, limit *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "recentActivePosts", limit)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RecentActivePosts is a free data retrieval call binding the contract method 0xccde89ce.
//
// Solidity: function recentActivePosts(uint256 limit) view returns(uint256[] ids)
func (_ThatsRekt *ThatsRektSession) RecentActivePosts(limit *big.Int) ([]*big.Int, error) {
	return _ThatsRekt.Contract.RecentActivePosts(&_ThatsRekt.CallOpts, limit)
}

// RecentActivePosts is a free data retrieval call binding the contract method 0xccde89ce.
//
// Solidity: function recentActivePosts(uint256 limit) view returns(uint256[] ids)
func (_ThatsRekt *ThatsRektCallerSession) RecentActivePosts(limit *big.Int) ([]*big.Int, error) {
	return _ThatsRekt.Contract.RecentActivePosts(&_ThatsRekt.CallOpts, limit)
}

// TailPostId is a free data retrieval call binding the contract method 0x8094c914.
//
// Solidity: function tailPostId() view returns(uint256)
func (_ThatsRekt *ThatsRektCaller) TailPostId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "tailPostId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TailPostId is a free data retrieval call binding the contract method 0x8094c914.
//
// Solidity: function tailPostId() view returns(uint256)
func (_ThatsRekt *ThatsRektSession) TailPostId() (*big.Int, error) {
	return _ThatsRekt.Contract.TailPostId(&_ThatsRekt.CallOpts)
}

// TailPostId is a free data retrieval call binding the contract method 0x8094c914.
//
// Solidity: function tailPostId() view returns(uint256)
func (_ThatsRekt *ThatsRektCallerSession) TailPostId() (*big.Int, error) {
	return _ThatsRekt.Contract.TailPostId(&_ThatsRekt.CallOpts)
}

// WhitelistAdmin is a free data retrieval call binding the contract method 0x4adbe551.
//
// Solidity: function whitelistAdmin() view returns(address)
func (_ThatsRekt *ThatsRektCaller) WhitelistAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "whitelistAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistAdmin is a free data retrieval call binding the contract method 0x4adbe551.
//
// Solidity: function whitelistAdmin() view returns(address)
func (_ThatsRekt *ThatsRektSession) WhitelistAdmin() (common.Address, error) {
	return _ThatsRekt.Contract.WhitelistAdmin(&_ThatsRekt.CallOpts)
}

// WhitelistAdmin is a free data retrieval call binding the contract method 0x4adbe551.
//
// Solidity: function whitelistAdmin() view returns(address)
func (_ThatsRekt *ThatsRektCallerSession) WhitelistAdmin() (common.Address, error) {
	return _ThatsRekt.Contract.WhitelistAdmin(&_ThatsRekt.CallOpts)
}

// WhitelistRemover is a free data retrieval call binding the contract method 0x68aa349f.
//
// Solidity: function whitelistRemover() view returns(address)
func (_ThatsRekt *ThatsRektCaller) WhitelistRemover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ThatsRekt.contract.Call(opts, &out, "whitelistRemover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistRemover is a free data retrieval call binding the contract method 0x68aa349f.
//
// Solidity: function whitelistRemover() view returns(address)
func (_ThatsRekt *ThatsRektSession) WhitelistRemover() (common.Address, error) {
	return _ThatsRekt.Contract.WhitelistRemover(&_ThatsRekt.CallOpts)
}

// WhitelistRemover is a free data retrieval call binding the contract method 0x68aa349f.
//
// Solidity: function whitelistRemover() view returns(address)
func (_ThatsRekt *ThatsRektCallerSession) WhitelistRemover() (common.Address, error) {
	return _ThatsRekt.Contract.WhitelistRemover(&_ThatsRekt.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ThatsRekt *ThatsRektTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ThatsRekt *ThatsRektSession) AcceptOwnership() (*types.Transaction, error) {
	return _ThatsRekt.Contract.AcceptOwnership(&_ThatsRekt.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ThatsRekt *ThatsRektTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ThatsRekt.Contract.AcceptOwnership(&_ThatsRekt.TransactOpts)
}

// AddAttackers is a paid mutator transaction binding the contract method 0x34d46b53.
//
// Solidity: function addAttackers(uint256 postId, address[] newAttackers) returns()
func (_ThatsRekt *ThatsRektTransactor) AddAttackers(opts *bind.TransactOpts, postId *big.Int, newAttackers []common.Address) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "addAttackers", postId, newAttackers)
}

// AddAttackers is a paid mutator transaction binding the contract method 0x34d46b53.
//
// Solidity: function addAttackers(uint256 postId, address[] newAttackers) returns()
func (_ThatsRekt *ThatsRektSession) AddAttackers(postId *big.Int, newAttackers []common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.AddAttackers(&_ThatsRekt.TransactOpts, postId, newAttackers)
}

// AddAttackers is a paid mutator transaction binding the contract method 0x34d46b53.
//
// Solidity: function addAttackers(uint256 postId, address[] newAttackers) returns()
func (_ThatsRekt *ThatsRektTransactorSession) AddAttackers(postId *big.Int, newAttackers []common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.AddAttackers(&_ThatsRekt.TransactOpts, postId, newAttackers)
}

// AddVictims is a paid mutator transaction binding the contract method 0x17ada0b0.
//
// Solidity: function addVictims(uint256 postId, address[] newVictims) returns()
func (_ThatsRekt *ThatsRektTransactor) AddVictims(opts *bind.TransactOpts, postId *big.Int, newVictims []common.Address) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "addVictims", postId, newVictims)
}

// AddVictims is a paid mutator transaction binding the contract method 0x17ada0b0.
//
// Solidity: function addVictims(uint256 postId, address[] newVictims) returns()
func (_ThatsRekt *ThatsRektSession) AddVictims(postId *big.Int, newVictims []common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.AddVictims(&_ThatsRekt.TransactOpts, postId, newVictims)
}

// AddVictims is a paid mutator transaction binding the contract method 0x17ada0b0.
//
// Solidity: function addVictims(uint256 postId, address[] newVictims) returns()
func (_ThatsRekt *ThatsRektTransactorSession) AddVictims(postId *big.Int, newVictims []common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.AddVictims(&_ThatsRekt.TransactOpts, postId, newVictims)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_ThatsRekt *ThatsRektTransactor) AddWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "addWhitelisted", account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_ThatsRekt *ThatsRektSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.AddWhitelisted(&_ThatsRekt.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_ThatsRekt *ThatsRektTransactorSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.AddWhitelisted(&_ThatsRekt.TransactOpts, account)
}

// AmendNote is a paid mutator transaction binding the contract method 0x5ef7c714.
//
// Solidity: function amendNote(uint256 postId, string newNote) returns()
func (_ThatsRekt *ThatsRektTransactor) AmendNote(opts *bind.TransactOpts, postId *big.Int, newNote string) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "amendNote", postId, newNote)
}

// AmendNote is a paid mutator transaction binding the contract method 0x5ef7c714.
//
// Solidity: function amendNote(uint256 postId, string newNote) returns()
func (_ThatsRekt *ThatsRektSession) AmendNote(postId *big.Int, newNote string) (*types.Transaction, error) {
	return _ThatsRekt.Contract.AmendNote(&_ThatsRekt.TransactOpts, postId, newNote)
}

// AmendNote is a paid mutator transaction binding the contract method 0x5ef7c714.
//
// Solidity: function amendNote(uint256 postId, string newNote) returns()
func (_ThatsRekt *ThatsRektTransactorSession) AmendNote(postId *big.Int, newNote string) (*types.Transaction, error) {
	return _ThatsRekt.Contract.AmendNote(&_ThatsRekt.TransactOpts, postId, newNote)
}

// AmendTitle is a paid mutator transaction binding the contract method 0x6cd0fc27.
//
// Solidity: function amendTitle(uint256 postId, string newTitle) returns()
func (_ThatsRekt *ThatsRektTransactor) AmendTitle(opts *bind.TransactOpts, postId *big.Int, newTitle string) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "amendTitle", postId, newTitle)
}

// AmendTitle is a paid mutator transaction binding the contract method 0x6cd0fc27.
//
// Solidity: function amendTitle(uint256 postId, string newTitle) returns()
func (_ThatsRekt *ThatsRektSession) AmendTitle(postId *big.Int, newTitle string) (*types.Transaction, error) {
	return _ThatsRekt.Contract.AmendTitle(&_ThatsRekt.TransactOpts, postId, newTitle)
}

// AmendTitle is a paid mutator transaction binding the contract method 0x6cd0fc27.
//
// Solidity: function amendTitle(uint256 postId, string newTitle) returns()
func (_ThatsRekt *ThatsRektTransactorSession) AmendTitle(postId *big.Int, newTitle string) (*types.Transaction, error) {
	return _ThatsRekt.Contract.AmendTitle(&_ThatsRekt.TransactOpts, postId, newTitle)
}

// Confirm is a paid mutator transaction binding the contract method 0x13a34b28.
//
// Solidity: function confirm(uint256 postId, uint8 direction) returns()
func (_ThatsRekt *ThatsRektTransactor) Confirm(opts *bind.TransactOpts, postId *big.Int, direction uint8) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "confirm", postId, direction)
}

// Confirm is a paid mutator transaction binding the contract method 0x13a34b28.
//
// Solidity: function confirm(uint256 postId, uint8 direction) returns()
func (_ThatsRekt *ThatsRektSession) Confirm(postId *big.Int, direction uint8) (*types.Transaction, error) {
	return _ThatsRekt.Contract.Confirm(&_ThatsRekt.TransactOpts, postId, direction)
}

// Confirm is a paid mutator transaction binding the contract method 0x13a34b28.
//
// Solidity: function confirm(uint256 postId, uint8 direction) returns()
func (_ThatsRekt *ThatsRektTransactorSession) Confirm(postId *big.Int, direction uint8) (*types.Transaction, error) {
	return _ThatsRekt.Contract.Confirm(&_ThatsRekt.TransactOpts, postId, direction)
}

// Initialize is a paid mutator transaction binding the contract method 0xb5df06b4.
//
// Solidity: function initialize(address initialOwner, address initialWhitelistAdmin, address initialWhitelistRemover, address initialPurgeAdmin, address initialPurgeRemover, address[] initialWhitelisters) returns()
func (_ThatsRekt *ThatsRektTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialWhitelistAdmin common.Address, initialWhitelistRemover common.Address, initialPurgeAdmin common.Address, initialPurgeRemover common.Address, initialWhitelisters []common.Address) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "initialize", initialOwner, initialWhitelistAdmin, initialWhitelistRemover, initialPurgeAdmin, initialPurgeRemover, initialWhitelisters)
}

// Initialize is a paid mutator transaction binding the contract method 0xb5df06b4.
//
// Solidity: function initialize(address initialOwner, address initialWhitelistAdmin, address initialWhitelistRemover, address initialPurgeAdmin, address initialPurgeRemover, address[] initialWhitelisters) returns()
func (_ThatsRekt *ThatsRektSession) Initialize(initialOwner common.Address, initialWhitelistAdmin common.Address, initialWhitelistRemover common.Address, initialPurgeAdmin common.Address, initialPurgeRemover common.Address, initialWhitelisters []common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.Initialize(&_ThatsRekt.TransactOpts, initialOwner, initialWhitelistAdmin, initialWhitelistRemover, initialPurgeAdmin, initialPurgeRemover, initialWhitelisters)
}

// Initialize is a paid mutator transaction binding the contract method 0xb5df06b4.
//
// Solidity: function initialize(address initialOwner, address initialWhitelistAdmin, address initialWhitelistRemover, address initialPurgeAdmin, address initialPurgeRemover, address[] initialWhitelisters) returns()
func (_ThatsRekt *ThatsRektTransactorSession) Initialize(initialOwner common.Address, initialWhitelistAdmin common.Address, initialWhitelistRemover common.Address, initialPurgeAdmin common.Address, initialPurgeRemover common.Address, initialWhitelisters []common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.Initialize(&_ThatsRekt.TransactOpts, initialOwner, initialWhitelistAdmin, initialWhitelistRemover, initialPurgeAdmin, initialPurgeRemover, initialWhitelisters)
}

// Post is a paid mutator transaction binding the contract method 0x65141a5d.
//
// Solidity: function post(uint256 expectedPostId, string title, address[] attackers_, address[] victims_, string note, uint64 attackedAt) returns(uint256 id)
func (_ThatsRekt *ThatsRektTransactor) Post(opts *bind.TransactOpts, expectedPostId *big.Int, title string, attackers_ []common.Address, victims_ []common.Address, note string, attackedAt uint64) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "post", expectedPostId, title, attackers_, victims_, note, attackedAt)
}

// Post is a paid mutator transaction binding the contract method 0x65141a5d.
//
// Solidity: function post(uint256 expectedPostId, string title, address[] attackers_, address[] victims_, string note, uint64 attackedAt) returns(uint256 id)
func (_ThatsRekt *ThatsRektSession) Post(expectedPostId *big.Int, title string, attackers_ []common.Address, victims_ []common.Address, note string, attackedAt uint64) (*types.Transaction, error) {
	return _ThatsRekt.Contract.Post(&_ThatsRekt.TransactOpts, expectedPostId, title, attackers_, victims_, note, attackedAt)
}

// Post is a paid mutator transaction binding the contract method 0x65141a5d.
//
// Solidity: function post(uint256 expectedPostId, string title, address[] attackers_, address[] victims_, string note, uint64 attackedAt) returns(uint256 id)
func (_ThatsRekt *ThatsRektTransactorSession) Post(expectedPostId *big.Int, title string, attackers_ []common.Address, victims_ []common.Address, note string, attackedAt uint64) (*types.Transaction, error) {
	return _ThatsRekt.Contract.Post(&_ThatsRekt.TransactOpts, expectedPostId, title, attackers_, victims_, note, attackedAt)
}

// PurgePost is a paid mutator transaction binding the contract method 0x18ec7472.
//
// Solidity: function purgePost(uint256 postId) returns()
func (_ThatsRekt *ThatsRektTransactor) PurgePost(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "purgePost", postId)
}

// PurgePost is a paid mutator transaction binding the contract method 0x18ec7472.
//
// Solidity: function purgePost(uint256 postId) returns()
func (_ThatsRekt *ThatsRektSession) PurgePost(postId *big.Int) (*types.Transaction, error) {
	return _ThatsRekt.Contract.PurgePost(&_ThatsRekt.TransactOpts, postId)
}

// PurgePost is a paid mutator transaction binding the contract method 0x18ec7472.
//
// Solidity: function purgePost(uint256 postId) returns()
func (_ThatsRekt *ThatsRektTransactorSession) PurgePost(postId *big.Int) (*types.Transaction, error) {
	return _ThatsRekt.Contract.PurgePost(&_ThatsRekt.TransactOpts, postId)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_ThatsRekt *ThatsRektTransactor) RemoveWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "removeWhitelisted", account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_ThatsRekt *ThatsRektSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.RemoveWhitelisted(&_ThatsRekt.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_ThatsRekt *ThatsRektTransactorSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.RemoveWhitelisted(&_ThatsRekt.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ThatsRekt *ThatsRektTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ThatsRekt *ThatsRektSession) RenounceOwnership() (*types.Transaction, error) {
	return _ThatsRekt.Contract.RenounceOwnership(&_ThatsRekt.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ThatsRekt *ThatsRektTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ThatsRekt.Contract.RenounceOwnership(&_ThatsRekt.TransactOpts)
}

// Retract is a paid mutator transaction binding the contract method 0x9fab6656.
//
// Solidity: function retract(uint256 postId) returns()
func (_ThatsRekt *ThatsRektTransactor) Retract(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "retract", postId)
}

// Retract is a paid mutator transaction binding the contract method 0x9fab6656.
//
// Solidity: function retract(uint256 postId) returns()
func (_ThatsRekt *ThatsRektSession) Retract(postId *big.Int) (*types.Transaction, error) {
	return _ThatsRekt.Contract.Retract(&_ThatsRekt.TransactOpts, postId)
}

// Retract is a paid mutator transaction binding the contract method 0x9fab6656.
//
// Solidity: function retract(uint256 postId) returns()
func (_ThatsRekt *ThatsRektTransactorSession) Retract(postId *big.Int) (*types.Transaction, error) {
	return _ThatsRekt.Contract.Retract(&_ThatsRekt.TransactOpts, postId)
}

// RevokePurgeAdmin is a paid mutator transaction binding the contract method 0xf6765ecd.
//
// Solidity: function revokePurgeAdmin() returns()
func (_ThatsRekt *ThatsRektTransactor) RevokePurgeAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "revokePurgeAdmin")
}

// RevokePurgeAdmin is a paid mutator transaction binding the contract method 0xf6765ecd.
//
// Solidity: function revokePurgeAdmin() returns()
func (_ThatsRekt *ThatsRektSession) RevokePurgeAdmin() (*types.Transaction, error) {
	return _ThatsRekt.Contract.RevokePurgeAdmin(&_ThatsRekt.TransactOpts)
}

// RevokePurgeAdmin is a paid mutator transaction binding the contract method 0xf6765ecd.
//
// Solidity: function revokePurgeAdmin() returns()
func (_ThatsRekt *ThatsRektTransactorSession) RevokePurgeAdmin() (*types.Transaction, error) {
	return _ThatsRekt.Contract.RevokePurgeAdmin(&_ThatsRekt.TransactOpts)
}

// RevokeWhitelistAdmin is a paid mutator transaction binding the contract method 0x7492b233.
//
// Solidity: function revokeWhitelistAdmin() returns()
func (_ThatsRekt *ThatsRektTransactor) RevokeWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "revokeWhitelistAdmin")
}

// RevokeWhitelistAdmin is a paid mutator transaction binding the contract method 0x7492b233.
//
// Solidity: function revokeWhitelistAdmin() returns()
func (_ThatsRekt *ThatsRektSession) RevokeWhitelistAdmin() (*types.Transaction, error) {
	return _ThatsRekt.Contract.RevokeWhitelistAdmin(&_ThatsRekt.TransactOpts)
}

// RevokeWhitelistAdmin is a paid mutator transaction binding the contract method 0x7492b233.
//
// Solidity: function revokeWhitelistAdmin() returns()
func (_ThatsRekt *ThatsRektTransactorSession) RevokeWhitelistAdmin() (*types.Transaction, error) {
	return _ThatsRekt.Contract.RevokeWhitelistAdmin(&_ThatsRekt.TransactOpts)
}

// SetPurgeAdmin is a paid mutator transaction binding the contract method 0x888852d5.
//
// Solidity: function setPurgeAdmin(address newAdmin) returns()
func (_ThatsRekt *ThatsRektTransactor) SetPurgeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "setPurgeAdmin", newAdmin)
}

// SetPurgeAdmin is a paid mutator transaction binding the contract method 0x888852d5.
//
// Solidity: function setPurgeAdmin(address newAdmin) returns()
func (_ThatsRekt *ThatsRektSession) SetPurgeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.SetPurgeAdmin(&_ThatsRekt.TransactOpts, newAdmin)
}

// SetPurgeAdmin is a paid mutator transaction binding the contract method 0x888852d5.
//
// Solidity: function setPurgeAdmin(address newAdmin) returns()
func (_ThatsRekt *ThatsRektTransactorSession) SetPurgeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.SetPurgeAdmin(&_ThatsRekt.TransactOpts, newAdmin)
}

// SetPurgeRemover is a paid mutator transaction binding the contract method 0xa9bba35d.
//
// Solidity: function setPurgeRemover(address newRemover) returns()
func (_ThatsRekt *ThatsRektTransactor) SetPurgeRemover(opts *bind.TransactOpts, newRemover common.Address) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "setPurgeRemover", newRemover)
}

// SetPurgeRemover is a paid mutator transaction binding the contract method 0xa9bba35d.
//
// Solidity: function setPurgeRemover(address newRemover) returns()
func (_ThatsRekt *ThatsRektSession) SetPurgeRemover(newRemover common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.SetPurgeRemover(&_ThatsRekt.TransactOpts, newRemover)
}

// SetPurgeRemover is a paid mutator transaction binding the contract method 0xa9bba35d.
//
// Solidity: function setPurgeRemover(address newRemover) returns()
func (_ThatsRekt *ThatsRektTransactorSession) SetPurgeRemover(newRemover common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.SetPurgeRemover(&_ThatsRekt.TransactOpts, newRemover)
}

// SetWhitelistAdmin is a paid mutator transaction binding the contract method 0xa8d49e64.
//
// Solidity: function setWhitelistAdmin(address newAdmin) returns()
func (_ThatsRekt *ThatsRektTransactor) SetWhitelistAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "setWhitelistAdmin", newAdmin)
}

// SetWhitelistAdmin is a paid mutator transaction binding the contract method 0xa8d49e64.
//
// Solidity: function setWhitelistAdmin(address newAdmin) returns()
func (_ThatsRekt *ThatsRektSession) SetWhitelistAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.SetWhitelistAdmin(&_ThatsRekt.TransactOpts, newAdmin)
}

// SetWhitelistAdmin is a paid mutator transaction binding the contract method 0xa8d49e64.
//
// Solidity: function setWhitelistAdmin(address newAdmin) returns()
func (_ThatsRekt *ThatsRektTransactorSession) SetWhitelistAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.SetWhitelistAdmin(&_ThatsRekt.TransactOpts, newAdmin)
}

// SetWhitelistRemover is a paid mutator transaction binding the contract method 0xba1b6ffd.
//
// Solidity: function setWhitelistRemover(address newRemover) returns()
func (_ThatsRekt *ThatsRektTransactor) SetWhitelistRemover(opts *bind.TransactOpts, newRemover common.Address) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "setWhitelistRemover", newRemover)
}

// SetWhitelistRemover is a paid mutator transaction binding the contract method 0xba1b6ffd.
//
// Solidity: function setWhitelistRemover(address newRemover) returns()
func (_ThatsRekt *ThatsRektSession) SetWhitelistRemover(newRemover common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.SetWhitelistRemover(&_ThatsRekt.TransactOpts, newRemover)
}

// SetWhitelistRemover is a paid mutator transaction binding the contract method 0xba1b6ffd.
//
// Solidity: function setWhitelistRemover(address newRemover) returns()
func (_ThatsRekt *ThatsRektTransactorSession) SetWhitelistRemover(newRemover common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.SetWhitelistRemover(&_ThatsRekt.TransactOpts, newRemover)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ThatsRekt *ThatsRektTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ThatsRekt *ThatsRektSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.TransferOwnership(&_ThatsRekt.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ThatsRekt *ThatsRektTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ThatsRekt.Contract.TransferOwnership(&_ThatsRekt.TransactOpts, newOwner)
}

// Unconfirm is a paid mutator transaction binding the contract method 0xec132ce3.
//
// Solidity: function unconfirm(uint256 postId) returns()
func (_ThatsRekt *ThatsRektTransactor) Unconfirm(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "unconfirm", postId)
}

// Unconfirm is a paid mutator transaction binding the contract method 0xec132ce3.
//
// Solidity: function unconfirm(uint256 postId) returns()
func (_ThatsRekt *ThatsRektSession) Unconfirm(postId *big.Int) (*types.Transaction, error) {
	return _ThatsRekt.Contract.Unconfirm(&_ThatsRekt.TransactOpts, postId)
}

// Unconfirm is a paid mutator transaction binding the contract method 0xec132ce3.
//
// Solidity: function unconfirm(uint256 postId) returns()
func (_ThatsRekt *ThatsRektTransactorSession) Unconfirm(postId *big.Int) (*types.Transaction, error) {
	return _ThatsRekt.Contract.Unconfirm(&_ThatsRekt.TransactOpts, postId)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ThatsRekt *ThatsRektTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ThatsRekt.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ThatsRekt *ThatsRektSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ThatsRekt.Contract.UpgradeToAndCall(&_ThatsRekt.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ThatsRekt *ThatsRektTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ThatsRekt.Contract.UpgradeToAndCall(&_ThatsRekt.TransactOpts, newImplementation, data)
}

// ThatsRektAttackersAddedIterator is returned from FilterAttackersAdded and is used to iterate over the raw logs and unpacked data for AttackersAdded events raised by the ThatsRekt contract.
type ThatsRektAttackersAddedIterator struct {
	Event *ThatsRektAttackersAdded // Event containing the contract specifics and raw log

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
func (it *ThatsRektAttackersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektAttackersAdded)
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
		it.Event = new(ThatsRektAttackersAdded)
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
func (it *ThatsRektAttackersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektAttackersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektAttackersAdded represents a AttackersAdded event raised by the ThatsRekt contract.
type ThatsRektAttackersAdded struct {
	PostId       *big.Int
	Amender      common.Address
	NewAttackers []common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttackersAdded is a free log retrieval operation binding the contract event 0x11e33fe659ce20067cdcc1c90a3b342aa497e29cfcb732a3eaddd3a2d3c39bb4.
//
// Solidity: event AttackersAdded(uint256 indexed postId, address indexed amender, address[] newAttackers)
func (_ThatsRekt *ThatsRektFilterer) FilterAttackersAdded(opts *bind.FilterOpts, postId []*big.Int, amender []common.Address) (*ThatsRektAttackersAddedIterator, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var amenderRule []interface{}
	for _, amenderItem := range amender {
		amenderRule = append(amenderRule, amenderItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "AttackersAdded", postIdRule, amenderRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektAttackersAddedIterator{contract: _ThatsRekt.contract, event: "AttackersAdded", logs: logs, sub: sub}, nil
}

// WatchAttackersAdded is a free log subscription operation binding the contract event 0x11e33fe659ce20067cdcc1c90a3b342aa497e29cfcb732a3eaddd3a2d3c39bb4.
//
// Solidity: event AttackersAdded(uint256 indexed postId, address indexed amender, address[] newAttackers)
func (_ThatsRekt *ThatsRektFilterer) WatchAttackersAdded(opts *bind.WatchOpts, sink chan<- *ThatsRektAttackersAdded, postId []*big.Int, amender []common.Address) (event.Subscription, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var amenderRule []interface{}
	for _, amenderItem := range amender {
		amenderRule = append(amenderRule, amenderItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "AttackersAdded", postIdRule, amenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektAttackersAdded)
				if err := _ThatsRekt.contract.UnpackLog(event, "AttackersAdded", log); err != nil {
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

// ParseAttackersAdded is a log parse operation binding the contract event 0x11e33fe659ce20067cdcc1c90a3b342aa497e29cfcb732a3eaddd3a2d3c39bb4.
//
// Solidity: event AttackersAdded(uint256 indexed postId, address indexed amender, address[] newAttackers)
func (_ThatsRekt *ThatsRektFilterer) ParseAttackersAdded(log types.Log) (*ThatsRektAttackersAdded, error) {
	event := new(ThatsRektAttackersAdded)
	if err := _ThatsRekt.contract.UnpackLog(event, "AttackersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektConfirmedIterator is returned from FilterConfirmed and is used to iterate over the raw logs and unpacked data for Confirmed events raised by the ThatsRekt contract.
type ThatsRektConfirmedIterator struct {
	Event *ThatsRektConfirmed // Event containing the contract specifics and raw log

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
func (it *ThatsRektConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektConfirmed)
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
		it.Event = new(ThatsRektConfirmed)
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
func (it *ThatsRektConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektConfirmed represents a Confirmed event raised by the ThatsRekt contract.
type ThatsRektConfirmed struct {
	PostId       *big.Int
	Confirmer    common.Address
	OldDirection uint8
	NewDirection uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterConfirmed is a free log retrieval operation binding the contract event 0x097ad824d639a422c3d350f244fc0f8c23b116959ede9dbc119114e0fae20222.
//
// Solidity: event Confirmed(uint256 indexed postId, address indexed confirmer, uint8 oldDirection, uint8 newDirection)
func (_ThatsRekt *ThatsRektFilterer) FilterConfirmed(opts *bind.FilterOpts, postId []*big.Int, confirmer []common.Address) (*ThatsRektConfirmedIterator, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var confirmerRule []interface{}
	for _, confirmerItem := range confirmer {
		confirmerRule = append(confirmerRule, confirmerItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "Confirmed", postIdRule, confirmerRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektConfirmedIterator{contract: _ThatsRekt.contract, event: "Confirmed", logs: logs, sub: sub}, nil
}

// WatchConfirmed is a free log subscription operation binding the contract event 0x097ad824d639a422c3d350f244fc0f8c23b116959ede9dbc119114e0fae20222.
//
// Solidity: event Confirmed(uint256 indexed postId, address indexed confirmer, uint8 oldDirection, uint8 newDirection)
func (_ThatsRekt *ThatsRektFilterer) WatchConfirmed(opts *bind.WatchOpts, sink chan<- *ThatsRektConfirmed, postId []*big.Int, confirmer []common.Address) (event.Subscription, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var confirmerRule []interface{}
	for _, confirmerItem := range confirmer {
		confirmerRule = append(confirmerRule, confirmerItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "Confirmed", postIdRule, confirmerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektConfirmed)
				if err := _ThatsRekt.contract.UnpackLog(event, "Confirmed", log); err != nil {
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

// ParseConfirmed is a log parse operation binding the contract event 0x097ad824d639a422c3d350f244fc0f8c23b116959ede9dbc119114e0fae20222.
//
// Solidity: event Confirmed(uint256 indexed postId, address indexed confirmer, uint8 oldDirection, uint8 newDirection)
func (_ThatsRekt *ThatsRektFilterer) ParseConfirmed(log types.Log) (*ThatsRektConfirmed, error) {
	event := new(ThatsRektConfirmed)
	if err := _ThatsRekt.contract.UnpackLog(event, "Confirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ThatsRekt contract.
type ThatsRektInitializedIterator struct {
	Event *ThatsRektInitialized // Event containing the contract specifics and raw log

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
func (it *ThatsRektInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektInitialized)
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
		it.Event = new(ThatsRektInitialized)
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
func (it *ThatsRektInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektInitialized represents a Initialized event raised by the ThatsRekt contract.
type ThatsRektInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ThatsRekt *ThatsRektFilterer) FilterInitialized(opts *bind.FilterOpts) (*ThatsRektInitializedIterator, error) {

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ThatsRektInitializedIterator{contract: _ThatsRekt.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ThatsRekt *ThatsRektFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ThatsRektInitialized) (event.Subscription, error) {

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektInitialized)
				if err := _ThatsRekt.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ThatsRekt *ThatsRektFilterer) ParseInitialized(log types.Log) (*ThatsRektInitialized, error) {
	event := new(ThatsRektInitialized)
	if err := _ThatsRekt.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the ThatsRekt contract.
type ThatsRektOwnershipTransferStartedIterator struct {
	Event *ThatsRektOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *ThatsRektOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektOwnershipTransferStarted)
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
		it.Event = new(ThatsRektOwnershipTransferStarted)
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
func (it *ThatsRektOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the ThatsRekt contract.
type ThatsRektOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ThatsRekt *ThatsRektFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ThatsRektOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektOwnershipTransferStartedIterator{contract: _ThatsRekt.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ThatsRekt *ThatsRektFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *ThatsRektOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektOwnershipTransferStarted)
				if err := _ThatsRekt.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ThatsRekt *ThatsRektFilterer) ParseOwnershipTransferStarted(log types.Log) (*ThatsRektOwnershipTransferStarted, error) {
	event := new(ThatsRektOwnershipTransferStarted)
	if err := _ThatsRekt.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ThatsRekt contract.
type ThatsRektOwnershipTransferredIterator struct {
	Event *ThatsRektOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ThatsRektOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektOwnershipTransferred)
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
		it.Event = new(ThatsRektOwnershipTransferred)
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
func (it *ThatsRektOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektOwnershipTransferred represents a OwnershipTransferred event raised by the ThatsRekt contract.
type ThatsRektOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ThatsRekt *ThatsRektFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ThatsRektOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektOwnershipTransferredIterator{contract: _ThatsRekt.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ThatsRekt *ThatsRektFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ThatsRektOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektOwnershipTransferred)
				if err := _ThatsRekt.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ThatsRekt *ThatsRektFilterer) ParseOwnershipTransferred(log types.Log) (*ThatsRektOwnershipTransferred, error) {
	event := new(ThatsRektOwnershipTransferred)
	if err := _ThatsRekt.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektPostCreatedIterator is returned from FilterPostCreated and is used to iterate over the raw logs and unpacked data for PostCreated events raised by the ThatsRekt contract.
type ThatsRektPostCreatedIterator struct {
	Event *ThatsRektPostCreated // Event containing the contract specifics and raw log

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
func (it *ThatsRektPostCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektPostCreated)
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
		it.Event = new(ThatsRektPostCreated)
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
func (it *ThatsRektPostCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektPostCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektPostCreated represents a PostCreated event raised by the ThatsRekt contract.
type ThatsRektPostCreated struct {
	Id         *big.Int
	Poster     common.Address
	AttackedAt uint64
	Title      string
	Attackers  []common.Address
	Victims    []common.Address
	Note       string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPostCreated is a free log retrieval operation binding the contract event 0x0d7ae440ca52974e1d1ce2edd77d29270dc9dd3ae72340834b3424a93e4998a2.
//
// Solidity: event PostCreated(uint256 indexed id, address indexed poster, uint64 attackedAt, string title, address[] attackers, address[] victims, string note)
func (_ThatsRekt *ThatsRektFilterer) FilterPostCreated(opts *bind.FilterOpts, id []*big.Int, poster []common.Address) (*ThatsRektPostCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var posterRule []interface{}
	for _, posterItem := range poster {
		posterRule = append(posterRule, posterItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "PostCreated", idRule, posterRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektPostCreatedIterator{contract: _ThatsRekt.contract, event: "PostCreated", logs: logs, sub: sub}, nil
}

// WatchPostCreated is a free log subscription operation binding the contract event 0x0d7ae440ca52974e1d1ce2edd77d29270dc9dd3ae72340834b3424a93e4998a2.
//
// Solidity: event PostCreated(uint256 indexed id, address indexed poster, uint64 attackedAt, string title, address[] attackers, address[] victims, string note)
func (_ThatsRekt *ThatsRektFilterer) WatchPostCreated(opts *bind.WatchOpts, sink chan<- *ThatsRektPostCreated, id []*big.Int, poster []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var posterRule []interface{}
	for _, posterItem := range poster {
		posterRule = append(posterRule, posterItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "PostCreated", idRule, posterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektPostCreated)
				if err := _ThatsRekt.contract.UnpackLog(event, "PostCreated", log); err != nil {
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

// ParsePostCreated is a log parse operation binding the contract event 0x0d7ae440ca52974e1d1ce2edd77d29270dc9dd3ae72340834b3424a93e4998a2.
//
// Solidity: event PostCreated(uint256 indexed id, address indexed poster, uint64 attackedAt, string title, address[] attackers, address[] victims, string note)
func (_ThatsRekt *ThatsRektFilterer) ParsePostCreated(log types.Log) (*ThatsRektPostCreated, error) {
	event := new(ThatsRektPostCreated)
	if err := _ThatsRekt.contract.UnpackLog(event, "PostCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektPostNoteAmendedIterator is returned from FilterPostNoteAmended and is used to iterate over the raw logs and unpacked data for PostNoteAmended events raised by the ThatsRekt contract.
type ThatsRektPostNoteAmendedIterator struct {
	Event *ThatsRektPostNoteAmended // Event containing the contract specifics and raw log

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
func (it *ThatsRektPostNoteAmendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektPostNoteAmended)
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
		it.Event = new(ThatsRektPostNoteAmended)
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
func (it *ThatsRektPostNoteAmendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektPostNoteAmendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektPostNoteAmended represents a PostNoteAmended event raised by the ThatsRekt contract.
type ThatsRektPostNoteAmended struct {
	PostId  *big.Int
	Amender common.Address
	NewNote string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPostNoteAmended is a free log retrieval operation binding the contract event 0x6b4b6748b092a36f538b5d936f48f9e52910f5b77b05297c90560423a14bb25c.
//
// Solidity: event PostNoteAmended(uint256 indexed postId, address indexed amender, string newNote)
func (_ThatsRekt *ThatsRektFilterer) FilterPostNoteAmended(opts *bind.FilterOpts, postId []*big.Int, amender []common.Address) (*ThatsRektPostNoteAmendedIterator, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var amenderRule []interface{}
	for _, amenderItem := range amender {
		amenderRule = append(amenderRule, amenderItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "PostNoteAmended", postIdRule, amenderRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektPostNoteAmendedIterator{contract: _ThatsRekt.contract, event: "PostNoteAmended", logs: logs, sub: sub}, nil
}

// WatchPostNoteAmended is a free log subscription operation binding the contract event 0x6b4b6748b092a36f538b5d936f48f9e52910f5b77b05297c90560423a14bb25c.
//
// Solidity: event PostNoteAmended(uint256 indexed postId, address indexed amender, string newNote)
func (_ThatsRekt *ThatsRektFilterer) WatchPostNoteAmended(opts *bind.WatchOpts, sink chan<- *ThatsRektPostNoteAmended, postId []*big.Int, amender []common.Address) (event.Subscription, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var amenderRule []interface{}
	for _, amenderItem := range amender {
		amenderRule = append(amenderRule, amenderItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "PostNoteAmended", postIdRule, amenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektPostNoteAmended)
				if err := _ThatsRekt.contract.UnpackLog(event, "PostNoteAmended", log); err != nil {
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

// ParsePostNoteAmended is a log parse operation binding the contract event 0x6b4b6748b092a36f538b5d936f48f9e52910f5b77b05297c90560423a14bb25c.
//
// Solidity: event PostNoteAmended(uint256 indexed postId, address indexed amender, string newNote)
func (_ThatsRekt *ThatsRektFilterer) ParsePostNoteAmended(log types.Log) (*ThatsRektPostNoteAmended, error) {
	event := new(ThatsRektPostNoteAmended)
	if err := _ThatsRekt.contract.UnpackLog(event, "PostNoteAmended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektPostPurgedIterator is returned from FilterPostPurged and is used to iterate over the raw logs and unpacked data for PostPurged events raised by the ThatsRekt contract.
type ThatsRektPostPurgedIterator struct {
	Event *ThatsRektPostPurged // Event containing the contract specifics and raw log

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
func (it *ThatsRektPostPurgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektPostPurged)
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
		it.Event = new(ThatsRektPostPurged)
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
func (it *ThatsRektPostPurgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektPostPurgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektPostPurged represents a PostPurged event raised by the ThatsRekt contract.
type ThatsRektPostPurged struct {
	PostId *big.Int
	By     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPostPurged is a free log retrieval operation binding the contract event 0x70a991e519cee55704541465ce57e817537363a613fcfc859739c03d70c235c3.
//
// Solidity: event PostPurged(uint256 indexed postId, address indexed by)
func (_ThatsRekt *ThatsRektFilterer) FilterPostPurged(opts *bind.FilterOpts, postId []*big.Int, by []common.Address) (*ThatsRektPostPurgedIterator, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "PostPurged", postIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektPostPurgedIterator{contract: _ThatsRekt.contract, event: "PostPurged", logs: logs, sub: sub}, nil
}

// WatchPostPurged is a free log subscription operation binding the contract event 0x70a991e519cee55704541465ce57e817537363a613fcfc859739c03d70c235c3.
//
// Solidity: event PostPurged(uint256 indexed postId, address indexed by)
func (_ThatsRekt *ThatsRektFilterer) WatchPostPurged(opts *bind.WatchOpts, sink chan<- *ThatsRektPostPurged, postId []*big.Int, by []common.Address) (event.Subscription, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "PostPurged", postIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektPostPurged)
				if err := _ThatsRekt.contract.UnpackLog(event, "PostPurged", log); err != nil {
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

// ParsePostPurged is a log parse operation binding the contract event 0x70a991e519cee55704541465ce57e817537363a613fcfc859739c03d70c235c3.
//
// Solidity: event PostPurged(uint256 indexed postId, address indexed by)
func (_ThatsRekt *ThatsRektFilterer) ParsePostPurged(log types.Log) (*ThatsRektPostPurged, error) {
	event := new(ThatsRektPostPurged)
	if err := _ThatsRekt.contract.UnpackLog(event, "PostPurged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektPostRemovedIterator is returned from FilterPostRemoved and is used to iterate over the raw logs and unpacked data for PostRemoved events raised by the ThatsRekt contract.
type ThatsRektPostRemovedIterator struct {
	Event *ThatsRektPostRemoved // Event containing the contract specifics and raw log

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
func (it *ThatsRektPostRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektPostRemoved)
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
		it.Event = new(ThatsRektPostRemoved)
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
func (it *ThatsRektPostRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektPostRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektPostRemoved represents a PostRemoved event raised by the ThatsRekt contract.
type ThatsRektPostRemoved struct {
	PostId *big.Int
	Reason uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPostRemoved is a free log retrieval operation binding the contract event 0x5718ae2ef8a84a4ac1944e4db68da2c2f99b2367a583836f2032da026b358c80.
//
// Solidity: event PostRemoved(uint256 indexed postId, uint8 reason)
func (_ThatsRekt *ThatsRektFilterer) FilterPostRemoved(opts *bind.FilterOpts, postId []*big.Int) (*ThatsRektPostRemovedIterator, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "PostRemoved", postIdRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektPostRemovedIterator{contract: _ThatsRekt.contract, event: "PostRemoved", logs: logs, sub: sub}, nil
}

// WatchPostRemoved is a free log subscription operation binding the contract event 0x5718ae2ef8a84a4ac1944e4db68da2c2f99b2367a583836f2032da026b358c80.
//
// Solidity: event PostRemoved(uint256 indexed postId, uint8 reason)
func (_ThatsRekt *ThatsRektFilterer) WatchPostRemoved(opts *bind.WatchOpts, sink chan<- *ThatsRektPostRemoved, postId []*big.Int) (event.Subscription, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "PostRemoved", postIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektPostRemoved)
				if err := _ThatsRekt.contract.UnpackLog(event, "PostRemoved", log); err != nil {
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

// ParsePostRemoved is a log parse operation binding the contract event 0x5718ae2ef8a84a4ac1944e4db68da2c2f99b2367a583836f2032da026b358c80.
//
// Solidity: event PostRemoved(uint256 indexed postId, uint8 reason)
func (_ThatsRekt *ThatsRektFilterer) ParsePostRemoved(log types.Log) (*ThatsRektPostRemoved, error) {
	event := new(ThatsRektPostRemoved)
	if err := _ThatsRekt.contract.UnpackLog(event, "PostRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektPostTitleAmendedIterator is returned from FilterPostTitleAmended and is used to iterate over the raw logs and unpacked data for PostTitleAmended events raised by the ThatsRekt contract.
type ThatsRektPostTitleAmendedIterator struct {
	Event *ThatsRektPostTitleAmended // Event containing the contract specifics and raw log

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
func (it *ThatsRektPostTitleAmendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektPostTitleAmended)
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
		it.Event = new(ThatsRektPostTitleAmended)
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
func (it *ThatsRektPostTitleAmendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektPostTitleAmendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektPostTitleAmended represents a PostTitleAmended event raised by the ThatsRekt contract.
type ThatsRektPostTitleAmended struct {
	PostId   *big.Int
	Amender  common.Address
	NewTitle string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPostTitleAmended is a free log retrieval operation binding the contract event 0xaae225037103bba935ab52a59332ced3e456790237b9b71dc31ce4357a9cdb6c.
//
// Solidity: event PostTitleAmended(uint256 indexed postId, address indexed amender, string newTitle)
func (_ThatsRekt *ThatsRektFilterer) FilterPostTitleAmended(opts *bind.FilterOpts, postId []*big.Int, amender []common.Address) (*ThatsRektPostTitleAmendedIterator, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var amenderRule []interface{}
	for _, amenderItem := range amender {
		amenderRule = append(amenderRule, amenderItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "PostTitleAmended", postIdRule, amenderRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektPostTitleAmendedIterator{contract: _ThatsRekt.contract, event: "PostTitleAmended", logs: logs, sub: sub}, nil
}

// WatchPostTitleAmended is a free log subscription operation binding the contract event 0xaae225037103bba935ab52a59332ced3e456790237b9b71dc31ce4357a9cdb6c.
//
// Solidity: event PostTitleAmended(uint256 indexed postId, address indexed amender, string newTitle)
func (_ThatsRekt *ThatsRektFilterer) WatchPostTitleAmended(opts *bind.WatchOpts, sink chan<- *ThatsRektPostTitleAmended, postId []*big.Int, amender []common.Address) (event.Subscription, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var amenderRule []interface{}
	for _, amenderItem := range amender {
		amenderRule = append(amenderRule, amenderItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "PostTitleAmended", postIdRule, amenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektPostTitleAmended)
				if err := _ThatsRekt.contract.UnpackLog(event, "PostTitleAmended", log); err != nil {
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

// ParsePostTitleAmended is a log parse operation binding the contract event 0xaae225037103bba935ab52a59332ced3e456790237b9b71dc31ce4357a9cdb6c.
//
// Solidity: event PostTitleAmended(uint256 indexed postId, address indexed amender, string newTitle)
func (_ThatsRekt *ThatsRektFilterer) ParsePostTitleAmended(log types.Log) (*ThatsRektPostTitleAmended, error) {
	event := new(ThatsRektPostTitleAmended)
	if err := _ThatsRekt.contract.UnpackLog(event, "PostTitleAmended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektPurgeAdminTransferredIterator is returned from FilterPurgeAdminTransferred and is used to iterate over the raw logs and unpacked data for PurgeAdminTransferred events raised by the ThatsRekt contract.
type ThatsRektPurgeAdminTransferredIterator struct {
	Event *ThatsRektPurgeAdminTransferred // Event containing the contract specifics and raw log

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
func (it *ThatsRektPurgeAdminTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektPurgeAdminTransferred)
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
		it.Event = new(ThatsRektPurgeAdminTransferred)
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
func (it *ThatsRektPurgeAdminTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektPurgeAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektPurgeAdminTransferred represents a PurgeAdminTransferred event raised by the ThatsRekt contract.
type ThatsRektPurgeAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPurgeAdminTransferred is a free log retrieval operation binding the contract event 0x00c5706759a5b9a68f03d9044a2d6fb4401210a0063438026a8bd5e1c97d3899.
//
// Solidity: event PurgeAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_ThatsRekt *ThatsRektFilterer) FilterPurgeAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*ThatsRektPurgeAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "PurgeAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektPurgeAdminTransferredIterator{contract: _ThatsRekt.contract, event: "PurgeAdminTransferred", logs: logs, sub: sub}, nil
}

// WatchPurgeAdminTransferred is a free log subscription operation binding the contract event 0x00c5706759a5b9a68f03d9044a2d6fb4401210a0063438026a8bd5e1c97d3899.
//
// Solidity: event PurgeAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_ThatsRekt *ThatsRektFilterer) WatchPurgeAdminTransferred(opts *bind.WatchOpts, sink chan<- *ThatsRektPurgeAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "PurgeAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektPurgeAdminTransferred)
				if err := _ThatsRekt.contract.UnpackLog(event, "PurgeAdminTransferred", log); err != nil {
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

// ParsePurgeAdminTransferred is a log parse operation binding the contract event 0x00c5706759a5b9a68f03d9044a2d6fb4401210a0063438026a8bd5e1c97d3899.
//
// Solidity: event PurgeAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_ThatsRekt *ThatsRektFilterer) ParsePurgeAdminTransferred(log types.Log) (*ThatsRektPurgeAdminTransferred, error) {
	event := new(ThatsRektPurgeAdminTransferred)
	if err := _ThatsRekt.contract.UnpackLog(event, "PurgeAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektPurgeRemoverTransferredIterator is returned from FilterPurgeRemoverTransferred and is used to iterate over the raw logs and unpacked data for PurgeRemoverTransferred events raised by the ThatsRekt contract.
type ThatsRektPurgeRemoverTransferredIterator struct {
	Event *ThatsRektPurgeRemoverTransferred // Event containing the contract specifics and raw log

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
func (it *ThatsRektPurgeRemoverTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektPurgeRemoverTransferred)
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
		it.Event = new(ThatsRektPurgeRemoverTransferred)
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
func (it *ThatsRektPurgeRemoverTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektPurgeRemoverTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektPurgeRemoverTransferred represents a PurgeRemoverTransferred event raised by the ThatsRekt contract.
type ThatsRektPurgeRemoverTransferred struct {
	PreviousRemover common.Address
	NewRemover      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPurgeRemoverTransferred is a free log retrieval operation binding the contract event 0x28ce69b5347f5af68f15277e8dc315bf2620588714c7f5b4e33c4454108f7638.
//
// Solidity: event PurgeRemoverTransferred(address indexed previousRemover, address indexed newRemover)
func (_ThatsRekt *ThatsRektFilterer) FilterPurgeRemoverTransferred(opts *bind.FilterOpts, previousRemover []common.Address, newRemover []common.Address) (*ThatsRektPurgeRemoverTransferredIterator, error) {

	var previousRemoverRule []interface{}
	for _, previousRemoverItem := range previousRemover {
		previousRemoverRule = append(previousRemoverRule, previousRemoverItem)
	}
	var newRemoverRule []interface{}
	for _, newRemoverItem := range newRemover {
		newRemoverRule = append(newRemoverRule, newRemoverItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "PurgeRemoverTransferred", previousRemoverRule, newRemoverRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektPurgeRemoverTransferredIterator{contract: _ThatsRekt.contract, event: "PurgeRemoverTransferred", logs: logs, sub: sub}, nil
}

// WatchPurgeRemoverTransferred is a free log subscription operation binding the contract event 0x28ce69b5347f5af68f15277e8dc315bf2620588714c7f5b4e33c4454108f7638.
//
// Solidity: event PurgeRemoverTransferred(address indexed previousRemover, address indexed newRemover)
func (_ThatsRekt *ThatsRektFilterer) WatchPurgeRemoverTransferred(opts *bind.WatchOpts, sink chan<- *ThatsRektPurgeRemoverTransferred, previousRemover []common.Address, newRemover []common.Address) (event.Subscription, error) {

	var previousRemoverRule []interface{}
	for _, previousRemoverItem := range previousRemover {
		previousRemoverRule = append(previousRemoverRule, previousRemoverItem)
	}
	var newRemoverRule []interface{}
	for _, newRemoverItem := range newRemover {
		newRemoverRule = append(newRemoverRule, newRemoverItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "PurgeRemoverTransferred", previousRemoverRule, newRemoverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektPurgeRemoverTransferred)
				if err := _ThatsRekt.contract.UnpackLog(event, "PurgeRemoverTransferred", log); err != nil {
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

// ParsePurgeRemoverTransferred is a log parse operation binding the contract event 0x28ce69b5347f5af68f15277e8dc315bf2620588714c7f5b4e33c4454108f7638.
//
// Solidity: event PurgeRemoverTransferred(address indexed previousRemover, address indexed newRemover)
func (_ThatsRekt *ThatsRektFilterer) ParsePurgeRemoverTransferred(log types.Log) (*ThatsRektPurgeRemoverTransferred, error) {
	event := new(ThatsRektPurgeRemoverTransferred)
	if err := _ThatsRekt.contract.UnpackLog(event, "PurgeRemoverTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ThatsRekt contract.
type ThatsRektUpgradedIterator struct {
	Event *ThatsRektUpgraded // Event containing the contract specifics and raw log

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
func (it *ThatsRektUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektUpgraded)
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
		it.Event = new(ThatsRektUpgraded)
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
func (it *ThatsRektUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektUpgraded represents a Upgraded event raised by the ThatsRekt contract.
type ThatsRektUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ThatsRekt *ThatsRektFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ThatsRektUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektUpgradedIterator{contract: _ThatsRekt.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ThatsRekt *ThatsRektFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ThatsRektUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektUpgraded)
				if err := _ThatsRekt.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ThatsRekt *ThatsRektFilterer) ParseUpgraded(log types.Log) (*ThatsRektUpgraded, error) {
	event := new(ThatsRektUpgraded)
	if err := _ThatsRekt.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektVictimsAddedIterator is returned from FilterVictimsAdded and is used to iterate over the raw logs and unpacked data for VictimsAdded events raised by the ThatsRekt contract.
type ThatsRektVictimsAddedIterator struct {
	Event *ThatsRektVictimsAdded // Event containing the contract specifics and raw log

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
func (it *ThatsRektVictimsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektVictimsAdded)
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
		it.Event = new(ThatsRektVictimsAdded)
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
func (it *ThatsRektVictimsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektVictimsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektVictimsAdded represents a VictimsAdded event raised by the ThatsRekt contract.
type ThatsRektVictimsAdded struct {
	PostId     *big.Int
	Amender    common.Address
	NewVictims []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVictimsAdded is a free log retrieval operation binding the contract event 0x6bb42a267ffcd2d73693fdcf84c1f13c887f2d4dba77e9477c0c4123eae655c8.
//
// Solidity: event VictimsAdded(uint256 indexed postId, address indexed amender, address[] newVictims)
func (_ThatsRekt *ThatsRektFilterer) FilterVictimsAdded(opts *bind.FilterOpts, postId []*big.Int, amender []common.Address) (*ThatsRektVictimsAddedIterator, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var amenderRule []interface{}
	for _, amenderItem := range amender {
		amenderRule = append(amenderRule, amenderItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "VictimsAdded", postIdRule, amenderRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektVictimsAddedIterator{contract: _ThatsRekt.contract, event: "VictimsAdded", logs: logs, sub: sub}, nil
}

// WatchVictimsAdded is a free log subscription operation binding the contract event 0x6bb42a267ffcd2d73693fdcf84c1f13c887f2d4dba77e9477c0c4123eae655c8.
//
// Solidity: event VictimsAdded(uint256 indexed postId, address indexed amender, address[] newVictims)
func (_ThatsRekt *ThatsRektFilterer) WatchVictimsAdded(opts *bind.WatchOpts, sink chan<- *ThatsRektVictimsAdded, postId []*big.Int, amender []common.Address) (event.Subscription, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}
	var amenderRule []interface{}
	for _, amenderItem := range amender {
		amenderRule = append(amenderRule, amenderItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "VictimsAdded", postIdRule, amenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektVictimsAdded)
				if err := _ThatsRekt.contract.UnpackLog(event, "VictimsAdded", log); err != nil {
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

// ParseVictimsAdded is a log parse operation binding the contract event 0x6bb42a267ffcd2d73693fdcf84c1f13c887f2d4dba77e9477c0c4123eae655c8.
//
// Solidity: event VictimsAdded(uint256 indexed postId, address indexed amender, address[] newVictims)
func (_ThatsRekt *ThatsRektFilterer) ParseVictimsAdded(log types.Log) (*ThatsRektVictimsAdded, error) {
	event := new(ThatsRektVictimsAdded)
	if err := _ThatsRekt.contract.UnpackLog(event, "VictimsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektWhitelistAdminTransferredIterator is returned from FilterWhitelistAdminTransferred and is used to iterate over the raw logs and unpacked data for WhitelistAdminTransferred events raised by the ThatsRekt contract.
type ThatsRektWhitelistAdminTransferredIterator struct {
	Event *ThatsRektWhitelistAdminTransferred // Event containing the contract specifics and raw log

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
func (it *ThatsRektWhitelistAdminTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektWhitelistAdminTransferred)
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
		it.Event = new(ThatsRektWhitelistAdminTransferred)
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
func (it *ThatsRektWhitelistAdminTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektWhitelistAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektWhitelistAdminTransferred represents a WhitelistAdminTransferred event raised by the ThatsRekt contract.
type ThatsRektWhitelistAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminTransferred is a free log retrieval operation binding the contract event 0x780008ae49c9be745883f9a2ca6d2e68e9063350aaa0782ee3f00316965d2515.
//
// Solidity: event WhitelistAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_ThatsRekt *ThatsRektFilterer) FilterWhitelistAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*ThatsRektWhitelistAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "WhitelistAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektWhitelistAdminTransferredIterator{contract: _ThatsRekt.contract, event: "WhitelistAdminTransferred", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminTransferred is a free log subscription operation binding the contract event 0x780008ae49c9be745883f9a2ca6d2e68e9063350aaa0782ee3f00316965d2515.
//
// Solidity: event WhitelistAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_ThatsRekt *ThatsRektFilterer) WatchWhitelistAdminTransferred(opts *bind.WatchOpts, sink chan<- *ThatsRektWhitelistAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "WhitelistAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektWhitelistAdminTransferred)
				if err := _ThatsRekt.contract.UnpackLog(event, "WhitelistAdminTransferred", log); err != nil {
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

// ParseWhitelistAdminTransferred is a log parse operation binding the contract event 0x780008ae49c9be745883f9a2ca6d2e68e9063350aaa0782ee3f00316965d2515.
//
// Solidity: event WhitelistAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_ThatsRekt *ThatsRektFilterer) ParseWhitelistAdminTransferred(log types.Log) (*ThatsRektWhitelistAdminTransferred, error) {
	event := new(ThatsRektWhitelistAdminTransferred)
	if err := _ThatsRekt.contract.UnpackLog(event, "WhitelistAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektWhitelistRemoverTransferredIterator is returned from FilterWhitelistRemoverTransferred and is used to iterate over the raw logs and unpacked data for WhitelistRemoverTransferred events raised by the ThatsRekt contract.
type ThatsRektWhitelistRemoverTransferredIterator struct {
	Event *ThatsRektWhitelistRemoverTransferred // Event containing the contract specifics and raw log

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
func (it *ThatsRektWhitelistRemoverTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektWhitelistRemoverTransferred)
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
		it.Event = new(ThatsRektWhitelistRemoverTransferred)
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
func (it *ThatsRektWhitelistRemoverTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektWhitelistRemoverTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektWhitelistRemoverTransferred represents a WhitelistRemoverTransferred event raised by the ThatsRekt contract.
type ThatsRektWhitelistRemoverTransferred struct {
	PreviousRemover common.Address
	NewRemover      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWhitelistRemoverTransferred is a free log retrieval operation binding the contract event 0x1238d49b83f1d9a79330ff8c39c6f65a914331d29985c3f4c1a7ae51e9aae2e8.
//
// Solidity: event WhitelistRemoverTransferred(address indexed previousRemover, address indexed newRemover)
func (_ThatsRekt *ThatsRektFilterer) FilterWhitelistRemoverTransferred(opts *bind.FilterOpts, previousRemover []common.Address, newRemover []common.Address) (*ThatsRektWhitelistRemoverTransferredIterator, error) {

	var previousRemoverRule []interface{}
	for _, previousRemoverItem := range previousRemover {
		previousRemoverRule = append(previousRemoverRule, previousRemoverItem)
	}
	var newRemoverRule []interface{}
	for _, newRemoverItem := range newRemover {
		newRemoverRule = append(newRemoverRule, newRemoverItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "WhitelistRemoverTransferred", previousRemoverRule, newRemoverRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektWhitelistRemoverTransferredIterator{contract: _ThatsRekt.contract, event: "WhitelistRemoverTransferred", logs: logs, sub: sub}, nil
}

// WatchWhitelistRemoverTransferred is a free log subscription operation binding the contract event 0x1238d49b83f1d9a79330ff8c39c6f65a914331d29985c3f4c1a7ae51e9aae2e8.
//
// Solidity: event WhitelistRemoverTransferred(address indexed previousRemover, address indexed newRemover)
func (_ThatsRekt *ThatsRektFilterer) WatchWhitelistRemoverTransferred(opts *bind.WatchOpts, sink chan<- *ThatsRektWhitelistRemoverTransferred, previousRemover []common.Address, newRemover []common.Address) (event.Subscription, error) {

	var previousRemoverRule []interface{}
	for _, previousRemoverItem := range previousRemover {
		previousRemoverRule = append(previousRemoverRule, previousRemoverItem)
	}
	var newRemoverRule []interface{}
	for _, newRemoverItem := range newRemover {
		newRemoverRule = append(newRemoverRule, newRemoverItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "WhitelistRemoverTransferred", previousRemoverRule, newRemoverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektWhitelistRemoverTransferred)
				if err := _ThatsRekt.contract.UnpackLog(event, "WhitelistRemoverTransferred", log); err != nil {
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

// ParseWhitelistRemoverTransferred is a log parse operation binding the contract event 0x1238d49b83f1d9a79330ff8c39c6f65a914331d29985c3f4c1a7ae51e9aae2e8.
//
// Solidity: event WhitelistRemoverTransferred(address indexed previousRemover, address indexed newRemover)
func (_ThatsRekt *ThatsRektFilterer) ParseWhitelistRemoverTransferred(log types.Log) (*ThatsRektWhitelistRemoverTransferred, error) {
	event := new(ThatsRektWhitelistRemoverTransferred)
	if err := _ThatsRekt.contract.UnpackLog(event, "WhitelistRemoverTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThatsRektWhitelistUpdatedIterator is returned from FilterWhitelistUpdated and is used to iterate over the raw logs and unpacked data for WhitelistUpdated events raised by the ThatsRekt contract.
type ThatsRektWhitelistUpdatedIterator struct {
	Event *ThatsRektWhitelistUpdated // Event containing the contract specifics and raw log

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
func (it *ThatsRektWhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThatsRektWhitelistUpdated)
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
		it.Event = new(ThatsRektWhitelistUpdated)
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
func (it *ThatsRektWhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThatsRektWhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThatsRektWhitelistUpdated represents a WhitelistUpdated event raised by the ThatsRekt contract.
type ThatsRektWhitelistUpdated struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistUpdated is a free log retrieval operation binding the contract event 0xf93f9a76c1bf3444d22400a00cb9fe990e6abe9dbb333fda48859cfee864543d.
//
// Solidity: event WhitelistUpdated(address indexed account, bool status)
func (_ThatsRekt *ThatsRektFilterer) FilterWhitelistUpdated(opts *bind.FilterOpts, account []common.Address) (*ThatsRektWhitelistUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ThatsRekt.contract.FilterLogs(opts, "WhitelistUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return &ThatsRektWhitelistUpdatedIterator{contract: _ThatsRekt.contract, event: "WhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchWhitelistUpdated is a free log subscription operation binding the contract event 0xf93f9a76c1bf3444d22400a00cb9fe990e6abe9dbb333fda48859cfee864543d.
//
// Solidity: event WhitelistUpdated(address indexed account, bool status)
func (_ThatsRekt *ThatsRektFilterer) WatchWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *ThatsRektWhitelistUpdated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ThatsRekt.contract.WatchLogs(opts, "WhitelistUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThatsRektWhitelistUpdated)
				if err := _ThatsRekt.contract.UnpackLog(event, "WhitelistUpdated", log); err != nil {
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

// ParseWhitelistUpdated is a log parse operation binding the contract event 0xf93f9a76c1bf3444d22400a00cb9fe990e6abe9dbb333fda48859cfee864543d.
//
// Solidity: event WhitelistUpdated(address indexed account, bool status)
func (_ThatsRekt *ThatsRektFilterer) ParseWhitelistUpdated(log types.Log) (*ThatsRektWhitelistUpdated, error) {
	event := new(ThatsRektWhitelistUpdated)
	if err := _ThatsRekt.contract.UnpackLog(event, "WhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
