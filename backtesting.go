// https://groww.in/trade-api/docs/curl/backtesting

package growwapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// GetExpiriesRequest represents the request for Client.GetExpiries
//
// https://groww.in/trade-api/docs/curl/backtesting#request-schema
type GetExpiriesRequest struct {
	// Stock Exchange
	Exchange Exchange `json:"exchange"`
	// Underlying symbol for which expiry dates are required (e.g., NIFTY, BANKNIFTY, RELIANCE)
	UnderlyingSymbol string `json:"underlying_symbol"`
	// Year for which expiry dates are required (2020 - current year). If year is not specified, current year is considered.
	Year int `json:"year"`
	// Month for which expiry dates are required (1-12). If month is not specified, expiries of the entire year is returned.
	Month int `json:"month"`
}

// GetExpiriesResponse represents the response for Client.GetExpiries
//
// https://groww.in/trade-api/docs/curl/backtesting#response-schema
type GetExpiriesResponse struct {
	Expiries []Time `json:"expiries"`
}

func (g GetExpiriesRequest) queryParams() url.Values {
	out := make(url.Values)
	out.Add("exchange", string(g.Exchange))
	out.Add("underlying_symbol", g.UnderlyingSymbol)

	if g.Year != 0 {
		out.Add("year", strconv.Itoa(g.Year))
	}

	if g.Month != 0 {
		out.Add("month", strconv.Itoa(g.Month))
	}

	return out
}

// GetExpiries : This API retrieves available expiry dates for derivatives instruments (FNO) for a given exchange and underlying symbol.
// Useful for backtesting options and futures strategies.
// Data of FNO instruments are available from 2020.
//
// https://groww.in/trade-api/docs/curl/backtesting#get-expiries
func (c *Client) GetExpiries(ctx context.Context, req GetExpiriesRequest) (GetExpiriesResponse, error) {
	const destination = "https://api.groww.in/v1/historical/expiries"
	return doGetRequest[GetExpiriesResponse](ctx, c, destination, req)
}

// GetContractsRequest represents the request for Client.GetContracts
//
// https://groww.in/trade-api/docs/curl/backtesting#request-schema-1
type GetContractsRequest struct {
	// Stock Exchange
	Exchange Exchange `json:"exchange"`
	// Underlying symbol for which expiry dates are required (e.g., NIFTY, BANKNIFTY, RELIANCE)
	UnderlyingSymbol string `json:"underlying_symbol"`
	// Expiry date for which contracts are required
	ExpiryDate time.Time `json:"expiry_date"`
}

// GetContractsResponse represents the response for Client.GetContracts
//
// https://groww.in/trade-api/docs/curl/backtesting#response-schema
type GetContractsResponse struct {
	Contracts []Time `json:"contracts"`
}

func (g GetContractsRequest) queryParams() url.Values {
	out := make(url.Values)
	out.Add("exchange", string(g.Exchange))
	out.Add("underlying_symbol", g.UnderlyingSymbol)
	out.Add("expiry_date", g.ExpiryDate.Format(time.DateOnly))
	return out
}

// GetContracts : This API retrieves available contract symbols for derivatives instruments for a given exchange, underlying symbol, and expiry date.
// Essential for backtesting specific options or futures contracts.
// Data of FNO instruments are available from 2020.
//
// https://groww.in/trade-api/docs/curl/backtesting#get-contracts
func (c *Client) GetContracts(ctx context.Context, req GetContractsRequest) (GetContractsResponse, error) {
	const destination = "https://api.groww.in/v1/historical/contracts"
	return doGetRequest[GetContractsResponse](ctx, c, destination, req)
}

