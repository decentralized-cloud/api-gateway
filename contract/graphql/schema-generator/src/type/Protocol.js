import { GraphQLEnumType } from 'graphql';

export default new GraphQLEnumType({
	name: 'Protocol',
	description: 'Protocol defines network protocols',
	values: {
		TCP: { value: 0, description: 'TCP protocol' },
		UDP: { value: 1, description: 'UDP protocol' },
		SCTP: { value: 2, description: 'SCTP protocol' },
	},
});
