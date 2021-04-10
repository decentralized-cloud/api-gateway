import { GraphQLObjectType, GraphQLNonNull, GraphQLString } from 'graphql';
import ConditionStatus from './ConditionStatus';

export default new GraphQLObjectType({
	name: 'ServiceCondition',
	description: ' Current service state of pod',
	fields: {
		type: { type: new GraphQLNonNull(GraphQLString), description: 'Type is the type of the condition' },
		status: { type: new GraphQLNonNull(ConditionStatus), description: 'Status is the status of the condition' },
		lastTransitionTime: {
			type: new GraphQLNonNull(GraphQLString),
			description: 'Last time the condition transitioned from one status to another',
		},
		reason: { type: new GraphQLNonNull(GraphQLString), description: 'Unique, one-word, CamelCase reason for the condition last transition' },
		message: { type: new GraphQLNonNull(GraphQLString), description: 'Human-readable message indicating details about last transition' },
	},
});
