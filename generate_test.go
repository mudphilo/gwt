package jwtfiltergolang

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestCreateToken(t *testing.T) {

	os.Setenv("DARIDE_JWT_SECRET","eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTY2MzAxODY0NSwiaWF0IjoxNjYzMDE4NjQ1fQ.q9SwFW4jkhSpQKupbFOZVwdzQKnnsI73BZJZT-lDr1E")
	os.Setenv("DARIDE_JWT_ISSUER","da-ride.com")
	os.Setenv("DARIDE_JWT_DURATION_HOURS","72")

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

	log.Printf(auth)
}

func TestCreateTokenWithTenant(t *testing.T) {

	_ = os.Setenv("DARIDE_JWT_SECRET", "eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTY2MzAxODY0NSwiaWF0IjoxNjYzMDE4NjQ1fQ.q9SwFW4jkhSpQKupbFOZVwdzQKnnsI73BZJZT-lDr1E")
	_ = os.Setenv("DARIDE_JWT_ISSUER", "da-ride.com")
	os.Setenv("DARIDE_JWT_DURATION_HOURS","72")

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

	auth, err := CreateTokenWithClient(1,"tenant",1,123, 1, "kamau", Role{
		Name:       "admin",
		Permission: permissions,
	})
	assert.NoError(t, err)

	if len(auth) == 0 {

		t.Fatalf(`Expeted token  = %s, Error %v`, auth, err)
	}

	log.Printf(auth)
}
