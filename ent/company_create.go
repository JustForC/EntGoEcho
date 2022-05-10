// Code generated by entc, DO NOT EDIT.

package ent

import (
	"CompanyAPI/ent/company"
	"CompanyAPI/ent/employee"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompanyCreate is the builder for creating a Company entity.
type CompanyCreate struct {
	config
	mutation *CompanyMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CompanyCreate) SetName(s string) *CompanyCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetService sets the "service" field.
func (cc *CompanyCreate) SetService(s string) *CompanyCreate {
	cc.mutation.SetService(s)
	return cc
}

// SetAddress sets the "address" field.
func (cc *CompanyCreate) SetAddress(s string) *CompanyCreate {
	cc.mutation.SetAddress(s)
	return cc
}

// AddEmployeeIDs adds the "employees" edge to the Employee entity by IDs.
func (cc *CompanyCreate) AddEmployeeIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddEmployeeIDs(ids...)
	return cc
}

// AddEmployees adds the "employees" edges to the Employee entity.
func (cc *CompanyCreate) AddEmployees(e ...*Employee) *CompanyCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cc.AddEmployeeIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cc *CompanyCreate) Mutation() *CompanyMutation {
	return cc.mutation
}

// Save creates the Company in the database.
func (cc *CompanyCreate) Save(ctx context.Context) (*Company, error) {
	var (
		err  error
		node *Company
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompanyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CompanyCreate) SaveX(ctx context.Context) *Company {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CompanyCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CompanyCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CompanyCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Company.name"`)}
	}
	if _, ok := cc.mutation.Service(); !ok {
		return &ValidationError{Name: "service", err: errors.New(`ent: missing required field "Company.service"`)}
	}
	if _, ok := cc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Company.address"`)}
	}
	return nil
}

func (cc *CompanyCreate) sqlSave(ctx context.Context) (*Company, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CompanyCreate) createSpec() (*Company, *sqlgraph.CreateSpec) {
	var (
		_node = &Company{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: company.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: company.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Service(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldService,
		})
		_node.Service = value
	}
	if value, ok := cc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldAddress,
		})
		_node.Address = value
	}
	if nodes := cc.mutation.EmployeesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CompanyCreateBulk is the builder for creating many Company entities in bulk.
type CompanyCreateBulk struct {
	config
	builders []*CompanyCreate
}

// Save creates the Company entities in the database.
func (ccb *CompanyCreateBulk) Save(ctx context.Context) ([]*Company, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Company, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompanyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CompanyCreateBulk) SaveX(ctx context.Context) []*Company {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CompanyCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CompanyCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}