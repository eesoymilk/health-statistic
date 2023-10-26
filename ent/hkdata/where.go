// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package hkdata

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.HKData {
	return predicate.HKData(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.HKData {
	return predicate.HKData(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.HKData {
	return predicate.HKData(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.HKData {
	return predicate.HKData(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.HKData {
	return predicate.HKData(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.HKData {
	return predicate.HKData(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.HKData {
	return predicate.HKData(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.HKData {
	return predicate.HKData(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.HKData {
	return predicate.HKData(sql.FieldContainsFold(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldType, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldValue, v))
}

// StartTimestamp applies equality check predicate on the "start_timestamp" field. It's identical to StartTimestampEQ.
func StartTimestamp(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldStartTimestamp, v))
}

// EndTimestamp applies equality check predicate on the "end_timestamp" field. It's identical to EndTimestampEQ.
func EndTimestamp(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldEndTimestamp, v))
}

// TimezoneID applies equality check predicate on the "timezone_id" field. It's identical to TimezoneIDEQ.
func TimezoneID(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldTimezoneID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.HKData {
	return predicate.HKData(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.HKData {
	return predicate.HKData(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.HKData {
	return predicate.HKData(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.HKData {
	return predicate.HKData(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.HKData {
	return predicate.HKData(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.HKData {
	return predicate.HKData(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.HKData {
	return predicate.HKData(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.HKData {
	return predicate.HKData(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.HKData {
	return predicate.HKData(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.HKData {
	return predicate.HKData(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.HKData {
	return predicate.HKData(sql.FieldContainsFold(FieldType, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.HKData {
	return predicate.HKData(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.HKData {
	return predicate.HKData(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.HKData {
	return predicate.HKData(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.HKData {
	return predicate.HKData(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.HKData {
	return predicate.HKData(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.HKData {
	return predicate.HKData(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.HKData {
	return predicate.HKData(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.HKData {
	return predicate.HKData(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.HKData {
	return predicate.HKData(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.HKData {
	return predicate.HKData(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.HKData {
	return predicate.HKData(sql.FieldContainsFold(FieldValue, v))
}

// StartTimestampEQ applies the EQ predicate on the "start_timestamp" field.
func StartTimestampEQ(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldStartTimestamp, v))
}

// StartTimestampNEQ applies the NEQ predicate on the "start_timestamp" field.
func StartTimestampNEQ(v string) predicate.HKData {
	return predicate.HKData(sql.FieldNEQ(FieldStartTimestamp, v))
}

// StartTimestampIn applies the In predicate on the "start_timestamp" field.
func StartTimestampIn(vs ...string) predicate.HKData {
	return predicate.HKData(sql.FieldIn(FieldStartTimestamp, vs...))
}

// StartTimestampNotIn applies the NotIn predicate on the "start_timestamp" field.
func StartTimestampNotIn(vs ...string) predicate.HKData {
	return predicate.HKData(sql.FieldNotIn(FieldStartTimestamp, vs...))
}

// StartTimestampGT applies the GT predicate on the "start_timestamp" field.
func StartTimestampGT(v string) predicate.HKData {
	return predicate.HKData(sql.FieldGT(FieldStartTimestamp, v))
}

// StartTimestampGTE applies the GTE predicate on the "start_timestamp" field.
func StartTimestampGTE(v string) predicate.HKData {
	return predicate.HKData(sql.FieldGTE(FieldStartTimestamp, v))
}

// StartTimestampLT applies the LT predicate on the "start_timestamp" field.
func StartTimestampLT(v string) predicate.HKData {
	return predicate.HKData(sql.FieldLT(FieldStartTimestamp, v))
}

// StartTimestampLTE applies the LTE predicate on the "start_timestamp" field.
func StartTimestampLTE(v string) predicate.HKData {
	return predicate.HKData(sql.FieldLTE(FieldStartTimestamp, v))
}

// StartTimestampContains applies the Contains predicate on the "start_timestamp" field.
func StartTimestampContains(v string) predicate.HKData {
	return predicate.HKData(sql.FieldContains(FieldStartTimestamp, v))
}

// StartTimestampHasPrefix applies the HasPrefix predicate on the "start_timestamp" field.
func StartTimestampHasPrefix(v string) predicate.HKData {
	return predicate.HKData(sql.FieldHasPrefix(FieldStartTimestamp, v))
}

// StartTimestampHasSuffix applies the HasSuffix predicate on the "start_timestamp" field.
func StartTimestampHasSuffix(v string) predicate.HKData {
	return predicate.HKData(sql.FieldHasSuffix(FieldStartTimestamp, v))
}

// StartTimestampEqualFold applies the EqualFold predicate on the "start_timestamp" field.
func StartTimestampEqualFold(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEqualFold(FieldStartTimestamp, v))
}

// StartTimestampContainsFold applies the ContainsFold predicate on the "start_timestamp" field.
func StartTimestampContainsFold(v string) predicate.HKData {
	return predicate.HKData(sql.FieldContainsFold(FieldStartTimestamp, v))
}

// EndTimestampEQ applies the EQ predicate on the "end_timestamp" field.
func EndTimestampEQ(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldEndTimestamp, v))
}

// EndTimestampNEQ applies the NEQ predicate on the "end_timestamp" field.
func EndTimestampNEQ(v string) predicate.HKData {
	return predicate.HKData(sql.FieldNEQ(FieldEndTimestamp, v))
}

// EndTimestampIn applies the In predicate on the "end_timestamp" field.
func EndTimestampIn(vs ...string) predicate.HKData {
	return predicate.HKData(sql.FieldIn(FieldEndTimestamp, vs...))
}

// EndTimestampNotIn applies the NotIn predicate on the "end_timestamp" field.
func EndTimestampNotIn(vs ...string) predicate.HKData {
	return predicate.HKData(sql.FieldNotIn(FieldEndTimestamp, vs...))
}

// EndTimestampGT applies the GT predicate on the "end_timestamp" field.
func EndTimestampGT(v string) predicate.HKData {
	return predicate.HKData(sql.FieldGT(FieldEndTimestamp, v))
}

// EndTimestampGTE applies the GTE predicate on the "end_timestamp" field.
func EndTimestampGTE(v string) predicate.HKData {
	return predicate.HKData(sql.FieldGTE(FieldEndTimestamp, v))
}

// EndTimestampLT applies the LT predicate on the "end_timestamp" field.
func EndTimestampLT(v string) predicate.HKData {
	return predicate.HKData(sql.FieldLT(FieldEndTimestamp, v))
}

// EndTimestampLTE applies the LTE predicate on the "end_timestamp" field.
func EndTimestampLTE(v string) predicate.HKData {
	return predicate.HKData(sql.FieldLTE(FieldEndTimestamp, v))
}

// EndTimestampContains applies the Contains predicate on the "end_timestamp" field.
func EndTimestampContains(v string) predicate.HKData {
	return predicate.HKData(sql.FieldContains(FieldEndTimestamp, v))
}

// EndTimestampHasPrefix applies the HasPrefix predicate on the "end_timestamp" field.
func EndTimestampHasPrefix(v string) predicate.HKData {
	return predicate.HKData(sql.FieldHasPrefix(FieldEndTimestamp, v))
}

// EndTimestampHasSuffix applies the HasSuffix predicate on the "end_timestamp" field.
func EndTimestampHasSuffix(v string) predicate.HKData {
	return predicate.HKData(sql.FieldHasSuffix(FieldEndTimestamp, v))
}

// EndTimestampEqualFold applies the EqualFold predicate on the "end_timestamp" field.
func EndTimestampEqualFold(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEqualFold(FieldEndTimestamp, v))
}

// EndTimestampContainsFold applies the ContainsFold predicate on the "end_timestamp" field.
func EndTimestampContainsFold(v string) predicate.HKData {
	return predicate.HKData(sql.FieldContainsFold(FieldEndTimestamp, v))
}

// TimezoneIDEQ applies the EQ predicate on the "timezone_id" field.
func TimezoneIDEQ(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEQ(FieldTimezoneID, v))
}

// TimezoneIDNEQ applies the NEQ predicate on the "timezone_id" field.
func TimezoneIDNEQ(v string) predicate.HKData {
	return predicate.HKData(sql.FieldNEQ(FieldTimezoneID, v))
}

// TimezoneIDIn applies the In predicate on the "timezone_id" field.
func TimezoneIDIn(vs ...string) predicate.HKData {
	return predicate.HKData(sql.FieldIn(FieldTimezoneID, vs...))
}

// TimezoneIDNotIn applies the NotIn predicate on the "timezone_id" field.
func TimezoneIDNotIn(vs ...string) predicate.HKData {
	return predicate.HKData(sql.FieldNotIn(FieldTimezoneID, vs...))
}

// TimezoneIDGT applies the GT predicate on the "timezone_id" field.
func TimezoneIDGT(v string) predicate.HKData {
	return predicate.HKData(sql.FieldGT(FieldTimezoneID, v))
}

// TimezoneIDGTE applies the GTE predicate on the "timezone_id" field.
func TimezoneIDGTE(v string) predicate.HKData {
	return predicate.HKData(sql.FieldGTE(FieldTimezoneID, v))
}

// TimezoneIDLT applies the LT predicate on the "timezone_id" field.
func TimezoneIDLT(v string) predicate.HKData {
	return predicate.HKData(sql.FieldLT(FieldTimezoneID, v))
}

// TimezoneIDLTE applies the LTE predicate on the "timezone_id" field.
func TimezoneIDLTE(v string) predicate.HKData {
	return predicate.HKData(sql.FieldLTE(FieldTimezoneID, v))
}

// TimezoneIDContains applies the Contains predicate on the "timezone_id" field.
func TimezoneIDContains(v string) predicate.HKData {
	return predicate.HKData(sql.FieldContains(FieldTimezoneID, v))
}

// TimezoneIDHasPrefix applies the HasPrefix predicate on the "timezone_id" field.
func TimezoneIDHasPrefix(v string) predicate.HKData {
	return predicate.HKData(sql.FieldHasPrefix(FieldTimezoneID, v))
}

// TimezoneIDHasSuffix applies the HasSuffix predicate on the "timezone_id" field.
func TimezoneIDHasSuffix(v string) predicate.HKData {
	return predicate.HKData(sql.FieldHasSuffix(FieldTimezoneID, v))
}

// TimezoneIDEqualFold applies the EqualFold predicate on the "timezone_id" field.
func TimezoneIDEqualFold(v string) predicate.HKData {
	return predicate.HKData(sql.FieldEqualFold(FieldTimezoneID, v))
}

// TimezoneIDContainsFold applies the ContainsFold predicate on the "timezone_id" field.
func TimezoneIDContainsFold(v string) predicate.HKData {
	return predicate.HKData(sql.FieldContainsFold(FieldTimezoneID, v))
}

// HasHealthkit applies the HasEdge predicate on the "healthkit" edge.
func HasHealthkit() predicate.HKData {
	return predicate.HKData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HealthkitTable, HealthkitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHealthkitWith applies the HasEdge predicate on the "healthkit" edge with a given conditions (other predicates).
func HasHealthkitWith(preds ...predicate.HealthKit) predicate.HKData {
	return predicate.HKData(func(s *sql.Selector) {
		step := newHealthkitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.HKData) predicate.HKData {
	return predicate.HKData(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.HKData) predicate.HKData {
	return predicate.HKData(func(s *sql.Selector) {
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
func Not(p predicate.HKData) predicate.HKData {
	return predicate.HKData(func(s *sql.Selector) {
		p(s.Not())
	})
}
