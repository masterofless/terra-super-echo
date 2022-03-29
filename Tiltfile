#
#   More info: https://docs.tilt.dev/api.html#api.warn
# print("""
# -----------------------------------------------------------------
# ✨ Hello Tilt! This appears in the (Tiltfile) pane whenever Tilt
   # evaluates this file.
# -----------------------------------------------------------------
# """.strip())
# warn('ℹ️ Open {tiltfile_path} in your favorite editor to get started.'.format(
    # tiltfile_path=config.main_path))


docker_build('go-super-echo', context='images/go-super-echo')
docker_build('nginx-super-echo', context='images/nginx-super-echo')
docker_build('node-super-echo', context='images/node-super-echo')

k8s_yaml(helm('./', name='terra-super-echo'))

k8s_resource('terra-super-echo-nginx', port_forwards='8000:80')
k8s_resource('terra-super-echo-node', port_forwards='8001:8080')
k8s_resource('terra-super-echo-go', port_forwards='8002:8080')

# Run local commands
#   Local commands can be helpful for one-time tasks like installing
#   project prerequisites. They can also manage long-lived processes
#   for non-containerized services or dependencies.
#
#   More info: https://docs.tilt.dev/local_resource.html
#
# local_resource('install-helm',
#                cmd='which helm > /dev/null || brew install helm',
#                # `cmd_bat`, when present, is used instead of `cmd` on Windows.
#                cmd_bat=[
#                    'powershell.exe',
#                    '-Noninteractive',
#                    '-Command',
#                    '& {if (!(Get-Command helm -ErrorAction SilentlyContinue)) {scoop install helm}}'
#                ]
# )


# Extensions are open-source, pre-packaged functions that extend Tilt
#
#   More info: https://github.com/tilt-dev/tilt-extensions
#
load('ext://git_resource', 'git_checkout')

