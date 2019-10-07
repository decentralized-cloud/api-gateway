import { GraphQLObjectType } from 'graphql';
import UserType from './User';
import { NodeField } from '../interface';

export default new GraphQLObjectType({
  name: 'Query',
  fields: {
    user: { type: UserType },
    node: NodeField,
  },
});
