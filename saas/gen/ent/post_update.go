// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/post"
	"saas/gen/ent/postcategory"
	"saas/gen/ent/poststatus"
	"saas/gen/ent/posttag"
	"saas/gen/ent/posttype"
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks     []Hook
	mutation  *PostMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PostUpdate) SetUpdatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pu *PostUpdate) ClearUpdatedAt() *PostUpdate {
	pu.mutation.ClearUpdatedAt()
	return pu
}

// SetAppID sets the "app_id" field.
func (pu *PostUpdate) SetAppID(s string) *PostUpdate {
	pu.mutation.SetAppID(s)
	return pu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillableAppID(s *string) *PostUpdate {
	if s != nil {
		pu.SetAppID(*s)
	}
	return pu
}

// ClearAppID clears the value of the "app_id" field.
func (pu *PostUpdate) ClearAppID() *PostUpdate {
	pu.mutation.ClearAppID()
	return pu
}

// SetName sets the "name" field.
func (pu *PostUpdate) SetName(s string) *PostUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PostUpdate) SetNillableName(s *string) *PostUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// ClearName clears the value of the "name" field.
func (pu *PostUpdate) ClearName() *PostUpdate {
	pu.mutation.ClearName()
	return pu
}

// SetSlug sets the "slug" field.
func (pu *PostUpdate) SetSlug(s string) *PostUpdate {
	pu.mutation.SetSlug(s)
	return pu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (pu *PostUpdate) SetNillableSlug(s *string) *PostUpdate {
	if s != nil {
		pu.SetSlug(*s)
	}
	return pu
}

// ClearSlug clears the value of the "slug" field.
func (pu *PostUpdate) ClearSlug() *PostUpdate {
	pu.mutation.ClearSlug()
	return pu
}

// SetPostStatusID sets the "post_status_id" field.
func (pu *PostUpdate) SetPostStatusID(s string) *PostUpdate {
	pu.mutation.SetPostStatusID(s)
	return pu
}

// SetNillablePostStatusID sets the "post_status_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillablePostStatusID(s *string) *PostUpdate {
	if s != nil {
		pu.SetPostStatusID(*s)
	}
	return pu
}

// ClearPostStatusID clears the value of the "post_status_id" field.
func (pu *PostUpdate) ClearPostStatusID() *PostUpdate {
	pu.mutation.ClearPostStatusID()
	return pu
}

// SetPostTypeID sets the "post_type_id" field.
func (pu *PostUpdate) SetPostTypeID(s string) *PostUpdate {
	pu.mutation.SetPostTypeID(s)
	return pu
}

// SetNillablePostTypeID sets the "post_type_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillablePostTypeID(s *string) *PostUpdate {
	if s != nil {
		pu.SetPostTypeID(*s)
	}
	return pu
}

// ClearPostTypeID clears the value of the "post_type_id" field.
func (pu *PostUpdate) ClearPostTypeID() *PostUpdate {
	pu.mutation.ClearPostTypeID()
	return pu
}

// SetPrimaryCategoryID sets the "primary_category_id" field.
func (pu *PostUpdate) SetPrimaryCategoryID(s string) *PostUpdate {
	pu.mutation.SetPrimaryCategoryID(s)
	return pu
}

// SetNillablePrimaryCategoryID sets the "primary_category_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillablePrimaryCategoryID(s *string) *PostUpdate {
	if s != nil {
		pu.SetPrimaryCategoryID(*s)
	}
	return pu
}

// ClearPrimaryCategoryID clears the value of the "primary_category_id" field.
func (pu *PostUpdate) ClearPrimaryCategoryID() *PostUpdate {
	pu.mutation.ClearPrimaryCategoryID()
	return pu
}

// SetHeadline sets the "headline" field.
func (pu *PostUpdate) SetHeadline(s string) *PostUpdate {
	pu.mutation.SetHeadline(s)
	return pu
}

// SetNillableHeadline sets the "headline" field if the given value is not nil.
func (pu *PostUpdate) SetNillableHeadline(s *string) *PostUpdate {
	if s != nil {
		pu.SetHeadline(*s)
	}
	return pu
}

