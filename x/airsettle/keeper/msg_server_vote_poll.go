package keeper

import (
	"context"
	"errors"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

func (k msgServer) VotePoll(goCtx context.Context, msg *types.MsgVotePoll) (*types.MsgVotePollResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	pollDetails, error := k.GetPollHelper(ctx, msg.ChainId, msg.PollId)
	if error != nil {
		return nil, error
	}

	var votesDoneByLength = len(pollDetails.VotesDoneBy)
	for i := 0; i < votesDoneByLength; i++ {
		votedBy := pollDetails.VotesDoneBy[i]
		if votedBy == msg.Creator {
			return &types.MsgVotePollResponse{}, errors.New("already voted")
		}
	}

	exeLayerDetails, found := k.GetExelayerById(ctx, pollDetails.ChainId)
	if !found {
		return &types.MsgVotePollResponse{}, errors.New("execution Layer not found")
	}
	var validatorsLength = len(exeLayerDetails.Validator)

	var isAuthenticValidator = false
	for i := 0; i < validatorsLength; i++ {
		validatorAddress := exeLayerDetails.Validator[i]
		if validatorAddress == msg.Creator {
			isAuthenticValidator = true
			break
		}
	}

	if !isAuthenticValidator {
		return &types.MsgVotePollResponse{}, errors.New("requester is not a validator")
	}

	var newVotesDoneBy = pollDetails.VotesDoneBy
	newVotesDoneBy = append(newVotesDoneBy, msg.Creator)

	var newVotes = pollDetails.Votes
	myStringVoteValue := strconv.FormatBool(msg.Vote)
	newVotes = append(newVotes, myStringVoteValue)

	var newIsComplete bool
	if len(newVotesDoneBy) >= int(pollDetails.TotalValidators) {

		newIsComplete = true

		var successVotePercentage float64
		var trueVoteCount int = 0
		for i := 0; i < len(pollDetails.Votes); i++ {
			value := pollDetails.Votes[i]
			if value == "true" {
				trueVoteCount++
			}
		}

		successVotePercentage = float64(trueVoteCount) / float64(pollDetails.TotalValidators) * 100

		if successVotePercentage >= 50.0 {
			var newValidators = exeLayerDetails.Validator
			var newVotingPower = exeLayerDetails.VotingPower
			newValidators = append(newValidators, pollDetails.NewValidator)
			newVotingPower = append(newVotingPower, 100)

			k.UpdateExecutionLayer(ctx, types.Exelayer{
				Validator:            newValidators,
				VotingPower:          newVotingPower,
				LatestBatch:          exeLayerDetails.LatestBatch,
				LatestMerkleRootHash: exeLayerDetails.LatestMerkleRootHash,
				VerificationKey:      exeLayerDetails.VerificationKey,
				ChainInfo:            exeLayerDetails.ChainInfo,
				Id:                   exeLayerDetails.Id,
				Creator:              exeLayerDetails.Creator,
			})
		}
	} else {
		newIsComplete = false
	}

	store := ctx.KVStore(k.storeKey)
	pollStore := prefix.NewStore(store, types.KeyPrefix(types.PollKeyPrefix+msg.ChainId+"/"))
	var poll = types.Poll{
		PollId:          pollDetails.PollId,
		ChainId:         pollDetails.ChainId,
		NewValidator:    pollDetails.NewValidator,
		VotesDoneBy:     newVotesDoneBy,
		Votes:           newVotes,
		TotalValidators: pollDetails.TotalValidators,
		IsComplete:      newIsComplete,
		StartDate:       pollDetails.StartDate,
		PollCreator:     pollDetails.PollCreator,
	}
	b := k.cdc.MustMarshal(&poll)
	pollStore.Set([]byte(pollDetails.PollId), b)

	return &types.MsgVotePollResponse{}, nil
}
