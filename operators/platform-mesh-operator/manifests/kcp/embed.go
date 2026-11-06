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

// Package kcp exposes the kcp manifests the platform-mesh-operator applies
// to bootstrap a platform (workspaces, WorkspaceTypes, APIExports, bindings,
// ...) as an embedded file system, so other modules, the integration tests
// in particular, use the deployed definitions instead of copies.
package kcp

import "embed"

// Manifests holds every file under manifests/kcp, keyed by its path relative
// to this directory, e.g. "02-root/workspace-type-org.yaml". Files are Go
// text/template sources the operator renders before applying; see
// subroutines.ReplaceTemplate for the function map.
//
//go:embed *.yaml */*.yaml
var Manifests embed.FS
