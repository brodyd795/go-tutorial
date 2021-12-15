package graph

import "github.com/brodyd795/go-tutorials/graphql-tutorial/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CharacterStore map[string]model.Character
}
