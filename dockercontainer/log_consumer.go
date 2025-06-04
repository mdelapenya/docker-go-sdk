package dockercontainer

const (
	// StdoutLog is the log type for STDOUT
	StdoutLog = "STDOUT"

	// StderrLog is the log type for STDERR
	StderrLog = "STDERR"
)

// LogProductionOption is a function that modifies a [Container].
type LogProductionOption func(*Container)

// Log represents a message that was created by a process,
// LogType is either "STDOUT" or "STDERR",
// Content is the byte contents of the message itself
type Log struct {
	LogType string
	Content []byte
}

// LogConsumer represents any object that can handle a Log.
// It is up to the LogConsumer instance what to do with the log.
type LogConsumer interface {
	Accept(Log)
}

// LogConsumerConfig is a configuration object for the producer/consumer pattern
type LogConsumerConfig struct {
	Opts      []LogProductionOption // options for the production of logs
	Consumers []LogConsumer         // consumers for the logs
}
