package mincoins

// Функция по умолчанию для нахождения минимального количества монет для суммы val
func MinCoins(val int, coins []int) []int {
	// Создаем массив для хранения результата
	res := make([]int, 0)
	// Начинаем с последней монеты в массиве
	i := len(coins) - 1
	// Проходим по монетам в порядке убывания
	for i >= 0 {
		// Добавляем монеты в массив результата, пока сумма val не станет меньше значения монеты
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i]) // добавляем монету в массив результата
		}
		i -= 1 // переходим к следующей монете
	}
	return res
}