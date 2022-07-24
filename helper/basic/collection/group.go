package collection

import "math"

func GroupIntSlice(sli []int, batchSize int) [][]int {
	totalCount := len(sli)

	if totalCount <= batchSize {
		return [][]int{sli}
	}

	cnt := int(math.Ceil(float64(totalCount) / float64(batchSize)))
	result := make([][]int, cnt)

	for i := 0; i < cnt; i++ {
		start := i * batchSize
		end := start + batchSize

		if end >= totalCount {
			result[i] = sli[start:totalCount]
		} else {
			result[i] = sli[start:end]
		}
	}

	return result
}

func GroupInt32Slice(sli []int32, batchSize int) [][]int32 {
	totalCount := len(sli)

	if totalCount <= batchSize {
		return [][]int32{sli}
	}

	cnt := int(math.Ceil(float64(totalCount) / float64(batchSize)))
	result := make([][]int32, cnt)

	for i := 0; i < cnt; i++ {
		start := i * batchSize
		end := start + batchSize

		if end >= totalCount {
			result[i] = sli[start:totalCount]
		} else {
			result[i] = sli[start:end]
		}
	}

	return result
}

func GroupInt64Slice(sli []int64, batchSize int) [][]int64 {
	totalCount := len(sli)

	if totalCount <= batchSize {
		return [][]int64{sli}
	}

	cnt := int(math.Ceil(float64(totalCount) / float64(batchSize)))
	result := make([][]int64, cnt)

	for i := 0; i < cnt; i++ {
		start := i * batchSize
		end := start + batchSize

		if end >= totalCount {
			result[i] = sli[start:totalCount]
		} else {
			result[i] = sli[start:end]
		}
	}

	return result
}

func GroupStrSlice(sli []string, batchSize int) [][]string {
	totalCount := len(sli)

	if totalCount <= batchSize {
		return [][]string{sli}
	}

	cnt := int(math.Ceil(float64(totalCount) / float64(batchSize)))
	result := make([][]string, cnt)

	for i := 0; i < cnt; i++ {
		start := i * batchSize
		end := start + batchSize

		if end >= totalCount {
			result[i] = sli[start:totalCount]
		} else {
			result[i] = sli[start:end]
		}
	}

	return result
}

func GroupFloat32Slice(sli []float32, batchSize int) [][]float32 {
	totalCount := len(sli)

	if totalCount <= batchSize {
		return [][]float32{sli}
	}

	cnt := int(math.Ceil(float64(totalCount) / float64(batchSize)))
	result := make([][]float32, cnt)

	for i := 0; i < cnt; i++ {
		start := i * batchSize
		end := start + batchSize

		if end >= totalCount {
			result[i] = sli[start:totalCount]
		} else {
			result[i] = sli[start:end]
		}
	}

	return result
}

func GroupFloat64Slice(sli []float64, batchSize int) [][]float64 {
	totalCount := len(sli)

	if totalCount <= batchSize {
		return [][]float64{sli}
	}

	cnt := int(math.Ceil(float64(totalCount) / float64(batchSize)))
	result := make([][]float64, cnt)

	for i := 0; i < cnt; i++ {
		start := i * batchSize
		end := start + batchSize

		if end >= totalCount {
			result[i] = sli[start:totalCount]
		} else {
			result[i] = sli[start:end]
		}
	}

	return result
}

func GroupAnySlice(sli []interface{}, batchSize int) [][]interface{} {
	totalCount := len(sli)

	if totalCount <= batchSize {
		return [][]interface{}{sli}
	}

	cnt := int(math.Ceil(float64(totalCount) / float64(batchSize)))
	result := make([][]interface{}, cnt)

	for i := 0; i < cnt; i++ {
		start := i * batchSize
		end := start + batchSize

		if end >= totalCount {
			result[i] = sli[start:totalCount]
		} else {
			result[i] = sli[start:end]
		}
	}

	return result
}
