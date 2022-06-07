package coverage

import (
	"os"
	"testing"
	"time"
	"fmt"
	"strconv"

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

	two:=People{p1,p2}
	three:=People{p1,p2,p3}
	tests:=[]struct{
		actual People
		expected int
	}{
		{two,2},
		{three,3},
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
	d1,err:=time.Parse(date,"1997-08-31")
	checkErr(err)
	p1:=Person{"Mustafa","Ustaz",d1}
	d2,err:=time.Parse(date,"1994-09-09")
	checkErr(err)
	d3,err:=time.Parse(date,"2004-01-30")
	checkErr(err)

	
	p2:=Person{"Mukhamet","Ustaz",d2}
	p3:=Person{"Ayzada","Ustaz",d3}
	
	p :=People{p1,p2,p3}

	tests:=[]struct{
		actual int
		current int
		expected bool
	}{
		{0, 0, true},
		{1, 2, false},
		{2, 2, true},
	}
	for i, v:=range tests{
		if p[v.current]==p[v.actual]&&p[v.actual]==p[v.current]!= v.expected{
			t.Errorf("[%d] expected: %v, got : %v", i,v.expected,p[v.current]==p[v.actual]&&p[v.actual]==p[v.current])
		}
	}
 }

func TestNew(t *testing.T){
	_,errs:=strconv.Atoi("epam")
	tests:=[]struct{
		s string
		expected *Matrix
		err error
	}{
		{"1 1 2\n3 5 8", &Matrix{2,3, []int{1,1,2,3,5,8}},nil},
		{"1 1 2\n3 5", nil, fmt.Errorf("Rows need to be the same length")},
		{"epam 1 2\n3 5 8", nil,errs},
	}
	for i,v:=range tests{
		got ,err:=New(v.s)
		if err!=nil && v.err!=nil && err.Error() != v.err.Error(){
			t.Errorf("[%d] Error happend while not expected: %s", i, err.Error())
		} 
		if !equal(got,v.expected){
			t.Errorf("[%d] expected: %v, got %d", i, v.expected, got)
		}
	}
}

func TestRows(t *testing.T){
	t1,err:=New("1 1 2\n3 5 8")
	checkErr(err)
	t2,err:=New("1 1 2\n3 5 8\n13 21 34")
	checkErr(err)
	tests:=[]struct{
		matrix *Matrix
		expected [][]int
	}{
		{t1, [][]int{{1,1,2}, {3,5,8}}},
		{t2, [][]int{{1,1,2}, {3,5,8},{13,21,34}}},
	}
	for i,v:=range tests{
		got:=v.matrix.Rows()
		if !getMatrix(got,v.expected){
			t.Errorf("[%d] expected: %v, got %d", i, v.expected,got)
		}
	}
}

func TestCols(t *testing.T){
	t1,err:=New("1 1 2\n3 5 8")
	checkErr(err)
	tests:=[]struct{
		matrix *Matrix
		expected [][]int
	}{
		{t1, [][]int{{1,3},{1,5}, {2,8}}},
	}
	for i,v:=range tests{
		got:=v.matrix.Cols()
		if !getMatrix(got,v.expected){
			t.Errorf("[%d] expected: %v, got %d", i, v.expected,got)
		}
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

func TestSet(t *testing.T){
	t1,err:=New("3 5 8\n13 21 34")
	checkErr(err)
	t2,err:=New("3 5 8\n13 21 34")
	checkErr(err)
	t3,err:=New("1 1 2\n3 5 8\n13 21 34")
	checkErr(err)

	tests := []struct{
		matrix *Matrix
		row,col,val int
		expected bool
	}{
		{t1,0, 2, 8, true},
		{t2,-1, -1, 10,false},
		{t3,2, 0, 13, true},
	}
	for i,v:=range tests{
		got :=v.matrix.Set(v.row,v.col,v.val)
		if got!=v.expected{
			t.Errorf("[%d] expected:%v, got %v",i,v.expected,got)
		}
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