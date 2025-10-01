package linter

import (
	"os"
)

type LintResult struct {
	FilePath string
	Issues   []string
	Error    error
}

var registry = NewRegistry()

func Init() {
	registry.Register(
		"checkTemplateVars",
		&LineRule{
			BaseRule: BaseRule{RuleType: Line},
			Handler:  CheckTemplateVars,
		},
	)

	registry.Register(
		"checkUnusedContextKeys",
		&FileRule{
			BaseRule: BaseRule{RuleType: File},
			Handler:  CheckUnusedContextKeys,
		},
	)

	registry.Register(
		"checkAsyncIncongruence",
		&FileRule{
			BaseRule: BaseRule{RuleType: File},
			Handler:  CheckAsyncIncongruence,
		},
	)
}

func Run(path string, cfg *Config, results chan<- LintResult) {
	// Placeholder for linting logic
	// This function would read the file, apply the rules from cfg,
	// and report any issues found.

	file, err := os.Open(path)
	if err != nil {
		results <- LintResult{FilePath: path, Error: err}
		return
	}
	defer file.Close()

	rules := GetEnabledRules(cfg)
	issues := []string{}
	for _, ruleName := range rules {
		rule, exists := registry.Get(ruleName)
		if !exists {
			continue
		}

		processor := &BaseHandler{}
		processor.
			Next(&LineHandler{}).
			Next(&FileHandler{})

		issues = append(issues, processor.Handle(rule, cfg, file)...)
	}

	results <- LintResult{FilePath: path, Issues: issues}
}
