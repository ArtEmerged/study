package piscine

func ListSize(l *List) int {
	var res int
	for l.Head != nil {
		res++
		l.Head = l.Head.Next
	}
	return res
}
