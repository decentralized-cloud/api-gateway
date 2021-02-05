import { GraphQLString, GraphQLNonNull, GraphQLID } from 'graphql';
import { mutationWithClientMutationId } from 'graphql-relay';
import { EdgeClusterConnection, EdgeClusterType } from '../type';

export default mutationWithClientMutationId({
	name: 'UpdateEdgeCluster',
	inputFields: {
		id: { type: new GraphQLNonNull(GraphQLID) },
		tenantID: { type: new GraphQLNonNull(GraphQLID) },
		name: { type: new GraphQLNonNull(GraphQLString) },
		clusterSecret: { type: new GraphQLNonNull(GraphQLString) },
		clusterType: { type: new GraphQLNonNull(EdgeClusterType) },
	},
	outputFields: {
		edgeCluster: { type: EdgeClusterConnection.edgeType },
	},
	mutateAndGetPayload: () => ({}),
});
