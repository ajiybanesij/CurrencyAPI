package Models

type CreateOffer struct {
	FromCurrency string `json:"fromCurrency" binding:"required"`
	ToCurrency   string `json:"toCurrency" binding:"required"`
}

type AcceptOffer struct {
	OfferID string  `json:"offerId" binding:"required"`
	Amount  float64 `json:"amount" binding:"required,gte=0"`
}
