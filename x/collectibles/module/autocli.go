package collectibles

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"divine/x/collectibles/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "Class",
					Use:            "class [class-id]",
					Short:          "Query class",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}},
				},

				{
					RpcMethod:      "Classes",
					Use:            "classes ",
					Short:          "Query classes",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "Collectible",
					Use:            "collectible [class-id] [token-id]",
					Short:          "Query collectible",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}, {ProtoField: "token_id"}},
				},

				{
					RpcMethod:      "CollectibleByUri",
					Use:            "collectible-by-uri [uri]",
					Short:          "Query collectible-by-uri",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "uri"}},
				},

				{
					RpcMethod:      "CollectiblesByClass",
					Use:            "collectibles-by-class [class-id]",
					Short:          "Query collectibles-by-class",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}},
				},

				{
					RpcMethod:      "CollectiblesByOwner",
					Use:            "collectibles-by-owner [owner]",
					Short:          "Query collectibles-by-owner",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}},
				},

				{
					RpcMethod:      "OwnerOf",
					Use:            "owner-of [class-id] [token-id]",
					Short:          "Query OwnerOf",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}, {ProtoField: "token_id"}},
				},

				{
					RpcMethod:      "Balance",
					Use:            "balance [class-id] [owner]",
					Short:          "Query Balance",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}, {ProtoField: "owner"}},
				},

				{
					RpcMethod:      "TotalSupply",
					Use:            "total-supply [class-id]",
					Short:          "Query TotalSupply",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}},
				},

				{
					RpcMethod:      "CollectiblesUri",
					Use:            "collectibles-uri [class-id] [token-id]",
					Short:          "Query collectibles-URI",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}, {ProtoField: "token_id"}},
				},

				{
					RpcMethod:      "ClassAdmin",
					Use:            "class-admin [class-id]",
					Short:          "Query class-admin",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}},
				},

				{
					RpcMethod:      "Authority",
					Use:            "authority ",
					Short:          "Query authority",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateClass",
					Use:            "create-class [name] [symbol] [uri] [admin]",
					Short:          "Send a create-class tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "symbol"}, {ProtoField: "uri"}, {ProtoField: "admin"}},
				},
				{
					RpcMethod:      "MintCollectible",
					Use:            "mint-collectible [class-id] [uri] [receiver]",
					Short:          "Send a mint-collectible tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}, {ProtoField: "uri"}, {ProtoField: "receiver"}},
				},
				{
					RpcMethod:      "BurnCollectible",
					Use:            "burn-collectible [class-id] [id]",
					Short:          "Send a burn-collectible tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "TransferCollectible",
					Use:            "transfer-collectible [class-id] [id] [receiver]",
					Short:          "Send a transfer-collectible tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}, {ProtoField: "id"}, {ProtoField: "receiver"}},
				},
				{
					RpcMethod:      "UpdateClassAdmin",
					Use:            "update-class-admin [class-id] [new-admin]",
					Short:          "Send a update-class-admin tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "class_id"}, {ProtoField: "new_admin"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
