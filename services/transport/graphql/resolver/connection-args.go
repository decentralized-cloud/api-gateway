// package resolver implements different GraphQL resolvers required by the GraphQL transport layer
package resolver

type connectionArgument struct {
	After  *string
	First  *int
	Before *string
	Last   *int
}
