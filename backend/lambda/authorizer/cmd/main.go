package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	my_cognito "github.com/ujjwal405/FileSharing/authorizer/cognito"
	my_db "github.com/ujjwal405/FileSharing/authorizer/dynamodb"

	"github.com/ujjwal405/FileSharing/authorizer/handler"
	"github.com/ujjwal405/FileSharing/authorizer/verifier"
)

var lambdaHandler *handler.AuthorizerHandler

func init() {

	cClient, err := my_cognito.NewCognitoClient()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}
	dClient, err := my_db.NewDynamoClient()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}
	verifier, err := verifier.NewVerifier()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}

	lambdaHandler = handler.NewLambdaHandler(cClient, verifier, dClient)
}

func handleAuthorize(ctx context.Context, event events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	authHeader := event.Headers[http.CanonicalHeaderKey("authorization")]
	idToken := event.Headers[http.CanonicalHeaderKey("x-id-token")] // ID Token is passed directly

	// Split Authorization header by space
	parts := strings.Split(authHeader, " ")
	if !(len(parts) == 2 && strings.ToLower(parts[0]) == "bearer") {
		return generatePolicyWithContext("Deny", event.MethodArn, nil), nil
	}

	newToken, err := lambdaHandler.Authorize(ctx, parts[1], idToken)
	if err != nil {
		log.Printf("failed to handle authorization: %v", err)
		return generatePolicyWithContext("Deny", event.MethodArn, newToken), nil
	}
	return generatePolicyWithContext("Allow", event.MethodArn, newToken), nil
}

func generatePolicyWithContext(effect, resource string, claims *handler.NewToken) events.APIGatewayCustomAuthorizerResponse {
	policy := events.APIGatewayCustomAuthorizerResponse{
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		},
	}
	if effect == "Allow" && claims != nil {
		contextMap := map[string]interface{}{
			"username": claims.Username,
		}

		if claims.AccessToken != "" {
			contextMap["access_token"] = claims.AccessToken
		}
		if claims.IdToken != "" {
			contextMap["id_token"] = claims.IdToken
		}
		if claims.GoogleLogin != "" {
			contextMap["google_login"] = claims.GoogleLogin
		}

		policy.Context = contextMap
	}

	return policy
}

func main() {
	lambda.Start(handleAuthorize)
}
