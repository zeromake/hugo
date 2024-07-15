// Copyright 2018 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resources

import (
	"context"

	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/tpl/css"
	"github.com/gohugoio/hugo/tpl/internal"
	"github.com/gohugoio/hugo/tpl/js"
)

const name = "resources"

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		ctx, err := New(d)
		if err != nil {
			// TODO(bep) no panic.
			panic(err)
		}

		ns := &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func(cctx context.Context, args ...any) (any, error) { return ctx, nil },
			OnCreated: func(m map[string]any) {
				for _, v := range m {
					switch v := v.(type) {
					case *css.Namespace:
						ctx.cssNs = v
					case *js.Namespace:
						ctx.jsNs = v
					}
				}
				if ctx.cssNs == nil {
					panic("css namespace not found")
				}
				if ctx.jsNs == nil {
					panic("js namespace not found")
				}
			},
		}

		ns.AddMethodMapping(ctx.Get,
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.GetRemote,
			nil,
			[][2]string{},
		)

		// Add aliases for the most common transformations.

		ns.AddMethodMapping(ctx.Fingerprint,
			[]string{"fingerprint"},
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Minify,
			[]string{"minify"},
			[][2]string{},
		)

		return ns
	}

	internal.AddTemplateFuncsNamespace(f)
}
