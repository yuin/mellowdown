package builder

type config struct {
	Site  siteConfig
	Theme themeConfig
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
