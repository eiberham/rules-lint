package linter

type RuleType int

const (
	Line RuleType = iota
	File
)

type RuleContext struct {
	Config     *Config
	Content    string
	LineNumber int
	FilePath   string
}

type Rule interface {
	Validate(ctx *RuleContext) string
	Type() RuleType
}

type BaseRule struct {
	RuleType RuleType
}

func (r *BaseRule) Type() RuleType {
	return r.RuleType
}

type LineRule struct {
	BaseRule
	Handler func(line string, lineNumber int, ctx *RuleContext) string
}

func (r *LineRule) Validate(ctx *RuleContext) string {
	if r.Handler != nil {
		return r.Handler(ctx.Content, ctx.LineNumber, ctx)
	}
	return ""
}

type FileRule struct {
	BaseRule
	Handler func(content string, ctx *RuleContext) string
}

func (r *FileRule) Validate(ctx *RuleContext) string {
	if r.Handler != nil {
		return r.Handler(ctx.Content, ctx)
	}
	return ""
}
