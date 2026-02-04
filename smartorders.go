// https://groww.in/trade-api/docs/curl/smart-orders

package growwapi

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// SmartOrderType represents the type of smart order
type SmartOrderType string

const (
	// SmartOrderTypeGtt - Good Till Triggered order
	SmartOrderTypeGtt SmartOrderType = "GTT"
	// SmartOrderTypeOco - One Cancels the Other order
	SmartOrderTypeOco SmartOrderType = "OCO"
)

// TriggerDirection represents the direction of price trigger
type TriggerDirection string

const (
	// TriggerDirectionUp - Trigger when price goes up
	TriggerDirectionUp TriggerDirection = "UP"
	// TriggerDirectionDown - Trigger when price goes down
	TriggerDirectionDown TriggerDirection = "DOWN"
)

// SmartOrderStatus represents the status of a smart order
type SmartOrderStatus string

const (
	// SmartOrderStatusActive - Smart order is active
	SmartOrderStatusActive SmartOrderStatus = "ACTIVE"
	// SmartOrderStatusCancelled - Smart order has been cancelled
	SmartOrderStatusCancelled SmartOrderStatus = "CANCELLED"
	// SmartOrderStatusCompleted - Smart order has been completed
	SmartOrderStatusCompleted SmartOrderStatus = "COMPLETED"
)

// Duration represents the validity duration of a smart order
type Duration string

const (
	// DurationDay - Valid until market close on the same trading day
	DurationDay Duration = "DAY"
)

// GttOrderDetails represents the order details for a GTT order
type GttOrderDetails struct {
	// Order type - LIMIT or SL
	OrderType OrderType `json:"order_type"`
	// Limit price (required for LIMIT/SL)
	Price float32 `json:"price,omitempty"`
	// Transaction type - BUY or SELL
	TransactionType TransactionType `json:"transaction_type"`
}

// CreateGttOrderRequest represents the request for creating a GTT order
//
// https://groww.in/trade-api/docs/curl/smart-orders#request-schema
type CreateGttOrderRequest struct {
	// Idempotency key, 8-20 alphanumeric chars with at most 2 hyphens
	ReferenceId string `json:"reference_id"`
	// Smart order type - must be GTT
	SmartOrderType SmartOrderType `json:"smart_order_type"`
	// Segment - CASH or FNO
	Segment Segment `json:"segment"`
	// Trading Symbol of the instrument as defined by the exchange
	TradingSymbol string `json:"trading_symbol"`
	// Order quantity
	Quantity int `json:"quantity"`
	// Trigger price threshold
	TriggerPrice float32 `json:"trigger_price,string"`
	// Trigger direction - UP or DOWN
	TriggerDirection TriggerDirection `json:"trigger_direction"`
	// Order details
	Order GttOrderDetails `json:"order"`
	// Product type - CNC or MIS
	ProductType Product `json:"product_type"`
	// Stock Exchange
	Exchange Exchange `json:"exchange"`
	// Validity duration
	Duration Duration `json:"duration"`
}

// OcoTargetOrder represents the target order details for an OCO order
type OcoTargetOrder struct {
	// Trigger price for target
	TriggerPrice float32 `json:"trigger_price,string"`
	// Order type - LIMIT or MARKET
	OrderType OrderType `json:"order_type"`
	// Limit price (required if order_type is LIMIT)
	Price float32 `json:"price,omitempty"`
}

// OcoStopLossOrder represents the stop loss order details for an OCO order
type OcoStopLossOrder struct {
	// Trigger price for stop loss
	TriggerPrice float32 `json:"trigger_price,string"`
	// Order type - SL or SL_M
	OrderType OrderType `json:"order_type"`
	// Limit price (required if order_type is SL)
	Price float32 `json:"price,omitempty"`
}

