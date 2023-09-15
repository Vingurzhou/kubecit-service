// Code generated by ent, DO NOT EDIT.

package lesson

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the lesson type in the database.
	Label = "lesson"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldReleasedTime holds the string denoting the released_time field in the database.
	FieldReleasedTime = "released_time"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStoragePath holds the string denoting the storage_path field in the database.
	FieldStoragePath = "storage_path"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldCourseware holds the string denoting the courseware field in the database.
	FieldCourseware = "courseware"
	// FieldIsFreePreview holds the string denoting the is_free_preview field in the database.
	FieldIsFreePreview = "is_free_preview"
	// FieldChapterID holds the string denoting the chapter_id field in the database.
	FieldChapterID = "chapter_id"
	// EdgeChapter holds the string denoting the chapter edge name in mutations.
	EdgeChapter = "chapter"
	// Table holds the table name of the lesson in the database.
	Table = "lessons"
	// ChapterTable is the table that holds the chapter relation/edge.
	ChapterTable = "lessons"
	// ChapterInverseTable is the table name for the Chapter entity.
	// It exists in this package in order to avoid circular dependency with the "chapter" package.
	ChapterInverseTable = "chapters"
	// ChapterColumn is the table column denoting the chapter relation/edge.
	ChapterColumn = "chapter_id"
)

// Columns holds all SQL columns for lesson fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldReleasedTime,
	FieldSort,
	FieldType,
	FieldStoragePath,
	FieldSource,
	FieldCourseware,
	FieldIsFreePreview,
	FieldChapterID,
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
	// DefaultReleasedTime holds the default value on creation for the "released_time" field.
	DefaultReleasedTime func() time.Time
	// UpdateDefaultReleasedTime holds the default value on update for the "released_time" field.
	UpdateDefaultReleasedTime func() time.Time
	// DefaultIsFreePreview holds the default value on creation for the "is_free_preview" field.
	DefaultIsFreePreview int
)

// OrderOption defines the ordering options for the Lesson queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByReleasedTime orders the results by the released_time field.
func ByReleasedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReleasedTime, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByStoragePath orders the results by the storage_path field.
func ByStoragePath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStoragePath, opts...).ToFunc()
}

// BySource orders the results by the source field.
func BySource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSource, opts...).ToFunc()
}

// ByCourseware orders the results by the courseware field.
func ByCourseware(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCourseware, opts...).ToFunc()
}

// ByIsFreePreview orders the results by the is_free_preview field.
func ByIsFreePreview(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFreePreview, opts...).ToFunc()
}

// ByChapterID orders the results by the chapter_id field.
func ByChapterID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChapterID, opts...).ToFunc()
}

// ByChapterField orders the results by chapter field.
func ByChapterField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChapterStep(), sql.OrderByField(field, opts...))
	}
}
func newChapterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChapterInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ChapterTable, ChapterColumn),
	)
}