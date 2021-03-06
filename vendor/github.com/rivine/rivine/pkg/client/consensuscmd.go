package client

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jimbersoftware/rivine/api"
	"github.com/jimbersoftware/rivine/types"
	"github.com/spf13/cobra"
)

var (
	consensusCmd = &cobra.Command{
		Use:   "consensus",
		Short: "Print the current state of consensus",
		Long:  "Print the current state of consensus such as current block, block height, and target.",
		Run:   Wrap(consensuscmd),
	}

	consensusTransactionCmd = &cobra.Command{
		Use:   "transaction <shortID>",
		Short: "Get an existing transaction",
		Long:  "Get an existing transaction from the blockchain, using its given shortID.",
		Run:   Wrap(consensustransactioncmd),
	}
)

// Consensuscmd is the handler for the command `rivinec consensus`.
// Prints the current state of consensus.
func consensuscmd() {
	var cg api.ConsensusGET
	err := _DefaultClient.httpClient.GetAPI("/consensus", &cg)
	if err != nil {
		Die("Could not get current consensus state:", err)
	}
	if cg.Synced {
		fmt.Printf(`Synced: %v
Block:  %v
Height: %v
Target: %v
`, YesNo(cg.Synced), cg.CurrentBlock, cg.Height, cg.Target)
	} else {
		estimatedHeight := EstimatedHeightAt(time.Now())
		estimatedProgress := float64(cg.Height) / float64(estimatedHeight) * 100
		if estimatedProgress > 99 {
			estimatedProgress = 99
		}
		fmt.Printf(`Synced: %v
Height: %v
Progress (estimated): %.f%%
`, YesNo(cg.Synced), cg.Height, estimatedProgress)
	}
}

// EstimatedHeightAt returns the estimated block height for the given time.
// Block height is estimated by calculating the minutes since a known block in
// the past and dividing by 10 minutes (the block time).
func EstimatedHeightAt(t time.Time) types.BlockHeight {
	block5e4Timestamp := time.Date(2016, time.May, 11, 19, 33, 0, 0, time.UTC)
	diff := t.Sub(block5e4Timestamp)
	estimatedHeight := 5e4 + (diff.Minutes() / 10)
	return types.BlockHeight(estimatedHeight + 0.5) // round to the nearest block
}

// consensustransactioncmd is the handler for the command `rivinec consensus transaction`.
// Prints the transaction found for the given id. If the ID is a long transaction ID, it also
// prints the short transaction ID for future reference
func consensustransactioncmd(id string) {
	var txn api.ConsensusGetTransaction

	err := _DefaultClient.httpClient.GetAPI("/consensus/transactions/"+id, &txn)
	if err != nil {
		Die("failed to get transaction:", err, "; ID:", id)
	}

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(txn)
	if err != nil {
		Die("failed to encode transaction:", err, "; ID:", id)
	}
}
