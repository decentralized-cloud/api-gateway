import { GraphQLEnumType } from 'graphql';

export default new GraphQLEnumType({
  name: 'EdgeClusterType',
  values: {
    K3S: { value: 0 },
  },
});
