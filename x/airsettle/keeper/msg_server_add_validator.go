package keeper

import (
	"context"
	"errors"
	"github.com/ComputerKeeda/airsettle/x/airsettle/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func (k msgServer) AddValidator(goCtx context.Context, msg *types.MsgAddValidator) (*types.MsgAddValidatorResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	newUUID := uuid.New().String()

	exeLayerDetails, found := k.GetExelayerById(ctx, msg.ChainId)

	if !found {
		return &types.MsgAddValidatorResponse{
			VotingPollId: "--",
		}, errors.New("execution layer not found")
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
		return &types.MsgAddValidatorResponse{
			VotingPollId: "--",
		}, errors.New("requester is not a validator")
	}

	store := ctx.KVStore(k.storeKey)
	dynamicPollIdKeyPrefix := types.PollKeyPrefix + msg.ChainId + "/"
	pollStore := prefix.NewStore(store, types.KeyPrefix(dynamicPollIdKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(pollStore, []byte{})
	var polls []types.Poll
	for ; iterator.Valid(); iterator.Next() {
		var poll types.Poll
		k.cdc.MustUnmarshal(iterator.Value(), &poll)
		polls = append(polls, poll)
	}
	iterator.Close()

	for i := 0; i < len(polls); i++ {
		poll := polls[i]
		if poll.NewValidator == msg.NewValidatorAddress && poll.ChainId == msg.ChainId {
			return &types.MsgAddValidatorResponse{
				VotingPollId: poll.PollId,
			}, errors.New("NewValidatorAddress is already in the list at PollId: " + poll.PollId)
		}
	}

	var computedIsComplete bool

	if validatorsLength < 2 {

		computedIsComplete = true

		exeLayerDetails.Validator = append(exeLayerDetails.Validator, msg.NewValidatorAddress)
		exeLayerDetails.VotingPower = append(exeLayerDetails.VotingPower, 100)

		k.UpdateExecutionLayer(ctx, exeLayerDetails)

	} else {
		computedIsComplete = false
	}

	var poll = types.Poll{
		PollId:          newUUID,
		ChainId:         msg.ChainId,
		NewValidator:    msg.NewValidatorAddress,
		VotesDoneBy:     []string{msg.Creator},
		Votes:           []string{"true"},
		TotalValidators: uint64(validatorsLength),
		IsComplete:      computedIsComplete,
		StartDate:       ctx.BlockTime().String(),
		PollCreator:     msg.Creator,
	}

	b := k.cdc.MustMarshal(&poll)
	pollStore.Set([]byte(newUUID), b)

	LogLoop([]string{"UUID created", newUUID})
	LogCreateFileOnPath(newUUID, "test/pollid.test.air")
	return &types.MsgAddValidatorResponse{
		VotingPollId: newUUID,
	}, nil

}
