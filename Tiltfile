# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_remote', 'helm_remote')
compile_opt = 'GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 '

# Compile example application
local_resource(
  'hello-world-compile',
  compile_opt + 'go build -o bin/hello-world main.go',
  deps=['.'],
  ignore=['bin', 'helm', 'Dockerfile', 'Tiltfile', 'README.md', 'LICENSE', '.gitignore'],
  labels="example-application",
)

# Build example docker image
docker_build_with_restart(
  'hello-world-image',
  '.',
  entrypoint=['/opt/app/bin/hello-world'],
  dockerfile='Dockerfile',
  only=[
    './bin',
  ],
  live_update=[
    sync('./bin', '/opt/app/bin'),
  ],
)

# Install example helm chart
k8s_yaml(helm('helm'))

# Label and port forwarding example applciation
k8s_resource(
  "hello-world",
  port_forwards='8080:8080',
  labels="example-application",
)
