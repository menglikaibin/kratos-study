// Code generated by entc, DO NOT EDIT.

package article

const (
	// Label holds the string label denoting the article type in the database.
	Label = "article"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCityID holds the string denoting the city_id field in the database.
	FieldCityID = "city_id"
	// FieldPublishedAt holds the string denoting the published_at field in the database.
	FieldPublishedAt = "published_at"
	// Table holds the table name of the article in the database.
	Table = "articles"
)

// Columns holds all SQL columns for article fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldContent,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldType,
	FieldCityID,
	FieldPublishedAt,
}

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
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent string
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy int64
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType int64
	// DefaultCityID holds the default value on creation for the "city_id" field.
	DefaultCityID int64
)
