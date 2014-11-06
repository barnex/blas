package blas

import (
	"fmt"
)

func ExampleMakeFloat32Matrix() {
	m := MakeFloat32Matrix(2, 3)
	fmt.Println(m)

	//Output:
	//[[0 0 0] [0 0 0]]
}

func ExampleMakeFloat64Matrix() {
	m := MakeFloat64Matrix(2, 3)
	fmt.Println(m)

	//Output:
	//[[0 0 0] [0 0 0]]
}

func ExampleMakeComplex64Matrix() {
	m := MakeComplex64Matrix(2, 3)
	fmt.Println(m)

	//Output:
	//[[(0+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]]
}

func ExampleMakeComplex128Matrix() {
	m := MakeComplex128Matrix(2, 3)
	fmt.Println(m)

	//Output:
	//[[(0+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]]
}

func ExampleSSubmatrix() {
	M := MakeFloat32Matrix(3, 4)
	for i := range M {
		for j := range M[i] {
			M[i][j] = float32(10*i + j)
		}
	}
	fmt.Println(M)

	k1, k2 := 0, 1 // offset
	n1, n2 := 2, 3 // size
	m := SSubmatrix(M, k1, k2, n1, n2)
	fmt.Println(m)

	//Output:
	//[[0 1 2 3] [10 11 12 13] [20 21 22 23]]
	//[[1 2 3] [11 12 13]]
}

func ExampleSSize() {
	M := MakeFloat32Matrix(3, 4)
	rows, cols, stride := SSize(M)
	fmt.Println("Original:  rows=", rows, "cols=", cols, "stride=", stride)

	k1, k2 := 0, 1 // offset
	n1, n2 := 2, 3 // size
	m := SSubmatrix(M, k1, k2, n1, n2)
	rows, cols, stride = SSize(m)
	fmt.Println("Submatrix: rows=", rows, "cols=", cols, "stride=", stride)

	//Output
	//Original:  rows= 3 cols= 4 stride= 4
	//Submatrix: rows= 2 cols= 3 stride= 4:
}

func ExampleDSubmatrix() {
	M := MakeFloat64Matrix(3, 4)
	for i := range M {
		for j := range M[i] {
			M[i][j] = float64(10*i + j)
		}
	}
	fmt.Println(M)

	k1, k2 := 0, 1 // offset
	n1, n2 := 2, 3 // size
	m := DSubmatrix(M, k1, k2, n1, n2)
	fmt.Println(m)

	//Output:
	//[[0 1 2 3] [10 11 12 13] [20 21 22 23]]
	//[[1 2 3] [11 12 13]]
}

func ExampleDSize() {
	M := MakeFloat64Matrix(3, 4)
	rows, cols, stride := DSize(M)
	fmt.Println("Original:  rows=", rows, "cols=", cols, "stride=", stride)

	k1, k2 := 0, 1 // offset
	n1, n2 := 2, 3 // size
	m := DSubmatrix(M, k1, k2, n1, n2)
	rows, cols, stride = DSize(m)
	fmt.Println("Submatrix: rows=", rows, "cols=", cols, "stride=", stride)

	//Output
	//Original:  rows= 3 cols= 4 stride= 4
	//Submatrix: rows= 2 cols= 3 stride= 4:
}

func ExampleCSubmatrix() {
	M := MakeComplex64Matrix(3, 4)
	for i := range M {
		for j := range M[i] {
			M[i][j] = complex(float32(i), float32(j))
		}
	}
	fmt.Println(M)

	k1, k2 := 0, 1 // offset
	n1, n2 := 2, 3 // size
	m := CSubmatrix(M, k1, k2, n1, n2)
	fmt.Println(m)

	//Output:
	//[[(0+0i) (0+1i) (0+2i) (0+3i)] [(1+0i) (1+1i) (1+2i) (1+3i)] [(2+0i) (2+1i) (2+2i) (2+3i)]]
	//[[(0+1i) (0+2i) (0+3i)] [(1+1i) (1+2i) (1+3i)]]
}

func ExampleCSize() {
	M := MakeComplex64Matrix(3, 4)
	rows, cols, stride := CSize(M)
	fmt.Println("Original:  rows=", rows, "cols=", cols, "stride=", stride)

	k1, k2 := 0, 1 // offset
	n1, n2 := 2, 3 // size
	m := CSubmatrix(M, k1, k2, n1, n2)
	rows, cols, stride = CSize(m)
	fmt.Println("Submatrix: rows=", rows, "cols=", cols, "stride=", stride)

	//Output
	//Original:  rows= 3 cols= 4 stride= 4
	//Submatrix: rows= 2 cols= 3 stride= 4:
}

func ExampleZSubmatrix() {
	M := MakeComplex128Matrix(3, 4)
	for i := range M {
		for j := range M[i] {
			M[i][j] = complex(float64(i), float64(j))
		}
	}
	fmt.Println(M)

	k1, k2 := 0, 1 // offset
	n1, n2 := 2, 3 // size
	m := ZSubmatrix(M, k1, k2, n1, n2)
	fmt.Println(m)

	//Output:
	//[[(0+0i) (0+1i) (0+2i) (0+3i)] [(1+0i) (1+1i) (1+2i) (1+3i)] [(2+0i) (2+1i) (2+2i) (2+3i)]]
	//[[(0+1i) (0+2i) (0+3i)] [(1+1i) (1+2i) (1+3i)]]
}

func ExampleZSize() {
	M := MakeComplex128Matrix(3, 4)
	rows, cols, stride := ZSize(M)
	fmt.Println("Original:  rows=", rows, "cols=", cols, "stride=", stride)

	k1, k2 := 0, 1 // offset
	n1, n2 := 2, 3 // size
	m := ZSubmatrix(M, k1, k2, n1, n2)
	rows, cols, stride = ZSize(m)
	fmt.Println("Submatrix: rows=", rows, "cols=", cols, "stride=", stride)

	//Output
	//Original:  rows= 3 cols= 4 stride= 4
	//Submatrix: rows= 2 cols= 3 stride= 4:
}