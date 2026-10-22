package nullable_test

import (
	"strings"
	"testing"

	"github.com/Valdenirmezadri/nullable"
	"github.com/graphql-go/graphql"
)

func TestGraphqlExtension(t *testing.T) {
	type user struct {
		Name nullable.String `json:"name"`
		Plan nullable.Uint   `json:"plan"`
	}

	query := func(nameType, planType graphql.Output) *graphql.Result {
		schema, err := graphql.NewSchema(graphql.SchemaConfig{
			Extensions: []graphql.Extension{nullable.GraphqlExtension{}},
			Query: graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: graphql.NewObject(graphql.ObjectConfig{Name: "user", Fields: graphql.Fields{
						"name": &graphql.Field{Type: nameType},
						"plan": &graphql.Field{Type: planType},
					}}),
					Resolve: func(graphql.ResolveParams) (interface{}, error) {
						return user{Name: nullable.NewString("ana"), Plan: nullable.NewUint(3)}, nil
					},
				},
			}}),
		})
		if err != nil {
			t.Fatal(err)
		}
		return graphql.Do(graphql.Params{Schema: schema, RequestString: `{ user { name plan } }`})
	}

	if r := query(nullable.GraphqlStringType, graphql.NewNonNull(nullable.GraphqlUintType)); len(r.Errors) > 0 {
		t.Fatalf("scalar certo não deve dar erro: %v", r.Errors)
	}

	r := query(graphql.String, graphql.Int)
	if len(r.Errors) != 2 {
		t.Fatalf("esperava 2 erros (name e plan), veio %v", r.Errors)
	}
	if !strings.Contains(r.Errors[0].Message, "user.") || !strings.Contains(r.Errors[0].Message, "nullable.") {
		t.Fatalf("mensagem sem campo/tipo: %s", r.Errors[0].Message)
	}
}

// tipo próprio em cima de um nullable (como models.NotificationStatus) herda
// IsNull() pelo Default embutido, mas tem scalar próprio: não é erro
func TestGraphqlExtensionSkipsDerivedTypes(t *testing.T) {
	type status nullable.Uint8

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Extensions: []graphql.Extension{nullable.GraphqlExtension{}},
		Query: graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(graphql.ResolveParams) (interface{}, error) {
					return status(nullable.NewUint8(2)), nil
				},
			},
		}}),
	})
	if err != nil {
		t.Fatal(err)
	}

	if r := graphql.Do(graphql.Params{Schema: schema, RequestString: `{ status }`}); len(r.Errors) > 0 {
		t.Fatalf("tipo derivado não deve dar erro: %v", r.Errors)
	}
}
