import { GraphQLEnumType } from 'graphql';

export default new GraphQLEnumType({
	name: 'EdgeClusterType',
	description: 'The different cluster types',
	values: {
		K3S: { value: 0, description: 'K3S cluster' },
	},
});
