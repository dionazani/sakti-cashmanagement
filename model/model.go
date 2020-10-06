package model

type AppResponseModel struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

type CallbackPaymentModel struct {
	Id              int32   `json:"id"`
	PartnerCode     string  `json:"partner"`
	ReferenceNumber string  `json:"referenceNumber"`
	AccountNumber   string  `json:"vaNumber"`
	Amount          float32 `json:"amount"`
	Step            string  `json:"step"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

type CallbackBniModel struct {
	TrxId         string  `json:"trxId`
	AccountNumber string  `json:"accountNumber"`
	Amount        float32 `json:"amount"`
}
