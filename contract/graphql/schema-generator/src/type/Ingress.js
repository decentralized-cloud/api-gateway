import { GraphQLObjectType, GraphQLString } from 'graphql';

export default new GraphQLObjectType({
	name: 'Ingress',
	description: 'Ingress represents the status of a load-balancer ingress point traffic intended for the service should be sent to an ingress point',
	fields: {
		ip: {
			type: GraphQLString,
			description: 'IP is set for load-balancer ingress points that are IP based, (typically GCE or OpenStack load-balancers)',
		},
		hostname: {
			type: GraphQLString,
			description: 'Hostname is set for load-balancer ingress points that are DNS based, (typically AWS load-balancers)',
		},
	},
});
