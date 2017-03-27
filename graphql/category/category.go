package category

import (
	"../../config"
	"../../crud"
	"../../models"
	"github.com/graphql-go/graphql"
)

var orm = config.GetEngine()
var tipos = map[string]graphql.Type{
	"id":         graphql.Int,
	"state":      graphql.Int,
	"nombre":     graphql.String,
	"lang":       graphql.String,
	"date":       graphql.String,
	"created_at": graphql.String,
	"updated_at": graphql.String,
}

var categoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: tipos["id"],
		},
		"state": &graphql.Field{
			Type: tipos["state"],
		},
		"nombre": &graphql.Field{
			Type: tipos["nombre"],
		},
		"lang": &graphql.Field{
			Type: tipos["lang"],
		},
		"date": &graphql.Field{
			Type: tipos["date"],
		},
		"created_at": &graphql.Field{
			Type: tipos["created_at"],
		},
		"updated_at": &graphql.Field{
			Type: tipos["updated_at"],
		},
	},
})

//GetData Obtiene los datos
var GetData = graphql.Fields{
	"categoryList": &graphql.Field{
		Type:        graphql.NewList(categoryType),
		Description: "List of category",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			data := make([]models.Category, 0)
			err := crud.GraphQLGet(params, orm, &data)
			return data, err
		},
	},
}

//SetData Establece los datos
var SetData = graphql.Fields{
	"createCategory": &graphql.Field{
		Type:        categoryType,
		Description: "Create new category",
		Args: graphql.FieldConfigArgument{
			"nombre": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(tipos["nombre"]),
			},
			"lang": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(tipos["lang"]),
			},
			"date": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(tipos["date"]),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			data := models.Category{}
			_, err := crud.GraphQLPut(params, orm, &data)
			return data, err
		},
	},
	"updateCategory": &graphql.Field{
		Type:        categoryType,
		Description: "Update existing category, mark it done or not done",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(tipos["id"]),
			},
			"nombre": &graphql.ArgumentConfig{
				Type: tipos["nombre"],
			},
			"lang": &graphql.ArgumentConfig{
				Type: tipos["lang"],
			},
			"date": &graphql.ArgumentConfig{
				Type: tipos["date"],
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			data := models.Category{}
			crud.GraphQLPost(params, orm, &data)
			return data, nil
		},
	},
	"deleteCategory": &graphql.Field{
		Type:        categoryType,
		Description: "Delete existing category",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(tipos["id"]),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			data := models.Category{}
			crud.GraphQLDelete(params, orm, &data)
			return data, nil
		},
	},
}
