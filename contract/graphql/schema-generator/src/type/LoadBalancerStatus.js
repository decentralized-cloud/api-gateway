import { GraphQLObjectType, GraphQLNonNull, GraphQLList } from 'graphql';
import LoadBalancerIngress from './LoadBalancerIngress';

export default new GraphQLObjectType({
	name: 'LoadBalancerStatus',
	description: 'LoadBalancerStatus represents the status of a load-balancer',
	fields: {
		ingress: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(LoadBalancerIngress))),
			description:
				'Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points',
		},
	},
});
