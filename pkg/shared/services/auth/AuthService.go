package authservice

import (
	"crypto/rand"
	"fmt"
	rsahelper "stock-contexts/pkg/shared/utils/rsa"

	"github.com/dgrijalva/jwt-go"
)

var jWTAuthSigton *JWTAuth = nil

// JWTAuth define repo struct
type JWTAuth struct {
	Realm         string
	SigningMethod string
	JWTParser     *jwt.Parser
}

type MapClaimsDTO struct {
	Id  string
	Aud string
	Iss string
}

func (j *JWTAuth) CreateRefreshToken() string {
	b := make([]byte, 256)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// SignJWT define a sign method
func (j *JWTAuth) SignJWT(userName string) (string, error) {
	signingKey, err := rsahelper.GetPrivatePemCert()
	if err != nil {
		return "", fmt.Errorf("Get signing key error: %v", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(signingKey)
	if err != nil {
		return "", fmt.Errorf("Get key error: %v", err)
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":  userName,
		"aud": "84932737.lazy-stock-screener.com",
		"iss": "https://lazy-stock-screener.com/",
	}).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Transform to token string error: %v", err)
	}
	return tokenString, nil
}

func (j *JWTAuth) VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := j.JWTParser.Parse(tokenString[7:], j.createJWTKey())
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Parse Key Error: %v", err)
	}

	token, err = j.TokenHandler(token)
	if err != nil {
		return nil, fmt.Errorf("Key Content Error: %v", err)
	}
	return token, nil
}

func (j *JWTAuth) createJWTKey() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		verifyingKey, err := rsahelper.GetPublicPemCert()
		if err != nil {
			return nil, fmt.Errorf("Get public key error: %v", err)
		}
		key, err := jwt.ParseRSAPublicKeyFromPEM(verifyingKey)
		if err != nil {
			return nil, fmt.Errorf("Parse public key error: %v", err)
		}
		return key, nil
	}
}

func (j *JWTAuth) TokenHandler(token *jwt.Token) (*jwt.Token, error) {
	// Verify 'aud' claim
	aud := "84932737.lazy-stock-screener.com"
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
	if !checkAud {
		return nil, fmt.Errorf("invalid audience")
	}
	// Verify 'iss' claim
	iss := "https://lazy-stock-screener.com/"
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
	if !checkIss {
		return nil, fmt.Errorf("invalid issuer")
	}
	return token, nil
}

// NewJWTAuth define
func NewJWTAuth() *JWTAuth {
	if jWTAuthSigton == nil {
		auth := &JWTAuth{
			Realm: "lazy-stock-screener.com",
			// Aud:           "lazy-stock-screener-identity",
			SigningMethod: "RS256",
			JWTParser: &jwt.Parser{
				ValidMethods: []string{"RS256"},
			},
		}
		jWTAuthSigton = auth
		return auth
	}
	return jWTAuthSigton
}
