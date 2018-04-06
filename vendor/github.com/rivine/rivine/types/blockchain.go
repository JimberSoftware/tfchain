package types

import (
	"github.com/rivine/rivine/build"
)

// BlockchainInfo contains information about a blockchain.
type BlockchainInfo struct {
	Name            string
	ChainVersion    build.ProtocolVersion
	ProtocolVersion build.ProtocolVersion
}

// DefaultBlockchainInfo returns the blockchain information
// for the default (Rivine) blockchain, using the version
// which is set as part of the build process.
func DefaultBlockchainInfo() BlockchainInfo {
	return BlockchainInfo{
		Name:            "Rivine",
		ChainVersion:    build.Version,
		ProtocolVersion: build.Version,
	}
}
