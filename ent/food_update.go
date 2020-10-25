// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"foodly/ent/food"
	"foodly/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// FoodUpdate is the builder for updating Food entities.
type FoodUpdate struct {
	config
	hooks      []Hook
	mutation   *FoodMutation
	predicates []predicate.Food
}

// Where adds a new predicate for the builder.
func (fu *FoodUpdate) Where(ps ...predicate.Food) *FoodUpdate {
	fu.predicates = append(fu.predicates, ps...)
	return fu
}

// SetName sets the name field.
func (fu *FoodUpdate) SetName(s string) *FoodUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetDesc sets the desc field.
func (fu *FoodUpdate) SetDesc(s string) *FoodUpdate {
	fu.mutation.SetDesc(s)
	return fu
}

// SetPrice sets the price field.
func (fu *FoodUpdate) SetPrice(s string) *FoodUpdate {
	fu.mutation.SetPrice(s)
	return fu
}

// SetImageName sets the image_name field.
func (fu *FoodUpdate) SetImageName(s string) *FoodUpdate {
	fu.mutation.SetImageName(s)
	return fu
}

// SetRestaurant sets the restaurant field.
func (fu *FoodUpdate) SetRestaurant(s string) *FoodUpdate {
	fu.mutation.SetRestaurant(s)
	return fu
}

// Mutation returns the FoodMutation object of the builder.
func (fu *FoodUpdate) Mutation() *FoodMutation {
	return fu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FoodUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FoodUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FoodUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FoodUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FoodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   food.Table,
			Columns: food.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: food.FieldID,
			},
		},
	}
	if ps := fu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldName,
		})
	}
	if value, ok := fu.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldDesc,
		})
	}
	if value, ok := fu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldPrice,
		})
	}
	if value, ok := fu.mutation.ImageName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldImageName,
		})
	}
	if value, ok := fu.mutation.Restaurant(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldRestaurant,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{food.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FoodUpdateOne is the builder for updating a single Food entity.
type FoodUpdateOne struct {
	config
	hooks    []Hook
	mutation *FoodMutation
}

// SetName sets the name field.
func (fuo *FoodUpdateOne) SetName(s string) *FoodUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetDesc sets the desc field.
func (fuo *FoodUpdateOne) SetDesc(s string) *FoodUpdateOne {
	fuo.mutation.SetDesc(s)
	return fuo
}

// SetPrice sets the price field.
func (fuo *FoodUpdateOne) SetPrice(s string) *FoodUpdateOne {
	fuo.mutation.SetPrice(s)
	return fuo
}

// SetImageName sets the image_name field.
func (fuo *FoodUpdateOne) SetImageName(s string) *FoodUpdateOne {
	fuo.mutation.SetImageName(s)
	return fuo
}

// SetRestaurant sets the restaurant field.
func (fuo *FoodUpdateOne) SetRestaurant(s string) *FoodUpdateOne {
	fuo.mutation.SetRestaurant(s)
	return fuo
}

// Mutation returns the FoodMutation object of the builder.
func (fuo *FoodUpdateOne) Mutation() *FoodMutation {
	return fuo.mutation
}

// Save executes the query and returns the updated entity.
func (fuo *FoodUpdateOne) Save(ctx context.Context) (*Food, error) {
	var (
		err  error
		node *Food
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FoodUpdateOne) SaveX(ctx context.Context) *Food {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FoodUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FoodUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FoodUpdateOne) sqlSave(ctx context.Context) (_node *Food, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   food.Table,
			Columns: food.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: food.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Food.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := fuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldName,
		})
	}
	if value, ok := fuo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldDesc,
		})
	}
	if value, ok := fuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldPrice,
		})
	}
	if value, ok := fuo.mutation.ImageName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldImageName,
		})
	}
	if value, ok := fuo.mutation.Restaurant(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldRestaurant,
		})
	}
	_node = &Food{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{food.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
