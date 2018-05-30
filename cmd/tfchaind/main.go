package main

import (
	"fmt"

	"github.com/jimbersoftware/tfchain/pkg/config"

	"github.com/jimbersoftware/rivine/pkg/daemon"
)

var (
	devnet      = "devnet"
	testnet     = "testnet"
	standardnet = "standard"
)

func main() {
	defaultDaemonConfig := daemon.DefaultConfig()
	defaultDaemonConfig.BlockchainInfo = config.GetBlockchainInfo()
	// Default network name, testnet for now since real network is not live yet
	defaultDaemonConfig.NetworkName = standardnet
	defaultDaemonConfig.CreateNetworConfig = SetupNetworks

	daemon.SetupDefaultDaemon(defaultDaemonConfig)
}

// SetupNetworks injects the correct chain constants and genesis nodes based on the chosen network
func SetupNetworks(name string) (daemon.NetworkConfig, error) {
	switch name {
	case standardnet:
		return daemon.NetworkConfig{
			Constants:      config.GetStandardnetGenesis(),
			BootstrapPeers: config.GetStandardnetBootstrapPeers(),
		}, nil
	case testnet:
		return daemon.NetworkConfig{
			Constants:      config.GetTestnetGenesis(),
			BootstrapPeers: config.GetTestnetBootstrapPeers(),
		}, nil
	case devnet:
		return daemon.NetworkConfig{
			Constants:      config.GetDevnetGenesis(),
			BootstrapPeers: nil,
		}, nil

	default:
		return daemon.NetworkConfig{}, fmt.Errorf("Netork name %q not recognized", name)
	}
}
