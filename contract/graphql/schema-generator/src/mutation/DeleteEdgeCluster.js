import { GraphQLNonNull, GraphQLID } from 'graphql';
import { mutationWithClientMutationId } from 'graphql-relay';

export default mutationWithClientMutationId({
	name: 'DeleteEdgeCluster',
	inputFields: {
		id: { type: new GraphQLNonNull(GraphQLID) },
	},
	outputFields: {
		deletedEdgeClusterID: { type: new GraphQLNonNull(GraphQLID) },
	},
	mutateAndGetPayload: () => ({}),
});
