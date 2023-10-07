package common

type Notifier struct {
	refreshChan chan interface{}
	closeChan   chan interface{}
}

func NewNotifier() *Notifier {
	return &Notifier{
		refreshChan: make(chan interface{}),
		closeChan:   make(chan interface{}),
	}
}

func (n *Notifier) RefreshChan() chan interface{} {
	return n.refreshChan
}

func (n *Notifier) CloseChan() chan interface{} {
	return n.closeChan
}
