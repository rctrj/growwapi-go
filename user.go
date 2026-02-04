// https://groww.in/trade-api/docs/curl/user

package growwapi

import (
	"context"
)

// UserProfile represents the user's profile information
//
// https://groww.in/trade-api/docs/curl/user#response-schema
type UserProfile struct {
	// Unique user identifier
	VendorUserId string `json:"vendor_user_id"`
	// Unique Client Code
	Ucc string `json:"ucc"`
	// NSE trading availability
	NseEnabled bool `json:"nse_enabled"`
	// BSE trading availability
	BseEnabled bool `json:"bse_enabled"`
	// Demat debit/pledge instruction status
	DdpiEnabled bool `json:"ddpi_enabled"`
	// Trading segments available (CASH, FNO, COMMODITY)
	ActiveSegments []Segment `json:"active_segments"`
}

// GetUserProfile : This API retrieves the user's profile information including their unique identifiers,
// trading capabilities across exchanges, enabled segments, and DDPI status.
//
// https://groww.in/trade-api/docs/curl/user#get-user-profile
func (c *Client) GetUserProfile(ctx context.Context) (UserProfile, error) {
	const destination = "https://api.groww.in/v1/user/detail"
	return doGetRequest[UserProfile](ctx, c, destination, nil)
}
