package auth

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

// ctxKey type would be used to put the claims in the context
type ctxKey int

const Key ctxKey = 1

// Auth struct privateKey field would be used to verify and generate token
type Auth struct {
	privateKey *rsa.PrivateKey
}

// NewAuth func set the privateKey in the Auth struct and returns the instance of it to the caller
func NewAuth(privateKey *rsa.PrivateKey) (*Auth, error) {

	if privateKey == nil {
		return nil, errors.New("private key cannot be nil")
	}

	a := Auth{
		privateKey: privateKey,
	}
	return &a, nil

}

func (a *Auth) GenerateToken(claims jwt.RegisteredClaims) (string, error) {

	//jwt.NewWithClaims takes a signingMethod and claims struct to generate a token on basis of it
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, &claims)

	//signing our token with our private key
	tokenStr, err := tkn.SignedString(a.privateKey)

	if err != nil {
		return "", fmt.Errorf("signing token %w", err)
	}

	return tokenStr, nil

}

// ValidateToken check whether token is valid or not
func (a *Auth) ValidateToken(tokenStr string) (jwt.RegisteredClaims, error) {

	var claims jwt.RegisteredClaims

	// verifying token with our public key // if token is valid, we fetch the data stored inside the token and put the data in claims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return a.privateKey.Public(), nil
	})

	if err != nil {
		return jwt.RegisteredClaims{}, fmt.Errorf("parsing token %w", err)
	}

	if !token.Valid {
		return jwt.RegisteredClaims{}, errors.New("invalid token")
	}

	//returning Claims if verification is successful
	return claims, nil

}

// ValidateSession method validate token by taking it out of the context
func (a *Auth) ValidateSession(ctx context.Context) (jwt.RegisteredClaims, error) {

	//taking the jwt token value out of the context
	tkn, ok := ctx.Value(Key).(string) //type assertion // making sure the value is of string type
	if !ok || tkn == "" {

		//jwt.RegisteredClaims contains all the info about token
		return jwt.RegisteredClaims{}, errors.New("token value not found")
	}

	// validating the token and if token is valid then we put the jwt payload in claims
	claims, err := a.ValidateToken(tkn)

	if err != nil {
		return jwt.RegisteredClaims{}, fmt.Errorf("invalid session %w", err)
	}
	return claims, nil
}

//jwt.RegisteredClaims{}
// iss (issuer): Issuer of the JWT
// sub (subject): Subject of the JWT (the users)
// aud (audience): Recipient for which the JWT is intended
// exp (expiration time): Time after which the JWT expires
// nbf (not before time): Time before which the JWT must not be accepted for processing
// iat (issued at time): Time at which the JWT was issued; can be used to determine age of the JWT
// jti (JWT ID): Unique identifier; can be used to prevent the JWT from being replayed (allows a token to be used only once)
