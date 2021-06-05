package middleware

import (
	"errors"
	"fmt"
	"log"

	jwthelper "stock-contexts/pkg/shared/utils/jwt"
	rsahelper "stock-contexts/pkg/shared/utils/rsa"

	"github.com/dgrijalva/jwt-go"
	routing "github.com/qiangxue/fasthttp-routing"
)

// Authorizing define jwt middleware
// func Authorizing() func(c *routing.Context) error {
// 	// maybe calling other authorizing server
// 	// TODO
// 	// after authorizing succeed
// 	return jwthelper.Sign()
// }

// Authenticating define a method
func Authenticating() func(c *routing.Context) error {
	return jwthelper.Verify("secret-key", jwthelper.VerifyOptions{
		Realm:         "lazy-stock-screener.com",
		SigningMethod: "RS256",
		ValidationKeyGetter: func() func(t *jwt.Token) (interface{}, error) {
			// extract public key from pem
			return func(t *jwt.Token) (interface{}, error) {
				verifyingKey, err := rsahelper.GetPublicPemCert()
				if err != nil {
					log.Fatalf("verifying key error: %s", err.Error())
				}
				// parse rsa public key from pem file
				key, err := jwt.ParseRSAPublicKeyFromPEM(verifyingKey)
				if err != nil {
					log.Fatalf("parsing RSAPublicKeyFromPEM error: %s", err.Error())
				}
				return key, nil
			}
		},
		TokenHandler: func(c *routing.Context, token *jwt.Token) error {
			// Verify 'aud' claim
			aud := "84932737.lazy-stock-screener.com"
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAud {
				return errors.New("invalid audience")
			}
			// Verify 'iss' claim
			iss := "https://lazy-stock-screener.com/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return errors.New("invalid issuer")
			}
			// Attach JWT token on header

			return nil
		},
	})
}

func Restricted() func(c *routing.Context) error {
	return func(c *routing.Context) error {
		// claims := c.Get("JWT").(*jwt.Token).Claims.(jwt.MapClaims)
		c.Write([]byte(fmt.Sprintf("Welcome, ID: %v!")))
		return nil
	}
}
