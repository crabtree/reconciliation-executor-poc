module github.com/crabtree/reconciliation-executor-helm

require helm.sh/helm/v3 v3.5.4

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)

go 1.16
