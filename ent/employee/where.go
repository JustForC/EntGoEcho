// Code generated by entc, DO NOT EDIT.

package employee

import (
	"CompanyAPI/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Salary applies equality check predicate on the "salary" field. It's identical to SalaryEQ.
func Salary(v int32) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalary), v))
	})
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SalaryEQ applies the EQ predicate on the "salary" field.
func SalaryEQ(v int32) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalary), v))
	})
}

// SalaryNEQ applies the NEQ predicate on the "salary" field.
func SalaryNEQ(v int32) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSalary), v))
	})
}

// SalaryIn applies the In predicate on the "salary" field.
func SalaryIn(vs ...int32) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSalary), v...))
	})
}

// SalaryNotIn applies the NotIn predicate on the "salary" field.
func SalaryNotIn(vs ...int32) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSalary), v...))
	})
}

// SalaryGT applies the GT predicate on the "salary" field.
func SalaryGT(v int32) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSalary), v))
	})
}

// SalaryGTE applies the GTE predicate on the "salary" field.
func SalaryGTE(v int32) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSalary), v))
	})
}

// SalaryLT applies the LT predicate on the "salary" field.
func SalaryLT(v int32) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSalary), v))
	})
}

// SalaryLTE applies the LTE predicate on the "salary" field.
func SalaryLTE(v int32) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSalary), v))
	})
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPosition), v))
	})
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPosition), v...))
	})
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPosition), v...))
	})
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPosition), v))
	})
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPosition), v))
	})
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPosition), v))
	})
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPosition), v))
	})
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPosition), v))
	})
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPosition), v))
	})
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPosition), v))
	})
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPosition), v))
	})
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPosition), v))
	})
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CompanyTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CompanyInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		p(s.Not())
	})
}