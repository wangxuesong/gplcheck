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

	SourceCommand struct {
		cmd
		Source string
	}

	LogCommand struct {
		cmd
		Entry LogEntry
	}

	StatusCommand struct {
		cmd
		Status string
	}

	ProgressStartCommand struct {
		cmd
		FileName string
		Total    int
	}

	ProgressUpdateCommand struct {
		cmd
		Progress int
		Total    int
	}

	ProgressEndCommand struct {
		cmd
	}
)

func (c *cmd) command() {}
