import { GraphQLEnumType } from 'graphql';

export default new GraphQLEnumType({
	name: 'Protocol',
	values: {
		TCP: { value: 0 },
		UDP: { value: 1 },
		SCTP: { value: 2 },
	},
});
