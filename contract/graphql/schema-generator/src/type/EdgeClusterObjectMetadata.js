import { GraphQLID, GraphQLObjectType, GraphQLNonNull, GraphQLString } from 'graphql';

export default new GraphQLObjectType({
	name: 'EdgeClusterObjectMetadata',
	description: 'Contains standard edge cluster objects metadata.',
	fields: {
		id: {
			type: new GraphQLNonNull(GraphQLID),
			description: 'The object unique identitfier',
		},
		name: {
			type: new GraphQLNonNull(GraphQLString),
			description: 'The object name',
		},
		namespace: {
			type: GraphQLString,
			description: 'The object namespace',
		},
	},
});
