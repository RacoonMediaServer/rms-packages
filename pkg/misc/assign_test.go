package misc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testAssignCase struct {
	Name  string
	Count int

	mutex int
}

func TestAssignFields(t *testing.T) {
	src := testAssignCase{
		Name:  "Test",
		Count: 133,
		mutex: 18,
	}
	dst := testAssignCase{}

	AssignFields(&src, &dst)

	assert.Equal(t, "Test", dst.Name)
	assert.Equal(t, 133, dst.Count)
	assert.Equal(t, 0, dst.mutex)
}
