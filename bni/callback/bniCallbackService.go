package bnicallback

import (
	myDb "sakti-cashmanagement/database"
	model "sakti-cashmanagement/model"
)

func Insert(callbackBniModel model.CallbackBniModel) (string, int64) {

	var result string = ""
	var id int64 = 0
	var callbackPaymentModel model.CallbackPaymentModel

	db, err := myDb.Connect()
	if err != nil {
		result = err.Error()
		return result, id
	}
	defer db.Close()

	callbackPaymentModel.ReferenceNumber = "00001"
	callbackPaymentModel.PartnerCode = "BNI"
	callbackPaymentModel.AccountNumber = callbackBniModel.AccountNumber
	callbackPaymentModel.Amount = callbackBniModel.Amount

	query := "INSERT INTO public.callback_payment (partner_code, reference_number, account_number, amount) VALUES($1, $2, $3, $4) RETURNING id"

	stmt, err := db.Prepare(query)
	if err != nil {
		result = err.Error()
		return result, id
	}
	defer stmt.Close()

	err = stmt.QueryRow(callbackPaymentModel.PartnerCode,
		callbackPaymentModel.ReferenceNumber,
		callbackPaymentModel.AccountNumber,
		callbackPaymentModel.Amount).Scan(&id)

	if err != nil {
		result = err.Error()
		return result, id
	}

	return result, id
}
