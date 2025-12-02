// https://groww.in/trade-api/docs/curl/annexures

package growwapi

// OrderStatus - https://groww.in/trade-api/docs/curl/annexures#order-status
type OrderStatus string

const (
	// OrderStatusNew - Order is newly created and pending for further processing
	OrderStatusNew OrderStatus = "NEW"

	// OrderStatusAcked - Order has been acknowledged by the system
	OrderStatusAcked OrderStatus = "ACKED"

	// OrderStatusTriggerPending - Order is waiting for a trigger event to be executed
	OrderStatusTriggerPending OrderStatus = "TRIGGER_PENDING"

	// OrderStatusApproved - Order has been approved and is ready for execution
	OrderStatusApproved OrderStatus = "APPROVED"

	// OrderStatusRejected - Order has been rejected by the system
	OrderStatusRejected OrderStatus = "REJECTED"

	// OrderStatusFailed - Order execution has failed
	OrderStatusFailed OrderStatus = "FAILED"

	// OrderStatusExecuted - Order has been successfully executed
	OrderStatusExecuted OrderStatus = "EXECUTED"

	// OrderStatusDeliveryAwaited - Order has been executed and waiting for delivery
	OrderStatusDeliveryAwaited OrderStatus = "DELIVERY_AWAITED"

	// OrderStatusCancelled - Order has been cancelled
	OrderStatusCancelled OrderStatus = "CANCELLED"

	// OrderStatusCancellationRequested - Request to cancel the order has been initiated
	OrderStatusCancellationRequested OrderStatus = "CANCELLATION_REQUESTED"

	// OrderStatusModificationRequested - Request to modify the order has been initiated
	OrderStatusModificationRequested OrderStatus = "MODIFICATION_REQUESTED"

	// OrderStatusCompleted - Order has been completed
	OrderStatusCompleted OrderStatus = "COMPLETED"
)

// AfterMarketOrderStatus - https://groww.in/trade-api/docs/curl/annexures#after-market-order-status
type AfterMarketOrderStatus = OrderStatus

const (
	// AfterMarketOrderStatusNa - Status not available
	AfterMarketOrderStatusNa AfterMarketOrderStatus = "NA"

	// AfterMarketOrderStatusPending - Order is pending for execution
	AfterMarketOrderStatusPending AfterMarketOrderStatus = "PENDING"

	// AfterMarketOrderStatusDispatched - Order has been dispatched for execution
	AfterMarketOrderStatusDispatched AfterMarketOrderStatus = "DISPATCHED"

	// AfterMarketOrderStatusParked - Order is parked for later execution
	AfterMarketOrderStatusParked AfterMarketOrderStatus = "PARKED"

	// AfterMarketOrderStatusPlaced - Order has been placed in the market
	AfterMarketOrderStatusPlaced AfterMarketOrderStatus = "PLACED"

	// AfterMarketOrderStatusFailed - Order execution has failed
	AfterMarketOrderStatusFailed AfterMarketOrderStatus = "FAILED"

	// AfterMarketOrderStatusMarket - Order is a market order
	AfterMarketOrderStatusMarket AfterMarketOrderStatus = "MARKET"
)

// Exchange - https://groww.in/trade-api/docs/curl/annexures#exchange
type Exchange string

const (
	// ExchangeBse - Bombay Stock Exchange - Asia's oldest exchange, known for SENSEX index
	ExchangeBse Exchange = "BSE"

	// ExchangeNse - National Stock Exchange - India's largest exchange by trading volume
	ExchangeNse Exchange = "NSE"
)

// Segment - https://groww.in/trade-api/docs/curl/annexures#segment
type Segment string

const (
	// SegmentCash - Regular equity market for trading stocks with delivery option
	SegmentCash Segment = "CASH"

	// SegmentFno - Futures and Options segment for trading derivatives contracts
	SegmentFno Segment = "FNO"
)

// OrderType - https://groww.in/trade-api/docs/curl/annexures#order-type
type OrderType string

