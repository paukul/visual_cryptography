package main
import (
  "image"
  "image/png"
  "image/color"
  "os"
  "bufio"
)

func main() {
  height, width := 480, 640
  image := image.NewRGBA(image.Rect(0, 0, width, height))

  for x := 0; x <= width; x++ {
    for y := 0; y <= height; y++ {
      alphaValue := 255
      if y % 2 == 0 {
        alphaValue = 0
      }
      image.SetRGBA(x, y, color.RGBA{R: 0, G: 0, B: 0, A: uint8(alphaValue)})
    }
  }

  file, err := os.Create("foo.png")
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

func readTemplate() (Image, Error) {
  file, err := os.Open("tux.png")
  if err != nil {
    return nil, err
  }
}
