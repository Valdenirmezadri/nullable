package nullable

import (
	"context"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

// GraphqlExtension fails the field when a nullable value goes out through a
// scalar that is not one of the nullable scalars. Without it the mistake is
// silent: nullable.String in graphql.String prints "{{ana true}}" and
// nullable.Uint in graphql.Int prints null.
//
// Register once per schema: graphql.SchemaConfig{Extensions: []graphql.Extension{nullable.GraphqlExtension{}}}
//
// It fails the query (error in the response), not the process: one forgotten
// field must not take the whole server down.
type GraphqlExtension struct{}

var nullableScalars = map[*graphql.Scalar]bool{
	GraphqlUintType:   true,
	GraphqlUint8Type:  true,
	GraphqlUint16Type: true,
	GraphqlUint32Type: true,
	GraphqlStringType: true,
	GraphqlTimeType:   true,
}

func (GraphqlExtension) ResolveFieldDidStart(ctx context.Context, info *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
	return ctx, func(value interface{}, _ error) {
		if !isNullableType(value) {
			return
		}

		returnType := info.ReturnType
		if nonNull, ok := returnType.(*graphql.NonNull); ok {
			returnType = nonNull.OfType
		}
		if scalar, ok := returnType.(*graphql.Scalar); ok && nullableScalars[scalar] {
			return
		}

		// graphql-go recovers this panic (it must be an error) and adds it to the response
		panic(fmt.Errorf("field %s.%s is %T but the schema uses %s: use the nullable.Graphql*Type scalar",
			info.ParentType.Name(), info.FieldName, value, info.ReturnType))
	}
}

// isNullableType matches only this package's types. A type built on one
// (type Status nullable.Uint8) also gets IsNull() through the embedded
// Default, but has its own Value/Scan and its own scalar, so it is skipped.
func isNullableType(value interface{}) bool {
	switch value.(type) {
	case Uint, Uint8, Uint16, Uint32, String, Time,
		*Uint, *Uint8, *Uint16, *Uint32, *String, *Time:
		return true
	}
	return false
}

func (GraphqlExtension) Name() string { return "nullable" }

func (GraphqlExtension) Init(ctx context.Context, _ *graphql.Params) context.Context { return ctx }

func (GraphqlExtension) ParseDidStart(ctx context.Context) (context.Context, graphql.ParseFinishFunc) {
	return ctx, func(error) {}
}

func (GraphqlExtension) ValidationDidStart(ctx context.Context) (context.Context, graphql.ValidationFinishFunc) {
	return ctx, func([]gqlerrors.FormattedError) {}
}

func (GraphqlExtension) ExecutionDidStart(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc) {
	return ctx, func(*graphql.Result) {}
}

func (GraphqlExtension) HasResult() bool { return false }

func (GraphqlExtension) GetResult(context.Context) interface{} { return nil }
