package adapters

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"log"
	"os"
)

type authExternalClient struct {
	Client *cognitoidentityprovider.CognitoIdentityProvider
}

type AuthExternalClient interface {
	NewUser(user User) error
	NewToken(cred *Credentials) (AuthResult, error)
}

func NewAuthExternalClient() (AuthExternalClient, error) {
	sess, err := NewAwsClient()
	if err != nil {
		return nil, err
	}

	cognitoClient := cognitoidentityprovider.New(sess)
	return &authExternalClient{Client: cognitoClient}, nil
}

func (auth *authExternalClient) NewUser(user User) error {
	input := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(os.Getenv("CLIENT_ID")),
		Password: aws.String(user.Password),
		Username: aws.String(user.CPF),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			}, {
				Name:  aws.String("name"),
				Value: aws.String(user.Name),
			},
		},
	}

	result, err := auth.Client.SignUp(input)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(result.String())
	return nil

}

func (auth *authExternalClient) NewToken(cred *Credentials) (AuthResult, error) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(cred.CPF),
			"PASSWORD": aws.String(cred.Password),
		},

		ClientId: aws.String(os.Getenv("CLIENT_ID")),
	}

	result, err := auth.Client.InitiateAuth(input)
	if err != nil {
		log.Println(err)
		return AuthResult{}, err

	}

	log.Println(result.AuthenticationResult.AccessToken)
	return AuthResult{
		AccessToken: result.AuthenticationResult.AccessToken,
		ExpiresIn:   result.AuthenticationResult.ExpiresIn,
		TokenType:   result.AuthenticationResult.TokenType,
	}, nil
}
