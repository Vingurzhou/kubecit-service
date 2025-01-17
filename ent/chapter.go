// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kubecit-service/ent/chapter"
	"kubecit-service/ent/course"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Chapter is the model entity for the Chapter schema.
type Chapter struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 章节名称
	Name string `json:"name,omitempty"`
	// 发布时间
	ReleasedTime time.Time `json:"released_time,omitempty"`
	// 章节描述
	Description string `json:"description,omitempty"`
	// 序号
	Sort int `json:"sort,omitempty"`
	// 是否有免费试看
	HasFreePreview int `json:"has_free_preview,omitempty"`
	// 课程id
	CourseID int `json:"course_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChapterQuery when eager-loading is set.
	Edges        ChapterEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ChapterEdges holds the relations/edges for other nodes in the graph.
type ChapterEdges struct {
	// Lessons holds the value of the lessons edge.
	Lessons []*Lesson `json:"lessons,omitempty"`
	// Course holds the value of the course edge.
	Course *Course `json:"course,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// LessonsOrErr returns the Lessons value or an error if the edge
// was not loaded in eager-loading.
func (e ChapterEdges) LessonsOrErr() ([]*Lesson, error) {
	if e.loadedTypes[0] {
		return e.Lessons, nil
	}
	return nil, &NotLoadedError{edge: "lessons"}
}

// CourseOrErr returns the Course value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChapterEdges) CourseOrErr() (*Course, error) {
	if e.loadedTypes[1] {
		if e.Course == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: course.Label}
		}
		return e.Course, nil
	}
	return nil, &NotLoadedError{edge: "course"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Chapter) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chapter.FieldID, chapter.FieldSort, chapter.FieldHasFreePreview, chapter.FieldCourseID:
			values[i] = new(sql.NullInt64)
		case chapter.FieldName, chapter.FieldDescription:
			values[i] = new(sql.NullString)
		case chapter.FieldReleasedTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Chapter fields.
func (c *Chapter) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chapter.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case chapter.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case chapter.FieldReleasedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field released_time", values[i])
			} else if value.Valid {
				c.ReleasedTime = value.Time
			}
		case chapter.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case chapter.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				c.Sort = int(value.Int64)
			}
		case chapter.FieldHasFreePreview:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field has_free_preview", values[i])
			} else if value.Valid {
				c.HasFreePreview = int(value.Int64)
			}
		case chapter.FieldCourseID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field course_id", values[i])
			} else if value.Valid {
				c.CourseID = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Chapter.
// This includes values selected through modifiers, order, etc.
func (c *Chapter) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryLessons queries the "lessons" edge of the Chapter entity.
func (c *Chapter) QueryLessons() *LessonQuery {
	return NewChapterClient(c.config).QueryLessons(c)
}

// QueryCourse queries the "course" edge of the Chapter entity.
func (c *Chapter) QueryCourse() *CourseQuery {
	return NewChapterClient(c.config).QueryCourse(c)
}

// Update returns a builder for updating this Chapter.
// Note that you need to call Chapter.Unwrap() before calling this method if this Chapter
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Chapter) Update() *ChapterUpdateOne {
	return NewChapterClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Chapter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Chapter) Unwrap() *Chapter {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Chapter is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Chapter) String() string {
	var builder strings.Builder
	builder.WriteString("Chapter(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("released_time=")
	builder.WriteString(c.ReleasedTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", c.Sort))
	builder.WriteString(", ")
	builder.WriteString("has_free_preview=")
	builder.WriteString(fmt.Sprintf("%v", c.HasFreePreview))
	builder.WriteString(", ")
	builder.WriteString("course_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CourseID))
	builder.WriteByte(')')
	return builder.String()
}

// Chapters is a parsable slice of Chapter.
type Chapters []*Chapter
