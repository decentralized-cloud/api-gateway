import { GraphQLEnumType } from 'graphql';

export default new GraphQLEnumType({
	name: 'ServiceType',
	description: 'ServiceType string describes ingress methods for a service',
	values: {
		ClusterIP: { value: 0, description: 'ClusterIP means a service will only be accessible inside the cluster, via the cluster IP' },
		NodePort: { value: 1, description: 'NodePort means a service will be exposed on one port of every node, in addition to ClusterIP type' },
		LoadBalancer: {
			value: 3,
			description:
				'LoadBalancer means a service will be exposed via an external load balancer (if the cloud provider supports it), in addition to NodePort type',
		},
		ExternalName: {
			value: 4,
			description:
				'ExternalName means a service consists of only a reference to an external name that kubedns or equivalent will return as a CNAME record, with no exposing or proxying of any pods involved',
		},
	},
});
