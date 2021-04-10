import { GraphQLObjectType, GraphQLNonNull } from 'graphql';
import ObjectMeta from './ObjectMeta';
import NodeStatus from './NodeStatus';

export default new GraphQLObjectType({
	name: 'EdgeClusterNode',
	description: 'Contains information about the edge cluster node',
	fields: {
		metadata: {
			type: new GraphQLNonNull(ObjectMeta),
			description: 'The node metadata',
		},
		status: {
			type: new GraphQLNonNull(NodeStatus),
			description: 'The most recently observed status of the node',
		},
	},
});
