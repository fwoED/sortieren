package quicksort

// QuickSort sortiert die gegebene Liste mit dem Quicksort-Algorithmus.
func QuickSort(list []int) {
	if len(list) < 2 {
		return
	}

	Index := len(list)- 1
	wechselwert := list[Index]

	links := 0
	for rechts := 0; rechts < Index; rechts++ {
		if list[rechts] < wechselwert {

			list[links], list[rechts] = list[rechts], list[links]
			links++
		}
	}

	list[links], list[Index] = list[Index], list[links]

	QuickSort(list[:links])
	QuickSort(list[links+1:])

}
