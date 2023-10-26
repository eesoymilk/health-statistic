// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/healthkit"
	"github.com/eesoymilk/health-statistic-api/ent/hkdata"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// HealthKitCreate is the builder for creating a HealthKit entity.
type HealthKitCreate struct {
	config
	mutation *HealthKitMutation
	hooks    []Hook
}

// SetStartTime sets the "start_time" field.
func (hkc *HealthKitCreate) SetStartTime(t time.Time) *HealthKitCreate {
	hkc.mutation.SetStartTime(t)
	return hkc
}

// SetEndTime sets the "end_time" field.
func (hkc *HealthKitCreate) SetEndTime(t time.Time) *HealthKitCreate {
	hkc.mutation.SetEndTime(t)
	return hkc
}

// SetID sets the "id" field.
func (hkc *HealthKitCreate) SetID(u uuid.UUID) *HealthKitCreate {
	hkc.mutation.SetID(u)
	return hkc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (hkc *HealthKitCreate) SetNillableID(u *uuid.UUID) *HealthKitCreate {
	if u != nil {
		hkc.SetID(*u)
	}
	return hkc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (hkc *HealthKitCreate) SetUserID(id string) *HealthKitCreate {
	hkc.mutation.SetUserID(id)
	return hkc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (hkc *HealthKitCreate) SetNillableUserID(id *string) *HealthKitCreate {
	if id != nil {
		hkc = hkc.SetUserID(*id)
	}
	return hkc
}

// SetUser sets the "user" edge to the User entity.
func (hkc *HealthKitCreate) SetUser(u *User) *HealthKitCreate {
	return hkc.SetUserID(u.ID)
}

// AddDatumIDs adds the "data" edge to the HKData entity by IDs.
func (hkc *HealthKitCreate) AddDatumIDs(ids ...string) *HealthKitCreate {
	hkc.mutation.AddDatumIDs(ids...)
	return hkc
}

// AddData adds the "data" edges to the HKData entity.
func (hkc *HealthKitCreate) AddData(h ...*HKData) *HealthKitCreate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hkc.AddDatumIDs(ids...)
}

// Mutation returns the HealthKitMutation object of the builder.
func (hkc *HealthKitCreate) Mutation() *HealthKitMutation {
	return hkc.mutation
}

// Save creates the HealthKit in the database.
func (hkc *HealthKitCreate) Save(ctx context.Context) (*HealthKit, error) {
	hkc.defaults()
	return withHooks(ctx, hkc.sqlSave, hkc.mutation, hkc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hkc *HealthKitCreate) SaveX(ctx context.Context) *HealthKit {
	v, err := hkc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hkc *HealthKitCreate) Exec(ctx context.Context) error {
	_, err := hkc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hkc *HealthKitCreate) ExecX(ctx context.Context) {
	if err := hkc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hkc *HealthKitCreate) defaults() {
	if _, ok := hkc.mutation.ID(); !ok {
		v := healthkit.DefaultID()
		hkc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hkc *HealthKitCreate) check() error {
	if _, ok := hkc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "HealthKit.start_time"`)}
	}
	if _, ok := hkc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`ent: missing required field "HealthKit.end_time"`)}
	}
	return nil
}

func (hkc *HealthKitCreate) sqlSave(ctx context.Context) (*HealthKit, error) {
	if err := hkc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hkc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hkc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	hkc.mutation.id = &_node.ID
	hkc.mutation.done = true
	return _node, nil
}

func (hkc *HealthKitCreate) createSpec() (*HealthKit, *sqlgraph.CreateSpec) {
	var (
		_node = &HealthKit{config: hkc.config}
		_spec = sqlgraph.NewCreateSpec(healthkit.Table, sqlgraph.NewFieldSpec(healthkit.FieldID, field.TypeUUID))
	)
	if id, ok := hkc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := hkc.mutation.StartTime(); ok {
		_spec.SetField(healthkit.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := hkc.mutation.EndTime(); ok {
		_spec.SetField(healthkit.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if nodes := hkc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   healthkit.UserTable,
			Columns: []string{healthkit.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_healthkit = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hkc.mutation.DataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   healthkit.DataTable,
			Columns: []string{healthkit.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hkdata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HealthKitCreateBulk is the builder for creating many HealthKit entities in bulk.
type HealthKitCreateBulk struct {
	config
	builders []*HealthKitCreate
}

// Save creates the HealthKit entities in the database.
func (hkcb *HealthKitCreateBulk) Save(ctx context.Context) ([]*HealthKit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hkcb.builders))
	nodes := make([]*HealthKit, len(hkcb.builders))
	mutators := make([]Mutator, len(hkcb.builders))
	for i := range hkcb.builders {
		func(i int, root context.Context) {
			builder := hkcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HealthKitMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hkcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hkcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hkcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hkcb *HealthKitCreateBulk) SaveX(ctx context.Context) []*HealthKit {
	v, err := hkcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hkcb *HealthKitCreateBulk) Exec(ctx context.Context) error {
	_, err := hkcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hkcb *HealthKitCreateBulk) ExecX(ctx context.Context) {
	if err := hkcb.Exec(ctx); err != nil {
		panic(err)
	}
}