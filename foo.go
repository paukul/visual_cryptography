package main
import (
  "image"
  "image/png"
  "image/color"
  "os"
  "bufio"
  "math/rand"
)

type CypherImage interface {
  Set(int, int, color.Color)
}

type CypherPixel []bool

func BlackPair() (CypherPixel, CypherPixel) {
  aPixel := NewCypherPixel()
  bPixel := make(CypherPixel, 4)
  for i := range aPixel {
    bPixel[i] = !aPixel[i]
  }
  return aPixel, bPixel
}

func WhitePair() (CypherPixel, CypherPixel) {
  aPixel := NewCypherPixel()
  bPixel := make(CypherPixel, 4)
  copy(bPixel, aPixel)
  return aPixel, bPixel
}

func NewCypherPixel() (CypherPixel) {
  px   := make(CypherPixel, 4)
  perm := rand.Perm(4)
  for i := range perm {
    px[i] = perm[i] % 2 == 0
  }
  return px
}

func main() {
  var templatePath string
  if len(os.Args) == 1 {
    templatePath = "tux.png"
  } else {
    templatePath = os.Args[1]
  }
  template, err := readTemplate(templatePath)
  if err != nil { panic(err) }
  templateBounds := template.Bounds()
  sourceHeight   := templateBounds.Max.Y
  sourceWidth    := templateBounds.Max.X
  cypherHeight   := sourceHeight * 2
  cypherWidth   := sourceWidth * 2

  image1 := image.NewRGBA(image.Rect(0, 0, cypherWidth, cypherHeight))
  image2 := image.NewRGBA(image.Rect(0, 0, cypherWidth, cypherHeight))
  for x := 0; x <= sourceWidth; x++ {
    for y := 0; y <= sourceHeight; y++ {
      setPixel(x, y, template.At(x, y), image1, image2)
    }
  }

  writeCypers(image1, image2)
}

func setPixel(x int, y int, col color.Color, cyperImage1 CypherImage, cyperImage2 CypherImage) {
  var px1, px2 CypherPixel
  grayValue := color.GrayModel.Convert(col).(color.Gray).Y
  if grayValue < 127 {
    px1, px2 = BlackPair()
  } else {
    px1, px2 = WhitePair()
  }

  i := 0
  for cX := 2*x; cX <= 2*x+1; cX++ {
    for cY := 2*y; cY <= 2*y+1; cY++ {
      color1 := colorForPixel(px1[i])
      color2 := colorForPixel(px2[i])
      cyperImage1.Set(cX, cY, color1)
      cyperImage2.Set(cX, cY, color2)
      i++
    }
  }
}

func colorForPixel(isBlack bool) (color.RGBA) {
  col := color.RGBA{}
  if isBlack { col.A = 255 }
  return col
}

func writeCypers(image1 image.Image, image2 image.Image) {
  writeCyper(image1, "foo1.png")
  writeCyper(image2, "foo2.png")
}

func writeCyper(image image.Image, filename string) {
  file, err := os.Create("out/" + filename)
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

func readTemplate(path string) (image.Image, error) {
  file, err := os.Open(path)
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
