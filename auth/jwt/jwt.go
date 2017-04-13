package jwt

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWT contains the signingSecret and utility methods
type JWT struct {
	signingSecret []byte
}

// New will return a JWT struct with the signingSecret populated, if not passed
// in it will load the secret from the environment (SAILFISH_AUTH_SECRET).
func New(args ...[]byte) JWT {
	secret := args[0]

	if secret != nil {
		return JWT{
			signingSecret: secret,
		}
	}

	return JWT{
		signingSecret: []byte("THIS_SECRET_SHOULD_BE_LOADED_FROM_THE_ENV"),
	}
}

// Generate will generate a signed JWT token continating the data passed in.
func (j *JWT) Generate(data map[string]string) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Duration(120 * 60)).UTC().Unix(),
	}

	for k, v := range data {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.signingSecret)
}

// Validate will verify that a JWT token is valid and has not been tampered with.
func (j *JWT) Validate(tokenStr string) (bool, interface{}) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return j.signingSecret, nil
	})

	if err != nil {
		fmt.Println("Unable to parse JWT")
		return false, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims
	}

	return false, false
}
