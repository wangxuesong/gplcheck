package common

type Notifier struct {
	refreshChan chan interface{}
	closeChan   chan interface{}
	logChan     chan Command
	commandChan chan Command
}

func NewNotifier() *Notifier {
	return &Notifier{
		refreshChan: make(chan interface{}),
		closeChan:   make(chan interface{}),
		logChan:     make(chan Command),
		commandChan: make(chan Command),
	}
}

func (n *Notifier) RefreshChan() chan interface{} {
	return n.refreshChan
}

func (n *Notifier) CloseChan() chan interface{} {
	return n.closeChan
}

func (n *Notifier) LogChan() chan Command {
	return n.logChan
}

func (n *Notifier) CommandChan() chan Command {
	return n.commandChan
}
