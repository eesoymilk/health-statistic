// Code generated by ent, DO NOT EDIT.

package answer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// EdgeQuestion holds the string denoting the question edge name in mutations.
	EdgeQuestion = "question"
	// EdgeUserQuestionnaire holds the string denoting the user_questionnaire edge name in mutations.
	EdgeUserQuestionnaire = "user_questionnaire"
	// Table holds the table name of the answer in the database.
	Table = "answers"
	// QuestionTable is the table that holds the question relation/edge.
	QuestionTable = "answers"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "questions"
	// QuestionColumn is the table column denoting the question relation/edge.
	QuestionColumn = "question_answers"
	// UserQuestionnaireTable is the table that holds the user_questionnaire relation/edge.
	UserQuestionnaireTable = "answers"
	// UserQuestionnaireInverseTable is the table name for the UserQuestionnaire entity.
	// It exists in this package in order to avoid circular dependency with the "userquestionnaire" package.
	UserQuestionnaireInverseTable = "user_questionnaires"
	// UserQuestionnaireColumn is the table column denoting the user_questionnaire relation/edge.
	UserQuestionnaireColumn = "user_questionnaire_answers"
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
	"user_questionnaire_answers",
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

// ByQuestionField orders the results by question field.
func ByQuestionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestionStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserQuestionnaireField orders the results by user_questionnaire field.
func ByUserQuestionnaireField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserQuestionnaireStep(), sql.OrderByField(field, opts...))
	}
}
func newQuestionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, QuestionTable, QuestionColumn),
	)
}
func newUserQuestionnaireStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserQuestionnaireInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserQuestionnaireTable, UserQuestionnaireColumn),
	)
}