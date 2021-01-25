// Copyright 2021 Authors of Cilium
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

package persistence

import (
	"os"

	"github.com/cilium/team-manager/pkg/config"

	"gopkg.in/yaml.v2"
)

func StoreState(file string, newConfig *config.Config) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewEncoder(f).Encode(newConfig)
}

func LoadState(file string) (*config.Config, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0440)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	storedConfig := config.Config{}
	err = yaml.NewDecoder(f).Decode(&storedConfig)
	if err != nil {
		return nil, err
	}
	return &storedConfig, nil
}