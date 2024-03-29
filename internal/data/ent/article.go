// Code generated by entc, DO NOT EDIT.

package ent

import (
	"blog/internal/data/ent/article"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Article is the model entity for the Article schema.
type Article struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	// 测试一哈
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	// 测试一哈
	Content string `json:"content,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	// 测试一哈
	CreatedBy int64 `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	// 测试一哈
	UpdatedBy int64 `json:"updated_by,omitempty"`
	// Type holds the value of the "type" field.
	// 测试一哈
	Type int64 `json:"type,omitempty"`
	// CityID holds the value of the "city_id" field.
	// 测试一哈
	CityID int64 `json:"city_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	// 测试一哈
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	// 测试一哈
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	// 测试一哈
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldID, article.FieldCreatedBy, article.FieldUpdatedBy, article.FieldType, article.FieldCityID:
			values[i] = new(sql.NullInt64)
		case article.FieldTitle, article.FieldContent:
			values[i] = new(sql.NullString)
		case article.FieldCreatedAt, article.FieldUpdatedAt, article.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Article", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Article fields.
func (a *Article) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case article.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case article.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				a.Content = value.String
			}
		case article.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				a.CreatedBy = value.Int64
			}
		case article.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				a.UpdatedBy = value.Int64
			}
		case article.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = value.Int64
			}
		case article.FieldCityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field city_id", values[i])
			} else if value.Valid {
				a.CityID = value.Int64
			}
		case article.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case article.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case article.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = new(time.Time)
				*a.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Article.
// Note that you need to call Article.Unwrap() before calling this method if this Article
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Article) Update() *ArticleUpdateOne {
	return (&ArticleClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Article entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Article) Unwrap() *Article {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Article is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Article) String() string {
	var builder strings.Builder
	builder.WriteString("Article(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", title=")
	builder.WriteString(a.Title)
	builder.WriteString(", content=")
	builder.WriteString(a.Content)
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedBy))
	builder.WriteString(", updated_by=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedBy))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", a.Type))
	builder.WriteString(", city_id=")
	builder.WriteString(fmt.Sprintf("%v", a.CityID))
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	if v := a.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Articles is a parsable slice of Article.
type Articles []*Article

func (a Articles) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
