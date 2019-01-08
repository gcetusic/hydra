// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	// "context"
	// "log"

	// "github.com/containerd/containerd"
	// "github.com/containerd/containerd/namespaces"
	// "github.com/containerd/containerd/oci"
	"github.com/spf13/cobra"
)

type ClientConfig struct {
	Host string `yaml:"host"`
}

var sourcePath, runtime string
var force, noCleanup bool

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// log.Printf("Build called for runtime %s", runtime)

		// client, err := containerd.New("/run/containerd/containerd.sock")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer client.Close()

		// context := namespaces.WithNamespace(context.Background(), "runtime")
		// image, err := client.Pull(context, "docker.io/library/alpine:latest", containerd.WithPullUnpack)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Printf("Successfully pulled %s image\n", image.Name())

		// container, err := client.NewContainer(context, "golang-2",
		// 	containerd.WithNewSnapshot("golang-rootfs1", image),
		// 	containerd.WithNewSpec(oci.WithImageConfig(image)),
		// )

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// defer container.Delete(context, containerd.WithSnapshotCleanup)
		// log.Printf("Successfully created container with ID %s and snapshot with ID redis-server-snapshot", container.ID())
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().StringVar(&sourcePath, "path", "", "<string,required> path to root application sources folder")
	buildCmd.Flags().StringVar(&runtime, "runtime", "", "<string,required> name of the runtime to use")
	buildCmd.Flags().BoolVar(&force, "force", false, "<bool, optional> force overwriting a previously existing")
	buildCmd.Flags().BoolVar(&noCleanup, "no-cleanup", false, "<bool, optional> for debugging; do not clean up artifacts for images that fail to build")
}
