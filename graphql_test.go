package nullable_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Valdenirmezadri/nullable"
	"github.com/graphql-go/graphql"
)

type event struct {
	ID nullable.Uint   `json:"id"`
	At nullable.Time   `json:"at"`
	Nm nullable.String `json:"nm"`
}

// roda uma mutation de verdade: argumento entra pelo scalar, campo sai pelo scalar
func run(t *testing.T, query string, vars map[string]interface{}) map[string]interface{} {
	t.Helper()

	eventType := graphql.NewObject(graphql.ObjectConfig{
		Name: "event",
		Fields: graphql.Fields{
			"id": &graphql.Field{Type: nullable.GraphqlUintType},
			"at": &graphql.Field{Type: nullable.GraphqlTimeType},
			"nm": &graphql.Field{Type: nullable.GraphqlStringType},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: graphql.Fields{
			"ok": &graphql.Field{Type: graphql.Boolean},
		}}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: graphql.Fields{
			"echo": &graphql.Field{
				Type: eventType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: nullable.GraphqlUintType},
					"at": &graphql.ArgumentConfig{Type: nullable.GraphqlTimeType},
					"nm": &graphql.ArgumentConfig{Type: nullable.GraphqlStringType},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					e := event{}
					// argumento chega como tipo comum; quem converte é o usecase
					id, _ := p.Args["id"].(int)
					at, _ := p.Args["at"].(time.Time)
					nm, _ := p.Args["nm"].(string)
					e.ID = nullable.NewUint(uint(id))
					e.At = nullable.NewTime(at)
					e.Nm = nullable.NewString(nm)
					return e, nil
				},
			},
		}}),
	})
	if err != nil {
		t.Fatal(err)
	}

	r := graphql.Do(graphql.Params{Schema: schema, RequestString: query, VariableValues: vars})
	if len(r.Errors) > 0 {
		t.Fatalf("%v", r.Errors)
	}
	raw, _ := json.Marshal(r.Data)
	var out map[string]map[string]interface{}
	json.Unmarshal(raw, &out)
	return out["echo"]
}

func TestValues(t *testing.T) {
	got := run(t, `mutation($id: nullableUint64Type, $at: nullableTimeType, $nm: nullableStringType) {
		echo(id: $id, at: $at, nm: $nm) { id at nm } }`,
		map[string]interface{}{"id": float64(7), "at": "2026-12-20T18:00:00-03:00", "nm": "ana"})

	if got["id"] != float64(7) || got["nm"] != "ana" {
		t.Fatalf("got %v", got)
	}
	at, _ := time.Parse(time.RFC3339, got["at"].(string))
	if !at.Equal(time.Date(2026, 12, 20, 21, 0, 0, 0, time.UTC)) {
		t.Fatalf("at = %v", got["at"])
	}
}

func TestNullNeverLeaves(t *testing.T) {
	for _, vars := range []map[string]interface{}{
		{"id": nil, "at": nil, "nm": nil},
		{"id": float64(0), "at": "", "nm": ""},
		{},
	} {
		got := run(t, `mutation($id: nullableUint64Type, $at: nullableTimeType, $nm: nullableStringType) {
			echo(id: $id, at: $at, nm: $nm) { id at nm } }`, vars)

		if got["id"] != float64(0) || got["nm"] != "" || got["at"] != "0001-01-01T00:00:00Z" {
			t.Fatalf("vars %v: got %v", vars, got)
		}
	}
}

func TestLiteral(t *testing.T) {
	got := run(t, `mutation { echo(id: 5, at: "2026-12-20T18:00:00Z", nm: "x") { id at nm } }`, nil)
	if got["id"] != float64(5) || got["at"] != "2026-12-20T18:00:00Z" {
		t.Fatalf("got %v", got)
	}
}

func TestInvalidIsRejected(t *testing.T) {
	r := graphql.Do(graphql.Params{
		Schema:        mustSchema(t),
		RequestString: `mutation { echo(at: "amanhã") { at } }`,
	})
	if len(r.Errors) == 0 {
		t.Fatal("data inválida deveria dar erro")
	}
}

func mustSchema(t *testing.T) graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: graphql.Fields{
			"ok": &graphql.Field{Type: graphql.Boolean},
		}}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: graphql.Fields{
			"echo": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{Name: "e", Fields: graphql.Fields{
					"at": &graphql.Field{Type: nullable.GraphqlTimeType},
				}}),
				Args:    graphql.FieldConfigArgument{"at": &graphql.ArgumentConfig{Type: nullable.GraphqlTimeType}},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) { return event{}, nil },
			},
		}}),
	})
	if err != nil {
		t.Fatal(err)
	}
	return schema
}
