package array_slice

func AppendSlice(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		//there is room to grow.
		z = x[:zlen]
	} else {
		//there is no room to grow, insufficient array
		//allocate new array,  double the len.
		zcap := zlen
		if zcap <= 2*len(x) {
			zcap = 2 * len(x)
		}

		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
