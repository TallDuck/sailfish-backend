package jwt_test

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tallduck/sailfish-backend/auth/jwt"
)

type TokenHeader struct {
	Alg string
	Typ string
}

type TokenBody struct {
	Exp      int
	Username string
}

var testSecret = []byte("TESTING123")

func TestGenerate(t *testing.T) {
	jwt := jwt.New(testSecret)

	username := "billybob"

	token, _ := jwt.Generate(map[string]string{
		"username": username,
	})

	tokenParts := strings.Split(token, ".")

	header := tokenParts[0]
	body := tokenParts[1]

	headerBytes, _ := base64.RawURLEncoding.DecodeString(header)
	bodyBytes, _ := base64.RawURLEncoding.DecodeString(body)

	var tokenHeader TokenHeader
	var tokenBody TokenBody

	json.Unmarshal(headerBytes, &tokenHeader)
	json.Unmarshal(bodyBytes, &tokenBody)

	validate, _ := jwt.Validate(token)

	assert := assert.New(t)

	assert.Equal("HS256", tokenHeader.Alg, "Token algorithm should be HS256")
	assert.Equal("JWT", tokenHeader.Typ, "Token type should be JWT")

	assert.Equal(username, tokenBody.Username, "Token Body Username should have same username passed in")

	assert.True(validate, "Token should be valid")
}

func TestValidate(t *testing.T) {
	jwt := jwt.New(testSecret)

	token, _ := jwt.Generate(map[string]string{})
	tamper := token + "touch"

	validate, _ := jwt.Validate(token)
	invalidate, _ := jwt.Validate(tamper)

	assert := assert.New(t)

	assert.True(validate, "Untouched token should be valid")
	assert.False(invalidate, "Token that has been tampered with should be invalid")
}
