package coverage

import (
	"os"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

const date ="2006-01-02"

func TestLen(t *testing.T){
	d1,err:=time.Parse(date,"1997-08-31")
	checkErr(err)
	d2,err:=time.Parse(date,"1994-09-09")
	checkErr(err)
	d3,err:=time.Parse(date,"2004-01-30")
	checkErr(err)
	
	p1:=Person{"Mustafa","Ustaz",d1}
	p2:=Person{"Mukhamet","Ustaz",d2}
	p3:=Person{"Ayzada","Ustaz",d3}

	three:=People{p1,p2,p3}
	zero:=People{}
	tests:=[]struct{
		actual People
		expected int
	}{
		{three,3},
		{zero,0},
	}
	for _,v:=range tests{
		if v.actual.Len()!=v.expected{
			t.Errorf("expected: %v, got: %v", v.expected,v.actual.Len())
		}
	}
}

 func TestLess(t *testing.T){
	d1,err:=time.Parse(date,"1997-08-31")
	checkErr(err)
	d2,err:=time.Parse(date,"1994-09-09")
	checkErr(err)
	d3,err:=time.Parse(date,"2004-01-30")
	checkErr(err)

	p1:=Person{"Mustafa","Ustaz",d1}
	p2:=Person{"Mukhamet","Ustaz",d2}
	p3:=Person{"Ayzada","Ustaz",d3}

	p :=People{p1,p2,p3}

	tests:=[]struct{
		actual int
		current int
		expected bool
	}{
		{0, 1, true},
		{0, 2, false},
		{1, 2, false},
		{2, 0, true},
	}
	for i, v:=range tests{
		if p.Less(v.actual,v.current)!= v.expected{
			t.Errorf("[%d] expected: %v, got : %v", i,v.expected, p.Less(v.actual, v.current))
		}
	}
 }

 func TestSwap(t *testing.T){
	t.Parallel()
	d1,err:=time.Parse(date,"1997-08-31")
	checkErr(err)
	d2,err:=time.Parse(date,"1994-09-09")
	checkErr(err)
	d3,err:=time.Parse(date,"2004-01-30")
	checkErr(err)
	
	p1:=Person{"Mustafa","Ustaz",d1}
	p2:=Person{"Mukhamet","Ustaz",d2}
	p3:=Person{"Ayzada","Ustaz",d3}
	
	p :=People{p1,p2,p3}
	p.Swap(1,2)
	assert.Equalf(t,p[1], p3,"Ayzada earlier")
	assert.Equalf(t,p[2], p2,"Mukhamet later")
 }

func TestNewNil(t *testing.T){
	t.Parallel()
	tests:=map[string]string{
		"nil" : "",
		"matrix" : "0 2\n0 5 8",
	}
	for i, test:=range tests{
		t.Run(i,func(t *testing.T){
			t.Parallel()
			expected,err:=New(test)
			assert.NotNil(t,err)
			assert.Nil(t,expected)
		})
	}
}

func TestNewNotNil(t *testing.T){
	t.Parallel()
	tests:=map[string]struct{
		matrix string
		row, col int
	}{
		"single row": {"0 1 2 3", 1, 4},
		"multiple rows": {"0 1 2\n3 4 5\n6 7 8", 3, 3},
		"multiple rows with leading and trailing spaces": {" 0 1 2\n 3 4 5 \n6 7 8 ", 3, 3},
	}
	for i,test:=range tests{
		t.Run(i,func(t *testing.T){
			t.Parallel()
			expected,err:=New(test.matrix)
			assert.Nil(t,err)
			assert.NotNil(t,expected)
			assert.NotNil(t,expected.data)
			assert.Equal(t,test.row,expected.rows)
			assert.Equal(t,test.col,expected.cols)
		})
	}
}

func TestRows(t *testing.T){
	t.Parallel()
	tests:=map[string]struct{
		matrix string
		expected [][]int
	}{
		"single row": {"0 1 2 3", [][]int{{0, 1, 2, 3}}},
		"multiple rows": {"0 1 2\n3 4 5\n6 7 8", [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}},
	}
	for i,v:=range tests{
		t.Run(i, func(t *testing.T){
			t.Parallel()
			m,_:=New(v.matrix)
			
			assert.Equal(t,v.expected,m.Rows())
		})
	}
}

func TestCols(t *testing.T){
	t.Parallel()
	tests:=map[string]struct{
		matrix string
		expected [][]int
	}{
		"single row": {"0 1 2 3", [][]int{{0}, {1}, {2}, {3}}},
		"multiple rows": {"0 1 2\n3 4 5\n6 7 8", [][]int{{0, 3, 6}, {1, 4, 7}, {2, 5, 8}}},
	}
	for i,v:=range tests{
		t.Run(i, func(t *testing.T){
			t.Parallel()
			m,_:=New(v.matrix)
			
			assert.Equal(t,v.expected,m.Cols())
		})
	}
}

func getMatrix(matrix1, matrix2 [][]int)bool{
	if len(matrix1)!= len(matrix2){
		return false
	}
	for i, r1:=range matrix1{
		r2:=matrix2[i]
		if len(r1)!=len(r2){
			return false
		}
		for j,v:=range r1{
			if v!=r2[j]{
				return false
			}
		}
	}
	return true
}

func TestSetFails(t *testing.T) {
	t.Parallel()

	matrix, _ := New("0 1 2")
	tData := map[string]struct{
		row int
		col int
	}{
		"negative row": {-1, 0},
		"row out of range": {1, 0},
		"negative col": {0, -1},
		"col out of range": {0, 3}}

	for name, value := range tData {
		tCase := value
		t.Run(name, func (t *testing.T) {
			t.Parallel()
			assert.False(t, matrix.Set(tCase.row, tCase.col, 100))
		})
	}
}

func TestSet(t *testing.T) {
	t.Parallel()

	tData := map[string]struct{
		input string
		row int
		col int
		value int
	}{
		"single row": {"0 1 2", 0, 1, 3},
		"multiple rows": {"0 1 2\n3 4 5\n6 7 8", 2, 2, 9}}

	for name, value := range tData {
		tCase := value
		t.Run(name, func (t *testing.T) {
			t.Parallel()
			matrix, _ := New(tCase.input)
			assert.True(t, matrix.Set(tCase.row, tCase.col, tCase.value))
			assert.Equal(t, tCase.value, matrix.Rows()[tCase.row][tCase.col])
		})
	}
}

func equal(matrix1, matrix2 *Matrix)bool{
	if matrix1==matrix2{
		return true
	}
	if matrix1.cols==matrix2.cols && matrix1.rows==matrix2.rows &&len(matrix1.data)==len(matrix2.data) {
		for i,v:=range matrix1.data{
			if v!=matrix2.data[i]{
				return false
			}
		}
		return true
	}
	return false
}

func checkErr(err error){
	if err != nil {
        panic(err)
    }
}