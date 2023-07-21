// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/mycard"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
)

// MyCardDelete is the builder for deleting a MyCard entity.
type MyCardDelete struct {
	config
	hooks    []Hook
	mutation *MyCardMutation
}

// Where appends a list predicates to the MyCardDelete builder.
func (mcd *MyCardDelete) Where(ps ...predicate.MyCard) *MyCardDelete {
	mcd.mutation.Where(ps...)
	return mcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mcd *MyCardDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mcd.sqlExec, mcd.mutation, mcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mcd *MyCardDelete) ExecX(ctx context.Context) int {
	n, err := mcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mcd *MyCardDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(mycard.Table, sqlgraph.NewFieldSpec(mycard.FieldID, field.TypeInt))
	if ps := mcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mcd.mutation.done = true
	return affected, err
}

// MyCardDeleteOne is the builder for deleting a single MyCard entity.
type MyCardDeleteOne struct {
	mcd *MyCardDelete
}

// Where appends a list predicates to the MyCardDelete builder.
func (mcdo *MyCardDeleteOne) Where(ps ...predicate.MyCard) *MyCardDeleteOne {
	mcdo.mcd.mutation.Where(ps...)
	return mcdo
}

// Exec executes the deletion query.
func (mcdo *MyCardDeleteOne) Exec(ctx context.Context) error {
	n, err := mcdo.mcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mycard.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mcdo *MyCardDeleteOne) ExecX(ctx context.Context) {
	if err := mcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
