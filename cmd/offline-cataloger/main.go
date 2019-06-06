// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

package main

import (
	"os"

	"github.com/kevinrizza/offline-cataloger/cmd/offline-cataloger/build"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	root := &cobra.Command{
		Use:   "offline-cataloger",
		Short: "A tool to generate operator catalogs for a Kubernetes cluster",
		Long: `
offline-cataloger is a CLI tool to generate self contained operator
catalogs. This application defines a set of tooling to pull operators
from hosted app-registries and make them available locally on a
Kubernetes cluster without the need to reach out to external sources
from the cluster.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if viper.GetBool("verbose") {
				log.SetLevel(log.DebugLevel)
				log.Debug("Debug logging is set")
			}
		},
	}

	root.AddCommand(build.NewCmd())

	root.PersistentFlags().Bool("verbose", false, "Enable verbose logging")
	if err := viper.BindPFlags(root.PersistentFlags()); err != nil {
		log.Fatalf("Failed to bind root flags: %v", err)
	}

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
