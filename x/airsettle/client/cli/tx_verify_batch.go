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

func CmdVerifyBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-batch [batch-number] [chain-id] [merkle-root-hash] [prev-merkle-root] [zk-proof]",
		Short: "Broadcast message verify_batch",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBatchNumber, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argChainId := args[1]
			argMerkleRootHash := args[2]
			argPrevMerkleRoot := args[3]
			argZkProof := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVerifyBatch(
				clientCtx.GetFromAddress().String(),
				argBatchNumber,
				argChainId,
				argMerkleRootHash,
				argPrevMerkleRoot,
				argZkProof,
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
