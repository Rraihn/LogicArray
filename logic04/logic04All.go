package logic04

import array2 "github.com/Rraihn/logic-dasar2/array"

func Logic04Soal01(n int) {
	// create array
	// name: array
	// isinya array2 (merujuk ke package array), NewNumberArray (variabel yang ada di create_array) ; isinya n,n
	array := array2.NewNumberArray(n, n)
	a := 2                            // set new var a ; value 2
	for i := 0; i < len(array); i++ { //looping baris
		for j := 0; j < len(array[i]); j++ { // looping kolom
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
	array2.PrintNumberArrayZeroEmpty(array)
}

func Logic04Soal02(n int) {
	// create array
	// name: array
	// isinya array2 (merujuk ke package array), NewNumberArray (variabel yang ada di create_array) ; isinya n,n
	array := array2.NewNumberArray(n, n)
	a := 1                   // set new var a ; value 1
	for i := 0; i < n; i++ { //looping baris
		for j := 0; j < n; j++ { // looping kolom
			if i%4 == 0 {
				array[j][i] = int32(a)
				a += 3
			} else if i%4 == 1 && j == n-1 {
				array[j][i] = int32(a)
				a += 3
			} else if i%4 == 2 {
				array[n-1-j][i] = int32(a)
				a += 3
			} else if i%4 == 3 && j == 0 {
				array[j][i] = int32(a)
				a += 3
			}
		}
	}
	// print array
	array2.PrintNumberArrayZeroEmpty(array)
}

func Logic04Soal03(n int) {
	array := array2.NewNumberArray(n, n)
	a := 1                   // set new var a ; value 1
	for i := 0; i < n; i++ { //looping baris
		for j := 0; j < n; j++ { // looping kolom
			if i%4 == 0 {
				array[n-1-j][i] = int32(a)
				a += 3
			} else if i%4 == 1 && j == 0 {
				array[j][i] = int32(a)
				a += 3
			} else if i%4 == 2 {
				array[j][i] = int32(a)
				a += 3
			} else if i%4 == 3 && j == n-1 {
				array[j][i] = int32(a)
				a += 3
			}
		}
	}
	// print array
	array2.PrintNumberArrayZeroEmpty(array)
}

func Logic04Soal04(n int) {
	array := array2.NewNumberArray(n, n)
	a := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%4 == 0 {
				array[i][n-1-j] = int32(a)
				a += 3
			} else if i%4 == 1 && j == 0 {
				array[i][j] = int32(a)
				a += 3
			} else if i%4 == 2 {
				array[i][j] = int32(a)
				a += 3
			} else if i%4 == 3 && j == n-1 {
				array[i][j] = int32(a)
				a += 3
			}
		}
	}
	// print array
	array2.PrintNumberArrayZeroEmpty(array)
}
