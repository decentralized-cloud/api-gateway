// packae relay implements common relay types used in the GraphQL transport layer
package relay

type ConnectionArgument struct {
	After  *string
	First  *int
	Before *string
	Last   *int
}