// GetHistoricalCandlesRequest represents the request for Client.GetHistoricalCandles
//
// https://groww.in/trade-api/docs/curl/backtesting#request-schema-2
type GetHistoricalCandlesRequest struct {
	// Stock Exchange
	Exchange Exchange `json:"exchange"`
	// Segment of the instrument such as CASH, FNO etc.
	Segment Segment `json:"segment"`
	// Groww symbol of the instrument for which historical data is required
	GrowwSymbol string `json:"groww_symbol"`
	// Start time from which data is required
	StartTime time.Time `json:"start_time"`
	// End time until which data is required
	EndTime time.Time `json:"end_time"`
	// Interval for which data is required.
	CandleInterval CandleInterval `json:"candle_interval"`
}

// Candle represents a candle date for an interval
//
// https://groww.in/trade-api/docs/curl/backtesting#response-schema-2
type Candle struct {
	// Timestamp of candle
	Timestamp Time
	Ohlcv
	// Open interest (only for FNO instruments, null for others)
	OpenInterest *float32
}

func (c *Candle) UnmarshalJSON(bytes []byte) error {
	asString := string(bytes)
	var arr [][]byte

	if err := json.Unmarshal(bytes, &arr); err != nil {
		return fmt.Errorf("parse into array(%q): %w", asString, err)
	}

	if len(arr) != 7 {
		return fmt.Errorf("expected 7 items, got %d for %q", len(arr), asString)
	}

	// parse timestamp to string
	var timestamp Time
	if err := json.Unmarshal(arr[0], &timestamp); err != nil {
		return fmt.Errorf("parse timestamp(%q): %w", asString, err)
	}

	// parse ohlcv
	ohlcvData := make([]float32, 5)
	for i := 0; i < 5; i++ {
		if err := json.Unmarshal(arr[i+1], &ohlcvData[i]); err != nil {
			return fmt.Errorf("parse ohlcvData(%q): %w", asString, err)
		}
	}

	// parse oi
	var oi *float32
	if err := json.Unmarshal(arr[6], &oi); err != nil {
		return fmt.Errorf("parse oi(%q): %w", asString, err)
	}

	c.Timestamp = timestamp
	c.Open = ohlcvData[0]
	c.High = ohlcvData[1]
	c.Low = ohlcvData[2]
	c.Close = ohlcvData[3]
	c.Volume = ohlcvData[4]
	c.OpenInterest = oi

	return nil
}

// HistoricalCandlesData represents the response of Client.GetHistoricalCandles
//
// https://groww.in/trade-api/docs/curl/backtesting#response-schema-2
type HistoricalCandlesData struct {
	// Candle data
	Candles []Candle `json:"candles"`
	// Closing price of the instrument
	ClosingPrice float32 `json:"closing_price"`
	// Start time
	StartTime Time `json:"start_time"`
	// End time
	EndTime Time `json:"end_time"`
	// Interval in minutes
	IntervalInMinutes int `json:"interval_in_minutes"`
}

func (g GetHistoricalCandlesRequest) queryParams() url.Values {
	out := make(url.Values)
	out.Add("exchange", string(g.Exchange))
	out.Add("segment", string(g.Segment))
	out.Add("groww_symbol", g.GrowwSymbol)
	out.Add("start_time", g.StartTime.Format(time.DateTime))
	out.Add("end_time", g.EndTime.Format(time.DateTime))
	out.Add("candle_interval", string(g.CandleInterval))
	return out
}

// GetHistoricalCandles Fetch historical candle data for backtesting trading strategies.
// This API provides:
//   - Historical OHLC (Open, High, Low, Close) data for all instruments.
//   - Volume for tradable instruments (Equities and FNO)
//   - Open Interest (OI) for FNO Data of Equities, Indices and FNO instruments are available from 2020.
//
// https://groww.in/trade-api/docs/curl/backtesting#get-historical-candle-data
func (c *Client) GetHistoricalCandles(ctx context.Context, req GetHistoricalCandlesRequest) (HistoricalCandlesData, error) {
	const destination = "https://api.groww.in/v1/historical/candles"
	return doGetRequest[HistoricalCandlesData](ctx, c, destination, req)
}
