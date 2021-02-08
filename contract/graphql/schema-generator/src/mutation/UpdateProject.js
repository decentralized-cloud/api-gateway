import { GraphQLString, GraphQLNonNull, GraphQLID } from 'graphql';
import { mutationWithClientMutationId } from 'graphql-relay';
import { ProjectConnection } from '../type';

export default mutationWithClientMutationId({
	name: 'UpdateProject',
	inputFields: {
		projectID: { type: new GraphQLNonNull(GraphQLID) },
		name: { type: new GraphQLNonNull(GraphQLString) },
	},
	outputFields: {
		project: { type: ProjectConnection.edgeType },
	},
	mutateAndGetPayload: () => ({}),
});
