import { GraphQLObjectType, GraphQLNonNull } from 'graphql';
import EdgeClusterObjectMetadata from './EdgeClusterObjectMetadata';
import EdgeClusterNodeStatus from './EdgeClusterNodeStatus';

export default new GraphQLObjectType({
	name: 'EdgeClusterNode',
	description: 'Contains information about the edge cluster node',
	fields: {
		metadata: {
			type: new GraphQLNonNull(EdgeClusterObjectMetadata),
			description: 'The node metadata',
		},
		status: {
			type: new GraphQLNonNull(EdgeClusterNodeStatus),
			description: 'The most recently observed status of the node',
		},
	},
});
