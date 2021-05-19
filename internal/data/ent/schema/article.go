package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Default("").Comment("测试一哈"),
		field.String("content").Default("").Comment("测试一哈"),
		field.Int64("created_by").Default(0).Comment("测试一哈"),
		field.Int64("updated_by").Default(0).Comment("测试一哈"),
		field.Int64("type").Default(0).Comment("测试一哈"),
		field.Int64("city_id").Default(0).Comment("测试一哈"),
		field.Time("created_at").Comment("测试一哈"),
		field.Time("updated_at").Comment("测试一哈"),
		field.Time("deleted_at").Nillable().Optional().Comment("测试一哈"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
