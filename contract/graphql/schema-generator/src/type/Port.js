import { GraphQLObjectType, GraphQLInt, GraphQLNonNull } from 'graphql';
import Protocol from './Protocol';

export default new GraphQLObjectType({
	name: 'Port',
	fields: {
		port: { type: new GraphQLNonNull(GraphQLInt) },
		protocol: { type: new GraphQLNonNull(Protocol) },
	},
});
