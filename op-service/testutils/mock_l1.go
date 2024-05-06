package testutils

import (
	"context"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
)

type MockL1Source struct {
	MockEthClient
}

func (m *MockL1Source) L1BlockRefByLabel(ctx context.Context, label eth.BlockLabel) (eth.L1BlockRef, error) {
	out := m.Mock.Called(label)
	return out.Get(0).(eth.L1BlockRef), out.Error(1)
}

func (m *MockL1Source) ExpectL1BlockRefByLabel(label eth.BlockLabel, ref eth.L1BlockRef, err error) {
	m.Mock.On("L1BlockRefByLabel", label).Once().Return(ref, err)
}

func (m *MockL1Source) L1BlockRefByNumber(ctx context.Context, num uint64) (eth.L1BlockRef, error) {
	out := m.Mock.Called(num)
	return out.Get(0).(eth.L1BlockRef), out.Error(1)
}

func (m *MockL1Source) ExpectL1BlockRefByNumber(num uint64, ref eth.L1BlockRef, err error) {
	m.Mock.On("L1BlockRefByNumber", num).Once().Return(ref, err)
}

func (m *MockL1Source) L1BlockRefByHash(ctx context.Context, hash common.Hash) (eth.L1BlockRef, error) {
	out := m.Mock.Called(hash)
	return out.Get(0).(eth.L1BlockRef), out.Error(1)
}

func (m *MockL1Source) ExpectL1BlockRefByHash(hash common.Hash, ref eth.L1BlockRef, err error) {
	m.Mock.On("L1BlockRefByHash", hash).Once().Return(ref, err)
}

func (m *MockL1Source) GoOrUpdatePreFetchReceipts(ctx context.Context, l1StartBlock uint64) error {
	out := m.Mock.MethodCalled("GoOrUpdatePreFetchReceipts", ctx, l1StartBlock)
	return *out[0].(*error)
}

func (m *MockL1Source) ExpectGoOrUpdatePreFetchReceipts(background context.Context, number uint64, err error) {
	m.Mock.On("GoOrUpdatePreFetchReceipts", background, number).Once().Return(&err)
}

func (m *MockL1Source) ClearReceiptsCacheBefore(blockNumber uint64) {
	m.Mock.MethodCalled("ClearReceiptsCacheBefore", blockNumber)
}

func (m *MockL1Source) ExpectClearReceiptsCacheBefore(blockNumber uint64) {
	m.Mock.On("ClearReceiptsCacheBefore", blockNumber).Once()
}