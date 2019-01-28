# mellowdown
An extensible markdown processor with visualization supports(PlantUML, PowerPoint etc...)

## Install
Just download an executable file from releases, or

```go
go get -u github.com/yuin/mellowdown/cmd/mellowdown
```

## Usage

```
  -addr string
        address like localhost:8000, this enables livereloading
  -file string
        Markdown file(Required)
  -format string
        Output format(html or pdf) (default "html")
  -lua string
        comma separeted lua renderers
  -out string
        Output Directory(Optional)
  -plantuml string
        PlantUML executable file path(Optional). If this value is empty, PLANTUML_PATH envvar value will be used as an executable file path (default "plantuml")
  -style string
        Style (Optional, available styles:github) (default "github")
  -syntax-highlight string
        Syntax Highlightinging Style (Optional, available styles:abap,algol,algol_nu,arduino,autumn,borland,bw,colorful,dracula,emacs,friendly,fruity,github,igor,lovelace,manni,monokai,monokailight,murphy,native,paraiso-dark,paraiso-light,pastie,perldoc,pygments,rainbow_dash,rrt,solarized-dark,solarized-dark256,solarized-light,swapoff,tango,trac,vim,vs,xcode) (default "monokailight")
  -wkhtmltopdf string
        Wkhtmltopdf executable file path(Optional). If this value is empty, WKHTMLTOPDF_PATH envvar value will be used as an executable file path
```

### Convert a markdown file to an html file

```
mellowdown render -file input.md
```

### Convert a markdown file to a pdf file
Requirements:

- [wkhtmltopdf](https://wkhtmltopdf.org/)

```
mellowdown render -file input.md -format pdf
```

### Watch markdown files in the current directory

```
mellowdown render -addr ":8000"
```

Open `*.html` in your browser and edit markdown files, mellowdown will
automatically convert markdown files and live reload the browser.

### Embed shapes in a PowerPoint file into an html file

    ```ppt
    file: images.pptx
    shape: image1
    ```

- `file` is a PowerPoint file path relative to the markdown file.
- `shape` is a `Alt Text` of the shape.

### PlantUML

    ```uml
    @startuml
    
    actor Promoter
    actor Entrant
    
    Promoter --> (Create Event)
    Promoter -> (Attend Event)
    
    Entrant --> (Find Event)
    (Attend Event) <- Entrant
    
    (Attend Event) <.. (Create Member)  : <<include>>
    
    @enduml
    ```

### Sample PlantUML launcher for windows

```
@echo off
set GRAPHVIZ_DOT=C:\PATH_TO\dot.exe
set PLANTUML_JAR=%~dp0plantuml.jar

java -jar %PLANTUML_JAR% -charset UTF-8 %*
```
