package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"point/x/lockenomics/types"
)

var (
	_ stakingTypes.StakingHooks = Hooks{}
)

// Hooks wrapper struct for the lockenomics keeper
type Hooks struct {
	k Keeper
}

// Return the wrapper struct
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (h Hooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) error {
	//TODO implement me
	panic("implement me")
}

func (h Hooks) BeforeValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) error {
	//TODO implement me
	panic("implement me")
}

func (h Hooks) AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	//TODO implement me
	panic("implement me")
}

func (h Hooks) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	//TODO implement me
	panic("implement me")
}

func (h Hooks) AfterValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	//TODO implement me
	panic("implement me")
}

func (h Hooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	//TODO implement me
	panic("implement me")
}

func (h Hooks) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	//TODO implement me
	panic("implement me")
}

func (h Hooks) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	//TODO implement me
	panic("implement me")
}

func (h Hooks) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) error {
	//TODO implement me
	panic("implement me")
}

// AfterDelegationModified is a wrapper for calling the Staking AfterDelegationModified
// hook on the module keeper
func (h Hooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	delegation, found := h.k.stakingKeeper.GetDelegation(ctx, delAddr, valAddr)
	if !found {
		return stakingTypes.ErrNoDelegation
	}

	if !delegation.Shares.IsPositive() {
		return stakingTypes.ErrNoDelegation
	}

	validator, found := h.k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return stakingTypes.ErrNoValidatorFound
	}

	tokens := validator.TokensFromShares(delegation.Shares)

	delegatedAmount := types.DelegatedAmount{
		Delegator: delAddr.String(),
		Validator: valAddr.String(),
		Amount:    uint64(tokens.RoundInt64()),
	}

	h.k.SetDelegatedAmount(ctx, delegatedAmount)

	return nil
}
