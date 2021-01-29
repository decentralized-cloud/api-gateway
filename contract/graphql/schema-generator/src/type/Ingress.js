import { GraphQLObjectType, GraphQLString } from 'graphql';

export default new GraphQLObjectType({
	name: 'Ingress',
	fields: {
		ip: { type: GraphQLString },
		hostname: { type: GraphQLString },
	},
});