// ClearHeadline clears the value of the "headline" field.
func (pu *PostUpdate) ClearHeadline() *PostUpdate {
	pu.mutation.ClearHeadline()
	return pu
}

// SetExcerpt sets the "excerpt" field.
func (pu *PostUpdate) SetExcerpt(s string) *PostUpdate {
	pu.mutation.SetExcerpt(s)
	return pu
}

// SetNillableExcerpt sets the "excerpt" field if the given value is not nil.
func (pu *PostUpdate) SetNillableExcerpt(s *string) *PostUpdate {
	if s != nil {
		pu.SetExcerpt(*s)
	}
	return pu
}

// ClearExcerpt clears the value of the "excerpt" field.
func (pu *PostUpdate) ClearExcerpt() *PostUpdate {
	pu.mutation.ClearExcerpt()
	return pu
}

// SetContent sets the "content" field.
func (pu *PostUpdate) SetContent(s string) *PostUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (pu *PostUpdate) SetNillableContent(s *string) *PostUpdate {
	if s != nil {
		pu.SetContent(*s)
	}
	return pu
}

// ClearContent clears the value of the "content" field.
func (pu *PostUpdate) ClearContent() *PostUpdate {
	pu.mutation.ClearContent()
	return pu
}

// SetMetaTitle sets the "meta_title" field.
func (pu *PostUpdate) SetMetaTitle(s string) *PostUpdate {
	pu.mutation.SetMetaTitle(s)
	return pu
}

// SetNillableMetaTitle sets the "meta_title" field if the given value is not nil.
func (pu *PostUpdate) SetNillableMetaTitle(s *string) *PostUpdate {
	if s != nil {
		pu.SetMetaTitle(*s)
	}
	return pu
}

// ClearMetaTitle clears the value of the "meta_title" field.
func (pu *PostUpdate) ClearMetaTitle() *PostUpdate {
	pu.mutation.ClearMetaTitle()
	return pu
}

// SetMetaDescr sets the "meta_descr" field.
func (pu *PostUpdate) SetMetaDescr(s string) *PostUpdate {
	pu.mutation.SetMetaDescr(s)
	return pu
}

// SetNillableMetaDescr sets the "meta_descr" field if the given value is not nil.
func (pu *PostUpdate) SetNillableMetaDescr(s *string) *PostUpdate {
	if s != nil {
		pu.SetMetaDescr(*s)
	}
	return pu
}

// ClearMetaDescr clears the value of the "meta_descr" field.
func (pu *PostUpdate) ClearMetaDescr() *PostUpdate {
	pu.mutation.ClearMetaDescr()
	return pu
}

// SetMetaCanonicalURL sets the "meta_canonical_url" field.
func (pu *PostUpdate) SetMetaCanonicalURL(s string) *PostUpdate {
	pu.mutation.SetMetaCanonicalURL(s)
	return pu
}

// SetNillableMetaCanonicalURL sets the "meta_canonical_url" field if the given value is not nil.
func (pu *PostUpdate) SetNillableMetaCanonicalURL(s *string) *PostUpdate {
	if s != nil {
		pu.SetMetaCanonicalURL(*s)
	}
	return pu
}

// ClearMetaCanonicalURL clears the value of the "meta_canonical_url" field.
func (pu *PostUpdate) ClearMetaCanonicalURL() *PostUpdate {
	pu.mutation.ClearMetaCanonicalURL()
	return pu
}

// SetMetaRobots sets the "meta_robots" field.
func (pu *PostUpdate) SetMetaRobots(s string) *PostUpdate {
	pu.mutation.SetMetaRobots(s)
	return pu
}

// SetNillableMetaRobots sets the "meta_robots" field if the given value is not nil.
func (pu *PostUpdate) SetNillableMetaRobots(s *string) *PostUpdate {
	if s != nil {
		pu.SetMetaRobots(*s)
	}
	return pu
}

// ClearMetaRobots clears the value of the "meta_robots" field.
func (pu *PostUpdate) ClearMetaRobots() *PostUpdate {
	pu.mutation.ClearMetaRobots()
	return pu
}

