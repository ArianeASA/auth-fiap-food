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

func handleGetToken(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	credential := &adapters.Credentials{}
	err := json.Unmarshal([]byte(req.Body), &credential)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
			//Headers:    headers(),
		}, nil
	}
	token, err := GetToken(credential)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
			//Headers:    headers(),
		}, nil
	}

	obj, err := json.Marshal(token)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(obj),
		//Headers:    headers(),
	}, nil

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
			Body:       err.Error(),
			Headers:    headers(),
		}, nil
	}

	err = CreateUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
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
		//Headers:    headers(),
	}, nil
}

func headers() map[string]string {
	connType := make(map[string]string)
	connType["Content-Type"] = "application/json"
	return connType
}

func main() {
	lambda.Start(router)
}
