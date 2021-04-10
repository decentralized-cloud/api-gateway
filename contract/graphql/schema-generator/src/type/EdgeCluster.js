import { GraphQLID, GraphQLObjectType, GraphQLString, GraphQLNonNull, GraphQLList } from 'graphql';
import { NodeInterface } from '../interface';
import Project from './EdgeClusterProject';
import ProvisionDetails from './ProvisionDetails';
import EdgeClusterType from './EdgeClusterType';
import EdgeClusterNode from './EdgeClusterNode';
import EdgeClusterPod from './EdgeClusterPod';
import EdgeClusterService from './EdgeClusterService';

export default new GraphQLObjectType({
	name: 'EdgeCluster',
	description: 'The edge cluster',
	fields: {
		id: { type: new GraphQLNonNull(GraphQLID), description: 'The unique edge cluster ID' },
		name: { type: new GraphQLNonNull(GraphQLString), description: 'The edge cluster name' },
		clusterSecret: { type: new GraphQLNonNull(GraphQLString), description: 'The cluster secrect value' },
		clusterType: { type: new GraphQLNonNull(EdgeClusterType), description: 'The cluster type' },
		project: { type: new GraphQLNonNull(Project), description: 'The project that owns the edge cluster' },
		provisionDetails: { type: new GraphQLNonNull(ProvisionDetails), description: 'The edge cluster provision details' },
		nodes: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(EdgeClusterNode))),
			description: 'The list of edge cluster nodes details',
		},
		pods: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(EdgeClusterPod))),
			description: 'The list of edge cluster pods details',
			args: {
				nodeName: { type: GraphQLString },
				namespace: { type: GraphQLString },
			},
		},
		services: {
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(EdgeClusterService))),
			description: 'The list of edge cluster services details',
			args: {
				namespace: { type: GraphQLString },
			},
		},
	},
	interfaces: [NodeInterface],
});
