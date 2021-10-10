// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (t *Todo) Children(ctx context.Context) ([]*Todo, error) {
	result, err := t.Edges.ChildrenOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryChildren().All(ctx)
	}
	return result, err
}

func (t *Todo) Parent(ctx context.Context) (*Todo, error) {
	result, err := t.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Todo) Owner(ctx context.Context) (*User, error) {
	result, err := t.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Todos(ctx context.Context) ([]*Todo, error) {
	result, err := u.Edges.TodosOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryTodos().All(ctx)
	}
	return result, err
}
