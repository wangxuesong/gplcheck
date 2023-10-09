package common

type (
	Command interface {
		command()
	}

	cmd struct{}

	ParseCommand struct {
		cmd
		FilePath string
	}

	LogCommand struct {
		cmd
		Entry LogEntry
	}
)

func (c *cmd) command() {}
