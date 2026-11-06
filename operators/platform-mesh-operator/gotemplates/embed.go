/*
Copyright The Platform Mesh Authors.

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

// Package gotemplates exposes the Go templates the platform-mesh-operator
// renders into cluster resources as an embedded file system: infra/ (the
// infrastructure components and their runtime resources) and components/ (the
// Platform Mesh components).
package gotemplates

import (
	"embed"
	"io/fs"
)

//go:embed infra components
var FS embed.FS

// Infra is the gotemplates/infra tree, rooted at that directory.
func Infra() fs.FS { return sub("infra") }

// Components is the gotemplates/components tree, rooted at that directory.
func Components() fs.FS { return sub("components") }

func sub(dir string) fs.FS {
	f, err := fs.Sub(FS, dir)
	if err != nil {
		panic(err) // the directories are embedded above
	}
	return f
}
