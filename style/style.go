package style

type styles map[string]string

var Styles styles

func init() {
	Styles = styles{}
	Styles["github"] = githubStyle
}

func Get(name string) string {
	v, ok := Styles[name]
	if !ok {
		return ""
	}
	return v
}

func Names() []string {
	ret := []string{}
	for k, _ := range Styles {
		ret = append(ret, k)
	}
	return ret
}
