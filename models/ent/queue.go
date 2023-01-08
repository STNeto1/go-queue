// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_models/ent/queue"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Queue is the model entity for the Queue schema.
type Queue struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref uuid.UUID `json:"ref,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the QueueQuery when eager-loading is set.
	Edges QueueEdges `json:"edges"`
}

// QueueEdges holds the relations/edges for other nodes in the graph.
type QueueEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e QueueEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e QueueEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[1] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Queue) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case queue.FieldName:
			values[i] = new(sql.NullString)
		case queue.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case queue.FieldID, queue.FieldRef:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Queue", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Queue fields.
func (q *Queue) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case queue.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				q.ID = *value
			}
		case queue.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				q.Name = value.String
			}
		case queue.FieldRef:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value != nil {
				q.Ref = *value
			}
		case queue.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				q.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Queue entity.
func (q *Queue) QueryUser() *UserQuery {
	return (&QueueClient{config: q.config}).QueryUser(q)
}

// QueryMessages queries the "messages" edge of the Queue entity.
func (q *Queue) QueryMessages() *MessageQuery {
	return (&QueueClient{config: q.config}).QueryMessages(q)
}

// Update returns a builder for updating this Queue.
// Note that you need to call Queue.Unwrap() before calling this method if this Queue
// was returned from a transaction, and the transaction was committed or rolled back.
func (q *Queue) Update() *QueueUpdateOne {
	return (&QueueClient{config: q.config}).UpdateOne(q)
}

// Unwrap unwraps the Queue entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (q *Queue) Unwrap() *Queue {
	_tx, ok := q.config.driver.(*txDriver)
	if !ok {
		panic("ent: Queue is not a transactional entity")
	}
	q.config.driver = _tx.drv
	return q
}

// String implements the fmt.Stringer.
func (q *Queue) String() string {
	var builder strings.Builder
	builder.WriteString("Queue(")
	builder.WriteString(fmt.Sprintf("id=%v, ", q.ID))
	builder.WriteString("name=")
	builder.WriteString(q.Name)
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", q.Ref))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(q.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Queues is a parsable slice of Queue.
type Queues []*Queue

func (q Queues) config(cfg config) {
	for _i := range q {
		q[_i].config = cfg
	}
}
