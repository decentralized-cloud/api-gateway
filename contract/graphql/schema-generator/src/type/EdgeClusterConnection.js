import { GraphQLInt } from 'graphql';
import { connectionDefinitions } from 'graphql-relay';
import EdgeCluster from './EdgeCluster';

export default connectionDefinitions({
	connectionFields: {
		totalCount: {
			type: GraphQLInt,
			description: 'Total number of edge clusters',
		},
	},
	name: 'EdgeClusterType',
	nodeType: EdgeCluster,
});
