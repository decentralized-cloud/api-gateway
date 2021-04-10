import { GraphQLObjectType, GraphQLNonNull, GraphQLList, GraphQLString } from 'graphql';
import PortStatus from './PortStatus';

export default new GraphQLObjectType({
	name: 'LoadBalancerIngress',
	description:
		'LoadBalancerIngress represents the status of a load-balancer ingress point traffic intended for the service should be sent to an ingress point',
	fields: {
		ip: { type: new GraphQLNonNull(GraphQLString), description: 'IP is set for load-balancer ingress points that are IP based' },
		hostname: { type: new GraphQLNonNull(GraphQLString), description: 'Hostname is set for load-balancer ingress points that are DNS based' },
		portStatus: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(PortStatus))),
			description: 'Ports is a list of records of service ports',
		},
	},
});
