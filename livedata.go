// https://groww.in/trade-api/docs/curl/live-data

package growwapi

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// QuoteRequest represents the request for Client.GetQuote
// https://groww.in/trade-api/docs/curl/live-data#request-schema
type QuoteRequest struct {
	// Stock Exchange
	Exchange Exchange `json:"exchange"`
	// Segment of the instrument such as CASH, FNO etc.
	Segment Segment `json:"segment"`
	// Trading Symbol of the instrument as defined by the exchange
	TradingSymbol string `json:"trading_symbol"`
}

type BookEntry struct {
	// Price of the book entry
	Price float32 `json:"price"`
	// Quantity of the book entry
	Quantity int `json:"quantity"`
}

type Quote struct {
	// Average price of the instrument in Rupees
	AveragePrice float32 `json:"average_price"`
	// Quantity of the bid
	BidQuantity int `json:"bid_quantity"`
	// Price of the bid
	BidPrice float32 `json:"bid_price"`
	// Day change in price
	DayChange float32 `json:"day_change"`
	// Day change percentage
	DayChangePerc float32 `json:"day_change_perc"`
	// High price range
	UpperCircuitLimit float32 `json:"upper_circuit_limit"`
	// Low price range
	LowerCircuitLimit float32 `json:"lower_circuit_limit"`
	// Ohlc for the instrument
	Ohlc ohlcString `json:"ohlc"`
	// Depth for the instrument
	Depth struct {
		// Buy book entries
		Buy []BookEntry `json:"buy"`
		// Sell book entries
		Sell []BookEntry `json:"sell"`
	} `json:"depth"`
	// High trade range
	HighTradeRange float32 `json:"high_trade_range"`
	// Implied volatility
	ImpliedVolatility float32 `json:"implied_volatility"`
	// Last trade quantity
	LastTradeQuantity int `json:"last_trade_quantity"`
	// Last trade time
	LastTradeTime Time `json:"last_trade_time"`
	// Low trade range
	LowTradeRange float32 `json:"low_trade_range"`
	// Last traded price
	LastPrice float32 `json:"last_price"`
	// Market capitalization
	MarketCap float32 `json:"market_cap"`
	// Offer price
	OfferPrice float32 `json:"offer_price"`
	// Quantity of the offer
	OfferQuantity int `json:"offer_quantity"`
	// Open interest day change
	OiDayChange float32 `json:"oi_day_change"`
	// Open interest day change percentage
	OiDayChangePercentage float32 `json:"oi_day_change_percentage"`
	// Open interest
	OpenInterest float32 `json:"open_interest"`
	// Previous open interest
	PreviousOpenInterest float32 `json:"previous_open_interest"`
	// Total buy quantity
	TotalBuyQuantity float32 `json:"total_buy_quantity"`
	// Total sell quantity
	TotalSellQuantity float32 `json:"total_sell_quantity"`
	// Volume of trades
	Volume int `json:"volume"`
	// 52-week high price
	Week52High float32 `json:"week_52_high"`
	// 52-week low price
	Week52Low float32 `json:"week_52_low"`
}

func (q QuoteRequest) queryParams() url.Values {
	out := make(url.Values)
	out.Add("exchange", string(q.Exchange))
	out.Add("segment", string(q.Segment))
	out.Add("trading_symbol", q.TradingSymbol)
	return out
}

// GetQuote : This API provides the complete live data snapshot for an instrument including the latest price, market depth, ohlc, market volumes and much more.
// If one requires only the latest price data then the Client.GetLtp api should be used.
// Similarly, if one is interested in getting only ohlc then Client.GetOhlc api should be used.
// Use the segment value FNO for derivatives and CASH for stocks and index.
// https://groww.in/trade-api/docs/curl/live-data#get-quote
func (c *Client) GetQuote(ctx context.Context, req QuoteRequest) (Quote, error) {
	const destination = "https://api.groww.in/v1/live-data/quote"
	return doGetRequest[Quote](ctx, c, destination, req)
}

// LtpRequest represents request for Client.GetLtp
// https://groww.in/trade-api/docs/curl/live-data#request-schema-1
type LtpRequest struct {
	// Segment of the instrument such as CASH, FNO etc.
	Segment Segment `json:"segment"`
	// Array of trading symbols with their respective exchanges.
	// For example: `NSE_RELIANCE` `BSE_SENSEX` `NSE_NIFTY25APR24100PE`
	ExchangeSymbols []string `json:"exchange_symbols"`
}

// Ltp represents a map of instrument to it's last traded price
// https://groww.in/trade-api/docs/curl/live-data#response-1
type Ltp map[string]float32

