package bnicallback

import (
	"encoding/json"
	"net/http"
	model "sakti-cashmanagement/model"
)

func Save(response http.ResponseWriter, request *http.Request) {

	var callbackPaymentModel model.CallbackPaymentModel
	var appResponseModel model.AppResponseModel

	err := json.NewDecoder(request.Body).Decode(&callbackPaymentModel)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	var result string = ""
	if callbackPaymentModel.Id == 0 {
		result = AddNew(callbackPaymentModel)
	}

	if len(result) > 0 {

		appResponseModel.Status = "001"
		appResponseModel.Message = result

		response.WriteHeader(http.StatusInternalServerError)
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(appResponseModel)

	} else {

		appResponseModel.Status = "000"
		appResponseModel.Message = ""

		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(appResponseModel)

	}

	return
}
