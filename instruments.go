// https://groww.in/trade-api/docs/curl/instruments

package growwapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gocarina/gocsv"
)

// Date represents a date coming from groww apis. It's generally represented as time.DateOnly
// Date.Time is nil if empty string was received
type Date struct {
	*time.Time
}

// Instrument - https://groww.in/trade-api/docs/curl/instruments#instrument-csv-columns
type Instrument struct {
	// The exchange where the instrument is traded
	Exchange Exchange `csv:"exchange"`
	// The unique token assigned to the instrument by the exchange
	ExchangeToken string `csv:"exchange_token"`
	// The trading symbol of the instrument to place orders with
	TradingSymbol string `csv:"trading_symbol"`
	// The symbol used by Groww to identify the instrument
	GrowwSymbol string `csv:"groww_symbol"`
	// The name of the instrument
	Name string `csv:"name"`
	// The type of the instrument
	InstrumentType InstrumentType `csv:"instrument_type"`
	// Segment of the instrument such as CASH, FNO etc.
	Segment Segment `csv:"segment"`
	// The series of the instrument (e.g., EQ, A, B, etc.)
	Series string `csv:"series"`
	// The ISIN (International Securities Identification number) of the instrument
	Isin string `csv:"isin"`
	// The symbol of the underlying asset (for derivatives). Empty for stocks and indices
	UnderlyingSymbol string `csv:"underlying_symbol"`
	// The exchange token of the underlying asset
	UnderlyingExchangeToken string `csv:"underlying_exchange_token"`
	// The minimum lot size for trading the instrument
	LotSize int `csv:"lot_size"`
	// The expiry date of the instrument (for Derivatives)
	ExpiryDate Date `csv:"expiry_date"`
	// The strike price of the instrument (for Options)
	StrikePrice int `csv:"strike_price"`
	// The minimum price movement for the instrument
	TickSize float32 `csv:"tick_size"`
	// The quantity that is frozen for trading
	FreezeQuantity int `csv:"freeze_quantity"`
	// Whether the instrument is reserved for trading
	IsReserved bool `csv:"is_reserved"`
	// Whether buying the instrument is allowed
	BuyAllowed bool `csv:"buy_allowed"`
	// Whether selling the instrument is allowed
	SellAllowed bool `csv:"sell_allowed"`
}

func (c *Client) Instruments() ([]Instrument, error) {
	return Instruments(c.httpClient)
}

// Instruments fetches and returns details of all instruments
func Instruments(httpClient *http.Client) ([]Instrument, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	const instrumentsUrl = "https://growwapi-assets.groww.in/instruments/instrument.csv"

	resp, err := httpClient.Get(instrumentsUrl)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Get(%q): %w", instrumentsUrl, err)
	}
	defer resp.Body.Close()

	var out []Instrument
	if err := gocsv.Unmarshal(resp.Body, &out); err != nil {
		return nil, fmt.Errorf("gocsv.Unmarshal(): %w", err)
	}

	return out, nil
}

func (d *Date) UnmarshalCSV(in string) error {
	if in == "" {
		d.Time = nil
		return nil
	}

	parsed, err := time.Parse(time.DateOnly, in)
	if err != nil {
		return fmt.Errorf("time.Parse(%q): %w", in, err)
	}

	d.Time = &parsed
	return nil
}
