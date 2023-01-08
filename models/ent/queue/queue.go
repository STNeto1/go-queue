// Code generated by ent, DO NOT EDIT.

package queue

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the queue type in the database.
	Label = "queue"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRef holds the string denoting the ref field in the database.
	FieldRef = "ref"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// Table holds the table name of the queue in the database.
	Table = "queues"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_queues"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// MessagesTable is the table that holds the messages relation/edge. The primary key declared below.
	MessagesTable = "queue_messages"
	// MessagesInverseTable is the table name for the QueueMessage entity.
	// It exists in this package in order to avoid circular dependency with the "queuemessage" package.
	MessagesInverseTable = "queue_messages"
)

// Columns holds all SQL columns for queue fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldRef,
	FieldCreatedAt,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "queue_id"}
	// MessagesPrimaryKey and MessagesColumn2 are the table columns denoting the
	// primary key for the messages relation (M2M).
	MessagesPrimaryKey = []string{"queue_id", "queue_message_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRef holds the default value on creation for the "ref" field.
	DefaultRef func() uuid.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
