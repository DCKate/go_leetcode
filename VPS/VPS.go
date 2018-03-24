//Package VPS Valid Perfect Square
package VPS

func nextPossibleValue(a int, b int) int {
	r := a + b
	if r%2 == 0 {
		return r / 2
	}
	return (r + 1) / 2

}

//IsPerfectSquare check if is perfect number
func IsPerfectSquare(num int) bool {
	end := num
	start := 0
	for end != start {
		tmp := nextPossibleValue(end, start)
		rr := tmp * tmp
		if rr == num {
			return true
		} else if rr > num && end != tmp {
			end = tmp
		} else if rr < num && start != tmp {
			start = tmp
		} else {
			break
		}

	}
	return false
}
