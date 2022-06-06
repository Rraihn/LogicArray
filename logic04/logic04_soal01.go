package logic04

import array2 "github.com/Rraihn/logic-dasar2/array"

func Logic04Soal01(n int) {
	// create array
	// name: array
	// isinya array2 (merujuk ke package array), NewNumberArray (variabel yang ada di create_array) ; isinya n,n
	array := array2.NewNumberArray(n, n)
	a := 2
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if i%4 == 0 {
				array[i][j] = int32(a)
				a++
			} else if i%4 == 1 && j == n-1 {
				array[i][j] = int32(a)
				a++
			} else if i%4 == 2 {
				array[i][n-1-j] = int32(a)
				a++
			} else if i%4 == 3 && j == 0 {
				array[i][j] = int32(a)
				a++
			}
		}
	}
	// print array
	array2.PrintNumberArray(array)
}
