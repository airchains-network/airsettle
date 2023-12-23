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

func CmdQVerifyBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "q-verify-batch [batch-number] [chain-id] [merkle-root-hash] [prev-merkle-root] [zk-proof]",
		Short: "Query q_verify_batch",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqBatchNumber, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			reqChainId := args[1]
			reqMerkleRootHash := args[2]
			reqPrevMerkleRoot := args[3]
			reqZkProof := args[4]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQVerifyBatchRequest{

				BatchNumber:    reqBatchNumber,
				ChainId:        reqChainId,
				MerkleRootHash: reqMerkleRootHash,
				PrevMerkleRoot: reqPrevMerkleRoot,
				ZkProof:        reqZkProof,
			}

			res, err := queryClient.QVerifyBatch(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
