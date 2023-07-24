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
	"github.com/eesoymilk/health-statistic-api/ent/mycard"
	"github.com/eesoymilk/health-statistic-api/ent/notification"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// MyCardCreate is the builder for creating a MyCard entity.
type MyCardCreate struct {
	config
	mutation *MyCardMutation
	hooks    []Hook
}

// SetCardPassword sets the "card_password" field.
func (mcc *MyCardCreate) SetCardPassword(s string) *MyCardCreate {
	mcc.mutation.SetCardPassword(s)
	return mcc
}

// SetCreatedAt sets the "created_at" field.
func (mcc *MyCardCreate) SetCreatedAt(t time.Time) *MyCardCreate {
	mcc.mutation.SetCreatedAt(t)
	return mcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcc *MyCardCreate) SetNillableCreatedAt(t *time.Time) *MyCardCreate {
	if t != nil {
		mcc.SetCreatedAt(*t)
	}
	return mcc
}

// SetTakenAt sets the "taken_at" field.
func (mcc *MyCardCreate) SetTakenAt(t time.Time) *MyCardCreate {
	mcc.mutation.SetTakenAt(t)
	return mcc
}

// SetNillableTakenAt sets the "taken_at" field if the given value is not nil.
func (mcc *MyCardCreate) SetNillableTakenAt(t *time.Time) *MyCardCreate {
	if t != nil {
		mcc.SetTakenAt(*t)
	}
	return mcc
}

// SetID sets the "id" field.
func (mcc *MyCardCreate) SetID(s string) *MyCardCreate {
	mcc.mutation.SetID(s)
	return mcc
}

// SetRecipientID sets the "recipient" edge to the User entity by ID.
func (mcc *MyCardCreate) SetRecipientID(id string) *MyCardCreate {
	mcc.mutation.SetRecipientID(id)
	return mcc
}

// SetNillableRecipientID sets the "recipient" edge to the User entity by ID if the given value is not nil.
func (mcc *MyCardCreate) SetNillableRecipientID(id *string) *MyCardCreate {
	if id != nil {
		mcc = mcc.SetRecipientID(*id)
	}
	return mcc
}

// SetRecipient sets the "recipient" edge to the User entity.
func (mcc *MyCardCreate) SetRecipient(u *User) *MyCardCreate {
	return mcc.SetRecipientID(u.ID)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (mcc *MyCardCreate) AddNotificationIDs(ids ...uuid.UUID) *MyCardCreate {
	mcc.mutation.AddNotificationIDs(ids...)
	return mcc
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (mcc *MyCardCreate) AddNotifications(n ...*Notification) *MyCardCreate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return mcc.AddNotificationIDs(ids...)
}

// Mutation returns the MyCardMutation object of the builder.
func (mcc *MyCardCreate) Mutation() *MyCardMutation {
	return mcc.mutation
}

// Save creates the MyCard in the database.
func (mcc *MyCardCreate) Save(ctx context.Context) (*MyCard, error) {
	mcc.defaults()
	return withHooks(ctx, mcc.sqlSave, mcc.mutation, mcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mcc *MyCardCreate) SaveX(ctx context.Context) *MyCard {
	v, err := mcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcc *MyCardCreate) Exec(ctx context.Context) error {
	_, err := mcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcc *MyCardCreate) ExecX(ctx context.Context) {
	if err := mcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcc *MyCardCreate) defaults() {
	if _, ok := mcc.mutation.CreatedAt(); !ok {
		v := mycard.DefaultCreatedAt()
		mcc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcc *MyCardCreate) check() error {
	if _, ok := mcc.mutation.CardPassword(); !ok {
		return &ValidationError{Name: "card_password", err: errors.New(`ent: missing required field "MyCard.card_password"`)}
	}
	if v, ok := mcc.mutation.CardPassword(); ok {
		if err := mycard.CardPasswordValidator(v); err != nil {
			return &ValidationError{Name: "card_password", err: fmt.Errorf(`ent: validator failed for field "MyCard.card_password": %w`, err)}
		}
	}
	if _, ok := mcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MyCard.created_at"`)}
	}
	if v, ok := mcc.mutation.ID(); ok {
		if err := mycard.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "MyCard.id": %w`, err)}
		}
	}
	return nil
}

func (mcc *MyCardCreate) sqlSave(ctx context.Context) (*MyCard, error) {
	if err := mcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected MyCard.ID type: %T", _spec.ID.Value)
		}
	}
	mcc.mutation.id = &_node.ID
	mcc.mutation.done = true
	return _node, nil
}

func (mcc *MyCardCreate) createSpec() (*MyCard, *sqlgraph.CreateSpec) {
	var (
		_node = &MyCard{config: mcc.config}
		_spec = sqlgraph.NewCreateSpec(mycard.Table, sqlgraph.NewFieldSpec(mycard.FieldID, field.TypeString))
	)
	if id, ok := mcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mcc.mutation.CardPassword(); ok {
		_spec.SetField(mycard.FieldCardPassword, field.TypeString, value)
		_node.CardPassword = value
	}
	if value, ok := mcc.mutation.CreatedAt(); ok {
		_spec.SetField(mycard.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mcc.mutation.TakenAt(); ok {
		_spec.SetField(mycard.FieldTakenAt, field.TypeTime, value)
		_node.TakenAt = value
	}
	if nodes := mcc.mutation.RecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   mycard.RecipientTable,
			Columns: []string{mycard.RecipientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.my_card_recipient = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mcc.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mycard.NotificationsTable,
			Columns: []string{mycard.NotificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MyCardCreateBulk is the builder for creating many MyCard entities in bulk.
type MyCardCreateBulk struct {
	config
	builders []*MyCardCreate
}

// Save creates the MyCard entities in the database.
func (mccb *MyCardCreateBulk) Save(ctx context.Context) ([]*MyCard, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mccb.builders))
	nodes := make([]*MyCard, len(mccb.builders))
	mutators := make([]Mutator, len(mccb.builders))
	for i := range mccb.builders {
		func(i int, root context.Context) {
			builder := mccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MyCardMutation)
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
					_, err = mutators[i+1].Mutate(root, mccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mccb *MyCardCreateBulk) SaveX(ctx context.Context) []*MyCard {
	v, err := mccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mccb *MyCardCreateBulk) Exec(ctx context.Context) error {
	_, err := mccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mccb *MyCardCreateBulk) ExecX(ctx context.Context) {
	if err := mccb.Exec(ctx); err != nil {
		panic(err)
	}
}
