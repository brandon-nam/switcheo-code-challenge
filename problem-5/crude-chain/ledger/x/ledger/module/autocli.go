package ledger

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "ledger/api/ledger/ledger"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowLedger",
					Use:            "show-ledger [id]",
					Short:          "Query show-ledger",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListLedger",
					Use:            "list-ledger",
					Short:          "Query list-ledger",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateLedger",
					Use:            "create-ledger [title] [body] [cost]",
					Short:          "Send a create-ledger tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "body"}, {ProtoField: "cost"}},
				},
				{
					RpcMethod:      "CreateLedger",
					Use:            "create-ledger [title] [body] [cost]",
					Short:          "Send a create-ledger tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "body"}, {ProtoField: "cost"}},
				},
				{
					RpcMethod:      "CreateLedger",
					Use:            "create-ledger [title] [body] [cost]",
					Short:          "Send a create-ledger tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "body"}, {ProtoField: "cost"}},
				},
				{
					RpcMethod:      "UpdateLedger",
					Use:            "update-ledger [title] [body] [cost] [id]",
					Short:          "Send a update-ledger tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "body"}, {ProtoField: "cost"}, {ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
