import { GraphQLObjectType, GraphQLNonNull } from 'graphql';
import ObjectMeta from './ObjectMeta';
import PodStatus from './PodStatus';
import PodSpec from './PodSpec';

export default new GraphQLObjectType({
	name: 'EdgeClusterPod',
	description: 'Contains information about the edge cluster pod',
	fields: {
		metadata: {
			type: new GraphQLNonNull(ObjectMeta),
			description: 'The pod metadata',
		},
		status: {
			type: new GraphQLNonNull(PodStatus),
			description: 'The most recently observed status of the pod',
		},
		spec: {
			type: new GraphQLNonNull(PodSpec),
			description: 'The specification of the desired behavior of the pod',
		},
	},
});
