import {
  GraphQLID, GraphQLObjectType, GraphQLString, GraphQLNonNull,
} from 'graphql';
import { NodeInterface } from '../interface';

export default new GraphQLObjectType({
  name: 'EdgeCluster',
  fields: {
    id: { type: new GraphQLNonNull(GraphQLID) },
    name: { type: new GraphQLNonNull(GraphQLString) },
  },
  interfaces: [NodeInterface],
});
