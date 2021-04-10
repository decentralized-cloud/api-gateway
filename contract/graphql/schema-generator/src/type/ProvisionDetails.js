import { GraphQLObjectType, GraphQLString } from 'graphql';
import LoadBalancerStatus from './LoadBalancerStatus';

export default new GraphQLObjectType({
	name: 'ProvisionDetails',
	description:
		'The edge cluster provision details contains details such as current status of the edge cluster as well as ingress address of the edge cluster to connect to',
	fields: {
		loadBalancer: {
			type: LoadBalancerStatus,
			description: 'LoadBalancer contains the current status of the load-balancer',
		},
		kubeconfigContent: {
			type: GraphQLString,
			description: 'The provisioned edge cluster kubeconfig content',
		},
	},
});
