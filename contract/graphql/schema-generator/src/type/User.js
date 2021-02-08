import { GraphQLID, GraphQLObjectType, GraphQLNonNull, GraphQLList } from 'graphql';
import { connectionArgs } from 'graphql-relay';
import { NodeInterface } from '../interface';
import Project from './Project';
import ProjectConnection from './ProjectConnection';
import EdgeCluster from './EdgeCluster';
import EdgeClusterConnection from './EdgeClusterConnection';
import SortingOptionPair from './SortingOptionPair';

export default new GraphQLObjectType({
	name: 'User',
	fields: {
		id: { type: new GraphQLNonNull(GraphQLID) },
		project: {
			type: Project,
			args: {
				projectID: { type: new GraphQLNonNull(GraphQLID) },
			},
		},
		projects: {
			type: ProjectConnection.connectionType,
			args: {
				...connectionArgs,
				projectIDs: { type: new GraphQLList(new GraphQLNonNull(GraphQLID)) },
				sortingOptions: { type: new GraphQLList(new GraphQLNonNull(SortingOptionPair)) },
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
				projectIDs: { type: new GraphQLList(new GraphQLNonNull(GraphQLID)) },
				sortingOptions: { type: new GraphQLList(new GraphQLNonNull(SortingOptionPair)) },
			},
		},
	},
	interfaces: [NodeInterface],
});
