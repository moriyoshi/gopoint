package main

import (
    "code.google.com/p/freetype-go/freetype"
    "code.google.com/p/freetype-go/freetype/truetype"
    "image"
    "image/draw"
    "image/color"
)

type DrawingContext struct {
    screen *Screen
    font *truetype.Font
    fontSize float64
    foreground image.Image
    background image.Image
    fctx *freetype.Context
}

func newDrawingContext(screen *Screen) *DrawingContext {
    return &DrawingContext {
        screen: screen,
        font: nil,
        fontSize: 12,
        foreground: image.Black,
        background: image.NewUniform(color.NRGBA { 255, 255, 255, 255 }),
        fctx: nil,
    }
}

func (dctx *DrawingContext) clear() {
    draw.Draw(dctx.screen.image, dctx.screen.bounds(), dctx.background, image.ZP, draw.Src)
}

func (dctx *DrawingContext) populateFreeTypeContext() {
    if (dctx.fctx != nil) { return }
    fctx := freetype.NewContext()
    fctx.SetDPI(90)
    fctx.SetFont(dctx.font)
    fctx.SetFontSize(dctx.fontSize)
    fctx.SetClip(dctx.screen.bounds())
    fctx.SetDst(dctx.screen.image)
    fctx.SetSrc(dctx.foreground)
    dctx.fctx = fctx
}

func (dctx *DrawingContext) setFont(font *truetype.Font, size float64) {
    dctx.font = font
    dctx.fontSize = size
    dctx.fctx = nil
}

func (dctx *DrawingContext) drawString(position image.Point, text string) error {
    dctx.populateFreeTypeContext()
    pt := freetype.Pt(position.X, position.Y + (int(dctx.fctx.PointToFix32(dctx.fontSize)) >> 8))
    _, err := dctx.fctx.DrawString(text, pt)
    return err
}

func (dctx *DrawingContext) drawImage(position image.Point, img image.Image) error {
    size := img.Bounds().Size()
    draw.Draw(dctx.screen.image, image.Rect(position.X, position.Y, position.X + size.X, position.Y + size.Y), img, image.ZP, draw.Src)
    return nil
}
