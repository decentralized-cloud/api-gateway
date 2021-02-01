import { GraphQLObjectType, GraphQLNonNull, GraphQLList } from 'graphql';
import EdgeClusterNodeCondition from './EdgeClusterNodeCondition';
import EdgeClusterNodeAddress from './EdgeClusterNodeAddress';
import EdgeClusterNodeSystemInfo from './EdgeClusterNodeSystemInfo';

export default new GraphQLObjectType({
	name: 'EdgeClusterNodeStatus',
	description: 'Contains information about the current status of a node',
	fields: {
		Conditions: {
			type: new GraphQLList(new GraphQLNonNull(EdgeClusterNodeCondition)),
			description: 'Conditions is an array of current observed node conditions',
		},
		Addresses: {
			type: new GraphQLList(new GraphQLNonNull(EdgeClusterNodeAddress)),
			description: 'Addresses is the list of addresses reachable to the node',
		},
		NodeInfo: {
			type: new GraphQLNonNull(EdgeClusterNodeSystemInfo),
			description: 'NodeInfo is the set of ids/uuids to uniquely identify the node',
		},
	},
});
