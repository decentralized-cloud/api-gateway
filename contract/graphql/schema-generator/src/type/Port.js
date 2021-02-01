import { GraphQLObjectType, GraphQLInt, GraphQLNonNull } from 'graphql';
import Protocol from './Protocol';

export default new GraphQLObjectType({
	name: 'Port',
	description: 'Port contains information on service port',
	fields: {
		port: { type: new GraphQLNonNull(GraphQLInt), description: 'The port number of the edge-cluster master port' },
		protocol: { type: new GraphQLNonNull(Protocol), description: 'The protocol of the edge-cluster master port' },
	},
});
