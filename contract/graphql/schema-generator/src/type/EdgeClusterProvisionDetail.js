import { GraphQLObjectType, GraphQLString, GraphQLList, GraphQLNonNull } from 'graphql';
import Ingress from './Ingress';
import Port from './Port';

export default new GraphQLObjectType({
	name: 'EdgeClusterProvisionDetail',
	fields: {
		ingress: { type: new GraphQLList(new GraphQLNonNull(Ingress)) },
		ports: { type: new GraphQLList(new GraphQLNonNull(Port)) },
		kubeconfigContent: { type: GraphQLString },
	},
});
