import { GraphQLObjectType, GraphQLNonNull, GraphQLString } from 'graphql';
import NodeConditionType from './NodeConditionType';
import ConditionStatus from './ConditionStatus';

export default new GraphQLObjectType({
	name: 'NodeCondition',
	description: ' Current service state of node',
	fields: {
		type: { type: new GraphQLNonNull(NodeConditionType), description: 'Type is the type of the condition' },
		status: { type: new GraphQLNonNull(ConditionStatus), description: 'Status is the status of the condition' },
		lastHeartbeatTime: { type: new GraphQLNonNull(GraphQLString), description: 'Last time we got an update on a given condition' },
		lastTransitionTime: {
			type: new GraphQLNonNull(GraphQLString),
			description: 'Last time the condition transitioned from one status to another',
		},
		reason: { type: new GraphQLNonNull(GraphQLString), description: 'Unique, one-word, CamelCase reason for the condition last transition' },
		message: { type: new GraphQLNonNull(GraphQLString), description: 'Human-readable message indicating details about last transition' },
	},
});
