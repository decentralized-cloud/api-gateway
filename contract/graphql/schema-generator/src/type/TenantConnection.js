import { GraphQLInt } from 'graphql';
import { connectionDefinitions } from 'graphql-relay';
import Tenant from './Tenant';

export default connectionDefinitions({
	connectionFields: {
		totalCount: {
			type: GraphQLInt,
			description: 'Total number of tenants',
		},
	},
	name: 'TenantType',
	nodeType: Tenant,
});
