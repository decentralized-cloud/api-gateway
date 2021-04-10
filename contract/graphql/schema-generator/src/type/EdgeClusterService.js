import { GraphQLObjectType, GraphQLNonNull } from 'graphql';
import ObjectMeta from './ObjectMeta';
import ServiceStatus from './ServiceStatus';
import ServiceSpec from './ServiceSpec';

export default new GraphQLObjectType({
	name: 'EdgeClusterService',
	description: 'Contains information about the edge cluster service',
	fields: {
		metadata: {
			type: new GraphQLNonNull(ObjectMeta),
			description: 'The service metadata',
		},
		status: {
			type: new GraphQLNonNull(ServiceStatus),
			description: 'The most recently observed status of the service',
		},
		spec: {
			type: new GraphQLNonNull(ServiceSpec),
			description: 'The specification of the desired behavior of the service',
		},
	},
});
