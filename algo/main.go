package algo

//Question:
//A certain bug's home is at position x. Help them get home from position 0 with minimum number of jumps.
//
//Description:
//The bug jumps according to the following rules:
//
//It can jump exactly a positions forward (to the right).
//It can jump exactly b positions backward (to the left).
//It cannot jump backward twice in a row.
//It cannot jump to any forbidden positions. (forbidden is an array of integers, where forbidden[i] is the point that the bug cannot jump to)
//The bug may jump forward beyond its home, but it cannot jump to positions numbered with negative integers.
//
//Note:
//If there is no possible sequence of jumps that lands the bug on position x, return -1.
//
//Constraints:
//1 <= forbidden.length <= 1000
//1 <= a, b, forbidden[i] <= 2000
//0 <= x <= 2000
//All the elements in forbidden are distinct.
//Position x is not forbidden.
//
//Example:
//Input: forbidden = [14,4,18,1,15], a(to the right) = 3, b(to the left) = 15, x = 9
//Output: 3
//0 -> 3 -> 3 -> 3 = 3
//
//
//Input: forbidden = [8,3,16,6,12,20], a = 15, b = 13, x = 11
//Output: -1

func main() {
	// inputs
	forbidden := []int{14, 4, 18, 1, 15} // blocked steps
	a := 3                               // move forward
	b := 15                              // move backwards
	x := 9                               // destination

	// start of code block
	path(forbidden, a, b, x)
}

func path(forbidden []int, a, b, x int) int {
	step := 0
	start := 0 // start
	for start != x {
		// move forward first
		start += a

		// check if forward is more than upper limit
		if start > x+a+b {
			// move back
			start -= b
		}

		// loop to check if step forward is forbidden
		for i := 0; i < len(forbidden); i++ {
			// if forward is inside forbidden move back
			if start == forbidden[i] {
				// reject step forward
				start -= a
			}
		}

		// add step
		step++
	}

	return step
}
