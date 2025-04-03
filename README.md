## `yamlpodspec`

Simple demonstration using [CUE](https://cuelang.org) with the [CUE Central
Registry](https://registry.cue.works), to validate that a YAML file is a valid
[Kubernetes
PodSpec](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#PodSpec),
using the
[cue.dev/x/k8s.io/api/core/v1](https://registry.cue.works/docs/cue.dev/x/k8s.io/api/core/v1)
package published to the CUE Central Registry.

Requires [CUE v0.13.0-alpha.3](https://cue.dev/docs/installing-cue/):

```
$ cue version
cue version v0.13.0-alpha.3
...
```

Run the demo locally:

```
cue login # only during the beta of the Central Registry

# cue vet directly against Central Registry schema
cue vet -c -d '#PodSpec' cue.dev/x/k8s.io/api/core/v1@latest data-good.yaml
cue vet -c -d '#PodSpec' cue.dev/x/k8s.io/api/core/v1@latest data-bad.yaml

# Using Go API
cue mod tidy
go mod tidy
go run . data-good.yaml
go run . data-bad.yaml
```

The `data-good.yaml` file passes validation for both `cmd/cue` and the Go program.
`data-bad.yaml` repo intentionally fails validation, for both `cmd/cue` and the
Go program.
