package model

type AppResponseModel struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

type CallbackBniModel struct {
	TrxId           string  `json:"trxId`
	AccountNumber   string  `json:"accountNumber"`
	Amount          float32 `json:"amount"`
	ReferenceNumber string  `json:"referenceNumber"`
	TransactionDate string  `json:"transactionDate"`
}
