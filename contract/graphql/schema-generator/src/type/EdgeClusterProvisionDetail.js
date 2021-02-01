import { GraphQLObjectType, GraphQLString, GraphQLList, GraphQLNonNull } from 'graphql';
import Ingress from './Ingress';
import Port from './Port';

export default new GraphQLObjectType({
	name: 'EdgeClusterProvisionDetail',
	description:
		'The edge cluster provision details contains details such as current status of the edge cluster as well as ingress address of the edge cluster to connect to',
	fields: {
		ingress: { type: new GraphQLList(new GraphQLNonNull(Ingress)), description: 'The ingress details of the edge cluster master node' },
		ports: { type: new GraphQLList(new GraphQLNonNull(Port)), description: 'The ingress details of the edge cluster master node' },
		kubeconfigContent: { type: GraphQLString, description: 'The provisioned edge cluster kubeconfig content' },
	},
});
