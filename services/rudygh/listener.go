package main

type Listener struct{}

func (l *Listener) Auth(line []byte, ack *bool) (err error) {
	return
}
