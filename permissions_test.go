package jwtfiltergolang

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// CreateToken  creates a token , checking
// for a valid return value.

func TestHasPermission(t *testing.T) {

	os.Setenv("JWT_SECRET","A10FD19C79E57532D6AEA472E42FE1D4FC509C3E0F6EEE71C4B70E82FICA432F")
	os.Setenv("JWT_ISSUER","da-ride.com")
	os.Setenv("JWT_DURATION_HOURS","72")

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

	// check permissions

	check := HasPermission(auth,"user","read","all")
	assert.Equal(t, true, check)

	check = HasPermission(auth,"driver","create","all")
	assert.Equal(t, false, check)

}
