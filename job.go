package jobs

import (
	protobuf "github.com/golang/protobuf/proto"
	json "github.com/json-iterator/go"
	proto "github.com/spiral/jobs/v2/proto"
)

// Handler handles job execution.
type Handler func(id string, j *Job) error

// ErrorHandler handles job execution errors.
type ErrorHandler func(id string, j *Job, err error)

// Job carries information about single job.
type Job struct {
	// Job contains name of job broker (usually PHP class).
	Job string `json:"job"`

	// Payload is string data (usually JSON) passed to Job broker.
	Payload string `json:"payload"`

	// Options contains set of PipelineOptions specific to job execution. Can be empty.
	Options *Options `json:"options,omitempty"`

	// Job Headers
	Headers map[string]string
}

// Body packs job payload into binary payload.
func (j *Job) Body() []byte {
	return []byte(j.Payload)
}

// Context packs job context (job, id) into binary payload.
func (j *Job) Context(id string) []byte {
	ctx, _ := json.Marshal(
		struct {
			ID      string            `json:"id"`
			Job     string            `json:"job"`
			Headers map[string]string `json:"headers"`
		}{ID: id, Job: j.Job, Headers: j.Headers},
	)

	return ctx
}

func (j *Job) ProtoUnmarshal(data []byte) (err error) {
	pJob := &proto.Job{}
	if err = protobuf.Unmarshal(data, pJob); err != nil {
		return err
	}

	j.Job = pJob.GetJob()
	j.Payload = string(pJob.GetPayload())
	j.Headers = pJob.GetHeaders()

	pOpt := pJob.GetOptions()
	if pOpt != nil {
		j.Options = &Options{}

		if pOpt.GetAttempts() != 0 {
			j.Options.Attempts = int(pOpt.GetAttempts())
		}

		if pOpt.GetDelay() != 0 {
			j.Options.Delay = int(pOpt.GetDelay())
		}

		if pOpt.GetPipeline() != "" {
			j.Options.Pipeline = pOpt.GetPipeline()
		}

		if pOpt.GetRetryDelay() != 0 {
			j.Options.RetryDelay = int(pOpt.GetRetryDelay())
		}

		if pOpt.GetTimeout() != 0 {
			j.Options.Timeout = int(pOpt.GetTimeout())
		}
	}

	return nil
}
