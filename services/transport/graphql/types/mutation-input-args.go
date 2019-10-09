// package types defines the different interfaces used in the GraphQL implementation
package types

import "github.com/graph-gophers/graphql-go"

type CreateTenantInput struct {
	Name             string
	ClientMutationId *string
}

type CreateTenantInputArgument struct {
	Input CreateTenantInput
}

type UpdateTenantInput struct {
	ID               graphql.ID
	Name             string
	ClientMutationId *string
}

type UpdateTenantInputArgument struct {
	Input UpdateTenantInput
}
