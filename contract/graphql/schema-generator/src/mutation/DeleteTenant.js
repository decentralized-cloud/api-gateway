import { GraphQLNonNull, GraphQLID } from 'graphql';
import { mutationWithClientMutationId } from 'graphql-relay';

export default mutationWithClientMutationId({
  name: 'DeleteTenant',
  inputFields: {
    tenantID: { type: new GraphQLNonNull(GraphQLID) },
  },
  outputFields: {
    deletedTenantID: { type: new GraphQLNonNull(GraphQLID) },
  },
  mutateAndGetPayload: () => ({}),
});
