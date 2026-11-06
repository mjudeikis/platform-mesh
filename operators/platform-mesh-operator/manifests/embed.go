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

// Package manifests exposes the manifests the platform-mesh-operator applies
// as an embedded file system: kcp/ (workspaces, WorkspaceTypes, APIExports,
// bindings, ...), features/ (the kcp manifests of each feature toggle) and
// k8s/ (cluster manifests). The binary and its manifests cannot disagree,
// and other modules, the integration tests in particular, use the deployed
// definitions instead of copies.
//
// The files are Go text/template sources the operator renders before
// applying them; see subroutines.ReplaceTemplate for the function map.
package manifests

import (
	"embed"
	"io/fs"
)

//go:embed kcp features k8s
var FS embed.FS

// KCP is the manifests/kcp tree, rooted at that directory.
func KCP() fs.FS { return sub("kcp") }

// Features is the manifests/features tree, rooted at that directory.
func Features() fs.FS { return sub("features") }

// K8s is the manifests/k8s tree, rooted at that directory.
func K8s() fs.FS { return sub("k8s") }

func sub(dir string) fs.FS {
	f, err := fs.Sub(FS, dir)
	if err != nil {
		panic(err) // the directories are embedded above
	}
	return f
}
