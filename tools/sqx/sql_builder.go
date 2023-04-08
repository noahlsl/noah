package sqx

import (
	"fmt"
	"strings"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/noahlsl/noah/tools/slicex"
	"github.com/pkg/errors"
)

type table interface {
	tableName() string
}
type SqlBuilder struct {
	selectDataset  *goqu.SelectDataset
	deleteDataset  *goqu.DeleteDataset
	updateDataset  *goqu.UpdateDataset
	installDataset *goqu.InsertDataset
	ex             []goqu.Expression
	asc            []string
	desc           []string
	groupBy        []string
	having         []goqu.Expression
	page           uint64
	limit          uint64
	tableName      string
}

func NewSelectBuilder(table string, fields ...string) *SqlBuilder {
	var f []interface{}
	for _, field := range fields {
		f = append(f, field)
	}
	return &SqlBuilder{
		selectDataset: goqu.Select(f...).From(table),
	}
}
func NewDeleteBuilder(table string) *SqlBuilder {
	return &SqlBuilder{
		deleteDataset: goqu.Delete(table),
	}
}
func NewUpdateBuilder(table string, val map[string]interface{}) *SqlBuilder {
	return &SqlBuilder{
		updateDataset: goqu.Update(table).Set(val),
	}
}
func NewInstallBuilder(table string, rows ...interface{}) *SqlBuilder {
	return &SqlBuilder{
		installDataset: goqu.Insert(table).Rows(rows...),
	}
}
func NewCountBuilder(table string, fields ...string) *SqlBuilder {
	var c = "1"
	if len(fields) == 1 {
		c = fields[0]
	}
	return &SqlBuilder{
		selectDataset: goqu.From(table).Select(goqu.COUNT(goqu.L(c))),
	}
}
func NewSumBuilder(table string, field string) *SqlBuilder {
	return &SqlBuilder{
		selectDataset: goqu.From(table).Select(goqu.SUM(goqu.I(field))),
	}
}
func (s *SqlBuilder) Or(ex ...goqu.Expression) *SqlBuilder {
	s.ex = append(s.ex, goqu.Or(ex...))
	return s
}
func (s *SqlBuilder) And(ex ...goqu.Expression) *SqlBuilder {
	s.ex = append(s.ex, goqu.And(ex...))
	return s
}
func (s *SqlBuilder) In(field string, params ...interface{}) *SqlBuilder {
	if len(params) == 0 {
		return s
	}

	var val []interface{}
	for _, param := range params {
		ps, ok := slicex.CreateAnyTypeSlice(param)
		if ok {
			val = append(val, ps...)
			continue
		}
		val = append(val, param)
	}

	s.ex = append(s.ex, goqu.C(field).In(val...))
	return s
}
func (s *SqlBuilder) NotIn(field string, params interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).NotIn(params))
	return s
}
func (s *SqlBuilder) Gt(field string, params interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).Gt(params))
	return s
}
func (s *SqlBuilder) Gte(field string, params interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).Gte(params))
	return s
}
func (s *SqlBuilder) Lt(field string, params interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).Lt(params))
	return s
}
func (s *SqlBuilder) Lte(field string, params interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).Lte(params))
	return s
}
func (s *SqlBuilder) Eq(field string, params interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).Eq(params))
	return s
}
func (s *SqlBuilder) Neq(field string, params interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).Neq(params))
	return s
}
func (s *SqlBuilder) Between(field string, start, stop interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).Between(goqu.Range(start, stop)))
	return s
}
func (s *SqlBuilder) NotBetween(field string, start, stop interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).NotBetween(goqu.Range(start, stop)))
	return s
}

// Like 两边模糊匹配
func (s *SqlBuilder) Like(field string, param interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).Like("%"+fmt.Sprintf("%v", param)+"%"))
	return s
}

// LLike 左边模糊匹配
func (s *SqlBuilder) LLike(field string, param interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).Like("%"+fmt.Sprintf("%v", param)))
	return s
}

// RLike 右边模糊匹配
func (s *SqlBuilder) RLike(field string, param interface{}) *SqlBuilder {
	s.ex = append(s.ex, goqu.C(field).Like(fmt.Sprintf("%v", param)+"%"))
	return s
}

func (s *SqlBuilder) Limit(l uint64) *SqlBuilder {
	s.limit = l
	return s
}
func (s *SqlBuilder) Page(l uint64) *SqlBuilder {
	s.page = l
	return s
}
func (s *SqlBuilder) ASC(fields ...string) *SqlBuilder {
	s.asc = append(s.asc, fields...)
	return s
}
func (s *SqlBuilder) Desc(fields ...string) *SqlBuilder {
	s.desc = append(s.desc, fields...)
	return s
}

// GroupBy 一般配合Count或者Sum使用
func (s *SqlBuilder) GroupBy(fields ...string) *SqlBuilder {
	s.groupBy = append(s.groupBy, fields...)
	return s
}
func (s *SqlBuilder) Having(expressions ...goqu.Expression) *SqlBuilder {
	s.having = append(s.having, expressions...)
	return s
}

func (s *SqlBuilder) CopyWhere() []goqu.Expression {
	return s.ex
}

func (s *SqlBuilder) WithWhere(exs []goqu.Expression) *SqlBuilder {
	s.ex = append(s.ex, exs...)
	return s
}

func (s *SqlBuilder) GetSql() string {
	sql, _, _ := s.ToSql()
	return sql
}

func (s *SqlBuilder) ToSql() (sql string, params []interface{}, err error) {
	// 查询语句
	if s.selectDataset != nil {
		se := s.selectDataset.Where(s.ex...)
		if s.tableName != "" {
			se = se.From(s.tableName)
		}
		if s.limit != 0 {
			se = se.Limit(uint(s.limit))
		}
		if s.page != 0 {
			if s.limit == 0 {
				s.limit = 20
			}
			s.page -= 1
			se = se.Offset(uint(s.page * s.limit))
		}
		for _, s2 := range s.asc {
			se = se.Order(goqu.C(s2).Asc())
		}
		for _, s2 := range s.desc {
			se = se.Order(goqu.C(s2).Desc())
		}
		for _, s2 := range s.groupBy {
			se = se.GroupBy(s2)
		}
		for _, s2 := range s.having {
			se = se.Having(s2)
		}
		sql, params, err = se.ToSQL()
		return Replace(sql), params, err

		// 删除语句
	} else if s.deleteDataset != nil {
		se := s.deleteDataset.Where(s.ex...)
		if s.tableName != "" {
			se = se.From(s.tableName)
		}
		if s.limit != 0 {
			se = se.Limit(uint(s.limit))
		}
		sql, params, err = se.ToSQL()
		return Replace(sql), params, err

		// 更新语句
	} else if s.updateDataset != nil {
		se := s.updateDataset.Where(s.ex...)
		if s.tableName != "" {
			se = se.From(s.tableName)
		}
		if s.limit != 0 {
			se = se.Limit(uint(s.limit))
		}
		sql, params, err = se.ToSQL()
		return Replace(sql), params, err

		// 插入语句
	} else if s.installDataset != nil {
		sql, params, err = s.installDataset.ToSQL()
		return Replace(sql), params, err
	}

	return "", nil, errors.New("No SqlBuilder Object")
}

func Replace(query string) string {
	s1 := strings.ReplaceAll(query, "\"`", "`")
	s2 := strings.ReplaceAll(s1, "`\"", "`")
	s3 := strings.ReplaceAll(s2, "\"", "`")
	return s3
}
