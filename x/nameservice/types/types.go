package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Initial Starting Price for a address that was never previously owned
// var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("locationtoken", 1)}

// Whois is a struct that contains all the metadata of a name
type Whois struct {
	Value string         `json:"value"` // name of street
	Owner sdk.AccAddress `json:"owner"`
	Lat   float64        `json:"lat"`
	Lon   float64        `json:"lon"`
	Zip   string         `json:"zip"`
}

// Returns a new Whois with the minprice as the price
func NewWhois() Whois {
	return Whois{
		// Lat: MinNamePrice,
	}
}

// implement fmt.Stringer
func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Lat: %s
Lon: %s
Zip: %s`, w.Owner, w.Value, w.Lat, w.Lon, w.Zip))
}
