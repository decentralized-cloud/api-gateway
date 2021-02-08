import { GraphQLInt } from 'graphql';
import { connectionDefinitions } from 'graphql-relay';
import Project from './Project';

export default connectionDefinitions({
	connectionFields: {
		totalCount: {
			type: GraphQLInt,
			description: 'Total number of projects',
		},
	},
	name: 'ProjectType',
	description: 'The project connection compatible with relay',
	nodeType: Project,
});
