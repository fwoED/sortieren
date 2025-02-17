package bubblesort

// BubbleSort sortiert the gegebene Liste mit dem Bubble-Sort-Algorithmus.
func BubbleSort(list []int) {

	for x := 1; x != 0; {
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				y := list[i+1]
				list[i+1] = list[i]
				list[i] = y
				x = 1
				break
			}
			x = 0
		}
	}
	return
}
