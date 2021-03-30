import { GraphQLObjectType, GraphQLNonNull, GraphQLList, GraphQLString } from 'graphql';
import EdgeClusterPodCondition from './EdgeClusterPodCondition';

export default new GraphQLObjectType({
	name: 'EdgeClusterPodStatus',
	description: 'Contains the most recently observed status of the existing edge cluster pod',
	fields: {
		hostIP: { type: new GraphQLNonNull(GraphQLString), description: 'IP address allocated to the pod. Routable at least within the cluster.' },
		podIP: { type: new GraphQLNonNull(GraphQLString), description: 'IP address allocated to the pod. Routable at least within the cluster.' },
		conditions: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(EdgeClusterPodCondition))),
			description: 'Current service state of edge cluster pod',
		},
	},
});