func (l LtpRequest) queryParams() url.Values {
	out := make(url.Values)

	for _, es := range l.ExchangeSymbols {
		out.Add("exchange_symbols", es)
	}

	out.Add("segment", string(l.Segment))
	return out
}

// GetLtp : The API can be used to get the latest price of an instrument.
// Use the segment value FNO for derivatives and CASH for stocks and indices.
// Upto 50 instruments are supported for each api call.
// https://groww.in/trade-api/docs/curl/live-data#get-ltp
func (c *Client) GetLtp(ctx context.Context, req LtpRequest) (Ltp, error) {
	const destination = "https://api.groww.in/v1/live-data/ltp"
	return doGetRequest[Ltp](ctx, c, destination, req)
}

// OhlcRequest represents request for Client.GetOhlc
// https://groww.in/trade-api/docs/curl/live-data#request-2
type OhlcRequest struct {
	// Segment of the instrument such as CASH, FNO etc.
	Segment Segment `json:"segment"`
	// Array of trading symbols with their respective exchanges.
	// For example: `NSE_RELIANCE` `BSE_SENSEX` `NSE_NIFTY25APR24100PE`
	ExchangeSymbols []string `json:"exchange_symbols"`
}

// OhlcResponse represents a map of instrument to ohlc
// https://groww.in/trade-api/docs/curl/live-data#response-2
type OhlcResponse map[string]ohlcString

func (o OhlcRequest) queryParams() url.Values {
	out := make(url.Values)

	for _, es := range o.ExchangeSymbols {
		out.Add("exchange_symbols", es)
	}

	out.Add("segment", string(o.Segment))
	return out
}

// GetOhlc : The API can be used to get the ohlc data of an instrument.
// Use the segment value FNO for derivatives and CASH for stocks and indices. Upto 50 instruments are supported for each API call.
// Note: The OHLC data retrieved using the OHLC API reflects the current time's OHLC (i.e., real-time snapshot).
// For interval-based OHLC data (e.g., 1-minute, 5-minute candles), please refer to the Backtesting APIs.
// https://groww.in/trade-api/docs/curl/live-data#get-ohlc
func (c *Client) GetOhlc(ctx context.Context, req OhlcRequest) (OhlcResponse, error) {
	const destination = "https://api.groww.in/v1/live-data/ohlc"
	return doGetRequest[OhlcResponse](ctx, c, destination, req)
}

// GetGreeksRequest represents the request for Client.GetGreeks
// https://groww.in/trade-api/docs/curl/live-data#request-schema-3
type GetGreeksRequest struct {
	// Stock Exchange - NSE or BSE
	Exchange string `json:"exchange"`
	// Underlying symbol for the contract such as NIFTY, BANKNIFTY, RELIANCE etc.
	Underlying string `json:"underlying"`
	// Trading Symbol of the FNO contract as defined by the exchange
	TradingSymbol string `json:"trading_symbol"`
	// Expiry date of the contract in YYYY-MM-DD format
	Expiry time.Time `json:"expiry"`
}

// Greeks are financial measures that help assess the risk and sensitivity of options contracts to various factors like underlying price changes, time decay, volatility, and interest rates
type Greeks struct {
	// Delta measures the rate of change of option price based on every 1 rupee change in the price of underlying.
	Delta float32 `json:"delta"`
	// Gamma measures the rate of change of delta with respect to underlying asset price.
	// Higher gamma means delta changes more rapidly
	Gamma float32 `json:"gamma"`
	// Theta measures the rate of time decay of option price.
	// Usually negative, indicating option value decreases over time
	Theta float32 `json:"theta"`
	// Vega measures the rate of change of option price based on every 1% change in implied volatility of the underlying asset
	Vega float32 `json:"vega"`
	// Rho measures the sensitivity of option price to changes in interest rates
	Rho float32 `json:"rho"`
	// Implied Volatility represents the market's expectation of future volatility, expressed as a percentage
	Iv float32 `json:"iv"`
}

// GetGreeks : This API provides the complete Greeks data for FNO (Futures and Options) contracts.
// Greeks are financial measures that help assess the risk and sensitivity of options contracts to various factors like underlying price changes, time decay, volatility, and interest rates.
// This API is specifically designed for derivatives trading and risk management.
func (c *Client) GetGreeks(ctx context.Context, req GetGreeksRequest) (Greeks, error) {
	destination := fmt.Sprintf(
		"https://api.groww.in/v1/live-data/greeks/exchange/%s/underlying/%s/trading_symbol/%s/expiry/%s",
		req.Exchange,
		req.Underlying,
		req.TradingSymbol,
		req.Expiry.Format(time.DateOnly),
	)
	return doGetRequest[Greeks](ctx, c, destination, nil)
}
