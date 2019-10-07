import { GraphQLString, GraphQLNonNull } from 'graphql';
import { mutationWithClientMutationId } from 'graphql-relay';
import { EdgeClusterConnection } from '../type';

export default mutationWithClientMutationId({
  name: 'CreateEdgeCluster',
  inputFields: {
    name: { type: new GraphQLNonNull(GraphQLString) },
  },
  outputFields: {
    edgeCluster: { type: EdgeClusterConnection.edgeType },
  },
  mutateAndGetPayload: () => ({}),
});
