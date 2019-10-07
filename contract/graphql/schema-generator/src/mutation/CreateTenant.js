import { GraphQLString, GraphQLNonNull } from 'graphql';
import { mutationWithClientMutationId } from 'graphql-relay';
import { TenantConnection } from '../type';

export default mutationWithClientMutationId({
  name: 'CreateTenant',
  inputFields: {
    name: { type: new GraphQLNonNull(GraphQLString) },
  },
  outputFields: {
    tenant: { type: TenantConnection.edgeType },
  },
  mutateAndGetPayload: () => ({}),
});