// SetCustom sets the "custom" field.
func (pu *PostUpdate) SetCustom(m map[string]interface{}) *PostUpdate {
	pu.mutation.SetCustom(m)
	return pu
}

// ClearCustom clears the value of the "custom" field.
func (pu *PostUpdate) ClearCustom() *PostUpdate {
	pu.mutation.ClearCustom()
	return pu
}

// SetPostStatus sets the "post_status" edge to the PostStatus entity.
func (pu *PostUpdate) SetPostStatus(p *PostStatus) *PostUpdate {
	return pu.SetPostStatusID(p.ID)
}

// SetPostType sets the "post_type" edge to the PostType entity.
func (pu *PostUpdate) SetPostType(p *PostType) *PostUpdate {
	return pu.SetPostTypeID(p.ID)
}

// SetPrimaryCategory sets the "primary_category" edge to the PostCategory entity.
func (pu *PostUpdate) SetPrimaryCategory(p *PostCategory) *PostUpdate {
	return pu.SetPrimaryCategoryID(p.ID)
}

// AddPostTagIDs adds the "post_tags" edge to the PostTag entity by IDs.
func (pu *PostUpdate) AddPostTagIDs(ids ...string) *PostUpdate {
	pu.mutation.AddPostTagIDs(ids...)
	return pu
}

// AddPostTags adds the "post_tags" edges to the PostTag entity.
func (pu *PostUpdate) AddPostTags(p ...*PostTag) *PostUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPostTagIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearPostStatus clears the "post_status" edge to the PostStatus entity.
func (pu *PostUpdate) ClearPostStatus() *PostUpdate {
	pu.mutation.ClearPostStatus()
	return pu
}

// ClearPostType clears the "post_type" edge to the PostType entity.
func (pu *PostUpdate) ClearPostType() *PostUpdate {
	pu.mutation.ClearPostType()
	return pu
}

// ClearPrimaryCategory clears the "primary_category" edge to the PostCategory entity.
func (pu *PostUpdate) ClearPrimaryCategory() *PostUpdate {
	pu.mutation.ClearPrimaryCategory()
	return pu
}

// ClearPostTags clears all "post_tags" edges to the PostTag entity.
func (pu *PostUpdate) ClearPostTags() *PostUpdate {
	pu.mutation.ClearPostTags()
	return pu
}

// RemovePostTagIDs removes the "post_tags" edge to PostTag entities by IDs.
func (pu *PostUpdate) RemovePostTagIDs(ids ...string) *PostUpdate {
	pu.mutation.RemovePostTagIDs(ids...)
	return pu
}

