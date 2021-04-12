import { GraphQLObjectType, GraphQLInt, GraphQLNonNull, GraphQLString } from 'graphql';
import Protocol from './Protocol';

export default new GraphQLObjectType({
	name: 'ServicePort',
	description: 'ServicePort contains information on service port',
	fields: {
		name: { type: new GraphQLNonNull(GraphQLString), description: 'The name of this port within the service' },
		protocol: {
			type: new GraphQLNonNull(Protocol),
			description: 'The IP protocol for this port',
		},
		port: { type: new GraphQLNonNull(GraphQLInt), description: 'The port that will be exposed by this service' },
		targetPort: {
			type: new GraphQLNonNull(GraphQLString),
			description: 'Number or name of the port to access on the pods targeted by the service',
		},
		nodePort: {
			type: new GraphQLNonNull(GraphQLInt),
			description: 'The port on each node on which this service is exposed when type is NodePort or LoadBalancer',
		},
	},
});
