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

type DeleteTenantInput struct {
	ID               graphql.ID
	ClientMutationId *string
}

type DeleteTenantInputArgument struct {
	Input DeleteTenantInput
}

type CreateEdgeClusterInput struct {
	Name             string
	ClientMutationId *string
}

type CreateEdgeClusterInputArgument struct {
	Input CreateEdgeClusterInput
}

type UpdateEdgeClusterInput struct {
	ID               graphql.ID
	Name             string
	ClientMutationId *string
}

type UpdateEdgeClusterInputArgument struct {
	Input UpdateEdgeClusterInput
}

type DeleteEdgeClusterInput struct {
	ID               graphql.ID
	ClientMutationId *string
}

type DeleteEdgeClusterInputArgument struct {
	Input DeleteEdgeClusterInput
}
