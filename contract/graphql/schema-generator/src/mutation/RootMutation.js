import { GraphQLObjectType } from 'graphql';
import createProject from './CreateProject';
import updateProject from './UpdateProject';
import deleteProject from './DeleteProject';
import createEdgeCluster from './CreateEdgeCluster';
import updateEdgeCluster from './UpdateEdgeCluster';
import deleteEdgeCluster from './DeleteEdgeCluster';

export default new GraphQLObjectType({
	name: 'Mutation',
	fields: {
		createProject,
		updateProject,
		deleteProject,
		createEdgeCluster,
		updateEdgeCluster,
		deleteEdgeCluster,
	},
});
