import { api } from './api';

const rbacModel = `[request_definition]
    r = sub, res, act, obj
    
    [policy_definition]
    p = sub, res, act, obj
    
    [role_definition]
    g = _, _
    
    [policy_effect]
    e = some(where (p.eft == allow)) && !some(where (p.eft == deny))
    
    [matchers]
    m = g(r.sub, p.sub) && globMatch(r.res, p.res) && globMatch(r.act, p.act) && keyMatch(r.obj, p.obj)`;

export const getRBACPolicies = async () => {
  const response = await api.get('policies');

  return {
    m: rbacModel,
    p: response.data,
  };
};
