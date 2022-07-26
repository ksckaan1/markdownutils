# markdownutils

## Installation

```go get github.com/ksckaan1/markdownutils```

## Features

> Important: Outputs are stylized.

### Classic Markdown

```md
# Hello

## Hello

### Hello

`inline code`
.
.
.
.
```

### Classed Code Area

    ```go
    package main
    import "fmt"
    func main(){
        fmt.Println("hello!")
    }
    ```

### Classed Code Area with Title

    ```go this is a title
    package main
    import "fmt"
    func main(){
        fmt.Println("hello!")
    }
    ```

### Hint Boxes

    :::info INFO TITLE
    This is a info message.
    :::

    :::success SUCCESS TITLE
    This is a success message.
    :::
    
    :::warning WARNING TITLE
    This is a warning message.
    :::

    :::danger DANGER TITLE
    This is a danger message.
    :::


![](readme/hintboxes.png)

### YouTube Video

    :::youtube vjHAvx47d98:::

