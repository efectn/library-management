// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/efectn/library-management/pkg/database/ent/predicate"
	"github.com/efectn/library-management/pkg/database/ent/role"
	"github.com/efectn/library-management/pkg/database/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// ClearPhone clears the value of the "phone" field.
func (uu *UserUpdate) ClearPhone() *UserUpdate {
	uu.mutation.ClearPhone()
	return uu
}

// SetCity sets the "city" field.
func (uu *UserUpdate) SetCity(s string) *UserUpdate {
	uu.mutation.SetCity(s)
	return uu
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCity(s *string) *UserUpdate {
	if s != nil {
		uu.SetCity(*s)
	}
	return uu
}

// ClearCity clears the value of the "city" field.
func (uu *UserUpdate) ClearCity() *UserUpdate {
	uu.mutation.ClearCity()
	return uu
}

// SetState sets the "state" field.
func (uu *UserUpdate) SetState(s string) *UserUpdate {
	uu.mutation.SetState(s)
	return uu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (uu *UserUpdate) SetNillableState(s *string) *UserUpdate {
	if s != nil {
		uu.SetState(*s)
	}
	return uu
}

// ClearState clears the value of the "state" field.
func (uu *UserUpdate) ClearState() *UserUpdate {
	uu.mutation.ClearState()
	return uu
}

// SetCountry sets the "country" field.
func (uu *UserUpdate) SetCountry(s string) *UserUpdate {
	uu.mutation.SetCountry(s)
	return uu
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCountry(s *string) *UserUpdate {
	if s != nil {
		uu.SetCountry(*s)
	}
	return uu
}

// ClearCountry clears the value of the "country" field.
func (uu *UserUpdate) ClearCountry() *UserUpdate {
	uu.mutation.ClearCountry()
	return uu
}

// SetAddress sets the "address" field.
func (uu *UserUpdate) SetAddress(s string) *UserUpdate {
	uu.mutation.SetAddress(s)
	return uu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAddress(s *string) *UserUpdate {
	if s != nil {
		uu.SetAddress(*s)
	}
	return uu
}

// ClearAddress clears the value of the "address" field.
func (uu *UserUpdate) ClearAddress() *UserUpdate {
	uu.mutation.ClearAddress()
	return uu
}

// SetZipCode sets the "zip_code" field.
func (uu *UserUpdate) SetZipCode(i int) *UserUpdate {
	uu.mutation.ResetZipCode()
	uu.mutation.SetZipCode(i)
	return uu
}

// SetNillableZipCode sets the "zip_code" field if the given value is not nil.
func (uu *UserUpdate) SetNillableZipCode(i *int) *UserUpdate {
	if i != nil {
		uu.SetZipCode(*i)
	}
	return uu
}

// AddZipCode adds i to the "zip_code" field.
func (uu *UserUpdate) AddZipCode(i int) *UserUpdate {
	uu.mutation.AddZipCode(i)
	return uu
}

// ClearZipCode clears the value of the "zip_code" field.
func (uu *UserUpdate) ClearZipCode() *UserUpdate {
	uu.mutation.ClearZipCode()
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uu *UserUpdate) AddRoleIDs(ids ...int) *UserUpdate {
	uu.mutation.AddRoleIDs(ids...)
	return uu
}

// AddRoles adds the "roles" edges to the Role entity.
func (uu *UserUpdate) AddRoles(r ...*Role) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddRoleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (uu *UserUpdate) ClearRoles() *UserUpdate {
	uu.mutation.ClearRoles()
	return uu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (uu *UserUpdate) RemoveRoleIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveRoleIDs(ids...)
	return uu
}

// RemoveRoles removes "roles" edges to Role entities.
func (uu *UserUpdate) RemoveRoles(r ...*Role) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := uu.defaults(); err != nil {
		return 0, err
	}
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() error {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
	}
	if uu.mutation.PhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPhone,
		})
	}
	if value, ok := uu.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCity,
		})
	}
	if uu.mutation.CityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldCity,
		})
	}
	if value, ok := uu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldState,
		})
	}
	if uu.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldState,
		})
	}
	if value, ok := uu.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCountry,
		})
	}
	if uu.mutation.CountryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldCountry,
		})
	}
	if value, ok := uu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAddress,
		})
	}
	if uu.mutation.AddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAddress,
		})
	}
	if value, ok := uu.mutation.ZipCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldZipCode,
		})
	}
	if value, ok := uu.mutation.AddedZipCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldZipCode,
		})
	}
	if uu.mutation.ZipCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldZipCode,
		})
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !uu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// ClearPhone clears the value of the "phone" field.
func (uuo *UserUpdateOne) ClearPhone() *UserUpdateOne {
	uuo.mutation.ClearPhone()
	return uuo
}

