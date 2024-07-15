# actions-runner-ephemeral

Ephemeral Runner Container Schduler

## Usage

```yaml
version: '3'

services:
  schduler:
    image: ghcr.io/5aaee9/actions-runner-ephemeral:main
    restart: always

    networks: [runner]
    environment:
      NAME_PREFIX: github-runner
      RUN_IMAGE: knatnetwork/github-runner:latest
      CONTAINER_COUNT: '4'
      NETWORK_NAME: github_runner_network
      ENV_LIST: |
        [
          "RUNNER_REGISTER_TO=Indexyz",
          "RUNNER_LABELS=kvm",
          "KMS_SERVER_ADDR=http://kms:3000",
          "ADDITIONAL_FLAGS=--ephemeral"
        ]
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock

  kms:
    networks: [runner]
    image: knatnetwork/github-runner-kms:latest
    restart: always
    environment:
      PAT_Indexyz: 'PAT_YOUR_PAT_HERE'

networks:
  runner: 
    name: github_runner_network
```