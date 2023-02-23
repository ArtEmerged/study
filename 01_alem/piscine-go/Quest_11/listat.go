package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	answer := l
	for i := 0; i < pos; i++ {
		if answer.Next == nil {
			return nil
		}
		answer = answer.Next
	}
	return answer
}
