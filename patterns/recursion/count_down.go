package recursion

import "fmt"

func CountDown(num int) {
	if num <= 0 {
		fmt.Println("All done!")

		return
	}

	fmt.Println("num: ", num)
	num--

	CountDown(num)
}
