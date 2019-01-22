package controller

//paging for ui paging
func paging(start, length, totalCount int) (int, int) {
	if length == 0 {
		length = 20
	}

	start = totalCount - start - length
	if start < 0 {
		length += start
		start = 0
	}
	return start, length
}
