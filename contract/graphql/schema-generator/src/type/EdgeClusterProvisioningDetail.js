import { GraphQLObjectType, GraphQLString } from 'graphql';
import EdgeClusterStatus from './EdgeClusterStatus';

export default new GraphQLObjectType({
  name: 'EdgeClusterProvisioningDetail',
  fields: {
    status: { type: EdgeClusterStatus },
    publicIPAddress: { type: GraphQLString },
    kubeconfigContent: { type: GraphQLString },
  },
});
