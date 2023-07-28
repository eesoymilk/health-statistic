// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package answer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the answer type in the database.
	Label = "answer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// EdgeChosen holds the string denoting the chosen edge name in mutations.
	EdgeChosen = "chosen"
	// EdgeQuestion holds the string denoting the question edge name in mutations.
	EdgeQuestion = "question"
	// EdgeQuestionnaireResponse holds the string denoting the questionnaire_response edge name in mutations.
	EdgeQuestionnaireResponse = "questionnaire_response"
	// Table holds the table name of the answer in the database.
	Table = "answers"
	// ChosenTable is the table that holds the chosen relation/edge.
	ChosenTable = "choices"
	// ChosenInverseTable is the table name for the Choice entity.
	// It exists in this package in order to avoid circular dependency with the "choice" package.
	ChosenInverseTable = "choices"
	// ChosenColumn is the table column denoting the chosen relation/edge.
	ChosenColumn = "answer_chosen"
	// QuestionTable is the table that holds the question relation/edge.
	QuestionTable = "answers"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "questions"
	// QuestionColumn is the table column denoting the question relation/edge.
	QuestionColumn = "question_answers"
	// QuestionnaireResponseTable is the table that holds the questionnaire_response relation/edge.
	QuestionnaireResponseTable = "answers"
	// QuestionnaireResponseInverseTable is the table name for the QuestionnaireResponse entity.
	// It exists in this package in order to avoid circular dependency with the "questionnaireresponse" package.
	QuestionnaireResponseInverseTable = "questionnaire_responses"
	// QuestionnaireResponseColumn is the table column denoting the questionnaire_response relation/edge.
	QuestionnaireResponseColumn = "questionnaire_response_answers"
)

// Columns holds all SQL columns for answer fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldBody,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "answers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"question_answers",
	"questionnaire_response_answers",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Answer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByChosenCount orders the results by chosen count.
func ByChosenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChosenStep(), opts...)
	}
}

// ByChosen orders the results by chosen terms.
func ByChosen(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChosenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByQuestionField orders the results by question field.
func ByQuestionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestionStep(), sql.OrderByField(field, opts...))
	}
}

// ByQuestionnaireResponseField orders the results by questionnaire_response field.
func ByQuestionnaireResponseField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestionnaireResponseStep(), sql.OrderByField(field, opts...))
	}
}
func newChosenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChosenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChosenTable, ChosenColumn),
	)
}
func newQuestionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, QuestionTable, QuestionColumn),
	)
}
func newQuestionnaireResponseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestionnaireResponseInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, QuestionnaireResponseTable, QuestionnaireResponseColumn),
	)
}
