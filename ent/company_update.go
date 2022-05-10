// Code generated by entc, DO NOT EDIT.

package ent

import (
	"CompanyAPI/ent/company"
	"CompanyAPI/ent/employee"
	"CompanyAPI/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompanyUpdate is the builder for updating Company entities.
type CompanyUpdate struct {
	config
	hooks    []Hook
	mutation *CompanyMutation
}

// Where appends a list predicates to the CompanyUpdate builder.
func (cu *CompanyUpdate) Where(ps ...predicate.Company) *CompanyUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CompanyUpdate) SetName(s string) *CompanyUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetService sets the "service" field.
func (cu *CompanyUpdate) SetService(s string) *CompanyUpdate {
	cu.mutation.SetService(s)
	return cu
}

// SetAddress sets the "address" field.
func (cu *CompanyUpdate) SetAddress(s string) *CompanyUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// AddEmployeeIDs adds the "employees" edge to the Employee entity by IDs.
func (cu *CompanyUpdate) AddEmployeeIDs(ids ...int) *CompanyUpdate {
	cu.mutation.AddEmployeeIDs(ids...)
	return cu
}

// AddEmployees adds the "employees" edges to the Employee entity.
func (cu *CompanyUpdate) AddEmployees(e ...*Employee) *CompanyUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.AddEmployeeIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cu *CompanyUpdate) Mutation() *CompanyMutation {
	return cu.mutation
}

// ClearEmployees clears all "employees" edges to the Employee entity.
func (cu *CompanyUpdate) ClearEmployees() *CompanyUpdate {
	cu.mutation.ClearEmployees()
	return cu
}

// RemoveEmployeeIDs removes the "employees" edge to Employee entities by IDs.
func (cu *CompanyUpdate) RemoveEmployeeIDs(ids ...int) *CompanyUpdate {
	cu.mutation.RemoveEmployeeIDs(ids...)
	return cu
}

// RemoveEmployees removes "employees" edges to Employee entities.
func (cu *CompanyUpdate) RemoveEmployees(e ...*Employee) *CompanyUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.RemoveEmployeeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CompanyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompanyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CompanyUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CompanyUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CompanyUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CompanyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   company.Table,
			Columns: company.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: company.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldName,
		})
	}
	if value, ok := cu.mutation.Service(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldService,
		})
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldAddress,
		})
	}
	if cu.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.EmployeesTable,
			Columns: []string{company.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedEmployeesIDs(); len(nodes) > 0 && !cu.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.EmployeesTable,
			Columns: []string{company.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.EmployeesTable,
			Columns: []string{company.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{company.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CompanyUpdateOne is the builder for updating a single Company entity.
type CompanyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CompanyMutation
}

// SetName sets the "name" field.
func (cuo *CompanyUpdateOne) SetName(s string) *CompanyUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetService sets the "service" field.
func (cuo *CompanyUpdateOne) SetService(s string) *CompanyUpdateOne {
	cuo.mutation.SetService(s)
	return cuo
}

// SetAddress sets the "address" field.
func (cuo *CompanyUpdateOne) SetAddress(s string) *CompanyUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// AddEmployeeIDs adds the "employees" edge to the Employee entity by IDs.
func (cuo *CompanyUpdateOne) AddEmployeeIDs(ids ...int) *CompanyUpdateOne {
	cuo.mutation.AddEmployeeIDs(ids...)
	return cuo
}

// AddEmployees adds the "employees" edges to the Employee entity.
func (cuo *CompanyUpdateOne) AddEmployees(e ...*Employee) *CompanyUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.AddEmployeeIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cuo *CompanyUpdateOne) Mutation() *CompanyMutation {
	return cuo.mutation
}

// ClearEmployees clears all "employees" edges to the Employee entity.
func (cuo *CompanyUpdateOne) ClearEmployees() *CompanyUpdateOne {
	cuo.mutation.ClearEmployees()
	return cuo
}

// RemoveEmployeeIDs removes the "employees" edge to Employee entities by IDs.
func (cuo *CompanyUpdateOne) RemoveEmployeeIDs(ids ...int) *CompanyUpdateOne {
	cuo.mutation.RemoveEmployeeIDs(ids...)
	return cuo
}

// RemoveEmployees removes "employees" edges to Employee entities.
func (cuo *CompanyUpdateOne) RemoveEmployees(e ...*Employee) *CompanyUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.RemoveEmployeeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CompanyUpdateOne) Select(field string, fields ...string) *CompanyUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Company entity.
func (cuo *CompanyUpdateOne) Save(ctx context.Context) (*Company, error) {
	var (
		err  error
		node *Company
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompanyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CompanyUpdateOne) SaveX(ctx context.Context) *Company {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CompanyUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CompanyUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CompanyUpdateOne) sqlSave(ctx context.Context) (_node *Company, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   company.Table,
			Columns: company.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: company.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Company.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, company.FieldID)
		for _, f := range fields {
			if !company.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != company.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldName,
		})
	}
	if value, ok := cuo.mutation.Service(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldService,
		})
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldAddress,
		})
	}
	if cuo.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.EmployeesTable,
			Columns: []string{company.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedEmployeesIDs(); len(nodes) > 0 && !cuo.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.EmployeesTable,
			Columns: []string{company.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.EmployeesTable,
			Columns: []string{company.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Company{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{company.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}