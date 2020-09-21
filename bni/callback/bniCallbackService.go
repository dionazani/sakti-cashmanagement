package bnicallback

import (
	db "sakti-cashmanagement/database"
	model "sakti-cashmanagement/model"
)

func AddNew(model model.CallbackPaymentModel) string {

	var result string = ""

	db, err := db.Connect()
	if err != nil {
		result = err.Error()
		return result
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO public.callback_payment (partner, reference_number, va_number, amount) VALUES($1, $2, $3, $4)",
		model.Partner, model.ReferenceNumber, model.VaNumber, model.Amount)

	if err != nil {
		result = err.Error()
		return result
	}

	return result
}
