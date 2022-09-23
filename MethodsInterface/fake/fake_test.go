//使用伪对象去测试
package fake

import (
	"fake/employee"
	"fmt"
	"testing"
)

type fakeStmtMaleCount struct {
	employee.Stmt
}

func (fakeStmtMaleCount) Exec(stmt string, args ...string) (employee.Result, error) {
	return employee.Result{Count: 5}, nil //伪对象直接返回
}

func TestEmployeeMaleCount(t *testing.T) {
	f := fakeStmtMaleCount{}
	c, _ := employee.MaleCount(f)
	if c != 5 {
		fmt.Printf("want:%d,actual:%d", 5, c)
		return
	}
}
