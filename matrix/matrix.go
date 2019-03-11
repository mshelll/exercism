package matrix

import (
	"strings"
	"strconv"
	"errors"
)


type Matrix struct {

	rows [][]int
	cols [][]int
}


func (m *Matrix) Rows() [][]int {

	var rows [][]int
	for _, row := range m.rows {
		var new_row = make([]int, len(row))
		copy(new_row, row)
		rows = append(rows, new_row)
	}

	return rows
}

func (m *Matrix) Cols() [][]int {

	var cols [][]int
	for _, col := range m.cols {
		var new_col = make([]int, len(col))
		copy(new_col, col)
		cols = append(cols, new_col)
	}

	return cols
}

func FindRow(line string) ([]int, error) {

	var row []int
	num_str := strings.Split(line, " ")
	for _, num := range num_str {
		if n, err := strconv.Atoi(num); err == nil {
			row = append(row, n)
		} else {
			return nil, errors.New("invalid number")
		}
	}

	return row, nil
}

func FindCols(rows [][]int, ncols int, col_len int) (out [][]int) {

	var cols [][]int
	for i:=0; i<ncols; i++ {

		var column = make([]int, col_len)
		for j:=0; j<col_len; j++ {

			column[j] = rows[j][i]
		}
		cols = append(cols, column)
	}

	return cols
}

func (m *Matrix) Validate() ( out bool, err error) {

	rowlen, prev_rowlen := 0, len(m.rows[0])

	for _, row := range m.rows[1:] {

		rowlen = len(row)
		if rowlen != prev_rowlen {
			return false, errors.New("incompatible rows")
		}
	}

	return true, nil
}

func (m *Matrix) Set(row, col, val int) (ok bool) {

	if row < 0 || row >= len(m.rows) ||
	   col < 0 || col >= len(m.cols) {

		   return false
	}
	m.rows[row][col] = val
	m.cols[col][row] = val

	return true

} 

func New(in string) ( out *Matrix, err error){

	lines := strings.Split(in, "\n")
	nrows := len(lines)
	var rows [][]int

	for _, line := range lines {
		
		row, err := FindRow(line)
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}

	var m = &Matrix{}
	m.rows = rows
	m.cols = FindCols(rows, len(rows[0]), nrows)

	valid, err := m.Validate()

	if !valid {
		return nil, err
	}

	return m, nil
}

