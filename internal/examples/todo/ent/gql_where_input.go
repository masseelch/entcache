// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"time"
	"todo/ent/predicate"
	"todo/ent/todo"
	"todo/ent/user"
)

// TodoWhereInput represents a where input for filtering Todo queries.
type TodoWhereInput struct {
	Not *TodoWhereInput   `json:"not,omitempty"`
	Or  []*TodoWhereInput `json:"or,omitempty"`
	And []*TodoWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "text" field predicates.
	Text             *string  `json:"text,omitempty"`
	TextNEQ          *string  `json:"textNEQ,omitempty"`
	TextIn           []string `json:"textIn,omitempty"`
	TextNotIn        []string `json:"textNotIn,omitempty"`
	TextGT           *string  `json:"textGT,omitempty"`
	TextGTE          *string  `json:"textGTE,omitempty"`
	TextLT           *string  `json:"textLT,omitempty"`
	TextLTE          *string  `json:"textLTE,omitempty"`
	TextContains     *string  `json:"textContains,omitempty"`
	TextHasPrefix    *string  `json:"textHasPrefix,omitempty"`
	TextHasSuffix    *string  `json:"textHasSuffix,omitempty"`
	TextEqualFold    *string  `json:"textEqualFold,omitempty"`
	TextContainsFold *string  `json:"textContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "status" field predicates.
	Status      *todo.Status  `json:"status,omitempty"`
	StatusNEQ   *todo.Status  `json:"statusNEQ,omitempty"`
	StatusIn    []todo.Status `json:"statusIn,omitempty"`
	StatusNotIn []todo.Status `json:"statusNotIn,omitempty"`

	// "priority" field predicates.
	Priority      *int  `json:"priority,omitempty"`
	PriorityNEQ   *int  `json:"priorityNEQ,omitempty"`
	PriorityIn    []int `json:"priorityIn,omitempty"`
	PriorityNotIn []int `json:"priorityNotIn,omitempty"`
	PriorityGT    *int  `json:"priorityGT,omitempty"`
	PriorityGTE   *int  `json:"priorityGTE,omitempty"`
	PriorityLT    *int  `json:"priorityLT,omitempty"`
	PriorityLTE   *int  `json:"priorityLTE,omitempty"`

	// "children" edge predicates.
	HasChildren     *bool             `json:"hasChildren,omitempty"`
	HasChildrenWith []*TodoWhereInput `json:"hasChildrenWith,omitempty"`

	// "parent" edge predicates.
	HasParent     *bool             `json:"hasParent,omitempty"`
	HasParentWith []*TodoWhereInput `json:"hasParentWith,omitempty"`

	// "owner" edge predicates.
	HasOwner     *bool             `json:"hasOwner,omitempty"`
	HasOwnerWith []*UserWhereInput `json:"hasOwnerWith,omitempty"`
}

// Filter applies the TodoWhereInput filter on the TodoQuery builder.
func (i *TodoWhereInput) Filter(q *TodoQuery) (*TodoQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering todos.
// An error is returned if the input is empty or invalid.
func (i *TodoWhereInput) P() (predicate.Todo, error) {
	var predicates []predicate.Todo
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, todo.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Todo, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, todo.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Todo, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, todo.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, todo.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, todo.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, todo.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, todo.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, todo.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, todo.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, todo.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, todo.IDLTE(*i.IDLTE))
	}
	if i.Text != nil {
		predicates = append(predicates, todo.TextEQ(*i.Text))
	}
	if i.TextNEQ != nil {
		predicates = append(predicates, todo.TextNEQ(*i.TextNEQ))
	}
	if len(i.TextIn) > 0 {
		predicates = append(predicates, todo.TextIn(i.TextIn...))
	}
	if len(i.TextNotIn) > 0 {
		predicates = append(predicates, todo.TextNotIn(i.TextNotIn...))
	}
	if i.TextGT != nil {
		predicates = append(predicates, todo.TextGT(*i.TextGT))
	}
	if i.TextGTE != nil {
		predicates = append(predicates, todo.TextGTE(*i.TextGTE))
	}
	if i.TextLT != nil {
		predicates = append(predicates, todo.TextLT(*i.TextLT))
	}
	if i.TextLTE != nil {
		predicates = append(predicates, todo.TextLTE(*i.TextLTE))
	}
	if i.TextContains != nil {
		predicates = append(predicates, todo.TextContains(*i.TextContains))
	}
	if i.TextHasPrefix != nil {
		predicates = append(predicates, todo.TextHasPrefix(*i.TextHasPrefix))
	}
	if i.TextHasSuffix != nil {
		predicates = append(predicates, todo.TextHasSuffix(*i.TextHasSuffix))
	}
	if i.TextEqualFold != nil {
		predicates = append(predicates, todo.TextEqualFold(*i.TextEqualFold))
	}
	if i.TextContainsFold != nil {
		predicates = append(predicates, todo.TextContainsFold(*i.TextContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, todo.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, todo.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, todo.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, todo.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, todo.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, todo.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, todo.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, todo.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.Status != nil {
		predicates = append(predicates, todo.StatusEQ(*i.Status))
	}
	if i.StatusNEQ != nil {
		predicates = append(predicates, todo.StatusNEQ(*i.StatusNEQ))
	}
	if len(i.StatusIn) > 0 {
		predicates = append(predicates, todo.StatusIn(i.StatusIn...))
	}
	if len(i.StatusNotIn) > 0 {
		predicates = append(predicates, todo.StatusNotIn(i.StatusNotIn...))
	}
	if i.Priority != nil {
		predicates = append(predicates, todo.PriorityEQ(*i.Priority))
	}
	if i.PriorityNEQ != nil {
		predicates = append(predicates, todo.PriorityNEQ(*i.PriorityNEQ))
	}
	if len(i.PriorityIn) > 0 {
		predicates = append(predicates, todo.PriorityIn(i.PriorityIn...))
	}
	if len(i.PriorityNotIn) > 0 {
		predicates = append(predicates, todo.PriorityNotIn(i.PriorityNotIn...))
	}
	if i.PriorityGT != nil {
		predicates = append(predicates, todo.PriorityGT(*i.PriorityGT))
	}
	if i.PriorityGTE != nil {
		predicates = append(predicates, todo.PriorityGTE(*i.PriorityGTE))
	}
	if i.PriorityLT != nil {
		predicates = append(predicates, todo.PriorityLT(*i.PriorityLT))
	}
	if i.PriorityLTE != nil {
		predicates = append(predicates, todo.PriorityLTE(*i.PriorityLTE))
	}

	if i.HasChildren != nil {
		p := todo.HasChildren()
		if !*i.HasChildren {
			p = todo.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasChildrenWith) > 0 {
		with := make([]predicate.Todo, 0, len(i.HasChildrenWith))
		for _, w := range i.HasChildrenWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, todo.HasChildrenWith(with...))
	}
	if i.HasParent != nil {
		p := todo.HasParent()
		if !*i.HasParent {
			p = todo.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasParentWith) > 0 {
		with := make([]predicate.Todo, 0, len(i.HasParentWith))
		for _, w := range i.HasParentWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, todo.HasParentWith(with...))
	}
	if i.HasOwner != nil {
		p := todo.HasOwner()
		if !*i.HasOwner {
			p = todo.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOwnerWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasOwnerWith))
		for _, w := range i.HasOwnerWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, todo.HasOwnerWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("todo/ent: empty predicate TodoWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return todo.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Not *UserWhereInput   `json:"not,omitempty"`
	Or  []*UserWhereInput `json:"or,omitempty"`
	And []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "todos" edge predicates.
	HasTodos     *bool             `json:"hasTodos,omitempty"`
	HasTodosWith []*TodoWhereInput `json:"hasTodosWith,omitempty"`
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, user.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, user.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, user.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, user.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, user.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, user.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, user.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, user.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, user.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, user.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, user.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, user.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, user.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasTodos != nil {
		p := user.HasTodos()
		if !*i.HasTodos {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasTodosWith) > 0 {
		with := make([]predicate.Todo, 0, len(i.HasTodosWith))
		for _, w := range i.HasTodosWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasTodosWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("todo/ent: empty predicate UserWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}
