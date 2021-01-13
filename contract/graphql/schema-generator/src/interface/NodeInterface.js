import { nodeDefinitions } from 'graphql-relay';

const { nodeInterface, nodeField } = nodeDefinitions(
	() => null,
	() => null
);

export const NodeInterface = nodeInterface;
export const NodeField = nodeField;
