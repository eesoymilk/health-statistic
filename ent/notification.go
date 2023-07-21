// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/eesoymilk/health-statistic-api/ent/notification"
)

// Notification is the model entity for the Notification schema.
type Notification struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SentAt holds the value of the "sent_at" field.
	SentAt time.Time `json:"sent_at,omitempty"`
	// ReadAt holds the value of the "read_at" field.
	ReadAt time.Time `json:"read_at,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NotificationQuery when eager-loading is set.
	Edges        NotificationEdges `json:"-"`
	selectValues sql.SelectValues
}

// NotificationEdges holds the relations/edges for other nodes in the graph.
type NotificationEdges struct {
	// Recipient holds the value of the recipient edge.
	Recipient []*User `json:"recipient,omitempty"`
	// Mycard holds the value of the mycard edge.
	Mycard []*MyCard `json:"mycard,omitempty"`
	// Price holds the value of the price edge.
	Price []*Price `json:"price,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// RecipientOrErr returns the Recipient value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationEdges) RecipientOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Recipient, nil
	}
	return nil, &NotLoadedError{edge: "recipient"}
}

// MycardOrErr returns the Mycard value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationEdges) MycardOrErr() ([]*MyCard, error) {
	if e.loadedTypes[1] {
		return e.Mycard, nil
	}
	return nil, &NotLoadedError{edge: "mycard"}
}

// PriceOrErr returns the Price value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationEdges) PriceOrErr() ([]*Price, error) {
	if e.loadedTypes[2] {
		return e.Price, nil
	}
	return nil, &NotLoadedError{edge: "price"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Notification) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notification.FieldID:
			values[i] = new(sql.NullInt64)
		case notification.FieldMessage:
			values[i] = new(sql.NullString)
		case notification.FieldSentAt, notification.FieldReadAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Notification fields.
func (n *Notification) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notification.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case notification.FieldSentAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sent_at", values[i])
			} else if value.Valid {
				n.SentAt = value.Time
			}
		case notification.FieldReadAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field read_at", values[i])
			} else if value.Valid {
				n.ReadAt = value.Time
			}
		case notification.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				n.Message = value.String
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Notification.
// This includes values selected through modifiers, order, etc.
func (n *Notification) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryRecipient queries the "recipient" edge of the Notification entity.
func (n *Notification) QueryRecipient() *UserQuery {
	return NewNotificationClient(n.config).QueryRecipient(n)
}

// QueryMycard queries the "mycard" edge of the Notification entity.
func (n *Notification) QueryMycard() *MyCardQuery {
	return NewNotificationClient(n.config).QueryMycard(n)
}

// QueryPrice queries the "price" edge of the Notification entity.
func (n *Notification) QueryPrice() *PriceQuery {
	return NewNotificationClient(n.config).QueryPrice(n)
}

// Update returns a builder for updating this Notification.
// Note that you need to call Notification.Unwrap() before calling this method if this Notification
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Notification) Update() *NotificationUpdateOne {
	return NewNotificationClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Notification entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Notification) Unwrap() *Notification {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Notification is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Notification) String() string {
	var builder strings.Builder
	builder.WriteString("Notification(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("sent_at=")
	builder.WriteString(n.SentAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("read_at=")
	builder.WriteString(n.ReadAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(n.Message)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (n *Notification) MarshalJSON() ([]byte, error) {
	type Alias Notification
	return json.Marshal(&struct {
		*Alias
		NotificationEdges
	}{
		Alias:             (*Alias)(n),
		NotificationEdges: n.Edges,
	})
}

// Notifications is a parsable slice of Notification.
type Notifications []*Notification
