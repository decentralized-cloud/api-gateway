import {
  GraphQLID, GraphQLObjectType, GraphQLString, GraphQLNonNull, GraphQLList,
} from 'graphql';
import { connectionArgs } from 'graphql-relay';
import { NodeInterface } from '../interface';
import EdgeCluster from './EdgeCluster';
import EdgeClusterConnection from './EdgeClusterConnection';

export default new GraphQLObjectType({
  name: 'Tenant',
  fields: {
    id: { type: new GraphQLNonNull(GraphQLID) },
    name: { type: GraphQLString },
    edgeCluster: {
      type: EdgeCluster,
      args: {
        edgeClusterId: { type: new GraphQLNonNull(GraphQLID) },
      },
    },
    edgeClusters: {
      type: EdgeClusterConnection.connectionType,
      args: {
        ...connectionArgs,
        edgeClusterIds: { type: new GraphQLList(new GraphQLNonNull(GraphQLID)) },
        sortOption: { type: GraphQLString },
      },
    },
  },
  interfaces: [NodeInterface],
});
