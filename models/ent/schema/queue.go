package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Queue holds the schema definition for the Queue entity.
type Queue struct {
	ent.Schema
}

// Fields of the Queue.
func (Queue) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.UUID("ref", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Queue.
func (Queue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("queues"),
		edge.To("messages", QueueMessage.Type),
	}
}
