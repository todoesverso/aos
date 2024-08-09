package models

type PositionalArgument struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Argument struct {
	Option      string              `yaml:"option"`
	Value       string              `yaml:"value"`
	Raw         string              `yaml:"raw"`
	Description string              `yaml:"description"`
	Positional  *PositionalArgument `yaml:"positional,omitempty"`
}

type YamlInput struct {
	Description string `yaml:"description"`
	Command     struct {
		Exec        string `yaml:"exec"`
		Description string `yaml:"description"`
	} `yaml:"command"`
	Arguments []Argument `yaml:"arguments"`
}
