// Copyright © 2019 The Offline-Cataloger Authors
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

package generate

import (
	"fmt"

	"github.com/kevinrizza/offline-cataloger/pkg/apis"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	endpointArg     = "endpoint"
	namespaceArg    = "namespace"
	authTokenArg    = "auth-token"
	defaultEndpoint = "https://quay.io/cnr"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-manifests",
		Short: "Generates manifests from appregistry namespace",
		RunE:  generateFunc,
	}

	cmd.Flags().StringP(authTokenArg, "a", "", "Authentication Token for App Registry endpoint")
	cmd.Flags().StringP(endpointArg, "e", "", "App Registry endpoint. Defaults to https://quay.io/cnr")

	return cmd
}

func generateFunc(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("command %s requires exactly one argument", cmd.CommandPath())
	}

	// Required Args
	namespace := args[0]

	// Optional Args
	authToken, _ := cmd.Flags().GetString(authTokenArg)

	endpoint, _ := cmd.Flags().GetString(endpointArg)
	if endpoint == "" {
		endpoint = defaultEndpoint
	}

	// Create the request to be handled by the builder
	request := &apis.GenerateManifestsRequest{
		Namespace:          namespace,
		AuthorizationToken: authToken,
		Endpoint:           endpoint,
	}

	log.Infof("Generating manifests from %s in namespace %s", request.Endpoint, request.Namespace)

	generateHandler, err := apis.NewGenerateHandler()
	if err != nil {
		return err
	}

	// Build
	err = generateHandler.Handle(request)
	if err != nil {
		return err
	}

	return nil
}
