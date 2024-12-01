// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"gqlsa/graph/model"
	"lace/jsonslice"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Student_name(ctx context.Context, field graphql.CollectedField, obj *model.Student) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Student_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Student_name(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Student",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var studentImplementors = []string{"Student"}

func (ec *executionContext) _Student(ctx context.Context, sel ast.SelectionSet, obj *model.Student) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, studentImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Student")
		case "name":
			out.Values[i] = ec._Student_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNTime2timeᚐTime(ctx context.Context, v interface{}) (time.Time, error) {
	res, err := graphql.UnmarshalTime(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNTime2timeᚐTime(ctx context.Context, sel ast.SelectionSet, v time.Time) graphql.Marshaler {
	res := graphql.MarshalTime(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNUint2uint(ctx context.Context, v interface{}) (uint, error) {
	res, err := graphql.UnmarshalUint(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNUint2uint(ctx context.Context, sel ast.SelectionSet, v uint) graphql.Marshaler {
	res := graphql.MarshalUint(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalOJsonSlice2laceᚋjsonsliceᚐJsonSlice(ctx context.Context, v interface{}) (jsonslice.JsonSlice, error) {
	if v == nil {
		return nil, nil
	}
	res, err := jsonslice.UnmarshalJsonSlice(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOJsonSlice2laceᚋjsonsliceᚐJsonSlice(ctx context.Context, sel ast.SelectionSet, v jsonslice.JsonSlice) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	res := jsonslice.MarshalJsonSlice(v)
	return res
}

func (ec *executionContext) unmarshalOMap2map(ctx context.Context, v interface{}) (map[string]interface{}, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalMap(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOMap2map(ctx context.Context, sel ast.SelectionSet, v map[string]interface{}) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	res := graphql.MarshalMap(v)
	return res
}

func (ec *executionContext) unmarshalOTime2timeᚐTime(ctx context.Context, v interface{}) (time.Time, error) {
	res, err := graphql.UnmarshalTime(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOTime2timeᚐTime(ctx context.Context, sel ast.SelectionSet, v time.Time) graphql.Marshaler {
	res := graphql.MarshalTime(v)
	return res
}

func (ec *executionContext) unmarshalOTime2ᚕtimeᚐTimeᚄ(ctx context.Context, v interface{}) ([]time.Time, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]time.Time, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNTime2timeᚐTime(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalOTime2ᚕtimeᚐTimeᚄ(ctx context.Context, sel ast.SelectionSet, v []time.Time) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNTime2timeᚐTime(ctx, sel, v[i])
	}

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) unmarshalOTime2ᚖtimeᚐTime(ctx context.Context, v interface{}) (*time.Time, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalTime(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOTime2ᚖtimeᚐTime(ctx context.Context, sel ast.SelectionSet, v *time.Time) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	res := graphql.MarshalTime(*v)
	return res
}

func (ec *executionContext) unmarshalOUint2uint(ctx context.Context, v interface{}) (uint, error) {
	res, err := graphql.UnmarshalUint(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOUint2uint(ctx context.Context, sel ast.SelectionSet, v uint) graphql.Marshaler {
	res := graphql.MarshalUint(v)
	return res
}

func (ec *executionContext) unmarshalOUint2ᚕuintᚄ(ctx context.Context, v interface{}) ([]uint, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]uint, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNUint2uint(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalOUint2ᚕuintᚄ(ctx context.Context, sel ast.SelectionSet, v []uint) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNUint2uint(ctx, sel, v[i])
	}

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) unmarshalOUint2ᚖuint(ctx context.Context, v interface{}) (*uint, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalUint(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOUint2ᚖuint(ctx context.Context, sel ast.SelectionSet, v *uint) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	res := graphql.MarshalUint(*v)
	return res
}

// endregion ***************************** type.gotpl *****************************
