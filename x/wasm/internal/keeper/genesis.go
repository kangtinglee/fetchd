package keeper

import (
	"github.com/CosmWasm/wasmd/x/wasm/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	// "github.com/CosmWasm/fetchd/x/wasm/internal/types"
)

// InitGenesis sets supply information for genesis.
//
// CONTRACT: all types of accounts must have been already initialized/created
func InitGenesis(ctx sdk.Context, keeper Keeper, data types.GenesisState) error {
	var maxCodeID uint64
	for i, code := range data.Codes {
		err := keeper.importCode(ctx, code.CodeID, code.CodeInfo, code.CodesBytes)
		if err != nil {
			return sdkerrors.Wrapf(err, "code %d with id: %d", i, code.CodeID)
		}
		if code.CodeID > maxCodeID {
			maxCodeID = code.CodeID
		}
	}

	var maxContractID int
	for i, contract := range data.Contracts {
		err := keeper.importContract(ctx, contract.ContractAddress, &contract.ContractInfo, contract.ContractState)
		if err != nil {
			return sdkerrors.Wrapf(err, "contract number %d", i)
		}
		maxContractID = i + 1 // not ideal but max(contractID) is not persisted otherwise
	}

	for i, seq := range data.Sequences {
		err := keeper.importAutoIncrementID(ctx, seq.IDKey, seq.Value)
		if err != nil {
			return sdkerrors.Wrapf(err, "sequence number %d", i)
		}
	}

	// sanity check seq values
	if keeper.peekAutoIncrementID(ctx, types.KeyLastCodeID) <= maxCodeID {
		return sdkerrors.Wrapf(types.ErrInvalid, "seq %s must be greater %d ", string(types.KeyLastCodeID), maxCodeID)
	}
	if keeper.peekAutoIncrementID(ctx, types.KeyLastInstanceID) <= uint64(maxContractID) {
		return sdkerrors.Wrapf(types.ErrInvalid, "seq %s must be greater %d ", string(types.KeyLastInstanceID), maxContractID)
	}

	return nil
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) types.GenesisState {
	var genState types.GenesisState

	maxCodeID := keeper.GetNextCodeID(ctx)
	for i := uint64(1); i < maxCodeID; i++ {
		if !keeper.containsCodeInfo(ctx, i) {
			continue
		}
		bytecode, err := keeper.GetByteCode(ctx, i)
		if err != nil {
			panic(err)
		}
		genState.Codes = append(genState.Codes, types.Code{
			CodeID:     i,
			CodeInfo:   *keeper.GetCodeInfo(ctx, i),
			CodesBytes: bytecode,
		})
	}

	keeper.ListContractInfo(ctx, func(addr sdk.AccAddress, contract types.ContractInfo) bool {
		contractStateIterator := keeper.GetContractState(ctx, addr)
		var state []types.Model
		for ; contractStateIterator.Valid(); contractStateIterator.Next() {
			m := types.Model{
				Key:   contractStateIterator.Key(),
				Value: contractStateIterator.Value(),
			}
			state = append(state, m)
		}

		genState.Contracts = append(genState.Contracts, types.Contract{
			ContractAddress: addr,
			ContractInfo:    contract,
			ContractState:   state,
		})

		return false
	})

	for _, k := range [][]byte{types.KeyLastCodeID, types.KeyLastInstanceID} {
		genState.Sequences = append(genState.Sequences, types.Sequence{
			IDKey: k,
			Value: keeper.peekAutoIncrementID(ctx, k),
		})
	}

	return genState
}
