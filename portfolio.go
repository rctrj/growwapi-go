// https://groww.in/trade-api/docs/curl/portfolio

package growwapi

import (
	"context"
	"net/url"
)

// Holding represents a stock holding in the user's DEMAT account
//
// https://groww.in/trade-api/docs/curl/portfolio#response-schema
type Holding struct {
	// ISIN of the instrument
	Isin string `json:"isin"`
	// Trading Symbol of the instrument as defined by the exchange
	TradingSymbol string `json:"trading_symbol"`
	// Quantity of the instrument held
	Quantity int `json:"quantity"`
	// Average price of the holding
	AveragePrice float32 `json:"average_price"`
	// Quantity pledged
	PledgeQuantity float32 `json:"pledge_quantity"`
	// Quantity locked in demat
	DematLockedQuantity float32 `json:"demat_locked_quantity"`
	// Quantity locked by Groww
	GrowwLockedQuantity float32 `json:"groww_locked_quantity"`
	// Quantity re-pledged
	RepledgeQuantity float32 `json:"repledge_quantity"`
	// T+1 quantity (trade day + 1)
	T1Quantity float32 `json:"t1_quantity"`
	// Free quantity in demat
	DematFreeQuantity float32 `json:"demat_free_quantity"`
	// Additional quantity from corporate actions
	CorporateActionAdditionalQuantity int `json:"corporate_action_additional_quantity"`
	// Quantity in active demat transfer
	ActiveDematTransferQuantity int `json:"active_demat_transfer_quantity"`
}

// holdingsResponse represents the holdings API response wrapper
type holdingsResponse struct {
	Holdings []Holding `json:"holdings"`
}

// GetHoldings : This API retrieves current stock holdings stored in a user's DEMAT account.
//
// https://groww.in/trade-api/docs/curl/portfolio#get-holdings
func (c *Client) GetHoldings(ctx context.Context) ([]Holding, error) {
	const destination = "https://api.groww.in/v1/holdings/user"
	resp, err := doGetRequest[holdingsResponse](ctx, c, destination, nil)
	if err != nil {
		return nil, err
	}
	return resp.Holdings, nil
}

// Position represents a position held by the user
//
// https://groww.in/trade-api/docs/curl/portfolio#response-schema-1
type Position struct {
	// Trading Symbol of the instrument as defined by the exchange
	TradingSymbol string `json:"trading_symbol"`
	// Segment of the instrument (CASH, FNO, COMMODITY)
	Segment Segment `json:"segment"`
	// Credit quantity
	CreditQuantity int `json:"credit_quantity"`
	// Credit price
	CreditPrice float32 `json:"credit_price"`
	// Debit quantity
	DebitQuantity int `json:"debit_quantity"`
	// Debit price
	DebitPrice float32 `json:"debit_price"`
	// Carry forward credit quantity
	CarryForwardCreditQuantity int `json:"carry_forward_credit_quantity"`
	// Carry forward credit price
	CarryForwardCreditPrice float32 `json:"carry_forward_credit_price"`
	// Carry forward debit quantity
	CarryForwardDebitQuantity int `json:"carry_forward_debit_quantity"`
	// Carry forward debit price
	CarryForwardDebitPrice float32 `json:"carry_forward_debit_price"`
	// Stock Exchange
	Exchange Exchange `json:"exchange"`
	// ISIN of the instrument
	SymbolIsin string `json:"symbol_isin"`
	// Net quantity
	Quantity int `json:"quantity"`
	// Product type
	Product Product `json:"product"`
	// Net carry forward quantity
	NetCarryForwardQuantity int `json:"net_carry_forward_quantity"`
	// Net price
	NetPrice float32 `json:"net_price"`
	// Net carry forward price
	NetCarryForwardPrice float32 `json:"net_carry_forward_price"`
	// Realised profit and loss
	RealisedPnl float32 `json:"realised_pnl"`
}

// positionsResponse represents the positions API response wrapper
type positionsResponse struct {
	Positions []Position `json:"positions"`
}

// GetPositionsRequest represents the request for Client.GetPositions
//
// https://groww.in/trade-api/docs/curl/portfolio#request-schema-1
type GetPositionsRequest struct {
	// [Optional] Segment of the instrument such as CASH, FNO etc.
	Segment Segment
}

func (g GetPositionsRequest) queryParams() url.Values {
	out := make(url.Values)
	if g.Segment != "" {
		out.Add("segment", string(g.Segment))
	}
	return out
}

// GetPositions : This API retrieves positions representing assets the user holds in their account.
//
// https://groww.in/trade-api/docs/curl/portfolio#get-position-for-user
func (c *Client) GetPositions(ctx context.Context, req GetPositionsRequest) ([]Position, error) {
	const destination = "https://api.groww.in/v1/positions/user"
	resp, err := doGetRequest[positionsResponse](ctx, c, destination, req)
	if err != nil {
		return nil, err
	}
	return resp.Positions, nil
}

// GetPositionBySymbolRequest represents the request for Client.GetPositionBySymbol
//
// https://groww.in/trade-api/docs/curl/portfolio#request-schema-2
type GetPositionBySymbolRequest struct {
	// Trading Symbol of the instrument as defined by the exchange
	TradingSymbol string
	// [Optional] Segment of the instrument such as CASH, FNO etc.
	Segment Segment
}

func (g GetPositionBySymbolRequest) queryParams() url.Values {
	out := make(url.Values)
	out.Add("trading_symbol", g.TradingSymbol)
	if g.Segment != "" {
		out.Add("segment", string(g.Segment))
	}
	return out
}

// GetPositionBySymbol : This API retrieves positions for a specific instrument based on trading symbol.
//
// https://groww.in/trade-api/docs/curl/portfolio#get-position-for-trading-symbol
func (c *Client) GetPositionBySymbol(ctx context.Context, req GetPositionBySymbolRequest) ([]Position, error) {
	const destination = "https://api.groww.in/v1/positions/trading-symbol"
	resp, err := doGetRequest[positionsResponse](ctx, c, destination, req)
	if err != nil {
		return nil, err
	}
	return resp.Positions, nil
}
