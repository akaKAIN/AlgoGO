package task_4

import (
	"fmt"
	"task_2"
)

/*
Определить, какое число в массиве встречается чаще всего.
def test(num):
	max = 0
	data = {}
	answer = ""
	l = [random.randrange(num) for i in range(num)]
	for i in l:
		data[str(i)] = data.get(str(i), 0) + 1
	for ind, val in data.items():
		if val > max:
			max = val
			answer = f'Число {ind} повторяется {val} раз(а)'
	print(answer)


*/

func FrequentCounter() {
	var randArr []int
	var max int
	var answer string

	data := make(map[int]int)
	randArr = task_2.GenArray()
	for _, val := range randArr {
		data[val]++
	}
	for key, val := range data {
		if val > max {
			max = val
			answer = fmt.Sprintf("\n\nЧисло %d встречается %d раз(а)\n", key, val)
		}
	}

	fmt.Println(answer)
}
