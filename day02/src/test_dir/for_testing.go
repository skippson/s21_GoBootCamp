package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Flags struct {
	mean   bool
	median bool
	mode   bool
	sd     bool
}

type Data struct {
	nums     []int
	flags    *Flags
	fail     bool
	failFlag string
}

func main() {
	argc := os.Args[1:]
	data := &Data{
		flags: &Flags{},
	}

	parseFlags(argc, data)

	data.nums = scan()
	sort.Ints(data.nums)

	print(data)
}

func parseFlags(argc []string, data *Data) {
	if len(argc) == 0 {
		data.flags.mean = true
		data.flags.median = true
		data.flags.mode = true
		data.flags.sd = true
	} else {
		for i := range argc {
			switch {
			case argc[i] == "1":
				data.flags.mean = true
			case argc[i] == "2":
				data.flags.median = true
			case argc[i] == "3":
				data.flags.mode = true
			case argc[i] == "4":
				data.flags.sd = true
			default:
				data.fail = true
				data.failFlag += argc[i] + " "
			}
		}
	}
}

func scan() []int {
	scanner := bufio.NewScanner(os.Stdin)
	var temple []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		num, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			fmt.Println("Invalid input. Enter only integer by new line")
			continue
		}
		if num <= -100000 || num >= 100000 {
			fmt.Println("Invalid input. Enter integer between -100000 and 100000")
			continue
		} else {
			temple = append(temple, int(num))
		}
	}

	return temple
}

func mean(nums []int) float64 {
	mean := 0.0
	if len(nums) != 0 {
		for _, num := range nums {
			mean += float64(num)
		}
		mean /= float64(len(nums))
	}

	return mean
}

func median(nums []int) float64 {
	median := 0.0
	if len(nums)%2 == 0 && len(nums) != 0 {
		median = float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2.0
	}

	if len(nums)%2 != 0 {
		median = float64(nums[len(nums)/2])
	}

	return median
}

func mode(nums []int) int {
	mode := 0
	table := make(map[int]int)

	if len(nums) != 0 {
		for _, num := range nums {
			table[num]++
		}
		idx := 0
		temple := 100000
		for num, i := range table {
			if i == idx && temple > num {
				temple = num
			}

			if i > idx {
				idx = i
				temple = num
			}
		}
		mode = temple
	}

	return mode
}

func sd(nums []int) float64 {
	sd := 0.0
	if len(nums) != 0 {
		mean := mean(nums)
		for _, num := range nums {
			sd += math.Pow(float64(num)-mean, 2)
		}

		sd = math.Sqrt(sd / float64(len(nums)))
	}

	return sd
}

func print(data *Data) {
	if data.flags.mean {
		fmt.Printf("Mean: %.2f\n", mean(data.nums))
	}

	if data.flags.median {
		fmt.Printf("Median: %.2f\n", median(data.nums))
	}

	if data.flags.mode {
		fmt.Printf("Mode: %d\n", mode(data.nums))
	}

	if data.flags.sd {
		fmt.Printf("SD: %.2f\n", sd(data.nums))
	}

	if data.fail {
		fmt.Printf("Invalid flag: %s\n", data.failFlag)
	}
}
