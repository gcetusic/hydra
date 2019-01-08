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
	// "fmt"
	// "os"

	// "github.com/firecracker-microvm/firecracker-go-sdk"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var fireCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// socketPath := "/tmp/firecracker.socket"
		// logFifo := "/tmp/firecracker/log.fifo"
		// metricsFifo := "/tmp/firecracker/metrics.fifo"
		// binPath := "/home/goc/build/firecracker"

		// cfg := firecracker.Config{
		// 	SocketPath:      socketPath,
		// 	KernelImagePath: "/home/goc/build/hello-vmlinux.bin",
		// 	RootDrive: firecracker.BlockDevice{
		// 		HostPath: "/home/goc/build/hello-rootfs.ext4",
		// 		Mode:     "rw",
		// 	},
		// 	LogFifo:     logFifo,
		// 	MetricsFifo: metricsFifo,
		// 	LogLevel:    "Debug",
		// 	CPUCount:    1,
		// 	CPUTemplate: firecracker.CPUTemplate(firecracker.CPUTemplateT2),
		// 	MemInMiB:    128,
		// 	Debug:       true,
		// }

		// // stdout will be directed to this file
		// stdoutPath := "/tmp/stdout.log"
		// stdout, err := os.OpenFile(stdoutPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		// if err != nil {
		// 	panic(fmt.Errorf("failed to create stdout file: %v", err))
		// }

		// // stderr will be directed to this file
		// stderrPath := "/tmp/stderr.log"
		// stderr, err := os.OpenFile(stderrPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		// if err != nil {
		// 	panic(fmt.Errorf("failed to create stderr file: %v", err))
		// }

		// ctx := context.Background()
		// // build our custom command that contains our two files to
		// // write to during process execution
		// vmCmd := firecracker.VMCommandBuilder{}.
		// 	WithBin(binPath).
		// 	WithSocketPath(socketPath).
		// 	WithStdout(stdout).
		// 	WithStderr(stderr).
		// 	Build(ctx)

		// machine, err := firecracker.NewMachine(cfg, firecracker.WithProcessRunner(vmCmd))
		// if err != nil {
		// 	panic(fmt.Errorf("failed to create new machine: %v", err))
		// }

		// ch, err := machine.Init(ctx)
		// if err != nil {
		// 	panic(fmt.Errorf("failed to initialize machine: %v", err))
		// }

		// if err := machine.StartInstance(ctx); err != nil {
		// 	panic(fmt.Errorf("Failed to start instance: %v", err))
		// }

		// go func() {
		// 	<-ch
		// 	os.Remove(cfg.SocketPath)
		// }()

		// fmt.Println("run called")

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	fireCmd.Flags().StringVar(&runtime, "runtime", "", "<string,required> name of the runtime to use")
}
