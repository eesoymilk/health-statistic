// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package price

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the price type in the database.
	Label = "price"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldTakenAt holds the string denoting the taken_at field in the database.
	FieldTakenAt = "taken_at"
	// EdgeRecipient holds the string denoting the recipient edge name in mutations.
	EdgeRecipient = "recipient"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// Table holds the table name of the price in the database.
	Table = "prices"
	// RecipientTable is the table that holds the recipient relation/edge.
	RecipientTable = "prices"
	// RecipientInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RecipientInverseTable = "users"
	// RecipientColumn is the table column denoting the recipient relation/edge.
	RecipientColumn = "price_recipient"
	// NotificationsTable is the table that holds the notifications relation/edge. The primary key declared below.
	NotificationsTable = "price_notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
)

// Columns holds all SQL columns for price fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldCreatedAt,
	FieldTakenAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "prices"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"price_recipient",
}

var (
	// NotificationsPrimaryKey and NotificationsColumn2 are the table columns denoting the
	// primary key for the notifications relation (M2M).
	NotificationsPrimaryKey = []string{"price_id", "notification_id"}
)

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

// OrderOption defines the ordering options for the Price queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByTakenAt orders the results by the taken_at field.
func ByTakenAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTakenAt, opts...).ToFunc()
}

// ByRecipientField orders the results by recipient field.
func ByRecipientField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRecipientStep(), sql.OrderByField(field, opts...))
	}
}

// ByNotificationsCount orders the results by notifications count.
func ByNotificationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationsStep(), opts...)
	}
}

// ByNotifications orders the results by notifications terms.
func ByNotifications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRecipientStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RecipientInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, RecipientTable, RecipientColumn),
	)
}
func newNotificationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, NotificationsTable, NotificationsPrimaryKey...),
	)
}
