package jwtfiltergolang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceContains(t *testing.T) {

	validActions := []string{"create", "read","update","delete"}

	assert.Equal(t,true,SliceContains(validActions,"create"))
	assert.Equal(t,false,SliceContains(validActions,"view"))

}
