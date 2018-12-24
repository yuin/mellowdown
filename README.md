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
-lua string
      comma separeted lua renderers
-out string
      Output Directory(Optional)
-plantuml string
      PlantUML Path(Optional). If PLANTUML_PATH envvar is not empty, this option will be overwritten by its value. (default "plantuml")
-style string
      Style (Optional, available styles:github) (default "github")
-syntax-highlight string
      Syntax Highlightinging Style (Optional, available styles:abap,algol,algol_nu,arduino,autumn,borland,bw,colorful,dracula,emacs,friendly,fruity,github,igor,lovelace,manni,monokai,monokailight,murphy,native,paraiso-dark,paraiso-light,pastie,perldoc,pygments,rainbow_dash,rrt,solarized-dark,solarized-dark256,solarized-light,swapoff,tango,trac,vim,vs,xcode) (default "monokailight")
```

### Convert a markdown file to an html file

```
mellowdown -file input.md
```

### Watch markdown files in the current directory

```
mellowdown -addr ":8000"
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
set PLANTUML_JAR=%~dp1plantuml.jar

java -jar %PLANTUML_JAR% -charset UTF-8 %*
```
