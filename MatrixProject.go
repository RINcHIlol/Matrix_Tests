package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	data    [][]int
	rows    int
	columns int
}

func (m *Matrix) NewMatrix(rows, columns int) *Matrix {
	m.data = make([][]int, rows)
	for i := range m.data {
		m.data[i] = make([]int, columns)
	}
	m.rows = rows
	m.columns = columns
	return m
}

func (m *Matrix) FillMatrix(number int) {
	for i := range m.data {
		for j := range m.data[i] {
			m.data[i][j] = number
		}
	}
}

func MultVector(vectorOne, vectorTwo []int) int {
	if len(vectorOne) != len(vectorTwo) {
		return 0 // или вернуть ошибку
	}
	var res int
	for i := range vectorOne {
		res += vectorOne[i] * vectorTwo[i]
	}
	return res
}

func (m *Matrix) CheckRowsBeforeInitMatrixWithRows(rows int) *Matrix {
	return m.NewMatrix(rows, rows)
}

func (m *Matrix) CheckColumnsBeforeInitMatrix(rows, columns int) *Matrix {
	return m.NewMatrix(rows, columns)
}

func (m *Matrix) CheckFillMatrixByNumber(rows, columns, fillNumber int) *Matrix {
	matrix := m.NewMatrix(rows, columns)
	matrix.FillMatrix(fillNumber)
	return matrix
}

func (m *Matrix) MultiplyMatrixScalar(scalar int) *Matrix {
	for i := range m.data {
		for j := range m.data[i] {
			m.data[i][j] = m.data[i][j] * scalar
		}
	}
	return m
}

func (m *Matrix) SumMatrix(matrix *Matrix) *Matrix {
	if m.rows != matrix.rows || m.columns != matrix.columns {
		fmt.Errorf("ohh.. unluck")
		return nil
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			m.data[i][j] = m.data[i][j] + matrix.data[i][j]
		}
	}
	return m
}

func (m *Matrix) GetRowByIndex(row int) []int {
	if row < 0 || row >= m.rows {
		return nil
	}

	return m.data[row]
}

func (m *Matrix) GetColumnByIndex(col int) []int {
	if col < 0 || col >= m.columns {
		return nil
	}

	columns := make([]int, m.rows)
	for i := 0; i < m.rows; i++ {
		columns[i] = m.data[i][col]
	}
	return columns
}

func (m *Matrix) MultiplyMatrix(matrixTwo *Matrix) [][]int {
	// Проверка совместимости размеров
	if m.columns != matrixTwo.rows {
		return nil
	}

	// Инициализация результирующей матрицы
	result := make([][]int, m.rows)
	for i := range result {
		result[i] = make([]int, matrixTwo.columns) // Инициализация строк
	}

	// Умножение матриц
	for i := 0; i < m.rows; i++ {
		for j := 0; j < matrixTwo.columns; j++ {
			result[i][j] = MultVector(m.GetRowByIndex(i), matrixTwo.GetColumnByIndex(j))
		}
	}

	return result
}

func (m *Matrix) ConvertToDiagonalMatrix() *Matrix {
	rows, columns := m.rows, m.columns
	for kIndex := 0; kIndex < rows; kIndex++ {
		for i := kIndex + 1; i < rows; i++ {
			koef := float64(m.data[i][kIndex]) / float64(m.data[kIndex][kIndex])
			for j := kIndex; j < columns; j++ {
				m.data[i][j] -= int(float64(m.data[kIndex][j]) * koef)
			}
		}
	}
	return m
}

func (m *Matrix) GetDeterMatrix() int {
	res := 1
	for i := 0; i < m.rows; i++ {
		res *= m.data[i][i]
	}
	return res
}

func (m *Matrix) String() string {
	var stringBuilder = strings.Builder{}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			stringBuilder.WriteString(strconv.Itoa(m.data[i][j]))
			stringBuilder.WriteString(" ")
		}
		stringBuilder.WriteRune('\n')
	}
	return stringBuilder.String()
}

func main() {
	matrix1 := new(Matrix)
	matrix1.data = [][]int{{20, 30, 30}, {10, 20, 30}, {10, 40, 20}}
	matrix1.rows = 3
	matrix1.columns = 3

	got := matrix1.ConvertToDiagonalMatrix()
	fmt.Println(got.String())
}
