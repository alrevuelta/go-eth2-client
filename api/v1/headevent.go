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

package v1

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

// HeadEvent is the data for the head event.
type HeadEvent struct {
	Slot            spec.Slot
	Block           spec.Root
	State           spec.Root
	EpochTransition bool
}

// headEventJSON is the spec representation of the struct.
type headEventJSON struct {
	Slot            string `json:"slot"`
	Block           string `json:"block"`
	State           string `json:"state"`
	EpochTransition bool   `json:"epoch_transition"`
}

// MarshalJSON implements json.Marshaler.
func (e *HeadEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&headEventJSON{
		Slot:            fmt.Sprintf("%d", e.Slot),
		Block:           fmt.Sprintf("%#x", e.Block),
		State:           fmt.Sprintf("%#x", e.State),
		EpochTransition: e.EpochTransition,
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *HeadEvent) UnmarshalJSON(input []byte) error {
	var err error

	var headEventJSON headEventJSON
	if err = json.Unmarshal(input, &headEventJSON); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}
	if headEventJSON.Slot == "" {
		return errors.New("slot missing")
	}
	slot, err := strconv.ParseUint(headEventJSON.Slot, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for slot")
	}
	e.Slot = spec.Slot(slot)
	if headEventJSON.Block == "" {
		return errors.New("block missing")
	}
	block, err := hex.DecodeString(strings.TrimPrefix(headEventJSON.Block, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid value for block")
	}
	if len(block) != rootLength {
		return fmt.Errorf("incorrect length %d for block", len(block))
	}
	copy(e.Block[:], block)
	if headEventJSON.State == "" {
		return errors.New("state missing")
	}
	state, err := hex.DecodeString(strings.TrimPrefix(headEventJSON.State, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid value for state")
	}
	if len(state) != rootLength {
		return fmt.Errorf("incorrect length %d for state", len(state))
	}
	copy(e.State[:], state)
	e.EpochTransition = headEventJSON.EpochTransition

	return nil
}

// String returns a string version of the structure.
func (e *HeadEvent) String() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(data)
}
