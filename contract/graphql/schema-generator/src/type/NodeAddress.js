import { GraphQLObjectType, GraphQLNonNull, GraphQLString } from 'graphql';
import NodeAddressType from './NodeAddressType';

export default new GraphQLObjectType({
	name: 'NodeAddress',
	description: 'The information for the edge cluster node address',
	fields: {
		nodeAddressType: {
			type: new GraphQLNonNull(NodeAddressType),
			description: 'Edge cluster node address type, one of Hostname, ExternalIP or InternalIP',
		},
		address: { type: new GraphQLNonNull(GraphQLString), description: 'The node address' },
	},
});
