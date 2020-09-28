package bnicallback

import (
	"encoding/json"
	"net/http"
	model "sakti-cashmanagement/model"
	"strconv"
)

func AddNew(response http.ResponseWriter, request *http.Request) {

	var callbackPaymentModel model.CallbackPaymentModel
	var appResponseModel model.AppResponseModel
	var id int64 = 0
	data := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&callbackPaymentModel)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	var result string = ""
	if callbackPaymentModel.Id == 0 {
		result, id = Insert(callbackPaymentModel)
		data["callbackPaymentId"] = strconv.FormatInt(id, 10)
	}

	if len(result) > 0 {

		appResponseModel.Status = "001"
		appResponseModel.Message = result

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(appResponseModel)

	} else {

		appResponseModel.Status = "000"
		appResponseModel.Message = ""
		appResponseModel.Data = data

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(appResponseModel)
	}

	return
}