// CreateOcoOrderRequest represents the request for creating an OCO order
//
// https://groww.in/trade-api/docs/curl/smart-orders#request-schema-1
type CreateOcoOrderRequest struct {
	// Idempotency key, 8-20 alphanumeric chars with at most 2 hyphens
	ReferenceId string `json:"reference_id"`
	// Smart order type - must be OCO
	SmartOrderType SmartOrderType `json:"smart_order_type"`
	// Segment - FNO (OCO for CASH segment not currently available)
	Segment Segment `json:"segment"`
	// Trading Symbol of the instrument as defined by the exchange
	TradingSymbol string `json:"trading_symbol"`
	// Quantity (must be <= abs(net_position_quantity))
	Quantity int `json:"quantity"`
	// Current net position quantity
	NetPositionQuantity int `json:"net_position_quantity"`
	// Transaction type - BUY or SELL
	TransactionType TransactionType `json:"transaction_type"`
	// Target order details
	Target OcoTargetOrder `json:"target"`
	// Stop loss order details
	StopLoss OcoStopLossOrder `json:"stop_loss"`
	// Product type - MIS for CASH, MIS or NRML for FNO
	ProductType Product `json:"product_type"`
	// Stock Exchange
	Exchange Exchange `json:"exchange"`
	// Validity duration
	Duration Duration `json:"duration"`
}

// SmartOrder represents a smart order (GTT or OCO)
//
// https://groww.in/trade-api/docs/curl/smart-orders#response-schema
type SmartOrder struct {
	// Smart order ID
	SmartOrderId string `json:"smart_order_id"`
	// Reference ID
	ReferenceId string `json:"reference_id"`
	// Smart order type
	SmartOrderType SmartOrderType `json:"smart_order_type"`
	// Status of the smart order
	Status SmartOrderStatus `json:"status"`
	// Segment
	Segment Segment `json:"segment"`
	// Trading symbol
	TradingSymbol string `json:"trading_symbol"`
	// Quantity
	Quantity int `json:"quantity"`
	// Trigger price (for GTT)
	TriggerPrice float32 `json:"trigger_price,omitempty"`
	// Trigger direction (for GTT)
	TriggerDirection TriggerDirection `json:"trigger_direction,omitempty"`
	// Order details (for GTT)
	Order *GttOrderDetails `json:"order,omitempty"`
	// Target order (for OCO)
	Target *OcoTargetOrder `json:"target,omitempty"`
	// Stop loss order (for OCO)
	StopLoss *OcoStopLossOrder `json:"stop_loss,omitempty"`
	// Product type
	ProductType Product `json:"product_type"`
	// Exchange
	Exchange Exchange `json:"exchange"`
	// Duration
	Duration Duration `json:"duration"`
	// Whether modification is allowed
	ModificationAllowed bool `json:"modification_allowed"`
	// Whether cancellation is allowed
	CancellationAllowed bool `json:"cancellation_allowed"`
	// Created at timestamp
	CreatedAt Time `json:"created_at"`
	// Updated at timestamp
	UpdatedAt Time `json:"updated_at"`
}

// CreateSmartOrder : This API is used to create a smart order (GTT or OCO).
// Note: COMMODITY segment is not supported; OCO for CASH segment is not currently available.
//
// https://groww.in/trade-api/docs/curl/smart-orders#create-gtt-order
// https://groww.in/trade-api/docs/curl/smart-orders#create-oco-order
func (c *Client) CreateSmartOrder(ctx context.Context, req any) (SmartOrder, error) {
	const destination = "https://api.groww.in/v1/order-advance/create"
	return doPostRequest[SmartOrder](ctx, c, destination, req)
}

// ModifyGttOrderRequest represents the request for modifying a GTT order
//
// https://groww.in/trade-api/docs/curl/smart-orders#request-schema-2
type ModifyGttOrderRequest struct {
	// Quantity
	Quantity int `json:"quantity,omitempty"`
	// Trigger price
	TriggerPrice float32 `json:"trigger_price,string,omitempty"`
	// Trigger direction
	TriggerDirection TriggerDirection `json:"trigger_direction,omitempty"`
	// Order details
	Order *GttOrderDetails `json:"order,omitempty"`
}

// ModifyOcoOrderRequest represents the request for modifying an OCO order
//
// https://groww.in/trade-api/docs/curl/smart-orders#request-schema-2
type ModifyOcoOrderRequest struct {
	// Quantity
	Quantity int `json:"quantity,omitempty"`
	// Duration
	Duration Duration `json:"duration,omitempty"`
	// Product type
	ProductType Product `json:"product_type,omitempty"`
	// Target trigger price
	Target *OcoTargetOrder `json:"target,omitempty"`
	// Stop loss trigger price
	StopLoss *OcoStopLossOrder `json:"stop_loss,omitempty"`
}

