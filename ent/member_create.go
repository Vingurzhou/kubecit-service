// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kubecit-service/ent/member"
	"kubecit-service/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberCreate is the builder for creating a Member entity.
type MemberCreate struct {
	config
	mutation *MemberMutation
	hooks    []Hook
}

// SetOrderNumber sets the "orderNumber" field.
func (mc *MemberCreate) SetOrderNumber(s string) *MemberCreate {
	mc.mutation.SetOrderNumber(s)
	return mc
}

// SetVipName sets the "vipName" field.
func (mc *MemberCreate) SetVipName(s string) *MemberCreate {
	mc.mutation.SetVipName(s)
	return mc
}

// SetVipId sets the "vipId" field.
func (mc *MemberCreate) SetVipId(s string) *MemberCreate {
	mc.mutation.SetVipId(s)
	return mc
}

// SetVipDesc sets the "vipDesc" field.
func (mc *MemberCreate) SetVipDesc(s string) *MemberCreate {
	mc.mutation.SetVipDesc(s)
	return mc
}

// SetStartTime sets the "startTime" field.
func (mc *MemberCreate) SetStartTime(t time.Time) *MemberCreate {
	mc.mutation.SetStartTime(t)
	return mc
}

// SetEndTime sets the "endTime" field.
func (mc *MemberCreate) SetEndTime(t time.Time) *MemberCreate {
	mc.mutation.SetEndTime(t)
	return mc
}

// SetIsExpired sets the "isExpired" field.
func (mc *MemberCreate) SetIsExpired(b bool) *MemberCreate {
	mc.mutation.SetIsExpired(b)
	return mc
}

// SetMemberId sets the "memberId" field.
func (mc *MemberCreate) SetMemberId(s string) *MemberCreate {
	mc.mutation.SetMemberId(s)
	return mc
}

// SetVipIcon sets the "vipIcon" field.
func (mc *MemberCreate) SetVipIcon(s string) *MemberCreate {
	mc.mutation.SetVipIcon(s)
	return mc
}

// SetID sets the "id" field.
func (mc *MemberCreate) SetID(s string) *MemberCreate {
	mc.mutation.SetID(s)
	return mc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mc *MemberCreate) SetUserID(id string) *MemberCreate {
	mc.mutation.SetUserID(id)
	return mc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (mc *MemberCreate) SetNillableUserID(id *string) *MemberCreate {
	if id != nil {
		mc = mc.SetUserID(*id)
	}
	return mc
}

// SetUser sets the "user" edge to the User entity.
func (mc *MemberCreate) SetUser(u *User) *MemberCreate {
	return mc.SetUserID(u.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (mc *MemberCreate) Mutation() *MemberMutation {
	return mc.mutation
}

// Save creates the Member in the database.
func (mc *MemberCreate) Save(ctx context.Context) (*Member, error) {
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MemberCreate) SaveX(ctx context.Context) *Member {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MemberCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MemberCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MemberCreate) check() error {
	if _, ok := mc.mutation.OrderNumber(); !ok {
		return &ValidationError{Name: "orderNumber", err: errors.New(`ent: missing required field "Member.orderNumber"`)}
	}
	if _, ok := mc.mutation.VipName(); !ok {
		return &ValidationError{Name: "vipName", err: errors.New(`ent: missing required field "Member.vipName"`)}
	}
	if _, ok := mc.mutation.VipId(); !ok {
		return &ValidationError{Name: "vipId", err: errors.New(`ent: missing required field "Member.vipId"`)}
	}
	if _, ok := mc.mutation.VipDesc(); !ok {
		return &ValidationError{Name: "vipDesc", err: errors.New(`ent: missing required field "Member.vipDesc"`)}
	}
	if _, ok := mc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "startTime", err: errors.New(`ent: missing required field "Member.startTime"`)}
	}
	if _, ok := mc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "endTime", err: errors.New(`ent: missing required field "Member.endTime"`)}
	}
	if _, ok := mc.mutation.IsExpired(); !ok {
		return &ValidationError{Name: "isExpired", err: errors.New(`ent: missing required field "Member.isExpired"`)}
	}
	if _, ok := mc.mutation.MemberId(); !ok {
		return &ValidationError{Name: "memberId", err: errors.New(`ent: missing required field "Member.memberId"`)}
	}
	if _, ok := mc.mutation.VipIcon(); !ok {
		return &ValidationError{Name: "vipIcon", err: errors.New(`ent: missing required field "Member.vipIcon"`)}
	}
	return nil
}

func (mc *MemberCreate) sqlSave(ctx context.Context) (*Member, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Member.ID type: %T", _spec.ID.Value)
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MemberCreate) createSpec() (*Member, *sqlgraph.CreateSpec) {
	var (
		_node = &Member{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(member.Table, sqlgraph.NewFieldSpec(member.FieldID, field.TypeString))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.OrderNumber(); ok {
		_spec.SetField(member.FieldOrderNumber, field.TypeString, value)
		_node.OrderNumber = value
	}
	if value, ok := mc.mutation.VipName(); ok {
		_spec.SetField(member.FieldVipName, field.TypeString, value)
		_node.VipName = value
	}
	if value, ok := mc.mutation.VipId(); ok {
		_spec.SetField(member.FieldVipId, field.TypeString, value)
		_node.VipId = value
	}
	if value, ok := mc.mutation.VipDesc(); ok {
		_spec.SetField(member.FieldVipDesc, field.TypeString, value)
		_node.VipDesc = value
	}
	if value, ok := mc.mutation.StartTime(); ok {
		_spec.SetField(member.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := mc.mutation.EndTime(); ok {
		_spec.SetField(member.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := mc.mutation.IsExpired(); ok {
		_spec.SetField(member.FieldIsExpired, field.TypeBool, value)
		_node.IsExpired = value
	}
	if value, ok := mc.mutation.MemberId(); ok {
		_spec.SetField(member.FieldMemberId, field.TypeString, value)
		_node.MemberId = value
	}
	if value, ok := mc.mutation.VipIcon(); ok {
		_spec.SetField(member.FieldVipIcon, field.TypeString, value)
		_node.VipIcon = value
	}
	if nodes := mc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   member.UserTable,
			Columns: []string{member.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_vip_member = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemberCreateBulk is the builder for creating many Member entities in bulk.
type MemberCreateBulk struct {
	config
	builders []*MemberCreate
}

// Save creates the Member entities in the database.
func (mcb *MemberCreateBulk) Save(ctx context.Context) ([]*Member, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Member, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MemberCreateBulk) SaveX(ctx context.Context) []*Member {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MemberCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MemberCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
