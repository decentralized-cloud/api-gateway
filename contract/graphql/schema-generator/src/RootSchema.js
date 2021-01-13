import { GraphQLSchema } from 'graphql';
import { RootMutation } from './mutation';
import { RootQuery } from './type';

export default function getRootSchema() {
	return new GraphQLSchema({
		query: RootQuery,
		mutation: RootMutation,
	});
}
