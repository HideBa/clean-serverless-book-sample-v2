package controller

import (
	"clean-serverless-book-sample-v2/registry"
	"clean-serverless-book-sample-v2/usecase"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type PostHelloRequest struct {
	Name string `json:"name"`
}

type HelloMessageResponse struct {
	Message string `json:"message"`
}

func HelloMessageSettigsValidator() *Validator {
	return &Validator{
		Settings: []*ValidatorSetting{
			{
				ArgName:      "name",
				ValidateTags: "required",
			},
		},
	}
}

func PostHello(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	validator := HelloMessageSettigsValidator()
	validErr := validator.ValidateBody(request.Body)
	if validErr != nil {
		return Response400(validErr)
	}

	var req PostHelloRequest
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return Response500(err)
	}

	h := registry.GetFactory().BuildCreateHelloMessage()
	res, err := h.Execute(&usecase.CreateHelloMessageRequest{Name: req.Name})
	if err != nil {
		return Response500(err)
	}

	return Response200(&HelloMessageResponse{
		Message: res.Message,
	})
}
