package main

import (
    gl "github.com/chsc/gogl/gl21"
    "github.com/jteeuwen/glfw"
    "code.google.com/p/freetype-go/freetype"
    "code.google.com/p/freetype-go/freetype/truetype"
    "io/ioutil"
    "image"
    "image/png"
    "os"
    "math"
)

func readFont(fileName string) (*truetype.Font, error) {
    bytes, err := ioutil.ReadFile(fileName)
    if err != nil { return nil, err }
    font, err := freetype.ParseFont(bytes)
    if err != nil { return nil, err }
    return font, nil
}

func drawSceneQuad(screens []*Screen) {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

    gl.Color4f(1, 1, 1, 1)

    if len(screens) > 0 {
        func () {
            gl.BindTexture(gl.TEXTURE_2D, screens[0].textureId)
            gl.Begin(gl.QUADS)
            defer gl.End()
            gl.Normal3f(0, 0, 1)
            gl.TexCoord2f(0, 1)
            gl.Vertex3f(-1, -1, 1)
            gl.TexCoord2f(1, 1)
            gl.Vertex3f(1, -1, 1)
            gl.TexCoord2f(1, 0)
            gl.Vertex3f(1, 1, 1)
            gl.TexCoord2f(0, 0)
            gl.Vertex3f(-1, 1, 1)
        }()
    }
    if len(screens) > 1 {
        func () {
            gl.BindTexture(gl.TEXTURE_2D, screens[1].textureId)
            gl.Begin(gl.QUADS)
            defer gl.End()
            gl.Normal3f(1, 0, 0)
            gl.TexCoord2f(0, 1)
            gl.Vertex3f(1, -1, 1)
            gl.TexCoord2f(1, 1)
            gl.Vertex3f(1, -1, -1)
            gl.TexCoord2f(1, 0)
            gl.Vertex3f(1, 1, -1)
            gl.TexCoord2f(0, 0)
            gl.Vertex3f(1, 1, 1)
        }()
    }
/*
    func () {
        gl.BindTexture(gl.TEXTURE_2D, screens[2].textureId)
        gl.Begin(gl.QUADS)
        defer gl.End()
        gl.Normal3f(0, 0, -1)
        gl.TexCoord2f(0, 1)
        gl.Vertex3f(1, -1, -1)
        gl.TexCoord2f(1, 1)
        gl.Vertex3f(-1, -1, -1)
        gl.TexCoord2f(1, 0)
        gl.Vertex3f(-1, 1, -1)
        gl.TexCoord2f(0, 0)
        gl.Vertex3f(1, 1, -1)
    }()
    func () {
        gl.BindTexture(gl.TEXTURE_2D, screens[3].textureId)
        gl.Begin(gl.QUADS)
        defer gl.End()
        gl.Normal3f(-1, 0, 0)
        gl.TexCoord2f(0, 1)
        gl.Vertex3f(-1, -1, -1)
        gl.TexCoord2f(1, 1)
        gl.Vertex3f(-1, -1, 1)
        gl.TexCoord2f(1, 0)
        gl.Vertex3f(-1, 1, 1)
        gl.TexCoord2f(0, 0)
        gl.Vertex3f(-1, 1, -1)
    }()
*/
}

type Object struct {
    position image.Point
    size float64
    text string
}

type Page struct {
    texts []Object
}

