# Giner

Giner is simple identicon gengerator

## Support image format

- ✅ PNG
- ✅ JPG
- ✏️ SVG (planning)

## Installation

```go
go get github.com/HungTP-Play/ginger-go@0.1.2
```

## Usage

```go
import (
 "log"

 "github.com/HungTP-Play/ginger-go/constant"
 "github.com/HungTP-Play/ginger-go/ginger"
)

func main() {
 info := "Tom_Ford_Black_Orchid"
 outputDir := "output"
 imageSize := 500
 padding := imageSize * 0.05
 _, err := ginger.DrawIdenticon(info, outputDir, imageSize, padding, constant.JPG)
 if err != nil {
  log.Printf(err.Error())
 }
}
```

## Output Example

![Identicon 1](https://i.ibb.co/f4J6wsh/Bacarat-Rouge-540.jpg)
![Identicon 2](https://i.ibb.co/KGfkvpx/Le-Labo-Santal-33.jpg)
![Identicon 3](https://i.ibb.co/3fwcJGS/Tom-Ford-Grey-Vetiver.jpg)
