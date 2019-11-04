package main

// CheckoutValidationRequest type
// LS will receive data like this from Checkout-Internal
type CheckoutValidationRequest struct {
	UserID   int
	Metadata MetadataCI
}

// CheckoutValidationResponse is a reponse from LS to Checkout-Internal
type CheckoutValidationResponse struct {
	Order    OrderInfo `json:"order"`
	Metadata string    `json:"metadata"`
}

// MetadataCI Type
type MetadataCI struct {
	Products []ProductFromCI `json:"products"`
	Shipping ShippingFromCI  `json:"shipping"`
}

// OrderInfo type
type OrderInfo struct {
	ID     int `json:"id"`     // ID is order_id that created in localservices
	Amount int `json:"amount"` // Amount is total of Product and Shipping prices
}

type ProductFromCI struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Volume        int    `json:"volume"`
	Dimension     string `json:"dimension,omitempty"`
	TotalWeight   int    `json:"total_weight"` // gram
	TotalPrice    int    `json:"total_price"`
	PriceOptionID int    `json:"price_option_id"`
	Qty           int    `json:"qty"`        // how many the product ordered
	QtyOption     int    `json:"qty_option"` // how many contents in a product
	SpecSet       string `json:"spec_set"`   // string json
	Notes         string `json:"notes"`
}

type ShippingFromCI2 struct {
	ID    int `json:"shipping_id"` // logistic rates service id
	SpID  int `json:"sp_id"`       // logistic rates service product id
	Price int `json:"shipping_price"`

	Insurance       int `json:"insurance"`
	InsurancePrice  int `json:"insurance_price"`
	OriginAddressID int `json:"origin_address_id"` // Tkp district id (partner location)
	DestAddressID   int `json:"dest_address_id"`   // Tkp user address id
}

type ShippingFromCI struct {
	ID              int `json:"shipping_id"` // logistic rates service id
	SpID            int `json:"sp_id"`       // logistic rates service product id
	Price           int `json:"shipping_price"`
	Insurance       int `json:"insurance"`
	InsurancePrice  int `json:"insurance_price"`
	Weight          int `json:"weight"`
	OrderValue      int `json:"order_value"`
	OriginAddressID int `json:"origin_address_id"` // Tkp district id (partner location)
	DestAddressID   int `json:"dest_address_id"`   // Tkp user address id
}
