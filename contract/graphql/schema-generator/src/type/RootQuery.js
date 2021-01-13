import { GraphQLObjectType } from 'graphql';
import UserType from './User';

export default new GraphQLObjectType({
	name: 'Query',
	fields: {
		user: { type: UserType },
	},
});
