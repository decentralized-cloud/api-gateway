import { GraphQLString, GraphQLNonNull, GraphQLID } from 'graphql';
import { mutationWithClientMutationId } from 'graphql-relay';
import { EdgeClusterConnection } from '../type';

export default mutationWithClientMutationId({
  name: 'UpdateEdgeCluster',
  inputFields: {
    id: { type: new GraphQLNonNull(GraphQLID) },
    name: { type: new GraphQLNonNull(GraphQLString) },
  },
  outputFields: {
    edgeCluster: { type: EdgeClusterConnection.edgeType },
  },
  mutateAndGetPayload: () => ({}),
});
