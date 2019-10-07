import {
  GraphQLID, GraphQLObjectType, GraphQLString, GraphQLNonNull, GraphQLList,
} from 'graphql';
import { connectionArgs } from 'graphql-relay';
import { NodeInterface } from '../interface';
import Tenant from './Tenant';
import TenantConnection from './TenantConnection';
import EdgeCluster from './EdgeCluster';
import EdgeClusterConnection from './EdgeClusterConnection';

export default new GraphQLObjectType({
  name: 'User',
  fields: {
    id: { type: new GraphQLNonNull(GraphQLID) },
    tenant: {
      type: Tenant,
      args: {
        tenantId: { type: new GraphQLNonNull(GraphQLID) },
      },
    },
    tenants: {
      type: TenantConnection.connectionType,
      args: {
        ...connectionArgs,
        tenantIds: { type: new GraphQLList(new GraphQLNonNull(GraphQLID)) },
        sortOption: { type: GraphQLString },
      },
    },
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
