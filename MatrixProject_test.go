package main

import (
	"testing"
)

func TestMatrix_CheckRowsBeforeInitMatrixWithRows(t *testing.T) {
	matrix := new(Matrix)
	got := matrix.CheckRowsBeforeInitMatrixWithRows(2)
	if len(got.data) != 2 {
		t.Errorf("CheckRowsBeforeInitMatrixWithRows() = %v, want %v", len(got.data), 2)
	}
	for _, row := range got.data {
		if len(row) != 2 {
			t.Errorf("len(raws) = %d, want 2", len(row))
		}
	}
}

func TestMatrix_CheckColumnsBeforeInitMatrix(t *testing.T) {
	matrix := new(Matrix)
	got := matrix.CheckColumnsBeforeInitMatrix(2, 4)
	if len(got.data) != 2 {
		t.Errorf("CheckColumnsBeforeInitMatrix() = %v, want %v", len(got.data), 2)
	}
	for _, row := range got.data {
		if len(row) != 4 {
			t.Errorf("len(raws) = %d, want 2", len(row))
		}
	}
}

func TestMatrix_CheckFillMatrixByNumber(t *testing.T) {
	matrix := new(Matrix)
	got := matrix.CheckFillMatrixByNumber(2, 2, 3)
	for _, rows := range got.data {
		for _, row := range rows {
			if row != 3 {
				t.Errorf("CheckFillMatrixByNumber() = %v, want %v", row, 3)
			}
		}
	}
}

func TestMatrix_MultiplyMatrixScalar(t *testing.T) {
	cases := []struct {
		name       string
		rows       int
		columns    int
		fillNumber int
		scalar     int
		want       [][]int
	}{
		{
			"first",
			2,
			2,
			2,
			3,
			[][]int{{6, 6}, {6, 6}},
		},
	}

	for _, tc := range cases {
		tc := tc
		matrix := new(Matrix)
		matrix.NewMatrix(tc.rows, tc.columns)
		matrix.FillMatrix(tc.fillNumber)
		t.Run(tc.name, func(t *testing.T) {
			got := matrix.MultiplyMatrixScalar(tc.scalar)
			if !compareMatrices(got.data, tc.want) {
				t.Errorf("MultiplyMatrix() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMatrix_SumMatrix(t *testing.T) {
	cases := []struct {
		name        string
		rows1       int
		columns1    int
		fillNumber1 int
		rows2       int
		columns2    int
		fillNumber2 int
		want        [][]int
	}{
		{
			"first",
			2,
			2,
			2,
			2,
			2,
			2,
			[][]int{{4, 4}, {4, 4}},
		},
	}
	for _, tc := range cases {
		tc := tc
		matrix := new(Matrix)
		matrix.NewMatrix(tc.rows1, tc.columns1)
		matrix.FillMatrix(tc.fillNumber1)
		matrix2 := new(Matrix)
		matrix2.NewMatrix(tc.rows2, tc.columns2)
		matrix2.FillMatrix(tc.fillNumber2)
		t.Run(tc.name, func(t *testing.T) {
			got := matrix.SumMatrix(matrix2)
			if !compareMatrices(got.data, tc.want) {
				t.Errorf("SumMatrix() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMatrix_GetRowByIndex(t *testing.T) {
	matrix := new(Matrix)
	matrix.NewMatrix(2, 2)
	matrix.FillMatrix(1)
	got := matrix.GetRowByIndex(0)
	if !compareSlices(got, []int{1, 1}) {
		t.Errorf("GetRowByIndex() = %v, want %v", got, []int{1, 1})
	}
}

func TestMatrix_GetColumnByIndex(t *testing.T) {
	matrix := new(Matrix)
	matrix.NewMatrix(2, 2)
	matrix.FillMatrix(1)
	got := matrix.GetRowByIndex(0)
	if !compareSlices(got, []int{1, 1}) {
		t.Errorf("GetRowByIndex() = %v, want %v", got, []int{1, 1})
	}
}

func TestMatrix_MultiplyMatrix(t *testing.T) {
	matrix := new(Matrix)
	matrix.NewMatrix(2, 2)
	matrix.FillMatrix(2)
	matrix2 := new(Matrix)
	matrix2.NewMatrix(2, 2)
	matrix2.FillMatrix(2)
	got := matrix.MultiplyMatrix(matrix2)
	want := [][]int{
		{8, 8},
		{8, 8},
	}
	if !compareMatrices(got, want) {
		t.Errorf("MultiplyMatrix() = %v, want %v", got, want)
	}
}

func TestMatrix_ConvertToDiagonalMatrix(t *testing.T) {
	matrix1 := new(Matrix)
	matrix1.data = [][]int{{20, 30, 30}, {10, 20, 30}, {10, 40, 20}}
	matrix1.rows = 3
	matrix1.columns = 3
	got := matrix1.ConvertToDiagonalMatrix()
	want := [][]int{
		{20, 30, 30},
		{0, 5, 15},
		{0, 0, -70},
	}
	if !compareMatrices(got.data, want) {
		t.Errorf("ConvertToDiagonaleMatrix() = %v, want %v", got.data, want)
	}
}

func TestMatrix_GetDeterMatrix(t *testing.T) {
	matrix1 := new(Matrix)
	matrix1.data = [][]int{{20, 30, 30}, {10, 20, 30}, {10, 40, 20}}
	matrix1.rows = 3
	matrix1.columns = 3
	got := matrix1.ConvertToDiagonalMatrix()
	res := got.GetDeterMatrix()
	want := -7000
	if res != want {
		t.Errorf("GetDeterMatrix() = %v, want %v", got, want)
	}
}

func TestMatrix_String(t *testing.T) {
	matrix1 := new(Matrix)
	matrix1.data = [][]int{{20, 30, 30}, {10, 20, 30}, {10, 40, 20}}
	matrix1.rows = 3
	matrix1.columns = 3
	got := matrix1.String()
	want := "20 30 30 \n10 20 30 \n10 40 20 \n"
	if got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func compareSlices(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func compareMatrices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
