package model

type GetResourceResponse struct {
	DiscountType   string  `json:"discount_type"`
	DiscountAmount float64 `json:"discount_amount"`
	TotalPrice     float64 `json:"total_price"`
	Description    string  `json:"description"`
}
