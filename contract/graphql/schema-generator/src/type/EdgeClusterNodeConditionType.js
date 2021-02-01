import { GraphQLEnumType } from 'graphql';

export default new GraphQLEnumType({
	name: 'EdgeClusterNodeConditionType',
	description: 'The valid conditions of node',
	values: {
		Ready: { value: 0, description: 'NodeReady means kubelet is healthy and ready to accept pods' },
		MemoryPressure: { value: 1, description: 'NodeMemoryPressure means the kubelet is under pressure due to insufficient available memory' },
		DiskPressure: { value: 2, description: 'NodeDiskPressure means the kubelet is under pressure due to insufficient available disk' },
		PIDPressure: { value: 3, description: 'NodePIDPressure means the kubelet is under pressure due to insufficient available PID' },
		NetworkUnavailable: { value: 4, description: 'NodeNetworkUnavailable means that network for the node is not correctly configured' },
	},
});
