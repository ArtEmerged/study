package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data, Next: l.Head}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Head = n
	}
}
