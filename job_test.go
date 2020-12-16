package jobs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJob_Body(t *testing.T) {
	j := &Job{Payload: "hello"}

	assert.Equal(t, []byte("hello"), j.Body())
}

func TestJob_Context(t *testing.T) {
	j := &Job{Job: "job"}

	assert.Equal(t, []byte(`{"id":"id","job":"job","headers":null}`), j.Context("id"))
}

func TestJob_Context_Headers(t *testing.T) {
	j := &Job{Job: "job", Headers: map[string]string{"test": "test"}}

	assert.Equal(t, []byte(`{"id":"id","job":"job","headers":{"test":"test"}}`), j.Context("id"))
}
