/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"fmt"
)

// Info contains versioning information.
// TODO: Add []string of api versions supported? It's still unclear
// how we'll want to distribute that information.
type Info struct {
	Major        string `json:"major" yaml:"major"`
	Minor        string `json:"minor" yaml:"minor"`
	GitVersion   string `json:"gitVersion" yaml:"gitVersion"`
	GitCommit    string `json:"gitCommit" yaml:"gitCommit"`
	GitTreeState string `json:"gitTreeState" yaml:"gitTreeState"`
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() Info {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the settings in pkg/version/base.go
	return Info{
		Major:        gitMajor,
		Minor:        gitMinor,
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
	}
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	commit := info.GitCommit
	if commit == "" {
		commit = "(unknown)"
	}
	return fmt.Sprintf("version %s.%s, build %s", info.Major, info.Minor, commit)
}
