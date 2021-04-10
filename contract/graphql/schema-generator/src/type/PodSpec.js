import { GraphQLObjectType, GraphQLNonNull, GraphQLString } from 'graphql';

export default new GraphQLObjectType({
	name: 'PodSpec',
	description: 'Contains the specification of the desired behavior of the existing edge cluster pod',
	fields: {
		nodeName: { type: new GraphQLNonNull(GraphQLString), description: 'The name of the node where the Pod is deployed into' },
	},
});
