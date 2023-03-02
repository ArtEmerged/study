package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for v := 0; v < len(table)-1; v++ {
			if table[v] > table[v+1] {
				table[v], table[v+1] = table[v+1], table[v]
			}
		}
	}
}
