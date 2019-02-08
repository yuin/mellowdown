package builder

type config struct {
	Site     siteConfig
	Resource resourceConfig
	Theme    themeConfig
}

type resourceConfig struct {
	Extras        []string
	ExtraPatterns []string

	Ignores        []string
	IgnorePatterns []string
}

type themeConfig struct {
	Name            string
	SyntaxHighlight string
}

type siteConfig struct {
	Name      string
	Copyright string
	Author    string
}
