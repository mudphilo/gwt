package jwtfiltergolang

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestTokenValidation(t *testing.T) {

	_ = os.Setenv("JWT_SECRET", "jWnZr4u7x!A%C*F-JaNdRgUkXp2s5v8y")
	_ = os.Setenv("JWT_ISSUER", "smestech.com")
	_ = os.Setenv("JWT_DURATION_HOURS", "72")

	var permissions []Permission
	permissions = append(permissions, Permission{
		Module:  "user",
		Scope: "all",
		Actions: []string{"create","read"},
	})

	permissions = append(permissions, Permission{
		Module:  "driver",
		Scope: "all",
		Actions: []string{"read"},
	})

	auth, err := CreateToken(123, 1, "kamau", Role{
		Name:       "admin",
		Permission: permissions,
	})

	assert.NoError(t, err)

	if len(auth) == 0 {

		t.Fatalf(`Expeted token  = %s, Error %v`, auth, err)
	}

	log.Printf("got JWT Token %s",auth)

	token, err := TokenValidation(auth)
	assert.NoError(t, err)

	assert.Equal(t, "kamau", token.Username)
	assert.Equal(t, int64(123), token.UserId)

}

func TestTokenValidationWrongTime(t *testing.T) {

	_ = os.Setenv("JWT_SECRET", "eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTY2MzAxODY0NSwiaWF0IjoxNjYzMDE4NjQ1fQ.q9SwFW4jkhSpQKupbFOZVwdzQKnnsI73BZJZT-lDr1E")
	_ = os.Setenv("JWT_ISSUER", "smestech.com")
	_ = os.Setenv("JWT_DURATION_HOURS", "a")

	var permissions []Permission
	permissions = append(permissions, Permission{
		Module:  "user",
		Actions: []string{"create","read"},
	})

	permissions = append(permissions, Permission{
		Module:  "driver",
		Actions: []string{"read"},
	})

	_, err := CreateToken(123, 1, "kamau", Role{
		Name:       "admin",
		Permission: permissions,
	})

	if err == nil {

		t.Fatalf("Expeted error got nil ")
	}

}

func TestTokenValidationWrongToken(t *testing.T) {

	os.Setenv("JWT_SECRET","eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTY2MzAxODY0NSwiaWF0IjoxNjYzMDE4NjQ1fQ.q9SwFW4jkhSpQKupbFOZVwdzQKnnsI73BZJZT-lDr1E")
	os.Setenv("JWT_ISSUER","smestech.com")
	os.Setenv("JWT_DURATION_HOURS","a")

	_, err := TokenValidation("wrongtoken here")
	if err == nil {

		t.Fatalf("Expeted error got nil ")
	}

}

func TestTokenValidationExpiredToken(t *testing.T) {

	os.Setenv("JWT_SECRET","eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTY2MzAxODY0NSwiaWF0IjoxNjYzMDE4NjQ1fQ.q9SwFW4jkhSpQKupbFOZVwdzQKnnsI73BZJZT-lDr1E")
	os.Setenv("JWT_ISSUER","smestech.com")
	os.Setenv("JWT_DURATION_HOURS","-10")

	var permissions []Permission
	permissions = append(permissions, Permission{
		Module:  "user",
		Actions: []string{"create","read"},
	})

	permissions = append(permissions, Permission{
		Module:  "driver",
		Actions: []string{"read"},
	})

	auth, err := CreateToken(123, 1, "kamau", Role{
		Name:       "admin",
		Permission: permissions,
	})

	assert.NoError(t,err)

	_, err = TokenValidation(auth)
	if err == nil {

		t.Fatalf("Expeted error got nil ")
	}

}

func TestValidaJavaToken(t *testing.T)  {

	os.Setenv("JWT_SECRET","dSgVkYp3s5v8y/B?E(H+MbQeThWmZq4t")
	os.Setenv("JWT_ISSUER","smestech.com")
	os.Setenv("JWT_DURATION_HOURS","10")

	javaToken := "eyJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE2NjgxODczMDgsImlzcyI6InNpdGUuY29tIiwidXNlcm5hbWUiOiJqYXZhIiwidXNlcl9pZCI6MSwidXNlcl9zdGF0dXMiOjEsInJvbGUiOnsibmFtZSI6ImFkbWluIiwicGVybWlzc2lvbiI6W3sibW9kdWxlIjoidXNlciIsInNjb3BlIjoiYWxsIiwiYWN0aW9ucyI6WyJjcmVhdGUiLCJyZWFkIiwidXBkYXRlIiwiZGVsZXRlIl19LHsibW9kdWxlIjoidHJpcCIsInNjb3BlIjoiYWxsIiwiYWN0aW9ucyI6WyJjcmVhdGUiLCJyZWFkIiwidXBkYXRlIiwiZGVsZXRlIl19LHsibW9kdWxlIjoicGF5bWVudCIsInNjb3BlIjoiYWxsIiwiYWN0aW9ucyI6WyJyZWFkIl19XX0sInRlbmFudF9pZCI6MSwiY2xpZW50X2lkIjoxLCJleHAiOjE2NjgyMjMzMDh9.IdssaLcy4-d6EkO6OFlWgZJ3a4XyEu8b30qVAslhPEk"

	token, err := TokenValidation(javaToken)
	assert.NoError(t, err)

	assert.Equal(t, "java", token.Username)
	assert.Equal(t, int64(1), token.UserId)

}