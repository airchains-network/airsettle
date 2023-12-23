package cli

import (
	"strconv"

	"github.com/ComputerKeeda/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-batch [batch-number] [chain-id] [inputs]",
		Short: "Broadcast message add_batch",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBatchNumber, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argChainId := args[1]
			argInputs := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddBatch(
				clientCtx.GetFromAddress().String(),
				argBatchNumber,
				argChainId,
				argInputs,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
