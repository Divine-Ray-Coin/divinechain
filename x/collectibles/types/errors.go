package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/collectibles module sentinel errors
var (
	ErrInvalidSigner        = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidPacketTimeout = errors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = errors.Register(ModuleName, 1501, "invalid version")
	ErrInvalidRequest       = errors.Register(ModuleName, 1502, "invalid request")
	ErrUnauthorized         = errors.Register(ModuleName, 1503, "unauthorized")
	ErrLocked               = errors.Register(ModuleName, 2, "collectible locked")
	ErrNotFound             = errors.Register(ModuleName, 3, "collectible not found")
	ErrNotOwner             = errors.Register(ModuleName, 4, "not collectible owner")
	ErrMintDisabled         = errors.Register(ModuleName, 5, "minting disabled")
	ErrBurnDisabled         = errors.Register(ModuleName, 6, "burning disabled")
	ErrClassExists          = errors.Register(ModuleName, 7, "collectible class already exists")
	ErrClassNotFound        = errors.Register(ModuleName, 8, "collectible class not found")
	ErrURIAlreadyExists     = errors.Register(ModuleName, 9, "uri already exists")
	ErrCollectibleNotFound  = errors.Register(ModuleName, 10, "collectible not found")
	ErrCollectibleLocked    = errors.Register(ModuleName, 11, "collectible locked")
	ErrChannelNotFound      = errors.Register(ModuleName, 12, "channel not found")
	ErrSequenceNotFound     = errors.Register(ModuleName, 13, "sequence not found")
)
