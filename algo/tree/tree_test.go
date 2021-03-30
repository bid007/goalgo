package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeserialize(t *testing.T) {
	assert.Equal(
		t, "3,1,<nil>,0,<nil>,<nil>,2,<nil>,<nil>",
		Serialize("3,1,<nil>,0,<nil>,<nil>,2,<nil>,<nil>").Deserialize())
}
