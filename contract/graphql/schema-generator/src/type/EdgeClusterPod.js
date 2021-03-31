import { GraphQLObjectType, GraphQLNonNull } from 'graphql';
import EdgeClusterObjectMetadata from './EdgeClusterObjectMetadata';
import EdgeClusterPodStatus from './EdgeClusterPodStatus';
import EdgeClusterPodSpec from './EdgeClusterPodSpec';

export default new GraphQLObjectType({
	name: 'EdgeClusterPod',
	description: 'Contains information about the edge cluster pod',
	fields: {
		metadata: {
			type: new GraphQLNonNull(EdgeClusterObjectMetadata),
			description: 'The pod metadata',
		},
		status: {
			type: new GraphQLNonNull(EdgeClusterPodStatus),
			description: 'The most recently observed status of the pod',
		},
		spec: {
			type: new GraphQLNonNull(EdgeClusterPodSpec),
			description: 'The specification of the desired behavior of the pod.',
		},
	},
});
