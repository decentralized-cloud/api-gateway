import { GraphQLObjectType, GraphQLNonNull, GraphQLList } from 'graphql';
import NodeCondition from './NodeCondition';
import NodeAddress from './NodeAddress';
import NodeSystemInfo from './NodeSystemInfo';

export default new GraphQLObjectType({
	name: 'NodeStatus',
	description: 'Contains information about the current status of a node',
	fields: {
		conditions: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(NodeCondition))),
			description: 'Conditions is an array of current observed node conditions',
		},
		addresses: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(NodeAddress))),
			description: 'Addresses is the list of addresses reachable to the node',
		},
		nodeInfo: {
			type: new GraphQLNonNull(NodeSystemInfo),
			description: 'NodeInfo is the set of ids/uuids to uniquely identify the node',
		},
	},
});
