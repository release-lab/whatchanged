package whatchanged

type EnumFormat string
type EnumPreset string

const (
	// format
	FormatMarkdown EnumFormat = "md"
	FormatJSON     EnumFormat = "json"
	// preset
	PresetDefault EnumPreset = "default"
	PresetFull    EnumPreset = "full"
)

type Options struct {
	Version      []string
	Branch       string
	Format       EnumFormat
	SkipFormat   bool
	Preset       EnumPreset // Priority: 1 The higher the level, the more priority
	TemplateFile string     // Priority: 2
	Template     string     // Priority: 3
	Silent       bool
}
