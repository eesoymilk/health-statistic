// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/healthkit"
	"github.com/eesoymilk/health-statistic-api/ent/hkdata"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/google/uuid"
)

// HKDataUpdate is the builder for updating HKData entities.
type HKDataUpdate struct {
	config
	hooks    []Hook
	mutation *HKDataMutation
}

// Where appends a list predicates to the HKDataUpdate builder.
func (hdu *HKDataUpdate) Where(ps ...predicate.HKData) *HKDataUpdate {
	hdu.mutation.Where(ps...)
	return hdu
}

// SetType sets the "type" field.
func (hdu *HKDataUpdate) SetType(s string) *HKDataUpdate {
	hdu.mutation.SetType(s)
	return hdu
}

// SetValue sets the "value" field.
func (hdu *HKDataUpdate) SetValue(s string) *HKDataUpdate {
	hdu.mutation.SetValue(s)
	return hdu
}

// SetStartTimestamp sets the "start_timestamp" field.
func (hdu *HKDataUpdate) SetStartTimestamp(s string) *HKDataUpdate {
	hdu.mutation.SetStartTimestamp(s)
	return hdu
}

// SetEndTimestamp sets the "end_timestamp" field.
func (hdu *HKDataUpdate) SetEndTimestamp(s string) *HKDataUpdate {
	hdu.mutation.SetEndTimestamp(s)
	return hdu
}

// SetTimezoneID sets the "timezone_id" field.
func (hdu *HKDataUpdate) SetTimezoneID(s string) *HKDataUpdate {
	hdu.mutation.SetTimezoneID(s)
	return hdu
}

// SetHealthkitID sets the "healthkit" edge to the HealthKit entity by ID.
func (hdu *HKDataUpdate) SetHealthkitID(id uuid.UUID) *HKDataUpdate {
	hdu.mutation.SetHealthkitID(id)
	return hdu
}

// SetNillableHealthkitID sets the "healthkit" edge to the HealthKit entity by ID if the given value is not nil.
func (hdu *HKDataUpdate) SetNillableHealthkitID(id *uuid.UUID) *HKDataUpdate {
	if id != nil {
		hdu = hdu.SetHealthkitID(*id)
	}
	return hdu
}

// SetHealthkit sets the "healthkit" edge to the HealthKit entity.
func (hdu *HKDataUpdate) SetHealthkit(h *HealthKit) *HKDataUpdate {
	return hdu.SetHealthkitID(h.ID)
}

// Mutation returns the HKDataMutation object of the builder.
func (hdu *HKDataUpdate) Mutation() *HKDataMutation {
	return hdu.mutation
}

