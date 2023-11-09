package adapters

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"reflect"
	"testing"
)

func TestNewAuthExternalClient(t *testing.T) {
	tests := []struct {
		name    string
		want    AuthExternalClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuthExternalClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuthExternalClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthExternalClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAwsClient(t *testing.T) {
	tests := []struct {
		name    string
		want    *session.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAwsClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAwsClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAwsClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authExternalClient_NewToken(t *testing.T) {
	type fields struct {
		Client *cognitoidentityprovider.CognitoIdentityProvider
	}
	type args struct {
		cred *Credentials
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    AuthResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := &authExternalClient{
				Client: tt.fields.Client,
			}
			got, err := auth.NewToken(tt.args.cred)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authExternalClient_NewUser(t *testing.T) {
	type fields struct {
		Client *cognitoidentityprovider.CognitoIdentityProvider
	}
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := &authExternalClient{
				Client: tt.fields.Client,
			}
			if err := auth.NewUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
