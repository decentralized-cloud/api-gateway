import { GraphQLObjectType, GraphQLNonNull, GraphQLString, GraphQLList } from 'graphql';
import ServiceType from './ServiceType';
import ServicePort from './ServicePort';

export default new GraphQLObjectType({
	name: 'ServiceSpec',
	description: 'Contains the specification of the desired behavior of the existing edge cluster service',
	fields: {
		ports: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(ServicePort))),
			description: 'The list of ports that are exposed by this service',
		},
		clusterIPs: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(GraphQLString))),
			description: 'clusterIPs is a list of IP addresses assigned to this service',
		},
		type: {
			type: new GraphQLNonNull(ServiceType),
			description: 'type determines how the Service is exposed',
		},
		externalIPs: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(GraphQLString))),
			description: 'externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service',
		},
		externalName: {
			type: GraphQLString,
			description:
				'externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record)',
		},
	},
});