// SetCity sets the "city" field.
func (uuo *UserUpdateOne) SetCity(s string) *UserUpdateOne {
	uuo.mutation.SetCity(s)
	return uuo
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCity(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCity(*s)
	}
	return uuo
}

// ClearCity clears the value of the "city" field.
func (uuo *UserUpdateOne) ClearCity() *UserUpdateOne {
	uuo.mutation.ClearCity()
	return uuo
}

// SetState sets the "state" field.
func (uuo *UserUpdateOne) SetState(s string) *UserUpdateOne {
	uuo.mutation.SetState(s)
	return uuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableState(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetState(*s)
	}
	return uuo
}

// ClearState clears the value of the "state" field.
func (uuo *UserUpdateOne) ClearState() *UserUpdateOne {
	uuo.mutation.ClearState()
	return uuo
}

// SetCountry sets the "country" field.
func (uuo *UserUpdateOne) SetCountry(s string) *UserUpdateOne {
	uuo.mutation.SetCountry(s)
	return uuo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCountry(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCountry(*s)
	}
	return uuo
}

// ClearCountry clears the value of the "country" field.
func (uuo *UserUpdateOne) ClearCountry() *UserUpdateOne {
	uuo.mutation.ClearCountry()
	return uuo
}

// SetAddress sets the "address" field.
func (uuo *UserUpdateOne) SetAddress(s string) *UserUpdateOne {
	uuo.mutation.SetAddress(s)
	return uuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAddress(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAddress(*s)
	}
	return uuo
}

// ClearAddress clears the value of the "address" field.
func (uuo *UserUpdateOne) ClearAddress() *UserUpdateOne {
	uuo.mutation.ClearAddress()
	return uuo
}

// SetZipCode sets the "zip_code" field.
func (uuo *UserUpdateOne) SetZipCode(i int) *UserUpdateOne {
	uuo.mutation.ResetZipCode()
	uuo.mutation.SetZipCode(i)
	return uuo
}

// SetNillableZipCode sets the "zip_code" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableZipCode(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetZipCode(*i)
	}
	return uuo
}

// AddZipCode adds i to the "zip_code" field.
func (uuo *UserUpdateOne) AddZipCode(i int) *UserUpdateOne {
	uuo.mutation.AddZipCode(i)
	return uuo
}

// ClearZipCode clears the value of the "zip_code" field.
func (uuo *UserUpdateOne) ClearZipCode() *UserUpdateOne {
	uuo.mutation.ClearZipCode()
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uuo *UserUpdateOne) AddRoleIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddRoleIDs(ids...)
	return uuo
}

// AddRoles adds the "roles" edges to the Role entity.
func (uuo *UserUpdateOne) AddRoles(r ...*Role) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddRoleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (uuo *UserUpdateOne) ClearRoles() *UserUpdateOne {
	uuo.mutation.ClearRoles()
	return uuo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (uuo *UserUpdateOne) RemoveRoleIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveRoleIDs(ids...)
	return uuo
}

// RemoveRoles removes "roles" edges to Role entities.
func (uuo *UserUpdateOne) RemoveRoles(r ...*Role) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveRoleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if err := uuo.defaults(); err != nil {
		return nil, err
	}
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() error {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
	}
	if uuo.mutation.PhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPhone,
		})
	}
	if value, ok := uuo.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCity,
		})
	}
	if uuo.mutation.CityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldCity,
		})
	}
	if value, ok := uuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldState,
		})
	}
	if uuo.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldState,
		})
	}
	if value, ok := uuo.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldCountry,
		})
	}
	if uuo.mutation.CountryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldCountry,
		})
	}
	if value, ok := uuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAddress,
		})
	}
	if uuo.mutation.AddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAddress,
		})
	}
	if value, ok := uuo.mutation.ZipCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldZipCode,
		})
	}
	if value, ok := uuo.mutation.AddedZipCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldZipCode,
		})
	}
	if uuo.mutation.ZipCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldZipCode,
		})
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !uuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
