import { GraphQLObjectType, GraphQLNonNull, GraphQLString } from 'graphql';
import EdgeClusterNodeConditionType from './EdgeClusterNodeConditionType';
import EdgeClusterNodeConditionStatus from './EdgeClusterNodeConditionStatus';

export default new GraphQLObjectType({
	name: 'EdgeClusterNodeCondition',
	description: 'Condition information for a node',
	fields: {
		type: { type: new GraphQLNonNull(EdgeClusterNodeConditionType), description: 'Type of node condition' },
		status: { type: new GraphQLNonNull(EdgeClusterNodeConditionStatus), description: 'Status of the condition, one of True, False, Unknown' },
		lastHeartbeatTime: { type: new GraphQLNonNull(GraphQLString), description: 'Last time we got an update on a given condition' },
		lastTransitionTime: { type: new GraphQLNonNull(GraphQLString), description: 'Last time the condition transit from one status to another' },
		reason: { type: new GraphQLNonNull(GraphQLString), description: '(brief) reason for the condition last transition' },
		message: { type: new GraphQLNonNull(GraphQLString), description: 'Human readable message indicating details about last transition' },
	},
});
