package response

import "go-fiber-test/models"

func BuildDataResponse(data map[string]string, message string, success bool, status int) models.APIResponse {
	var apiResponse models.APIResponse

	apiResponse.Data = data
	apiResponse.Message = message
	apiResponse.Status = status
	apiResponse.Success = success

	return apiResponse
}
