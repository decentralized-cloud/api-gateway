import { GraphQLObjectType, GraphQLInt, GraphQLNonNull, GraphQLString } from 'graphql';
import Protocol from './Protocol';

export default new GraphQLObjectType({
	name: 'PortStatus',
	description: 'PortStatus represents the error condition of a service port',
	fields: {
		port: { type: new GraphQLNonNull(GraphQLInt), description: 'Port is the port number of the service port of which status is recorded here' },
		protocol: {
			type: new GraphQLNonNull(Protocol),
			description: 'Protocol is the protocol of the service port of which status is recorded here',
		},
		error: { type: GraphQLString, description: 'Error is to record the problem with the service port' },
	},
});
