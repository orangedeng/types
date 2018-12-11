# Rancher API Document
---

## API Specification

> API-Spec Document is From `https://github.com/rancher/api-spec/blob/master/specification.md`

## API List: ##

## Object Schema Name: `NodePool`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `NodePool` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/nodePools`|
|Input| `null` |
|Output| List of [resource object `NodePool`](#object-NodePool) |
### Get `NodePool`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/nodePools/<resource_id>` |
|Input| `null`|
|Output| [Resource object `NodePool`](#object-NodePool) |
### Create `NodePool`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/nodePools` |
|Input| [Resource object `NodePool`](#object-NodePool) |
|Output| [Resource object `NodePool`](#object-NodePool) after create |
### Update `NodePool`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/nodePools/<resource_id>` |
|Input| [Resource object `NodePool`](#object-NodePool) to update |
|Output| [Resource object `NodePool`](#object-NodePool) after delete |
### Delete `NodePool`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/nodePools/<resource_id>` |
|Input| `null` |
|Output| [Resource object `NodePool`](#object-NodePool) after delete |
---

## Object Schema Name: `Node`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `Node` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/nodes`|
|Input| `null` |
|Output| List of [resource object `Node`](#object-Node) |
### Get `Node`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/nodes/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Node`](#object-Node) |
### Create `Node`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/nodes` |
|Input| [Resource object `Node`](#object-Node) |
|Output| [Resource object `Node`](#object-Node) after create |
### Update `Node`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/nodes/<resource_id>` |
|Input| [Resource object `Node`](#object-Node) to update |
|Output| [Resource object `Node`](#object-Node) after delete |
### Delete `Node`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/nodes/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Node`](#object-Node) after delete |
### Resource Action `cordon`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/nodes/<resource_id>/cordon` |
|Input| `null` |
|Output| `null` |
### Resource Action `drain`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/nodes/<resource_id>/drain` |
|Input| [object `nodeDrainInput`](#object-nodeDrainInput) |
|Output| `null` |
### Resource Action `stopDrain`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/nodes/<resource_id>/stopDrain` |
|Input| `null` |
|Output| `null` |
### Resource Action `uncordon`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/nodes/<resource_id>/uncordon` |
|Input| `null` |
|Output| `null` |
---

## Object Schema Name: `NodeDriver`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `NodeDriver` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/nodeDrivers`|
|Input| `null` |
|Output| List of [resource object `NodeDriver`](#object-NodeDriver) |
### Get `NodeDriver`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/nodeDrivers/<resource_id>` |
|Input| `null`|
|Output| [Resource object `NodeDriver`](#object-NodeDriver) |
### Create `NodeDriver`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/nodeDrivers` |
|Input| [Resource object `NodeDriver`](#object-NodeDriver) |
|Output| [Resource object `NodeDriver`](#object-NodeDriver) after create |
### Update `NodeDriver`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/nodeDrivers/<resource_id>` |
|Input| [Resource object `NodeDriver`](#object-NodeDriver) to update |
|Output| [Resource object `NodeDriver`](#object-NodeDriver) after delete |
### Delete `NodeDriver`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/nodeDrivers/<resource_id>` |
|Input| `null` |
|Output| [Resource object `NodeDriver`](#object-NodeDriver) after delete |
### Resource Action `activate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/nodeDrivers/<resource_id>/activate` |
|Input| `null` |
|Output| [object `nodeDriver`](#object-nodeDriver) |
### Resource Action `deactivate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/nodeDrivers/<resource_id>/deactivate` |
|Input| `null` |
|Output| [object `nodeDriver`](#object-nodeDriver) |
---

## Object Schema Name: `NodeTemplate`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `NodeTemplate` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/nodeTemplates`|
|Input| `null` |
|Output| List of [resource object `NodeTemplate`](#object-NodeTemplate) |
### Get `NodeTemplate`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/nodeTemplates/<resource_id>` |
|Input| `null`|
|Output| [Resource object `NodeTemplate`](#object-NodeTemplate) |
### Create `NodeTemplate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/nodeTemplates` |
|Input| [Resource object `NodeTemplate`](#object-NodeTemplate) |
|Output| [Resource object `NodeTemplate`](#object-NodeTemplate) after create |
### Update `NodeTemplate`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/nodeTemplates/<resource_id>` |
|Input| [Resource object `NodeTemplate`](#object-NodeTemplate) to update |
|Output| [Resource object `NodeTemplate`](#object-NodeTemplate) after delete |
### Delete `NodeTemplate`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/nodeTemplates/<resource_id>` |
|Input| `null` |
|Output| [Resource object `NodeTemplate`](#object-NodeTemplate) after delete |
---

## Object Schema Name: `Project`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `Project` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects`|
|Input| `null` |
|Output| List of [resource object `Project`](#object-Project) |
### Get `Project`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Project`](#object-Project) |
### Create `Project`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects` |
|Input| [Resource object `Project`](#object-Project) |
|Output| [Resource object `Project`](#object-Project) after create |
### Update `Project`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<resource_id>` |
|Input| [Resource object `Project`](#object-Project) to update |
|Output| [Resource object `Project`](#object-Project) after delete |
### Delete `Project`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Project`](#object-Project) after delete |
### Resource Action `disableMonitoring`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<resource_id>/disableMonitoring` |
|Input| `null` |
|Output| `null` |
### Resource Action `enableMonitoring`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<resource_id>/enableMonitoring` |
|Input| [object `monitoringInput`](#object-monitoringInput) |
|Output| `null` |
### Resource Action `exportYaml`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<resource_id>/exportYaml` |
|Input| `null` |
|Output| `null` |
### Resource Action `setpodsecuritypolicytemplate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<resource_id>/setpodsecuritypolicytemplate` |
|Input| [object `setPodSecurityPolicyTemplateInput`](#object-setPodSecurityPolicyTemplateInput) |
|Output| [object `project`](#object-project) |
---

## Object Schema Name: `GlobalRole`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `GlobalRole` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/globalRoles`|
|Input| `null` |
|Output| List of [resource object `GlobalRole`](#object-GlobalRole) |
### Get `GlobalRole`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/globalRoles/<resource_id>` |
|Input| `null`|
|Output| [Resource object `GlobalRole`](#object-GlobalRole) |
### Update `GlobalRole`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/globalRoles/<resource_id>` |
|Input| [Resource object `GlobalRole`](#object-GlobalRole) to update |
|Output| [Resource object `GlobalRole`](#object-GlobalRole) after delete |
---

## Object Schema Name: `GlobalRoleBinding`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `GlobalRoleBinding` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/globalRoleBindings`|
|Input| `null` |
|Output| List of [resource object `GlobalRoleBinding`](#object-GlobalRoleBinding) |
### Get `GlobalRoleBinding`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/globalRoleBindings/<resource_id>` |
|Input| `null`|
|Output| [Resource object `GlobalRoleBinding`](#object-GlobalRoleBinding) |
### Create `GlobalRoleBinding`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/globalRoleBindings` |
|Input| [Resource object `GlobalRoleBinding`](#object-GlobalRoleBinding) |
|Output| [Resource object `GlobalRoleBinding`](#object-GlobalRoleBinding) after create |
### Update `GlobalRoleBinding`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/globalRoleBindings/<resource_id>` |
|Input| [Resource object `GlobalRoleBinding`](#object-GlobalRoleBinding) to update |
|Output| [Resource object `GlobalRoleBinding`](#object-GlobalRoleBinding) after delete |
### Delete `GlobalRoleBinding`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/globalRoleBindings/<resource_id>` |
|Input| `null` |
|Output| [Resource object `GlobalRoleBinding`](#object-GlobalRoleBinding) after delete |
---

## Object Schema Name: `RoleTemplate`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `RoleTemplate` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/roleTemplates`|
|Input| `null` |
|Output| List of [resource object `RoleTemplate`](#object-RoleTemplate) |
### Get `RoleTemplate`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/roleTemplates/<resource_id>` |
|Input| `null`|
|Output| [Resource object `RoleTemplate`](#object-RoleTemplate) |
### Create `RoleTemplate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/roleTemplates` |
|Input| [Resource object `RoleTemplate`](#object-RoleTemplate) |
|Output| [Resource object `RoleTemplate`](#object-RoleTemplate) after create |
### Update `RoleTemplate`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/roleTemplates/<resource_id>` |
|Input| [Resource object `RoleTemplate`](#object-RoleTemplate) to update |
|Output| [Resource object `RoleTemplate`](#object-RoleTemplate) after delete |
### Delete `RoleTemplate`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/roleTemplates/<resource_id>` |
|Input| `null` |
|Output| [Resource object `RoleTemplate`](#object-RoleTemplate) after delete |
---

## Object Schema Name: `PodSecurityPolicyTemplate`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `PodSecurityPolicyTemplate` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/podSecurityPolicyTemplates`|
|Input| `null` |
|Output| List of [resource object `PodSecurityPolicyTemplate`](#object-PodSecurityPolicyTemplate) |
### Get `PodSecurityPolicyTemplate`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/podSecurityPolicyTemplates/<resource_id>` |
|Input| `null`|
|Output| [Resource object `PodSecurityPolicyTemplate`](#object-PodSecurityPolicyTemplate) |
### Create `PodSecurityPolicyTemplate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/podSecurityPolicyTemplates` |
|Input| [Resource object `PodSecurityPolicyTemplate`](#object-PodSecurityPolicyTemplate) |
|Output| [Resource object `PodSecurityPolicyTemplate`](#object-PodSecurityPolicyTemplate) after create |
### Update `PodSecurityPolicyTemplate`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/podSecurityPolicyTemplates/<resource_id>` |
|Input| [Resource object `PodSecurityPolicyTemplate`](#object-PodSecurityPolicyTemplate) to update |
|Output| [Resource object `PodSecurityPolicyTemplate`](#object-PodSecurityPolicyTemplate) after delete |
### Delete `PodSecurityPolicyTemplate`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/podSecurityPolicyTemplates/<resource_id>` |
|Input| `null` |
|Output| [Resource object `PodSecurityPolicyTemplate`](#object-PodSecurityPolicyTemplate) after delete |
---

## Object Schema Name: `PodSecurityPolicyTemplateProjectBinding`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `PodSecurityPolicyTemplateProjectBinding` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/podSecurityPolicyTemplateProjectBindings`|
|Input| `null` |
|Output| List of [resource object `PodSecurityPolicyTemplateProjectBinding`](#object-PodSecurityPolicyTemplateProjectBinding) |
### Create `PodSecurityPolicyTemplateProjectBinding`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/podSecurityPolicyTemplateProjectBindings` |
|Input| [Resource object `PodSecurityPolicyTemplateProjectBinding`](#object-PodSecurityPolicyTemplateProjectBinding) |
|Output| [Resource object `PodSecurityPolicyTemplateProjectBinding`](#object-PodSecurityPolicyTemplateProjectBinding) after create |
---

## Object Schema Name: `ClusterRoleTemplateBinding`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ClusterRoleTemplateBinding` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusterRoleTemplateBindings`|
|Input| `null` |
|Output| List of [resource object `ClusterRoleTemplateBinding`](#object-ClusterRoleTemplateBinding) |
### Get `ClusterRoleTemplateBinding`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusterRoleTemplateBindings/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ClusterRoleTemplateBinding`](#object-ClusterRoleTemplateBinding) |
### Create `ClusterRoleTemplateBinding`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterRoleTemplateBindings` |
|Input| [Resource object `ClusterRoleTemplateBinding`](#object-ClusterRoleTemplateBinding) |
|Output| [Resource object `ClusterRoleTemplateBinding`](#object-ClusterRoleTemplateBinding) after create |
### Update `ClusterRoleTemplateBinding`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusterRoleTemplateBindings/<resource_id>` |
|Input| [Resource object `ClusterRoleTemplateBinding`](#object-ClusterRoleTemplateBinding) to update |
|Output| [Resource object `ClusterRoleTemplateBinding`](#object-ClusterRoleTemplateBinding) after delete |
### Delete `ClusterRoleTemplateBinding`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusterRoleTemplateBindings/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ClusterRoleTemplateBinding`](#object-ClusterRoleTemplateBinding) after delete |
---

## Object Schema Name: `ProjectRoleTemplateBinding`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ProjectRoleTemplateBinding` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projectRoleTemplateBindings`|
|Input| `null` |
|Output| List of [resource object `ProjectRoleTemplateBinding`](#object-ProjectRoleTemplateBinding) |
### Get `ProjectRoleTemplateBinding`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projectRoleTemplateBindings/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ProjectRoleTemplateBinding`](#object-ProjectRoleTemplateBinding) |
### Create `ProjectRoleTemplateBinding`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectRoleTemplateBindings` |
|Input| [Resource object `ProjectRoleTemplateBinding`](#object-ProjectRoleTemplateBinding) |
|Output| [Resource object `ProjectRoleTemplateBinding`](#object-ProjectRoleTemplateBinding) after create |
### Update `ProjectRoleTemplateBinding`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projectRoleTemplateBindings/<resource_id>` |
|Input| [Resource object `ProjectRoleTemplateBinding`](#object-ProjectRoleTemplateBinding) to update |
|Output| [Resource object `ProjectRoleTemplateBinding`](#object-ProjectRoleTemplateBinding) after delete |
### Delete `ProjectRoleTemplateBinding`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projectRoleTemplateBindings/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ProjectRoleTemplateBinding`](#object-ProjectRoleTemplateBinding) after delete |
---

## Object Schema Name: `Cluster`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `Cluster` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusters`|
|Input| `null` |
|Output| List of [resource object `Cluster`](#object-Cluster) |
### Get `Cluster`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusters/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Cluster`](#object-Cluster) |
### Create `Cluster`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusters` |
|Input| [Resource object `Cluster`](#object-Cluster) |
|Output| [Resource object `Cluster`](#object-Cluster) after create |
### Update `Cluster`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusters/<resource_id>` |
|Input| [Resource object `Cluster`](#object-Cluster) to update |
|Output| [Resource object `Cluster`](#object-Cluster) after delete |
### Delete `Cluster`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusters/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Cluster`](#object-Cluster) after delete |
### Resource Action `disableMonitoring`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusters/<resource_id>/disableMonitoring` |
|Input| `null` |
|Output| `null` |
### Resource Action `enableMonitoring`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusters/<resource_id>/enableMonitoring` |
|Input| [object `monitoringInput`](#object-monitoringInput) |
|Output| `null` |
### Resource Action `exportYaml`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusters/<resource_id>/exportYaml` |
|Input| `null` |
|Output| [object `exportOutput`](#object-exportOutput) |
### Resource Action `generateKubeconfig`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusters/<resource_id>/generateKubeconfig` |
|Input| `null` |
|Output| [object `generateKubeConfigOutput`](#object-generateKubeConfigOutput) |
### Resource Action `importYaml`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusters/<resource_id>/importYaml` |
|Input| [object `importClusterYamlInput`](#object-importClusterYamlInput) |
|Output| [object `importYamlOutput`](#object-importYamlOutput) |
---

## Object Schema Name: `ClusterRegistrationToken`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ClusterRegistrationToken` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusterRegistrationTokens`|
|Input| `null` |
|Output| List of [resource object `ClusterRegistrationToken`](#object-ClusterRegistrationToken) |
### Get `ClusterRegistrationToken`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusterRegistrationTokens/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ClusterRegistrationToken`](#object-ClusterRegistrationToken) |
### Create `ClusterRegistrationToken`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterRegistrationTokens` |
|Input| [Resource object `ClusterRegistrationToken`](#object-ClusterRegistrationToken) |
|Output| [Resource object `ClusterRegistrationToken`](#object-ClusterRegistrationToken) after create |
### Update `ClusterRegistrationToken`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusterRegistrationTokens/<resource_id>` |
|Input| [Resource object `ClusterRegistrationToken`](#object-ClusterRegistrationToken) to update |
|Output| [Resource object `ClusterRegistrationToken`](#object-ClusterRegistrationToken) after delete |
### Delete `ClusterRegistrationToken`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusterRegistrationTokens/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ClusterRegistrationToken`](#object-ClusterRegistrationToken) after delete |
---

## Object Schema Name: `Catalog`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `Catalog` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/catalogs`|
|Input| `null` |
|Output| List of [resource object `Catalog`](#object-Catalog) |
### Get `Catalog`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/catalogs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Catalog`](#object-Catalog) |
### Create `Catalog`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/catalogs` |
|Input| [Resource object `Catalog`](#object-Catalog) |
|Output| [Resource object `Catalog`](#object-Catalog) after create |
### Update `Catalog`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/catalogs/<resource_id>` |
|Input| [Resource object `Catalog`](#object-Catalog) to update |
|Output| [Resource object `Catalog`](#object-Catalog) after delete |
### Delete `Catalog`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/catalogs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Catalog`](#object-Catalog) after delete |
### Resource Action `refresh`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/catalogs/<resource_id>/refresh` |
|Input| `null` |
|Output| `null` |
### Resource List Action `refresh`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/catalogs/<resource_id>` |
|Input| `null` |
|Output| `null` |
---

## Object Schema Name: `Template`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `Template` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/templates`|
|Input| `null` |
|Output| List of [resource object `Template`](#object-Template) |
### Get `Template`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/templates/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Template`](#object-Template) |
### Create `Template`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/templates` |
|Input| [Resource object `Template`](#object-Template) |
|Output| [Resource object `Template`](#object-Template) after create |
### Update `Template`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/templates/<resource_id>` |
|Input| [Resource object `Template`](#object-Template) to update |
|Output| [Resource object `Template`](#object-Template) after delete |
### Delete `Template`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/templates/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Template`](#object-Template) after delete |
---

## Object Schema Name: `TemplateVersion`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `TemplateVersion` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/templateVersions`|
|Input| `null` |
|Output| List of [resource object `TemplateVersion`](#object-TemplateVersion) |
### Get `TemplateVersion`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/templateVersions/<resource_id>` |
|Input| `null`|
|Output| [Resource object `TemplateVersion`](#object-TemplateVersion) |
### Create `TemplateVersion`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/templateVersions` |
|Input| [Resource object `TemplateVersion`](#object-TemplateVersion) |
|Output| [Resource object `TemplateVersion`](#object-TemplateVersion) after create |
### Update `TemplateVersion`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/templateVersions/<resource_id>` |
|Input| [Resource object `TemplateVersion`](#object-TemplateVersion) to update |
|Output| [Resource object `TemplateVersion`](#object-TemplateVersion) after delete |
### Delete `TemplateVersion`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/templateVersions/<resource_id>` |
|Input| `null` |
|Output| [Resource object `TemplateVersion`](#object-TemplateVersion) after delete |
---

## Object Schema Name: `TemplateContent`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `TemplateContent` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/templateContents`|
|Input| `null` |
|Output| List of [resource object `TemplateContent`](#object-TemplateContent) |
### Get `TemplateContent`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/templateContents/<resource_id>` |
|Input| `null`|
|Output| [Resource object `TemplateContent`](#object-TemplateContent) |
### Create `TemplateContent`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/templateContents` |
|Input| [Resource object `TemplateContent`](#object-TemplateContent) |
|Output| [Resource object `TemplateContent`](#object-TemplateContent) after create |
### Update `TemplateContent`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/templateContents/<resource_id>` |
|Input| [Resource object `TemplateContent`](#object-TemplateContent) to update |
|Output| [Resource object `TemplateContent`](#object-TemplateContent) after delete |
### Delete `TemplateContent`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/templateContents/<resource_id>` |
|Input| `null` |
|Output| [Resource object `TemplateContent`](#object-TemplateContent) after delete |
---

## Object Schema Name: `Group`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `Group` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/groups`|
|Input| `null` |
|Output| List of [resource object `Group`](#object-Group) |
### Get `Group`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/groups/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Group`](#object-Group) |
### Create `Group`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/groups` |
|Input| [Resource object `Group`](#object-Group) |
|Output| [Resource object `Group`](#object-Group) after create |
### Update `Group`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/groups/<resource_id>` |
|Input| [Resource object `Group`](#object-Group) to update |
|Output| [Resource object `Group`](#object-Group) after delete |
### Delete `Group`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/groups/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Group`](#object-Group) after delete |
---

## Object Schema Name: `GroupMember`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `GroupMember` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/groupMembers`|
|Input| `null` |
|Output| List of [resource object `GroupMember`](#object-GroupMember) |
### Get `GroupMember`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/groupMembers/<resource_id>` |
|Input| `null`|
|Output| [Resource object `GroupMember`](#object-GroupMember) |
### Create `GroupMember`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/groupMembers` |
|Input| [Resource object `GroupMember`](#object-GroupMember) |
|Output| [Resource object `GroupMember`](#object-GroupMember) after create |
### Update `GroupMember`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/groupMembers/<resource_id>` |
|Input| [Resource object `GroupMember`](#object-GroupMember) to update |
|Output| [Resource object `GroupMember`](#object-GroupMember) after delete |
### Delete `GroupMember`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/groupMembers/<resource_id>` |
|Input| `null` |
|Output| [Resource object `GroupMember`](#object-GroupMember) after delete |
---

## Object Schema Name: `Principal`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `Principal` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/principals`|
|Input| `null` |
|Output| List of [resource object `Principal`](#object-Principal) |
### Get `Principal`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/principals/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Principal`](#object-Principal) |
### Resource List Action `search`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/principals/<resource_id>` |
|Input| [object `searchPrincipalsInput`](#object-searchPrincipalsInput) |
|Output| [object `collection`](#object-collection) |
---

## Object Schema Name: `User`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `User` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/users`|
|Input| `null` |
|Output| List of [resource object `User`](#object-User) |
### Get `User`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/users/<resource_id>` |
|Input| `null`|
|Output| [Resource object `User`](#object-User) |
### Create `User`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/users` |
|Input| [Resource object `User`](#object-User) |
|Output| [Resource object `User`](#object-User) after create |
### Update `User`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/users/<resource_id>` |
|Input| [Resource object `User`](#object-User) to update |
|Output| [Resource object `User`](#object-User) after delete |
### Delete `User`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/users/<resource_id>` |
|Input| `null` |
|Output| [Resource object `User`](#object-User) after delete |
### Resource Action `setpassword`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/users/<resource_id>/setpassword` |
|Input| [object `setPasswordInput`](#object-setPasswordInput) |
|Output| [object `user`](#object-user) |
### Resource List Action `changepassword`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/users/<resource_id>` |
|Input| [object `changePasswordInput`](#object-changePasswordInput) |
|Output| `null` |
---

## Object Schema Name: `AuthConfig`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `AuthConfig` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/authConfigs`|
|Input| `null` |
|Output| List of [resource object `AuthConfig`](#object-AuthConfig) |
### Get `AuthConfig`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/authConfigs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `AuthConfig`](#object-AuthConfig) |
### Update `AuthConfig`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/authConfigs/<resource_id>` |
|Input| [Resource object `AuthConfig`](#object-AuthConfig) to update |
|Output| [Resource object `AuthConfig`](#object-AuthConfig) after delete |
### Delete `AuthConfig`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/authConfigs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `AuthConfig`](#object-AuthConfig) after delete |
---

## Object Schema Name: `LdapConfig`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `LdapConfig` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/ldapConfigs`|
|Input| `null` |
|Output| List of [resource object `LdapConfig`](#object-LdapConfig) |
### Get `LdapConfig`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/ldapConfigs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `LdapConfig`](#object-LdapConfig) |
### Create `LdapConfig`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/ldapConfigs` |
|Input| [Resource object `LdapConfig`](#object-LdapConfig) |
|Output| [Resource object `LdapConfig`](#object-LdapConfig) after create |
### Update `LdapConfig`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/ldapConfigs/<resource_id>` |
|Input| [Resource object `LdapConfig`](#object-LdapConfig) to update |
|Output| [Resource object `LdapConfig`](#object-LdapConfig) after delete |
### Delete `LdapConfig`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/ldapConfigs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `LdapConfig`](#object-LdapConfig) after delete |
---

## Object Schema Name: `Token`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `Token` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/tokens`|
|Input| `null` |
|Output| List of [resource object `Token`](#object-Token) |
### Get `Token`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/tokens/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Token`](#object-Token) |
### Create `Token`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/tokens` |
|Input| [Resource object `Token`](#object-Token) |
|Output| [Resource object `Token`](#object-Token) after create |
### Update `Token`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/tokens/<resource_id>` |
|Input| [Resource object `Token`](#object-Token) to update |
|Output| [Resource object `Token`](#object-Token) after delete |
### Delete `Token`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/tokens/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Token`](#object-Token) after delete |
### Resource List Action `logout`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/tokens/<resource_id>` |
|Input| `null` |
|Output| `null` |
---

## Object Schema Name: `DynamicSchema`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `DynamicSchema` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/dynamicSchemas`|
|Input| `null` |
|Output| List of [resource object `DynamicSchema`](#object-DynamicSchema) |
### Get `DynamicSchema`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/dynamicSchemas/<resource_id>` |
|Input| `null`|
|Output| [Resource object `DynamicSchema`](#object-DynamicSchema) |
### Create `DynamicSchema`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/dynamicSchemas` |
|Input| [Resource object `DynamicSchema`](#object-DynamicSchema) |
|Output| [Resource object `DynamicSchema`](#object-DynamicSchema) after create |
### Update `DynamicSchema`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/dynamicSchemas/<resource_id>` |
|Input| [Resource object `DynamicSchema`](#object-DynamicSchema) to update |
|Output| [Resource object `DynamicSchema`](#object-DynamicSchema) after delete |
### Delete `DynamicSchema`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/dynamicSchemas/<resource_id>` |
|Input| `null` |
|Output| [Resource object `DynamicSchema`](#object-DynamicSchema) after delete |
---

## Object Schema Name: `Preference`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `Preference` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/preferences`|
|Input| `null` |
|Output| List of [resource object `Preference`](#object-Preference) |
### Get `Preference`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/preferences/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Preference`](#object-Preference) |
### Create `Preference`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/preferences` |
|Input| [Resource object `Preference`](#object-Preference) |
|Output| [Resource object `Preference`](#object-Preference) after create |
### Update `Preference`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/preferences/<resource_id>` |
|Input| [Resource object `Preference`](#object-Preference) to update |
|Output| [Resource object `Preference`](#object-Preference) after delete |
### Delete `Preference`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/preferences/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Preference`](#object-Preference) after delete |
---

## Object Schema Name: `ProjectNetworkPolicy`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ProjectNetworkPolicy` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projectNetworkPolicies`|
|Input| `null` |
|Output| List of [resource object `ProjectNetworkPolicy`](#object-ProjectNetworkPolicy) |
### Get `ProjectNetworkPolicy`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projectNetworkPolicies/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ProjectNetworkPolicy`](#object-ProjectNetworkPolicy) |
---

## Object Schema Name: `ClusterLogging`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ClusterLogging` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusterLoggings`|
|Input| `null` |
|Output| List of [resource object `ClusterLogging`](#object-ClusterLogging) |
### Get `ClusterLogging`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusterLoggings/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ClusterLogging`](#object-ClusterLogging) |
### Create `ClusterLogging`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterLoggings` |
|Input| [Resource object `ClusterLogging`](#object-ClusterLogging) |
|Output| [Resource object `ClusterLogging`](#object-ClusterLogging) after create |
### Update `ClusterLogging`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusterLoggings/<resource_id>` |
|Input| [Resource object `ClusterLogging`](#object-ClusterLogging) to update |
|Output| [Resource object `ClusterLogging`](#object-ClusterLogging) after delete |
### Delete `ClusterLogging`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusterLoggings/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ClusterLogging`](#object-ClusterLogging) after delete |
---

## Object Schema Name: `ProjectLogging`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ProjectLogging` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projectLoggings`|
|Input| `null` |
|Output| List of [resource object `ProjectLogging`](#object-ProjectLogging) |
### Get `ProjectLogging`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projectLoggings/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ProjectLogging`](#object-ProjectLogging) |
### Create `ProjectLogging`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectLoggings` |
|Input| [Resource object `ProjectLogging`](#object-ProjectLogging) |
|Output| [Resource object `ProjectLogging`](#object-ProjectLogging) after create |
### Update `ProjectLogging`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projectLoggings/<resource_id>` |
|Input| [Resource object `ProjectLogging`](#object-ProjectLogging) to update |
|Output| [Resource object `ProjectLogging`](#object-ProjectLogging) after delete |
### Delete `ProjectLogging`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projectLoggings/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ProjectLogging`](#object-ProjectLogging) after delete |
---

## Object Schema Name: `ListenConfig`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ListenConfig` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/listenConfigs`|
|Input| `null` |
|Output| List of [resource object `ListenConfig`](#object-ListenConfig) |
### Get `ListenConfig`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/listenConfigs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ListenConfig`](#object-ListenConfig) |
### Create `ListenConfig`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/listenConfigs` |
|Input| [Resource object `ListenConfig`](#object-ListenConfig) |
|Output| [Resource object `ListenConfig`](#object-ListenConfig) after create |
### Update `ListenConfig`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/listenConfigs/<resource_id>` |
|Input| [Resource object `ListenConfig`](#object-ListenConfig) to update |
|Output| [Resource object `ListenConfig`](#object-ListenConfig) after delete |
### Delete `ListenConfig`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/listenConfigs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ListenConfig`](#object-ListenConfig) after delete |
---

## Object Schema Name: `Setting`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `Setting` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/settings`|
|Input| `null` |
|Output| List of [resource object `Setting`](#object-Setting) |
### Get `Setting`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/settings/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Setting`](#object-Setting) |
### Create `Setting`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/settings` |
|Input| [Resource object `Setting`](#object-Setting) |
|Output| [Resource object `Setting`](#object-Setting) after create |
### Update `Setting`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/settings/<resource_id>` |
|Input| [Resource object `Setting`](#object-Setting) to update |
|Output| [Resource object `Setting`](#object-Setting) after delete |
### Delete `Setting`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/settings/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Setting`](#object-Setting) after delete |
---

## Object Schema Name: `ClusterAlert`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ClusterAlert` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusterAlerts`|
|Input| `null` |
|Output| List of [resource object `ClusterAlert`](#object-ClusterAlert) |
### Get `ClusterAlert`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusterAlerts/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ClusterAlert`](#object-ClusterAlert) |
### Create `ClusterAlert`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterAlerts` |
|Input| [Resource object `ClusterAlert`](#object-ClusterAlert) |
|Output| [Resource object `ClusterAlert`](#object-ClusterAlert) after create |
### Update `ClusterAlert`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusterAlerts/<resource_id>` |
|Input| [Resource object `ClusterAlert`](#object-ClusterAlert) to update |
|Output| [Resource object `ClusterAlert`](#object-ClusterAlert) after delete |
### Delete `ClusterAlert`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusterAlerts/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ClusterAlert`](#object-ClusterAlert) after delete |
---

## Object Schema Name: `ProjectAlert`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ProjectAlert` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projectAlerts`|
|Input| `null` |
|Output| List of [resource object `ProjectAlert`](#object-ProjectAlert) |
### Get `ProjectAlert`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projectAlerts/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ProjectAlert`](#object-ProjectAlert) |
### Create `ProjectAlert`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectAlerts` |
|Input| [Resource object `ProjectAlert`](#object-ProjectAlert) |
|Output| [Resource object `ProjectAlert`](#object-ProjectAlert) after create |
### Update `ProjectAlert`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projectAlerts/<resource_id>` |
|Input| [Resource object `ProjectAlert`](#object-ProjectAlert) to update |
|Output| [Resource object `ProjectAlert`](#object-ProjectAlert) after delete |
### Delete `ProjectAlert`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projectAlerts/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ProjectAlert`](#object-ProjectAlert) after delete |
---

## Object Schema Name: `Notifier`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `Notifier` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/notifiers`|
|Input| `null` |
|Output| List of [resource object `Notifier`](#object-Notifier) |
### Get `Notifier`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/notifiers/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Notifier`](#object-Notifier) |
### Create `Notifier`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/notifiers` |
|Input| [Resource object `Notifier`](#object-Notifier) |
|Output| [Resource object `Notifier`](#object-Notifier) after create |
### Update `Notifier`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/notifiers/<resource_id>` |
|Input| [Resource object `Notifier`](#object-Notifier) to update |
|Output| [Resource object `Notifier`](#object-Notifier) after delete |
### Delete `Notifier`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/notifiers/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Notifier`](#object-Notifier) after delete |
### Resource Action `send`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/notifiers/<resource_id>/send` |
|Input| [object `notification`](#object-notification) |
|Output| `null` |
### Resource List Action `send`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/notifiers/<resource_id>` |
|Input| [object `notification`](#object-notification) |
|Output| `null` |
---

## Object Schema Name: `ClusterAlertGroup`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ClusterAlertGroup` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusterAlertGroups`|
|Input| `null` |
|Output| List of [resource object `ClusterAlertGroup`](#object-ClusterAlertGroup) |
### Get `ClusterAlertGroup`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusterAlertGroups/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ClusterAlertGroup`](#object-ClusterAlertGroup) |
### Create `ClusterAlertGroup`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterAlertGroups` |
|Input| [Resource object `ClusterAlertGroup`](#object-ClusterAlertGroup) |
|Output| [Resource object `ClusterAlertGroup`](#object-ClusterAlertGroup) after create |
### Update `ClusterAlertGroup`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusterAlertGroups/<resource_id>` |
|Input| [Resource object `ClusterAlertGroup`](#object-ClusterAlertGroup) to update |
|Output| [Resource object `ClusterAlertGroup`](#object-ClusterAlertGroup) after delete |
### Delete `ClusterAlertGroup`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusterAlertGroups/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ClusterAlertGroup`](#object-ClusterAlertGroup) after delete |
---

## Object Schema Name: `ProjectAlertGroup`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ProjectAlertGroup` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projectAlertGroups`|
|Input| `null` |
|Output| List of [resource object `ProjectAlertGroup`](#object-ProjectAlertGroup) |
### Get `ProjectAlertGroup`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projectAlertGroups/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ProjectAlertGroup`](#object-ProjectAlertGroup) |
### Create `ProjectAlertGroup`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectAlertGroups` |
|Input| [Resource object `ProjectAlertGroup`](#object-ProjectAlertGroup) |
|Output| [Resource object `ProjectAlertGroup`](#object-ProjectAlertGroup) after create |
### Update `ProjectAlertGroup`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projectAlertGroups/<resource_id>` |
|Input| [Resource object `ProjectAlertGroup`](#object-ProjectAlertGroup) to update |
|Output| [Resource object `ProjectAlertGroup`](#object-ProjectAlertGroup) after delete |
### Delete `ProjectAlertGroup`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projectAlertGroups/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ProjectAlertGroup`](#object-ProjectAlertGroup) after delete |
---

## Object Schema Name: `ClusterAlertRule`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ClusterAlertRule` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusterAlertRules`|
|Input| `null` |
|Output| List of [resource object `ClusterAlertRule`](#object-ClusterAlertRule) |
### Get `ClusterAlertRule`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusterAlertRules/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ClusterAlertRule`](#object-ClusterAlertRule) |
### Create `ClusterAlertRule`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterAlertRules` |
|Input| [Resource object `ClusterAlertRule`](#object-ClusterAlertRule) |
|Output| [Resource object `ClusterAlertRule`](#object-ClusterAlertRule) after create |
### Update `ClusterAlertRule`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusterAlertRules/<resource_id>` |
|Input| [Resource object `ClusterAlertRule`](#object-ClusterAlertRule) to update |
|Output| [Resource object `ClusterAlertRule`](#object-ClusterAlertRule) after delete |
### Delete `ClusterAlertRule`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusterAlertRules/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ClusterAlertRule`](#object-ClusterAlertRule) after delete |
### Resource Action `activate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterAlertRules/<resource_id>/activate` |
|Input| `null` |
|Output| `null` |
### Resource Action `deactivate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterAlertRules/<resource_id>/deactivate` |
|Input| `null` |
|Output| `null` |
### Resource Action `mute`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterAlertRules/<resource_id>/mute` |
|Input| `null` |
|Output| `null` |
### Resource Action `unmute`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterAlertRules/<resource_id>/unmute` |
|Input| `null` |
|Output| `null` |
---

## Object Schema Name: `ProjectAlertRule`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ProjectAlertRule` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projectAlertRules`|
|Input| `null` |
|Output| List of [resource object `ProjectAlertRule`](#object-ProjectAlertRule) |
### Get `ProjectAlertRule`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projectAlertRules/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ProjectAlertRule`](#object-ProjectAlertRule) |
### Create `ProjectAlertRule`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectAlertRules` |
|Input| [Resource object `ProjectAlertRule`](#object-ProjectAlertRule) |
|Output| [Resource object `ProjectAlertRule`](#object-ProjectAlertRule) after create |
### Update `ProjectAlertRule`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projectAlertRules/<resource_id>` |
|Input| [Resource object `ProjectAlertRule`](#object-ProjectAlertRule) to update |
|Output| [Resource object `ProjectAlertRule`](#object-ProjectAlertRule) after delete |
### Delete `ProjectAlertRule`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projectAlertRules/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ProjectAlertRule`](#object-ProjectAlertRule) after delete |
### Resource Action `activate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectAlertRules/<resource_id>/activate` |
|Input| `null` |
|Output| `null` |
### Resource Action `deactivate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectAlertRules/<resource_id>/deactivate` |
|Input| `null` |
|Output| `null` |
### Resource Action `mute`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectAlertRules/<resource_id>/mute` |
|Input| `null` |
|Output| `null` |
### Resource Action `unmute`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectAlertRules/<resource_id>/unmute` |
|Input| `null` |
|Output| `null` |
---

## Object Schema Name: `ComposeConfig`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ComposeConfig` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/composeConfigs`|
|Input| `null` |
|Output| List of [resource object `ComposeConfig`](#object-ComposeConfig) |
### Get `ComposeConfig`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/composeConfigs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ComposeConfig`](#object-ComposeConfig) |
### Create `ComposeConfig`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/composeConfigs` |
|Input| [Resource object `ComposeConfig`](#object-ComposeConfig) |
|Output| [Resource object `ComposeConfig`](#object-ComposeConfig) after create |
### Update `ComposeConfig`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/composeConfigs/<resource_id>` |
|Input| [Resource object `ComposeConfig`](#object-ComposeConfig) to update |
|Output| [Resource object `ComposeConfig`](#object-ComposeConfig) after delete |
### Delete `ComposeConfig`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/composeConfigs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ComposeConfig`](#object-ComposeConfig) after delete |
---

## Object Schema Name: `ProjectCatalog`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ProjectCatalog` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projectCatalogs`|
|Input| `null` |
|Output| List of [resource object `ProjectCatalog`](#object-ProjectCatalog) |
### Get `ProjectCatalog`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projectCatalogs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ProjectCatalog`](#object-ProjectCatalog) |
### Create `ProjectCatalog`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectCatalogs` |
|Input| [Resource object `ProjectCatalog`](#object-ProjectCatalog) |
|Output| [Resource object `ProjectCatalog`](#object-ProjectCatalog) after create |
### Update `ProjectCatalog`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projectCatalogs/<resource_id>` |
|Input| [Resource object `ProjectCatalog`](#object-ProjectCatalog) to update |
|Output| [Resource object `ProjectCatalog`](#object-ProjectCatalog) after delete |
### Delete `ProjectCatalog`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projectCatalogs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ProjectCatalog`](#object-ProjectCatalog) after delete |
### Resource Action `refresh`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectCatalogs/<resource_id>/refresh` |
|Input| `null` |
|Output| `null` |
### Resource List Action `refresh`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectCatalogs/<resource_id>` |
|Input| `null` |
|Output| `null` |
---

## Object Schema Name: `ClusterCatalog`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ClusterCatalog` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusterCatalogs`|
|Input| `null` |
|Output| List of [resource object `ClusterCatalog`](#object-ClusterCatalog) |
### Get `ClusterCatalog`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusterCatalogs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ClusterCatalog`](#object-ClusterCatalog) |
### Create `ClusterCatalog`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterCatalogs` |
|Input| [Resource object `ClusterCatalog`](#object-ClusterCatalog) |
|Output| [Resource object `ClusterCatalog`](#object-ClusterCatalog) after create |
### Update `ClusterCatalog`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusterCatalogs/<resource_id>` |
|Input| [Resource object `ClusterCatalog`](#object-ClusterCatalog) to update |
|Output| [Resource object `ClusterCatalog`](#object-ClusterCatalog) after delete |
### Delete `ClusterCatalog`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusterCatalogs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ClusterCatalog`](#object-ClusterCatalog) after delete |
### Resource Action `refresh`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterCatalogs/<resource_id>/refresh` |
|Input| `null` |
|Output| `null` |
### Resource List Action `refresh`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterCatalogs/<resource_id>` |
|Input| `null` |
|Output| `null` |
---

## Object Schema Name: `MultiClusterApp`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `MultiClusterApp` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/multiClusterApps`|
|Input| `null` |
|Output| List of [resource object `MultiClusterApp`](#object-MultiClusterApp) |
### Get `MultiClusterApp`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/multiClusterApps/<resource_id>` |
|Input| `null`|
|Output| [Resource object `MultiClusterApp`](#object-MultiClusterApp) |
### Create `MultiClusterApp`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/multiClusterApps` |
|Input| [Resource object `MultiClusterApp`](#object-MultiClusterApp) |
|Output| [Resource object `MultiClusterApp`](#object-MultiClusterApp) after create |
### Update `MultiClusterApp`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/multiClusterApps/<resource_id>` |
|Input| [Resource object `MultiClusterApp`](#object-MultiClusterApp) to update |
|Output| [Resource object `MultiClusterApp`](#object-MultiClusterApp) after delete |
### Delete `MultiClusterApp`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/multiClusterApps/<resource_id>` |
|Input| `null` |
|Output| [Resource object `MultiClusterApp`](#object-MultiClusterApp) after delete |
---

## Object Schema Name: `GlobalDNS`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `GlobalDNS` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/globalDnses`|
|Input| `null` |
|Output| List of [resource object `GlobalDNS`](#object-GlobalDNS) |
### Get `GlobalDNS`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/globalDnses/<resource_id>` |
|Input| `null`|
|Output| [Resource object `GlobalDNS`](#object-GlobalDNS) |
### Create `GlobalDNS`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/globalDnses` |
|Input| [Resource object `GlobalDNS`](#object-GlobalDNS) |
|Output| [Resource object `GlobalDNS`](#object-GlobalDNS) after create |
### Update `GlobalDNS`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/globalDnses/<resource_id>` |
|Input| [Resource object `GlobalDNS`](#object-GlobalDNS) to update |
|Output| [Resource object `GlobalDNS`](#object-GlobalDNS) after delete |
### Delete `GlobalDNS`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/globalDnses/<resource_id>` |
|Input| `null` |
|Output| [Resource object `GlobalDNS`](#object-GlobalDNS) after delete |
---

## Object Schema Name: `GlobalDNSProvider`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `GlobalDNSProvider` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/globalDnsProviders`|
|Input| `null` |
|Output| List of [resource object `GlobalDNSProvider`](#object-GlobalDNSProvider) |
### Get `GlobalDNSProvider`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/globalDnsProviders/<resource_id>` |
|Input| `null`|
|Output| [Resource object `GlobalDNSProvider`](#object-GlobalDNSProvider) |
### Create `GlobalDNSProvider`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/globalDnsProviders` |
|Input| [Resource object `GlobalDNSProvider`](#object-GlobalDNSProvider) |
|Output| [Resource object `GlobalDNSProvider`](#object-GlobalDNSProvider) after create |
### Update `GlobalDNSProvider`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/globalDnsProviders/<resource_id>` |
|Input| [Resource object `GlobalDNSProvider`](#object-GlobalDNSProvider) to update |
|Output| [Resource object `GlobalDNSProvider`](#object-GlobalDNSProvider) after delete |
### Delete `GlobalDNSProvider`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/globalDnsProviders/<resource_id>` |
|Input| `null` |
|Output| [Resource object `GlobalDNSProvider`](#object-GlobalDNSProvider) after delete |
---

## Object Schema Name: `KontainerDriver`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `KontainerDriver` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/kontainerDrivers`|
|Input| `null` |
|Output| List of [resource object `KontainerDriver`](#object-KontainerDriver) |
### Get `KontainerDriver`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/kontainerDrivers/<resource_id>` |
|Input| `null`|
|Output| [Resource object `KontainerDriver`](#object-KontainerDriver) |
### Create `KontainerDriver`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/kontainerDrivers` |
|Input| [Resource object `KontainerDriver`](#object-KontainerDriver) |
|Output| [Resource object `KontainerDriver`](#object-KontainerDriver) after create |
### Update `KontainerDriver`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/kontainerDrivers/<resource_id>` |
|Input| [Resource object `KontainerDriver`](#object-KontainerDriver) to update |
|Output| [Resource object `KontainerDriver`](#object-KontainerDriver) after delete |
### Delete `KontainerDriver`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/kontainerDrivers/<resource_id>` |
|Input| `null` |
|Output| [Resource object `KontainerDriver`](#object-KontainerDriver) after delete |
### Resource Action `activate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/kontainerDrivers/<resource_id>/activate` |
|Input| `null` |
|Output| `null` |
### Resource Action `deactivate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/kontainerDrivers/<resource_id>/deactivate` |
|Input| `null` |
|Output| `null` |
---

## Object Schema Name: `MonitorMetric`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `MonitorMetric` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/monitorMetrics`|
|Input| `null` |
|Output| List of [resource object `MonitorMetric`](#object-MonitorMetric) |
### Get `MonitorMetric`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/monitorMetrics/<resource_id>` |
|Input| `null`|
|Output| [Resource object `MonitorMetric`](#object-MonitorMetric) |
### Create `MonitorMetric`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/monitorMetrics` |
|Input| [Resource object `MonitorMetric`](#object-MonitorMetric) |
|Output| [Resource object `MonitorMetric`](#object-MonitorMetric) after create |
### Update `MonitorMetric`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/monitorMetrics/<resource_id>` |
|Input| [Resource object `MonitorMetric`](#object-MonitorMetric) to update |
|Output| [Resource object `MonitorMetric`](#object-MonitorMetric) after delete |
### Delete `MonitorMetric`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/monitorMetrics/<resource_id>` |
|Input| `null` |
|Output| [Resource object `MonitorMetric`](#object-MonitorMetric) after delete |
### Resource List Action `listclustermetricname`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/monitorMetrics/<resource_id>` |
|Input| [object `clusterMetricNamesInput`](#object-clusterMetricNamesInput) |
|Output| [object `metricNamesOutput`](#object-metricNamesOutput) |
### Resource List Action `listprojectmetricname`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/monitorMetrics/<resource_id>` |
|Input| [object `projectMetricNamesInput`](#object-projectMetricNamesInput) |
|Output| [object `metricNamesOutput`](#object-metricNamesOutput) |
### Resource List Action `querycluster`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/monitorMetrics/<resource_id>` |
|Input| [object `queryClusterMetricInput`](#object-queryClusterMetricInput) |
|Output| [object `queryMetricOutput`](#object-queryMetricOutput) |
### Resource List Action `queryproject`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/monitorMetrics/<resource_id>` |
|Input| [object `queryProjectMetricInput`](#object-queryProjectMetricInput) |
|Output| [object `queryMetricOutput`](#object-queryMetricOutput) |
---

## Object Schema Name: `ClusterMonitorGraph`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ClusterMonitorGraph` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusterMonitorGraphs`|
|Input| `null` |
|Output| List of [resource object `ClusterMonitorGraph`](#object-ClusterMonitorGraph) |
### Get `ClusterMonitorGraph`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusterMonitorGraphs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ClusterMonitorGraph`](#object-ClusterMonitorGraph) |
### Create `ClusterMonitorGraph`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterMonitorGraphs` |
|Input| [Resource object `ClusterMonitorGraph`](#object-ClusterMonitorGraph) |
|Output| [Resource object `ClusterMonitorGraph`](#object-ClusterMonitorGraph) after create |
### Update `ClusterMonitorGraph`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusterMonitorGraphs/<resource_id>` |
|Input| [Resource object `ClusterMonitorGraph`](#object-ClusterMonitorGraph) to update |
|Output| [Resource object `ClusterMonitorGraph`](#object-ClusterMonitorGraph) after delete |
### Delete `ClusterMonitorGraph`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusterMonitorGraphs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ClusterMonitorGraph`](#object-ClusterMonitorGraph) after delete |
### Resource List Action `query`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusterMonitorGraphs/<resource_id>` |
|Input| [object `queryGraphInput`](#object-queryGraphInput) |
|Output| [object `queryClusterGraphOutput`](#object-queryClusterGraphOutput) |
---

## Object Schema Name: `ProjectMonitorGraph`
### K8s Object Meta
|||
|---|---|
|API Group| `management.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| Management |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/management.cattle.io/v3 |
### Get `ProjectMonitorGraph` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projectMonitorGraphs`|
|Input| `null` |
|Output| List of [resource object `ProjectMonitorGraph`](#object-ProjectMonitorGraph) |
### Get `ProjectMonitorGraph`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projectMonitorGraphs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ProjectMonitorGraph`](#object-ProjectMonitorGraph) |
### Create `ProjectMonitorGraph`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectMonitorGraphs` |
|Input| [Resource object `ProjectMonitorGraph`](#object-ProjectMonitorGraph) |
|Output| [Resource object `ProjectMonitorGraph`](#object-ProjectMonitorGraph) after create |
### Update `ProjectMonitorGraph`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projectMonitorGraphs/<resource_id>` |
|Input| [Resource object `ProjectMonitorGraph`](#object-ProjectMonitorGraph) to update |
|Output| [Resource object `ProjectMonitorGraph`](#object-ProjectMonitorGraph) after delete |
### Delete `ProjectMonitorGraph`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projectMonitorGraphs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ProjectMonitorGraph`](#object-ProjectMonitorGraph) after delete |
### Resource List Action `query`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projectMonitorGraphs/<resource_id>` |
|Input| [object `queryGraphInput`](#object-queryGraphInput) |
|Output| [object `queryProjectGraphOutput`](#object-queryProjectGraphOutput) |
---

## Object Schema Name: `Namespace`
### K8s Object Meta
|||
|---|---|
|API Group| `cluster.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/core/v1 |
### Get `Namespace` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusters/<cluster_id>/namespaces`|
|Input| `null` |
|Output| List of [resource object `Namespace`](#object-Namespace) |
### Get `Namespace`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusters/<cluster_id>/namespaces/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Namespace`](#object-Namespace) |
### Create `Namespace`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusters/<cluster_id>/namespaces` |
|Input| [Resource object `Namespace`](#object-Namespace) |
|Output| [Resource object `Namespace`](#object-Namespace) after create |
### Update `Namespace`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusters/<cluster_id>/namespaces/<resource_id>` |
|Input| [Resource object `Namespace`](#object-Namespace) to update |
|Output| [Resource object `Namespace`](#object-Namespace) after delete |
### Delete `Namespace`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusters/<cluster_id>/namespaces/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Namespace`](#object-Namespace) after delete |
### Resource Action `move`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusters/<cluster_id>/namespaces/<resource_id>/move` |
|Input| [object `namespaceMove`](#object-namespaceMove) |
|Output| `null` |
---

## Object Schema Name: `PersistentVolume`
### K8s Object Meta
|||
|---|---|
|API Group| `cluster.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/core/v1 |
### Get `PersistentVolume` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusters/<cluster_id>/persistentVolumes`|
|Input| `null` |
|Output| List of [resource object `PersistentVolume`](#object-PersistentVolume) |
### Get `PersistentVolume`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusters/<cluster_id>/persistentVolumes/<resource_id>` |
|Input| `null`|
|Output| [Resource object `PersistentVolume`](#object-PersistentVolume) |
### Create `PersistentVolume`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusters/<cluster_id>/persistentVolumes` |
|Input| [Resource object `PersistentVolume`](#object-PersistentVolume) |
|Output| [Resource object `PersistentVolume`](#object-PersistentVolume) after create |
### Update `PersistentVolume`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusters/<cluster_id>/persistentVolumes/<resource_id>` |
|Input| [Resource object `PersistentVolume`](#object-PersistentVolume) to update |
|Output| [Resource object `PersistentVolume`](#object-PersistentVolume) after delete |
### Delete `PersistentVolume`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusters/<cluster_id>/persistentVolumes/<resource_id>` |
|Input| `null` |
|Output| [Resource object `PersistentVolume`](#object-PersistentVolume) after delete |
---

## Object Schema Name: `StorageClass`
### K8s Object Meta
|||
|---|---|
|API Group| `cluster.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/storage/v1 |
### Get `StorageClass` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/clusters/<cluster_id>/storageClasses`|
|Input| `null` |
|Output| List of [resource object `StorageClass`](#object-StorageClass) |
### Get `StorageClass`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/clusters/<cluster_id>/storageClasses/<resource_id>` |
|Input| `null`|
|Output| [Resource object `StorageClass`](#object-StorageClass) |
### Create `StorageClass`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/clusters/<cluster_id>/storageClasses` |
|Input| [Resource object `StorageClass`](#object-StorageClass) |
|Output| [Resource object `StorageClass`](#object-StorageClass) after create |
### Update `StorageClass`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/clusters/<cluster_id>/storageClasses/<resource_id>` |
|Input| [Resource object `StorageClass`](#object-StorageClass) to update |
|Output| [Resource object `StorageClass`](#object-StorageClass) after delete |
### Delete `StorageClass`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/clusters/<cluster_id>/storageClasses/<resource_id>` |
|Input| `null` |
|Output| [Resource object `StorageClass`](#object-StorageClass) after delete |
---

## Object Schema Name: `PersistentVolumeClaim`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/core/v1 |
### Get `PersistentVolumeClaim` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/persistentVolumeClaims`|
|Input| `null` |
|Output| List of [resource object `PersistentVolumeClaim`](#object-PersistentVolumeClaim) |
### Get `PersistentVolumeClaim`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/persistentVolumeClaims/<resource_id>` |
|Input| `null`|
|Output| [Resource object `PersistentVolumeClaim`](#object-PersistentVolumeClaim) |
### Create `PersistentVolumeClaim`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/persistentVolumeClaims` |
|Input| [Resource object `PersistentVolumeClaim`](#object-PersistentVolumeClaim) |
|Output| [Resource object `PersistentVolumeClaim`](#object-PersistentVolumeClaim) after create |
### Update `PersistentVolumeClaim`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/persistentVolumeClaims/<resource_id>` |
|Input| [Resource object `PersistentVolumeClaim`](#object-PersistentVolumeClaim) to update |
|Output| [Resource object `PersistentVolumeClaim`](#object-PersistentVolumeClaim) after delete |
### Delete `PersistentVolumeClaim`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/persistentVolumeClaims/<resource_id>` |
|Input| `null` |
|Output| [Resource object `PersistentVolumeClaim`](#object-PersistentVolumeClaim) after delete |
---

## Object Schema Name: `ConfigMap`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/core/v1 |
### Get `ConfigMap` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/configMaps`|
|Input| `null` |
|Output| List of [resource object `ConfigMap`](#object-ConfigMap) |
### Get `ConfigMap`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/configMaps/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ConfigMap`](#object-ConfigMap) |
### Create `ConfigMap`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/configMaps` |
|Input| [Resource object `ConfigMap`](#object-ConfigMap) |
|Output| [Resource object `ConfigMap`](#object-ConfigMap) after create |
### Update `ConfigMap`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/configMaps/<resource_id>` |
|Input| [Resource object `ConfigMap`](#object-ConfigMap) to update |
|Output| [Resource object `ConfigMap`](#object-ConfigMap) after delete |
### Delete `ConfigMap`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/configMaps/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ConfigMap`](#object-ConfigMap) after delete |
---

## Object Schema Name: `Ingress`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/extensions/v1beta1 |
### Get `Ingress` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/ingresses`|
|Input| `null` |
|Output| List of [resource object `Ingress`](#object-Ingress) |
### Get `Ingress`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/ingresses/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Ingress`](#object-Ingress) |
### Create `Ingress`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/ingresses` |
|Input| [Resource object `Ingress`](#object-Ingress) |
|Output| [Resource object `Ingress`](#object-Ingress) after create |
### Update `Ingress`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/ingresses/<resource_id>` |
|Input| [Resource object `Ingress`](#object-Ingress) to update |
|Output| [Resource object `Ingress`](#object-Ingress) after delete |
### Delete `Ingress`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/ingresses/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Ingress`](#object-Ingress) after delete |
---

## Object Schema Name: `Secret`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/core/v1 |
### Get `Secret` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/secrets`|
|Input| `null` |
|Output| List of [resource object `Secret`](#object-Secret) |
### Get `Secret`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/secrets/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Secret`](#object-Secret) |
### Create `Secret`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/secrets` |
|Input| [Resource object `Secret`](#object-Secret) |
|Output| [Resource object `Secret`](#object-Secret) after create |
### Update `Secret`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/secrets/<resource_id>` |
|Input| [Resource object `Secret`](#object-Secret) to update |
|Output| [Resource object `Secret`](#object-Secret) after delete |
### Delete `Secret`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/secrets/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Secret`](#object-Secret) after delete |
---

## Object Schema Name: `ServiceAccountToken`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `ServiceAccountToken` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/serviceAccountTokens`|
|Input| `null` |
|Output| List of [resource object `ServiceAccountToken`](#object-ServiceAccountToken) |
### Get `ServiceAccountToken`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/serviceAccountTokens/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ServiceAccountToken`](#object-ServiceAccountToken) |
### Create `ServiceAccountToken`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/serviceAccountTokens` |
|Input| [Resource object `ServiceAccountToken`](#object-ServiceAccountToken) |
|Output| [Resource object `ServiceAccountToken`](#object-ServiceAccountToken) after create |
### Update `ServiceAccountToken`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/serviceAccountTokens/<resource_id>` |
|Input| [Resource object `ServiceAccountToken`](#object-ServiceAccountToken) to update |
|Output| [Resource object `ServiceAccountToken`](#object-ServiceAccountToken) after delete |
### Delete `ServiceAccountToken`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/serviceAccountTokens/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ServiceAccountToken`](#object-ServiceAccountToken) after delete |
---

## Object Schema Name: `DockerCredential`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `DockerCredential` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/dockerCredentials`|
|Input| `null` |
|Output| List of [resource object `DockerCredential`](#object-DockerCredential) |
### Get `DockerCredential`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/dockerCredentials/<resource_id>` |
|Input| `null`|
|Output| [Resource object `DockerCredential`](#object-DockerCredential) |
### Create `DockerCredential`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/dockerCredentials` |
|Input| [Resource object `DockerCredential`](#object-DockerCredential) |
|Output| [Resource object `DockerCredential`](#object-DockerCredential) after create |
### Update `DockerCredential`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/dockerCredentials/<resource_id>` |
|Input| [Resource object `DockerCredential`](#object-DockerCredential) to update |
|Output| [Resource object `DockerCredential`](#object-DockerCredential) after delete |
### Delete `DockerCredential`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/dockerCredentials/<resource_id>` |
|Input| `null` |
|Output| [Resource object `DockerCredential`](#object-DockerCredential) after delete |
---

## Object Schema Name: `Certificate`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `Certificate` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/certificates`|
|Input| `null` |
|Output| List of [resource object `Certificate`](#object-Certificate) |
### Get `Certificate`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/certificates/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Certificate`](#object-Certificate) |
### Create `Certificate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/certificates` |
|Input| [Resource object `Certificate`](#object-Certificate) |
|Output| [Resource object `Certificate`](#object-Certificate) after create |
### Update `Certificate`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/certificates/<resource_id>` |
|Input| [Resource object `Certificate`](#object-Certificate) to update |
|Output| [Resource object `Certificate`](#object-Certificate) after delete |
### Delete `Certificate`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/certificates/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Certificate`](#object-Certificate) after delete |
---

## Object Schema Name: `BasicAuth`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `BasicAuth` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/basicAuths`|
|Input| `null` |
|Output| List of [resource object `BasicAuth`](#object-BasicAuth) |
### Get `BasicAuth`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/basicAuths/<resource_id>` |
|Input| `null`|
|Output| [Resource object `BasicAuth`](#object-BasicAuth) |
### Create `BasicAuth`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/basicAuths` |
|Input| [Resource object `BasicAuth`](#object-BasicAuth) |
|Output| [Resource object `BasicAuth`](#object-BasicAuth) after create |
### Update `BasicAuth`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/basicAuths/<resource_id>` |
|Input| [Resource object `BasicAuth`](#object-BasicAuth) to update |
|Output| [Resource object `BasicAuth`](#object-BasicAuth) after delete |
### Delete `BasicAuth`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/basicAuths/<resource_id>` |
|Input| `null` |
|Output| [Resource object `BasicAuth`](#object-BasicAuth) after delete |
---

## Object Schema Name: `SSHAuth`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `SSHAuth` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/sshAuths`|
|Input| `null` |
|Output| List of [resource object `SSHAuth`](#object-SSHAuth) |
### Get `SSHAuth`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/sshAuths/<resource_id>` |
|Input| `null`|
|Output| [Resource object `SSHAuth`](#object-SSHAuth) |
### Create `SSHAuth`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/sshAuths` |
|Input| [Resource object `SSHAuth`](#object-SSHAuth) |
|Output| [Resource object `SSHAuth`](#object-SSHAuth) after create |
### Update `SSHAuth`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/sshAuths/<resource_id>` |
|Input| [Resource object `SSHAuth`](#object-SSHAuth) to update |
|Output| [Resource object `SSHAuth`](#object-SSHAuth) after delete |
### Delete `SSHAuth`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/sshAuths/<resource_id>` |
|Input| `null` |
|Output| [Resource object `SSHAuth`](#object-SSHAuth) after delete |
---

## Object Schema Name: `NamespacedSecret`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/core/v1 |
### Get `NamespacedSecret` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/namespacedSecrets`|
|Input| `null` |
|Output| List of [resource object `NamespacedSecret`](#object-NamespacedSecret) |
### Get `NamespacedSecret`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/namespacedSecrets/<resource_id>` |
|Input| `null`|
|Output| [Resource object `NamespacedSecret`](#object-NamespacedSecret) |
### Create `NamespacedSecret`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/namespacedSecrets` |
|Input| [Resource object `NamespacedSecret`](#object-NamespacedSecret) |
|Output| [Resource object `NamespacedSecret`](#object-NamespacedSecret) after create |
### Update `NamespacedSecret`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/namespacedSecrets/<resource_id>` |
|Input| [Resource object `NamespacedSecret`](#object-NamespacedSecret) to update |
|Output| [Resource object `NamespacedSecret`](#object-NamespacedSecret) after delete |
### Delete `NamespacedSecret`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/namespacedSecrets/<resource_id>` |
|Input| `null` |
|Output| [Resource object `NamespacedSecret`](#object-NamespacedSecret) after delete |
---

## Object Schema Name: `NamespacedServiceAccountToken`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `NamespacedServiceAccountToken` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/namespacedServiceAccountTokens`|
|Input| `null` |
|Output| List of [resource object `NamespacedServiceAccountToken`](#object-NamespacedServiceAccountToken) |
### Get `NamespacedServiceAccountToken`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/namespacedServiceAccountTokens/<resource_id>` |
|Input| `null`|
|Output| [Resource object `NamespacedServiceAccountToken`](#object-NamespacedServiceAccountToken) |
### Create `NamespacedServiceAccountToken`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/namespacedServiceAccountTokens` |
|Input| [Resource object `NamespacedServiceAccountToken`](#object-NamespacedServiceAccountToken) |
|Output| [Resource object `NamespacedServiceAccountToken`](#object-NamespacedServiceAccountToken) after create |
### Update `NamespacedServiceAccountToken`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/namespacedServiceAccountTokens/<resource_id>` |
|Input| [Resource object `NamespacedServiceAccountToken`](#object-NamespacedServiceAccountToken) to update |
|Output| [Resource object `NamespacedServiceAccountToken`](#object-NamespacedServiceAccountToken) after delete |
### Delete `NamespacedServiceAccountToken`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/namespacedServiceAccountTokens/<resource_id>` |
|Input| `null` |
|Output| [Resource object `NamespacedServiceAccountToken`](#object-NamespacedServiceAccountToken) after delete |
---

## Object Schema Name: `NamespacedDockerCredential`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `NamespacedDockerCredential` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/namespacedDockerCredentials`|
|Input| `null` |
|Output| List of [resource object `NamespacedDockerCredential`](#object-NamespacedDockerCredential) |
### Get `NamespacedDockerCredential`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/namespacedDockerCredentials/<resource_id>` |
|Input| `null`|
|Output| [Resource object `NamespacedDockerCredential`](#object-NamespacedDockerCredential) |
### Create `NamespacedDockerCredential`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/namespacedDockerCredentials` |
|Input| [Resource object `NamespacedDockerCredential`](#object-NamespacedDockerCredential) |
|Output| [Resource object `NamespacedDockerCredential`](#object-NamespacedDockerCredential) after create |
### Update `NamespacedDockerCredential`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/namespacedDockerCredentials/<resource_id>` |
|Input| [Resource object `NamespacedDockerCredential`](#object-NamespacedDockerCredential) to update |
|Output| [Resource object `NamespacedDockerCredential`](#object-NamespacedDockerCredential) after delete |
### Delete `NamespacedDockerCredential`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/namespacedDockerCredentials/<resource_id>` |
|Input| `null` |
|Output| [Resource object `NamespacedDockerCredential`](#object-NamespacedDockerCredential) after delete |
---

## Object Schema Name: `NamespacedCertificate`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `NamespacedCertificate` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/namespacedCertificates`|
|Input| `null` |
|Output| List of [resource object `NamespacedCertificate`](#object-NamespacedCertificate) |
### Get `NamespacedCertificate`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/namespacedCertificates/<resource_id>` |
|Input| `null`|
|Output| [Resource object `NamespacedCertificate`](#object-NamespacedCertificate) |
### Create `NamespacedCertificate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/namespacedCertificates` |
|Input| [Resource object `NamespacedCertificate`](#object-NamespacedCertificate) |
|Output| [Resource object `NamespacedCertificate`](#object-NamespacedCertificate) after create |
### Update `NamespacedCertificate`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/namespacedCertificates/<resource_id>` |
|Input| [Resource object `NamespacedCertificate`](#object-NamespacedCertificate) to update |
|Output| [Resource object `NamespacedCertificate`](#object-NamespacedCertificate) after delete |
### Delete `NamespacedCertificate`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/namespacedCertificates/<resource_id>` |
|Input| `null` |
|Output| [Resource object `NamespacedCertificate`](#object-NamespacedCertificate) after delete |
---

## Object Schema Name: `NamespacedBasicAuth`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `NamespacedBasicAuth` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/namespacedBasicAuths`|
|Input| `null` |
|Output| List of [resource object `NamespacedBasicAuth`](#object-NamespacedBasicAuth) |
### Get `NamespacedBasicAuth`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/namespacedBasicAuths/<resource_id>` |
|Input| `null`|
|Output| [Resource object `NamespacedBasicAuth`](#object-NamespacedBasicAuth) |
### Create `NamespacedBasicAuth`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/namespacedBasicAuths` |
|Input| [Resource object `NamespacedBasicAuth`](#object-NamespacedBasicAuth) |
|Output| [Resource object `NamespacedBasicAuth`](#object-NamespacedBasicAuth) after create |
### Update `NamespacedBasicAuth`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/namespacedBasicAuths/<resource_id>` |
|Input| [Resource object `NamespacedBasicAuth`](#object-NamespacedBasicAuth) to update |
|Output| [Resource object `NamespacedBasicAuth`](#object-NamespacedBasicAuth) after delete |
### Delete `NamespacedBasicAuth`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/namespacedBasicAuths/<resource_id>` |
|Input| `null` |
|Output| [Resource object `NamespacedBasicAuth`](#object-NamespacedBasicAuth) after delete |
---

## Object Schema Name: `NamespacedSSHAuth`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `NamespacedSSHAuth` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/namespacedSshAuths`|
|Input| `null` |
|Output| List of [resource object `NamespacedSSHAuth`](#object-NamespacedSSHAuth) |
### Get `NamespacedSSHAuth`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/namespacedSshAuths/<resource_id>` |
|Input| `null`|
|Output| [Resource object `NamespacedSSHAuth`](#object-NamespacedSSHAuth) |
### Create `NamespacedSSHAuth`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/namespacedSshAuths` |
|Input| [Resource object `NamespacedSSHAuth`](#object-NamespacedSSHAuth) |
|Output| [Resource object `NamespacedSSHAuth`](#object-NamespacedSSHAuth) after create |
### Update `NamespacedSSHAuth`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/namespacedSshAuths/<resource_id>` |
|Input| [Resource object `NamespacedSSHAuth`](#object-NamespacedSSHAuth) to update |
|Output| [Resource object `NamespacedSSHAuth`](#object-NamespacedSSHAuth) after delete |
### Delete `NamespacedSSHAuth`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/namespacedSshAuths/<resource_id>` |
|Input| `null` |
|Output| [Resource object `NamespacedSSHAuth`](#object-NamespacedSSHAuth) after delete |
---

## Object Schema Name: `Service`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/core/v1 |
### Get `Service` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/services`|
|Input| `null` |
|Output| List of [resource object `Service`](#object-Service) |
### Get `Service`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/services/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Service`](#object-Service) |
### Create `Service`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/services` |
|Input| [Resource object `Service`](#object-Service) |
|Output| [Resource object `Service`](#object-Service) after create |
### Update `Service`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/services/<resource_id>` |
|Input| [Resource object `Service`](#object-Service) to update |
|Output| [Resource object `Service`](#object-Service) after delete |
### Delete `Service`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/services/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Service`](#object-Service) after delete |
---

## Object Schema Name: `DNSRecord`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/core/v1 |
### Get `DNSRecord` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/dnsRecords`|
|Input| `null` |
|Output| List of [resource object `DNSRecord`](#object-DNSRecord) |
### Get `DNSRecord`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/dnsRecords/<resource_id>` |
|Input| `null`|
|Output| [Resource object `DNSRecord`](#object-DNSRecord) |
### Create `DNSRecord`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/dnsRecords` |
|Input| [Resource object `DNSRecord`](#object-DNSRecord) |
|Output| [Resource object `DNSRecord`](#object-DNSRecord) after create |
### Update `DNSRecord`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/dnsRecords/<resource_id>` |
|Input| [Resource object `DNSRecord`](#object-DNSRecord) to update |
|Output| [Resource object `DNSRecord`](#object-DNSRecord) after delete |
### Delete `DNSRecord`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/dnsRecords/<resource_id>` |
|Input| `null` |
|Output| [Resource object `DNSRecord`](#object-DNSRecord) after delete |
---

## Object Schema Name: `Pod`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/core/v1 |
### Get `Pod` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/pods`|
|Input| `null` |
|Output| List of [resource object `Pod`](#object-Pod) |
### Get `Pod`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/pods/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Pod`](#object-Pod) |
### Create `Pod`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/pods` |
|Input| [Resource object `Pod`](#object-Pod) |
|Output| [Resource object `Pod`](#object-Pod) after create |
### Update `Pod`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/pods/<resource_id>` |
|Input| [Resource object `Pod`](#object-Pod) to update |
|Output| [Resource object `Pod`](#object-Pod) after delete |
### Delete `Pod`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/pods/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Pod`](#object-Pod) after delete |
---

## Object Schema Name: `Deployment`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/apps/v1beta2 |
### Get `Deployment` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/deployments`|
|Input| `null` |
|Output| List of [resource object `Deployment`](#object-Deployment) |
### Get `Deployment`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/deployments/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Deployment`](#object-Deployment) |
### Create `Deployment`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/deployments` |
|Input| [Resource object `Deployment`](#object-Deployment) |
|Output| [Resource object `Deployment`](#object-Deployment) after create |
### Update `Deployment`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/deployments/<resource_id>` |
|Input| [Resource object `Deployment`](#object-Deployment) to update |
|Output| [Resource object `Deployment`](#object-Deployment) after delete |
### Delete `Deployment`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/deployments/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Deployment`](#object-Deployment) after delete |
### Resource Action `pause`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/deployments/<resource_id>/pause` |
|Input| `null` |
|Output| `null` |
### Resource Action `resume`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/deployments/<resource_id>/resume` |
|Input| `null` |
|Output| `null` |
### Resource Action `rollback`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/deployments/<resource_id>/rollback` |
|Input| [object `deploymentRollbackInput`](#object-deploymentRollbackInput) |
|Output| `null` |
---

## Object Schema Name: `ReplicationController`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/core/v1 |
### Get `ReplicationController` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/replicationControllers`|
|Input| `null` |
|Output| List of [resource object `ReplicationController`](#object-ReplicationController) |
### Get `ReplicationController`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/replicationControllers/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ReplicationController`](#object-ReplicationController) |
---

## Object Schema Name: `ReplicaSet`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/extensions/v1beta1 |
### Get `ReplicaSet` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/replicaSets`|
|Input| `null` |
|Output| List of [resource object `ReplicaSet`](#object-ReplicaSet) |
### Get `ReplicaSet`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/replicaSets/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ReplicaSet`](#object-ReplicaSet) |
### Create `ReplicaSet`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/replicaSets` |
|Input| [Resource object `ReplicaSet`](#object-ReplicaSet) |
|Output| [Resource object `ReplicaSet`](#object-ReplicaSet) after create |
### Update `ReplicaSet`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/replicaSets/<resource_id>` |
|Input| [Resource object `ReplicaSet`](#object-ReplicaSet) to update |
|Output| [Resource object `ReplicaSet`](#object-ReplicaSet) after delete |
### Delete `ReplicaSet`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/replicaSets/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ReplicaSet`](#object-ReplicaSet) after delete |
---

## Object Schema Name: `StatefulSet`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/apps/v1beta2 |
### Get `StatefulSet` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/statefulSets`|
|Input| `null` |
|Output| List of [resource object `StatefulSet`](#object-StatefulSet) |
### Get `StatefulSet`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/statefulSets/<resource_id>` |
|Input| `null`|
|Output| [Resource object `StatefulSet`](#object-StatefulSet) |
### Create `StatefulSet`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/statefulSets` |
|Input| [Resource object `StatefulSet`](#object-StatefulSet) |
|Output| [Resource object `StatefulSet`](#object-StatefulSet) after create |
### Update `StatefulSet`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/statefulSets/<resource_id>` |
|Input| [Resource object `StatefulSet`](#object-StatefulSet) to update |
|Output| [Resource object `StatefulSet`](#object-StatefulSet) after delete |
### Delete `StatefulSet`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/statefulSets/<resource_id>` |
|Input| `null` |
|Output| [Resource object `StatefulSet`](#object-StatefulSet) after delete |
---

## Object Schema Name: `DaemonSet`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/apps/v1beta2 |
### Get `DaemonSet` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/daemonSets`|
|Input| `null` |
|Output| List of [resource object `DaemonSet`](#object-DaemonSet) |
### Get `DaemonSet`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/daemonSets/<resource_id>` |
|Input| `null`|
|Output| [Resource object `DaemonSet`](#object-DaemonSet) |
### Create `DaemonSet`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/daemonSets` |
|Input| [Resource object `DaemonSet`](#object-DaemonSet) |
|Output| [Resource object `DaemonSet`](#object-DaemonSet) after create |
### Update `DaemonSet`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/daemonSets/<resource_id>` |
|Input| [Resource object `DaemonSet`](#object-DaemonSet) to update |
|Output| [Resource object `DaemonSet`](#object-DaemonSet) after delete |
### Delete `DaemonSet`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/daemonSets/<resource_id>` |
|Input| `null` |
|Output| [Resource object `DaemonSet`](#object-DaemonSet) after delete |
---

## Object Schema Name: `Job`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/batch/v1 |
### Get `Job` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/jobs`|
|Input| `null` |
|Output| List of [resource object `Job`](#object-Job) |
### Get `Job`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/jobs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Job`](#object-Job) |
### Create `Job`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/jobs` |
|Input| [Resource object `Job`](#object-Job) |
|Output| [Resource object `Job`](#object-Job) after create |
### Update `Job`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/jobs/<resource_id>` |
|Input| [Resource object `Job`](#object-Job) to update |
|Output| [Resource object `Job`](#object-Job) after delete |
### Delete `Job`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/jobs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Job`](#object-Job) after delete |
---

## Object Schema Name: `CronJob`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/k8s.io/api/batch/v1beta1 |
### Get `CronJob` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/cronJobs`|
|Input| `null` |
|Output| List of [resource object `CronJob`](#object-CronJob) |
### Get `CronJob`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/cronJobs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `CronJob`](#object-CronJob) |
### Create `CronJob`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/cronJobs` |
|Input| [Resource object `CronJob`](#object-CronJob) |
|Output| [Resource object `CronJob`](#object-CronJob) after create |
### Update `CronJob`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/cronJobs/<resource_id>` |
|Input| [Resource object `CronJob`](#object-CronJob) to update |
|Output| [Resource object `CronJob`](#object-CronJob) after delete |
### Delete `CronJob`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/cronJobs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `CronJob`](#object-CronJob) after delete |
---

## Object Schema Name: `Workload`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `Workload` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/workloads`|
|Input| `null` |
|Output| List of [resource object `Workload`](#object-Workload) |
### Get `Workload`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/workloads/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Workload`](#object-Workload) |
### Create `Workload`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/workloads` |
|Input| [Resource object `Workload`](#object-Workload) |
|Output| [Resource object `Workload`](#object-Workload) after create |
### Update `Workload`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/workloads/<resource_id>` |
|Input| [Resource object `Workload`](#object-Workload) to update |
|Output| [Resource object `Workload`](#object-Workload) after delete |
### Delete `Workload`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/workloads/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Workload`](#object-Workload) after delete |
### Resource Action `pause`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/workloads/<resource_id>/pause` |
|Input| `null` |
|Output| `null` |
### Resource Action `resume`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/workloads/<resource_id>/resume` |
|Input| `null` |
|Output| `null` |
### Resource Action `rollback`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/workloads/<resource_id>/rollback` |
|Input| [object `rollbackRevision`](#object-rollbackRevision) |
|Output| `null` |
---

## Object Schema Name: `App`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `App` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/apps`|
|Input| `null` |
|Output| List of [resource object `App`](#object-App) |
### Get `App`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/apps/<resource_id>` |
|Input| `null`|
|Output| [Resource object `App`](#object-App) |
### Create `App`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/apps` |
|Input| [Resource object `App`](#object-App) |
|Output| [Resource object `App`](#object-App) after create |
### Update `App`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/apps/<resource_id>` |
|Input| [Resource object `App`](#object-App) to update |
|Output| [Resource object `App`](#object-App) after delete |
### Delete `App`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/apps/<resource_id>` |
|Input| `null` |
|Output| [Resource object `App`](#object-App) after delete |
### Resource Action `rollback`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/apps/<resource_id>/rollback` |
|Input| [object `rollbackRevision`](#object-rollbackRevision) |
|Output| `null` |
### Resource Action `upgrade`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/apps/<resource_id>/upgrade` |
|Input| [object `appUpgradeConfig`](#object-appUpgradeConfig) |
|Output| `null` |
---

## Object Schema Name: `AppRevision`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `AppRevision` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/appRevisions`|
|Input| `null` |
|Output| List of [resource object `AppRevision`](#object-AppRevision) |
### Get `AppRevision`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/appRevisions/<resource_id>` |
|Input| `null`|
|Output| [Resource object `AppRevision`](#object-AppRevision) |
### Create `AppRevision`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/appRevisions` |
|Input| [Resource object `AppRevision`](#object-AppRevision) |
|Output| [Resource object `AppRevision`](#object-AppRevision) after create |
### Update `AppRevision`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/appRevisions/<resource_id>` |
|Input| [Resource object `AppRevision`](#object-AppRevision) to update |
|Output| [Resource object `AppRevision`](#object-AppRevision) after delete |
### Delete `AppRevision`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/appRevisions/<resource_id>` |
|Input| `null` |
|Output| [Resource object `AppRevision`](#object-AppRevision) after delete |
---

## Object Schema Name: `SourceCodeProvider`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `false` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `SourceCodeProvider` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/sourceCodeProviders`|
|Input| `null` |
|Output| List of [resource object `SourceCodeProvider`](#object-SourceCodeProvider) |
### Get `SourceCodeProvider`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/sourceCodeProviders/<resource_id>` |
|Input| `null`|
|Output| [Resource object `SourceCodeProvider`](#object-SourceCodeProvider) |
### Update `SourceCodeProvider`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/sourceCodeProviders/<resource_id>` |
|Input| [Resource object `SourceCodeProvider`](#object-SourceCodeProvider) to update |
|Output| [Resource object `SourceCodeProvider`](#object-SourceCodeProvider) after delete |
### Delete `SourceCodeProvider`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/sourceCodeProviders/<resource_id>` |
|Input| `null` |
|Output| [Resource object `SourceCodeProvider`](#object-SourceCodeProvider) after delete |
---

## Object Schema Name: `SourceCodeProviderConfig`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `SourceCodeProviderConfig` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/sourceCodeProviderConfigs`|
|Input| `null` |
|Output| List of [resource object `SourceCodeProviderConfig`](#object-SourceCodeProviderConfig) |
### Get `SourceCodeProviderConfig`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/sourceCodeProviderConfigs/<resource_id>` |
|Input| `null`|
|Output| [Resource object `SourceCodeProviderConfig`](#object-SourceCodeProviderConfig) |
### Update `SourceCodeProviderConfig`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/sourceCodeProviderConfigs/<resource_id>` |
|Input| [Resource object `SourceCodeProviderConfig`](#object-SourceCodeProviderConfig) to update |
|Output| [Resource object `SourceCodeProviderConfig`](#object-SourceCodeProviderConfig) after delete |
### Delete `SourceCodeProviderConfig`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/sourceCodeProviderConfigs/<resource_id>` |
|Input| `null` |
|Output| [Resource object `SourceCodeProviderConfig`](#object-SourceCodeProviderConfig) after delete |
---

## Object Schema Name: `SourceCodeCredential`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `SourceCodeCredential` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/sourceCodeCredentials`|
|Input| `null` |
|Output| List of [resource object `SourceCodeCredential`](#object-SourceCodeCredential) |
### Get `SourceCodeCredential`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/sourceCodeCredentials/<resource_id>` |
|Input| `null`|
|Output| [Resource object `SourceCodeCredential`](#object-SourceCodeCredential) |
### Create `SourceCodeCredential`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/sourceCodeCredentials` |
|Input| [Resource object `SourceCodeCredential`](#object-SourceCodeCredential) |
|Output| [Resource object `SourceCodeCredential`](#object-SourceCodeCredential) after create |
### Delete `SourceCodeCredential`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/sourceCodeCredentials/<resource_id>` |
|Input| `null` |
|Output| [Resource object `SourceCodeCredential`](#object-SourceCodeCredential) after delete |
### Resource Action `logout`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/sourceCodeCredentials/<resource_id>/logout` |
|Input| `null` |
|Output| `null` |
### Resource Action `refreshrepos`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/sourceCodeCredentials/<resource_id>/refreshrepos` |
|Input| `null` |
|Output| `null` |
---

## Object Schema Name: `Pipeline`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `Pipeline` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/pipelines`|
|Input| `null` |
|Output| List of [resource object `Pipeline`](#object-Pipeline) |
### Get `Pipeline`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/pipelines/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Pipeline`](#object-Pipeline) |
### Create `Pipeline`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/pipelines` |
|Input| [Resource object `Pipeline`](#object-Pipeline) |
|Output| [Resource object `Pipeline`](#object-Pipeline) after create |
### Update `Pipeline`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/pipelines/<resource_id>` |
|Input| [Resource object `Pipeline`](#object-Pipeline) to update |
|Output| [Resource object `Pipeline`](#object-Pipeline) after delete |
### Delete `Pipeline`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/pipelines/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Pipeline`](#object-Pipeline) after delete |
### Resource Action `activate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/pipelines/<resource_id>/activate` |
|Input| `null` |
|Output| `null` |
### Resource Action `deactivate`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/pipelines/<resource_id>/deactivate` |
|Input| `null` |
|Output| `null` |
### Resource Action `pushconfig`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/pipelines/<resource_id>/pushconfig` |
|Input| [object `pushPipelineConfigInput`](#object-pushPipelineConfigInput) |
|Output| `null` |
### Resource Action `run`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/pipelines/<resource_id>/run` |
|Input| [object `runPipelineInput`](#object-runPipelineInput) |
|Output| `null` |
---

## Object Schema Name: `PipelineExecution`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `PipelineExecution` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/pipelineExecutions`|
|Input| `null` |
|Output| List of [resource object `PipelineExecution`](#object-PipelineExecution) |
### Get `PipelineExecution`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/pipelineExecutions/<resource_id>` |
|Input| `null`|
|Output| [Resource object `PipelineExecution`](#object-PipelineExecution) |
### Create `PipelineExecution`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/pipelineExecutions` |
|Input| [Resource object `PipelineExecution`](#object-PipelineExecution) |
|Output| [Resource object `PipelineExecution`](#object-PipelineExecution) after create |
### Update `PipelineExecution`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/pipelineExecutions/<resource_id>` |
|Input| [Resource object `PipelineExecution`](#object-PipelineExecution) to update |
|Output| [Resource object `PipelineExecution`](#object-PipelineExecution) after delete |
### Delete `PipelineExecution`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/pipelineExecutions/<resource_id>` |
|Input| `null` |
|Output| [Resource object `PipelineExecution`](#object-PipelineExecution) after delete |
### Resource Action `rerun`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/pipelineExecutions/<resource_id>/rerun` |
|Input| `null` |
|Output| `null` |
### Resource Action `stop`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/pipelineExecutions/<resource_id>/stop` |
|Input| `null` |
|Output| `null` |
---

## Object Schema Name: `PipelineSetting`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `PipelineSetting` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/pipelineSettings`|
|Input| `null` |
|Output| List of [resource object `PipelineSetting`](#object-PipelineSetting) |
### Get `PipelineSetting`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/pipelineSettings/<resource_id>` |
|Input| `null`|
|Output| [Resource object `PipelineSetting`](#object-PipelineSetting) |
### Create `PipelineSetting`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/pipelineSettings` |
|Input| [Resource object `PipelineSetting`](#object-PipelineSetting) |
|Output| [Resource object `PipelineSetting`](#object-PipelineSetting) after create |
### Update `PipelineSetting`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/pipelineSettings/<resource_id>` |
|Input| [Resource object `PipelineSetting`](#object-PipelineSetting) to update |
|Output| [Resource object `PipelineSetting`](#object-PipelineSetting) after delete |
### Delete `PipelineSetting`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/pipelineSettings/<resource_id>` |
|Input| `null` |
|Output| [Resource object `PipelineSetting`](#object-PipelineSetting) after delete |
---

## Object Schema Name: `SourceCodeRepository`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/apis/project.cattle.io/v3 |
### Get `SourceCodeRepository` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/sourceCodeRepositories`|
|Input| `null` |
|Output| List of [resource object `SourceCodeRepository`](#object-SourceCodeRepository) |
### Get `SourceCodeRepository`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/sourceCodeRepositories/<resource_id>` |
|Input| `null`|
|Output| [Resource object `SourceCodeRepository`](#object-SourceCodeRepository) |
### Create `SourceCodeRepository`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/sourceCodeRepositories` |
|Input| [Resource object `SourceCodeRepository`](#object-SourceCodeRepository) |
|Output| [Resource object `SourceCodeRepository`](#object-SourceCodeRepository) after create |
### Delete `SourceCodeRepository`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/sourceCodeRepositories/<resource_id>` |
|Input| `null` |
|Output| [Resource object `SourceCodeRepository`](#object-SourceCodeRepository) after delete |
---

## Object Schema Name: `Prometheus`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/github.com/coreos/prometheus-operator/pkg/client/monitoring/v1 |
### Get `Prometheus` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/prometheuses`|
|Input| `null` |
|Output| List of [resource object `Prometheus`](#object-Prometheus) |
### Get `Prometheus`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/prometheuses/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Prometheus`](#object-Prometheus) |
### Create `Prometheus`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/prometheuses` |
|Input| [Resource object `Prometheus`](#object-Prometheus) |
|Output| [Resource object `Prometheus`](#object-Prometheus) after create |
### Update `Prometheus`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/prometheuses/<resource_id>` |
|Input| [Resource object `Prometheus`](#object-Prometheus) to update |
|Output| [Resource object `Prometheus`](#object-Prometheus) after delete |
### Delete `Prometheus`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/prometheuses/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Prometheus`](#object-Prometheus) after delete |
---

## Object Schema Name: `ServiceMonitor`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/github.com/coreos/prometheus-operator/pkg/client/monitoring/v1 |
### Get `ServiceMonitor` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/serviceMonitors`|
|Input| `null` |
|Output| List of [resource object `ServiceMonitor`](#object-ServiceMonitor) |
### Get `ServiceMonitor`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/serviceMonitors/<resource_id>` |
|Input| `null`|
|Output| [Resource object `ServiceMonitor`](#object-ServiceMonitor) |
### Create `ServiceMonitor`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/serviceMonitors` |
|Input| [Resource object `ServiceMonitor`](#object-ServiceMonitor) |
|Output| [Resource object `ServiceMonitor`](#object-ServiceMonitor) after create |
### Update `ServiceMonitor`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/serviceMonitors/<resource_id>` |
|Input| [Resource object `ServiceMonitor`](#object-ServiceMonitor) to update |
|Output| [Resource object `ServiceMonitor`](#object-ServiceMonitor) after delete |
### Delete `ServiceMonitor`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/serviceMonitors/<resource_id>` |
|Input| `null` |
|Output| [Resource object `ServiceMonitor`](#object-ServiceMonitor) after delete |
---

## Object Schema Name: `PrometheusRule`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/github.com/coreos/prometheus-operator/pkg/client/monitoring/v1 |
### Get `PrometheusRule` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/prometheusRules`|
|Input| `null` |
|Output| List of [resource object `PrometheusRule`](#object-PrometheusRule) |
### Get `PrometheusRule`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/prometheusRules/<resource_id>` |
|Input| `null`|
|Output| [Resource object `PrometheusRule`](#object-PrometheusRule) |
### Create `PrometheusRule`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/prometheusRules` |
|Input| [Resource object `PrometheusRule`](#object-PrometheusRule) |
|Output| [Resource object `PrometheusRule`](#object-PrometheusRule) after create |
### Update `PrometheusRule`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/prometheusRules/<resource_id>` |
|Input| [Resource object `PrometheusRule`](#object-PrometheusRule) to update |
|Output| [Resource object `PrometheusRule`](#object-PrometheusRule) after delete |
### Delete `PrometheusRule`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/prometheusRules/<resource_id>` |
|Input| `null` |
|Output| [Resource object `PrometheusRule`](#object-PrometheusRule) after delete |
---

## Object Schema Name: `Alertmanager`
### K8s Object Meta
|||
|---|---|
|API Group| `project.cattle.io` |
|API Vesion| `v3` |
|Storage Scope| User Cluster |
|Namespaced| `true` |
|Package Name | github.com/rancher/types/vendor/github.com/coreos/prometheus-operator/pkg/client/monitoring/v1 |
### Get `Alertmanager` List
|||
|---|---|
|Method| `Get`|
|Path|`/v3/projects/<project_id>/alertmanagers`|
|Input| `null` |
|Output| List of [resource object `Alertmanager`](#object-Alertmanager) |
### Get `Alertmanager`
|||
|---|---|
|Method| `Get` |
|Path| `/v3/projects/<project_id>/alertmanagers/<resource_id>` |
|Input| `null`|
|Output| [Resource object `Alertmanager`](#object-Alertmanager) |
### Create `Alertmanager`
|||
|---|---|
|Method| `Post` |
|Path| `/v3/projects/<project_id>/alertmanagers` |
|Input| [Resource object `Alertmanager`](#object-Alertmanager) |
|Output| [Resource object `Alertmanager`](#object-Alertmanager) after create |
### Update `Alertmanager`
|||
|---|---|
|Method| `Put` |
|Path| `/v3/projects/<project_id>/alertmanagers/<resource_id>` |
|Input| [Resource object `Alertmanager`](#object-Alertmanager) to update |
|Output| [Resource object `Alertmanager`](#object-Alertmanager) after delete |
### Delete `Alertmanager`
|||
|---|---|
|Method| `Delete` |
|Path| `/v3/projects/<project_id>/alertmanagers/<resource_id>` |
|Input| `null` |
|Output| [Resource object `Alertmanager`](#object-Alertmanager) after delete |
---


## Resource Object Reference: ##

### Object `Taint`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Effect|string|true|true|true|false|
|Key|string|true|true|true|false|
|TimeAdded|date|true|true|true|false|
|Value|string|true|true|true|false|

### Object `ConfigMapNodeConfigSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|KubeletConfigKey|string|true|true|true|false|
|Name|string|true|true|true|false|
|Namespace|string|true|true|true|false|
|ResourceVersion|string|true|true|true|false|
|UID|string|true|true|true|false|

### Object `NodeConfigSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ConfigMap|configMapNodeConfigSource|true|true|true|false|

### Object `InternalNodeSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|PodCidr|string|false|false|true|false|
|ProviderId|string|false|false|true|false|
|Taints|array[taint]|false|true|true|false|
|Unschedulable|boolean|false|true|false|false|

### Object `NodeCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastHeartbeatTime|date|true|true|true|false|
|LastTransitionTime|date|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `NodeAddress`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Address|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `DaemonEndpoint`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Port|int|true|true|false|false|

### Object `NodeDaemonEndpoints`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|KubeletEndpoint|daemonEndpoint|true|true|true|false|

### Object `NodeSystemInfo`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Architecture|string|true|true|true|false|
|BootID|string|true|true|true|false|
|ContainerRuntimeVersion|string|true|true|true|false|
|KernelVersion|string|true|true|true|false|
|KubeProxyVersion|string|true|true|true|false|
|KubeletVersion|string|true|true|true|false|
|MachineID|string|true|true|true|false|
|OperatingSystem|string|true|true|true|false|
|OSImage|string|true|true|true|false|
|SystemUUID|string|true|true|true|false|

### Object `ContainerImage`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Names|array[string]|true|true|true|false|
|SizeBytes|int|true|true|false|false|

### Object `AttachedVolume`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|

### Object `NodeConfigStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Active|nodeConfigSource|true|true|true|false|
|Assigned|nodeConfigSource|true|true|true|false|
|Error|string|true|true|true|false|
|LastKnownGood|nodeConfigSource|true|true|true|false|

### Object `CPUInfo`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Count|int|true|true|false|false|

### Object `MemoryInfo`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MemTotalKiB|int|true|true|false|false|

### Object `OSInfo`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DockerVersion|string|true|true|true|false|
|KernelVersion|string|true|true|true|false|
|OperatingSystem|string|true|true|true|false|

### Object `KubernetesInfo`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|KubeProxyVersion|string|true|true|true|false|
|KubeletVersion|string|true|true|true|false|

### Object `NodeInfo`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CPU|cpuInfo|true|true|true|false|
|Kubernetes|kubernetesInfo|true|true|true|false|
|Memory|memoryInfo|true|true|true|false|
|OS|osInfo|true|true|true|false|

### Object `InternalNodeStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Allocatable|map[string]|true|true|true|false|
|Capacity|map[string]|true|true|true|false|
|Config|nodeConfigStatus|true|true|true|false|
|ExternalIPAddress|string|true|true|true|false|
|Hostname|string|true|true|true|false|
|Info|nodeInfo|true|true|true|false|
|IPAddress|string|true|true|true|false|
|NodeConditions|array[nodeCondition]|true|true|true|false|
|VolumesAttached|map[attachedVolume]|true|true|true|false|
|VolumesInUse|array[string]|true|true|true|false|

### Object `PublicEndpoint`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Addresses|array[string]|false|false|true|false|
|AllNodes|boolean|false|false|false|false|
|Hostname|string|false|false|true|false|
|IngressID|reference[ingress]|false|false|true|false|
|NodeID|reference[/v3/schemas/node]|false|false|true|false|
|Path|string|false|false|true|false|
|PodID|reference[pod]|false|false|true|false|
|Port|int|false|false|false|false|
|Protocol|string|false|false|true|false|
|ServiceID|reference[service]|false|false|true|false|

### Object `OwnerReference`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|APIVersion|string|true|true|true|false|
|BlockOwnerDeletion|boolean|true|true|true|false|
|Controller|boolean|true|true|true|false|
|Kind|string|true|true|true|false|
|Name|string|true|true|true|false|
|UID|string|true|true|true|false|

### Object `Initializer`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|

### Object `ListMeta`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Continue|string|true|true|true|false|
|ResourceVersion|string|true|true|true|false|
|SelfLink|string|true|true|true|false|

### Object `StatusCause`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Field|string|true|true|true|false|
|Message|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `StatusDetails`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Causes|array[statusCause]|true|true|true|false|
|Group|string|true|true|true|false|
|Kind|string|true|true|true|false|
|Name|string|true|true|true|false|
|RetryAfterSeconds|int|true|true|false|false|
|UID|string|true|true|true|false|

### Object `Status`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|APIVersion|string|true|true|true|false|
|Code|int|true|true|false|false|
|Details|statusDetails|true|true|true|false|
|Kind|string|true|true|true|false|
|Message|string|true|true|true|false|
|ListMeta|listMeta|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|

### Object `Initializers`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Pending|array[initializer]|true|true|true|false|
|Result|status|true|true|true|false|

### Object `ObjectMeta`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|Finalizers|array[string]|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|Namespace|string|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|SelfLink|string|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `NodePoolSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterID|reference[cluster]|true|false|true|true|
|ControlPlane|boolean|true|true|false|false|
|DisplayName|string|true|true|true|false|
|Etcd|boolean|true|true|false|false|
|HostnamePrefix|string|true|true|false|true|
|NodeAnnotations|map[string]|true|true|true|false|
|NodeLabels|map[string]|true|true|true|false|
|NodeTemplateID|reference[nodeTemplate]|true|true|false|true|
|Quantity|int|true|true|false|true|
|Worker|boolean|true|true|false|false|

### Object `Condition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|string|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `NodePoolStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Conditions|array[condition]|false|false|true|false|

### Object `NodePool`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClusterID|reference[cluster]|true|false|true|true|
|ControlPlane|boolean|true|true|false|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DisplayName|string|true|true|true|false|
|Etcd|boolean|true|true|false|false|
|HostnamePrefix|string|true|true|false|true|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeAnnotations|map[string]|true|true|true|false|
|NodeLabels|map[string]|true|true|true|false|
|NodeTemplateID|reference[nodeTemplate]|true|true|false|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Quantity|int|true|true|false|true|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Status|nodePoolStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|
|Worker|boolean|true|true|false|false|

### Object `NodeDrainInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DeleteLocalData|boolean|true|true|false|false|
|Force|boolean|true|true|false|false|
|GracePeriod|int|true|true|false|false|
|IgnoreDaemonSets|boolean|true|true|false|false|
|Timeout|int|true|true|false|false|

### Object `CustomConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Address|string|true|true|true|false|
|DockerSocket|string|true|true|true|false|
|InternalAddress|string|true|true|true|false|
|Label|map[string]|true|true|true|false|
|SSHKey|password|true|true|true|false|
|User|string|true|true|true|false|

### Object `NodeSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ControlPlane|boolean|true|false|false|false|
|CurrentNodeAnnotations|map[string]|true|true|true|false|
|CurrentNodeLabels|map[string]|true|true|true|false|
|CustomConfig|customConfig|true|true|true|false|
|Description|string|true|true|true|false|
|DesiredNodeAnnotations|map[string]|true|true|true|false|
|DesiredNodeLabels|map[string]|true|true|true|false|
|DesiredNodeUnschedulable|string|true|true|true|false|
|DisplayName|string|true|true|true|false|
|Etcd|boolean|true|false|false|false|
|Imported|boolean|true|true|false|false|
|NodeDrainInput|nodeDrainInput|true|true|true|false|
|NodePoolID|reference[nodePool]|false|false|true|false|
|NodeTemplateID|reference[nodeTemplate]|true|false|true|false|
|PodCidr|string|false|false|true|false|
|ProviderId|string|false|false|true|false|
|RequestedHostname|hostname|true|false|true|true|
|Taints|array[taint]|false|true|true|false|
|Unschedulable|boolean|false|true|false|false|
|Worker|boolean|true|false|false|false|

### Object `NodeTemplateSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AuthCertificateAuthority|string|true|true|true|false|
|AuthKey|string|true|true|true|false|
|Description|string|true|true|true|false|
|DisplayName|string|true|true|true|false|
|DockerVersion|string|true|true|true|false|
|Driver|string|false|false|true|false|
|EngineEnv|map[string]|true|true|true|false|
|EngineInsecureRegistry|array[string]|true|true|true|false|
|EngineInstallURL|string|true|true|true|false|
|EngineLabel|map[string]|true|true|true|false|
|EngineOpt|map[string]|true|true|true|false|
|EngineRegistryMirror|array[string]|true|true|true|false|
|EngineStorageDriver|string|true|true|true|false|
|UseInternalIPAddress|boolean|true|false|false|false|

### Object `RKEConfigNode`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Address|string|true|true|true|false|
|DockerSocket|string|true|true|true|false|
|HostnameOverride|string|true|true|true|false|
|InternalAddress|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|NodeID|reference[node]|true|true|true|false|
|Port|string|true|true|true|false|
|Role|array[enum]|true|true|true|false|
|SSHAgentAuth|boolean|true|true|false|false|
|SSHKey|password|true|true|true|false|
|SSHKeyPath|string|true|true|true|false|
|User|string|true|true|true|false|

### Object `DockerInfo`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Architecture|string|true|true|true|false|
|CgroupDriver|string|true|true|true|false|
|Debug|boolean|true|true|false|false|
|DockerRootDir|string|true|true|true|false|
|Driver|string|true|true|true|false|
|ExperimentalBuild|boolean|true|true|false|false|
|HTTPProxy|string|true|true|true|false|
|HTTPSProxy|string|true|true|true|false|
|ID|string|true|true|true|false|
|IndexServerAddress|string|true|true|true|false|
|KernelVersion|string|true|true|true|false|
|Labels|array[string]|true|true|true|false|
|LoggingDriver|string|true|true|true|false|
|Name|string|true|true|true|false|
|NoProxy|string|true|true|true|false|
|OperatingSystem|string|true|true|true|false|
|OSType|string|true|true|true|false|
|ServerVersion|string|true|true|true|false|

### Object `NodeStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Allocatable|map[string]|false|false|true|false|
|Capacity|map[string]|false|false|true|false|
|Conditions|array[nodeCondition]|false|false|true|false|
|DockerInfo|dockerInfo|false|false|true|false|
|ExternalIPAddress|string|false|false|true|false|
|Hostname|string|false|false|true|false|
|Info|nodeInfo|false|false|true|false|
|IPAddress|string|false|false|true|false|
|Limits|map[string]|false|false|true|false|
|NodeAnnotations|map[string]|false|false|true|false|
|NodeLabels|map[string]|false|false|true|false|
|NodeName|string|false|false|true|false|
|NodeTaints|array[taint]|false|false|true|false|
|Requested|map[string]|false|false|true|false|
|NodeConfig|rkeConfigNode|false|false|true|false|
|VolumesAttached|map[attachedVolume]|false|false|true|false|
|VolumesInUse|array[string]|false|false|true|false|

### Object `Node`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Allocatable|map[string]|false|false|true|false|
|Annotations|map[string]|false|false|true|false|
|Capacity|map[string]|false|false|true|false|
|ClusterID|reference[cluster]|true|false|true|true|
|Conditions|array[nodeCondition]|false|false|true|false|
|ControlPlane|boolean|true|false|false|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|CustomConfig|customConfig|true|true|true|false|
|Description|string|true|true|true|false|
|DockerInfo|dockerInfo|false|false|true|false|
|Etcd|boolean|true|false|false|false|
|ExternalIPAddress|string|false|false|true|false|
|Hostname|string|false|false|true|false|
|Id|dnsLabel|false|false|true|false|
|Imported|boolean|true|true|false|false|
|Info|nodeInfo|false|false|true|false|
|IPAddress|string|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Limits|map[string]|false|false|true|false|
|Name|string|true|true|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeName|string|false|false|true|false|
|NodePoolID|reference[nodePool]|false|false|true|false|
|NodeTaints|array[taint]|false|false|true|false|
|NodeTemplateID|reference[nodeTemplate]|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PodCidr|string|false|false|true|false|
|ProviderId|string|false|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|Removed|date|false|false|true|false|
|Requested|map[string]|false|false|true|false|
|RequestedHostname|hostname|true|false|true|true|
|SshUser|string|false|false|true|false|
|State|string|false|false|false|false|
|Taints|array[taint]|false|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|Unschedulable|boolean|false|false|false|false|
|UUID|string|false|false|true|false|
|VolumesAttached|map[attachedVolume]|false|false|true|false|
|VolumesInUse|array[string]|false|false|true|false|
|Worker|boolean|true|false|false|false|

### Object `NodeDriverSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Active|boolean|true|true|false|false|
|Builtin|boolean|true|true|false|false|
|Checksum|string|true|true|true|false|
|Description|string|true|true|true|false|
|DisplayName|string|true|true|true|false|
|ExternalID|string|true|true|true|false|
|UIURL|string|true|true|true|false|
|URL|string|true|true|true|true|
|WhitelistDomains|array[string]|true|true|true|false|

### Object `NodeDriverStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Conditions|array[condition]|false|false|true|false|

### Object `NodeDriver`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Active|boolean|true|true|false|false|
|Annotations|map[string]|true|true|true|false|
|Builtin|boolean|true|true|false|false|
|Checksum|string|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|ExternalID|string|true|true|true|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|string|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Status|nodeDriverStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UIURL|string|true|true|true|false|
|URL|string|true|true|true|true|
|UUID|string|false|false|true|false|
|WhitelistDomains|array[string]|true|true|true|false|

### Object `NodeTemplateCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|string|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `NodeTemplateStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Conditions|array[nodeTemplateCondition]|false|false|true|false|

### Object `NodeTemplate`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|AuthCertificateAuthority|string|true|true|true|false|
|AuthKey|string|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|DockerVersion|string|true|true|true|false|
|Driver|string|false|false|true|false|
|EngineEnv|map[string]|true|true|true|false|
|EngineInsecureRegistry|array[string]|true|true|true|false|
|EngineInstallURL|string|true|true|true|false|
|EngineLabel|map[string]|true|true|true|false|
|EngineOpt|map[string]|true|true|true|false|
|EngineRegistryMirror|array[string]|true|true|true|false|
|EngineStorageDriver|string|true|true|true|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|string|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Status|nodeTemplateStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UseInternalIPAddress|boolean|true|false|false|false|
|UUID|string|false|false|true|false|

### Object `ProjectCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|string|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `MonitoringCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|string|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `MonitoringStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Conditions|array[monitoringCondition]|true|true|true|false|
|GrafanaEndpoint|string|true|true|true|false|

### Object `ProjectStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Conditions|array[projectCondition]|false|false|true|false|
|MonitoringStatus|monitoringStatus|false|false|true|false|
|PodSecurityPolicyTemplateName|string|false|false|true|false|

### Object `SetPodSecurityPolicyTemplateInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|PodSecurityPolicyTemplateName|reference[podSecurityPolicyTemplate]|true|true|true|true|

### Object `ImportYamlOutput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Message|string|true|true|true|false|

### Object `ResourceQuotaLimit`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ConfigMaps|string|true|true|true|false|
|LimitsCPU|string|true|true|true|false|
|LimitsMemory|string|true|true|true|false|
|PersistentVolumeClaims|string|true|true|true|false|
|Pods|string|true|true|true|false|
|ReplicationControllers|string|true|true|true|false|
|RequestsCPU|string|true|true|true|false|
|RequestsMemory|string|true|true|true|false|
|RequestsStorage|string|true|true|true|false|
|Secrets|string|true|true|true|false|
|Services|string|true|true|true|false|
|ServicesLoadBalancers|string|true|true|true|false|
|ServicesNodePorts|string|true|true|true|false|

### Object `ProjectResourceQuota`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Limit|resourceQuotaLimit|true|true|true|false|
|UsedLimit|resourceQuotaLimit|true|true|true|false|

### Object `NamespaceResourceQuota`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Limit|resourceQuotaLimit|true|true|true|false|

### Object `ProjectSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterID|reference[cluster]|true|true|true|true|
|Description|string|true|true|true|false|
|DisplayName|string|true|true|true|true|
|EnableProjectMonitoring|boolean|true|true|false|false|
|NamespaceDefaultResourceQuota|namespaceResourceQuota|true|true|true|false|
|ResourceQuota|projectResourceQuota|true|true|true|false|

### Object `Project`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClusterID|reference[cluster]|true|true|true|true|
|Conditions|array[projectCondition]|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|EnableProjectMonitoring|boolean|true|true|false|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|MonitoringStatus|monitoringStatus|false|false|true|false|
|Name|string|true|true|true|true|
|NamespaceDefaultResourceQuota|namespaceResourceQuota|true|true|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PodSecurityPolicyTemplateName|string|false|false|true|false|
|Removed|date|false|false|true|false|
|ResourceQuota|projectResourceQuota|true|true|true|false|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `PolicyRule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|APIGroups|array[string]|true|true|true|false|
|NonResourceURLs|array[string]|true|true|true|false|
|ResourceNames|array[string]|true|true|true|false|
|Resources|array[string]|true|true|true|false|
|Verbs|array[string]|true|true|true|false|

### Object `GlobalRole`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|false|true|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|string|true|true|true|true|
|NewUserDefault|boolean|true|true|false|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|Rules|array[policyRule]|true|false|true|false|
|UUID|string|false|false|true|false|

### Object `GlobalRoleBinding`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|GlobalRoleID|reference[globalRole]|true|true|true|true|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|UserID|reference[user]|true|true|true|true|
|UUID|string|false|false|true|false|

### Object `RoleTemplate`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Builtin|boolean|false|false|false|false|
|ClusterCreatorDefault|boolean|true|true|false|true|
|Context|string|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|External|boolean|true|true|false|false|
|Hidden|boolean|true|true|false|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Locked|boolean|true|true|false|false|
|Name|string|true|true|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectCreatorDefault|boolean|true|true|false|true|
|Removed|date|false|false|true|false|
|RoleTemplateIDs|array[reference[roleTemplate]]|true|true|true|false|
|Rules|array[policyRule]|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `HostPortRange`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Max|int|true|true|false|false|
|Min|int|true|true|false|false|

### Object `SELinuxOptions`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Level|string|true|true|true|false|
|Role|string|true|true|true|false|
|Type|string|true|true|true|false|
|User|string|true|true|true|false|

### Object `SELinuxStrategyOptions`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Rule|string|true|true|true|false|
|SELinuxOptions|seLinuxOptions|true|true|true|false|

### Object `IDRange`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Max|int|true|true|false|false|
|Min|int|true|true|false|false|

### Object `RunAsUserStrategyOptions`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Ranges|array[idRange]|true|true|true|false|
|Rule|string|true|true|true|false|

### Object `SupplementalGroupsStrategyOptions`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Ranges|array[idRange]|true|true|true|false|
|Rule|string|true|true|true|false|

### Object `FSGroupStrategyOptions`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Ranges|array[idRange]|true|true|true|false|
|Rule|string|true|true|true|false|

### Object `AllowedHostPath`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|PathPrefix|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|

### Object `AllowedFlexVolume`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Driver|string|true|true|true|false|

### Object `PodSecurityPolicySpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AllowPrivilegeEscalation|boolean|true|true|true|false|
|AllowedCapabilities|array[string]|true|true|true|false|
|AllowedFlexVolumes|array[allowedFlexVolume]|true|true|true|false|
|AllowedHostPaths|array[allowedHostPath]|true|true|true|false|
|AllowedProcMountTypes|array[string]|true|true|true|false|
|AllowedUnsafeSysctls|array[string]|true|true|true|false|
|DefaultAddCapabilities|array[string]|true|true|true|false|
|DefaultAllowPrivilegeEscalation|boolean|true|true|true|false|
|ForbiddenSysctls|array[string]|true|true|true|false|
|FSGroup|fsGroupStrategyOptions|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|HostPorts|array[hostPortRange]|true|true|true|false|
|Privileged|boolean|true|true|false|false|
|ReadOnlyRootFilesystem|boolean|true|true|false|false|
|RequiredDropCapabilities|array[string]|true|true|true|false|
|RunAsUser|runAsUserStrategyOptions|true|true|true|false|
|SELinux|seLinuxStrategyOptions|true|true|true|false|
|SupplementalGroups|supplementalGroupsStrategyOptions|true|true|true|false|
|Volumes|array[string]|true|true|true|false|

### Object `PodSecurityPolicyTemplate`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AllowPrivilegeEscalation|boolean|true|true|true|false|
|AllowedCapabilities|array[string]|true|true|true|false|
|AllowedFlexVolumes|array[allowedFlexVolume]|true|true|true|false|
|AllowedHostPaths|array[allowedHostPath]|true|true|true|false|
|AllowedProcMountTypes|array[string]|true|true|true|false|
|AllowedUnsafeSysctls|array[string]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DefaultAddCapabilities|array[string]|true|true|true|false|
|DefaultAllowPrivilegeEscalation|boolean|true|true|true|false|
|Description|string|true|true|true|false|
|ForbiddenSysctls|array[string]|true|true|true|false|
|FSGroup|fsGroupStrategyOptions|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|HostPorts|array[hostPortRange]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Privileged|boolean|true|true|false|false|
|ReadOnlyRootFilesystem|boolean|true|true|false|false|
|Removed|date|false|false|true|false|
|RequiredDropCapabilities|array[string]|true|true|true|false|
|RunAsUser|runAsUserStrategyOptions|true|true|true|false|
|SELinux|seLinuxStrategyOptions|true|true|true|false|
|SupplementalGroups|supplementalGroupsStrategyOptions|true|true|true|false|
|UUID|string|false|false|true|false|
|Volumes|array[string]|true|true|true|false|

### Object `PodSecurityPolicyTemplateProjectBinding`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PodSecurityPolicyTemplateName|reference[podSecurityPolicyTemplate]|true|true|true|true|
|Removed|date|false|false|true|false|
|TargetProjectName|reference[project]|true|true|true|true|
|UUID|string|false|false|true|false|

### Object `ClusterRoleTemplateBinding`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClusterID|reference[cluster]|true|true|true|true|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|GroupID|reference[group]|true|true|true|false|
|GroupPrincipalID|reference[principal]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|RoleTemplateID|reference[roleTemplate]|true|true|true|true|
|UserID|reference[user]|true|true|true|false|
|UserPrincipalID|reference[principal]|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `ProjectRoleTemplateBinding`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|GroupID|reference[group]|true|true|true|false|
|GroupPrincipalID|reference[principal]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|true|
|Removed|date|false|false|true|false|
|RoleTemplateID|reference[roleTemplate]|true|true|true|true|
|UserID|reference[user]|true|true|true|false|
|UserPrincipalID|reference[principal]|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `ImportedConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|KubeConfig|password|true|true|true|false|

### Object `ETCDService`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CACert|string|true|true|true|false|
|Cert|string|true|true|true|false|
|Creation|string|true|true|true|false|
|ExternalURLs|array[string]|true|true|true|false|
|ExtraArgs|map[string]|true|true|true|false|
|ExtraBinds|array[string]|true|true|true|false|
|ExtraEnv|array[string]|true|true|true|false|
|Image|string|true|true|true|false|
|Key|string|true|true|true|false|
|Path|string|true|true|true|false|
|Retention|string|true|true|true|false|
|Snapshot|boolean|true|true|true|false|

### Object `KubeAPIService`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ExtraArgs|map[string]|true|true|true|false|
|ExtraBinds|array[string]|true|true|true|false|
|ExtraEnv|array[string]|true|true|true|false|
|Image|string|true|true|true|false|
|PodSecurityPolicy|boolean|true|true|false|false|
|ServiceClusterIPRange|string|true|true|true|false|
|ServiceNodePortRange|string|true|true|true|false|

### Object `KubeControllerService`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterCIDR|string|true|true|true|false|
|ExtraArgs|map[string]|true|true|true|false|
|ExtraBinds|array[string]|true|true|true|false|
|ExtraEnv|array[string]|true|true|true|false|
|Image|string|true|true|true|false|
|ServiceClusterIPRange|string|true|true|true|false|

### Object `SchedulerService`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ExtraArgs|map[string]|true|true|true|false|
|ExtraBinds|array[string]|true|true|true|false|
|ExtraEnv|array[string]|true|true|true|false|
|Image|string|true|true|true|false|

### Object `KubeletService`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterDNSServer|string|true|true|true|false|
|ClusterDomain|string|true|true|true|false|
|ExtraArgs|map[string]|true|true|true|false|
|ExtraBinds|array[string]|true|true|true|false|
|ExtraEnv|array[string]|true|true|true|false|
|FailSwapOn|boolean|true|true|false|false|
|Image|string|true|true|true|false|
|InfraContainerImage|string|true|true|true|false|

### Object `KubeproxyService`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ExtraArgs|map[string]|true|true|true|false|
|ExtraBinds|array[string]|true|true|true|false|
|ExtraEnv|array[string]|true|true|true|false|
|Image|string|true|true|true|false|

### Object `RKEConfigServices`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Etcd|etcdService|true|true|true|false|
|KubeAPI|kubeAPIService|true|true|true|false|
|KubeController|kubeControllerService|true|true|true|false|
|Kubelet|kubeletService|true|true|true|false|
|Kubeproxy|kubeproxyService|true|true|true|false|
|Scheduler|schedulerService|true|true|true|false|

### Object `CalicoNetworkProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CloudProvider|string|true|true|true|false|

### Object `CanalNetworkProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Iface|string|true|true|true|false|

### Object `FlannelNetworkProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Iface|string|true|true|true|false|

### Object `WeaveNetworkProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Password|string|true|true|true|false|

### Object `NetworkConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CalicoNetworkProvider|calicoNetworkProvider|true|true|true|false|
|CanalNetworkProvider|canalNetworkProvider|true|true|true|false|
|FlannelNetworkProvider|flannelNetworkProvider|true|true|true|false|
|Options|map[string]|true|true|true|false|
|Plugin|string|true|true|true|false|
|WeaveNetworkProvider|weaveNetworkProvider|true|true|true|false|

### Object `AuthnConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Options|map[string]|true|true|true|false|
|SANs|array[string]|true|true|true|false|
|Strategy|string|true|true|true|false|

### Object `RKESystemImages`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Alpine|string|true|true|true|false|
|CalicoCNI|string|true|true|true|false|
|CalicoControllers|string|true|true|true|false|
|CalicoCtl|string|true|true|true|false|
|CalicoNode|string|true|true|true|false|
|CanalCNI|string|true|true|true|false|
|CanalFlannel|string|true|true|true|false|
|CanalNode|string|true|true|true|false|
|CertDownloader|string|true|true|true|false|
|DNSmasq|string|true|true|true|false|
|Etcd|string|true|true|true|false|
|Flannel|string|true|true|true|false|
|FlannelCNI|string|true|true|true|false|
|Ingress|string|true|true|true|false|
|IngressBackend|string|true|true|true|false|
|KubeDNS|string|true|true|true|false|
|KubeDNSAutoscaler|string|true|true|true|false|
|KubeDNSSidecar|string|true|true|true|false|
|Kubernetes|string|true|true|true|false|
|KubernetesServicesSidecar|string|true|true|true|false|
|MetricsServer|string|true|true|true|false|
|NginxProxy|string|true|true|true|false|
|PodInfraContainer|string|true|true|true|false|
|WeaveCNI|string|true|true|true|false|
|WeaveNode|string|true|true|true|false|

### Object `AuthzConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Mode|string|true|true|true|false|
|Options|map[string]|true|true|true|false|

### Object `PrivateRegistry`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|IsDefault|boolean|true|true|false|false|
|Password|password|true|true|true|false|
|URL|string|true|true|true|false|
|User|string|true|true|true|false|

### Object `IngressConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ExtraArgs|map[string]|true|true|true|false|
|NodeSelector|map[string]|true|true|true|false|
|Options|map[string]|true|true|true|false|
|Provider|string|true|true|true|false|

### Object `AWSCloudProvider`
--- 


### Object `AzureCloudProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AADClientCertPassword|string|true|true|true|false|
|AADClientCertPath|string|true|true|true|false|
|AADClientID|string|true|true|true|false|
|AADClientSecret|string|true|true|true|false|
|Cloud|string|true|true|true|false|
|CloudProviderBackoff|boolean|true|true|false|false|
|CloudProviderBackoffDuration|int|true|true|false|false|
|CloudProviderBackoffExponent|int|true|true|false|false|
|CloudProviderBackoffJitter|int|true|true|false|false|
|CloudProviderBackoffRetries|int|true|true|false|false|
|CloudProviderRateLimit|boolean|true|true|false|false|
|CloudProviderRateLimitBucket|int|true|true|false|false|
|CloudProviderRateLimitQPS|int|true|true|false|false|
|Location|string|true|true|true|false|
|MaximumLoadBalancerRuleCount|int|true|true|false|false|
|PrimaryAvailabilitySetName|string|true|true|true|false|
|PrimaryScaleSetName|string|true|true|true|false|
|ResourceGroup|string|true|true|true|false|
|RouteTableName|string|true|true|true|false|
|SecurityGroupName|string|true|true|true|false|
|SubnetName|string|true|true|true|false|
|SubscriptionID|string|true|true|true|false|
|TenantID|string|true|true|true|false|
|UseInstanceMetadata|boolean|true|true|false|false|
|UseManagedIdentityExtension|boolean|true|true|false|false|
|VMType|string|true|true|true|false|
|VnetName|string|true|true|true|false|
|VnetResourceGroup|string|true|true|true|false|

### Object `GlobalOpenstackOpts`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AuthURL|string|true|true|true|false|
|CAFile|string|true|true|true|false|
|DomainID|string|true|true|true|false|
|DomainName|string|true|true|true|false|
|Password|string|true|true|true|false|
|Region|string|true|true|true|false|
|TenantID|string|true|true|true|false|
|TenantName|string|true|true|true|false|
|TrustID|string|true|true|true|false|
|UserID|string|true|true|true|false|
|Username|string|true|true|true|false|

### Object `LoadBalancerOpenstackOpts`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CreateMonitor|boolean|true|true|false|false|
|FloatingNetworkID|string|true|true|true|false|
|LBMethod|string|true|true|true|false|
|LBProvider|string|true|true|true|false|
|LBVersion|string|true|true|true|false|
|ManageSecurityGroups|boolean|true|true|false|false|
|MonitorDelay|int|true|true|false|false|
|MonitorMaxRetries|int|true|true|false|false|
|MonitorTimeout|int|true|true|false|false|
|SubnetID|string|true|true|true|false|
|UseOctavia|boolean|true|true|false|false|

### Object `BlockStorageOpenstackOpts`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BSVersion|string|true|true|true|false|
|IgnoreVolumeAZ|boolean|true|true|false|false|
|TrustDevicePath|boolean|true|true|false|false|

### Object `RouteOpenstackOpts`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|RouterID|string|true|true|true|false|

### Object `MetadataOpenstackOpts`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|RequestTimeout|int|true|true|false|false|
|SearchOrder|string|true|true|true|false|

### Object `OpenstackCloudProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BlockStorage|blockStorageOpenstackOpts|true|true|true|false|
|Global|globalOpenstackOpts|true|true|true|false|
|LoadBalancer|loadBalancerOpenstackOpts|true|true|true|false|
|Metadata|metadataOpenstackOpts|true|true|true|false|
|Route|routeOpenstackOpts|true|true|true|false|

### Object `GlobalVsphereOpts`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Datacenter|string|true|true|true|false|
|Datacenters|string|true|true|true|false|
|DefaultDatastore|string|true|true|true|false|
|InsecureFlag|boolean|true|true|false|false|
|Password|string|true|true|true|false|
|VCenterPort|string|true|true|true|false|
|VCenterIP|string|true|true|true|false|
|RoundTripperCount|int|true|true|false|false|
|User|string|true|true|true|false|
|VMName|string|true|true|true|false|
|VMUUID|string|true|true|true|false|
|WorkingDir|string|true|true|true|false|

### Object `VirtualCenterConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Datacenters|string|true|true|true|false|
|Password|string|true|true|true|false|
|VCenterPort|string|true|true|true|false|
|RoundTripperCount|int|true|true|false|false|
|User|string|true|true|true|false|

### Object `NetworkVshpereOpts`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|PublicNetwork|string|true|true|true|false|

### Object `DiskVsphereOpts`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|SCSIControllerType|string|true|true|true|false|

### Object `WorkspaceVsphereOpts`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Datacenter|string|true|true|true|false|
|DefaultDatastore|string|true|true|true|false|
|Folder|string|true|true|true|false|
|ResourcePoolPath|string|true|true|true|false|
|VCenterIP|string|true|true|true|false|

### Object `VsphereCloudProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Disk|diskVsphereOpts|true|true|true|false|
|Global|globalVsphereOpts|true|true|true|false|
|Network|networkVshpereOpts|true|true|true|false|
|VirtualCenter|map[virtualCenterConfig]|true|true|true|false|
|Workspace|workspaceVsphereOpts|true|true|true|false|

### Object `CloudProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AWSCloudProvider|awsCloudProvider|true|true|true|false|
|AzureCloudProvider|azureCloudProvider|true|true|true|false|
|CustomCloudProvider|string|true|true|true|false|
|Name|string|true|true|true|false|
|OpenstackCloudProvider|openstackCloudProvider|true|true|true|false|
|VsphereCloudProvider|vsphereCloudProvider|true|true|true|false|

### Object `BastionHost`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Address|string|true|true|true|false|
|Port|string|true|true|true|false|
|SSHAgentAuth|boolean|true|true|false|false|
|SSHKey|password|true|true|true|false|
|SSHKeyPath|string|true|true|true|false|
|User|string|true|true|true|false|

### Object `MonitoringConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Options|map[string]|true|true|true|false|
|Provider|string|true|true|true|false|

### Object `RancherKubernetesEngineConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AddonJobTimeout|int|true|true|false|false|
|Addons|string|true|true|true|false|
|AddonsInclude|array[string]|true|true|true|false|
|Authentication|authnConfig|true|true|true|false|
|Authorization|authzConfig|true|true|true|false|
|BastionHost|bastionHost|true|true|true|false|
|CloudProvider|cloudProvider|true|true|true|false|
|ClusterName|string|true|true|true|false|
|IgnoreDockerVersion|boolean|true|true|false|false|
|Ingress|ingressConfig|true|true|true|false|
|Version|string|true|true|true|false|
|Monitoring|monitoringConfig|true|true|true|false|
|Network|networkConfig|true|true|true|false|
|Nodes|array[rkeConfigNode]|true|true|true|false|
|PrefixPath|string|true|true|true|false|
|PrivateRegistries|array[privateRegistry]|true|true|true|false|
|Services|rkeConfigServices|true|true|true|false|
|SSHAgentAuth|boolean|true|true|false|false|
|SSHKeyPath|string|true|true|true|false|

### Object `ClusterSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AmazonElasticContainerServiceConfig|map[json]|true|true|true|false|
|AzureKubernetesServiceConfig|map[json]|true|true|true|false|
|DefaultClusterRoleForProjectMembers|reference[roleTemplate]|true|true|true|false|
|DefaultPodSecurityPolicyTemplateID|reference[podSecurityPolicyTemplate]|true|true|true|false|
|Description|string|true|true|true|false|
|DesiredAgentImage|string|true|true|true|false|
|DisplayName|string|true|true|true|true|
|DockerRootDir|string|true|true|true|false|
|EnableClusterAlerting|boolean|true|true|false|false|
|EnableClusterMonitoring|boolean|true|true|false|false|
|EnableNetworkPolicy|boolean|true|true|true|false|
|GenericEngineConfig|map[json]|true|true|true|false|
|GoogleKubernetesEngineConfig|map[json]|true|true|true|false|
|ImportedConfig|importedConfig|false|false|true|false|
|Internal|boolean|false|false|false|false|
|RancherKubernetesEngineConfig|rancherKubernetesEngineConfig|true|true|true|false|

### Object `ClusterCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|string|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `ComponentCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Error|string|true|true|true|false|
|Message|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `ClusterComponentStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Conditions|array[componentCondition]|true|true|true|false|
|Name|string|true|true|true|false|

### Object `Info`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BuildDate|string|true|true|true|false|
|Compiler|string|true|true|true|false|
|GitCommit|string|true|true|true|false|
|GitTreeState|string|true|true|true|false|
|GitVersion|string|true|true|true|false|
|GoVersion|string|true|true|true|false|
|Major|string|true|true|true|false|
|Minor|string|true|true|true|false|
|Platform|string|true|true|true|false|

### Object `LoadBalancerCapabilities`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Enabled|boolean|true|true|false|false|
|HealthCheckSupported|boolean|true|true|false|false|
|ProtocolsSupported|array[string]|true|true|true|false|
|Provider|string|true|true|true|false|

### Object `IngressCapabilities`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CustomDefaultBackend|boolean|true|true|false|false|
|IngressProvider|string|true|true|true|false|

### Object `Capabilities`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|IngressCapabilities|array[ingressCapabilities]|true|true|true|false|
|LoadBalancerCapabilities|loadBalancerCapabilities|true|true|true|false|
|NodePoolScalingSupported|boolean|true|true|false|false|
|NodePortRange|string|true|true|true|false|

### Object `ClusterStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AgentImage|string|false|false|true|false|
|Allocatable|map[string]|false|false|true|false|
|APIEndpoint|string|false|false|true|false|
|AppliedEnableNetworkPolicy|boolean|false|false|false|false|
|AppliedPodSecurityPolicyTemplateName|string|false|false|true|false|
|AppliedSpec|clusterSpec|false|false|true|false|
|CACert|string|false|false|true|false|
|Capabilities|capabilities|false|false|true|false|
|Capacity|map[string]|false|false|true|false|
|ComponentStatuses|array[clusterComponentStatus]|false|false|true|false|
|Conditions|array[clusterCondition]|false|false|true|false|
|Driver|string|false|false|true|false|
|FailedSpec|clusterSpec|false|false|true|false|
|Limits|map[string]|false|false|true|false|
|MonitoringStatus|monitoringStatus|false|false|true|false|
|Requested|map[string]|false|false|true|false|
|Version|info|false|false|true|false|

### Object `Cluster`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AgentImage|string|false|false|true|false|
|Allocatable|map[string]|false|false|true|false|
|Annotations|map[string]|true|true|true|false|
|APIEndpoint|string|false|false|true|false|
|AppliedEnableNetworkPolicy|boolean|false|false|false|false|
|AppliedPodSecurityPolicyTemplateName|string|false|false|true|false|
|AppliedSpec|clusterSpec|false|false|true|false|
|CACert|string|false|false|true|false|
|Capabilities|capabilities|false|false|true|false|
|Capacity|map[string]|false|false|true|false|
|ComponentStatuses|array[clusterComponentStatus]|false|false|true|false|
|Conditions|array[clusterCondition]|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DefaultClusterRoleForProjectMembers|reference[roleTemplate]|true|true|true|false|
|DefaultPodSecurityPolicyTemplateID|reference[podSecurityPolicyTemplate]|true|true|true|false|
|Description|string|true|true|true|false|
|DesiredAgentImage|string|true|true|true|false|
|DockerRootDir|string|true|true|true|false|
|Driver|string|false|false|true|false|
|EnableClusterAlerting|boolean|true|true|false|false|
|EnableClusterMonitoring|boolean|true|true|false|false|
|EnableNetworkPolicy|boolean|true|true|true|false|
|FailedSpec|clusterSpec|false|false|true|false|
|Id|dnsLabel|false|false|true|false|
|ImportedConfig|importedConfig|false|false|true|false|
|Internal|boolean|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Limits|map[string]|false|false|true|false|
|MonitoringStatus|monitoringStatus|false|false|true|false|
|Name|dnsLabel|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|RancherKubernetesEngineConfig|rancherKubernetesEngineConfig|true|true|true|false|
|Removed|date|false|false|true|false|
|Requested|map[string]|false|false|true|false|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|
|Version|info|false|false|true|false|

### Object `ClusterRegistrationTokenSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterID|reference[cluster]|true|true|true|true|

### Object `ClusterRegistrationTokenStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Command|string|false|false|true|false|
|InsecureCommand|string|false|false|true|false|
|ManifestURL|string|false|false|true|false|
|NodeCommand|string|false|false|true|false|
|Token|string|false|false|true|false|
|WindowsNodeCommand|string|false|false|true|false|

### Object `ClusterRegistrationToken`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClusterID|reference[cluster]|true|true|true|true|
|Command|string|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|InsecureCommand|string|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|ManifestURL|string|false|false|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeCommand|string|false|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Token|string|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|
|WindowsNodeCommand|string|false|false|true|false|

### Object `GenerateKubeConfigOutput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Config|string|true|true|true|false|

### Object `ImportClusterYamlInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DefaultNamespace|string|true|true|true|false|
|Namespace|string|true|true|true|false|
|ProjectID|reference[project]|true|true|true|false|
|YAML|string|true|true|true|false|

### Object `ExportOutput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|YAMLOutput|string|true|true|true|false|

### Object `MonitoringInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Answers|map[string]|true|true|true|false|

### Object `CatalogSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Branch|string|true|true|true|false|
|CatalogKind|string|true|true|true|false|
|Description|string|true|true|true|false|
|Password|password|true|true|true|false|
|URL|string|true|true|true|true|
|Username|string|true|true|true|false|

### Object `VersionCommits`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Value|map[string]|true|true|true|false|

### Object `CatalogCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|string|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `CatalogStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Commit|string|false|false|true|false|
|Conditions|array[catalogCondition]|false|false|true|false|
|HelmVersionCommits|map[versionCommits]|false|false|true|false|
|LastRefreshTimestamp|string|false|false|true|false|

### Object `Catalog`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Branch|string|true|true|true|false|
|Commit|string|false|false|true|false|
|Conditions|array[catalogCondition]|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Kind|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|LastRefreshTimestamp|string|false|false|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Password|password|true|true|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|URL|string|true|true|true|true|
|Username|string|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `SubQuestion`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Default|string|true|true|true|false|
|Description|string|true|true|true|false|
|Group|string|true|true|true|false|
|InvalidChars|string|true|true|true|false|
|Label|string|true|true|true|false|
|Max|int|true|true|false|false|
|MaxLength|int|true|true|false|false|
|Min|int|true|true|false|false|
|MinLength|int|true|true|false|false|
|Options|array[string]|true|true|true|false|
|Required|boolean|true|true|false|false|
|ShowIf|string|true|true|true|false|
|Type|string|true|true|true|false|
|ValidChars|string|true|true|true|false|
|Variable|string|true|true|true|false|

### Object `Question`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Default|string|true|true|true|false|
|Description|string|true|true|true|false|
|Group|string|true|true|true|false|
|InvalidChars|string|true|true|true|false|
|Label|string|true|true|true|false|
|Max|int|true|true|false|false|
|MaxLength|int|true|true|false|false|
|Min|int|true|true|false|false|
|MinLength|int|true|true|false|false|
|Options|array[string]|true|true|true|false|
|Required|boolean|true|true|false|false|
|ShowIf|string|true|true|true|false|
|ShowSubquestionIf|string|true|true|true|false|
|Subquestions|array[subQuestion]|true|true|true|false|
|Type|string|true|true|true|false|
|ValidChars|string|true|true|true|false|
|Variable|string|true|true|true|false|

### Object `TemplateVersionSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AppReadme|string|true|true|true|false|
|Digest|string|true|true|true|false|
|ExternalID|string|true|true|true|false|
|Files|map[string]|true|true|true|false|
|KubeVersion|string|true|true|true|false|
|Questions|array[question]|true|true|true|false|
|RancherVersion|string|true|true|true|false|
|Readme|string|true|true|true|false|
|RequiredNamespace|string|true|true|true|false|
|UpgradeVersionLinks|map[string]|true|true|true|false|
|Version|string|true|true|true|false|

### Object `TemplateSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CatalogID|reference[catalog]|true|true|true|false|
|Categories|array[string]|true|true|true|false|
|Category|string|true|true|true|false|
|ClusterCatalogID|reference[clusterCatalog]|true|true|true|false|
|ClusterID|reference[cluster]|true|true|true|true|
|DefaultTemplateVersionID|reference[templateVersion]|true|true|true|false|
|DefaultVersion|string|true|true|true|false|
|Description|string|true|true|true|false|
|DisplayName|string|true|true|true|false|
|FolderName|string|true|true|true|false|
|Icon|string|true|true|true|false|
|IconFilename|string|true|true|true|false|
|Maintainer|string|true|true|true|false|
|Path|string|true|true|true|false|
|ProjectCatalogID|reference[projectCatalog]|true|true|true|false|
|ProjectID|reference[project]|true|true|true|true|
|ProjectURL|string|true|true|true|false|
|Readme|string|true|true|true|false|
|UpgradeFrom|string|true|true|true|false|
|Versions|array[templateVersionSpec]|true|true|true|false|

### Object `TemplateStatus`
--- 


### Object `Template`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|CatalogID|reference[catalog]|true|true|true|false|
|Categories|array[string]|true|true|true|false|
|Category|string|true|true|true|false|
|ClusterCatalogID|reference[clusterCatalog]|true|true|true|false|
|ClusterID|reference[cluster]|true|true|true|true|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DefaultTemplateVersionID|reference[templateVersion]|true|true|true|false|
|DefaultVersion|string|true|true|true|false|
|Description|string|true|true|true|false|
|FolderName|string|true|true|true|false|
|Icon|string|true|true|true|false|
|IconFilename|string|true|true|true|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Maintainer|string|true|true|true|false|
|Name|string|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Path|string|true|true|true|false|
|ProjectCatalogID|reference[projectCatalog]|true|true|true|false|
|ProjectID|reference[project]|true|true|true|true|
|ProjectURL|string|true|true|true|false|
|Readme|string|true|true|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Status|templateStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UpgradeFrom|string|true|true|true|false|
|UUID|string|false|false|true|false|
|VersionLinks|map[string]|true|true|true|false|
|Versions|array[templateVersionSpec]|true|true|true|false|

### Object `TemplateVersionStatus`
--- 


### Object `TemplateVersion`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|AppReadme|string|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Digest|string|true|true|true|false|
|ExternalID|string|true|true|true|false|
|Files|map[string]|true|true|true|false|
|KubeVersion|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Questions|array[question]|true|true|true|false|
|RancherVersion|string|true|true|true|false|
|Readme|string|true|true|true|false|
|Removed|date|false|false|true|false|
|RequiredNamespace|string|true|true|true|false|
|State|string|false|false|false|false|
|Status|templateVersionStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UpgradeVersionLinks|map[string]|true|true|true|false|
|UUID|string|false|false|true|false|
|Version|string|true|true|true|false|

### Object `TemplateContent`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Data|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `Group`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|string|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `GroupMember`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|GroupID|reference[group]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PrincipalID|reference[principal]|true|true|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `Principal`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|ExtraInfo|map[string]|true|true|true|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|LoginName|string|true|true|true|false|
|Me|boolean|true|true|false|false|
|MemberOf|boolean|true|true|false|false|
|Name|string|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PrincipalType|string|true|true|true|false|
|ProfilePicture|string|true|true|true|false|
|ProfileURL|string|true|true|true|false|
|Provider|string|true|true|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `SearchPrincipalsInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|false|true|
|PrincipalType|enum|true|true|true|false|

### Object `ChangePasswordInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CurrentPassword|string|true|true|true|true|
|NewPassword|string|true|true|true|true|

### Object `SetPasswordInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|NewPassword|string|true|true|true|true|

### Object `UserSpec`
--- 


### Object `UserCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|string|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `UserStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Conditions|array[userCondition]|false|false|true|false|

### Object `User`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Conditions|array[userCondition]|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Enabled|boolean|true|true|true|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Me|boolean|true|true|false|false|
|MustChangePassword|boolean|true|true|false|false|
|Name|string|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Password|string|true|false|true|false|
|PrincipalIDs|array[reference[principal]]|true|true|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|Username|string|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `AuthConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessMode|enum|true|true|false|true|
|AllowedPrincipalIDs|array[reference[principal]]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|Type|string|true|false|true|false|
|UUID|string|false|false|true|false|

### Object `LocalConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessMode|enum|true|true|false|true|
|AllowedPrincipalIDs|array[reference[principal]]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|Type|string|true|false|true|false|
|UUID|string|false|false|true|false|

### Object `GithubConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessMode|enum|true|true|false|true|
|AllowedPrincipalIDs|array[reference[principal]]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|ClientID|string|true|true|true|true|
|ClientSecret|password|true|true|true|true|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|TLS|boolean|true|true|false|false|
|Type|string|true|false|true|false|
|UUID|string|false|false|true|false|

### Object `GithubConfigTestOutput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|RedirectURL|string|true|true|true|false|

### Object `GithubConfigApplyInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Code|string|true|true|true|false|
|Enabled|boolean|true|true|false|false|
|GithubConfig|githubConfig|true|true|true|false|

### Object `AzureADConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessMode|enum|true|true|false|true|
|AllowedPrincipalIDs|array[reference[principal]]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|ApplicationID|string|true|true|false|true|
|ApplicationSecret|password|true|true|false|true|
|AuthEndpoint|string|true|true|false|true|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|Endpoint|string|true|true|false|true|
|GraphEndpoint|string|true|true|false|true|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|RancherURL|string|true|true|false|true|
|Removed|date|false|false|true|false|
|TenantID|string|true|true|false|true|
|TokenEndpoint|string|true|true|false|true|
|Type|string|true|false|true|false|
|UUID|string|false|false|true|false|

### Object `AzureADConfigTestOutput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|RedirectURL|string|true|true|true|false|

### Object `AzureADConfigApplyInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Code|string|true|true|true|false|
|Config|azureADConfig|true|true|true|false|

### Object `ActiveDirectoryConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessMode|enum|true|true|false|true|
|AllowedPrincipalIDs|array[reference[principal]]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Certificate|string|true|true|true|false|
|ConnectionTimeout|int|true|true|false|true|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DefaultLoginDomain|string|true|true|true|false|
|Enabled|boolean|true|true|false|false|
|GroupDNAttribute|string|true|true|true|true|
|GroupMemberMappingAttribute|string|true|true|true|true|
|GroupMemberUserAttribute|string|true|true|true|true|
|GroupNameAttribute|string|true|true|true|true|
|GroupObjectClass|string|true|true|true|true|
|GroupSearchAttribute|string|true|true|true|true|
|GroupSearchBase|string|true|true|true|false|
|GroupSearchFilter|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NestedGroupMembershipEnabled|boolean|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Port|int|true|true|false|false|
|Removed|date|false|false|true|false|
|Servers|array[string]|true|true|true|true|
|ServiceAccountPassword|password|true|true|true|true|
|ServiceAccountUsername|string|true|true|true|true|
|TLS|boolean|true|true|false|false|
|Type|string|true|false|true|false|
|UserDisabledBitMask|int|true|true|false|false|
|UserEnabledAttribute|string|true|true|true|true|
|UserLoginAttribute|string|true|true|true|true|
|UserNameAttribute|string|true|true|true|true|
|UserObjectClass|string|true|true|true|true|
|UserSearchAttribute|string|true|true|true|true|
|UserSearchBase|string|true|true|true|true|
|UserSearchFilter|string|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `ActiveDirectoryTestAndApplyInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDirectoryConfig|activeDirectoryConfig|true|true|true|false|
|Enabled|boolean|true|true|false|false|
|Password|string|true|true|true|false|
|Username|string|true|true|true|false|

### Object `OpenLdapConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessMode|enum|true|true|false|true|
|AllowedPrincipalIDs|array[reference[principal]]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Certificate|string|true|true|true|false|
|ConnectionTimeout|int|true|true|false|true|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|GroupDNAttribute|string|true|true|false|false|
|GroupMemberMappingAttribute|string|true|true|false|true|
|GroupMemberUserAttribute|string|true|true|false|false|
|GroupNameAttribute|string|true|true|false|true|
|GroupObjectClass|string|true|true|false|true|
|GroupSearchAttribute|string|true|true|false|true|
|GroupSearchBase|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NestedGroupMembershipEnabled|boolean|true|true|false|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Port|int|true|true|false|true|
|Removed|date|false|false|true|false|
|Servers|array[string]|true|true|false|true|
|ServiceAccountDistinguishedName|string|true|true|true|true|
|ServiceAccountPassword|password|true|true|true|true|
|TLS|boolean|true|true|false|true|
|Type|string|true|false|true|false|
|UserDisabledBitMask|int|true|true|false|false|
|UserEnabledAttribute|string|true|true|true|false|
|UserLoginAttribute|string|true|true|false|true|
|UserMemberAttribute|string|true|true|false|true|
|UserNameAttribute|string|true|true|false|true|
|UserObjectClass|string|true|true|false|true|
|UserSearchAttribute|string|true|true|false|true|
|UserSearchBase|string|true|true|false|true|
|UUID|string|false|false|true|false|

### Object `LdapConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessMode|enum|true|true|false|true|
|AllowedPrincipalIDs|array[reference[principal]]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Certificate|string|true|true|true|false|
|ConnectionTimeout|int|true|true|false|true|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|GroupDNAttribute|string|true|true|false|false|
|GroupMemberMappingAttribute|string|true|true|false|true|
|GroupMemberUserAttribute|string|true|true|false|false|
|GroupNameAttribute|string|true|true|false|true|
|GroupObjectClass|string|true|true|false|true|
|GroupSearchAttribute|string|true|true|false|true|
|GroupSearchBase|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NestedGroupMembershipEnabled|boolean|true|true|false|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Port|int|true|true|false|true|
|Removed|date|false|false|true|false|
|Servers|array[string]|true|true|false|true|
|ServiceAccountDistinguishedName|string|true|true|true|true|
|ServiceAccountPassword|password|true|true|true|true|
|TLS|boolean|true|true|false|true|
|Type|string|true|false|true|false|
|UserDisabledBitMask|int|true|true|false|false|
|UserEnabledAttribute|string|true|true|true|false|
|UserLoginAttribute|string|true|true|false|true|
|UserMemberAttribute|string|true|true|false|true|
|UserNameAttribute|string|true|true|false|true|
|UserObjectClass|string|true|true|false|true|
|UserSearchAttribute|string|true|true|false|true|
|UserSearchBase|string|true|true|false|true|
|UUID|string|false|false|true|false|

### Object `OpenLdapTestAndApplyInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LdapConfig|ldapConfig|true|true|true|false|
|Password|password|true|true|true|true|
|Username|string|true|true|true|false|

### Object `FreeIpaConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessMode|enum|true|true|false|true|
|AllowedPrincipalIDs|array[reference[principal]]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Certificate|string|true|true|true|false|
|ConnectionTimeout|int|true|true|false|true|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|GroupDNAttribute|string|true|true|false|false|
|GroupMemberMappingAttribute|string|true|true|false|true|
|GroupMemberUserAttribute|string|true|true|false|false|
|GroupNameAttribute|string|true|true|false|true|
|GroupObjectClass|string|true|true|false|true|
|GroupSearchAttribute|string|true|true|false|true|
|GroupSearchBase|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Port|int|true|true|false|true|
|Removed|date|false|false|true|false|
|Servers|array[string]|true|true|false|true|
|ServiceAccountDistinguishedName|string|true|true|true|true|
|ServiceAccountPassword|password|true|true|true|true|
|TLS|boolean|true|true|false|true|
|Type|string|true|false|true|false|
|UserDisabledBitMask|int|true|true|false|false|
|UserEnabledAttribute|string|true|true|true|false|
|UserLoginAttribute|string|true|true|false|true|
|UserMemberAttribute|string|true|true|false|true|
|UserNameAttribute|string|true|true|false|true|
|UserObjectClass|string|true|true|false|true|
|UserSearchAttribute|string|true|true|false|true|
|UserSearchBase|string|true|true|false|true|
|UUID|string|false|false|true|false|

### Object `FreeIpaTestAndApplyInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LdapConfig|ldapConfig|true|true|true|false|
|Password|password|true|true|true|true|
|Username|string|true|true|true|false|

### Object `PingConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessMode|enum|true|true|false|true|
|AllowedPrincipalIDs|array[reference[principal]]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DisplayNameField|string|true|true|true|true|
|Enabled|boolean|true|true|false|false|
|GroupsField|string|true|true|true|true|
|IDPMetadataContent|string|true|true|true|true|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|RancherAPIHost|string|true|true|true|true|
|Removed|date|false|false|true|false|
|SpCert|string|true|true|true|true|
|SpKey|password|true|true|true|true|
|Type|string|true|false|true|false|
|UIDField|string|true|true|true|true|
|UserNameField|string|true|true|true|true|
|UUID|string|false|false|true|false|

### Object `ADFSConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessMode|enum|true|true|false|true|
|AllowedPrincipalIDs|array[reference[principal]]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DisplayNameField|string|true|true|true|true|
|Enabled|boolean|true|true|false|false|
|GroupsField|string|true|true|true|true|
|IDPMetadataContent|string|true|true|true|true|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|RancherAPIHost|string|true|true|true|true|
|Removed|date|false|false|true|false|
|SpCert|string|true|true|true|true|
|SpKey|password|true|true|true|true|
|Type|string|true|false|true|false|
|UIDField|string|true|true|true|true|
|UserNameField|string|true|true|true|true|
|UUID|string|false|false|true|false|

### Object `KeyCloakConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessMode|enum|true|true|false|true|
|AllowedPrincipalIDs|array[reference[principal]]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DisplayNameField|string|true|true|true|true|
|Enabled|boolean|true|true|false|false|
|GroupsField|string|true|true|true|true|
|IDPMetadataContent|string|true|true|true|true|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|RancherAPIHost|string|true|true|true|true|
|Removed|date|false|false|true|false|
|SpCert|string|true|true|true|true|
|SpKey|password|true|true|true|true|
|Type|string|true|false|true|false|
|UIDField|string|true|true|true|true|
|UserNameField|string|true|true|true|true|
|UUID|string|false|false|true|false|

### Object `SamlConfigTestInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FinalRedirectURL|string|true|true|true|false|

### Object `SamlConfigTestOutput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|IdpRedirectURL|string|true|true|true|false|

### Object `Token`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|AuthProvider|string|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Current|boolean|true|true|false|false|
|Description|string|true|true|true|false|
|Expired|boolean|true|true|false|false|
|ExpiresAt|string|true|true|true|false|
|GroupPrincipals|array[reference[principal]]|true|true|true|false|
|IsDerived|boolean|true|true|false|false|
|Labels|map[string]|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProviderInfo|map[string]|true|true|true|false|
|Removed|date|false|false|true|false|
|Token|string|true|false|true|false|
|TTLMillis|int|true|true|false|false|
|UserID|reference[user]|true|true|true|false|
|UserPrincipal|reference[principal]|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `Values`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BoolValue|boolean|true|true|false|false|
|IntValue|int|true|true|false|false|
|StringSliceValue|array[string]|true|true|true|false|
|StringValue|string|true|true|true|false|

### Object `Field`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Create|boolean|true|true|false|false|
|Default|values|true|true|true|false|
|Description|string|true|true|true|false|
|DynamicField|boolean|true|true|false|false|
|InvalidChars|string|true|true|true|false|
|Max|int|true|true|false|false|
|MaxLength|int|true|true|false|false|
|Min|int|true|true|false|false|
|MinLength|int|true|true|false|false|
|Nullable|boolean|true|true|false|false|
|Options|array[string]|true|true|true|false|
|Required|boolean|true|true|false|false|
|Type|string|true|true|true|false|
|Unique|boolean|true|true|false|false|
|Update|boolean|true|true|false|false|
|ValidChars|string|true|true|true|false|

### Object `Action`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Input|string|true|true|true|false|
|Output|string|true|true|true|false|

### Object `Filter`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Modifiers|array[string]|true|true|true|false|

### Object `DynamicSchemaSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CollectionActions|map[action]|true|true|true|false|
|CollectionFields|map[field]|true|true|true|false|
|CollectionFilters|map[filter]|true|true|true|false|
|CollectionMethods|array[string]|true|true|true|false|
|Embed|boolean|true|true|false|false|
|EmbedType|string|true|true|true|false|
|IncludeableLinks|array[string]|true|true|true|false|
|PluralName|string|true|true|true|false|
|ResourceActions|map[action]|true|true|true|false|
|ResourceFields|map[field]|true|true|true|false|
|ResourceMethods|array[string]|true|true|true|false|
|SchemaName|string|true|true|true|false|

### Object `DynamicSchemaStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Fake|string|false|false|true|false|

### Object `DynamicSchema`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|CollectionActions|map[action]|true|true|true|false|
|CollectionFields|map[field]|true|true|true|false|
|CollectionFilters|map[filter]|true|true|true|false|
|CollectionMethods|array[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Embed|boolean|true|true|false|false|
|EmbedType|string|true|true|true|false|
|IncludeableLinks|array[string]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PluralName|string|true|true|true|false|
|Removed|date|false|false|true|false|
|ResourceActions|map[action]|true|true|true|false|
|ResourceFields|map[field]|true|true|true|false|
|ResourceMethods|array[string]|true|true|true|false|
|SchemaName|string|true|true|true|false|
|State|string|false|false|false|false|
|Status|dynamicSchemaStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `Preference`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|
|Value|string|true|true|true|true|

### Object `UserAttribute`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|GroupPrincipals|map[principal]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|UserName|string|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `ProjectNetworkPolicySpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Description|string|true|true|true|false|
|ProjectID|reference[project]|true|true|true|true|

### Object `ProjectNetworkPolicyStatus`
--- 


### Object `ProjectNetworkPolicy`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|true|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Status|projectNetworkPolicyStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `ElasticsearchConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AuthPassword|string|true|true|true|false|
|AuthUserName|string|true|true|true|false|
|Certificate|string|true|true|true|false|
|ClientCert|string|true|true|true|false|
|ClientKey|string|true|true|true|false|
|ClientKeyPass|string|true|true|true|false|
|DateFormat|enum|true|true|true|true|
|Endpoint|string|true|true|true|true|
|IndexPrefix|string|true|true|true|true|
|SSLVerify|boolean|true|true|false|false|
|SSLVersion|enum|true|true|true|false|

### Object `SplunkConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Certificate|string|true|true|true|false|
|ClientCert|string|true|true|true|false|
|ClientKey|string|true|true|true|false|
|ClientKeyPass|string|true|true|true|false|
|Endpoint|string|true|true|true|true|
|Index|string|true|true|true|false|
|Source|string|true|true|true|false|
|SSLVerify|boolean|true|true|false|false|
|Token|string|true|true|true|true|

### Object `KafkaConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BrokerEndpoints|array[string]|true|true|true|false|
|Certificate|string|true|true|true|false|
|ClientCert|string|true|true|true|false|
|ClientKey|string|true|true|true|false|
|Topic|string|true|true|true|true|
|ZookeeperEndpoint|string|true|true|true|false|

### Object `SyslogConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Certificate|string|true|true|true|false|
|ClientCert|string|true|true|true|false|
|ClientKey|string|true|true|true|false|
|Endpoint|string|true|true|true|true|
|Program|string|true|true|true|false|
|Protocol|enum|true|true|true|false|
|Severity|enum|true|true|true|false|
|SSLVerify|boolean|true|true|false|false|
|Token|string|true|true|true|false|

### Object `FluentServer`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Endpoint|string|true|true|true|true|
|Hostname|string|true|true|true|false|
|Password|string|true|true|true|false|
|SharedKey|string|true|true|true|false|
|Standby|boolean|true|true|false|false|
|Username|string|true|true|true|false|
|Weight|int|true|true|false|false|

### Object `FluentForwarderConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Certificate|string|true|true|true|false|
|Compress|boolean|true|true|false|false|
|EnableTLS|boolean|true|true|false|false|
|FluentServers|array[fluentServer]|true|true|true|true|

### Object `ClusterLoggingSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterID|reference[cluster]|true|true|true|false|
|DisplayName|string|true|true|true|false|
|ElasticsearchConfig|elasticsearchConfig|true|true|true|false|
|FluentForwarderConfig|fluentForwarderConfig|true|true|true|false|
|KafkaConfig|kafkaConfig|true|true|true|false|
|OutputFlushInterval|int|true|true|false|false|
|OutputTags|map[string]|true|true|true|false|
|SplunkConfig|splunkConfig|true|true|true|false|
|SyslogConfig|syslogConfig|true|true|true|false|

### Object `LoggingCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|string|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `ClusterLoggingStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AppliedSpec|clusterLoggingSpec|false|false|true|false|
|Conditions|array[loggingCondition]|false|false|true|false|
|FailedSpec|clusterLoggingSpec|false|false|true|false|

### Object `ClusterLogging`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|AppliedSpec|clusterLoggingSpec|false|false|true|false|
|ClusterID|reference[cluster]|true|true|true|false|
|Conditions|array[loggingCondition]|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|ElasticsearchConfig|elasticsearchConfig|true|true|true|false|
|FailedSpec|clusterLoggingSpec|false|false|true|false|
|FluentForwarderConfig|fluentForwarderConfig|true|true|true|false|
|Id|dnsLabel|false|false|true|false|
|KafkaConfig|kafkaConfig|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|string|true|true|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OutputFlushInterval|int|true|true|false|false|
|OutputTags|map[string]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|SplunkConfig|splunkConfig|true|true|true|false|
|State|string|false|false|false|false|
|SyslogConfig|syslogConfig|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `ProjectLoggingSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DisplayName|string|true|true|true|false|
|ElasticsearchConfig|elasticsearchConfig|true|true|true|false|
|FluentForwarderConfig|fluentForwarderConfig|true|true|true|false|
|KafkaConfig|kafkaConfig|true|true|true|false|
|OutputFlushInterval|int|true|true|false|false|
|OutputTags|map[string]|true|true|true|false|
|ProjectID|reference[project]|true|true|true|false|
|SplunkConfig|splunkConfig|true|true|true|false|
|SyslogConfig|syslogConfig|true|true|true|false|

### Object `ProjectLoggingStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AppliedSpec|projectLoggingSpec|false|false|true|false|
|Conditions|array[loggingCondition]|false|false|true|false|

### Object `ProjectLogging`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|ElasticsearchConfig|elasticsearchConfig|true|true|true|false|
|FluentForwarderConfig|fluentForwarderConfig|true|true|true|false|
|Id|dnsLabel|false|false|true|false|
|KafkaConfig|kafkaConfig|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|string|true|true|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OutputFlushInterval|int|true|true|false|false|
|OutputTags|map[string]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|false|
|Removed|date|false|false|true|false|
|SplunkConfig|splunkConfig|true|true|true|false|
|State|string|false|false|false|false|
|Status|projectLoggingStatus|false|false|true|false|
|SyslogConfig|syslogConfig|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `ListenConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Algorithm|string|false|false|true|false|
|Annotations|map[string]|true|true|true|false|
|CACerts|string|true|true|true|false|
|Cert|string|true|true|true|false|
|CertFingerprint|string|false|false|true|false|
|CN|string|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Domains|array[string]|true|true|true|false|
|Enabled|boolean|true|true|false|false|
|ExpiresAt|string|false|false|true|false|
|GeneratedCerts|map[string]|false|false|true|false|
|Id|dnsLabel|false|false|true|false|
|IssuedAt|string|false|false|true|false|
|Issuer|string|false|false|true|false|
|Key|string|true|true|true|false|
|KeySize|int|false|false|false|false|
|KnownIPs|array[string]|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Mode|enum|true|true|true|false|
|Name|string|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|SerialNumber|string|false|false|true|false|
|SubjectAlternativeNames|array[string]|false|false|true|false|
|TOS|array[string]|true|true|true|false|
|UUID|string|false|false|true|false|
|Version|int|false|false|false|false|

### Object `Setting`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Customized|boolean|false|false|false|false|
|Default|string|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|
|Value|string|true|true|true|true|

### Object `Recipient`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|NotifierID|reference[notifier]|true|true|true|true|
|NotifierType|enum|true|true|true|true|
|Recipient|string|true|true|true|false|

### Object `TargetNode`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Condition|enum|true|true|true|true|
|CPUThreshold|int|true|true|false|false|
|MemThreshold|int|true|true|false|false|
|NodeID|reference[node]|true|true|true|false|
|Selector|map[string]|true|true|true|false|

### Object `TargetSystemService`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Condition|enum|true|true|true|true|

### Object `TargetEvent`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|EventType|enum|true|true|true|true|
|ResourceKind|enum|true|true|true|true|

### Object `ClusterAlertSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterID|reference[cluster]|true|true|true|false|
|Description|string|true|true|true|false|
|DisplayName|string|true|true|true|true|
|InitialWaitSeconds|int|true|true|false|true|
|Recipients|array[recipient]|true|true|true|true|
|RepeatIntervalSeconds|int|true|true|false|true|
|Severity|enum|true|true|true|true|
|TargetEvent|targetEvent|true|true|true|false|
|TargetNode|targetNode|true|true|true|false|
|TargetSystemService|targetSystemService|true|true|true|false|

### Object `AlertStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AlertState|enum|false|false|true|false|

### Object `ClusterAlert`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClusterID|reference[cluster]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|DisplayName|string|true|true|true|true|
|InitialWaitSeconds|int|true|true|false|true|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Recipients|array[recipient]|true|true|true|true|
|Removed|date|false|false|true|false|
|RepeatIntervalSeconds|int|true|true|false|true|
|Severity|enum|true|true|true|true|
|State|string|false|false|false|false|
|Status|alertStatus|false|false|true|false|
|TargetEvent|targetEvent|true|true|true|false|
|TargetNode|targetNode|true|true|true|false|
|TargetSystemService|targetSystemService|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `TargetWorkload`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AvailablePercentage|int|true|true|false|true|
|Selector|map[string]|true|true|true|false|
|WorkloadID|string|true|true|true|false|

### Object `TargetPod`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Condition|enum|true|true|true|true|
|PodID|reference[/v3/projects/schemas/pod]|true|true|true|true|
|RestartIntervalSeconds|int|true|true|false|false|
|RestartTimes|int|true|true|false|false|

### Object `ProjectAlertSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Description|string|true|true|true|false|
|DisplayName|string|true|true|true|true|
|InitialWaitSeconds|int|true|true|false|true|
|ProjectID|reference[project]|true|true|true|false|
|Recipients|array[recipient]|true|true|true|true|
|RepeatIntervalSeconds|int|true|true|false|true|
|Severity|enum|true|true|true|true|
|TargetPod|targetPod|true|true|true|false|
|TargetWorkload|targetWorkload|true|true|true|false|

### Object `ProjectAlert`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|DisplayName|string|true|true|true|true|
|InitialWaitSeconds|int|true|true|false|true|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|false|
|Recipients|array[recipient]|true|true|true|true|
|Removed|date|false|false|true|false|
|RepeatIntervalSeconds|int|true|true|false|true|
|Severity|enum|true|true|true|true|
|State|string|false|false|false|false|
|Status|alertStatus|false|false|true|false|
|TargetPod|targetPod|true|true|true|false|
|TargetWorkload|targetWorkload|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `SMTPConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DefaultRecipient|string|true|true|true|true|
|Host|hostname|true|true|true|true|
|Password|string|true|true|true|false|
|Port|int|true|true|false|true|
|Sender|string|true|true|true|true|
|TLS|boolean|true|true|false|true|
|Username|string|true|true|true|false|

### Object `SlackConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DefaultRecipient|string|true|true|true|true|
|URL|string|true|true|true|true|

### Object `PagerdutyConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ServiceKey|string|true|true|true|true|

### Object `WebhookConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|URL|string|true|true|true|true|

### Object `Notification`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Message|string|true|true|true|false|
|PagerdutyConfig|pagerdutyConfig|true|true|true|false|
|SlackConfig|slackConfig|true|true|true|false|
|SMTPConfig|smtpConfig|true|true|true|false|
|WebhookConfig|webhookConfig|true|true|true|false|

### Object `NotifierSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterID|reference[cluster]|true|true|true|false|
|Description|string|true|true|true|false|
|DisplayName|string|true|true|true|true|
|PagerdutyConfig|pagerdutyConfig|true|true|true|false|
|SlackConfig|slackConfig|true|true|true|false|
|SMTPConfig|smtpConfig|true|true|true|false|
|WebhookConfig|webhookConfig|true|true|true|false|

### Object `NotifierStatus`
--- 


### Object `Notifier`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClusterID|reference[cluster]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|string|true|true|true|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PagerdutyConfig|pagerdutyConfig|true|true|true|false|
|Removed|date|false|false|true|false|
|SlackConfig|slackConfig|true|true|true|false|
|SMTPConfig|smtpConfig|true|true|true|false|
|State|string|false|false|false|false|
|Status|notifierStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|
|WebhookConfig|webhookConfig|true|true|true|false|

### Object `ClusterGroupSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterID|reference[cluster]|true|true|true|false|
|Description|string|true|true|true|false|
|DisplayName|string|true|true|true|true|
|GroupIntervalSeconds|int|true|true|false|true|
|GroupWaitSeconds|int|true|true|false|true|
|Recipients|array[recipient]|true|true|true|true|
|RepeatIntervalSeconds|int|true|true|false|true|

### Object `ClusterAlertGroup`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AlertState|enum|false|false|true|false|
|Annotations|map[string]|true|true|true|false|
|ClusterID|reference[cluster]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|GroupIntervalSeconds|int|true|true|false|true|
|GroupWaitSeconds|int|true|true|false|true|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|string|true|true|true|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Recipients|array[recipient]|true|true|true|true|
|Removed|date|false|false|true|false|
|RepeatIntervalSeconds|int|true|true|false|true|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `ProjectGroupSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Description|string|true|true|true|false|
|DisplayName|string|true|true|true|true|
|GroupIntervalSeconds|int|true|true|false|true|
|GroupWaitSeconds|int|true|true|false|true|
|ProjectID|reference[project]|true|true|true|false|
|Recipients|array[recipient]|true|true|true|true|
|RepeatIntervalSeconds|int|true|true|false|true|

### Object `ProjectAlertGroup`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AlertState|enum|false|false|true|false|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|GroupIntervalSeconds|int|true|true|false|true|
|GroupWaitSeconds|int|true|true|false|true|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|string|true|true|true|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|false|
|Recipients|array[recipient]|true|true|true|true|
|Removed|date|false|false|true|false|
|RepeatIntervalSeconds|int|true|true|false|true|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `NodeRule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Condition|enum|true|true|true|true|
|CPUThreshold|int|true|true|false|false|
|MemThreshold|int|true|true|false|false|
|NodeID|reference[node]|true|true|true|false|
|Selector|map[string]|true|true|true|false|

### Object `EventRule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|EventType|enum|true|true|true|true|
|ResourceKind|enum|true|true|true|true|

### Object `SystemServiceRule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Condition|enum|true|true|true|true|

### Object `MetricRule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Comparison|enum|true|true|true|false|
|Description|string|true|true|true|false|
|Duration|string|true|true|true|false|
|Expression|string|true|true|true|true|
|LegendFormat|string|true|true|true|false|
|Step|int|true|true|false|false|
|ThresholdValue|float|true|true|true|false|

### Object `ClusterAlertRuleSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterID|reference[cluster]|true|true|true|false|
|DisplayName|string|true|true|true|false|
|EventRule|eventRule|true|true|true|false|
|GroupID|reference[clusterAlertGroup]|true|true|true|false|
|GroupIntervalSeconds|int|true|true|false|true|
|GroupWaitSeconds|int|true|true|false|true|
|MetricRule|metricRule|true|true|true|false|
|NodeRule|nodeRule|true|true|true|false|
|RepeatIntervalSeconds|int|true|true|false|true|
|Severity|enum|true|true|true|true|
|SystemServiceRule|systemServiceRule|true|true|true|false|

### Object `ClusterAlertRule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AlertState|enum|false|false|true|false|
|Annotations|map[string]|true|true|true|false|
|ClusterID|reference[cluster]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|EventRule|eventRule|true|true|true|false|
|GroupID|reference[clusterAlertGroup]|true|true|true|false|
|GroupIntervalSeconds|int|true|true|false|true|
|GroupWaitSeconds|int|true|true|false|true|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|MetricRule|metricRule|true|true|true|false|
|Name|string|true|true|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeRule|nodeRule|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|RepeatIntervalSeconds|int|true|true|false|true|
|Severity|enum|true|true|true|true|
|State|string|false|false|false|false|
|SystemServiceRule|systemServiceRule|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `PodRule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Condition|enum|true|true|true|true|
|PodID|reference[/v3/projects/schemas/pod]|true|true|true|true|
|RestartIntervalSeconds|int|true|true|false|false|
|RestartTimes|int|true|true|false|false|

### Object `WorkloadRule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AvailablePercentage|int|true|true|false|true|
|Selector|map[string]|true|true|true|false|
|WorkloadID|string|true|true|true|false|

### Object `ProjectAlertRuleSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DisplayName|string|true|true|true|false|
|GroupID|reference[projectAlertGroup]|true|true|true|false|
|GroupIntervalSeconds|int|true|true|false|true|
|GroupWaitSeconds|int|true|true|false|true|
|MetricRule|metricRule|true|true|true|false|
|PodRule|podRule|true|true|true|false|
|ProjectID|reference[project]|true|true|true|false|
|RepeatIntervalSeconds|int|true|true|false|true|
|Severity|enum|true|true|true|true|
|WorkloadRule|workloadRule|true|true|true|false|

### Object `ProjectAlertRule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AlertState|enum|false|false|true|false|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|GroupID|reference[projectAlertGroup]|true|true|true|false|
|GroupIntervalSeconds|int|true|true|false|true|
|GroupWaitSeconds|int|true|true|false|true|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|MetricRule|metricRule|true|true|true|false|
|Name|string|true|true|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PodRule|podRule|true|true|true|false|
|ProjectID|reference[project]|true|true|true|false|
|Removed|date|false|false|true|false|
|RepeatIntervalSeconds|int|true|true|false|true|
|Severity|enum|true|true|true|true|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|
|WorkloadRule|workloadRule|true|true|true|false|

### Object `ComposeSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|RancherCompose|string|true|true|true|false|

### Object `ComposeCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|string|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `ComposeStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Conditions|array[composeCondition]|false|false|true|false|

### Object `ComposeConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|RancherCompose|string|true|true|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Status|composeStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `ProjectCatalog`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Branch|string|true|true|true|false|
|Commit|string|false|false|true|false|
|Conditions|array[catalogCondition]|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Kind|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|LastRefreshTimestamp|string|false|false|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Password|password|true|true|true|false|
|ProjectID|reference[project]|true|true|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|URL|string|true|true|true|true|
|Username|string|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `ClusterCatalog`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Branch|string|true|true|true|false|
|ClusterID|reference[cluster]|true|true|true|true|
|Commit|string|false|false|true|false|
|Conditions|array[catalogCondition]|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Kind|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|LastRefreshTimestamp|string|false|false|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Password|password|true|true|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|URL|string|true|true|true|true|
|Username|string|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `Answer`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterID|reference[cluster]|true|true|true|false|
|ProjectID|reference[project]|true|true|true|false|
|Values|map[string]|true|true|true|true|

### Object `Target`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AppID|reference[v3/projects/schemas/app]|true|true|true|false|
|Healthstate|string|true|true|true|false|
|ProjectID|reference[project]|true|true|true|true|

### Object `MultiClusterAppSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Answers|array[answer]|true|true|true|false|
|Targets|array[target]|true|true|true|true|
|TemplateVersionID|reference[templateVersion]|true|true|true|true|

### Object `MultiClusterAppStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Healthstate|string|false|false|true|false|

### Object `MultiClusterApp`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Answers|array[answer]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Status|multiClusterAppStatus|false|false|true|false|
|Targets|array[target]|true|true|true|true|
|TemplateVersionID|reference[templateVersion]|true|true|true|true|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `GlobalDNSSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FQDN|string|true|true|true|false|
|MultiClusterAppID|reference[multiClusterApp]|true|true|true|false|
|ProjectIDs|array[reference[project]]|true|true|true|false|
|ProviderID|reference[globalDnsProvider]|true|true|true|false|

### Object `GlobalDNSStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Endpoints|array[string]|false|false|true|false|

### Object `GlobalDNS`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|FQDN|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|MultiClusterAppID|reference[multiClusterApp]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectIDs|array[reference[project]]|true|true|true|false|
|ProviderID|reference[globalDnsProvider]|true|true|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Status|globalDnsStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `Route53ProviderConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessKey|string|true|true|true|false|
|RootDomain|string|true|true|true|true|
|SecretKey|password|true|true|true|false|

### Object `GlobalDNSProviderSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Route53ProviderConfig|route53ProviderConfig|true|true|true|false|

### Object `GlobalDNSProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|Route53ProviderConfig|route53ProviderConfig|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `KontainerDriverSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Active|boolean|true|true|false|false|
|BuiltIn|boolean|true|false|false|false|
|Checksum|string|true|true|true|false|
|UIURL|string|true|true|true|false|
|URL|string|true|true|true|true|
|WhitelistDomains|array[string]|true|true|true|false|

### Object `KontainerDriverStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActualURL|string|false|false|true|false|
|Conditions|array[condition]|false|false|true|false|
|DisplayName|string|false|false|true|false|
|ExecutablePath|string|false|false|true|false|

### Object `KontainerDriver`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Active|boolean|true|true|false|false|
|ActualURL|string|false|false|true|false|
|Annotations|map[string]|true|true|true|false|
|BuiltIn|boolean|true|false|false|false|
|Checksum|string|true|true|true|false|
|Conditions|array[condition]|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|ExecutablePath|string|false|false|true|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|string|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UIURL|string|true|true|true|false|
|URL|string|true|true|true|true|
|UUID|string|false|false|true|false|
|WhitelistDomains|array[string]|true|true|true|false|

### Object `QueryGraphInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Filters|map[string]|true|true|true|false|
|From|string|true|true|true|false|
|Interval|string|true|true|true|false|
|IsDetails|boolean|true|true|false|false|
|MetricParams|map[string]|true|true|true|false|
|To|string|true|true|true|false|

### Object `QueryClusterGraph`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|GraphName|reference[clusterMonitorGraph]|true|true|true|false|
|Series|array[reference[timeSeries]]|true|true|true|false|

### Object `QueryClusterGraphOutput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Data|array[queryClusterGraph]|true|true|true|false|
|Type|string|true|true|true|false|

### Object `QueryProjectGraph`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|GraphName|reference[projectMonitorGraph]|true|true|true|false|
|Series|array[reference[timeSeries]]|true|true|true|false|

### Object `QueryProjectGraphOutput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Data|array[queryProjectGraph]|true|true|true|false|
|Type|string|true|true|true|false|

### Object `QueryClusterMetricInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterName|reference[cluster]|true|true|true|false|
|Expr|string|true|true|true|true|
|From|string|true|true|true|false|
|Interval|string|true|true|true|false|
|To|string|true|true|true|false|

### Object `QueryProjectMetricInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Expr|string|true|true|true|true|
|From|string|true|true|true|false|
|Interval|string|true|true|true|false|
|ProjectName|reference[project]|true|true|true|false|
|To|string|true|true|true|false|

### Object `QueryMetricOutput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Series|array[reference[timeSeries]]|true|true|true|false|
|Type|string|true|true|true|false|

### Object `ClusterMetricNamesInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterName|reference[cluster]|true|true|true|false|

### Object `ProjectMetricNamesInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ProjectName|reference[project]|true|true|true|false|

### Object `MetricNamesOutput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Names|array[string]|true|true|true|false|
|Type|string|true|true|true|false|

### Object `TimeSeries`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|
|Points|array[array[float]]|true|true|true|false|

### Object `MonitorMetricSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Description|string|true|true|true|false|
|Expression|string|true|true|true|true|
|LegendFormat|string|true|true|true|false|

### Object `MonitorMetric`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Expression|string|true|true|true|true|
|Labels|map[string]|true|true|true|false|
|LegendFormat|string|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `YAxis`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Unit|string|true|true|true|false|

### Object `ClusterMonitorGraphSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterID|reference[cluster]|true|true|true|false|
|Description|string|true|true|true|false|
|DetailsMetricsSelector|map[string]|true|true|true|false|
|DisplayResourceType|enum|true|true|true|false|
|GraphType|enum|true|true|true|false|
|MetricsSelector|map[string]|true|true|true|false|
|Priority|int|true|true|false|false|
|ResourceType|enum|true|true|true|false|
|YAxis|yAxis|true|true|true|false|

### Object `ClusterMonitorGraph`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClusterID|reference[cluster]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|DetailsMetricsSelector|map[string]|true|true|true|false|
|DisplayResourceType|enum|true|true|true|false|
|GraphType|enum|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|MetricsSelector|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Priority|int|true|true|false|false|
|Removed|date|false|false|true|false|
|ResourceType|enum|true|true|true|false|
|UUID|string|false|false|true|false|
|YAxis|yAxis|true|true|true|false|

### Object `ProjectMonitorGraphSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Description|string|true|true|true|false|
|DetailsMetricsSelector|map[string]|true|true|true|false|
|DisplayResourceType|enum|true|true|true|false|
|GraphType|enum|true|true|true|false|
|MetricsSelector|map[string]|true|true|true|false|
|Priority|int|true|true|false|false|
|ProjectID|reference[project]|true|true|true|false|
|ResourceType|enum|true|true|true|false|
|YAxis|yAxis|true|true|true|false|

### Object `ProjectMonitorGraph`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|DetailsMetricsSelector|map[string]|true|true|true|false|
|DisplayResourceType|enum|true|true|true|false|
|GraphType|enum|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|MetricsSelector|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Priority|int|true|true|false|false|
|ProjectID|reference[project]|true|true|true|false|
|Removed|date|false|false|true|false|
|ResourceType|enum|true|true|true|false|
|UUID|string|false|false|true|false|
|YAxis|yAxis|true|true|true|false|

### Object `ResourceQuotaLimit`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ConfigMaps|string|true|true|true|false|
|LimitsCPU|string|true|true|true|false|
|LimitsMemory|string|true|true|true|false|
|PersistentVolumeClaims|string|true|true|true|false|
|Pods|string|true|true|true|false|
|ReplicationControllers|string|true|true|true|false|
|RequestsCPU|string|true|true|true|false|
|RequestsMemory|string|true|true|true|false|
|RequestsStorage|string|true|true|true|false|
|Secrets|string|true|true|true|false|
|Services|string|true|true|true|false|
|ServicesLoadBalancers|string|true|true|true|false|
|ServicesNodePorts|string|true|true|true|false|

### Object `NamespaceResourceQuota`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Limit|resourceQuotaLimit|true|true|true|false|

### Object `OwnerReference`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|APIVersion|string|true|true|true|false|
|BlockOwnerDeletion|boolean|true|true|true|false|
|Controller|boolean|true|true|true|false|
|Kind|string|true|true|true|false|
|Name|string|true|true|true|false|
|UID|string|true|true|true|false|

### Object `Initializer`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|

### Object `ListMeta`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Continue|string|true|true|true|false|
|ResourceVersion|string|true|true|true|false|
|SelfLink|string|true|true|true|false|

### Object `StatusCause`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Field|string|true|true|true|false|
|Message|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `StatusDetails`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Causes|array[statusCause]|true|true|true|false|
|Group|string|true|true|true|false|
|Kind|string|true|true|true|false|
|Name|string|true|true|true|false|
|RetryAfterSeconds|int|true|true|false|false|
|UID|string|true|true|true|false|

### Object `Status`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|APIVersion|string|true|true|true|false|
|Code|int|true|true|false|false|
|Details|statusDetails|true|true|true|false|
|Kind|string|true|true|true|false|
|Message|string|true|true|true|false|
|ListMeta|listMeta|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|

### Object `Initializers`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Pending|array[initializer]|true|true|true|false|
|Result|status|true|true|true|false|

### Object `ObjectMeta`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|Finalizers|array[string]|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|Namespace|string|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|SelfLink|string|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `NamespaceSpec`
--- 


### Object `NamespaceStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Phase|string|false|false|true|false|

### Object `Namespace`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|ResourceQuota|namespaceResourceQuota|true|true|true|false|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `NamespaceMove`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ProjectID|string|true|true|true|false|

### Object `GCEPersistentDiskVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|Partition|int|true|true|false|false|
|PDName|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|

### Object `AWSElasticBlockStoreVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|Partition|int|true|true|false|false|
|ReadOnly|boolean|true|true|false|false|
|VolumeID|string|true|true|true|false|

### Object `HostPathVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Kind|enum|true|true|true|false|
|Path|string|true|true|true|false|

### Object `GlusterfsVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|EndpointsName|string|true|true|true|false|
|Path|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|

### Object `NFSVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Path|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|Server|string|true|true|true|false|

### Object `SecretReference`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|
|Namespace|string|true|true|true|false|

### Object `RBDPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|RBDImage|string|true|true|true|false|
|Keyring|string|true|true|true|false|
|CephMonitors|array[string]|true|true|true|false|
|RBDPool|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|secretReference|true|true|true|false|
|RadosUser|string|true|true|true|false|

### Object `ISCSIPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DiscoveryCHAPAuth|boolean|true|true|false|false|
|SessionCHAPAuth|boolean|true|true|false|false|
|FSType|string|true|true|true|false|
|InitiatorName|string|true|true|true|false|
|IQN|string|true|true|true|false|
|ISCSIInterface|string|true|true|true|false|
|Lun|int|true|true|false|false|
|Portals|array[string]|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|secretReference|true|true|true|false|
|TargetPortal|string|true|true|true|false|

### Object `CinderPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|secretReference|true|true|true|false|
|VolumeID|string|true|true|true|false|

### Object `CephFSPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Monitors|array[string]|true|true|true|false|
|Path|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretFile|string|true|true|true|false|
|SecretRef|secretReference|true|true|true|false|
|User|string|true|true|true|false|

### Object `FCVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|Lun|int|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|TargetWWNs|array[string]|true|true|true|false|
|WWIDs|array[string]|true|true|true|false|

### Object `FlockerVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DatasetName|string|true|true|true|false|
|DatasetUUID|string|true|true|true|false|

### Object `FlexPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Driver|string|true|true|true|false|
|FSType|string|true|true|true|false|
|Options|map[string]|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|secretReference|true|true|true|false|

### Object `AzureFilePersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ReadOnly|boolean|true|true|false|false|
|SecretName|string|true|true|true|false|
|SecretNamespace|string|true|true|true|false|
|ShareName|string|true|true|true|false|

### Object `VsphereVirtualDiskVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|StoragePolicyID|string|true|true|true|false|
|StoragePolicyName|string|true|true|true|false|
|VolumePath|string|true|true|true|false|

### Object `QuobyteVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Group|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|Registry|string|true|true|true|false|
|User|string|true|true|true|false|
|Volume|string|true|true|true|false|

### Object `AzureDiskVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CachingMode|string|true|true|true|false|
|DiskName|string|true|true|true|false|
|DataDiskURI|string|true|true|true|false|
|FSType|string|true|true|true|false|
|Kind|string|true|true|true|false|
|ReadOnly|boolean|true|true|true|false|

### Object `PhotonPersistentDiskVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|PdID|string|true|true|true|false|

### Object `PortworxVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|VolumeID|string|true|true|true|false|

### Object `ScaleIOPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|Gateway|string|true|true|true|false|
|ProtectionDomain|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|secretReference|true|true|true|false|
|SSLEnabled|boolean|true|true|false|false|
|StorageMode|string|true|true|true|false|
|StoragePool|string|true|true|true|false|
|System|string|true|true|true|false|
|VolumeName|string|true|true|true|false|

### Object `LocalVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|Path|string|true|true|true|false|

### Object `ObjectReference`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|APIVersion|string|true|true|true|false|
|FieldPath|string|true|true|true|false|
|Kind|string|true|true|true|false|
|Name|string|true|true|true|false|
|Namespace|string|true|true|true|false|
|ResourceVersion|string|true|true|true|false|
|UID|string|true|true|true|false|

### Object `StorageOSPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|objectReference|true|true|true|false|
|VolumeName|string|true|true|true|false|
|VolumeNamespace|string|true|true|true|false|

### Object `CSIPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ControllerPublishSecretRef|secretReference|true|true|true|false|
|Driver|string|true|true|true|false|
|FSType|string|true|true|true|false|
|NodePublishSecretRef|secretReference|true|true|true|false|
|NodeStageSecretRef|secretReference|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|VolumeAttributes|map[string]|true|true|true|false|
|VolumeHandle|string|true|true|true|false|

### Object `NodeSelectorRequirement`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Key|string|true|true|true|false|
|Operator|string|true|true|true|false|
|Values|array[string]|true|true|true|false|

### Object `NodeSelectorTerm`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MatchExpressions|array[nodeSelectorRequirement]|true|true|true|false|
|MatchFields|array[nodeSelectorRequirement]|true|true|true|false|

### Object `NodeSelector`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|NodeSelectorTerms|array[nodeSelectorTerm]|true|true|true|false|

### Object `VolumeNodeAffinity`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Required|nodeSelector|true|true|true|false|

### Object `PersistentVolumeSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessModes|array[string]|true|true|true|false|
|AWSElasticBlockStore|awsElasticBlockStoreVolumeSource|true|true|true|false|
|AzureDisk|azureDiskVolumeSource|true|true|true|false|
|AzureFile|azureFilePersistentVolumeSource|true|true|true|false|
|Capacity|map[string]|true|true|true|false|
|CephFS|cephFSPersistentVolumeSource|true|true|true|false|
|Cinder|cinderPersistentVolumeSource|true|true|true|false|
|ClaimRef|objectReference|true|true|true|false|
|CSI|csiPersistentVolumeSource|true|true|true|false|
|FC|fcVolumeSource|true|true|true|false|
|FlexVolume|flexPersistentVolumeSource|true|true|true|false|
|Flocker|flockerVolumeSource|true|true|true|false|
|GCEPersistentDisk|gcePersistentDiskVolumeSource|true|true|true|false|
|Glusterfs|glusterfsVolumeSource|true|true|true|false|
|HostPath|hostPathVolumeSource|true|true|true|false|
|ISCSI|iscsiPersistentVolumeSource|true|true|true|false|
|Local|localVolumeSource|true|true|true|false|
|MountOptions|array[string]|true|true|true|false|
|NFS|nfsVolumeSource|true|true|true|false|
|NodeAffinity|volumeNodeAffinity|true|true|true|false|
|PersistentVolumeReclaimPolicy|string|true|true|true|false|
|PhotonPersistentDisk|photonPersistentDiskVolumeSource|true|true|true|false|
|PortworxVolume|portworxVolumeSource|true|true|true|false|
|Quobyte|quobyteVolumeSource|true|true|true|false|
|RBD|rbdPersistentVolumeSource|true|true|true|false|
|ScaleIO|scaleIOPersistentVolumeSource|true|true|true|false|
|StorageClassID|reference[storageClass]|true|true|true|false|
|StorageOS|storageOSPersistentVolumeSource|true|true|true|false|
|VolumeMode|string|true|true|true|false|
|VsphereVolume|vsphereVirtualDiskVolumeSource|true|true|true|false|

### Object `PersistentVolumeStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Message|string|false|false|true|false|
|Phase|string|false|false|true|false|
|Reason|string|false|false|true|false|

### Object `PersistentVolume`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessModes|array[string]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|AWSElasticBlockStore|awsElasticBlockStoreVolumeSource|true|true|true|false|
|AzureDisk|azureDiskVolumeSource|true|true|true|false|
|AzureFile|azureFilePersistentVolumeSource|true|true|true|false|
|Capacity|map[string]|true|true|true|false|
|CephFS|cephFSPersistentVolumeSource|true|true|true|false|
|Cinder|cinderPersistentVolumeSource|true|true|true|false|
|ClaimRef|objectReference|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|CSI|csiPersistentVolumeSource|true|true|true|false|
|Description|string|true|true|true|false|
|FC|fcVolumeSource|true|true|true|false|
|FlexVolume|flexPersistentVolumeSource|true|true|true|false|
|Flocker|flockerVolumeSource|true|true|true|false|
|GCEPersistentDisk|gcePersistentDiskVolumeSource|true|true|true|false|
|Glusterfs|glusterfsVolumeSource|true|true|true|false|
|HostPath|hostPathVolumeSource|true|true|true|false|
|ISCSI|iscsiPersistentVolumeSource|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Local|localVolumeSource|true|true|true|false|
|MountOptions|array[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NFS|nfsVolumeSource|true|true|true|false|
|NodeAffinity|volumeNodeAffinity|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PersistentVolumeReclaimPolicy|string|true|true|true|false|
|PhotonPersistentDisk|photonPersistentDiskVolumeSource|true|true|true|false|
|PortworxVolume|portworxVolumeSource|true|true|true|false|
|Quobyte|quobyteVolumeSource|true|true|true|false|
|RBD|rbdPersistentVolumeSource|true|true|true|false|
|Removed|date|false|false|true|false|
|ScaleIO|scaleIOPersistentVolumeSource|true|true|true|false|
|State|string|false|false|false|false|
|Status|persistentVolumeStatus|false|false|true|false|
|StorageClassID|reference[storageClass]|true|true|true|false|
|StorageOS|storageOSPersistentVolumeSource|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|
|VolumeMode|string|true|true|true|false|
|VsphereVolume|vsphereVirtualDiskVolumeSource|true|true|true|false|

### Object `TopologySelectorLabelRequirement`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Key|string|true|true|true|false|
|Values|array[string]|true|true|true|false|

### Object `TopologySelectorTerm`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MatchLabelExpressions|array[topologySelectorLabelRequirement]|true|true|true|false|

### Object `StorageClass`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AllowVolumeExpansion|boolean|true|true|true|false|
|AllowedTopologies|array[topologySelectorTerm]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|MountOptions|array[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Parameters|map[string]|true|true|true|false|
|Provisioner|string|true|true|true|false|
|ReclaimPolicy|enum|true|true|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|
|VolumeBindingMode|string|true|true|true|false|

### Object `PersistentVolumeClaimVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|PersistentVolumeClaimID|reference[persistentVolumeClaim]|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|

### Object `KeyToPath`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Key|string|true|true|true|false|
|Mode|int|true|true|true|false|
|Path|string|true|true|true|false|

### Object `SecretVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DefaultMode|int|true|true|true|false|
|Items|array[keyToPath]|true|true|true|false|
|Optional|boolean|true|true|true|false|
|SecretName|string|true|true|true|false|

### Object `VolumeMount`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MountPath|string|true|true|true|true|
|MountPropagation|string|true|true|true|false|
|Name|string|true|true|true|true|
|ReadOnly|boolean|true|true|false|false|
|SubPath|string|true|true|true|false|

### Object `HostPathVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Kind|enum|true|true|true|false|
|Path|string|true|true|true|false|

### Object `EmptyDirVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Medium|string|true|true|true|false|
|SizeLimit|string|true|true|true|false|

### Object `GCEPersistentDiskVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|Partition|int|true|true|false|false|
|PDName|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|

### Object `AWSElasticBlockStoreVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|Partition|int|true|true|false|false|
|ReadOnly|boolean|true|true|false|false|
|VolumeID|string|true|true|true|false|

### Object `GitRepoVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Directory|string|true|true|true|false|
|Repository|string|true|true|true|false|
|Revision|string|true|true|true|false|

### Object `NFSVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Path|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|Server|string|true|true|true|false|

### Object `LocalObjectReference`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|

### Object `ISCSIVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DiscoveryCHAPAuth|boolean|true|true|false|false|
|SessionCHAPAuth|boolean|true|true|false|false|
|FSType|string|true|true|true|false|
|InitiatorName|string|true|true|true|false|
|IQN|string|true|true|true|false|
|ISCSIInterface|string|true|true|true|false|
|Lun|int|true|true|false|false|
|Portals|array[string]|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|localObjectReference|true|true|true|false|
|TargetPortal|string|true|true|true|false|

### Object `GlusterfsVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|EndpointsName|string|true|true|true|false|
|Path|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|

### Object `RBDVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|RBDImage|string|true|true|true|false|
|Keyring|string|true|true|true|false|
|CephMonitors|array[string]|true|true|true|false|
|RBDPool|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|localObjectReference|true|true|true|false|
|RadosUser|string|true|true|true|false|

### Object `FlexVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Driver|string|true|true|true|false|
|FSType|string|true|true|true|false|
|Options|map[string]|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|localObjectReference|true|true|true|false|

### Object `CinderVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|localObjectReference|true|true|true|false|
|VolumeID|string|true|true|true|false|

### Object `CephFSVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Monitors|array[string]|true|true|true|false|
|Path|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretFile|string|true|true|true|false|
|SecretRef|localObjectReference|true|true|true|false|
|User|string|true|true|true|false|

### Object `FlockerVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DatasetName|string|true|true|true|false|
|DatasetUUID|string|true|true|true|false|

### Object `ObjectFieldSelector`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|APIVersion|string|true|true|true|false|
|FieldPath|string|true|true|true|false|

### Object `ResourceFieldSelector`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ContainerName|string|true|true|true|false|
|Divisor|string|true|true|true|false|
|Resource|string|true|true|true|false|

### Object `DownwardAPIVolumeFile`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FieldRef|objectFieldSelector|true|true|true|false|
|Mode|int|true|true|true|false|
|Path|string|true|true|true|false|
|ResourceFieldRef|resourceFieldSelector|true|true|true|false|

### Object `DownwardAPIVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DefaultMode|int|true|true|true|false|
|Items|array[downwardAPIVolumeFile]|true|true|true|false|

### Object `FCVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|Lun|int|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|TargetWWNs|array[string]|true|true|true|false|
|WWIDs|array[string]|true|true|true|false|

### Object `AzureFileVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ReadOnly|boolean|true|true|false|false|
|SecretName|string|true|true|true|false|
|ShareName|string|true|true|true|false|

### Object `ConfigMapVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DefaultMode|int|true|true|true|false|
|Items|array[keyToPath]|true|true|true|false|
|Name|string|true|true|true|false|
|Optional|boolean|true|true|true|false|

### Object `VsphereVirtualDiskVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|StoragePolicyID|string|true|true|true|false|
|StoragePolicyName|string|true|true|true|false|
|VolumePath|string|true|true|true|false|

### Object `QuobyteVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Group|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|Registry|string|true|true|true|false|
|User|string|true|true|true|false|
|Volume|string|true|true|true|false|

### Object `AzureDiskVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CachingMode|string|true|true|true|false|
|DiskName|string|true|true|true|false|
|DataDiskURI|string|true|true|true|false|
|FSType|string|true|true|true|false|
|Kind|string|true|true|true|false|
|ReadOnly|boolean|true|true|true|false|

### Object `PhotonPersistentDiskVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|PdID|string|true|true|true|false|

### Object `SecretProjection`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Items|array[keyToPath]|true|true|true|false|
|Name|string|true|true|true|false|
|Optional|boolean|true|true|true|false|

### Object `DownwardAPIProjection`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Items|array[downwardAPIVolumeFile]|true|true|true|false|

### Object `ConfigMapProjection`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Items|array[keyToPath]|true|true|true|false|
|Name|string|true|true|true|false|
|Optional|boolean|true|true|true|false|

### Object `ServiceAccountTokenProjection`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Audience|string|true|true|true|false|
|ExpirationSeconds|int|true|true|true|false|
|Path|string|true|true|true|false|

### Object `VolumeProjection`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ConfigMap|configMapProjection|true|true|true|false|
|DownwardAPI|downwardAPIProjection|true|true|true|false|
|Secret|secretProjection|true|true|true|false|
|ServiceAccountToken|serviceAccountTokenProjection|true|true|true|false|

### Object `ProjectedVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DefaultMode|int|true|true|true|false|
|Sources|array[volumeProjection]|true|true|true|false|

### Object `PortworxVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|VolumeID|string|true|true|true|false|

### Object `ScaleIOVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|Gateway|string|true|true|true|false|
|ProtectionDomain|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|localObjectReference|true|true|true|false|
|SSLEnabled|boolean|true|true|false|false|
|StorageMode|string|true|true|true|false|
|StoragePool|string|true|true|true|false|
|System|string|true|true|true|false|
|VolumeName|string|true|true|true|false|

### Object `StorageOSVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|localObjectReference|true|true|true|false|
|VolumeName|string|true|true|true|false|
|VolumeNamespace|string|true|true|true|false|

### Object `Volume`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AWSElasticBlockStore|awsElasticBlockStoreVolumeSource|true|true|true|false|
|AzureDisk|azureDiskVolumeSource|true|true|true|false|
|AzureFile|azureFileVolumeSource|true|true|true|false|
|CephFS|cephFSVolumeSource|true|true|true|false|
|Cinder|cinderVolumeSource|true|true|true|false|
|ConfigMap|configMapVolumeSource|true|true|true|false|
|DownwardAPI|downwardAPIVolumeSource|true|true|true|false|
|EmptyDir|emptyDirVolumeSource|true|true|true|false|
|FC|fcVolumeSource|true|true|true|false|
|FlexVolume|flexVolumeSource|true|true|true|false|
|Flocker|flockerVolumeSource|true|true|true|false|
|GCEPersistentDisk|gcePersistentDiskVolumeSource|true|true|true|false|
|GitRepo|gitRepoVolumeSource|true|true|true|false|
|Glusterfs|glusterfsVolumeSource|true|true|true|false|
|HostPath|hostPathVolumeSource|true|true|true|false|
|ISCSI|iscsiVolumeSource|true|true|true|false|
|Name|string|true|true|true|false|
|NFS|nfsVolumeSource|true|true|true|false|
|PersistentVolumeClaim|persistentVolumeClaimVolumeSource|true|true|true|false|
|PhotonPersistentDisk|photonPersistentDiskVolumeSource|true|true|true|false|
|PortworxVolume|portworxVolumeSource|true|true|true|false|
|Projected|projectedVolumeSource|true|true|true|false|
|Quobyte|quobyteVolumeSource|true|true|true|false|
|RBD|rbdVolumeSource|true|true|true|false|
|ScaleIO|scaleIOVolumeSource|true|true|true|false|
|Secret|secretVolumeSource|true|true|true|false|
|StorageOS|storageOSVolumeSource|true|true|true|false|
|VsphereVolume|vsphereVirtualDiskVolumeSource|true|true|true|false|

### Object `SecretReference`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|
|Namespace|string|true|true|true|false|

### Object `RBDPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|RBDImage|string|true|true|true|false|
|Keyring|string|true|true|true|false|
|CephMonitors|array[string]|true|true|true|false|
|RBDPool|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|secretReference|true|true|true|false|
|RadosUser|string|true|true|true|false|

### Object `ISCSIPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DiscoveryCHAPAuth|boolean|true|true|false|false|
|SessionCHAPAuth|boolean|true|true|false|false|
|FSType|string|true|true|true|false|
|InitiatorName|string|true|true|true|false|
|IQN|string|true|true|true|false|
|ISCSIInterface|string|true|true|true|false|
|Lun|int|true|true|false|false|
|Portals|array[string]|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|secretReference|true|true|true|false|
|TargetPortal|string|true|true|true|false|

### Object `CinderPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|secretReference|true|true|true|false|
|VolumeID|string|true|true|true|false|

### Object `CephFSPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Monitors|array[string]|true|true|true|false|
|Path|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretFile|string|true|true|true|false|
|SecretRef|secretReference|true|true|true|false|
|User|string|true|true|true|false|

### Object `FlexPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Driver|string|true|true|true|false|
|FSType|string|true|true|true|false|
|Options|map[string]|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|secretReference|true|true|true|false|

### Object `AzureFilePersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ReadOnly|boolean|true|true|false|false|
|SecretName|string|true|true|true|false|
|SecretNamespace|string|true|true|true|false|
|ShareName|string|true|true|true|false|

### Object `ScaleIOPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|Gateway|string|true|true|true|false|
|ProtectionDomain|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|secretReference|true|true|true|false|
|SSLEnabled|boolean|true|true|false|false|
|StorageMode|string|true|true|true|false|
|StoragePool|string|true|true|true|false|
|System|string|true|true|true|false|
|VolumeName|string|true|true|true|false|

### Object `LocalVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|Path|string|true|true|true|false|

### Object `ObjectReference`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|APIVersion|string|true|true|true|false|
|FieldPath|string|true|true|true|false|
|Kind|string|true|true|true|false|
|Name|string|true|true|true|false|
|Namespace|string|true|true|true|false|
|ResourceVersion|string|true|true|true|false|
|UID|string|true|true|true|false|

### Object `StorageOSPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|FSType|string|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|SecretRef|objectReference|true|true|true|false|
|VolumeName|string|true|true|true|false|
|VolumeNamespace|string|true|true|true|false|

### Object `CSIPersistentVolumeSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ControllerPublishSecretRef|secretReference|true|true|true|false|
|Driver|string|true|true|true|false|
|FSType|string|true|true|true|false|
|NodePublishSecretRef|secretReference|true|true|true|false|
|NodeStageSecretRef|secretReference|true|true|true|false|
|ReadOnly|boolean|true|true|false|false|
|VolumeAttributes|map[string]|true|true|true|false|
|VolumeHandle|string|true|true|true|false|

### Object `NodeSelectorRequirement`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Key|string|true|true|true|false|
|Operator|string|true|true|true|false|
|Values|array[string]|true|true|true|false|

### Object `NodeSelectorTerm`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MatchExpressions|array[nodeSelectorRequirement]|true|true|true|false|
|MatchFields|array[nodeSelectorRequirement]|true|true|true|false|

### Object `NodeSelector`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|NodeSelectorTerms|array[nodeSelectorTerm]|true|true|true|false|

### Object `VolumeNodeAffinity`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Required|nodeSelector|true|true|true|false|

### Object `PersistentVolumeSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessModes|array[string]|true|true|true|false|
|AWSElasticBlockStore|awsElasticBlockStoreVolumeSource|true|true|true|false|
|AzureDisk|azureDiskVolumeSource|true|true|true|false|
|AzureFile|azureFilePersistentVolumeSource|true|true|true|false|
|Capacity|map[string]|true|true|true|false|
|CephFS|cephFSPersistentVolumeSource|true|true|true|false|
|Cinder|cinderPersistentVolumeSource|true|true|true|false|
|ClaimRef|objectReference|true|true|true|false|
|CSI|csiPersistentVolumeSource|true|true|true|false|
|FC|fcVolumeSource|true|true|true|false|
|FlexVolume|flexPersistentVolumeSource|true|true|true|false|
|Flocker|flockerVolumeSource|true|true|true|false|
|GCEPersistentDisk|gcePersistentDiskVolumeSource|true|true|true|false|
|Glusterfs|glusterfsVolumeSource|true|true|true|false|
|HostPath|hostPathVolumeSource|true|true|true|false|
|ISCSI|iscsiPersistentVolumeSource|true|true|true|false|
|Local|localVolumeSource|true|true|true|false|
|MountOptions|array[string]|true|true|true|false|
|NFS|nfsVolumeSource|true|true|true|false|
|NodeAffinity|volumeNodeAffinity|true|true|true|false|
|PersistentVolumeReclaimPolicy|string|true|true|true|false|
|PhotonPersistentDisk|photonPersistentDiskVolumeSource|true|true|true|false|
|PortworxVolume|portworxVolumeSource|true|true|true|false|
|Quobyte|quobyteVolumeSource|true|true|true|false|
|RBD|rbdPersistentVolumeSource|true|true|true|false|
|ScaleIO|scaleIOPersistentVolumeSource|true|true|true|false|
|StorageClassID|reference[/v3/cluster/storageClass]|true|true|true|false|
|StorageOS|storageOSPersistentVolumeSource|true|true|true|false|
|VolumeMode|string|true|true|true|false|
|VsphereVolume|vsphereVirtualDiskVolumeSource|true|true|true|false|

### Object `LabelSelectorRequirement`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Key|string|true|true|true|false|
|Operator|string|true|true|true|false|
|Values|array[string]|true|true|true|false|

### Object `LabelSelector`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MatchExpressions|array[labelSelectorRequirement]|true|true|true|false|
|MatchLabels|map[string]|true|true|true|false|

### Object `ResourceRequirements`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Limits|map[string]|true|true|true|false|
|Requests|map[string]|true|true|true|false|

### Object `TypedLocalObjectReference`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|APIGroup|string|true|true|true|false|
|Kind|string|true|true|true|false|
|Name|string|true|true|true|false|

### Object `PersistentVolumeClaimSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessModes|array[enum]|true|true|true|false|
|DataSource|typedLocalObjectReference|true|true|true|false|
|Resources|resourceRequirements|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|StorageClassID|reference[/v3/cluster/storageClass]|true|true|true|false|
|VolumeID|reference[/v3/cluster/persistentVolume]|true|true|true|false|
|VolumeMode|string|true|true|true|false|

### Object `OwnerReference`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|APIVersion|string|true|true|true|false|
|BlockOwnerDeletion|boolean|true|true|true|false|
|Controller|boolean|true|true|true|false|
|Kind|string|true|true|true|false|
|Name|string|true|true|true|false|
|UID|string|true|true|true|false|

### Object `Initializer`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|

### Object `ListMeta`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Continue|string|true|true|true|false|
|ResourceVersion|string|true|true|true|false|
|SelfLink|string|true|true|true|false|

### Object `StatusCause`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Field|string|true|true|true|false|
|Message|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `StatusDetails`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Causes|array[statusCause]|true|true|true|false|
|Group|string|true|true|true|false|
|Kind|string|true|true|true|false|
|Name|string|true|true|true|false|
|RetryAfterSeconds|int|true|true|false|false|
|UID|string|true|true|true|false|

### Object `Status`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|APIVersion|string|true|true|true|false|
|Code|int|true|true|false|false|
|Details|statusDetails|true|true|true|false|
|Kind|string|true|true|true|false|
|Message|string|true|true|true|false|
|ListMeta|listMeta|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|

### Object `Initializers`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Pending|array[initializer]|true|true|true|false|
|Result|status|true|true|true|false|

### Object `ObjectMeta`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|Finalizers|array[string]|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|Namespace|string|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|SelfLink|string|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `PersistentVolumeClaimCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastProbeTime|date|true|true|true|false|
|LastTransitionTime|date|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `PersistentVolumeClaimStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessModes|array[string]|false|false|true|false|
|Capacity|map[string]|false|false|true|false|
|Conditions|array[persistentVolumeClaimCondition]|false|false|true|false|
|Phase|string|false|false|true|false|

### Object `PersistentVolumeClaim`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessModes|array[enum]|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DataSource|typedLocalObjectReference|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|Resources|resourceRequirements|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|State|string|false|false|false|false|
|Status|persistentVolumeClaimStatus|false|false|true|false|
|StorageClassID|reference[/v3/cluster/storageClass]|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|
|VolumeID|reference[/v3/cluster/persistentVolume]|true|true|true|false|
|VolumeMode|string|true|true|true|false|

### Object `ConfigMap`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|BinaryData|map[base64]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Data|map[string]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `IngressBackend`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ServiceID|reference[service]|true|true|true|false|
|TargetPort|intOrString|true|true|true|false|
|WorkloadIDs|array[reference[workload]]|true|true|true|false|

### Object `HTTPIngressPath`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ServiceID|reference[service]|true|true|true|false|
|TargetPort|intOrString|true|true|true|false|
|WorkloadIDs|array[reference[workload]]|true|true|true|false|

### Object `HTTPIngressRuleValue`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Paths|map[httpIngressPath]|true|true|true|false|

### Object `IngressRule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Host|string|true|true|true|false|
|Paths|map[ingressBackend]|true|true|true|false|

### Object `IngressTLS`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CertificateID|reference[certificate]|true|true|true|false|
|Hosts|array[string]|true|true|true|false|

### Object `IngressSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Backend|ingressBackend|true|true|true|false|
|Rules|array[ingressRule]|true|true|true|false|
|TLS|array[ingressTLS]|true|true|true|false|

### Object `LoadBalancerIngress`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Hostname|string|true|true|true|false|
|IP|string|true|true|true|false|

### Object `LoadBalancerStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Ingress|array[loadBalancerIngress]|true|true|true|false|

### Object `IngressStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LoadBalancer|loadBalancerStatus|false|false|true|false|

### Object `Ingress`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DefaultBackend|ingressBackend|true|true|true|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|false|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|Removed|date|false|false|true|false|
|Rules|array[ingressRule]|true|true|true|false|
|State|string|false|false|false|false|
|Status|ingressStatus|false|false|true|false|
|TLS|array[ingressTLS]|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `Secret`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Data|map[base64]|true|true|true|false|
|Description|string|true|true|true|false|
|Kind|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|StringData|map[string]|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `ServiceAccountToken`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccountName|string|true|true|true|false|
|AccountUID|string|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|CACRT|string|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|Token|string|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `RegistryCredential`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Auth|string|true|true|true|false|
|Description|string|true|true|true|false|
|Password|string|true|true|true|false|
|Username|string|true|true|true|false|

### Object `DockerCredential`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Registries|map[registryCredential]|true|true|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `Certificate`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Algorithm|string|false|false|true|false|
|Annotations|map[string]|true|true|true|false|
|CertFingerprint|string|false|false|true|false|
|Certs|string|true|true|true|false|
|CN|string|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|ExpiresAt|string|false|false|true|false|
|IssuedAt|string|false|false|true|false|
|Issuer|string|false|false|true|false|
|Key|string|true|true|true|false|
|KeySize|string|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|SerialNumber|string|false|false|true|false|
|SubjectAlternativeNames|array[string]|false|false|true|false|
|UUID|string|false|false|true|false|
|Version|string|false|false|true|false|

### Object `BasicAuth`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Password|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|Username|string|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `SSHAuth`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Fingerprint|string|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PrivateKey|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `NamespacedSecret`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Data|map[base64]|true|true|true|false|
|Description|string|true|true|true|false|
|Kind|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|StringData|map[string]|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `NamespacedServiceAccountToken`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccountName|string|true|true|true|false|
|AccountUID|string|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|CACRT|string|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|Token|string|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `NamespacedDockerCredential`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Registries|map[registryCredential]|true|true|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `NamespacedCertificate`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Algorithm|string|false|false|true|false|
|Annotations|map[string]|true|true|true|false|
|CertFingerprint|string|false|false|true|false|
|Certs|string|true|true|true|false|
|CN|string|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|ExpiresAt|string|false|false|true|false|
|IssuedAt|string|false|false|true|false|
|Issuer|string|false|false|true|false|
|Key|string|true|true|true|false|
|KeySize|string|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|SerialNumber|string|false|false|true|false|
|SubjectAlternativeNames|array[string]|false|false|true|false|
|UUID|string|false|false|true|false|
|Version|string|false|false|true|false|

### Object `NamespacedBasicAuth`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Password|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|Username|string|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `NamespacedSSHAuth`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Fingerprint|string|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PrivateKey|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `Service`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClusterIp|string|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|ExternalIPs|array[string]|true|true|true|false|
|ExternalTrafficPolicy|string|true|true|true|false|
|HealthCheckNodePort|int|true|true|false|false|
|Hostname|string|true|true|true|false|
|IPAddresses|array[string]|true|true|true|false|
|Kind|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|LoadBalancerIP|string|true|true|true|false|
|LoadBalancerSourceRanges|array[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Ports|array[servicePort]|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|PublishNotReadyAddresses|boolean|true|true|false|false|
|Removed|date|false|false|true|false|
|Selector|map[string]|true|true|true|false|
|SessionAffinity|string|true|true|true|false|
|SessionAffinityConfig|sessionAffinityConfig|true|true|true|false|
|State|string|false|false|false|false|
|TargetDNSRecordIDs|array[reference[dnsRecord]]|true|true|true|false|
|TargetWorkloadIDs|array[reference[workload]]|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|
|WorkloadID|reference[workload]|false|false|true|false|

### Object `ServicePort`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|
|NodePort|int|true|true|false|false|
|Port|int|true|true|false|false|
|Protocol|string|true|true|true|false|
|TargetPort|intOrString|true|true|true|false|

### Object `ClientIPConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|TimeoutSeconds|int|true|true|true|false|

### Object `SessionAffinityConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClientIP|clientIPConfig|true|true|true|false|

### Object `ServiceSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClusterIp|string|true|true|true|false|
|ExternalIPs|array[string]|true|true|true|false|
|ExternalTrafficPolicy|string|true|true|true|false|
|HealthCheckNodePort|int|true|true|false|false|
|Hostname|string|true|true|true|false|
|LoadBalancerIP|string|true|true|true|false|
|LoadBalancerSourceRanges|array[string]|true|true|true|false|
|Ports|array[servicePort]|true|true|true|false|
|PublishNotReadyAddresses|boolean|true|true|false|false|
|Selector|map[string]|true|true|true|false|
|ServiceKind|string|true|true|true|false|
|SessionAffinity|string|true|true|true|false|
|SessionAffinityConfig|sessionAffinityConfig|true|true|true|false|

### Object `ServiceStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LoadBalancer|loadBalancerStatus|false|false|true|false|

### Object `DNSRecord`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClusterIp|string|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|Hostname|string|true|true|true|false|
|IPAddresses|array[string]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabelRestricted|true|false|false|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Ports|array[servicePort]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|Removed|date|false|false|true|false|
|Selector|map[string]|true|true|true|false|
|State|string|false|false|false|false|
|TargetDNSRecordIDs|array[reference[dnsRecord]]|true|true|true|false|
|TargetWorkloadIDs|array[reference[workload]]|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|
|WorkloadID|reference[workload]|false|false|true|false|

### Object `ContainerPort`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ContainerPort|int|true|true|false|false|
|DNSName|string|true|true|true|false|
|HostIp|string|true|true|true|false|
|Kind|enum|true|true|true|false|
|Name|string|true|true|true|false|
|Protocol|string|true|true|true|false|
|SourcePort|int|true|true|false|false|

### Object `Capabilities`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CapAdd|array[enum]|true|true|true|false|
|CapDrop|array[enum]|true|true|true|false|

### Object `PublicEndpoint`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Addresses|array[string]|false|false|true|false|
|AllNodes|boolean|false|false|false|false|
|Hostname|string|false|false|true|false|
|IngressID|reference[ingress]|false|false|true|false|
|NodeID|reference[/v3/schemas/node]|false|false|true|false|
|Path|string|false|false|true|false|
|PodID|reference[pod]|false|false|true|false|
|Port|int|false|false|false|false|
|Protocol|string|false|false|true|false|
|ServiceID|reference[service]|false|false|true|false|

### Object `WorkloadMetric`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Path|string|true|true|true|false|
|Port|int|true|true|false|false|
|Schema|enum|true|true|true|false|

### Object `ExecAction`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Command|array[string]|true|true|true|false|

### Object `HTTPHeader`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|
|Value|string|true|true|true|false|

### Object `HTTPGetAction`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|HTTPHeaders|array[httpHeader]|true|true|true|false|
|Path|string|true|true|true|false|
|Port|intOrString|true|true|true|false|
|Scheme|string|true|true|true|false|

### Object `TCPSocketAction`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Host|string|true|true|true|false|
|Port|intOrString|true|true|true|false|

### Object `Handler`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Command|array[string]|true|true|true|false|
|Host|string|true|true|true|false|
|HTTPHeaders|array[httpHeader]|true|true|true|false|
|Path|string|true|true|true|false|
|Port|intOrString|true|true|true|false|
|Scheme|string|true|true|true|false|
|TCP|boolean|true|true|false|false|

### Object `Probe`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Command|array[string]|true|true|true|false|
|FailureThreshold|int|true|true|false|false|
|Host|string|true|true|true|false|
|HTTPHeaders|array[httpHeader]|true|true|true|false|
|InitialDelaySeconds|int|true|true|false|false|
|Path|string|true|true|true|false|
|PeriodSeconds|int|true|true|false|false|
|Port|intOrString|true|true|true|false|
|Scheme|string|true|true|true|false|
|SuccessThreshold|int|true|true|false|false|
|TCP|boolean|true|true|false|false|
|TimeoutSeconds|int|true|true|false|false|

### Object `ConfigMapEnvSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|
|Optional|boolean|true|true|true|false|

### Object `SecretEnvSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|
|Optional|boolean|true|true|true|false|

### Object `EnvFromSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ConfigMapRef|configMapEnvSource|true|true|true|false|
|Prefix|string|true|true|true|false|
|SecretRef|secretEnvSource|true|true|true|false|

### Object `ConfigMapKeySelector`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Key|string|true|true|true|false|
|Name|string|true|true|true|false|
|Optional|boolean|true|true|true|false|

### Object `SecretKeySelector`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Key|string|true|true|true|false|
|Name|string|true|true|true|false|
|Optional|boolean|true|true|true|false|

### Object `EnvVarSource`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ConfigMapKeyRef|configMapKeySelector|true|true|true|false|
|FieldRef|objectFieldSelector|true|true|true|false|
|ResourceFieldRef|resourceFieldSelector|true|true|true|false|
|SecretKeyRef|secretKeySelector|true|true|true|false|

### Object `EnvVar`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|
|Value|string|true|true|true|false|
|ValueFrom|envVarSource|true|true|true|false|

### Object `VolumeDevice`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DevicePath|string|true|true|true|false|
|Name|string|true|true|true|false|

### Object `Lifecycle`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|PostStart|handler|true|true|true|false|
|PreStop|handler|true|true|true|false|

### Object `SELinuxOptions`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Level|string|true|true|true|false|
|Role|string|true|true|true|false|
|Type|string|true|true|true|false|
|User|string|true|true|true|false|

### Object `SecurityContext`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AllowPrivilegeEscalation|boolean|true|true|true|false|
|CapAdd|array[enum]|true|true|true|false|
|CapDrop|array[enum]|true|true|true|false|
|Privileged|boolean|true|true|true|false|
|ProcMount|string|true|true|true|false|
|ReadOnly|boolean|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|Uid|int|true|true|true|false|

### Object `EnvironmentFrom`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Optional|boolean|true|true|false|false|
|Prefix|string|true|true|true|false|
|Source|enum|true|true|true|false|
|SourceKey|string|true|true|true|false|
|SourceName|string|true|true|true|false|
|TargetKey|string|true|true|true|false|

### Object `Container`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AllowPrivilegeEscalation|boolean|true|true|true|false|
|CapAdd|array[enum]|true|true|true|false|
|CapDrop|array[enum]|true|true|true|false|
|Command|array[string]|true|true|true|false|
|Entrypoint|array[string]|true|true|true|false|
|Environment|map[string]|true|true|true|false|
|EnvironmentFrom|array[environmentFrom]|true|true|true|false|
|ExitCode|int|true|true|true|false|
|Image|string|true|true|true|false|
|ImagePullPolicy|string|true|true|true|false|
|InitContainer|boolean|true|true|false|false|
|LivenessProbe|probe|true|true|true|false|
|Name|string|true|true|true|false|
|Ports|array[containerPort]|true|true|true|false|
|PostStart|handler|true|true|true|false|
|PreStop|handler|true|true|true|false|
|Privileged|boolean|true|true|true|false|
|ProcMount|string|true|true|true|false|
|ReadOnly|boolean|true|true|true|false|
|ReadinessProbe|probe|true|true|true|false|
|Resources|resourceRequirements|true|true|true|false|
|RestartCount|int|true|true|false|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|State|string|true|true|true|false|
|Stdin|boolean|true|true|false|false|
|StdinOnce|boolean|true|true|false|false|
|TerminationMessagePath|string|true|true|true|false|
|TerminationMessagePolicy|string|true|true|true|false|
|Transitioning|string|true|true|true|false|
|TransitioningMessage|string|true|true|true|false|
|TTY|boolean|true|true|false|false|
|Uid|int|true|true|true|false|
|VolumeDevices|array[volumeDevice]|true|true|true|false|
|VolumeMounts|array[volumeMount]|true|true|true|false|
|WorkingDir|string|true|true|true|false|

### Object `Sysctl`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|
|Value|string|true|true|true|false|

### Object `PodSecurityContext`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|Uid|int|true|true|true|false|

### Object `PreferredSchedulingTerm`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Preference|nodeSelectorTerm|true|true|true|false|
|Weight|int|true|true|false|false|

### Object `NodeAffinity`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|PreferredDuringSchedulingIgnoredDuringExecution|array[preferredSchedulingTerm]|true|true|true|false|
|RequiredDuringSchedulingIgnoredDuringExecution|nodeSelector|true|true|true|false|

### Object `PodAffinityTerm`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LabelSelector|labelSelector|true|true|true|false|
|Namespaces|array[string]|true|true|true|false|
|TopologyKey|string|true|true|true|false|

### Object `WeightedPodAffinityTerm`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|PodAffinityTerm|podAffinityTerm|true|true|true|false|
|Weight|int|true|true|false|false|

### Object `PodAffinity`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|PreferredDuringSchedulingIgnoredDuringExecution|array[weightedPodAffinityTerm]|true|true|true|false|
|RequiredDuringSchedulingIgnoredDuringExecution|array[podAffinityTerm]|true|true|true|false|

### Object `PodAntiAffinity`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|PreferredDuringSchedulingIgnoredDuringExecution|array[weightedPodAffinityTerm]|true|true|true|false|
|RequiredDuringSchedulingIgnoredDuringExecution|array[podAffinityTerm]|true|true|true|false|

### Object `Affinity`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|NodeAffinity|nodeAffinity|true|true|true|false|
|PodAffinity|podAffinity|true|true|true|false|
|PodAntiAffinity|podAntiAffinity|true|true|true|false|

### Object `Toleration`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Effect|string|true|true|true|false|
|Key|string|true|true|true|false|
|Operator|string|true|true|true|false|
|TolerationSeconds|int|true|true|true|false|
|Value|string|true|true|true|false|

### Object `HostAlias`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Hostnames|array[string]|true|true|true|false|
|IP|string|true|true|true|false|

### Object `PodDNSConfigOption`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|false|
|Value|string|true|true|true|false|

### Object `PodDNSConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Nameservers|array[string]|true|true|true|false|
|Options|array[podDNSConfigOption]|true|true|true|false|
|Searches|array[string]|true|true|true|false|

### Object `PodReadinessGate`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ConditionType|string|true|true|true|false|

### Object `NodeScheduling`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|Preferred|array[string]|true|true|true|false|
|RequireAll|array[string]|true|true|true|false|
|RequireAny|array[string]|true|true|true|false|

### Object `Scheduling`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Node|nodeScheduling|true|true|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|Scheduler|string|true|true|true|false|
|Tolerate|array[toleration]|true|true|true|false|

### Object `PodSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Uid|int|true|true|true|false|
|Volumes|array[volume]|true|true|true|false|

### Object `PodCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastProbeTime|date|true|true|true|false|
|LastTransitionTime|date|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `ContainerStateWaiting`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|

### Object `ContainerStateRunning`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|StartedAt|date|true|true|true|false|

### Object `ContainerStateTerminated`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ContainerID|string|true|true|true|false|
|ExitCode|int|true|true|false|false|
|FinishedAt|date|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Signal|int|true|true|false|false|
|StartedAt|date|true|true|true|false|

### Object `ContainerState`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Running|containerStateRunning|true|true|true|false|
|Terminated|containerStateTerminated|true|true|true|false|
|Waiting|containerStateWaiting|true|true|true|false|

### Object `ContainerStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ContainerID|string|true|true|true|false|
|Image|string|true|true|true|false|
|ImageID|string|true|true|true|false|
|LastTerminationState|containerState|true|true|true|false|
|Name|string|true|true|true|false|
|Ready|boolean|true|true|false|false|
|RestartCount|int|true|true|false|false|
|State|containerState|true|true|true|false|

### Object `PodStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Conditions|array[podCondition]|false|false|true|false|
|ContainerStatuses|array[containerStatus]|false|false|true|false|
|InitContainerStatuses|array[containerStatus]|false|false|true|false|
|Message|string|false|false|true|false|
|NodeIp|string|false|false|true|false|
|NominatedNodeName|string|false|false|true|false|
|Phase|string|false|false|true|false|
|PodIp|string|false|false|true|false|
|QOSClass|string|false|false|true|false|
|Reason|string|false|false|true|false|
|StartTime|date|false|false|true|false|

### Object `Pod`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|Removed|date|false|false|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|State|string|false|false|false|false|
|Status|podStatus|false|false|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|Uid|int|true|true|true|false|
|UUID|string|false|false|true|false|
|Volumes|array[volume]|true|true|true|false|
|WorkloadID|reference[workload]|true|true|true|false|
|WorkloadMetrics|array[workloadMetric]|false|false|true|false|

### Object `PodTemplateSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|ObjectMeta|objectMeta|true|true|true|false|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Uid|int|true|true|true|false|
|Volumes|array[volume]|true|true|true|false|

### Object `RollingUpdateDeployment`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MaxSurge|intOrString|true|true|true|false|
|MaxUnavailable|intOrString|true|true|true|false|

### Object `DeploymentStrategy`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MaxSurge|intOrString|true|true|true|false|
|MaxUnavailable|intOrString|true|true|true|false|
|Strategy|enum|true|true|true|false|

### Object `DeploymentConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MaxSurge|intOrString|true|true|true|false|
|MaxUnavailable|intOrString|true|true|true|false|
|MinReadySeconds|int|true|true|false|false|
|ProgressDeadlineSeconds|int|true|true|true|false|
|RevisionHistoryLimit|int|true|true|true|false|
|Strategy|enum|true|true|true|false|

### Object `DeploymentSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|DeploymentConfig|deploymentConfig|true|true|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|ObjectMeta|objectMeta|true|true|true|false|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|Paused|boolean|true|true|false|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|Scale|int|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Uid|int|true|true|true|false|
|Volumes|array[volume]|true|true|true|false|

### Object `DeploymentRollbackInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ReplicaSetID|reference[replicaSet]|true|true|true|false|

### Object `DeploymentCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|date|true|true|true|false|
|LastUpdateTime|date|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `DeploymentStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AvailableReplicas|int|false|false|false|false|
|CollisionCount|int|false|false|true|false|
|Conditions|array[deploymentCondition]|false|false|true|false|
|ObservedGeneration|int|false|false|false|false|
|ReadyReplicas|int|false|false|false|false|
|Replicas|int|false|false|false|false|
|UnavailableReplicas|int|false|false|false|false|
|UpdatedReplicas|int|false|false|false|false|

### Object `Deployment`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DeploymentConfig|deploymentConfig|true|true|true|false|
|DeploymentStatus|deploymentStatus|false|false|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabelRestricted|true|false|false|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Paused|boolean|true|true|false|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|Removed|date|false|false|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|Scale|int|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|State|string|false|false|false|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|Uid|int|true|true|true|false|
|UUID|string|false|false|true|false|
|Volumes|array[volume]|true|true|true|false|
|WorkloadAnnotations|map[string]|true|true|true|false|
|WorkloadLabels|map[string]|true|true|true|false|
|WorkloadMetrics|array[workloadMetric]|true|true|true|false|

### Object `ReplicationControllerConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MinReadySeconds|int|true|true|false|false|

### Object `ReplicationControllerSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|ObjectMeta|objectMeta|true|true|true|false|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|ReplicationControllerConfig|replicationControllerConfig|true|true|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|Scale|int|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|map[string]|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Uid|int|true|true|true|false|
|Volumes|array[volume]|true|true|true|false|

### Object `ReplicationControllerCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|date|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `ReplicationControllerStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AvailableReplicas|int|false|false|false|false|
|Conditions|array[replicationControllerCondition]|false|false|true|false|
|FullyLabeledReplicas|int|false|false|false|false|
|ObservedGeneration|int|false|false|false|false|
|ReadyReplicas|int|false|false|false|false|
|Replicas|int|false|false|false|false|

### Object `ReplicationController`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabelRestricted|true|false|false|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|Removed|date|false|false|true|false|
|ReplicationControllerConfig|replicationControllerConfig|true|true|true|false|
|ReplicationControllerStatus|replicationControllerStatus|false|false|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|Scale|int|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|map[string]|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|State|string|false|false|false|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|Uid|int|true|true|true|false|
|UUID|string|false|false|true|false|
|Volumes|array[volume]|true|true|true|false|
|WorkloadAnnotations|map[string]|true|true|true|false|
|WorkloadLabels|map[string]|true|true|true|false|
|WorkloadMetrics|array[workloadMetric]|true|true|true|false|

### Object `ReplicaSetConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MinReadySeconds|int|true|true|false|false|

### Object `ReplicaSetSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ReplicaSetConfig|replicaSetConfig|true|true|true|false|
|Scale|int|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|Template|podTemplateSpec|true|true|true|false|

### Object `ReplicaSetCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|date|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `ReplicaSetStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AvailableReplicas|int|false|false|false|false|
|Conditions|array[replicaSetCondition]|false|false|true|false|
|FullyLabeledReplicas|int|false|false|false|false|
|ObservedGeneration|int|false|false|false|false|
|ReadyReplicas|int|false|false|false|false|
|Replicas|int|false|false|false|false|

### Object `ReplicaSet`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabelRestricted|true|false|false|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|Removed|date|false|false|true|false|
|ReplicaSetConfig|replicaSetConfig|true|true|true|false|
|ReplicaSetStatus|replicaSetStatus|false|false|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|Scale|int|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|State|string|false|false|false|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|Uid|int|true|true|true|false|
|UUID|string|false|false|true|false|
|Volumes|array[volume]|true|true|true|false|
|WorkloadAnnotations|map[string]|true|true|true|false|
|WorkloadLabels|map[string]|true|true|true|false|
|WorkloadMetrics|array[workloadMetric]|true|true|true|false|

### Object `RollingUpdateStatefulSetStrategy`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Partition|int|true|true|true|false|

### Object `StatefulSetUpdateStrategy`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Partition|int|true|true|true|false|
|Strategy|enum|true|true|true|false|

### Object `StatefulSetConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Partition|int|true|true|true|false|
|PodManagementPolicy|enum|true|true|true|false|
|RevisionHistoryLimit|int|true|true|true|false|
|ServiceName|string|true|true|true|false|
|Strategy|enum|true|true|true|false|
|VolumeClaimTemplates|array[persistentVolumeClaim]|true|true|true|false|

### Object `StatefulSetSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|ObjectMeta|objectMeta|true|true|true|false|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|Scale|int|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|StatefulSetConfig|statefulSetConfig|true|true|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Uid|int|true|true|true|false|
|Volumes|array[volume]|true|true|true|false|

### Object `StatefulSetCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|date|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `StatefulSetStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CollisionCount|int|false|false|true|false|
|Conditions|array[statefulSetCondition]|false|false|true|false|
|CurrentReplicas|int|false|false|false|false|
|CurrentRevision|string|false|false|true|false|
|ObservedGeneration|int|false|false|false|false|
|ReadyReplicas|int|false|false|false|false|
|Replicas|int|false|false|false|false|
|UpdateRevision|string|false|false|true|false|
|UpdatedReplicas|int|false|false|false|false|

### Object `StatefulSet`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabelRestricted|true|false|false|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|Removed|date|false|false|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|Scale|int|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|State|string|false|false|false|false|
|StatefulSetConfig|statefulSetConfig|true|true|true|false|
|StatefulSetStatus|statefulSetStatus|false|false|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|Uid|int|true|true|true|false|
|UUID|string|false|false|true|false|
|Volumes|array[volume]|true|true|true|false|
|WorkloadAnnotations|map[string]|true|true|true|false|
|WorkloadLabels|map[string]|true|true|true|false|
|WorkloadMetrics|array[workloadMetric]|true|true|true|false|

### Object `RollingUpdateDaemonSet`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MaxUnavailable|intOrString|true|true|true|false|

### Object `DaemonSetUpdateStrategy`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MaxUnavailable|intOrString|true|true|true|false|
|Strategy|enum|true|true|true|false|

### Object `DaemonSetConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|MaxUnavailable|intOrString|true|true|true|false|
|MinReadySeconds|int|true|true|false|false|
|RevisionHistoryLimit|int|true|true|true|false|
|Strategy|enum|true|true|true|false|

### Object `DaemonSetSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|DaemonSetConfig|daemonSetConfig|true|true|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|ObjectMeta|objectMeta|true|true|true|false|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Uid|int|true|true|true|false|
|Volumes|array[volume]|true|true|true|false|

### Object `DaemonSetCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|date|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `DaemonSetStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CollisionCount|int|false|false|true|false|
|Conditions|array[daemonSetCondition]|false|false|true|false|
|CurrentNumberScheduled|int|false|false|false|false|
|DesiredNumberScheduled|int|false|false|false|false|
|NumberAvailable|int|false|false|false|false|
|NumberMisscheduled|int|false|false|false|false|
|NumberReady|int|false|false|false|false|
|NumberUnavailable|int|false|false|false|false|
|ObservedGeneration|int|false|false|false|false|
|UpdatedNumberScheduled|int|false|false|false|false|

### Object `DaemonSet`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DaemonSetConfig|daemonSetConfig|true|true|true|false|
|DaemonSetStatus|daemonSetStatus|false|false|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabelRestricted|true|false|false|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|Removed|date|false|false|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|State|string|false|false|false|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|Uid|int|true|true|true|false|
|UUID|string|false|false|true|false|
|Volumes|array[volume]|true|true|true|false|
|WorkloadAnnotations|map[string]|true|true|true|false|
|WorkloadLabels|map[string]|true|true|true|false|
|WorkloadMetrics|array[workloadMetric]|true|true|true|false|

### Object `JobConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|BackoffLimit|int|true|true|true|false|
|Completions|int|true|true|true|false|
|ManualSelector|boolean|true|true|true|false|
|Parallelism|int|true|true|true|false|

### Object `JobSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|JobConfig|jobConfig|true|true|true|false|
|ObjectMeta|objectMeta|true|true|true|false|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|TTLSecondsAfterFinished|int|true|true|true|false|
|Uid|int|true|true|true|false|
|Volumes|array[volume]|true|true|true|false|

### Object `JobCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastProbeTime|date|true|true|true|false|
|LastTransitionTime|date|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `JobStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Active|int|false|false|false|false|
|CompletionTime|date|false|false|true|false|
|Conditions|array[jobCondition]|false|false|true|false|
|Failed|int|false|false|false|false|
|StartTime|date|false|false|true|false|
|Succeeded|int|false|false|false|false|

### Object `Job`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|JobConfig|jobConfig|true|true|true|false|
|JobStatus|jobStatus|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabelRestricted|true|false|false|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|Removed|date|false|false|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|State|string|false|false|false|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|TTLSecondsAfterFinished|int|true|true|true|false|
|Uid|int|true|true|true|false|
|UUID|string|false|false|true|false|
|Volumes|array[volume]|true|true|true|false|
|WorkloadAnnotations|map[string]|true|true|true|false|
|WorkloadLabels|map[string]|true|true|true|false|
|WorkloadMetrics|array[workloadMetric]|true|true|true|false|

### Object `JobTemplateSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|JobConfig|jobConfig|true|true|true|false|
|JobMetadata|objectMeta|true|true|true|false|
|ObjectMeta|objectMeta|true|true|true|false|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|TTLSecondsAfterFinished|int|true|true|true|false|
|Uid|int|true|true|true|false|
|Volumes|array[volume]|true|true|true|false|

### Object `CronJobConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ConcurrencyPolicy|enum|true|true|true|false|
|FailedJobsHistoryLimit|int|true|true|true|false|
|JobAnnotations|map[string]|true|true|true|false|
|JobConfig|jobConfig|true|true|true|false|
|JobLabels|map[string]|true|true|true|false|
|Schedule|string|true|true|true|false|
|StartingDeadlineSeconds|int|true|true|true|false|
|SuccessfulJobsHistoryLimit|int|true|true|true|false|
|Suspend|boolean|true|true|true|false|

### Object `CronJobSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|CronJobConfig|cronJobConfig|true|true|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|ObjectMeta|objectMeta|true|true|true|false|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|TTLSecondsAfterFinished|int|true|true|true|false|
|Uid|int|true|true|true|false|
|Volumes|array[volume]|true|true|true|false|

### Object `CronJobStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Active|array[objectReference]|false|false|true|false|
|LastScheduleTime|date|false|false|true|false|

### Object `CronJob`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|CronJobConfig|cronJobConfig|true|true|true|false|
|CronJobStatus|cronJobStatus|false|false|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabelRestricted|true|false|false|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|Removed|date|false|false|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|State|string|false|false|false|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|TTLSecondsAfterFinished|int|true|true|true|false|
|Uid|int|true|true|true|false|
|UUID|string|false|false|true|false|
|Volumes|array[volume]|true|true|true|false|
|WorkloadAnnotations|map[string]|true|true|true|false|
|WorkloadLabels|map[string]|true|true|true|false|
|WorkloadMetrics|array[workloadMetric]|true|true|true|false|

### Object `Workload`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ActiveDeadlineSeconds|int|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|AutomountServiceAccountToken|boolean|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|CronJobConfig|cronJobConfig|true|true|true|false|
|CronJobStatus|cronJobStatus|false|false|true|false|
|DaemonSetConfig|daemonSetConfig|true|true|true|false|
|DaemonSetStatus|daemonSetStatus|false|false|true|false|
|DeploymentConfig|deploymentConfig|true|true|true|false|
|DeploymentStatus|deploymentStatus|false|false|true|false|
|DNSConfig|podDNSConfig|true|true|true|false|
|DNSPolicy|string|true|true|true|false|
|Fsgid|int|true|true|true|false|
|Gids|array[int]|true|true|true|false|
|HostAliases|array[hostAlias]|true|true|true|false|
|HostIPC|boolean|true|true|false|false|
|HostNetwork|boolean|true|true|false|false|
|HostPID|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|JobConfig|jobConfig|true|true|true|false|
|JobStatus|jobStatus|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabelRestricted|true|false|false|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeID|reference[/v3/schemas/node]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Paused|boolean|true|true|false|false|
|Priority|int|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|PublicEndpoints|array[publicEndpoint]|false|false|true|false|
|ReadinessGates|array[podReadinessGate]|true|true|true|false|
|Removed|date|false|false|true|false|
|ReplicaSetConfig|replicaSetConfig|true|true|true|false|
|ReplicaSetStatus|replicaSetStatus|false|false|true|false|
|ReplicationControllerConfig|replicationControllerConfig|true|true|true|false|
|ReplicationControllerStatus|replicationControllerStatus|false|false|true|false|
|RestartPolicy|string|true|true|true|false|
|RunAsGroup|int|true|true|true|false|
|RunAsNonRoot|boolean|true|true|true|false|
|RuntimeClassName|string|true|true|true|false|
|Scale|int|true|true|true|false|
|SchedulerName|string|true|true|true|false|
|Scheduling|scheduling|true|true|true|false|
|Selector|labelSelector|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ShareProcessNamespace|boolean|true|true|true|false|
|State|string|false|false|false|false|
|StatefulSetConfig|statefulSetConfig|true|true|true|false|
|StatefulSetStatus|statefulSetStatus|false|false|true|false|
|Subdomain|string|true|true|true|false|
|Sysctls|array[sysctl]|true|true|true|false|
|TerminationGracePeriodSeconds|int|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|TTLSecondsAfterFinished|int|true|true|true|false|
|Uid|int|true|true|true|false|
|UUID|string|false|false|true|false|
|Volumes|array[volume]|true|true|true|false|
|WorkloadAnnotations|map[string]|true|true|true|false|
|WorkloadLabels|map[string]|true|true|true|false|
|WorkloadMetrics|array[workloadMetric]|true|true|true|false|

### Object `AppUpgradeConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Answers|map[string]|true|true|true|false|
|ExternalID|string|true|true|true|false|

### Object `RollbackRevision`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|RevisionID|reference[/v3/project/schemas/apprevision]|true|true|true|false|

### Object `AppSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Answers|map[string]|true|true|true|false|
|AppRevisionID|reference[/v3/project/schemas/apprevision]|true|true|true|false|
|Description|string|true|true|true|false|
|ExternalID|string|true|true|true|false|
|Files|map[string]|true|true|true|false|
|MultiClusterAppID|reference[/v3/schemas/multiclusterapp]|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|true|true|false|
|Prune|boolean|true|true|false|false|
|TargetNamespace|string|true|true|true|false|

### Object `AppCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|string|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `AppStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AppliedFiles|map[string]|false|false|true|false|
|Conditions|array[appCondition]|false|false|true|false|
|LastAppliedTemplates|string|false|false|true|false|
|Notes|string|false|false|true|false|

### Object `App`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Answers|map[string]|true|true|true|false|
|AppRevisionID|reference[/v3/project/schemas/apprevision]|true|true|true|false|
|AppliedFiles|map[string]|false|false|true|false|
|Conditions|array[appCondition]|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|ExternalID|string|true|true|true|false|
|Files|map[string]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|LastAppliedTemplates|string|false|false|true|false|
|MultiClusterAppID|reference[/v3/schemas/multiclusterapp]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|Notes|string|false|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|true|true|false|
|Prune|boolean|true|true|false|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|TargetNamespace|string|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `AppRevisionSpec`
--- 


### Object `AppRevisionStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Answers|map[string]|false|false|true|false|
|Digest|string|false|false|true|false|
|ExternalID|string|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|false|false|true|false|

### Object `AppRevision`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Removed|date|false|false|true|false|
|State|string|false|false|false|false|
|Status|appRevisionStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|

### Object `AuthAppInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClientID|string|true|true|true|true|
|ClientSecret|string|true|true|true|true|
|Code|string|true|true|true|true|
|Host|string|true|true|true|false|
|InheritGlobal|boolean|true|true|false|false|
|RedirectURL|string|true|true|true|false|
|SourceCodeType|string|true|true|true|true|
|TLS|boolean|true|true|false|false|

### Object `AuthUserInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Code|string|true|true|true|true|
|RedirectURL|string|true|true|true|false|
|SourceCodeType|string|true|true|true|true|

### Object `RunPipelineInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Branch|string|true|true|true|false|

### Object `SourceCodeConfig`
--- 


### Object `RunScriptConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Image|string|true|true|true|true|
|ShellScript|string|true|true|true|false|

### Object `PublishImageConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BuildContext|string|true|true|true|true|
|DockerfilePath|string|true|true|true|true|
|PushRemote|boolean|true|true|false|false|
|Registry|string|true|true|true|false|
|Tag|string|true|true|true|true|

### Object `ApplyYamlConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Content|string|true|true|true|false|
|Namespace|string|true|true|true|false|
|Path|string|true|true|true|false|

### Object `EnvFrom`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|SourceKey|string|true|true|true|true|
|SourceName|string|true|true|true|true|
|TargetKey|string|true|true|true|false|

### Object `Constraint`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Exclude|array[string]|true|true|true|false|
|Include|array[string]|true|true|true|false|

### Object `Constraints`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Branch|constraint|true|true|true|false|
|Event|constraint|true|true|true|false|

### Object `Step`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ApplyYamlConfig|applyYamlConfig|true|true|true|false|
|Env|map[string]|true|true|true|false|
|EnvFrom|array[envFrom]|true|true|true|false|
|Privileged|boolean|true|true|false|false|
|PublishImageConfig|publishImageConfig|true|true|true|false|
|RunScriptConfig|runScriptConfig|true|true|true|false|
|SourceCodeConfig|sourceCodeConfig|true|true|true|false|
|When|constraints|true|true|true|false|

### Object `Stage`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Name|string|true|true|true|true|
|Steps|array[step]|true|true|true|true|
|When|constraints|true|true|true|false|

### Object `Recipient`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Notifier|string|true|true|true|false|
|Recipient|string|true|true|true|false|

### Object `PipelineNotification`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Condition|array[string]|true|true|true|false|
|Message|string|true|true|true|false|
|Recipients|array[recipient]|true|true|true|false|

### Object `PipelineConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Branch|constraint|true|true|true|false|
|Notification|pipelineNotification|true|true|true|false|
|Stages|array[stage]|true|true|true|false|
|Timeout|int|true|true|false|false|

### Object `PushPipelineConfigInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Configs|map[pipelineConfig]|true|true|true|false|

### Object `GithubApplyInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClientID|string|true|true|true|false|
|ClientSecret|string|true|true|true|false|
|Code|string|true|true|true|false|
|Hostname|string|true|true|true|false|
|InheritAuth|boolean|true|true|false|false|
|RedirectURL|string|true|true|true|false|
|TLS|boolean|true|true|false|false|

### Object `GitlabApplyInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClientID|string|true|true|true|false|
|ClientSecret|string|true|true|true|false|
|Code|string|true|true|true|false|
|Hostname|string|true|true|true|false|
|RedirectURL|string|true|true|true|false|
|TLS|boolean|true|true|false|false|

### Object `BitbucketCloudApplyInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|ClientID|string|true|true|true|false|
|ClientSecret|string|true|true|true|false|
|Code|string|true|true|true|false|
|Hostname|string|true|true|true|false|
|RedirectURL|string|true|true|true|false|
|TLS|boolean|true|true|false|false|

### Object `BitbucketServerApplyInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Hostname|string|true|true|true|false|
|OAuthToken|string|true|true|true|false|
|OAuthVerifier|string|true|true|true|false|
|RedirectURL|string|true|true|true|false|
|TLS|boolean|true|true|false|false|

### Object `BitbucketServerRequestLoginInput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Hostname|string|true|true|true|false|
|RedirectURL|string|true|true|true|false|
|TLS|boolean|true|true|false|false|

### Object `BitbucketServerRequestLoginOutput`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LoginURL|string|true|true|true|false|

### Object `SourceCodeProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|false|
|Removed|date|false|false|true|false|
|Type|enum|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `GithubProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|false|
|RedirectURL|string|true|true|true|false|
|Removed|date|false|false|true|false|
|Type|enum|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `GitlabProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|false|
|RedirectURL|string|true|true|true|false|
|Removed|date|false|false|true|false|
|Type|enum|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `BitbucketCloudProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|false|
|RedirectURL|string|true|true|true|false|
|Removed|date|false|false|true|false|
|Type|enum|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `BitbucketServerProvider`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|false|
|RedirectURL|string|true|true|true|false|
|Removed|date|false|false|true|false|
|Type|enum|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `SourceCodeProviderConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|true|
|Removed|date|false|false|true|false|
|Type|enum|true|false|true|false|
|UUID|string|false|false|true|false|

### Object `GithubPipelineConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClientID|string|true|false|true|false|
|ClientSecret|password|true|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|Inherit|boolean|true|false|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|true|
|Removed|date|false|false|true|false|
|TLS|boolean|true|true|false|false|
|Type|enum|true|false|true|false|
|UUID|string|false|false|true|false|

### Object `GitlabPipelineConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClientID|string|true|false|true|false|
|ClientSecret|password|true|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|true|
|RedirectURL|string|true|false|true|false|
|Removed|date|false|false|true|false|
|TLS|boolean|true|true|false|false|
|Type|enum|true|false|true|false|
|UUID|string|false|false|true|false|

### Object `BitbucketCloudPipelineConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ClientID|string|true|false|true|false|
|ClientSecret|password|true|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|true|
|RedirectURL|string|true|false|true|false|
|Removed|date|false|false|true|false|
|Type|enum|true|false|true|false|
|UUID|string|false|false|true|false|

### Object `BitbucketServerPipelineConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|ConsumerKey|string|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Enabled|boolean|true|true|false|false|
|Hostname|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PrivateKey|password|true|true|true|false|
|ProjectID|reference[project]|true|true|true|true|
|PublicKey|string|true|true|true|false|
|RedirectURL|string|true|true|true|false|
|Removed|date|false|false|true|false|
|TLS|boolean|true|true|false|false|
|Type|enum|true|false|true|false|
|UUID|string|false|false|true|false|

### Object `PipelineSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DisplayName|string|true|true|true|false|
|ProjectID|reference[project]|true|true|true|true|
|RepositoryURL|string|true|true|true|false|
|SourceCodeCredentialID|reference[sourceCodeCredential]|true|false|true|false|
|TriggerWebhookPr|boolean|true|true|false|false|
|TriggerWebhookPush|boolean|true|true|false|false|
|TriggerWebhookTag|boolean|true|true|false|false|

### Object `SourceCodeCredentialSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessToken|string|true|false|true|false|
|AvatarURL|string|true|true|true|false|
|DisplayName|string|true|true|true|true|
|Expiry|string|true|true|true|false|
|GitCloneToken|string|true|false|true|false|
|GitLoginName|string|true|true|true|false|
|HTMLURL|string|true|true|true|false|
|LoginName|string|true|true|true|false|
|ProjectID|reference[project]|true|true|true|false|
|RefreshToken|string|true|false|true|false|
|SourceCodeType|enum|true|true|true|true|
|UserID|reference[user]|true|true|true|true|

### Object `SourceCodeCredentialStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Logout|boolean|false|false|false|false|

### Object `SourceCodeCredential`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessToken|string|true|false|true|false|
|Annotations|map[string]|true|true|true|false|
|AvatarURL|string|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DisplayName|string|true|true|true|true|
|Expiry|string|true|true|true|false|
|GitCloneToken|string|true|false|true|false|
|GitLoginName|string|true|true|true|false|
|HTMLURL|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|LoginName|string|true|true|true|false|
|Logout|boolean|false|false|false|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|false|
|RefreshToken|string|true|false|true|false|
|Removed|date|false|false|true|false|
|SourceCodeType|enum|true|true|true|true|
|State|string|false|false|false|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UserID|reference[user]|true|true|true|true|
|UUID|string|false|false|true|false|

### Object `PipelineStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastExecutionID|string|false|false|true|false|
|LastRunState|string|false|false|true|false|
|LastStarted|string|false|false|true|false|
|NextRun|int|false|false|false|false|
|NextStart|string|false|false|true|false|
|PipelineState|enum|false|false|true|true|
|SourceCodeCredential|sourceCodeCredential|false|false|true|false|
|Token|string|false|false|true|false|
|WebHookID|string|false|false|true|false|

### Object `Pipeline`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Id|dnsLabel|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|LastExecutionID|string|false|false|true|false|
|LastRunState|string|false|false|true|false|
|LastStarted|string|false|false|true|false|
|Name|string|true|true|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NextRun|int|false|false|false|false|
|NextStart|string|false|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PipelineState|enum|false|false|true|true|
|ProjectID|reference[project]|true|true|true|true|
|Removed|date|false|false|true|false|
|RepositoryURL|string|true|true|true|false|
|SourceCodeCredential|sourceCodeCredential|false|false|true|false|
|SourceCodeCredentialID|reference[sourceCodeCredential]|true|false|true|false|
|State|string|false|false|false|false|
|Token|string|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|TriggerWebhookPr|boolean|true|true|false|false|
|TriggerWebhookPush|boolean|true|true|false|false|
|TriggerWebhookTag|boolean|true|true|false|false|
|UUID|string|false|false|true|false|
|WebHookID|string|false|false|true|false|

### Object `PipelineExecutionSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Author|string|true|true|true|false|
|AvatarURL|string|true|true|true|false|
|Branch|string|true|true|true|false|
|Commit|string|true|true|true|false|
|Email|string|true|true|true|false|
|Event|string|true|true|true|false|
|HTMLLink|string|true|true|true|false|
|Message|string|true|true|true|false|
|PipelineConfig|pipelineConfig|true|true|true|true|
|PipelineID|reference[pipeline]|true|true|true|true|
|ProjectID|reference[project]|true|true|true|true|
|Ref|string|true|true|true|false|
|RepositoryURL|string|true|true|true|false|
|Run|int|true|true|false|true|
|Title|string|true|true|true|false|
|TriggerUserID|reference[user]|true|true|true|false|
|TriggeredBy|enum|true|true|true|true|

### Object `PipelineCondition`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|LastTransitionTime|string|true|true|true|false|
|LastUpdateTime|string|true|true|true|false|
|Message|string|true|true|true|false|
|Reason|string|true|true|true|false|
|Status|string|true|true|true|false|
|Type|string|true|true|true|false|

### Object `StepStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Ended|string|true|true|true|false|
|Started|string|true|true|true|false|
|State|string|true|true|true|false|

### Object `StageStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Ended|string|true|true|true|false|
|Started|string|true|true|true|false|
|State|string|true|true|true|false|
|Steps|array[stepStatus]|true|true|true|false|

### Object `PipelineExecutionStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Conditions|array[pipelineCondition]|false|false|true|false|
|Ended|string|false|false|true|false|
|ExecutionState|string|false|false|true|false|
|Stages|array[stageStatus]|false|false|true|false|
|Started|string|false|false|true|false|

### Object `PipelineExecution`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Author|string|true|true|true|false|
|AvatarURL|string|true|true|true|false|
|Branch|string|true|true|true|false|
|Commit|string|true|true|true|false|
|Conditions|array[pipelineCondition]|false|false|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Email|string|true|true|true|false|
|Ended|string|false|false|true|false|
|Event|string|true|true|true|false|
|ExecutionState|string|false|false|true|false|
|HTMLLink|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Message|string|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PipelineConfig|pipelineConfig|true|true|true|true|
|PipelineID|reference[pipeline]|true|true|true|true|
|ProjectID|reference[project]|true|true|true|true|
|Ref|string|true|true|true|false|
|Removed|date|false|false|true|false|
|RepositoryURL|string|true|true|true|false|
|Run|int|true|true|false|true|
|Stages|array[stageStatus]|false|false|true|false|
|Started|string|false|false|true|false|
|State|string|false|false|false|false|
|Title|string|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|TriggerUserID|reference[user]|true|true|true|false|
|TriggeredBy|enum|true|true|true|true|
|UUID|string|false|false|true|false|

### Object `PipelineSetting`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Customized|boolean|false|false|false|false|
|Default|string|false|false|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[project]|true|true|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|
|Value|string|true|true|true|true|

### Object `RepoPerm`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Admin|boolean|true|true|false|false|
|Pull|boolean|true|true|false|false|
|Push|boolean|true|true|false|false|

### Object `SourceCodeRepositorySpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|DefaultBranch|string|true|true|true|false|
|Language|string|true|true|true|false|
|Permissions|repoPerm|true|true|true|false|
|ProjectID|reference[project]|true|true|true|false|
|SourceCodeCredentialID|reference[sourceCodeCredential]|true|true|true|true|
|SourceCodeType|enum|true|true|true|true|
|URL|string|true|true|true|false|
|UserID|reference[user]|true|true|true|true|

### Object `SourceCodeRepositoryStatus`
--- 


### Object `SourceCodeRepository`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|DefaultBranch|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Language|string|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Permissions|repoPerm|true|true|true|false|
|ProjectID|reference[project]|true|true|true|false|
|Removed|date|false|false|true|false|
|SourceCodeCredentialID|reference[sourceCodeCredential]|true|true|true|true|
|SourceCodeType|enum|true|true|true|true|
|State|string|false|false|false|false|
|Status|sourceCodeRepositoryStatus|false|false|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|URL|string|true|true|true|false|
|UserID|reference[user]|true|true|true|true|
|UUID|string|false|false|true|false|

### Object `StorageSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|EmptyDir|emptyDirVolumeSource|true|true|true|false|
|VolumeClaimTemplate|persistentVolumeClaim|true|true|true|false|

### Object `TLSConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|CAFile|string|true|true|true|false|
|CertFile|string|true|true|true|false|
|InsecureSkipVerify|boolean|true|true|false|false|
|KeyFile|string|true|true|true|false|
|ServerName|string|true|true|true|false|

### Object `AlertmanagerEndpoints`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BearerTokenFile|string|true|true|true|false|
|Name|string|true|true|true|false|
|Namespace|string|true|true|true|false|
|PathPrefix|string|true|true|true|false|
|Port|intOrString|true|true|true|false|
|Scheme|string|true|true|true|false|
|TLSConfig|tlsConfig|true|true|true|false|

### Object `AlertingSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Alertmanagers|array[alertmanagerEndpoints]|true|true|true|false|

### Object `RelabelConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Action|string|true|true|true|false|
|Modulus|int|true|true|false|false|
|Regex|string|true|true|true|false|
|Replacement|string|true|true|true|false|
|Separator|string|true|true|true|false|
|SourceLabels|array[string]|true|true|true|false|
|TargetLabel|string|true|true|true|false|

### Object `QueueConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BatchSendDeadline|string|true|true|true|false|
|Capacity|int|true|true|false|false|
|MaxBackoff|string|true|true|true|false|
|MaxRetries|int|true|true|false|false|
|MaxSamplesPerSend|int|true|true|false|false|
|MaxShards|int|true|true|false|false|
|MinBackoff|string|true|true|true|false|

### Object `RemoteWriteSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BasicAuth|basicAuth|true|true|true|false|
|BearerToken|string|true|true|true|false|
|BearerTokenFile|string|true|true|true|false|
|ProxyURL|string|true|true|true|false|
|QueueConfig|queueConfig|true|true|true|false|
|RemoteTimeout|string|true|true|true|false|
|TLSConfig|tlsConfig|true|true|true|false|
|URL|string|true|true|true|false|
|WriteRelabelConfigs|array[relabelConfig]|true|true|true|false|

### Object `RemoteReadSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BasicAuth|basicAuth|true|true|true|false|
|BearerToken|string|true|true|true|false|
|BearerTokenFile|string|true|true|true|false|
|ProxyURL|string|true|true|true|false|
|ReadRecent|boolean|true|true|false|false|
|RemoteTimeout|string|true|true|true|false|
|RequiredMatchers|map[string]|true|true|true|false|
|TLSConfig|tlsConfig|true|true|true|false|
|URL|string|true|true|true|false|

### Object `APIServerConfig`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BasicAuth|basicAuth|true|true|true|false|
|BearerToken|string|true|true|true|false|
|BearerTokenFile|string|true|true|true|false|
|Host|string|true|true|true|false|
|TLSConfig|tlsConfig|true|true|true|false|

### Object `ThanosGCSSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Bucket|string|true|true|true|false|
|SecretKey|secretKeySelector|true|true|true|false|

### Object `ThanosS3Spec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AccessKey|secretKeySelector|true|true|true|false|
|Bucket|string|true|true|true|false|
|EncryptSSE|boolean|true|true|true|false|
|Endpoint|string|true|true|true|false|
|Insecure|boolean|true|true|true|false|
|SecretKey|secretKeySelector|true|true|true|false|
|SignatureVersion2|boolean|true|true|true|false|

### Object `ThanosSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|BaseImage|string|true|true|true|false|
|GCS|thanosGCSSpec|true|true|true|false|
|Peers|string|true|true|true|false|
|Resources|resourceRequirements|true|true|true|false|
|S3|thanosS3Spec|true|true|true|false|
|SHA|string|true|true|true|false|
|Tag|string|true|true|true|false|
|Version|string|true|true|true|false|

### Object `PrometheusSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AdditionalAlertManagerConfigs|secretKeySelector|true|true|true|false|
|AdditionalAlertRelabelConfigs|secretKeySelector|true|true|true|false|
|AdditionalScrapeConfigs|secretKeySelector|true|true|true|false|
|Affinity|affinity|true|true|true|false|
|Alerting|alertingSpec|true|true|true|false|
|BaseImage|string|true|true|true|false|
|ConfigMaps|array[string]|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|EvaluationInterval|string|true|true|true|false|
|ExternalLabels|map[string]|true|true|true|false|
|ExternalURL|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|ListenLocal|boolean|true|true|false|false|
|LogLevel|enum|true|true|true|false|
|NodeSelector|map[string]|true|true|true|false|
|PodMetadata|objectMeta|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|RemoteRead|array[remoteReadSpec]|true|true|true|false|
|RemoteWrite|array[remoteWriteSpec]|true|true|true|false|
|Replicas|int|true|true|true|false|
|Resources|resourceRequirements|true|true|true|false|
|Retention|string|true|true|true|false|
|RoutePrefix|string|true|true|true|false|
|RuleSelector|labelSelector|true|true|true|false|
|ScrapeInterval|string|true|true|true|false|
|Secrets|array[string]|true|true|true|false|
|SecurityContext|podSecurityContext|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ServiceMonitorSelector|labelSelector|true|true|true|false|
|SHA|string|true|true|true|false|
|Storage|storageSpec|true|true|true|false|
|Tag|string|true|true|true|false|
|Tolerations|array[toleration]|true|true|true|false|
|Version|string|true|true|true|false|

### Object `PrometheusStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AvailableReplicas|int|false|false|false|false|
|Paused|boolean|false|false|false|false|
|Replicas|int|false|false|false|false|
|UnavailableReplicas|int|false|false|false|false|
|UpdatedReplicas|int|false|false|false|false|

### Object `Prometheus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AdditionalAlertManagerConfigs|secretKeySelector|true|true|true|false|
|AdditionalAlertRelabelConfigs|secretKeySelector|true|true|true|false|
|AdditionalScrapeConfigs|secretKeySelector|true|true|true|false|
|Affinity|affinity|true|true|true|false|
|Alerting|alertingSpec|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|BaseImage|string|true|true|true|false|
|ConfigMaps|array[string]|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Description|string|true|true|true|false|
|EvaluationInterval|string|true|true|true|false|
|ExternalLabels|map[string]|true|true|true|false|
|ExternalURL|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|ListenLocal|boolean|true|true|false|false|
|LogLevel|enum|true|true|true|false|
|Name|dnsLabelRestricted|true|false|false|true|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeSelector|map[string]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PodMetadata|objectMeta|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|RemoteRead|array[remoteReadSpec]|true|true|true|false|
|RemoteWrite|array[remoteWriteSpec]|true|true|true|false|
|Removed|date|false|false|true|false|
|Replicas|int|true|true|true|false|
|Resources|resourceRequirements|true|true|true|false|
|Retention|string|true|true|true|false|
|RoutePrefix|string|true|true|true|false|
|RuleSelector|labelSelector|true|true|true|false|
|ScrapeInterval|string|true|true|true|false|
|Secrets|array[string]|true|true|true|false|
|SecurityContext|podSecurityContext|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|ServiceMonitorSelector|labelSelector|true|true|true|false|
|SHA|string|true|true|true|false|
|State|string|false|false|false|false|
|Storage|storageSpec|true|true|true|false|
|Tag|string|true|true|true|false|
|Tolerations|array[toleration]|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|
|Version|string|true|true|true|false|

### Object `Endpoint`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Interval|string|true|true|true|false|
|Params|map[array[string]]|true|true|true|false|
|Path|string|true|true|true|false|
|RelabelConfigs|array[relabelConfig]|true|true|true|false|
|Scheme|string|true|true|true|false|
|ScrapeTimeout|string|true|true|true|false|
|TargetPort|intOrString|true|true|true|false|

### Object `NamespaceSelector`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Any|boolean|true|true|false|false|
|MatchNames|array[string]|true|true|true|false|

### Object `ServiceMonitorSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Endpoints|array[endpoint]|true|true|true|false|
|JobLabel|string|true|true|true|false|
|NamespaceSelector|array[string]|true|true|true|false|
|PodTargetLabels|array[string]|true|true|true|false|
|SampleLimit|int|true|true|false|false|
|Selector|labelSelector|true|true|true|false|
|TargetLabels|array[string]|true|true|true|false|

### Object `ServiceMonitor`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Endpoints|array[endpoint]|true|true|true|false|
|Id|dnsLabel|false|false|true|false|
|JobLabel|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|string|true|true|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NamespaceSelector|array[string]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|PodTargetLabels|array[string]|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|SampleLimit|int|true|true|false|false|
|Selector|labelSelector|true|true|true|false|
|TargetLabels|array[string]|true|true|true|false|
|TargetService|string|true|true|true|false|
|TargetWorkload|string|true|true|true|false|
|UUID|string|false|false|true|false|

### Object `Rule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Alert|string|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|Expr|intOrString|true|true|true|false|
|For|string|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Record|string|true|true|true|false|

### Object `RuleGroup`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Interval|string|true|true|true|false|
|Name|string|true|true|true|false|
|Rules|array[rule]|true|true|true|false|

### Object `PrometheusRuleSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Groups|array[ruleGroup]|true|true|true|false|

### Object `PrometheusRule`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|Annotations|map[string]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|Groups|array[ruleGroup]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|UUID|string|false|false|true|false|

### Object `AlertmanagerSpec`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AdditionalPeers|array[string]|true|true|true|false|
|Affinity|affinity|true|true|true|false|
|BaseImage|string|true|true|true|false|
|ConfigMaps|array[string]|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|ExternalURL|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|ListenLocal|boolean|true|true|false|false|
|LogLevel|string|true|true|true|false|
|NodeSelector|map[string]|true|true|true|false|
|Paused|boolean|true|true|false|false|
|PodMetadata|objectMeta|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|Replicas|int|true|true|true|false|
|Resources|resourceRequirements|true|true|true|false|
|Retention|string|true|true|true|false|
|RoutePrefix|string|true|true|true|false|
|Secrets|array[string]|true|true|true|false|
|SecurityContext|podSecurityContext|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|SHA|string|true|true|true|false|
|Storage|storageSpec|true|true|true|false|
|Tag|string|true|true|true|false|
|Tolerations|array[toleration]|true|true|true|false|
|Version|string|true|true|true|false|

### Object `AlertmanagerStatus`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AvailableReplicas|int|false|false|false|false|
|Paused|boolean|false|false|false|false|
|Replicas|int|false|false|false|false|
|UnavailableReplicas|int|false|false|false|false|
|UpdatedReplicas|int|false|false|false|false|

### Object `Alertmanager`
--- 
Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|
|AdditionalPeers|array[string]|true|true|true|false|
|Affinity|affinity|true|true|true|false|
|Annotations|map[string]|true|true|true|false|
|BaseImage|string|true|true|true|false|
|ConfigMaps|array[string]|true|true|true|false|
|Containers|array[container]|true|true|true|false|
|Created|date|false|false|true|false|
|CreatorID|reference[/v3/schemas/user]|false|false|false|false|
|ExternalURL|string|true|true|true|false|
|ImagePullSecrets|array[localObjectReference]|true|true|true|false|
|Labels|map[string]|true|true|true|false|
|ListenLocal|boolean|true|true|false|false|
|LogLevel|string|true|true|true|false|
|Name|dnsLabel|true|false|true|false|
|NamespaceId|reference[/v3/clusters/schemas/namespace]|true|false|true|true|
|NodeSelector|map[string]|true|true|true|false|
|OwnerReferences|array[ownerReference]|false|false|true|false|
|Paused|boolean|true|true|false|false|
|PodMetadata|objectMeta|true|true|true|false|
|PriorityClassName|string|true|true|true|false|
|ProjectID|reference[/v3/schemas/project]|true|false|true|false|
|Removed|date|false|false|true|false|
|Replicas|int|true|true|true|false|
|Resources|resourceRequirements|true|true|true|false|
|Retention|string|true|true|true|false|
|RoutePrefix|string|true|true|true|false|
|Secrets|array[string]|true|true|true|false|
|SecurityContext|podSecurityContext|true|true|true|false|
|ServiceAccountName|string|true|true|true|false|
|SHA|string|true|true|true|false|
|State|string|false|false|false|false|
|Storage|storageSpec|true|true|true|false|
|Tag|string|true|true|true|false|
|Tolerations|array[toleration]|true|true|true|false|
|Transitioning|enum|false|false|false|false|
|TransitioningMessage|string|false|false|false|false|
|UUID|string|false|false|true|false|
|Version|string|true|true|true|false|

---
