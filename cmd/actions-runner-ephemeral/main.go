package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

var NamePrefix = os.Getenv("NAME_PREFIX")
var RequestImage = os.Getenv("RUN_IMAGE")
var EnvList = os.Getenv("ENV_LIST")
var ContainerCount = os.Getenv("CONTAINER_COUNT")
var NetworkName = os.Getenv("NETWORK_NAME")

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containerCount, err := strconv.ParseInt(ContainerCount, 10, 64)
	if err != nil {
		panic(err)
	}

	var envs []string
	err = json.Unmarshal([]byte(EnvList), &envs)
	if err != nil {
		panic(err)
	}

	networkID := ""

	if len(NetworkName) > 0 {
		networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{
			Filters: filters.NewArgs(filters.Arg("name", NetworkName)),
		})
		if err != nil {
			panic(err)
		}

		for _, net := range networks {
			networkID = net.ID
		}
	}

	for {
		ctx := context.Background()

		containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}

		containersFiltered := lo.Filter(containers, func(c types.Container, _ int) bool {
			return len(lo.Filter(c.Names, func(n string, _ int) bool {
				return strings.HasPrefix(strings.TrimPrefix(n, "/"), NamePrefix)
			})) > 0
		})

		createCount := int(containerCount) - len(containersFiltered)

		if createCount > 0 {
			for i := 0; i < createCount; i++ {
				name := strings.Split(uuid.New().String(), "-")[0]
				containerName := fmt.Sprintf("%s-%s", NamePrefix, name)
				hostConfig := &container.HostConfig{
					AutoRemove: true,
					Mounts: []mount.Mount{
						{
							Type:   mount.TypeBind,
							Source: "/var/run/docker.sock",
							Target: "/var/run/docker.sock",
						},
					},
				}

				if len(networkID) > 0 {
					hostConfig.NetworkMode = container.NetworkMode(networkID)
				}

				res, err := cli.ContainerCreate(ctx, &container.Config{
					Image: RequestImage,
					Env:   envs,
				}, hostConfig, &network.NetworkingConfig{}, nil, containerName)

				if err != nil {
					panic(err)
				}

				err = cli.ContainerStart(ctx, res.ID, types.ContainerStartOptions{})

				if err != nil {
					panic(err)
				}
				fmt.Printf("Container %s started\n", containerName)
			}
		}

		time.Sleep(time.Second)
	}
}
