package bnicallback

import (
	db "sakti-cashmanagement/database"
	model "sakti-cashmanagement/model"
)

func Insert(model model.CallbackPaymentModel) (string, int64) {

	var result string = ""
	var id int64 = 0

	db, err := db.Connect()
	if err != nil {
		result = err.Error()
		return result, id
	}
	defer db.Close()

	query := "INSERT INTO public.callback_payment (partner, reference_number, va_number, amount) VALUES($1, $2, $3, $4) RETURNING id"

	stmt, err := db.Prepare(query)
	if err != nil {
		result = err.Error()
		return result, id
	}
	defer stmt.Close()

	err = stmt.QueryRow(model.Partner, model.ReferenceNumber, model.VaNumber, model.Amount).Scan(&id)
	if err != nil {
		result = err.Error()
		return result, id
	}

	return result, id
}
