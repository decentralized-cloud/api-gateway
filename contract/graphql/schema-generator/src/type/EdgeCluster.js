import { GraphQLID, GraphQLObjectType, GraphQLString, GraphQLNonNull, GraphQLList } from 'graphql';
import { NodeInterface } from '../interface';
import Tenant from './EdgeClusterTenant';
import EdgeClusterProvisionDetail from './EdgeClusterProvisionDetail';
import EdgeClusterType from './EdgeClusterType';
import EdgeClusterNodeStatus from './EdgeClusterNodeStatus';

export default new GraphQLObjectType({
	name: 'EdgeCluster',
	description: 'The edge cluster',
	fields: {
		id: { type: new GraphQLNonNull(GraphQLID), description: 'The unique edge cluster ID' },
		name: { type: new GraphQLNonNull(GraphQLString), description: 'The edge cluster name' },
		clusterSecret: { type: new GraphQLNonNull(GraphQLString), description: 'The cluster secrect value' },
		clusterType: { type: new GraphQLNonNull(EdgeClusterType), description: 'The cluster type' },
		tenant: { type: new GraphQLNonNull(Tenant), description: 'The tenant that owns the edge cluster' },
		provisionDetail: { type: EdgeClusterProvisionDetail, description: 'The edge cluster provision details' },
		nodes: {
			type: new GraphQLList(new GraphQLNonNull(EdgeClusterNodeStatus)),
			description: 'The list of an existing edge cluster nodes details',
		},
	},
	interfaces: [NodeInterface],
});
