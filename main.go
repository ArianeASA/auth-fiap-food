package main

import (
	"auth-fiap-food/adapters"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"strings"
	"time"

	"net/http"
)

type errorResponse struct {
	Cause string
}

func newError(msg string) errorResponse {
	return errorResponse{Cause: msg}
}

func handleGetToken(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	credential := &adapters.Credentials{}
	err := json.Unmarshal([]byte(req.Body), &credential)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers(),
			Body:       getJson(newError(err.Error())),
		}, nil
	}

	token, err := GetToken(credential)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       getJson(newError(err.Error())),
			Headers:    headers(),
		}, nil
	}

	//obj, err := json.Marshal(token)
	//if err != nil {
	//	return events.APIGatewayProxyResponse{}, err
	//}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       getJson(token),
		Headers:    headers(),
	}, nil

}

func getJson(obj interface{}) string {
	objJson, _ := json.Marshal(obj)
	return string(objJson)
}

func GetToken(cred *adapters.Credentials) (adapters.AuthResult, error) {
	authClient, err := adapters.NewAuthExternalClient()
	if err != nil {
		return adapters.AuthResult{}, nil
	}

	token, err := authClient.NewToken(cred)
	if err != nil {
		return adapters.AuthResult{}, err
	}

	return token, nil
}

func CreateUser(user adapters.User) error {
	authClient, err := adapters.NewAuthExternalClient()
	if err != nil {
		return nil
	}

	err = authClient.NewUser(user)
	if err != nil {
		return err
	}

	return nil
}

func handleCreateUser(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	var user adapters.User
	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       getJson(newError(err.Error())),
			Headers:    headers(),
		}, nil
	}

	err = CreateUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       getJson(newError(err.Error())),
			Headers:    headers(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "Created",
		Headers:    headers(),
	}, nil
}

func router(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	httpRequest := req.RequestContext.HTTP
	log.Printf("EVENT : %v", req)

	if strings.HasSuffix(httpRequest.Path, "/users") {
		if httpRequest.Method == "POST" {
			return handleCreateUser(req)
		}
	}

	if strings.HasSuffix(httpRequest.Path, "/users/token") {
		if httpRequest.Method == "POST" {
			return handleGetToken(req)
		}
	}
	log.Println(fmt.Sprintf("Test novo código %s ", time.Now().String()))
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
		Headers:    headers(),
	}, nil
}

func headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func main() {
	lambda.Start(router)
}