// RemovePostTags removes "post_tags" edges to PostTag entities.
func (pu *PostUpdate) RemovePostTags(p ...*PostTag) *PostUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePostTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PostUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok && !pu.mutation.UpdatedAtCleared() {
		v := post.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PostUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PostUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pu.mutation.CreatedAtCleared() {
		_spec.ClearField(post.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.UpdatedAtCleared() {
		_spec.ClearField(post.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.AppID(); ok {
		_spec.SetField(post.FieldAppID, field.TypeString, value)
	}
	if pu.mutation.AppIDCleared() {
		_spec.ClearField(post.FieldAppID, field.TypeString)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(post.FieldName, field.TypeString, value)
	}
	if pu.mutation.NameCleared() {
		_spec.ClearField(post.FieldName, field.TypeString)
	}
	if value, ok := pu.mutation.Slug(); ok {
		_spec.SetField(post.FieldSlug, field.TypeString, value)
	}
	if pu.mutation.SlugCleared() {
		_spec.ClearField(post.FieldSlug, field.TypeString)
	}
	if value, ok := pu.mutation.Headline(); ok {
		_spec.SetField(post.FieldHeadline, field.TypeString, value)
	}
	if pu.mutation.HeadlineCleared() {
		_spec.ClearField(post.FieldHeadline, field.TypeString)
	}
	if value, ok := pu.mutation.Excerpt(); ok {
		_spec.SetField(post.FieldExcerpt, field.TypeString, value)
	}
	if pu.mutation.ExcerptCleared() {
		_spec.ClearField(post.FieldExcerpt, field.TypeString)
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if pu.mutation.ContentCleared() {
		_spec.ClearField(post.FieldContent, field.TypeString)
	}
	if value, ok := pu.mutation.MetaTitle(); ok {
		_spec.SetField(post.FieldMetaTitle, field.TypeString, value)
	}
	if pu.mutation.MetaTitleCleared() {
		_spec.ClearField(post.FieldMetaTitle, field.TypeString)
	}
	if value, ok := pu.mutation.MetaDescr(); ok {
		_spec.SetField(post.FieldMetaDescr, field.TypeString, value)
	}
	if pu.mutation.MetaDescrCleared() {
		_spec.ClearField(post.FieldMetaDescr, field.TypeString)
	}
	if value, ok := pu.mutation.MetaCanonicalURL(); ok {
		_spec.SetField(post.FieldMetaCanonicalURL, field.TypeString, value)
	}
	if pu.mutation.MetaCanonicalURLCleared() {
		_spec.ClearField(post.FieldMetaCanonicalURL, field.TypeString)
	}
	if value, ok := pu.mutation.MetaRobots(); ok {
		_spec.SetField(post.FieldMetaRobots, field.TypeString, value)
	}
	if pu.mutation.MetaRobotsCleared() {
		_spec.ClearField(post.FieldMetaRobots, field.TypeString)
	}
	if value, ok := pu.mutation.Custom(); ok {
		_spec.SetField(post.FieldCustom, field.TypeJSON, value)
	}
	if pu.mutation.CustomCleared() {
		_spec.ClearField(post.FieldCustom, field.TypeJSON)
	}
	if pu.mutation.PostStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PostStatusTable,
			Columns: []string{post.PostStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poststatus.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PostStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PostStatusTable,
			Columns: []string{post.PostStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poststatus.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PostTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PostTypeTable,
			Columns: []string{post.PostTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttype.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PostTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PostTypeTable,
			Columns: []string{post.PostTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttype.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PrimaryCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PrimaryCategoryTable,
			Columns: []string{post.PrimaryCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PrimaryCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PrimaryCategoryTable,
			Columns: []string{post.PrimaryCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PostTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.PostTagsTable,
			Columns: post.PostTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttag.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedPostTagsIDs(); len(nodes) > 0 && !pu.mutation.PostTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.PostTagsTable,
			Columns: post.PostTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PostTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.PostTagsTable,
			Columns: post.PostTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PostMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PostUpdateOne) SetUpdatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (puo *PostUpdateOne) ClearUpdatedAt() *PostUpdateOne {
	puo.mutation.ClearUpdatedAt()
	return puo
}

// SetAppID sets the "app_id" field.
func (puo *PostUpdateOne) SetAppID(s string) *PostUpdateOne {
	puo.mutation.SetAppID(s)
	return puo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableAppID(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetAppID(*s)
	}
	return puo
}

// ClearAppID clears the value of the "app_id" field.
func (puo *PostUpdateOne) ClearAppID() *PostUpdateOne {
	puo.mutation.ClearAppID()
	return puo
}

// SetName sets the "name" field.
func (puo *PostUpdateOne) SetName(s string) *PostUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableName(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// ClearName clears the value of the "name" field.
func (puo *PostUpdateOne) ClearName() *PostUpdateOne {
	puo.mutation.ClearName()
	return puo
}

// SetSlug sets the "slug" field.
func (puo *PostUpdateOne) SetSlug(s string) *PostUpdateOne {
	puo.mutation.SetSlug(s)
	return puo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableSlug(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetSlug(*s)
	}
	return puo
}

// ClearSlug clears the value of the "slug" field.
func (puo *PostUpdateOne) ClearSlug() *PostUpdateOne {
	puo.mutation.ClearSlug()
	return puo
}

// SetPostStatusID sets the "post_status_id" field.
func (puo *PostUpdateOne) SetPostStatusID(s string) *PostUpdateOne {
	puo.mutation.SetPostStatusID(s)
	return puo
}

// SetNillablePostStatusID sets the "post_status_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillablePostStatusID(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetPostStatusID(*s)
	}
	return puo
}

// ClearPostStatusID clears the value of the "post_status_id" field.
func (puo *PostUpdateOne) ClearPostStatusID() *PostUpdateOne {
	puo.mutation.ClearPostStatusID()
	return puo
}

// SetPostTypeID sets the "post_type_id" field.
func (puo *PostUpdateOne) SetPostTypeID(s string) *PostUpdateOne {
	puo.mutation.SetPostTypeID(s)
	return puo
}

// SetNillablePostTypeID sets the "post_type_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillablePostTypeID(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetPostTypeID(*s)
	}
	return puo
}

// ClearPostTypeID clears the value of the "post_type_id" field.
func (puo *PostUpdateOne) ClearPostTypeID() *PostUpdateOne {
	puo.mutation.ClearPostTypeID()
	return puo
}

// SetPrimaryCategoryID sets the "primary_category_id" field.
func (puo *PostUpdateOne) SetPrimaryCategoryID(s string) *PostUpdateOne {
	puo.mutation.SetPrimaryCategoryID(s)
	return puo
}

// SetNillablePrimaryCategoryID sets the "primary_category_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillablePrimaryCategoryID(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetPrimaryCategoryID(*s)
	}
	return puo
}

// ClearPrimaryCategoryID clears the value of the "primary_category_id" field.
func (puo *PostUpdateOne) ClearPrimaryCategoryID() *PostUpdateOne {
	puo.mutation.ClearPrimaryCategoryID()
	return puo
}

// SetHeadline sets the "headline" field.
func (puo *PostUpdateOne) SetHeadline(s string) *PostUpdateOne {
	puo.mutation.SetHeadline(s)
	return puo
}

// SetNillableHeadline sets the "headline" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableHeadline(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetHeadline(*s)
	}
	return puo
}

// ClearHeadline clears the value of the "headline" field.
func (puo *PostUpdateOne) ClearHeadline() *PostUpdateOne {
	puo.mutation.ClearHeadline()
	return puo
}

// SetExcerpt sets the "excerpt" field.
func (puo *PostUpdateOne) SetExcerpt(s string) *PostUpdateOne {
	puo.mutation.SetExcerpt(s)
	return puo
}

// SetNillableExcerpt sets the "excerpt" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableExcerpt(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetExcerpt(*s)
	}
	return puo
}

// ClearExcerpt clears the value of the "excerpt" field.
func (puo *PostUpdateOne) ClearExcerpt() *PostUpdateOne {
	puo.mutation.ClearExcerpt()
	return puo
}

// SetContent sets the "content" field.
func (puo *PostUpdateOne) SetContent(s string) *PostUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableContent(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetContent(*s)
	}
	return puo
}

// ClearContent clears the value of the "content" field.
func (puo *PostUpdateOne) ClearContent() *PostUpdateOne {
	puo.mutation.ClearContent()
	return puo
}

// SetMetaTitle sets the "meta_title" field.
func (puo *PostUpdateOne) SetMetaTitle(s string) *PostUpdateOne {
	puo.mutation.SetMetaTitle(s)
	return puo
}

// SetNillableMetaTitle sets the "meta_title" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableMetaTitle(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetMetaTitle(*s)
	}
	return puo
}

