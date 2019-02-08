# インデックス

```toc
depth: 2
title: Table of contents
```

[ ](#label1)

## level2 {#custom-id}
## level2-1 

{{ label("ラベル#label1") }}

[ラベル2]{#label2}

### level3 

- [リンク](./hoge/sub.md)

ほげほげ

![画像](./gazou.jpg)

![桜][a]
[a]:./sakura.jpg

## level2-2


``` go
func (r *ChromaRenderer) Accept(info string) bool {
	return len(info) != 0
}

func (r *ChromaRenderer) Render(w io.Writer, node *blackfriday.Node) error {
	lexer := lexers.Get(getLang(noge))
	formatter := html.New(html.WithClasses())
	if err := formatter.Format(w, styles.Monokai, lexer); err != nil {
		return err
	}
	return nil
}
```
