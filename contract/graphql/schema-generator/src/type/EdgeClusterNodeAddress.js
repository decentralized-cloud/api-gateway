import { GraphQLObjectType, GraphQLNonNull, GraphQLString } from 'graphql';
import EdgeClusterNodeAddressType from './EdgeClusterNodeAddressType';

export default new GraphQLObjectType({
	name: 'EdgeClusterNodeAddress',
	description: 'The information for the edge cluster node address',
	fields: {
		NodeAddressType: {
			type: new GraphQLNonNull(EdgeClusterNodeAddressType),
			description: 'Edge cluster node address type, one of Hostname, ExternalIP or InternalIP',
		},
		Address: { type: new GraphQLNonNull(GraphQLString), description: 'The node address' },
	},
});
