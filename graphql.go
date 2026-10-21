package nullable

// GraphQL (graphql-go) scalars for the nullable types.
//
// Same rule as the types: the GraphQL side never sees null, only the zero of
// the type (0, "", 0001-01-01T00:00:00Z). Input "" or null becomes the null
// value, which is saved as NULL in the database.
//
// Use: &graphql.Field{Type: nullable.GraphqlTimeType}

import (
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// names kept from htchat-server so existing schemas do not change.
// One instance per type: graphql-go rejects two scalars with the same name.
var (
	GraphqlUintType = newScalar[uint]("nullableUint64Type", "nullable.Uint as an unsigned integer, 0 when null",
		func(s string) (Uint, error) {
			v, err := strconv.ParseUint(s, 10, 64)
			return NewUint(uint(v)), err
		})

	GraphqlUint8Type = newScalar[uint8]("nullableUint8Type", "nullable.Uint8 as an unsigned integer, 0 when null",
		func(s string) (Uint8, error) {
			v, err := strconv.ParseUint(s, 10, 8)
			return NewUint8(uint8(v)), err
		})

	GraphqlUint16Type = newScalar[uint16]("nullableUint16Type", "nullable.Uint16 as an unsigned integer, 0 when null",
		func(s string) (Uint16, error) {
			v, err := strconv.ParseUint(s, 10, 16)
			return NewUint16(uint16(v)), err
		})

	GraphqlUint32Type = newScalar[uint32]("nullableUint32Type", "nullable.Uint32 as an unsigned integer, 0 when null",
		func(s string) (Uint32, error) {
			v, err := strconv.ParseUint(s, 10, 32)
			return NewUint32(uint32(v)), err
		})

	GraphqlStringType = newScalar[string]("nullableStringType", `nullable.String as a string, "" when null`,
		func(s string) (String, error) {
			return NewString(s), nil
		})

	GraphqlTimeType = newScalar[time.Time]("nullableTimeType", "nullable.Time as an RFC 3339 string, 0001-01-01T00:00:00Z when null",
		func(s string) (Time, error) {
			v, err := time.Parse(time.RFC3339Nano, s)
			return NewTime(v), err
		})
)

// newScalar builds a scalar whose Serialize calls Get() and whose inputs go
// through parse. Invalid input returns nil, so graphql-go rejects the argument.
func newScalar[T comparable, N interface{ Get() T }](name, description string, parse func(string) (N, error)) *graphql.Scalar {
	parseAny := func(value interface{}) interface{} {
		s, ok := scalarString(value)
		if !ok {
			return nil
		}
		if s == "" {
			var null N
			return null
		}
		parsed, err := parse(s)
		if err != nil {
			return nil
		}
		return parsed
	}

	return graphql.NewScalar(graphql.ScalarConfig{
		Name:        name,
		Description: description,
		Serialize: func(value interface{}) interface{} {
			if g, ok := value.(interface{ Get() T }); ok {
				return g.Get()
			}
			var zero T
			return zero
		},
		ParseValue: parseAny,
		ParseLiteral: func(valueAST ast.Value) interface{} {
			return parseAny(valueAST)
		},
	})
}

// scalarString normalizes what arrives from variables (JSON) and from literals
func scalarString(value interface{}) (string, bool) {
	switch v := value.(type) {
	case string:
		return v, true
	case *string:
		if v == nil {
			return "", false
		}
		return *v, true
	case int:
		return strconv.Itoa(v), true
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), true
	case *ast.IntValue:
		return v.Value, true
	case *ast.StringValue:
		return v.Value, true
	}
	return "", false
}
