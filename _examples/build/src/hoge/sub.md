# sub

``` ppt
file: images.pptx
shape: image1
```

```uml
@startuml
[*] --> active

active -right-> inactive : disable
inactive -left-> active  : enable

inactive --> closed  : close
active --> closed  : close

closed --> [*]

@enduml
```

[ ](#custom-id)

![sub桜][a]
[a]:./sub_sakura.jpg

