package bnicallback

import (
	myDb "sakti-cashmanagement/database"
	model "sakti-cashmanagement/model"
)

func Insert(callbackBniModel model.CallbackBniModel) (string, int64) {

	var result string = ""
	var id int64 = 0

	db, err := myDb.Connect()
	if err != nil {
		result = err.Error()
		return result, id
	}
	defer db.Close()

	var referenceNumber string = callbackBniModel.ReferenceNumber
	var partnerCode string = "BNI"
	var accountNumber string = callbackBniModel.AccountNumber
	var amount float32 = callbackBniModel.Amount

	query := "INSERT INTO public.callback_payment (partner_code, reference_number, account_number, amount) VALUES($1, $2, $3, $4) RETURNING id"

	stmt, err := db.Prepare(query)
	if err != nil {
		result = err.Error()
		return result, id
	}
	defer stmt.Close()

	err = stmt.QueryRow(partnerCode,
		referenceNumber,
		accountNumber,
		amount).Scan(&id)

	if err != nil {
		result = err.Error()
		return result, id
	}

	return result, id
}
