import { GraphQLEnumType } from 'graphql';

export default new GraphQLEnumType({
	name: 'EdgeClusterPodConditionType',
	description: 'The edge cluster pod condition',
	values: {
		ContainersReady: { value: 0 },
		PodInitialized: { value: 1 },
		PodReady: { value: 2 },
		PodScheduled: { value: 3 },
	},
});
