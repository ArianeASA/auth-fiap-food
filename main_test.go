package main

import (
	"auth-fiap-food/adapters"
	"github.com/aws/aws-lambda-go/events"
	"reflect"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		user adapters.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetToken(t *testing.T) {
	type args struct {
		cred *adapters.Credentials
	}
	tests := []struct {
		name    string
		args    args
		want    adapters.AuthResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetToken(tt.args.cred)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleCreateUser(t *testing.T) {
	type args struct {
		req events.APIGatewayV2HTTPRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handleCreateUser(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleCreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleCreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleGetToken(t *testing.T) {
	type args struct {
		req events.APIGatewayV2HTTPRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handleGetToken(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleGetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleGetToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headers(t *testing.T) {
	tests := []struct {
		name string
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := headers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("headers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_router(t *testing.T) {
	type args struct {
		req events.APIGatewayV2HTTPRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := router(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("router() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("router() got = %v, want %v", got, tt.want)
			}
		})
	}
}
