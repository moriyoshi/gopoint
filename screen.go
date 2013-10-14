package main

import (
    gl "github.com/chsc/gogl/gl21"
    "image"
    "os"
    "bufio"
    "image/png"
)

type Screen struct {
    scene *Scene
    image *image.RGBA
    textureId gl.Uint
}

func newScreen(scene *Scene) *Screen {
    retval := Screen {
        scene: scene,
        image: image.NewRGBA(image.Rect(0, 0, scene.size.X, scene.size.Y)),
        textureId: scene.newTexture(),
    }

    return &retval
}

func (screen *Screen) destroy() {
    screen.scene.deleteTexture(screen.textureId)
}

func (screen *Screen) bounds() image.Rectangle {
    return screen.image.Bounds()
}

func (screen *Screen) realize() {
    screenSize := screen.bounds().Size()
    gl.BindTexture(gl.TEXTURE_2D, screen.textureId)
    gl.TexImage2D(
        gl.TEXTURE_2D,
        0, 4 /* gl.RGBA */,
        gl.Sizei(screenSize.X), gl.Sizei(screenSize.Y),
        0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(&screen.image.Pix[0]),
    )
}

func (screen *Screen) writeAs(file string) error {
    f, err := os.Create(file)
    if err != nil { return err }
    defer f.Close()
    b := bufio.NewWriter(f)
    err = png.Encode(b, screen.image)
    if err != nil { return err }
    err = b.Flush()
    if err != nil { return err }
    return nil
}
