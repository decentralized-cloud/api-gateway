import { GraphQLObjectType } from 'graphql';
import createTenant from './CreateTenant';
import updateTenant from './UpdateTenant';
import deleteTenant from './DeleteTenant';
import createEdgeCluster from './CreateEdgeCluster';
import updateEdgeCluster from './UpdateEdgeCluster';
import deleteEdgeCluster from './DeleteEdgeCluster';

export default new GraphQLObjectType({
	name: 'Mutation',
	fields: {
		createTenant,
		updateTenant,
		deleteTenant,
		createEdgeCluster,
		updateEdgeCluster,
		deleteEdgeCluster,
	},
});
