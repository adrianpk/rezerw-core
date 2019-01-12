package core

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

var (
	ct    = "Content-Type"
	ctVal = "application/json"
)

// RenderOk - Render Ok.
func RenderOk(res []byte) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			ct: ctVal,
		},
		Body: string(res),
	}, nil
}

// ErrorResult - Error result
type ErrorResult struct {
	Err string `json:"err"`
}

// RenderError - Render error.
func RenderError(err error) (events.APIGatewayProxyResponse, error) {
	log.Print(err.Error())
	res, err2 := json.Marshal(ErrorResult{Err: err.Error()})
	log.Print(err2.Error())
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers: map[string]string{
			ct: ctVal,
		},
		Body: string(res),
	}, err
}
