package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	st := NewStack[int]()
	assert.NotNil(t, st)
	assert.True(t, st.IsEmpty())
}

func TestStackPush(t *testing.T) {
	st := NewStack[int]()
	st.Push(10)
	assert.False(t, st.IsEmpty())
	assert.Equal(t, 10, st.Top())
}

func TestStackPop(t *testing.T) {
	st := NewStack[int]()
	st.Push(20)
	st.Push(30)

	err := st.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 20, st.Top())
	assert.False(t, st.IsEmpty())

	err = st.Pop()
	assert.NoError(t, err)
	assert.True(t, st.IsEmpty())

	err = st.Pop()
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid operation: empty stack")
}

func TestStackTop(t *testing.T) {
	st := NewStack[int]()
	st.Push(40)
	assert.Equal(t, 40, st.Top())

	st.Push(50)
	assert.Equal(t, 50, st.Top())

	err := st.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 40, st.Top())
}

func TestStackIsEmpty(t *testing.T) {
	st := NewStack[int]()
	assert.True(t, st.IsEmpty())

	st.Push(60)
	assert.False(t, st.IsEmpty())

	err := st.Pop()
	assert.NoError(t, err)
	assert.True(t, st.IsEmpty())
}
