// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"saas/gen/ent/media"
	"saas/gen/ent/mediable"
	"saas/gen/ent/post"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (m *MediaQuery) CollectFields(ctx context.Context, satisfies ...string) (*MediaQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return m, nil
	}
	if err := m.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MediaQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(media.Columns))
		selectedFields = []string{media.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "mediables":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MediableClient{config: m.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, mediableImplementors)...); err != nil {
				return err
			}
			m.WithNamedMediables(alias, func(wq *MediableQuery) {
				*wq = *query
			})
		case "appID":
			if _, ok := fieldSeen[media.FieldAppID]; !ok {
				selectedFields = append(selectedFields, media.FieldAppID)
				fieldSeen[media.FieldAppID] = struct{}{}
			}
		case "disk":
			if _, ok := fieldSeen[media.FieldDisk]; !ok {
				selectedFields = append(selectedFields, media.FieldDisk)
				fieldSeen[media.FieldDisk] = struct{}{}
			}
		case "directory":
			if _, ok := fieldSeen[media.FieldDirectory]; !ok {
				selectedFields = append(selectedFields, media.FieldDirectory)
				fieldSeen[media.FieldDirectory] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[media.FieldName]; !ok {
				selectedFields = append(selectedFields, media.FieldName)
				fieldSeen[media.FieldName] = struct{}{}
			}
		case "originalName":
			if _, ok := fieldSeen[media.FieldOriginalName]; !ok {
				selectedFields = append(selectedFields, media.FieldOriginalName)
				fieldSeen[media.FieldOriginalName] = struct{}{}
			}
		case "extension":
			if _, ok := fieldSeen[media.FieldExtension]; !ok {
				selectedFields = append(selectedFields, media.FieldExtension)
				fieldSeen[media.FieldExtension] = struct{}{}
			}
		case "mimeType":
			if _, ok := fieldSeen[media.FieldMimeType]; !ok {
				selectedFields = append(selectedFields, media.FieldMimeType)
				fieldSeen[media.FieldMimeType] = struct{}{}
			}
		case "aggregateType":
			if _, ok := fieldSeen[media.FieldAggregateType]; !ok {
				selectedFields = append(selectedFields, media.FieldAggregateType)
				fieldSeen[media.FieldAggregateType] = struct{}{}
			}
		case "size":
			if _, ok := fieldSeen[media.FieldSize]; !ok {
				selectedFields = append(selectedFields, media.FieldSize)
				fieldSeen[media.FieldSize] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[media.FieldDescription]; !ok {
				selectedFields = append(selectedFields, media.FieldDescription)
				fieldSeen[media.FieldDescription] = struct{}{}
			}
		case "isVariant":
			if _, ok := fieldSeen[media.FieldIsVariant]; !ok {
				selectedFields = append(selectedFields, media.FieldIsVariant)
				fieldSeen[media.FieldIsVariant] = struct{}{}
			}
		case "variantName":
			if _, ok := fieldSeen[media.FieldVariantName]; !ok {
				selectedFields = append(selectedFields, media.FieldVariantName)
				fieldSeen[media.FieldVariantName] = struct{}{}
			}
		case "originalMediaID":
			if _, ok := fieldSeen[media.FieldOriginalMediaID]; !ok {
				selectedFields = append(selectedFields, media.FieldOriginalMediaID)
				fieldSeen[media.FieldOriginalMediaID] = struct{}{}
			}
		case "checksum":
			if _, ok := fieldSeen[media.FieldChecksum]; !ok {
				selectedFields = append(selectedFields, media.FieldChecksum)
				fieldSeen[media.FieldChecksum] = struct{}{}
			}
		case "workspaceID":
			if _, ok := fieldSeen[media.FieldWorkspaceID]; !ok {
				selectedFields = append(selectedFields, media.FieldWorkspaceID)
				fieldSeen[media.FieldWorkspaceID] = struct{}{}
			}
		case "alt":
			if _, ok := fieldSeen[media.FieldAlt]; !ok {
				selectedFields = append(selectedFields, media.FieldAlt)
				fieldSeen[media.FieldAlt] = struct{}{}
			}
		case "uid":
			if _, ok := fieldSeen[media.FieldUID]; !ok {
				selectedFields = append(selectedFields, media.FieldUID)
				fieldSeen[media.FieldUID] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[media.FieldStatus]; !ok {
				selectedFields = append(selectedFields, media.FieldStatus)
				fieldSeen[media.FieldStatus] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		m.Select(selectedFields...)
	}
	return nil
}

type mediaPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MediaPaginateOption
}

func newMediaPaginateArgs(rv map[string]any) *mediaPaginateArgs {
	args := &mediaPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (m *MediableQuery) CollectFields(ctx context.Context, satisfies ...string) (*MediableQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return m, nil
	}
	if err := m.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MediableQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(mediable.Columns))
		selectedFields = []string{mediable.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "media":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MediaClient{config: m.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, mediaImplementors)...); err != nil {
				return err
			}
			m.withMedia = query
			if _, ok := fieldSeen[mediable.FieldMediaID]; !ok {
				selectedFields = append(selectedFields, mediable.FieldMediaID)
				fieldSeen[mediable.FieldMediaID] = struct{}{}
			}
		case "appID":
			if _, ok := fieldSeen[mediable.FieldAppID]; !ok {
				selectedFields = append(selectedFields, mediable.FieldAppID)
				fieldSeen[mediable.FieldAppID] = struct{}{}
			}
		case "mediaID":
			if _, ok := fieldSeen[mediable.FieldMediaID]; !ok {
				selectedFields = append(selectedFields, mediable.FieldMediaID)
				fieldSeen[mediable.FieldMediaID] = struct{}{}
			}
		case "mediableID":
			if _, ok := fieldSeen[mediable.FieldMediableID]; !ok {
				selectedFields = append(selectedFields, mediable.FieldMediableID)
				fieldSeen[mediable.FieldMediableID] = struct{}{}
			}
		case "mediableType":
			if _, ok := fieldSeen[mediable.FieldMediableType]; !ok {
				selectedFields = append(selectedFields, mediable.FieldMediableType)
				fieldSeen[mediable.FieldMediableType] = struct{}{}
			}
		case "tag":
			if _, ok := fieldSeen[mediable.FieldTag]; !ok {
				selectedFields = append(selectedFields, mediable.FieldTag)
				fieldSeen[mediable.FieldTag] = struct{}{}
			}
		case "order":
			if _, ok := fieldSeen[mediable.FieldOrder]; !ok {
				selectedFields = append(selectedFields, mediable.FieldOrder)
				fieldSeen[mediable.FieldOrder] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		m.Select(selectedFields...)
	}
	return nil
}

type mediablePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MediablePaginateOption
}

func newMediablePaginateArgs(rv map[string]any) *mediablePaginateArgs {
	args := &mediablePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (po *PostQuery) CollectFields(ctx context.Context, satisfies ...string) (*PostQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return po, nil
	}
	if err := po.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return po, nil
}

func (po *PostQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(post.Columns))
		selectedFields = []string{post.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "name":
			if _, ok := fieldSeen[post.FieldName]; !ok {
				selectedFields = append(selectedFields, post.FieldName)
				fieldSeen[post.FieldName] = struct{}{}
			}
		case "slug":
			if _, ok := fieldSeen[post.FieldSlug]; !ok {
				selectedFields = append(selectedFields, post.FieldSlug)
				fieldSeen[post.FieldSlug] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		po.Select(selectedFields...)
	}
	return nil
}

type postPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PostPaginateOption
}

func newPostPaginateArgs(rv map[string]any) *postPaginateArgs {
	args := &postPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PostWhereInput); ok {
		args.opts = append(args.opts, WithPostFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
