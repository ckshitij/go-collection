package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	assert.NotNil(t, q)
	assert.True(t, q.IsEmpty())
}

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(10)
	assert.False(t, q.IsEmpty())
	assert.Equal(t, 10, q.Front())
	assert.Equal(t, 10, q.Back())

	q.Enqueue(20)
	assert.Equal(t, 10, q.Front())
	assert.Equal(t, 20, q.Back())
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)

	err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 20, q.Front())
	assert.False(t, q.IsEmpty())

	err = q.Dequeue()
	assert.NoError(t, err)
	assert.True(t, q.IsEmpty())

	err = q.Dequeue()
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid operation: empty Queue")
}

func TestQueueFront(t *testing.T) {
	q := NewQueue[int]()
	assert.Equal(t, 0, q.Front()) // Default zero value for int when queue is empty

	q.Enqueue(30)
	assert.Equal(t, 30, q.Front())

	q.Enqueue(40)
	assert.Equal(t, 30, q.Front()) // Front should remain the same

	q.Dequeue()
	assert.Equal(t, 40, q.Front()) // Front should update after dequeue
}

func TestQueueBack(t *testing.T) {
	q := NewQueue[int]()
	assert.Equal(t, 0, q.Back()) // Default zero value for int when queue is empty

	q.Enqueue(50)
	assert.Equal(t, 50, q.Back())

	q.Enqueue(60)
	assert.Equal(t, 60, q.Back())

	q.Dequeue()
	assert.Equal(t, 60, q.Back()) // Back should not change after dequeue
}
