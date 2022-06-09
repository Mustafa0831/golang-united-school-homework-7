package coverage

import (
	"os"
	"testing"
	"time"
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

func TestLen(t *testing.T){
	var p People
	p = append(p, Person{firstName: "firstName", lastName: "Test_lastName", birthDay: time.Now() })
	p = append(p, Person{firstName: "firstName1", lastName: "Test_lastName1", birthDay: time.Now().Add(time.Minute * 2) })
	length := p.Len()
	if length != 2 {
		t.Error("eln shouldn't be equal 2 ", length)
	}

}

func TestLess(t *testing.T){
	var p People
	p = append(p, Person{firstName: "firstNam", lastName: "Test_lastName", birthDay: time.Now().Add(time.Minute *2) })
	p = append(p, Person{firstName: "firstName1", lastName: "Test_lastName1", birthDay: time.Now() })

	
	less := p.Less(0,1)

	if less != true{
		t.Error(" ", less)
	}
}

func TestLessError(t *testing.T){
	var p People
	t1 := time.Now()
	p = append(p, Person{firstName: "firstName", lastName: "Test_lastName", birthDay: t1 })
	p = append(p, Person{firstName: "firstName", lastName: "Test_lastName1", birthDay: t1 })

	
	less := p.Less(0,1)

	if less != true{
		t.Error("if it's not true ", less)
	}
}

func TestLessError1(t *testing.T){
	var p People
	t1 := time.Now()
	p = append(p, Person{firstName: "firstName", lastName: "Test_lastName", birthDay: t1 })
	p = append(p, Person{firstName: "firstName1", lastName: "Test_lastName1", birthDay: t1 })

	
	err := p.Less(0,1)

	if err != true{
		t.Error("if it's not true  ", err)
	}
}


func TestSwap(t *testing.T){
	var p People
	p = append(p, Person{firstName: "firstNam", lastName: "Test_lastName", birthDay: time.Now().Add(time.Minute *2) })
	p = append(p, Person{firstName: "firstName1", lastName: "Test_lastName1", birthDay: time.Now() })

	p0 := p[0]

	p.Swap(0,1)


	if p0 == p[0]{

		t.Error("didn't changed", p[0])
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func TestNewMatrix(t *testing.T){

	_, err := New("10 20 30 40 50 60 70 80 90")
	if err != nil {
		t.Error("can't create matrix")
	}


}

func TestNewMatrixError1(t *testing.T){

	m, err1 := New("10 20 30 40 50 60 70 80 90 \n 10 20 30 \n")
	if err1 == nil || m != nil{
		t.Error("can't create matrix ")
	}


}

func TestNewMatrixError2(t *testing.T){
	
	m0, err2 := New(" ")
	if err2 == nil || m0 != nil{
		t.Error("can't create matrix ")
	}

}

func TestNewMatrixError3(t *testing.T){
	
	m0, err3 := New("10 20 30 40 50 60 70 80 90 \n 10, 20, 30, 40, 50, 60, 70, 80, 90")
	if err3 == nil || m0 != nil{
		t.Error("can't create matrix ")
	}

}

func TestNewMatrixCols(t *testing.T){

	

	m, err := New("10 20 30 40 50 60 70 80 90")
	if err != nil {
		t.Error("can't create matrix")
	}

	var cols0 [][]int
	for i :=0; i < 9; i++{
		cols0 = append(cols0, []int{(i+1)*10})
	} 

	cols := m.Cols()



	b := compareSlices(cols, cols0)
	if b != true {
		t.Error("can't compare slice")
	}

}

func TestNewMatrixRows(t *testing.T){

	

	m, err := New("10 20 30 40 50 60 70 80 90")
	if err != nil {
		t.Error("can't get new slice")
	}


	s := []int{10,20,30,40,50,60,70,80,90}

	var rows0 [][]int
		rows0 = append(rows0, s)

	rows := m.Rows()

	b0 := compareSlicesR(rows, rows0)
	if b0 != true {
		t.Error("can't get new slice")
	}
}

func TestNewMatrixSet(t *testing.T){

	

	m, err := New("10 20 30 40 50 60 70 80 90")
	if err != nil {
		t.Error("can't set new matrix")
	}


	n := m.Set(0,1, 77)

	n0 := m.Set(-1,10,77)

	if !n || m.data[1] != 77 {
		t.Error("can't set new matrix")
	}

	if n0 != false {
		t.Error("can't set new matrix")
	}

}


func compareSlices(s1, s2 [][]int) bool {
	if len(s1) != len(s2) {
	return false
	}
	
	for i := 0; i < len(s1); i++ {
	if s1[i][0] != s2[i][0] {
	return false
	}
	}
	
	return true
   }

   func compareSlicesR(s1, s2 [][]int) bool {
	if len(s1) != len(s2) {
	return false
	}
	
	for i := 0; i < len(s1); i++ {
	if s1[0][i] != s2[0][i] {
	return false
	}
	}
	
	return true
   }