package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ExecutionlayersKeyPrefix is the prefix to retrieve all Executionlayers
	ExecutionLayerKey = "Exelayer/id/"
	VerificationKey   = "Exelayer/vk/"
	CounterStoreKey   = "Exelayer/counter/"
	ChainAdminKey     = "ExelayerChain/admin/"

	ExeLayerChainKey = "ExelayerChain/value"
	PollKeyPrefix    = "Poll/value/"
)

// ExecutionlayersKey returns the store key to retrieve a Executionlayers from the index fields
func ExeLayerKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
