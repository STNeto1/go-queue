package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

const (
	// QueueMessageStatusPending is the status of a message that is waiting to be processed.
	QueueMessageStatusPending = "pending"
	// QueueMessageStatusProcessing is the status of a message that is currently being processed.
	QueueMessageStatusProcessing = "processing"
	// QueueMessageStatusProcessed is the status of a message that has been processed.
	QueueMessageStatusProcessed = "processed"
	// QueueMessageStatusFailed is the status of a message that has failed to be processed.
	QueueMessageStatusFailed = "failed"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("body"),
		field.String("content_type"),
		field.String("status").Default(QueueMessageStatusPending),
		field.Uint("retries").Default(0),
		field.Uint("max_retries").Default(5),
		field.Time("available_from").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("queue", Queue.Type).Ref("messages").Unique(),
	}
}
