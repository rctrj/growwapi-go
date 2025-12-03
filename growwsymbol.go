// https://groww.in/trade-api/docs/curl/backtesting#groww-symbol

package growwapi

import (
	"fmt"
	"time"
)

type (
	growwSymbolGenerator byte
	optionType           string
)

const (
	// GrowwSymbol can be used to generate groww symbol
	GrowwSymbol growwSymbolGenerator = 0

	OptionTypeCE optionType = "CE"
	OptionTypePE optionType = "PE"
)

func (g growwSymbolGenerator) Equity(exchange Exchange, tradingSymbol string) string {
	return fmt.Sprintf("%s-%s", exchange, tradingSymbol)
}

func (g growwSymbolGenerator) Index(exchange Exchange, tradingSymbol string) string {
	return fmt.Sprintf("%s-%s", exchange, tradingSymbol)
}

func (g growwSymbolGenerator) Future(exchange Exchange, tradingSymbol string, expiry time.Time) string {
	return fmt.Sprintf("%s-%s-%s-FUT", exchange, tradingSymbol, expiry.Format("02Jan06"))
}

func (g growwSymbolGenerator) Option(
	exchange Exchange,
	tradingSymbol string,
	expiry time.Time,
	strikePrice float32,
	optionType optionType,
) string {
	return fmt.Sprintf(
		"%s-%s-%s-%v-%s",
		exchange,
		tradingSymbol,
		expiry.Format("02Jan06"),
		strikePrice,
		optionType,
	)
}

func (g growwSymbolGenerator) CallOption(
	exchange Exchange,
	tradingSymbol string,
	expiry time.Time,
	strikePrice float32,
) string {
	return g.Option(exchange, tradingSymbol, expiry, strikePrice, OptionTypeCE)
}

func (g growwSymbolGenerator) PutOption(
	exchange Exchange,
	tradingSymbol string,
	expiry time.Time,
	strikePrice float32,
) string {
	return g.Option(exchange, tradingSymbol, expiry, strikePrice, OptionTypePE)
}
