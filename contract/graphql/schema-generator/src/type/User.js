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
        tenantID: { type: new GraphQLNonNull(GraphQLID) },
      },
    },
    tenants: {
      type: TenantConnection.connectionType,
      args: {
        ...connectionArgs,
        tenantIDs: { type: new GraphQLList(new GraphQLNonNull(GraphQLID)) },
        sortOption: { type: GraphQLString },
      },
    },
    edgeCluster: {
      type: EdgeCluster,
      args: {
        edgeClusterID: { type: new GraphQLNonNull(GraphQLID) },
      },
    },
    edgeClusters: {
      type: EdgeClusterConnection.connectionType,
      args: {
        ...connectionArgs,
        edgeClusterIDs: { type: new GraphQLList(new GraphQLNonNull(GraphQLID)) },
        sortOption: { type: GraphQLString },
      },
    },
  },
  interfaces: [NodeInterface],
});
