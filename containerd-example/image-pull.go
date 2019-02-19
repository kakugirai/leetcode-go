package containerd_example

import (
	"context"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

func main() {
	if err := redisExample(); err != nil {
		log.Fatal(err)
	}
}

func redisExample() error {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()

	ctx := namespaces.WithNamespace(context.Background(), "example")
	image, err := client.Pull(ctx, "docker.io/library/redis:alpine", containerd.WithPullUnpack)
	if err != nil {
		return err
	}
	log.Printf("Successfully pulled %s image\n", image.Name())

	return nil
}
