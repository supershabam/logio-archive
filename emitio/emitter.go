package emitio

import "io"
import "time"

func NewEmitter(service []string, opts ...EmitterOption) (*StdEmitter, error) {
	return &StdEmitter{
		service: service,
	}, nil
}

type emitterOption struct {
	out io.Writer
}

type EmitterOption func(*emitterOption)

func WithOutput(out io.Writer) EmitterOption {
	return func(option *emitterOption) {
		option.out = out
	}
}

type Event interface {
	Start() time.Time
	Duration() time.Duration
	Name() string
	Baggage() map[string]string
	AddCause(Cause)
	Carrier() Carrier
	Severity()
}

type Carrier interface {
	Event() event
}

type valuer interface{}

type causer interface{}

type event struct {
	Name     string            // who
	Fields   map[string]valuer // what
	Start    time.Time         // when
	Duration time.Duration
	Causes   []causer // why
}

func (e *event) Induce() Carrier {}

type Severity int

type Emitter interface {
	Emit(Severity, Event)
}

type StdEmitter struct {
	service []string
}
