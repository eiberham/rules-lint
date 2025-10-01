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
