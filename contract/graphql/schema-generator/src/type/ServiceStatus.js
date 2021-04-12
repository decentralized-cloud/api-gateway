import { GraphQLObjectType, GraphQLNonNull, GraphQLList } from 'graphql';
import LoadBalancerStatus from './LoadBalancerStatus';
import ServiceCondition from './ServiceCondition';

export default new GraphQLObjectType({
	name: 'ServiceStatus',
	description: 'Contains the most recently observed status of the existing edge cluster service',
	fields: {
		loadBalancer: {
			type: LoadBalancerStatus,
			description: 'LoadBalancer contains the current status of the load-balancer',
		},
		conditions: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(ServiceCondition))),
			description: 'Current service state of service',
		},
	},
});
