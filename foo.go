package main
import (
  "image"
  "image/png"
  "image/color"
  "os"
  "fmt"
  "bufio"
)

type CypherImage interface {
  Set(int, int, color.Color)
}

func main() {
  template, err := readTemplate()
  if err != nil { panic(err) }
  templateBounds := template.Bounds()
  sourceHeight   := templateBounds.Max.Y
  sourceWidth    := templateBounds.Max.X
  fmt.Printf("sourceWidth: %d\n", sourceWidth)
  fmt.Printf("sourceHeight: %d\n", sourceHeight)
  cypherHeight   := sourceHeight * 2
  cypherWidth   := sourceWidth * 2

  image := image.NewRGBA(image.Rect(0, 0, cypherWidth, cypherHeight))
  for x := 0; x <= sourceWidth; x++ {
    for y := 0; y <= sourceHeight; y++ {
      setPixel(x, y, template.At(x, y), image)
    }
  }

  writeCyper(image)
}

func setPixel(x int, y int, color color.Color, cyperImage CypherImage) {
  for cX := 2*x; cX <= 2*x+1; cX++ {
    for cY := 2*y; cY <= 2*y+1; cY++ {
      cyperImage.Set(cX, cY, color)
    }
  }
}

func writeCyper(image image.Image) {
  file, err := os.Create("out/foo.png")
  if err != nil { panic(err) }
  defer func() {
    if err := file.Close(); err != nil {
      panic(err)
    }
  }()

  writer := bufio.NewWriter(file)
  png.Encode(writer, image)
  if err = writer.Flush(); err != nil { panic(err) }
}

func readTemplate() (image.Image, error) {
  file, err := os.Open("tux.png")
  if err != nil {
    return nil, err
  }

  defer func() {
    if err := file.Close(); err != nil {
      panic(err)
    }
  }()

  reader := bufio.NewReader(file)
  return png.Decode(reader)
}
