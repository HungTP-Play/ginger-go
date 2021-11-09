# Ginger

Giner is simple identicon gengerator

This project is inspired by [Github Indenticon](https://github.blog/2013-08-14-identicons/), but the current implement is inspired from [Don Park: 9-block IP Identification](https://web.archive.org/web/20080703155519/http://www.docuverse.com/blog/donpark/2007/01/18/visual-security-9-block-ip-identification).

## Support image format

- ✅ PNG
- ✅ JPG
- ✏️ SVG (planning)

## Support Identicon Style

- ✅ Don Park: 9-block IP Identification style
- ✅ Github Identicon style
- ✏️ [Jdenticon](https://jdenticon.com/) style (planning)

## Installation

```shell
go get github.com/HungTP-Play/ginger-go@0.1.2
```

## Usage

### Don Park Style

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
 padding := float64(imageSize) * 0.05
 multiColor := true
 _, err := ginger.DrawIdenticon(info, outputDir, imageSize, int(padding), constant.JPG, multiColor)
 if err != nil {
  log.Printf(err.Error())
 }
}
```

Output Example

One color

![Identicon 1](https://i.ibb.co/N3FnVPN/Identicon-Bacarat-Rouge-540.png)
![Identicon 2](https://i.ibb.co/5n3TZJY/Identicon-Bvlgari-Aqva-Pour-Homme.png)
![Identicon 3](https://i.ibb.co/BtP3GYh/Identicon-Clue-De-Nuit-Pour-Homme.png)
![Identicon 4](https://i.ibb.co/7CjKnL1/Identicon-Vesace-Eros.png)

Multi-color

<img src="https://i.ibb.co/ysJ3J6V/Identicon-Bacarat-Rouge-540.png" width="200"/>
<img src="https://i.ibb.co/2MZ1Nsb/Identicon-Bvlgari-Aqva-Pour-Homme.png" width="200"/>
<img src="https://i.ibb.co/FHQ8vdp/Identicon-Tom-Ford-Black-Orchid.png" width="200"/>
<img src="https://i.ibb.co/y6smtWb/Identicon-Vesace-Eros.png" width="200"/>

### Github Style

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
 padding := float64(imageSize) * 0.05
 _, err := ginger.DrawGithubIdenticon(info, outputDir, imageSize, int(padding), constant.JPG)
 if err != nil {
  log.Printf(err.Error())
 }
}
```

Output example

![Github Identicon 1](https://i.ibb.co/wpxB7xZ/Github-Bacarat-Rouge-540.png)
![Github Identicon 2](https://i.ibb.co/LNPwxMB/Github-Bvlgari-Aqva-Pour-Homme.png)
![Github Identicon  3](https://i.ibb.co/0GYMDZK/Github-Clue-De-Nuit-Pour-Homme.png)
![Github Identicon  4](https://i.ibb.co/X4NB6w1/Github-Vesace-Eros.png)

![Alt](https://repobeats.axiom.co/api/embed/b117db333826b60ce8518ff41bc2729aff1e9b08.svg "Repobeats analytics image")