// ClearMetaTitle clears the value of the "meta_title" field.
func (puo *PostUpdateOne) ClearMetaTitle() *PostUpdateOne {
	puo.mutation.ClearMetaTitle()
	return puo
}

// SetMetaDescr sets the "meta_descr" field.
func (puo *PostUpdateOne) SetMetaDescr(s string) *PostUpdateOne {
	puo.mutation.SetMetaDescr(s)
	return puo
}

// SetNillableMetaDescr sets the "meta_descr" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableMetaDescr(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetMetaDescr(*s)
	}
	return puo
}

// ClearMetaDescr clears the value of the "meta_descr" field.
func (puo *PostUpdateOne) ClearMetaDescr() *PostUpdateOne {
	puo.mutation.ClearMetaDescr()
	return puo
}

// SetMetaCanonicalURL sets the "meta_canonical_url" field.
func (puo *PostUpdateOne) SetMetaCanonicalURL(s string) *PostUpdateOne {
	puo.mutation.SetMetaCanonicalURL(s)
	return puo
}

// SetNillableMetaCanonicalURL sets the "meta_canonical_url" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableMetaCanonicalURL(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetMetaCanonicalURL(*s)
	}
	return puo
}

// ClearMetaCanonicalURL clears the value of the "meta_canonical_url" field.
func (puo *PostUpdateOne) ClearMetaCanonicalURL() *PostUpdateOne {
	puo.mutation.ClearMetaCanonicalURL()
	return puo
}

