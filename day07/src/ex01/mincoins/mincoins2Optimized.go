package mincoins

// Минимальное количество монет для суммы val,
// оптимизации достигается за счет изменения алгоритма
func MinCoins2Optimized(val int, coins []int) []int {
	// Если сумма равна 0, то не нужно никаких монет
	if val == 0 {
		return []int{}
	}

	// Если массив монет пуст, то невозможно достичь суммы
	if coins == nil {
		return nil
	}

	// Если массив монет пуст, то невозможно достичь суммы
	if len(coins) == 0 {
		return []int{}
	}

	// Создаем массив сумм, где индекс - сумма, а значение - минимальное количество монет для этой суммы
	sums := make([]int, val+1)
	for i := 1; i <= val; i++ {
		sums[i] = val + 1 // инициализируем значения как максимально возможное количество монет
	}
	sums[0] = 0 // для суммы 0 нужно 0 монет

	// Проходим по каждой сумме от 1 до val и обновляем значения в массиве сумм
	for i := 1; i <= val; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && sums[i-coin]+1 < sums[i] {
				sums[i] = sums[i-coin] + 1 // если можно достичь суммы i с помощью монеты coin, то обновляем значение
			}
		}
	}

	// Создаем массив для хранения результата
	ans := make([]int, 0)
	idx := val
	// Проходим от суммы val до 0 и добавляем монеты в массив результата
	for idx > 0 {
		for _, coin := range coins {
			if idx-coin >= 0 && sums[idx] == sums[idx-coin]+1 {
				ans = append(ans, coin) // добавляем монету в массив результата
				idx -= coin // уменьшаем сумму на значение монеты
				break
			}
		}
	}

	return ans
}