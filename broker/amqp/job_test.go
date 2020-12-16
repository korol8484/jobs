package amqp

import (
	"testing"

	"github.com/spiral/jobs/v2"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func Test_Unpack_Errors(t *testing.T) {
	_, _, _, err := unpack(amqp.Delivery{
		Headers: map[string]interface{}{},
	})
	assert.Error(t, err)

	_, _, _, err = unpack(amqp.Delivery{
		Headers: map[string]interface{}{
			"rr-id": "id",
		},
	})
	assert.Error(t, err)

	_, _, _, err = unpack(amqp.Delivery{
		Headers: map[string]interface{}{
			"rr-id":      "id",
			"rr-attempt": int64(0),
		},
	})
	assert.Error(t, err)
}

func Test_Unpack_Headers(t *testing.T) {
	table := amqp.Delivery{
		Headers: amqp.Table{
			"rr-id":      "id",
			"rr-job":     "job",
			"rr-attempt": int64(1),
			"test":       "test",
			"test1":      "test1",
		},
	}

	_, _, j, err := unpack(table)

	assert.NoError(t, err)
	assert.Equal(t, []byte(`{"id":"id","job":"job","headers":{"test":"test","test1":"test1"}}`), j.Context("id"))
}

func Test_Pack_Headers(t *testing.T) {
	id := "id"
	attempt := 1

	j := &jobs.Job{
		Job:     "job",
		Payload: "data",
		Options: &jobs.Options{
			Pipeline:   "",
			Delay:      0,
			Attempts:   0,
			RetryDelay: 0,
			Timeout:    0,
		},
		Headers: map[string]string{
			"test":  "test",
			"test1": "test1",
		},
	}

	table := amqp.Table{
		"rr-id":          id,
		"rr-job":         j.Job,
		"rr-attempt":     int64(attempt),
		"rr-maxAttempts": int64(j.Options.Attempts),
		"rr-timeout":     int64(j.Options.Timeout),
		"rr-delay":       int64(j.Options.Delay),
		"rr-retryDelay":  int64(j.Options.RetryDelay),
		"test":           "test",
		"test1":          "test1",
	}

	testTable := pack(id, attempt, j)
	assert.Equal(t, testTable, table)
}
