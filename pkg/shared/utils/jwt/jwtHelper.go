package jwthelper

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// JWTTokenHandler handles the parsed JWT token.
type JWTTokenHandler func(*routing.Context, *jwt.Token) error

// VerificationKeyHandler gets a dynamic VerificationKey.
type VerificationKeyHandler func() func(t *jwt.Token) (interface{}, error)

// VerifyOptions represents the options that can be used with the JWT handler.
type VerifyOptions struct {
	Realm               string
	SigningMethod       string
	TokenHandler        JWTTokenHandler
	ValidationKeyGetter VerificationKeyHandler
}

func Encrypt() {
	// rcpt := jose.Recipient{
	// 	Algorithm:  jose.PBES2_HS256_A128KW,
	// 	Key:        "mypassphrase",
	// 	PBES2Count: 4096,
	// 	PBES2Salt:  []byte{"123123123"},
	// }
}

// func Sign() routing.Handler {
// 	return func(ctx *routing.Context) error {
// 		signingKey, err := rsahelper.GetPrivatePemCert()
// 		if err != nil {
// 			return fmt.Errorf("Get signing key error: %v", err)
// 		}
// 		key, err := jwt.ParseRSAPrivateKeyFromPEM(signingKey)
// 		if err != nil {
// 			return fmt.Errorf("Get key error: %v", err)
// 		}
// 		tokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
// 			"id":  "222",
// 			"aud": "84932737.lazy-stock-screener.com",
// 			"iss": "https://lazy-stock-screener.com/",
// 		}).SignedString(key)
// 		if err != nil {
// 			return fmt.Errorf("Transform to token string error: %v", err)
// 		}
// 		// ctx.Response.Header.WriteTo(os.Stdout)
// 		ctx.Response.Header.Set("Authorization", "Bearer "+tokenString)
// 		ctx.Write([]byte(tokenString))
// 		return nil
// 	}
// }

func getJWTKeyFunc(opt VerifyOptions, verificationKey interface{}) jwt.Keyfunc {
	if opt.ValidationKeyGetter == nil {
		return func(t *jwt.Token) (interface{}, error) {
			return verificationKey, nil
		}
	}
	return opt.ValidationKeyGetter()
}

func Verify(verificationKey interface{}, options ...VerifyOptions) routing.Handler {
	var opt VerifyOptions
	var keyFunc jwt.Keyfunc
	var message string
	if len(options) > 0 {
		opt = options[0]
	}
	keyFunc = getJWTKeyFunc(opt, verificationKey)
	parser := &jwt.Parser{
		ValidMethods: []string{opt.SigningMethod},
	}
	return func(ctx *routing.Context) error {
		// ctx.Request.Header.WriteTo(os.Stdout)
		header := string(ctx.Request.Header.Peek("Authorization"))
		if strings.HasPrefix(header, "Bearer ") {
			if token, err := parser.Parse(header[7:], keyFunc); err == nil && token.Valid {
				if err = opt.TokenHandler(ctx, token); err != nil {
					message = err.Error()
				}
				return nil
			}
		}
		ctx.Response.Header.Set("WWW-Authenticate", `Bearer realm="`+opt.Realm+`"`)
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		ctx.SetBodyString(message)
		ctx.Abort()
		return nil
	}
}
