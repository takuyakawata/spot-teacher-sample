// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/emailverification"
)

// EmailVerificationCreate is the builder for creating a EmailVerification entity.
type EmailVerificationCreate struct {
	config
	mutation *EmailVerificationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (evc *EmailVerificationCreate) SetCreatedAt(t time.Time) *EmailVerificationCreate {
	evc.mutation.SetCreatedAt(t)
	return evc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (evc *EmailVerificationCreate) SetNillableCreatedAt(t *time.Time) *EmailVerificationCreate {
	if t != nil {
		evc.SetCreatedAt(*t)
	}
	return evc
}

// SetUpdatedAt sets the "updated_at" field.
func (evc *EmailVerificationCreate) SetUpdatedAt(t time.Time) *EmailVerificationCreate {
	evc.mutation.SetUpdatedAt(t)
	return evc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (evc *EmailVerificationCreate) SetNillableUpdatedAt(t *time.Time) *EmailVerificationCreate {
	if t != nil {
		evc.SetUpdatedAt(*t)
	}
	return evc
}

// SetEmail sets the "email" field.
func (evc *EmailVerificationCreate) SetEmail(s string) *EmailVerificationCreate {
	evc.mutation.SetEmail(s)
	return evc
}

// SetToken sets the "token" field.
func (evc *EmailVerificationCreate) SetToken(s string) *EmailVerificationCreate {
	evc.mutation.SetToken(s)
	return evc
}

// SetExpiredAt sets the "expired_at" field.
func (evc *EmailVerificationCreate) SetExpiredAt(t time.Time) *EmailVerificationCreate {
	evc.mutation.SetExpiredAt(t)
	return evc
}

// SetID sets the "id" field.
func (evc *EmailVerificationCreate) SetID(i int64) *EmailVerificationCreate {
	evc.mutation.SetID(i)
	return evc
}

// Mutation returns the EmailVerificationMutation object of the builder.
func (evc *EmailVerificationCreate) Mutation() *EmailVerificationMutation {
	return evc.mutation
}

// Save creates the EmailVerification in the database.
func (evc *EmailVerificationCreate) Save(ctx context.Context) (*EmailVerification, error) {
	evc.defaults()
	return withHooks(ctx, evc.sqlSave, evc.mutation, evc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (evc *EmailVerificationCreate) SaveX(ctx context.Context) *EmailVerification {
	v, err := evc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (evc *EmailVerificationCreate) Exec(ctx context.Context) error {
	_, err := evc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (evc *EmailVerificationCreate) ExecX(ctx context.Context) {
	if err := evc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (evc *EmailVerificationCreate) defaults() {
	if _, ok := evc.mutation.CreatedAt(); !ok {
		v := emailverification.DefaultCreatedAt()
		evc.mutation.SetCreatedAt(v)
	}
	if _, ok := evc.mutation.UpdatedAt(); !ok {
		v := emailverification.DefaultUpdatedAt()
		evc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (evc *EmailVerificationCreate) check() error {
	if _, ok := evc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EmailVerification.created_at"`)}
	}
	if _, ok := evc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "EmailVerification.updated_at"`)}
	}
	if _, ok := evc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "EmailVerification.email"`)}
	}
	if _, ok := evc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "EmailVerification.token"`)}
	}
	if _, ok := evc.mutation.ExpiredAt(); !ok {
		return &ValidationError{Name: "expired_at", err: errors.New(`ent: missing required field "EmailVerification.expired_at"`)}
	}
	if v, ok := evc.mutation.ID(); ok {
		if err := emailverification.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "EmailVerification.id": %w`, err)}
		}
	}
	return nil
}

func (evc *EmailVerificationCreate) sqlSave(ctx context.Context) (*EmailVerification, error) {
	if err := evc.check(); err != nil {
		return nil, err
	}
	_node, _spec := evc.createSpec()
	if err := sqlgraph.CreateNode(ctx, evc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	evc.mutation.id = &_node.ID
	evc.mutation.done = true
	return _node, nil
}

func (evc *EmailVerificationCreate) createSpec() (*EmailVerification, *sqlgraph.CreateSpec) {
	var (
		_node = &EmailVerification{config: evc.config}
		_spec = sqlgraph.NewCreateSpec(emailverification.Table, sqlgraph.NewFieldSpec(emailverification.FieldID, field.TypeInt64))
	)
	if id, ok := evc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := evc.mutation.CreatedAt(); ok {
		_spec.SetField(emailverification.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := evc.mutation.UpdatedAt(); ok {
		_spec.SetField(emailverification.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := evc.mutation.Email(); ok {
		_spec.SetField(emailverification.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := evc.mutation.Token(); ok {
		_spec.SetField(emailverification.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := evc.mutation.ExpiredAt(); ok {
		_spec.SetField(emailverification.FieldExpiredAt, field.TypeTime, value)
		_node.ExpiredAt = value
	}
	return _node, _spec
}

// EmailVerificationCreateBulk is the builder for creating many EmailVerification entities in bulk.
type EmailVerificationCreateBulk struct {
	config
	err      error
	builders []*EmailVerificationCreate
}

// Save creates the EmailVerification entities in the database.
func (evcb *EmailVerificationCreateBulk) Save(ctx context.Context) ([]*EmailVerification, error) {
	if evcb.err != nil {
		return nil, evcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(evcb.builders))
	nodes := make([]*EmailVerification, len(evcb.builders))
	mutators := make([]Mutator, len(evcb.builders))
	for i := range evcb.builders {
		func(i int, root context.Context) {
			builder := evcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailVerificationMutation)
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
					_, err = mutators[i+1].Mutate(root, evcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, evcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, evcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (evcb *EmailVerificationCreateBulk) SaveX(ctx context.Context) []*EmailVerification {
	v, err := evcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (evcb *EmailVerificationCreateBulk) Exec(ctx context.Context) error {
	_, err := evcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (evcb *EmailVerificationCreateBulk) ExecX(ctx context.Context) {
	if err := evcb.Exec(ctx); err != nil {
		panic(err)
	}
}
