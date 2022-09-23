//这里定义接口和方法，方法的实现需要依赖外部数据库
package employee

type Result struct {
	Count int
}

func (r Result) Int() int {
	return r.Count
}

type Rows []struct{} //表示任意结构体---对应于数据库的记录

type Stmt interface {
	Clode() error
	NumInput() int
	Exec(stmt string, args ...string) (Result, error)
	Query(args []string) (Rows, error)
}

func MaleCount(s Stmt) (int, error) { //参数是接口（注意接口也是一种类型）
	result, err := s.Exec("select count(*) from employee_tab where gender=?", "1")
	if err != nil {
		return 0, err
	}
	return result.Int(), nil
}
