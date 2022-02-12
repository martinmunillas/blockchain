package main

import (
	"fmt"
	"os"

	"github.com/martinmunillas/munichain/node"
	"github.com/spf13/cobra"
)

func runCmd() *cobra.Command {

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Runs the node.",
		Run: func(cmd *cobra.Command, args []string) {
			dataDir, _ := cmd.Flags().GetString(flagDataDir)
			port, _ := cmd.Flags().GetUint64(flagPort)
			fmt.Println("Launching munichain node and its HTTP API...")

			bootstrap := node.PeerNode{
				IP:          "0.0.0.0",
				Port:        8080,
				IsBootstrap: true,
				IsActive:    true,
			}

			n := node.Node{
				DataDir: dataDir,
				Port:    port,
				KnownPeers: []node.PeerNode{
					bootstrap,
				},
			}
			err := n.Run()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	addDefaultRequiredFlags(runCmd)
	runCmd.Flags().Uint64(flagPort, node.DefaultHttpPort, "exposed HTTP port for communication with peers")

	return runCmd
}
