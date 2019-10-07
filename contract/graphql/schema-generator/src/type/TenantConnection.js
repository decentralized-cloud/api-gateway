import { connectionDefinitions } from 'graphql-relay';
import Tenant from './Tenant';

export default connectionDefinitions({
  name: 'TenantType',
  nodeType: Tenant,
});
