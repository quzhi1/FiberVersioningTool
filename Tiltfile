# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_resource', 'helm_resource')
compile_opt = 'GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 '

# Swag
local_resource(
  'update-swagger',
  'swag init -g main.go --output docs',
  deps=['main.go'],
  labels="hello-world",
)

# Compile example application
local_resource(
  'hello-world-compile',
  compile_opt + 'go build -o bin/hello-world main.go',
  deps=['main.go', 'docs'],
  ignore=['bin', 'helm', 'Dockerfile', 'Tiltfile', 'README.md', 'LICENSE', '.gitignore'],
  resource_deps=['update-swagger'],
  labels="hello-world",
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
helm_resource(
  'hello-world',
  'helm',
  image_deps=['hello-world-image'],
  image_keys=[('image.repository', 'image.tag')],
  port_forwards=['8080:8080'],
  labels="hello-world",
)
