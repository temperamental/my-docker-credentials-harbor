# docker-credentials-harbor

Simple docker credentials plugin

## Build

```bash
go build -o docker-credentials-harbor
```

## Usage

Define the following environment variables

```bash
export HARBOR_CLIENT_NAME="username"
export HARBOR_CLIENT_SECRET="ha-ha-as-if"
```

Put the plugin `docker-credentials-harbor` in a directory that is in your $PATH.

Edit your Docker daemon configuration file and add the following line:

```json
plugins: /path/to/plugin
```

You may need to restart the Docker daemon to pick up this change.
