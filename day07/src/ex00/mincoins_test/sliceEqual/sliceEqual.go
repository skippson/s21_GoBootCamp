package sliceequal

// Проверяет, равны ли два среза целых чисел
func SliceEqual(a, b []int) bool {
	// Если длины срезов не равны, то они не равны
	if len(a) != len(b) {
		return false
	}

	// Проходим по элементам срезов и сравниваем их
	for i := range a {
		// Если хотя бы один элемент не равен, то срезы не равны
		if a[i] != b[i] {
			return false
		}
	}

	// Если все элементы равны, то срезы равны
	return true
}