// ModifySmartOrder : This API is used to modify a smart order (GTT or OCO).
// For fundamental changes (symbol, segment, type), cancel and recreate the order.
//
// https://groww.in/trade-api/docs/curl/smart-orders#modify-smart-order
func (c *Client) ModifySmartOrder(ctx context.Context, smartOrderId string, req any) (SmartOrder, error) {
	destination := fmt.Sprintf("https://api.groww.in/v1/order-advance/modify/%s", smartOrderId)
	return doPutRequest[SmartOrder](ctx, c, destination, req)
}

// CancelSmartOrderRequest represents the request for cancelling a smart order
type CancelSmartOrderRequest struct {
	// Segment - CASH or FNO
	Segment Segment
	// Smart order type - GTT or OCO
	SmartOrderType SmartOrderType
	// Smart order ID
	SmartOrderId string
}

// CancelSmartOrder : This API is used to cancel a smart order (GTT or OCO).
//
// https://groww.in/trade-api/docs/curl/smart-orders#cancel-smart-order
func (c *Client) CancelSmartOrder(ctx context.Context, req CancelSmartOrderRequest) (SmartOrder, error) {
	destination := fmt.Sprintf(
		"https://api.groww.in/v1/order-advance/cancel/%s/%s/%s",
		req.Segment,
		req.SmartOrderType,
		req.SmartOrderId,
	)
	return doPostRequest[SmartOrder](ctx, c, destination, nil)
}

// GetSmartOrderRequest represents the request for getting a smart order
type GetSmartOrderRequest struct {
	// Segment - CASH or FNO
	Segment Segment
	// Smart order type - GTT or OCO
	SmartOrderType SmartOrderType
	// Smart order ID
	SmartOrderId string
}

// GetSmartOrder : This API retrieves the details of a smart order.
//
// https://groww.in/trade-api/docs/curl/smart-orders#get-smart-order
func (c *Client) GetSmartOrder(ctx context.Context, req GetSmartOrderRequest) (SmartOrder, error) {
	destination := fmt.Sprintf(
		"https://api.groww.in/v1/order-advance/status/%s/%s/internal/%s",
		req.Segment,
		req.SmartOrderType,
		req.SmartOrderId,
	)
	return doGetRequest[SmartOrder](ctx, c, destination, nil)
}

// ListSmartOrdersRequest represents the request for listing smart orders
//
// https://groww.in/trade-api/docs/curl/smart-orders#request-schema-4
type ListSmartOrdersRequest struct {
	// [Optional] Segment filter
	Segment Segment
	// Smart order type - GTT or OCO (default: OCO)
	SmartOrderType SmartOrderType
	// [Optional] Status filter - ACTIVE, CANCELLED, COMPLETED (default: ACTIVE)
	Status SmartOrderStatus
	// [Optional] Page number starting at 0 (default: 0)
	Page int
	// [Optional] Records per page, 1-50 (default: 10)
	PageSize int
	// [Optional] Start date time in ISO8601 format
	StartDateTime string
	// [Optional] End date time in ISO8601 format (max 1-month range from start)
	EndDateTime string
}

func (l ListSmartOrdersRequest) queryParams() url.Values {
	out := make(url.Values)

	if l.Segment != "" {
		out.Add("segment", string(l.Segment))
	}

	if l.SmartOrderType != "" {
		out.Add("smart_order_type", string(l.SmartOrderType))
	}

	if l.Status != "" {
		out.Add("status", string(l.Status))
	}

	if l.Page != 0 {
		out.Add("page", strconv.Itoa(l.Page))
	}

	if l.PageSize != 0 {
		out.Add("page_size", strconv.Itoa(l.PageSize))
	}

	if l.StartDateTime != "" {
		out.Add("start_date_time", l.StartDateTime)
	}

	if l.EndDateTime != "" {
		out.Add("end_date_time", l.EndDateTime)
	}

	return out
}

// ListSmartOrders : This API retrieves a list of smart orders.
//
// https://groww.in/trade-api/docs/curl/smart-orders#list-smart-orders
func (c *Client) ListSmartOrders(ctx context.Context, req ListSmartOrdersRequest) ([]SmartOrder, error) {
	const destination = "https://api.groww.in/v1/order-advance/list"
	return doGetRequest[[]SmartOrder](ctx, c, destination, req)
}
