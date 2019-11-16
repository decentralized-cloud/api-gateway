import { GraphQLString, GraphQLNonNull, GraphQLID } from 'graphql';
import { mutationWithClientMutationId } from 'graphql-relay';
import { EdgeClusterConnection } from '../type';

export default mutationWithClientMutationId({
  name: 'CreateEdgeCluster',
  inputFields: {
    tenantID: { type: new GraphQLNonNull(GraphQLID) },
    name: { type: new GraphQLNonNull(GraphQLString) },
    k3SClusterSecret: { type: new GraphQLNonNull(GraphQLString) },
  },
  outputFields: {
    edgeCluster: { type: EdgeClusterConnection.edgeType },
  },
  mutateAndGetPayload: () => ({}),
});
