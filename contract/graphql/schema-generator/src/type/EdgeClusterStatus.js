import { GraphQLEnumType } from 'graphql';

export default new GraphQLEnumType({
	name: 'EdgeClusterStatus',
	values: {
		Provisioning: { value: 0 },
		Ready: { value: 1 },
		Deleting: { value: 2 },
	},
});
