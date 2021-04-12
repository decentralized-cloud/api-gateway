import { GraphQLEnumType } from 'graphql';

export default new GraphQLEnumType({
	name: 'ConditionStatus',
	description: 'These are valid condition statuses',
	values: {
		True: { value: 0, description: 'True means a resource is in the condition' },
		False: { value: 1, description: 'False means a resource is not in the condition' },
		Unknown: { value: 2, description: 'Unknown means kubernetes cannot decide if a resource is in the condition or not' },
	},
});
