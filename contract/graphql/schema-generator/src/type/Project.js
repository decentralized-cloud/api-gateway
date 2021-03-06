import { GraphQLID, GraphQLObjectType, GraphQLString, GraphQLNonNull, GraphQLList } from 'graphql';
import { connectionArgs } from 'graphql-relay';
import { NodeInterface } from '../interface';
import EdgeCluster from './EdgeCluster';
import EdgeClusterConnection from './EdgeClusterConnection';

export default new GraphQLObjectType({
	name: 'Project',
	fields: {
		id: { type: new GraphQLNonNull(GraphQLID) },
		name: { type: new GraphQLNonNull(GraphQLString) },
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
