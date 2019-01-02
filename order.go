package poloniex

import (
	"encoding/json"
	"errors"
	"strconv"
)

// OrderBook is a list of open orders for a currency
type OrderBook struct {
	Asks     []*Order `json:"asks"`
	Bids     []*Order `json:"bids"`
	IsFrozen int      `json:"isFrozen,string"`
	Error    string   `json:"error"`
}

// Order is an order in an OrderBook
type Order struct {
	Rate   float64
	Amount float64
}

// UnmarshalJSON parses the order as poloniex gives it
func (o *Order) UnmarshalJSON(data []byte) error {
	vals := []interface{}{}
	if err := json.Unmarshal(data, &vals); err != nil {
		return err
	}
	if len(vals) != 2 {
		return errors.New("invalid order data")
	}
	if rateStr, ok := vals[0].(string); ok {
		rate, err := strconv.ParseFloat(rateStr, 64)
		if err != nil {
			return err
		}
		o.Rate = rate
	} else {
		return errors.New("invalid rate value")
	}

	if amount, ok := vals[1].(float64); ok {
		o.Amount = amount
	} else {
		return errors.New("invalid amount value")
	}

	return nil
}

// OpenOrder is an open order for the account
type OpenOrder struct {
	OrderNumber int64   `json:"orderNumber,string"`
	Type        string  `json:"type"`
	Rate        float64 `json:"rate,string"`
	Amount      float64 `json:"amount,string"`
	Total       float64 `json:"total,string"`
}