const (
	// OrderTypeLimit - Specify exact price, may not get filled immediately but ensures price control
	OrderTypeLimit OrderType = "LIMIT"

	// OrderTypeMarket - Immediate execution at best available price, no price guarantee
	OrderTypeMarket OrderType = "MARKET"

	// OrderTypeStopLoss - Stop Loss - Protection order that triggers at specified price to limit losses
	OrderTypeStopLoss OrderType = "SL"

	// OrderTypeStopLossMarket - Stop Loss Market - Market order triggered at specified price to limit losses
	OrderTypeStopLossMarket OrderType = "SL_M"
)

// Product - https://groww.in/trade-api/docs/curl/annexures#product
type Product string

const (
	// ProductCnc - Cash and Carry - For delivery-based equity trading with full upfront payment
	ProductCnc Product = "CNC"

	// ProductMis - Margin Intraday Square-off - Higher leverage but must close by day end
	ProductMis Product = "MIS"

	// ProductNormal - Regular margin trading allowing overnight positions with standard leverage
	ProductNormal Product = "NRML"
)

// TransactionType - https://groww.in/trade-api/docs/curl/annexures#transaction-type
type TransactionType string

const (
	// TransactionTypeBuy - Long position - Profit from price increase, loss from price decrease
	TransactionTypeBuy TransactionType = "BUY"

	// TransactionTypeSell - Short position - Profit from price decrease, loss from price increase
	TransactionTypeSell TransactionType = "SELL"
)

// Validity - https://groww.in/trade-api/docs/curl/annexures#validity
type Validity string

// ValidityDay - Valid until market close on the same trading day
const ValidityDay Validity = "DAY"

// CandleInterval - https://groww.in/trade-api/docs/curl/annexures#candle-interval
type CandleInterval string

const (
	// CandleInterval1Min - 1 minute interval
	CandleInterval1Min CandleInterval = "1minute"

	// CandleInterval2Min - 2 minute interval
	CandleInterval2Min CandleInterval = "2minute"

	// CandleInterval3Min - 3 minute interval
	CandleInterval3Min CandleInterval = "3minute"

	// CandleInterval5Min - 5 minute interval
	CandleInterval5Min CandleInterval = "5minute"

	// CandleInterval10Min - 10 minute interval
	CandleInterval10Min CandleInterval = "10minute"

	// CandleInterval15Min - 15 minute interval
	CandleInterval15Min CandleInterval = "15minute"

	// CandleInterval30Min - 30 minute interval
	CandleInterval30Min CandleInterval = "30minute"

	// CandleInterval1Hour - 1 hour interval
	CandleInterval1Hour CandleInterval = "1hour"

	// CandleInterval4Hour - 4 hour interval
	CandleInterval4Hour CandleInterval = "4hour"

	// CandleInterval1Day - 1 day interval
	CandleInterval1Day CandleInterval = "1day"

	// CandleInterval1Week - 1 week interval
	CandleInterval1Week CandleInterval = "1week"

	// CandleInterval1Month - 1 month interval
	CandleInterval1Month CandleInterval = "1month"
)

// InstrumentType - https://groww.in/trade-api/docs/curl/annexures#instrument-type
type InstrumentType string

const (
	// InstrumentTypeEquity - Equity - Represents ownership in a company
	InstrumentTypeEquity InstrumentType = "EQ"

	// InstrumentTypeIndex - Index - Composite value of a group of stocks representing a market
	InstrumentTypeIndex InstrumentType = "IDX"

	// InstrumentTypeFutures - Futures - Derivatives contract to buy/sell an asset at a future date
	InstrumentTypeFutures InstrumentType = "FUT"

	// InstrumentTypeCallOption - Call Option - Derivatives contract giving the right to buy an asset
	InstrumentTypeCallOption InstrumentType = "CE"

	// InstrumentTypePutOption - Put Option - Derivatives contract giving the right to sell an asset
	InstrumentTypePutOption InstrumentType = "PE"
)
