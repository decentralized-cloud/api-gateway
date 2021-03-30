import { GraphQLID, GraphQLObjectType, GraphQLString, GraphQLNonNull, GraphQLList } from 'graphql';
import { NodeInterface } from '../interface';
import Project from './EdgeClusterProject';
import EdgeClusterProvisionDetail from './EdgeClusterProvisionDetail';
import EdgeClusterType from './EdgeClusterType';
import EdgeClusterNode from './EdgeClusterNode';

export default new GraphQLObjectType({
	name: 'EdgeCluster',
	description: 'The edge cluster',
	fields: {
		id: { type: new GraphQLNonNull(GraphQLID), description: 'The unique edge cluster ID' },
		name: { type: new GraphQLNonNull(GraphQLString), description: 'The edge cluster name' },
		clusterSecret: { type: new GraphQLNonNull(GraphQLString), description: 'The cluster secrect value' },
		clusterType: { type: new GraphQLNonNull(EdgeClusterType), description: 'The cluster type' },
		project: { type: new GraphQLNonNull(Project), description: 'The project that owns the edge cluster' },
		provisionDetail: { type: EdgeClusterProvisionDetail, description: 'The edge cluster provision details' },
		nodes: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(EdgeClusterNode))),
			description: 'The list of edge cluster nodes details',
		},
	},
	interfaces: [NodeInterface],
});
