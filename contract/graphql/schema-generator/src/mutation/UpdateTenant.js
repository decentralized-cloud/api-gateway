import { GraphQLString, GraphQLNonNull, GraphQLID } from 'graphql';
import { mutationWithClientMutationId } from 'graphql-relay';
import { TenantConnection } from '../type';

export default mutationWithClientMutationId({
	name: 'UpdateTenant',
	inputFields: {
		tenantID: { type: new GraphQLNonNull(GraphQLID) },
		name: { type: new GraphQLNonNull(GraphQLString) },
	},
	outputFields: {
		tenant: { type: TenantConnection.edgeType },
	},
	mutateAndGetPayload: () => ({}),
});
