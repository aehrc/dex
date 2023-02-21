// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dexidp/dex/storage/ent/db/devicerequest"
	"github.com/dexidp/dex/storage/ent/db/predicate"
)

// DeviceRequestDelete is the builder for deleting a DeviceRequest entity.
type DeviceRequestDelete struct {
	config
	hooks    []Hook
	mutation *DeviceRequestMutation
}

// Where appends a list predicates to the DeviceRequestDelete builder.
func (drd *DeviceRequestDelete) Where(ps ...predicate.DeviceRequest) *DeviceRequestDelete {
	drd.mutation.Where(ps...)
	return drd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (drd *DeviceRequestDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(drd.hooks) == 0 {
		affected, err = drd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			drd.mutation = mutation
			affected, err = drd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(drd.hooks) - 1; i >= 0; i-- {
			if drd.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = drd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, drd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (drd *DeviceRequestDelete) ExecX(ctx context.Context) int {
	n, err := drd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (drd *DeviceRequestDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: devicerequest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: devicerequest.FieldID,
			},
		},
	}
	if ps := drd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, drd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DeviceRequestDeleteOne is the builder for deleting a single DeviceRequest entity.
type DeviceRequestDeleteOne struct {
	drd *DeviceRequestDelete
}

// Exec executes the deletion query.
func (drdo *DeviceRequestDeleteOne) Exec(ctx context.Context) error {
	n, err := drdo.drd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{devicerequest.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (drdo *DeviceRequestDeleteOne) ExecX(ctx context.Context) {
	drdo.drd.ExecX(ctx)
}
