package model

type AppResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type CallbackPaymentModel struct {
	Id              int32   `json:"id"`
	Partner         string  `json:"partner"`
	ReferenceNumber string  `json:"referenceNumber"`
	VaNumber        string  `json:"vaNumber"`
	Amount          float32 `json:"amount"`
	Step            string  `json:"step"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}
