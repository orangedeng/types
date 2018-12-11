# Rancher API Document
---

## API Specification

> API-Spec Document is From `https://github.com/rancher/api-spec/blob/master/specification.md`

## API List: ##
{{range $index, $schema := .Schemas}}{{if not ($schema.CanList nil)}}
## Object Schema Name: `{{ $schema.CodeName}}`
### K8s Object Meta
|||
|---|---|
|API Group| `{{$schema.Version.Group}}` |
|API Vesion| `{{$schema.Version.Version}}` |
|Storage Scope| {{if eq $schema.Version.Group "management.cattle.io"}}Management{{else}}User Cluster{{end}} |
|Namespaced| `{{if eq $schema.Scope "namespace"}}true{{else}}false{{end}}` |
|Package Name | {{$schema.PkgName}} |
### Get `{{$schema.CodeName}}` List
|||
|---|---|
|Method| `Get`|
|Path|`{{ $.GetPath $schema.Version.Path}}{{$schema.PluralName}}`|
|Input| `null` |
|Output| List of [resource object `{{$schema.CodeName}}`](#object-{{$schema.CodeName}}) |{{if not ($schema.CanGet nil)}}
### Get `{{$schema.CodeName}}`
|||
|---|---|
|Method| `Get` |
|Path| `{{ $.GetPath $schema.Version.Path}}{{$schema.PluralName}}/<resource_id>` |
|Input| `null`|
|Output| [Resource object `{{$schema.CodeName}}`](#object-{{$schema.CodeName}}) |{{end}}{{if not ($schema.CanCreate nil)}}
### Create `{{$schema.CodeName}}`
|||
|---|---|
|Method| `Post` |
|Path| `{{ $.GetPath $schema.Version.Path}}{{$schema.PluralName}}` |
|Input| [Resource object `{{$schema.CodeName}}`](#object-{{$schema.CodeName}}) |
|Output| [Resource object `{{$schema.CodeName}}`](#object-{{$schema.CodeName}}) after create |{{end}}{{if not ($schema.CanUpdate nil)}}
### Update `{{$schema.CodeName}}`
|||
|---|---|
|Method| `Put` |
|Path| `{{ $.GetPath $schema.Version.Path}}{{$schema.PluralName}}/<resource_id>` |
|Input| [Resource object `{{$schema.CodeName}}`](#object-{{$schema.CodeName}}) to update |
|Output| [Resource object `{{$schema.CodeName}}`](#object-{{$schema.CodeName}}) after delete |{{end}}{{if not ($schema.CanDelete nil)}}
### Delete `{{$schema.CodeName}}`
|||
|---|---|
|Method| `Delete` |
|Path| `{{ $.GetPath $schema.Version.Path}}{{$schema.PluralName}}/<resource_id>` |
|Input| `null` |
|Output| [Resource object `{{$schema.CodeName}}`](#object-{{$schema.CodeName}}) after delete |{{end}}{{range $actionName, $action :=$schema.ResourceActions}}
### Resource Action `{{$actionName}}`
|||
|---|---|
|Method| `Post` |
|Path| `{{ $.GetPath $schema.Version.Path}}{{$schema.PluralName}}/<resource_id>/{{$actionName}}` |
|Input| {{if ne (len $action.Input) 0}}[object `{{$action.Input}}`](#object-{{$action.Input}}){{else}}`null`{{end}} |
|Output| {{if ne (len $action.Output) 0}}[object `{{$action.Output}}`](#object-{{$action.Output}}){{else}}`null`{{end}} |{{end}}{{range $actionName, $action :=$schema.CollectionActions}}
### Resource List Action `{{$actionName}}`
|||
|---|---|
|Method| `Post` |
|Path| `{{ $.GetPath $schema.Version.Path}}{{$schema.PluralName}}/<resource_id>` |
|Input| {{if ne (len $action.Input) 0}}[object `{{$action.Input}}`](#object-{{$action.Input}}){{else}}`null`{{end}} |
|Output| {{if ne (len $action.Output) 0}}[object `{{$action.Output}}`](#object-{{$action.Output}}){{else}}`null`{{end}} |{{end}}
---
{{end}}{{end}}

## Resource Object Reference: ##
{{range $index, $schema := .Schemas}}
### Object `{{ $schema.CodeName}}`
--- 
{{with $schema.ResourceFields}}Object Attribute:

|Name|Type|Can Create|Can Update|Can Null|Required for Input|
|---|---|---|---|---|---|{{range $iindex, $field := $schema.ResourceFields}}
|{{ $field.CodeName}}|{{ $field.Type}}|{{$field.Create}}|{{$field.Update}}|{{$field.Nullable}}|{{$field.Required}}|{{end}}{{end}}
{{end}}
---
