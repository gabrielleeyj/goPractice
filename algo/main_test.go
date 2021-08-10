package algo

import "testing"

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
type Bug struct {
	position           int
	stepback           bool
	forbiddenPositions []int
}

func (b *Bug) back(j int) {
	if b.stepback {
		return
	}
	b.position -= j
	b.stepback = true
}

func (b *Bug) forward(j int) {
	b.position += j
	b.stepback = false
}

func (b *Bug) getCurrentPosition() int {
	return b.position
}

func NewBug(forbiddenPosition []int) Bug {
	return Bug{forbiddenPositions: forbiddenPosition}
}

func TestMoveForward(t *testing.T) {
	var b Bug
	if b.getCurrentPosition() != 0 {
		t.FailNow()
	}
	b.forward(15)
	if b.getCurrentPosition() != 15 {
		t.FailNow()
	}
}

func TestMoveBackward(t *testing.T) {
	var b Bug
	if b.getCurrentPosition() != 0 {
		t.FailNow()
	}
	b.back(3)
	if b.getCurrentPosition() != -3 {
		t.FailNow()
	}
}

func TestMoveBackTwice(t *testing.T) {
	var b Bug
	b.back(3)
	t.Logf("Moved back: %d, current position: %d", 3, b.getCurrentPosition())
	at := b.getCurrentPosition()
	b.back(3)
	t.Logf("Moved back: %d, current position: %d", 3, b.getCurrentPosition())
	if at != b.getCurrentPosition() {
		t.FailNow()
	}
}

func TestMoveBackAfterForward(t *testing.T) {
	var b Bug
	at := b.getCurrentPosition()
	b.back(3)
	t.Logf("Moved back: %d, current position: %d", 3, b.getCurrentPosition())
	if at == b.getCurrentPosition() {
		t.FailNow()
	}
	b.forward(1)
	t.Logf("Moved forward: %d, current position: %d", 1, b.getCurrentPosition())
	b.back(1)
	t.Logf("Moved back: %d, current position: %d", 1, b.getCurrentPosition())
	if at == b.getCurrentPosition() {
		t.FailNow()
	}

}

func TestForbiddenPos(t *testing.T) {
	b := NewBug([]int{14, 4, 18, 1, 15})
	at := b.getCurrentPosition()
	b.forward(14)
	if at != b.getCurrentPosition() {
		t.Fail()
	}
	b.forward(4)
	if at != b.getCurrentPosition() {
		t.Fail()
	}
	b.forward(1)
	if at != b.getCurrentPosition() {
		t.Fail()
	}
	b.forward(15)
	if at != b.getCurrentPosition() {
		t.Fail()
	}
	b.forward(18)
	if at != b.getCurrentPosition() {
		t.Fail()
	}
	b.forward(2)
	if at == b.getCurrentPosition() {
		t.Fail()
	}
}