var pages = []Page {
    Page {
        texts: []Object {
            Object {
                position: image.Point { 60, 340 },
                size: 60,
                text: "Goでプレゼンツールを作った話",
            },
            Object {
                position: image.Point { 60, 450 },
                size: 30,
                text: "Moriyoshi Koizumi <mozo@mozo.jp>",
            },
        },
    },
    Page {
        texts: []Object {
            Object {
                position: image.Point { 500, 40 },
                size: 60,
                text: "自己紹介",
            },
            Object {
                position: image.Point { 60, 200 },
                size: 40,
                text: "・かつてGoにパッチを送ったりした",
            },
            Object {
                position: image.Point { 60, 300 },
                size: 40,
                text: "・日々の仕事はPythonを使っています",
            },
        },
    },
    Page {
        texts: []Object {
            Object {
                position: image.Point { 230, 40 },
                size: 60,
                text: "このツールを作った動機",
            },
            Object {
                position: image.Point { 60, 200 },
                size: 40,
                text: "・PHPでプレゼンツールを作った事がある",
            },
            Object {
                position: image.Point { 60, 300 },
                size: 40,
                text: "・ジョークで作ってみたものの、意外と便利であった",
            },
            Object {
                position: image.Point { 60, 400 },
                size: 40,
                text: "・Goでもやってみよう",
            },
        },
    },
    Page {
        texts: []Object {
            Object {
                position: image.Point { 340, 40 },
                size: 60,
                text: "このツールの特徴",
            },
            Object {
                position: image.Point { 60, 200 },
                size: 40,
                text: "・プレゼン内容をソースコード自体に記述",
            },
            Object {
                position: image.Point { 60, 300 },
                size: 40,
                text: "・プレゼン自体が実行ファイルなので",
            },
            Object {
                position: image.Point { 112, 360 },
                size: 40,
                text: "配布が容易",
            },
            Object {
                position: image.Point { 60, 460 },
                size: 40,
                text: "・実行ファイルは静的リンクされているので",
            },
            Object {
                position: image.Point { 112, 520 },
                size: 40,
                text: "見る人も安心",
            },
        },
    },
    Page {
        texts: []Object {
            Object {
                position: image.Point { 340, 40 },
                size: 60,
                text: "使っている技術",
            },
            Object {
                position: image.Point { 60, 200 },
                size: 40,
                text: "・GoGL",
            },
            Object {
                position: image.Point { 112, 260 },
                size: 30,
                text: "http://github.com/chsc/gogl",
            },
            Object {
                position: image.Point { 60, 360 },
                size: 40,
                text: "・GoGL",
            },
            Object {
                position: image.Point { 112, 420 },
                size: 30,
                text: "http://github.com/jteeuwen/glfw",
            },
            Object {
                position: image.Point { 60, 520 },
                size: 40,
                text: "・freetype-go",
            },
            Object {
                position: image.Point { 112, 580 },
                size: 30,
                text: "http://code.google.com/p/freetype-go",
            },
        },
    },
    Page {
        texts: []Object {
            Object {
                position: image.Point { 300, 40 },
                size: 60,
                text: "やっていること",
            },
            Object {
                position: image.Point { 60, 200 },
                size: 40,
                text: "1. ウィンドウの生成 (glfw) とOpenGLの初期化",
            },
            Object {
                position: image.Point { 60, 300 },
                size: 40,
                text: "2. 各ページをfreetype-goを使って画像としてレンダ",
            },
            Object {
                position: image.Point { 112, 360 },
                size: 40,
                text: "リング",
            },
            Object {
                position: image.Point { 60, 440 },
                size: 40,
                text: "3. レンダリングした画像を元にテクスチャ生成",
            },
            Object {
                position: image.Point { 60, 540 },
                size: 40,
                text: "4. 直角に組み合わされた長方形にテクスチャを張り",
            },
            Object {
                position: image.Point { 112, 600 },
                size: 40,
                text: "つけ",
            },
            Object {
                position: image.Point { 60, 700 },
                size: 40,
                text: "5. シーンを表示",
            },
        },
    },
    Page {
        texts: []Object {
            Object {
                position: image.Point { 580, 400 },
                size: 60,
                text: "デモ",
            },
        },
    },
}

func readImage(file string) (image.Image, error) {
    f, err := os.Open(file)
    if err != nil { return nil, err }
    defer f.Close()
    return png.Decode(f)
}

func realMain(scene *Scene) {
    font, err := readFont("mplus-1c-regular.ttf")
    if err != nil {
        showError(err)
        return
    }

    gopher, err := readImage("gopher.png")
    if err != nil {
        showError(err)
        return
    }

    screens := make([]*Screen, 0)
    for _, page := range(pages) {
        screen := newScreen(scene)
        screens = append(screens, screen)
        dctx := newDrawingContext(screen)
        dctx.clear()
        for _, text := range(page.texts) {
            dctx.setFont(font, text.size)
            dctx.drawImage(image.Point { 900, 600 }, gopher)
            dctx.drawString(text.position, text.text)
        }
        screen.realize()
    }

    page := 0
    r := 0.
    state := "stop"
    for glfw.WindowParam(glfw.Opened) == 1 {
        if state == "next" {
            r += 3
            if r >= 90 {
                r = 0
                state = "stop"
                page += 1
            }
        }
        gl.MatrixMode(gl.MODELVIEW)
        gl.LoadIdentity()
        gl.Translatef(0, 0, gl.Float(-2.01 - math.Sqrt(2) * math.Cos(math.Pi * (math.Mod(r, 90.) - 45.) / 180.) + 1))
        gl.Rotatef(gl.Float(-r), 0, 1., 0)
        drawSceneQuad(screens[page:])
        glfw.SwapBuffers()
        if glfw.Key(glfw.KeyEnter) == glfw.KeyPress {
            state = "next"
        }
        if glfw.Key(glfw.KeyEsc) == glfw.KeyPress { break }
    }
}

func main() {
    if err := glfw.Init(); err != nil {
        showError(err)
        return
    }
    defer glfw.Terminate()

    glfw.OpenWindowHint(glfw.WindowNoResize, 1)


    vidmode := glfw.DesktopMode()
    width, height := vidmode.W, vidmode.H

    if err := glfw.OpenWindow(width, height, 0, 0, 0, 0, 16, 0, glfw.Windowed /*glfw.Fullscreen*/); err != nil {
        showError(err)
        return
    }

    defer glfw.CloseWindow()

    glfw.SetSwapInterval(1)
    glfw.SetWindowTitle("gopoint")

    if err := gl.Init(); err != nil {
       showError(err)
        return
    }

    scene := newScene(image.Point { X: width, Y: height })
    defer scene.destroy()

    realMain(scene)
}
