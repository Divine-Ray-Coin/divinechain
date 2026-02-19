package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibckeeper "github.com/cosmos/ibc-go/v10/modules/core/keeper"

	"divine/x/collectibles/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema collections.Schema
	Params collections.Item[types.Params]

	Port collections.Item[string]

	ibcKeeperFn func() *ibckeeper.Keeper

	Classes          collections.Map[string, types.Class]
	ClassSequences   collections.Map[string, uint64]
	ClassSupply      collections.Map[string, uint64]
	Collectibles     collections.Map[collections.Pair[string, string], types.Collectible]
	CollectibleByURI collections.Map[string, types.CollectibleIndex]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,
	ibcKeeperFn func() *ibckeeper.Keeper,

) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,

		ibcKeeperFn:      ibcKeeperFn,
		Port:             collections.NewItem(sb, types.PortKey, "port", collections.StringValue),
		Params:           collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		Classes:          collections.NewMap(sb, types.ClassKey, "class", collections.StringKey, codec.CollValue[types.Class](cdc)),
		ClassSequences:   collections.NewMap(sb, types.ClassSequenceKey, "class_sequence", collections.StringKey, collections.Uint64Value),
		ClassSupply:      collections.NewMap(sb, types.ClassSupplyKey, "class_supply", collections.StringKey, collections.Uint64Value),
		Collectibles:     collections.NewMap(sb, types.CollectibleKey, "collectible", collections.PairKeyCodec(collections.StringKey, collections.StringKey), codec.CollValue[types.Collectible](cdc)),
		CollectibleByURI: collections.NewMap(sb, types.CollectibleURIKey, "collectible_by_uri", collections.StringKey, codec.CollValue[types.CollectibleIndex](cdc)),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() []byte {
	return k.authority
}

func (k Keeper) GetParams(ctx sdk.Context) (types.Params, error) {
	return k.Params.Get(ctx)
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	return k.Params.Set(ctx, params)
}

func (k Keeper) SetClass(ctx sdk.Context, class types.Class) error {
	return k.Classes.Set(ctx, class.ClassId, class)
}

func (k Keeper) GetClass(ctx sdk.Context, classID string) (types.Class, error) {
	return k.Classes.Get(ctx, classID)
}

func (k Keeper) SetSupply(ctx sdk.Context, classID string, quantity uint64) error {
	return k.ClassSupply.Set(ctx, classID, quantity)
}

// GetSupply returns the total supply of a class
func (k Keeper) GetSupply(ctx sdk.Context, classID string) (uint64, error) {
	return k.ClassSupply.Get(ctx, classID)
}

func (k Keeper) NextCollectibleID(ctx sdk.Context, classID string) (string, error) {
	seq, err := k.ClassSequences.Get(ctx, classID)
	if err != nil {
		seq = 0
	}

	seq++

	if err := k.ClassSequences.Set(ctx, classID, seq); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", seq), nil
}

func (k Keeper) SetCollectible(ctx sdk.Context, c types.Collectible) error {
	key := collections.Join(c.ClassId, c.Id)

	if err := k.Collectibles.Set(ctx, key, c); err != nil {
		return err
	}

	index := types.CollectibleIndex{
		ClassId:       c.ClassId,
		CollectibleId: c.Id,
	}
	// Update URI index
	return k.CollectibleByURI.Set(ctx, c.Uri, index)
}

func (k Keeper) GetCollectible(ctx sdk.Context, classID, id string) (types.Collectible, error) {
	return k.Collectibles.Get(ctx, collections.Join(classID, id))
}

func (k Keeper) RemoveCollectible(ctx sdk.Context, classID, id string) error {
	key := collections.Join(classID, id)

	c, err := k.GetCollectible(ctx, classID, id)
	if err != nil {
		return err
	}

	if err := k.Collectibles.Remove(ctx, key); err != nil {
		return err
	}

	return k.CollectibleByURI.Remove(ctx, c.Uri)
}

func (k Keeper) GetCollectibleByURI(ctx sdk.Context, uri string) (types.Collectible, error) {
	key, err := k.CollectibleByURI.Get(ctx, uri)
	if err != nil {
		return types.Collectible{}, err
	}

	return k.GetCollectible(ctx, key.ClassId, key.CollectibleId)
}

func (k Keeper) IBCKeeper() *ibckeeper.Keeper {
	return k.ibcKeeperFn()
}
