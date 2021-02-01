import { GraphQLObjectType, GraphQLNonNull, GraphQLString } from 'graphql';

export default new GraphQLObjectType({
	name: 'EdgeClusterNodeSystemInfo',
	description: 'contains a set of ids/uuids to uniquely identify the node',
	fields: {
		MachineID: {
			type: new GraphQLNonNull(GraphQLString),
			description: 'MachineID reported by the node. For unique machine identification in the cluster this field is preferred',
		},
		SystemUUID: {
			type: new GraphQLNonNull(GraphQLString),
			description:
				'SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat host',
		},
		BootID: { type: new GraphQLNonNull(GraphQLString), description: 'Boot ID reported by the node' },
		KernelVersion: {
			type: new GraphQLNonNull(GraphQLString),
			description: 'Kernel Version reported by the node from "uname -r" (e.g. 3.16.0-0.bpo.4-amd64).',
		},
		OSImage: {
			type: new GraphQLNonNull(GraphQLString),
			description: 'OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy))',
		},
		ContainerRuntimeVersion: {
			type: new GraphQLNonNull(GraphQLString),
			description: 'ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0)',
		},
		KubeletVersion: { type: new GraphQLNonNull(GraphQLString), description: 'Kubelet Version reported by the node' },
		KubeProxyVersion: { type: new GraphQLNonNull(GraphQLString), description: 'KubeProxy Version reported by the node' },
		OperatingSystem: { type: new GraphQLNonNull(GraphQLString), description: 'The Operating System reported by the node' },
		Architecture: { type: new GraphQLNonNull(GraphQLString), description: 'The Architecture reported by the node' },
	},
});
