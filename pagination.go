package util

import "math"

// GetPage get page parameters, returns pagenum, pagesize
func GetPage(page int, pageSize int) (int, int) {
	if pageSize == 0 {
		pageSize = 20
	}

	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result, pageSize
}

// GetTotalPage get total page by totalCount and pageSize
func GetTotalPage(totalCount int, pageSize int) int {
	return int(math.Ceil(float64(totalCount) / float64(pageSize)))
}
