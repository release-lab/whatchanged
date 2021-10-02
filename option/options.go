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
	PresetSimple  Preset = "simple"
)

type Options struct {
	Version      []string
	Branch       string
	Format       Format
	SkipFormat   bool
	Preset       Preset // Priority: 1 The higher the level, the more priority
	TemplateFile string // Priority: 2
	Template     string // Priority: 3
	Silent       bool
}
