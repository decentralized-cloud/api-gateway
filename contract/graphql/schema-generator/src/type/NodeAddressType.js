import { GraphQLEnumType } from 'graphql';

export default new GraphQLEnumType({
	name: 'NodeAddressType',
	description: 'The valid address type of edge cluster node',
	values: {
		Hostname: { value: 0 },
		ExternalIP: { value: 1 },
		InternalIP: { value: 2 },
		ExternalDNS: { value: 3 },
		InternalDNS: { value: 4 },
	},
});
