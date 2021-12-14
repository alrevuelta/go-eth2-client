// Copyright © 2021 Attestant Limited.
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

package multi

import (
	"context"

	consensusclient "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

// FarFutureEpoch provides the far future epoch of the chain.
func (s *Service) FarFutureEpoch(ctx context.Context) (phase0.Epoch, error) {
	res, err := s.doCall(ctx, func(ctx context.Context, client consensusclient.Service) (interface{}, error) {
		aggregate, err := client.(consensusclient.FarFutureEpochProvider).FarFutureEpoch(ctx)
		if err != nil {
			return nil, err
		}
		return aggregate, nil
	}, nil)
	if err != nil {
		return 0, err
	}
	if res == nil {
		return 0, nil
	}
	return res.(phase0.Epoch), nil
}
