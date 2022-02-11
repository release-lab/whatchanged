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
	Version      []string   `json:"version"`
	Branch       string     `json:"branch"`
	Format       EnumFormat `json:"format"`
	SkipFormat   bool       `json:"skip_format"`
	Preset       EnumPreset `json:"preset"`        // Priority: 1 The higher the level, the more priority
	TemplateFile string     `json:"template_file"` // Priority: 2
	Template     string     `json:"template"`      // Priority: 3
	Silent       bool       `json:"silent"`
}
