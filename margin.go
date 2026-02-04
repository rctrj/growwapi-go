// https://groww.in/trade-api/docs/curl/margin

package growwapi

import (
	"context"
	"net/url"
)

// FnoMarginDetails represents the margin details for Futures and Options segment
type FnoMarginDetails struct {
	// SPAN margin requirement
	SpanMargin float32 `json:"span_margin"`
	// Exposure margin requirement
	ExposureMargin float32 `json:"exposure_margin"`
	// Available balance
	AvailableBalance float32 `json:"available_balance"`
	// Used margin
	UsedMargin float32 `json:"used_margin"`
}

// EquityMarginDetails represents the margin details for Equity segment
type EquityMarginDetails struct {
	// CNC available balance
	CncAvailableBalance float32 `json:"cnc_available_balance"`
	// CNC used margin
	CncUsedMargin float32 `json:"cnc_used_margin"`
	// MIS available balance
	MisAvailableBalance float32 `json:"mis_available_balance"`
	// MIS used margin
	MisUsedMargin float32 `json:"mis_used_margin"`
}

// CommodityMarginDetails represents the margin details for Commodity segment
type CommodityMarginDetails struct {
	// SPAN margin requirement
	SpanMargin float32 `json:"span_margin"`
	// Exposure margin requirement
	ExposureMargin float32 `json:"exposure_margin"`
	// Available balance
	AvailableBalance float32 `json:"available_balance"`
	// Used margin
	UsedMargin float32 `json:"used_margin"`
	// Mark to market value
	M2mValue float32 `json:"m2m_value"`
}

// UserMargin represents the margin details for a user
//
// https://groww.in/trade-api/docs/curl/margin#response-schema
type UserMargin struct {
	// Available liquid funds
	ClearCash float32 `json:"clear_cash"`
	// Total margin utilized
	NetMarginUsed float32 `json:"net_margin_used"`
	// Associated brokerage and charges
	BrokerageAndCharges float32 `json:"brokerage_and_charges"`
	// Pledged securities value used
	CollateralUsed float32 `json:"collateral_used"`
	// Unused collateral value
	CollateralAvailable float32 `json:"collateral_available"`
	// Additional margin granted
	AdhocMargin float32 `json:"adhoc_margin"`
	// Futures & Options margin breakdown
	FnoMarginDetails FnoMarginDetails `json:"fno_margin_details"`
	// Stock market margin breakdown
	EquityMarginDetails EquityMarginDetails `json:"equity_margin_details"`
	// Commodity-specific margins
	CommodityMarginDetails CommodityMarginDetails `json:"commodity_margin_details"`
}

// GetUserMargin : This API retrieves the margin details for a user.
//
// https://groww.in/trade-api/docs/curl/margin#get-available-user-margin
func (c *Client) GetUserMargin(ctx context.Context) (UserMargin, error) {
	const destination = "https://api.groww.in/v1/margins/detail/user"
	return doGetRequest[UserMargin](ctx, c, destination, nil)
}

// OrderMarginRequest represents a single order for margin calculation
//
// https://groww.in/trade-api/docs/curl/margin#request-schema-1
type OrderMarginRequest struct {
	// Trading Symbol of the instrument as defined by the exchange
	TradingSymbol string `json:"trading_symbol"`
	// Transaction type - BUY or SELL
	TransactionType TransactionType `json:"transaction_type"`
	// Order quantity
	Quantity int `json:"quantity"`
	// Price for limit orders
	Price float32 `json:"price,omitempty"`
	// Order type - LIMIT or MARKET
	OrderType OrderType `json:"order_type"`
	// Product type - CNC, MIS, NRML
	Product Product `json:"product"`
	// Stock Exchange
	Exchange Exchange `json:"exchange"`
}

// RequiredMarginRequest represents the request for Client.GetRequiredMargin
type RequiredMarginRequest struct {
	// Segment of the instrument such as CASH, FNO etc.
	Segment Segment
	// Orders to calculate margin for (can be a basket of orders for FNO/COMMODITY)
	Orders []OrderMarginRequest
}

func (r RequiredMarginRequest) queryParams() url.Values {
	out := make(url.Values)
	out.Add("segment", string(r.Segment))
	return out
}

// RequiredMargin represents the margin requirement for an order or basket of orders
//
// https://groww.in/trade-api/docs/curl/margin#response-schema-1
type RequiredMargin struct {
	// Exposure margin needed
	ExposureRequired float32 `json:"exposure_required"`
	// SPAN margin requirement
	SpanRequired float32 `json:"span_required"`
	// Premium for options
	OptionBuyPremium float32 `json:"option_buy_premium"`
	// Transaction costs
	BrokerageAndCharges float32 `json:"brokerage_and_charges"`
	// Combined margin obligation
	TotalRequirement float32 `json:"total_requirement"`
	// CNC-specific margin requirement
	CashCncMarginRequired float32 `json:"cash_cnc_margin_required"`
	// MIS-specific margin requirement
	CashMisMarginRequired float32 `json:"cash_mis_margin_required"`
	// Delivery-related margin requirement
	PhysicalDeliveryMarginRequirement float32 `json:"physical_delivery_margin_requirement"`
}

// GetRequiredMargin : This API calculates margin requirements for single orders or basket of orders.
// Basket margin calculation is supported for FNO and COMMODITY segments.
//
// https://groww.in/trade-api/docs/curl/margin#required-margin-for-order
func (c *Client) GetRequiredMargin(ctx context.Context, req RequiredMarginRequest) (RequiredMargin, error) {
	destination := "https://api.groww.in/v1/margins/detail/orders"
	params := req.queryParams()

	fullUrl := destination + "?" + params.Encode()
	return doPostRequest[RequiredMargin](ctx, c, fullUrl, req.Orders)
}
