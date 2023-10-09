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

	ClearCommand struct {
		cmd
	}

	LogCommand struct {
		cmd
		Entry LogEntry
	}
)

func (c *cmd) command() {}