// SetMetaRobots sets the "meta_robots" field.
func (puo *PostUpdateOne) SetMetaRobots(s string) *PostUpdateOne {
	puo.mutation.SetMetaRobots(s)
	return puo
}

// SetNillableMetaRobots sets the "meta_robots" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableMetaRobots(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetMetaRobots(*s)
	}
	return puo
}

// ClearMetaRobots clears the value of the "meta_robots" field.
func (puo *PostUpdateOne) ClearMetaRobots() *PostUpdateOne {
	puo.mutation.ClearMetaRobots()
	return puo
}

// SetCustom sets the "custom" field.
func (puo *PostUpdateOne) SetCustom(m map[string]interface{}) *PostUpdateOne {
	puo.mutation.SetCustom(m)
	return puo
}

// ClearCustom clears the value of the "custom" field.
func (puo *PostUpdateOne) ClearCustom() *PostUpdateOne {
	puo.mutation.ClearCustom()
	return puo
}

// SetPostStatus sets the "post_status" edge to the PostStatus entity.
func (puo *PostUpdateOne) SetPostStatus(p *PostStatus) *PostUpdateOne {
	return puo.SetPostStatusID(p.ID)
}

// SetPostType sets the "post_type" edge to the PostType entity.
func (puo *PostUpdateOne) SetPostType(p *PostType) *PostUpdateOne {
	return puo.SetPostTypeID(p.ID)
}

// SetPrimaryCategory sets the "primary_category" edge to the PostCategory entity.
func (puo *PostUpdateOne) SetPrimaryCategory(p *PostCategory) *PostUpdateOne {
	return puo.SetPrimaryCategoryID(p.ID)
}

// AddPostTagIDs adds the "post_tags" edge to the PostTag entity by IDs.
func (puo *PostUpdateOne) AddPostTagIDs(ids ...string) *PostUpdateOne {
	puo.mutation.AddPostTagIDs(ids...)
	return puo
}

// AddPostTags adds the "post_tags" edges to the PostTag entity.
func (puo *PostUpdateOne) AddPostTags(p ...*PostTag) *PostUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPostTagIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearPostStatus clears the "post_status" edge to the PostStatus entity.
func (puo *PostUpdateOne) ClearPostStatus() *PostUpdateOne {
	puo.mutation.ClearPostStatus()
	return puo
}

// ClearPostType clears the "post_type" edge to the PostType entity.
func (puo *PostUpdateOne) ClearPostType() *PostUpdateOne {
	puo.mutation.ClearPostType()
	return puo
}

// ClearPrimaryCategory clears the "primary_category" edge to the PostCategory entity.
func (puo *PostUpdateOne) ClearPrimaryCategory() *PostUpdateOne {
	puo.mutation.ClearPrimaryCategory()
	return puo
}

// ClearPostTags clears all "post_tags" edges to the PostTag entity.
func (puo *PostUpdateOne) ClearPostTags() *PostUpdateOne {
	puo.mutation.ClearPostTags()
	return puo
}

// RemovePostTagIDs removes the "post_tags" edge to PostTag entities by IDs.
func (puo *PostUpdateOne) RemovePostTagIDs(ids ...string) *PostUpdateOne {
	puo.mutation.RemovePostTagIDs(ids...)
	return puo
}

