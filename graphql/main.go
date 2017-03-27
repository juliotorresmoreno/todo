package graphql

import "fmt"
import "./category"
import "github.com/graphql-go/graphql"

var schema graphql.Schema

//ExecuteQuery Ejecuta las consultas
func ExecuteQuery(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func concat(store graphql.Fields, append graphql.Fields) graphql.Fields {
	for i, v := range append {
		store[i] = v
	}
	return store
}

func init() {
	var query = make(graphql.Fields, 0)
	var mutation = make(graphql.Fields, 0)
	query = concat(query, category.GetData)
	mutation = concat(mutation, category.SetData)

	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: query,
	})

	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: mutation,
	})

	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}
