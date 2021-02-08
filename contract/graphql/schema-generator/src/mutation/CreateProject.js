import { GraphQLString, GraphQLNonNull } from 'graphql';
import { mutationWithClientMutationId } from 'graphql-relay';
import { ProjectConnection } from '../type';

export default mutationWithClientMutationId({
	name: 'CreateProject',
	inputFields: {
		name: { type: new GraphQLNonNull(GraphQLString) },
	},
	outputFields: {
		project: { type: ProjectConnection.edgeType },
	},
	mutateAndGetPayload: () => ({}),
});
