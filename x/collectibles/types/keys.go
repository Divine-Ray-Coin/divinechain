package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "collectibles"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// GovModuleName duplicates the gov module's name to avoid a dependency with x/gov.
	// It should be synced with the gov module's name if it is ever changed.
	// See: https://github.com/cosmos/cosmos-sdk/blob/v0.52.0-beta.2/x/gov/types/keys.go#L9
	GovModuleName = "gov"

	// Version defines the current version the IBC module supports
	Version = "collectibles-1"

	// PortID is the default port id that module binds to
	PortID = "collectibles"

	EventTypeCreateClass            = "create_class"
	EventTypeUpdateClassAdmin       = "update_class"
	EventTypeMintCollectible        = "mint_collectible"
	EventTypeTransferCollectible    = "transfer_collectible"
	EventTypeBurnCollectible        = "burn_collectible"
	EventTypeIBCTransferCollectible = "ibc_transfer_collectible"

	AttributeKeyClassID       = "class_id"
	AttributeKeyClassName     = "class_name"
	AttributeKeyCollectibleID = "collectible_id"
	AttributeKeyClassAdmin    = "class_admin"
	AttributeKeyOwner         = "owner"
	AttributeKeyReceiver      = "receiver"
	AttributeKeyURI           = "uri"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = collections.NewPrefix("collectibles-port-")

	// ParamsKey is the prefix to retrieve all Params
	ParamsKey = collections.NewPrefix("p_collectibles")

	ClassKey = collections.NewPrefix("class/")

	// CollectibleKey stores collectibles by (classID, collectibleID)
	CollectibleKey = collections.NewPrefix("collectible/")

	// CollectibleURIKey maps URI → (classID, collectibleID)
	CollectibleURIKey = collections.NewPrefix("collectible-uri/")

	// ClassSequenceKey stores per-class auto increment counters
	ClassSequenceKey = collections.NewPrefix("class-seq/")

	ClassSupplyKey = collections.NewPrefix("class-supply/")
)
