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
	Preset       Preset
	TemplateFile string // Priority is higher than preset
	Silent       bool
}
