package daemon

import (
	"github.com/jimbersoftware/rivine/modules"
	"github.com/jimbersoftware/rivine/types"
)

// NetworkConfig are variables for a particular chain. Currently, these are genesis constants and bootstrap peers
type NetworkConfig struct {
	// Blockchain Constants for this network
	Constants types.ChainConstants
	// BootstrapPeers for this network
	BootstrapPeers []modules.NetAddress
}
