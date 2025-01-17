// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kubecit-service/ent/chapter"
	"kubecit-service/ent/course"
	"kubecit-service/ent/lesson"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChapterCreate is the builder for creating a Chapter entity.
type ChapterCreate struct {
	config
	mutation *ChapterMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *ChapterCreate) SetName(s string) *ChapterCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetReleasedTime sets the "released_time" field.
func (cc *ChapterCreate) SetReleasedTime(t time.Time) *ChapterCreate {
	cc.mutation.SetReleasedTime(t)
	return cc
}

// SetNillableReleasedTime sets the "released_time" field if the given value is not nil.
func (cc *ChapterCreate) SetNillableReleasedTime(t *time.Time) *ChapterCreate {
	if t != nil {
		cc.SetReleasedTime(*t)
	}
	return cc
}

// SetDescription sets the "description" field.
func (cc *ChapterCreate) SetDescription(s string) *ChapterCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetSort sets the "sort" field.
func (cc *ChapterCreate) SetSort(i int) *ChapterCreate {
	cc.mutation.SetSort(i)
	return cc
}

// SetHasFreePreview sets the "has_free_preview" field.
func (cc *ChapterCreate) SetHasFreePreview(i int) *ChapterCreate {
	cc.mutation.SetHasFreePreview(i)
	return cc
}

// SetNillableHasFreePreview sets the "has_free_preview" field if the given value is not nil.
func (cc *ChapterCreate) SetNillableHasFreePreview(i *int) *ChapterCreate {
	if i != nil {
		cc.SetHasFreePreview(*i)
	}
	return cc
}

// SetCourseID sets the "course_id" field.
func (cc *ChapterCreate) SetCourseID(i int) *ChapterCreate {
	cc.mutation.SetCourseID(i)
	return cc
}

// SetNillableCourseID sets the "course_id" field if the given value is not nil.
func (cc *ChapterCreate) SetNillableCourseID(i *int) *ChapterCreate {
	if i != nil {
		cc.SetCourseID(*i)
	}
	return cc
}

// AddLessonIDs adds the "lessons" edge to the Lesson entity by IDs.
func (cc *ChapterCreate) AddLessonIDs(ids ...int) *ChapterCreate {
	cc.mutation.AddLessonIDs(ids...)
	return cc
}

// AddLessons adds the "lessons" edges to the Lesson entity.
func (cc *ChapterCreate) AddLessons(l ...*Lesson) *ChapterCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cc.AddLessonIDs(ids...)
}

// SetCourse sets the "course" edge to the Course entity.
func (cc *ChapterCreate) SetCourse(c *Course) *ChapterCreate {
	return cc.SetCourseID(c.ID)
}

// Mutation returns the ChapterMutation object of the builder.
func (cc *ChapterCreate) Mutation() *ChapterMutation {
	return cc.mutation
}

// Save creates the Chapter in the database.
func (cc *ChapterCreate) Save(ctx context.Context) (*Chapter, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChapterCreate) SaveX(ctx context.Context) *Chapter {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChapterCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChapterCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChapterCreate) defaults() {
	if _, ok := cc.mutation.ReleasedTime(); !ok {
		v := chapter.DefaultReleasedTime()
		cc.mutation.SetReleasedTime(v)
	}
	if _, ok := cc.mutation.HasFreePreview(); !ok {
		v := chapter.DefaultHasFreePreview
		cc.mutation.SetHasFreePreview(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChapterCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Chapter.name"`)}
	}
	if _, ok := cc.mutation.ReleasedTime(); !ok {
		return &ValidationError{Name: "released_time", err: errors.New(`ent: missing required field "Chapter.released_time"`)}
	}
	if _, ok := cc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Chapter.description"`)}
	}
	if _, ok := cc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Chapter.sort"`)}
	}
	if _, ok := cc.mutation.HasFreePreview(); !ok {
		return &ValidationError{Name: "has_free_preview", err: errors.New(`ent: missing required field "Chapter.has_free_preview"`)}
	}
	return nil
}

func (cc *ChapterCreate) sqlSave(ctx context.Context) (*Chapter, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChapterCreate) createSpec() (*Chapter, *sqlgraph.CreateSpec) {
	var (
		_node = &Chapter{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(chapter.Table, sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(chapter.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.ReleasedTime(); ok {
		_spec.SetField(chapter.FieldReleasedTime, field.TypeTime, value)
		_node.ReleasedTime = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(chapter.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.Sort(); ok {
		_spec.SetField(chapter.FieldSort, field.TypeInt, value)
		_node.Sort = value
	}
	if value, ok := cc.mutation.HasFreePreview(); ok {
		_spec.SetField(chapter.FieldHasFreePreview, field.TypeInt, value)
		_node.HasFreePreview = value
	}
	if nodes := cc.mutation.LessonsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chapter.LessonsTable,
			Columns: []string{chapter.LessonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lesson.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chapter.CourseTable,
			Columns: []string{chapter.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CourseID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChapterCreateBulk is the builder for creating many Chapter entities in bulk.
type ChapterCreateBulk struct {
	config
	builders []*ChapterCreate
}

// Save creates the Chapter entities in the database.
func (ccb *ChapterCreateBulk) Save(ctx context.Context) ([]*Chapter, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Chapter, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChapterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChapterCreateBulk) SaveX(ctx context.Context) []*Chapter {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChapterCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChapterCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
