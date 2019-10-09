// package types defines the different interfaces used in the GraphQL implementation
package types

type ConnectionArgument struct {
	After  *string
	First  *int
	Before *string
	Last   *int
}
