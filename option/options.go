package option

type Format string
type Preset string

const (
	// format
	FormatMarkdown Format = "md"
	FormatJSON     Format = "json"
	// preset
	PresetDefault Preset = "default"
	PresetFull    Preset = "full"
)

type Options struct {
	Version      string
	Format       Format
	Preset       Preset // Priority: 1 The higher the level, the more priority
	TemplateFile string // Priority: 2
	Template     string // Priority: 3
	Silent       bool
}
