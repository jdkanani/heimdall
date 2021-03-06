// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/maticnetwork/bor/common"
	heimdalltypes "github.com/maticnetwork/heimdall/types"

	mock "github.com/stretchr/testify/mock"

	rootchain "github.com/maticnetwork/heimdall/contracts/rootchain"

	stakemanager "github.com/maticnetwork/heimdall/contracts/stakemanager"

	statesender "github.com/maticnetwork/heimdall/contracts/statesender"

	types "github.com/maticnetwork/bor/core/types"
)

// IContractCaller is an autogenerated mock type for the IContractCaller type
type IContractCaller struct {
	mock.Mock
}

// CurrentAccountStateRoot provides a mock function with given fields:
func (_m *IContractCaller) CurrentAccountStateRoot() ([32]byte, error) {
	ret := _m.Called()

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func() [32]byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CurrentHeaderBlock provides a mock function with given fields:
func (_m *IContractCaller) CurrentHeaderBlock() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CurrentSpanNumber provides a mock function with given fields:
func (_m *IContractCaller) CurrentSpanNumber() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// CurrentStateCounter provides a mock function with given fields:
func (_m *IContractCaller) CurrentStateCounter() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// DecodeNewHeaderBlockEvent provides a mock function with given fields: _a0, _a1
func (_m *IContractCaller) DecodeNewHeaderBlockEvent(_a0 *types.Receipt, _a1 uint64) (*rootchain.RootchainNewHeaderBlock, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *rootchain.RootchainNewHeaderBlock
	if rf, ok := ret.Get(0).(func(*types.Receipt, uint64) *rootchain.RootchainNewHeaderBlock); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rootchain.RootchainNewHeaderBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeSignerUpdateEvent provides a mock function with given fields: _a0, _a1
func (_m *IContractCaller) DecodeSignerUpdateEvent(_a0 *types.Receipt, _a1 uint64) (*stakemanager.StakemanagerSignerChange, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stakemanager.StakemanagerSignerChange
	if rf, ok := ret.Get(0).(func(*types.Receipt, uint64) *stakemanager.StakemanagerSignerChange); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakemanager.StakemanagerSignerChange)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeValidatorStakeUpdateEvent provides a mock function with given fields: _a0, _a1
func (_m *IContractCaller) DecodeValidatorStakeUpdateEvent(_a0 *types.Receipt, _a1 uint64) (*stakemanager.StakemanagerStakeUpdate, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stakemanager.StakemanagerStakeUpdate
	if rf, ok := ret.Get(0).(func(*types.Receipt, uint64) *stakemanager.StakemanagerStakeUpdate); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakemanager.StakemanagerStakeUpdate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeValidatorTopupFeesEvent provides a mock function with given fields: _a0, _a1
func (_m *IContractCaller) DecodeValidatorTopupFeesEvent(_a0 *types.Receipt, _a1 uint64) (*stakemanager.StakemanagerTopUpFee, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stakemanager.StakemanagerTopUpFee
	if rf, ok := ret.Get(0).(func(*types.Receipt, uint64) *stakemanager.StakemanagerTopUpFee); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakemanager.StakemanagerTopUpFee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncodeStateSyncedEvent provides a mock function with given fields: _a0
func (_m *IContractCaller) EncodeStateSyncedEvent(_a0 *types.Log) (*statesender.StatesenderStateSynced, error) {
	ret := _m.Called(_a0)

	var r0 *statesender.StatesenderStateSynced
	if rf, ok := ret.Get(0).(func(*types.Log) *statesender.StatesenderStateSynced); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*statesender.StatesenderStateSynced)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Log) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: address
func (_m *IContractCaller) GetBalance(address common.Address) (*big.Int, error) {
	ret := _m.Called(address)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address) *big.Int); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockNumberFromTxHash provides a mock function with given fields: _a0
func (_m *IContractCaller) GetBlockNumberFromTxHash(_a0 common.Hash) (*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Hash) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCheckpointSign provides a mock function with given fields: txHash
func (_m *IContractCaller) GetCheckpointSign(txHash common.Hash) ([]byte, []byte, []byte, error) {
	ret := _m.Called(txHash)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(common.Hash) []byte); ok {
		r0 = rf(txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(common.Hash) []byte); ok {
		r1 = rf(txHash)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 []byte
	if rf, ok := ret.Get(2).(func(common.Hash) []byte); ok {
		r2 = rf(txHash)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]byte)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(common.Hash) error); ok {
		r3 = rf(txHash)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetConfirmedTxReceipt provides a mock function with given fields: _a0
func (_m *IContractCaller) GetConfirmedTxReceipt(_a0 common.Hash) (*types.Receipt, error) {
	ret := _m.Called(_a0)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(common.Hash) *types.Receipt); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHeaderInfo provides a mock function with given fields: headerID
func (_m *IContractCaller) GetHeaderInfo(headerID uint64) (common.Hash, uint64, uint64, uint64, heimdalltypes.HeimdallAddress, error) {
	ret := _m.Called(headerID)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(uint64) common.Hash); ok {
		r0 = rf(headerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(uint64) uint64); ok {
		r1 = rf(headerID)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 uint64
	if rf, ok := ret.Get(2).(func(uint64) uint64); ok {
		r2 = rf(headerID)
	} else {
		r2 = ret.Get(2).(uint64)
	}

	var r3 uint64
	if rf, ok := ret.Get(3).(func(uint64) uint64); ok {
		r3 = rf(headerID)
	} else {
		r3 = ret.Get(3).(uint64)
	}

	var r4 heimdalltypes.HeimdallAddress
	if rf, ok := ret.Get(4).(func(uint64) heimdalltypes.HeimdallAddress); ok {
		r4 = rf(headerID)
	} else {
		if ret.Get(4) != nil {
			r4 = ret.Get(4).(heimdalltypes.HeimdallAddress)
		}
	}

	var r5 error
	if rf, ok := ret.Get(5).(func(uint64) error); ok {
		r5 = rf(headerID)
	} else {
		r5 = ret.Error(5)
	}

	return r0, r1, r2, r3, r4, r5
}

// GetLastChildBlock provides a mock function with given fields:
func (_m *IContractCaller) GetLastChildBlock() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMainChainBlock provides a mock function with given fields: _a0
func (_m *IContractCaller) GetMainChainBlock(_a0 *big.Int) (*types.Header, error) {
	ret := _m.Called(_a0)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(*big.Int) *types.Header); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMainTxReceipt provides a mock function with given fields: _a0
func (_m *IContractCaller) GetMainTxReceipt(_a0 common.Hash) (*types.Receipt, error) {
	ret := _m.Called(_a0)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(common.Hash) *types.Receipt); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaticChainBlock provides a mock function with given fields: _a0
func (_m *IContractCaller) GetMaticChainBlock(_a0 *big.Int) (*types.Header, error) {
	ret := _m.Called(_a0)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(*big.Int) *types.Header); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaticTxReceipt provides a mock function with given fields: _a0
func (_m *IContractCaller) GetMaticTxReceipt(_a0 common.Hash) (*types.Receipt, error) {
	ret := _m.Called(_a0)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(common.Hash) *types.Receipt); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSpanDetails provides a mock function with given fields: id
func (_m *IContractCaller) GetSpanDetails(id *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	ret := _m.Called(id)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Int); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *big.Int
	if rf, ok := ret.Get(1).(func(*big.Int) *big.Int); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	var r2 *big.Int
	if rf, ok := ret.Get(2).(func(*big.Int) *big.Int); ok {
		r2 = rf(id)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*big.Int)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(*big.Int) error); ok {
		r3 = rf(id)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetValidatorInfo provides a mock function with given fields: valID
func (_m *IContractCaller) GetValidatorInfo(valID heimdalltypes.ValidatorID) (heimdalltypes.Validator, error) {
	ret := _m.Called(valID)

	var r0 heimdalltypes.Validator
	if rf, ok := ret.Get(0).(func(heimdalltypes.ValidatorID) heimdalltypes.Validator); ok {
		r0 = rf(valID)
	} else {
		r0 = ret.Get(0).(heimdalltypes.Validator)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(heimdalltypes.ValidatorID) error); ok {
		r1 = rf(valID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsTxConfirmed provides a mock function with given fields: _a0
func (_m *IContractCaller) IsTxConfirmed(_a0 common.Hash) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Hash) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SendCheckpoint provides a mock function with given fields: voteSignBytes, sigs, txData
func (_m *IContractCaller) SendCheckpoint(voteSignBytes []byte, sigs []byte, txData []byte) {
	_m.Called(voteSignBytes, sigs, txData)
}
