import {
  GraphQLID, GraphQLObjectType, GraphQLString, GraphQLNonNull,
} from 'graphql';
import { NodeInterface } from '../interface';
import Tenant from './EdgeClusterTenant';
import EdgeClusterProvisioningDetail from './EdgeClusterProvisioningDetail';

export default new GraphQLObjectType({
  name: 'EdgeCluster',
  fields: {
    id: { type: new GraphQLNonNull(GraphQLID) },
    name: { type: new GraphQLNonNull(GraphQLString) },
    clusterSecret: { type: new GraphQLNonNull(GraphQLString) },
    tenant: { type: new GraphQLNonNull(Tenant) },
    provisioningDetail: { type: EdgeClusterProvisioningDetail },
  },
  interfaces: [NodeInterface],
});