// ClearHealthkit clears the "healthkit" edge to the HealthKit entity.
func (hdu *HKDataUpdate) ClearHealthkit() *HKDataUpdate {
	hdu.mutation.ClearHealthkit()
	return hdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hdu *HKDataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hdu.sqlSave, hdu.mutation, hdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hdu *HKDataUpdate) SaveX(ctx context.Context) int {
	affected, err := hdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hdu *HKDataUpdate) Exec(ctx context.Context) error {
	_, err := hdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hdu *HKDataUpdate) ExecX(ctx context.Context) {
	if err := hdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hdu *HKDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(hkdata.Table, hkdata.Columns, sqlgraph.NewFieldSpec(hkdata.FieldID, field.TypeUUID))
	if ps := hdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hdu.mutation.GetType(); ok {
		_spec.SetField(hkdata.FieldType, field.TypeString, value)
	}
	if value, ok := hdu.mutation.Value(); ok {
		_spec.SetField(hkdata.FieldValue, field.TypeString, value)
	}
	if value, ok := hdu.mutation.StartTimestamp(); ok {
		_spec.SetField(hkdata.FieldStartTimestamp, field.TypeString, value)
	}
	if value, ok := hdu.mutation.EndTimestamp(); ok {
		_spec.SetField(hkdata.FieldEndTimestamp, field.TypeString, value)
	}
	if value, ok := hdu.mutation.TimezoneID(); ok {
		_spec.SetField(hkdata.FieldTimezoneID, field.TypeString, value)
	}
	if hdu.mutation.HealthkitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hkdata.HealthkitTable,
			Columns: []string{hkdata.HealthkitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(healthkit.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hdu.mutation.HealthkitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hkdata.HealthkitTable,
			Columns: []string{hkdata.HealthkitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(healthkit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hkdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hdu.mutation.done = true
	return n, nil
}

// HKDataUpdateOne is the builder for updating a single HKData entity.
type HKDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HKDataMutation
}

// SetType sets the "type" field.
func (hduo *HKDataUpdateOne) SetType(s string) *HKDataUpdateOne {
	hduo.mutation.SetType(s)
	return hduo
}

// SetValue sets the "value" field.
func (hduo *HKDataUpdateOne) SetValue(s string) *HKDataUpdateOne {
	hduo.mutation.SetValue(s)
	return hduo
}

// SetStartTimestamp sets the "start_timestamp" field.
func (hduo *HKDataUpdateOne) SetStartTimestamp(s string) *HKDataUpdateOne {
	hduo.mutation.SetStartTimestamp(s)
	return hduo
}

// SetEndTimestamp sets the "end_timestamp" field.
func (hduo *HKDataUpdateOne) SetEndTimestamp(s string) *HKDataUpdateOne {
	hduo.mutation.SetEndTimestamp(s)
	return hduo
}

// SetTimezoneID sets the "timezone_id" field.
func (hduo *HKDataUpdateOne) SetTimezoneID(s string) *HKDataUpdateOne {
	hduo.mutation.SetTimezoneID(s)
	return hduo
}

// SetHealthkitID sets the "healthkit" edge to the HealthKit entity by ID.
func (hduo *HKDataUpdateOne) SetHealthkitID(id uuid.UUID) *HKDataUpdateOne {
	hduo.mutation.SetHealthkitID(id)
	return hduo
}

// SetNillableHealthkitID sets the "healthkit" edge to the HealthKit entity by ID if the given value is not nil.
func (hduo *HKDataUpdateOne) SetNillableHealthkitID(id *uuid.UUID) *HKDataUpdateOne {
	if id != nil {
		hduo = hduo.SetHealthkitID(*id)
	}
	return hduo
}

// SetHealthkit sets the "healthkit" edge to the HealthKit entity.
func (hduo *HKDataUpdateOne) SetHealthkit(h *HealthKit) *HKDataUpdateOne {
	return hduo.SetHealthkitID(h.ID)
}

// Mutation returns the HKDataMutation object of the builder.
func (hduo *HKDataUpdateOne) Mutation() *HKDataMutation {
	return hduo.mutation
}

// ClearHealthkit clears the "healthkit" edge to the HealthKit entity.
func (hduo *HKDataUpdateOne) ClearHealthkit() *HKDataUpdateOne {
	hduo.mutation.ClearHealthkit()
	return hduo
}

// Where appends a list predicates to the HKDataUpdate builder.
func (hduo *HKDataUpdateOne) Where(ps ...predicate.HKData) *HKDataUpdateOne {
	hduo.mutation.Where(ps...)
	return hduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hduo *HKDataUpdateOne) Select(field string, fields ...string) *HKDataUpdateOne {
	hduo.fields = append([]string{field}, fields...)
	return hduo
}

// Save executes the query and returns the updated HKData entity.
func (hduo *HKDataUpdateOne) Save(ctx context.Context) (*HKData, error) {
	return withHooks(ctx, hduo.sqlSave, hduo.mutation, hduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hduo *HKDataUpdateOne) SaveX(ctx context.Context) *HKData {
	node, err := hduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hduo *HKDataUpdateOne) Exec(ctx context.Context) error {
	_, err := hduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hduo *HKDataUpdateOne) ExecX(ctx context.Context) {
	if err := hduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hduo *HKDataUpdateOne) sqlSave(ctx context.Context) (_node *HKData, err error) {
	_spec := sqlgraph.NewUpdateSpec(hkdata.Table, hkdata.Columns, sqlgraph.NewFieldSpec(hkdata.FieldID, field.TypeUUID))
	id, ok := hduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HKData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hkdata.FieldID)
		for _, f := range fields {
			if !hkdata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hkdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hduo.mutation.GetType(); ok {
		_spec.SetField(hkdata.FieldType, field.TypeString, value)
	}
	if value, ok := hduo.mutation.Value(); ok {
		_spec.SetField(hkdata.FieldValue, field.TypeString, value)
	}
	if value, ok := hduo.mutation.StartTimestamp(); ok {
		_spec.SetField(hkdata.FieldStartTimestamp, field.TypeString, value)
	}
	if value, ok := hduo.mutation.EndTimestamp(); ok {
		_spec.SetField(hkdata.FieldEndTimestamp, field.TypeString, value)
	}
	if value, ok := hduo.mutation.TimezoneID(); ok {
		_spec.SetField(hkdata.FieldTimezoneID, field.TypeString, value)
	}
	if hduo.mutation.HealthkitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hkdata.HealthkitTable,
			Columns: []string{hkdata.HealthkitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(healthkit.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hduo.mutation.HealthkitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hkdata.HealthkitTable,
			Columns: []string{hkdata.HealthkitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(healthkit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HKData{config: hduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hkdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	hduo.mutation.done = true
	return _node, nil
}
