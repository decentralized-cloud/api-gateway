import { connectionDefinitions } from 'graphql-relay';
import EdgeCluster from './EdgeCluster';

export default connectionDefinitions({
  name: 'EdgeClusterType',
  nodeType: EdgeCluster,
});
