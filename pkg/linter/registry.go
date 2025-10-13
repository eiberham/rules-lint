/*
Copyright 2025 The Rules-Lint Authors.

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

package linter

type Registry struct {
	rules map[string]Rule
}

func NewRegistry() *Registry {
	return &Registry{
		rules: make(map[string]Rule),
	}
}

func (r *Registry) Register(name string, rule Rule) {
	r.rules[name] = rule
}

func (r *Registry) Get(name string) (Rule, bool) {
	rule, exists := r.rules[name]
	return rule, exists
}
