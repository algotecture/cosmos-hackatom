package nameservice

import (
	"github.com/cosmos/sdk-application-tutorial/x/nameservice/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgBuyLocation = types.NewMsgBuyLocation
	NewMsgSetDAG      = types.NewMsgSetDAG
	NewWhois          = types.NewWhois
	ModuleCdc         = types.ModuleCdc
	RegisterCodec     = types.RegisterCodec
)

type (
	MsgSetDAG       = types.MsgSetDAG
	MsgBuyLocation  = types.MsgBuyLocation
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResLocations
	Whois           = types.Whois
)
