// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package mycard

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.MyCard {
	return predicate.MyCard(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.MyCard {
	return predicate.MyCard(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.MyCard {
	return predicate.MyCard(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.MyCard {
	return predicate.MyCard(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.MyCard {
	return predicate.MyCard(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.MyCard {
	return predicate.MyCard(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.MyCard {
	return predicate.MyCard(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.MyCard {
	return predicate.MyCard(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.MyCard {
	return predicate.MyCard(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.MyCard {
	return predicate.MyCard(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.MyCard {
	return predicate.MyCard(sql.FieldContainsFold(FieldID, id))
}

// CardPassword applies equality check predicate on the "card_password" field. It's identical to CardPasswordEQ.
func CardPassword(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldEQ(FieldCardPassword, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldEQ(FieldCreatedAt, v))
}

// TakenAt applies equality check predicate on the "taken_at" field. It's identical to TakenAtEQ.
func TakenAt(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldEQ(FieldTakenAt, v))
}

// CardPasswordEQ applies the EQ predicate on the "card_password" field.
func CardPasswordEQ(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldEQ(FieldCardPassword, v))
}

// CardPasswordNEQ applies the NEQ predicate on the "card_password" field.
func CardPasswordNEQ(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldNEQ(FieldCardPassword, v))
}

// CardPasswordIn applies the In predicate on the "card_password" field.
func CardPasswordIn(vs ...string) predicate.MyCard {
	return predicate.MyCard(sql.FieldIn(FieldCardPassword, vs...))
}

// CardPasswordNotIn applies the NotIn predicate on the "card_password" field.
func CardPasswordNotIn(vs ...string) predicate.MyCard {
	return predicate.MyCard(sql.FieldNotIn(FieldCardPassword, vs...))
}

// CardPasswordGT applies the GT predicate on the "card_password" field.
func CardPasswordGT(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldGT(FieldCardPassword, v))
}

// CardPasswordGTE applies the GTE predicate on the "card_password" field.
func CardPasswordGTE(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldGTE(FieldCardPassword, v))
}

// CardPasswordLT applies the LT predicate on the "card_password" field.
func CardPasswordLT(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldLT(FieldCardPassword, v))
}

// CardPasswordLTE applies the LTE predicate on the "card_password" field.
func CardPasswordLTE(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldLTE(FieldCardPassword, v))
}

// CardPasswordContains applies the Contains predicate on the "card_password" field.
func CardPasswordContains(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldContains(FieldCardPassword, v))
}

// CardPasswordHasPrefix applies the HasPrefix predicate on the "card_password" field.
func CardPasswordHasPrefix(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldHasPrefix(FieldCardPassword, v))
}

// CardPasswordHasSuffix applies the HasSuffix predicate on the "card_password" field.
func CardPasswordHasSuffix(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldHasSuffix(FieldCardPassword, v))
}

// CardPasswordEqualFold applies the EqualFold predicate on the "card_password" field.
func CardPasswordEqualFold(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldEqualFold(FieldCardPassword, v))
}

// CardPasswordContainsFold applies the ContainsFold predicate on the "card_password" field.
func CardPasswordContainsFold(v string) predicate.MyCard {
	return predicate.MyCard(sql.FieldContainsFold(FieldCardPassword, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldLTE(FieldCreatedAt, v))
}

// TakenAtEQ applies the EQ predicate on the "taken_at" field.
func TakenAtEQ(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldEQ(FieldTakenAt, v))
}

// TakenAtNEQ applies the NEQ predicate on the "taken_at" field.
func TakenAtNEQ(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldNEQ(FieldTakenAt, v))
}

// TakenAtIn applies the In predicate on the "taken_at" field.
func TakenAtIn(vs ...time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldIn(FieldTakenAt, vs...))
}

// TakenAtNotIn applies the NotIn predicate on the "taken_at" field.
func TakenAtNotIn(vs ...time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldNotIn(FieldTakenAt, vs...))
}

// TakenAtGT applies the GT predicate on the "taken_at" field.
func TakenAtGT(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldGT(FieldTakenAt, v))
}

// TakenAtGTE applies the GTE predicate on the "taken_at" field.
func TakenAtGTE(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldGTE(FieldTakenAt, v))
}

// TakenAtLT applies the LT predicate on the "taken_at" field.
func TakenAtLT(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldLT(FieldTakenAt, v))
}

// TakenAtLTE applies the LTE predicate on the "taken_at" field.
func TakenAtLTE(v time.Time) predicate.MyCard {
	return predicate.MyCard(sql.FieldLTE(FieldTakenAt, v))
}

// TakenAtIsNil applies the IsNil predicate on the "taken_at" field.
func TakenAtIsNil() predicate.MyCard {
	return predicate.MyCard(sql.FieldIsNull(FieldTakenAt))
}

// TakenAtNotNil applies the NotNil predicate on the "taken_at" field.
func TakenAtNotNil() predicate.MyCard {
	return predicate.MyCard(sql.FieldNotNull(FieldTakenAt))
}

// HasRecipient applies the HasEdge predicate on the "recipient" edge.
func HasRecipient() predicate.MyCard {
	return predicate.MyCard(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RecipientTable, RecipientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRecipientWith applies the HasEdge predicate on the "recipient" edge with a given conditions (other predicates).
func HasRecipientWith(preds ...predicate.User) predicate.MyCard {
	return predicate.MyCard(func(s *sql.Selector) {
		step := newRecipientStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotifications applies the HasEdge predicate on the "notifications" edge.
func HasNotifications() predicate.MyCard {
	return predicate.MyCard(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotificationsWith applies the HasEdge predicate on the "notifications" edge with a given conditions (other predicates).
func HasNotificationsWith(preds ...predicate.Notification) predicate.MyCard {
	return predicate.MyCard(func(s *sql.Selector) {
		step := newNotificationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MyCard) predicate.MyCard {
	return predicate.MyCard(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MyCard) predicate.MyCard {
	return predicate.MyCard(func(s *sql.Selector) {
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
func Not(p predicate.MyCard) predicate.MyCard {
	return predicate.MyCard(func(s *sql.Selector) {
		p(s.Not())
	})
}