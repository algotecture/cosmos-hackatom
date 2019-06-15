package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// TODO MsgSetDAG defines a message
type MsgSetDAG struct {
	Name  string         `json:"name"`
	Value string         `json:"value"` //
	Owner sdk.AccAddress `json:"owner"` // should be only owner of last append.. but let's skip it
}

// TODO !!
// NewMsgSetDAG is a constructor function for MsgSetDAG
func NewMsgSetDAG(name string, value string, owner sdk.AccAddress) MsgSetDAG {
	return MsgSetDAG{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgSetDAG) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetDAG) Type() string { return "set_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetDAG) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetDAG) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetDAG) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgBuyLocation defines the MsgBuyLocation message
type MsgBuyLocation struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
	Lat   string         `json:"lat"`
	Lon   string         `json:"lon"`
}

// NewMsgBuyLocation is the constructor function for MsgBuyLocation
func NewMsgBuyLocation(name string, bid sdk.Coins, buyer sdk.AccAddress, lat string, lon string) MsgBuyLocation {
	return MsgBuyLocation{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
		Lat:   lat,
		Lon:   lon,
	}
}

// Route should return the name of the module
func (msg MsgBuyLocation) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBuyLocation) Type() string { return "buy_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyLocation) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if len(msg.Lat) == 0 {
		return sdk.ErrUnknownRequest("Lat cannot be empty")
	}
	if len(msg.Lon) == 0 {
		return sdk.ErrUnknownRequest("Lon cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBuyLocation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyLocation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}
