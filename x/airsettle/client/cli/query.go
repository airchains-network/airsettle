package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/airchains-network/airsettle/x/airsettle/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group airsettle queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdVerificationKey())

	cmd.AddCommand(CmdShowExecutionLayerByAddress())

	cmd.AddCommand(CmdShowExecutionLayerById())

	cmd.AddCommand(CmdListAllExecutionLayers())

	cmd.AddCommand(CmdGetBatch())

	cmd.AddCommand(CmdListPolls())

	cmd.AddCommand(CmdGetPoll())

	cmd.AddCommand(CmdQVerifyBatch())

	// this line is used by starport scaffolding # 1

	return cmd
}
