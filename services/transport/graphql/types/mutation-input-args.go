// package types defines the different interfaces used in the GraphQL implementation
package types

type CreateTenantInput struct {
	Name             string
	ClientMutationId *string
}

type CreateTenantInputArgument struct {
	Input CreateTenantInput
}
