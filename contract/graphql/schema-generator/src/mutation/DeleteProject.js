import { GraphQLNonNull, GraphQLID } from 'graphql';
import { mutationWithClientMutationId } from 'graphql-relay';

export default mutationWithClientMutationId({
	name: 'DeleteProject',
	inputFields: {
		projectID: { type: new GraphQLNonNull(GraphQLID) },
	},
	outputFields: {
		deletedProjectID: { type: new GraphQLNonNull(GraphQLID) },
	},
	mutateAndGetPayload: () => ({}),
});
