// Copyright 2019 Istio Authors
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

package fixtures

import (
	"istio.io/istio/galley/pkg/config/analysis"
	"istio.io/istio/galley/pkg/config/analysis/diag"
	"istio.io/istio/galley/pkg/config/meta/schema/collection"
	"istio.io/istio/galley/pkg/config/resource"
)

// Context is a test fixture of analysis.Context
type Context struct {
	Entries []*resource.Entry
	Reports []diag.Message
}

var _ analysis.Context = &Context{}

// Report implements analysis.Context
func (ctx *Context) Report(c collection.Name, t diag.Message) {
	ctx.Reports = append(ctx.Reports, t)
}

// Find implements analysis.Context
func (ctx *Context) Find(c collection.Name, name resource.Name) *resource.Entry { return nil }

// Exists implements analysis.Context
func (ctx *Context) Exists(c collection.Name, name resource.Name) bool { return false }

// ForEach implements analysis.Context
func (ctx *Context) ForEach(c collection.Name, fn analysis.IteratorFn) {
	for _, r := range ctx.Entries {
		fn(r)
	}
}

// Canceled implements analysis.Context
func (ctx *Context) Canceled() bool { return false }
