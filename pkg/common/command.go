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
)

func (c *cmd) command() {}
