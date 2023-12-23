package cli

import (
	"strconv"

	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-batch [batch-number] [chain-id]",
		Short: "Query get_batch",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqBatchNumber, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			reqChainId := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetBatchRequest{

				BatchNumber: reqBatchNumber,
				ChainId:     reqChainId,
			}

			res, err := queryClient.GetBatch(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