// RemovePostTags removes "post_tags" edges to PostTag entities.
func (puo *PostUpdateOne) RemovePostTags(p ...*PostTag) *PostUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePostTagIDs(ids...)
}

// Where appends a list predicates to the PostUpdate builder.
func (puo *PostUpdateOne) Where(ps ...predicate.Post) *PostUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PostUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok && !puo.mutation.UpdatedAtCleared() {
		v := post.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PostUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PostUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if puo.mutation.CreatedAtCleared() {
		_spec.ClearField(post.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.UpdatedAtCleared() {
		_spec.ClearField(post.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.AppID(); ok {
		_spec.SetField(post.FieldAppID, field.TypeString, value)
	}
	if puo.mutation.AppIDCleared() {
		_spec.ClearField(post.FieldAppID, field.TypeString)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(post.FieldName, field.TypeString, value)
	}
	if puo.mutation.NameCleared() {
		_spec.ClearField(post.FieldName, field.TypeString)
	}
	if value, ok := puo.mutation.Slug(); ok {
		_spec.SetField(post.FieldSlug, field.TypeString, value)
	}
	if puo.mutation.SlugCleared() {
		_spec.ClearField(post.FieldSlug, field.TypeString)
	}
	if value, ok := puo.mutation.Headline(); ok {
		_spec.SetField(post.FieldHeadline, field.TypeString, value)
	}
	if puo.mutation.HeadlineCleared() {
		_spec.ClearField(post.FieldHeadline, field.TypeString)
	}
	if value, ok := puo.mutation.Excerpt(); ok {
		_spec.SetField(post.FieldExcerpt, field.TypeString, value)
	}
	if puo.mutation.ExcerptCleared() {
		_spec.ClearField(post.FieldExcerpt, field.TypeString)
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if puo.mutation.ContentCleared() {
		_spec.ClearField(post.FieldContent, field.TypeString)
	}
	if value, ok := puo.mutation.MetaTitle(); ok {
		_spec.SetField(post.FieldMetaTitle, field.TypeString, value)
	}
	if puo.mutation.MetaTitleCleared() {
		_spec.ClearField(post.FieldMetaTitle, field.TypeString)
	}
	if value, ok := puo.mutation.MetaDescr(); ok {
		_spec.SetField(post.FieldMetaDescr, field.TypeString, value)
	}
	if puo.mutation.MetaDescrCleared() {
		_spec.ClearField(post.FieldMetaDescr, field.TypeString)
	}
	if value, ok := puo.mutation.MetaCanonicalURL(); ok {
		_spec.SetField(post.FieldMetaCanonicalURL, field.TypeString, value)
	}
	if puo.mutation.MetaCanonicalURLCleared() {
		_spec.ClearField(post.FieldMetaCanonicalURL, field.TypeString)
	}
	if value, ok := puo.mutation.MetaRobots(); ok {
		_spec.SetField(post.FieldMetaRobots, field.TypeString, value)
	}
	if puo.mutation.MetaRobotsCleared() {
		_spec.ClearField(post.FieldMetaRobots, field.TypeString)
	}
	if value, ok := puo.mutation.Custom(); ok {
		_spec.SetField(post.FieldCustom, field.TypeJSON, value)
	}
	if puo.mutation.CustomCleared() {
		_spec.ClearField(post.FieldCustom, field.TypeJSON)
	}
	if puo.mutation.PostStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PostStatusTable,
			Columns: []string{post.PostStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poststatus.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PostStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PostStatusTable,
			Columns: []string{post.PostStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poststatus.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PostTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PostTypeTable,
			Columns: []string{post.PostTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttype.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PostTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PostTypeTable,
			Columns: []string{post.PostTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttype.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PrimaryCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PrimaryCategoryTable,
			Columns: []string{post.PrimaryCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PrimaryCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PrimaryCategoryTable,
			Columns: []string{post.PrimaryCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PostTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.PostTagsTable,
			Columns: post.PostTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttag.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedPostTagsIDs(); len(nodes) > 0 && !puo.mutation.PostTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.PostTagsTable,
			Columns: post.PostTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PostTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.PostTagsTable,
			Columns: post.PostTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posttag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
