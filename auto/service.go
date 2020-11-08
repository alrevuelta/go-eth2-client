// Copyright © 2020 Attestant Limited.
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

package auto

import (
	"context"

	client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/prysmgrpc"
	standardhttp "github.com/attestantio/go-eth2-client/standardhttp/v1"
	"github.com/attestantio/go-eth2-client/tekuhttp"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	zerologger "github.com/rs/zerolog/log"
)

// log is a service-wide logger.
var log zerolog.Logger

// New creates a new Ethereum 2 client service, trying different implementations at the given address.
func New(ctx context.Context, params ...Parameter) (client.Service, error) {
	parameters, err := parseAndCheckParameters(params...)
	if err != nil {
		return nil, errors.Wrap(err, "problem with parameters")
	}

	// Set logging.
	log = zerologger.With().Str("service", "client").Str("impl", "auto").Logger()
	if parameters.logLevel != log.GetLevel() {
		log = log.Level(parameters.logLevel)
	}

	// Try standard.
	standardClient, err := tryStandard(ctx, parameters)
	if err == nil {
		return standardClient, nil
	}

	// Try prysm.
	prysmClient, err := tryPrysm(ctx, parameters)
	if err == nil {
		return prysmClient, nil
	}

	// Try teku.
	tekuClient, err := tryTeku(ctx, parameters)
	if err == nil {
		return tekuClient, nil
	}

	// No luck
	return nil, errors.New("failed to connect to Ethereum 2 client with any known method")
}

func tryStandard(ctx context.Context, parameters *parameters) (*standardhttp.Service, error) {
	standardhttpParameters := make([]standardhttp.Parameter, 0)
	standardhttpParameters = append(standardhttpParameters, standardhttp.WithLogLevel(parameters.logLevel))
	standardhttpParameters = append(standardhttpParameters, standardhttp.WithAddress(parameters.address))
	standardhttpParameters = append(standardhttpParameters, standardhttp.WithTimeout(parameters.timeout))
	client, err := standardhttp.New(ctx, standardhttpParameters...)
	if err != nil {
		return nil, errors.Wrap(err, "failed when trying to open connection with standard API")
	}
	return client, nil
}

func tryPrysm(ctx context.Context, parameters *parameters) (*prysmgrpc.Service, error) {
	prysmParameters := make([]prysmgrpc.Parameter, 0)
	prysmParameters = append(prysmParameters, prysmgrpc.WithLogLevel(parameters.logLevel))
	prysmParameters = append(prysmParameters, prysmgrpc.WithAddress(parameters.address))
	prysmParameters = append(prysmParameters, prysmgrpc.WithTimeout(parameters.timeout))
	client, err := prysmgrpc.New(ctx, prysmParameters...)
	if err != nil {
		return nil, errors.Wrap(err, "failed when trying to open connection to prysm")
	}
	return client, nil
}

func tryTeku(ctx context.Context, parameters *parameters) (*tekuhttp.Service, error) {
	tekuParameters := make([]tekuhttp.Parameter, 0)
	tekuParameters = append(tekuParameters, tekuhttp.WithLogLevel(parameters.logLevel))
	tekuParameters = append(tekuParameters, tekuhttp.WithAddress(parameters.address))
	tekuParameters = append(tekuParameters, tekuhttp.WithTimeout(parameters.timeout))
	client, err := tekuhttp.New(ctx, tekuParameters...)
	if err != nil {
		return nil, errors.Wrap(err, "failed when trying to open connection to teku")
	}
	return client, nil
}
