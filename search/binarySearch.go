package search

type BinaryInt struct {
	IntArray []int
	Low      int
	High     int
	Mid      int
}

func BinaryIntFactory(a []int) *BinaryInt {
	var bi BinaryInt
	bi.IntArray = a
	l := len(a)
	bi.High = l - 1
	bi.Low = 0
	bi.Mid = (l - 1) / 2
	return &bi
}

func (a *BinaryInt) BinarySearchInt(find int) int {
	p := a.IntArray[a.Mid]
	if find == p {
		return a.Mid
	} else if a.Low == a.Mid && a.High == a.Mid {
		return -1
	}
	if find > p {
		a.Low = a.Mid + 1
		a.Mid = (a.Low + a.High) / 2
		return a.BinarySearchInt(find)
	}
	a.High = a.Mid - 1
	a.Mid = (a.Low + a.High) / 2
	return a.BinarySearchInt(find)
}
