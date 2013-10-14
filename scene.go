package main

import (
    gl "github.com/chsc/gogl/gl21"
    "image"
)

type Scene struct {
    textures UintVec
    size image.Point
}

func newScene(size image.Point) *Scene {
    gl.Enable(gl.TEXTURE_2D)
    gl.Enable(gl.DEPTH_TEST)
    gl.Enable(gl.LIGHTING)
    gl.ClearColor(1., 1., 1., 1.)
    gl.ClearDepth(1)
    gl.DepthFunc(gl.LEQUAL)

    ambientColor := []gl.Float { 0.5, 0.5, 0.5, 1 }
    diffusionColor := []gl.Float { 1, 1, 1, 1 }
    lightPosition := []gl.Float { 0, 0, 10, 0 }
    gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambientColor[0])
    gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffusionColor[0])
    gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
    gl.Enable(gl.LIGHT0)

    gl.Viewport(0, 0, gl.Sizei(size.X), gl.Sizei(size.Y))
    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()
    gl.Frustum(-1, 1, -1, 1, 1., 20.0)
    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()

    return &Scene { size: size }
}

func (scene *Scene) newTexture() gl.Uint {
    var retval gl.Uint
    gl.GenTextures(1, &retval)
    gl.BindTexture(gl.TEXTURE_2D, retval)
    gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
    gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
    scene.textures.Add(retval)
    return retval
}

func (scene *Scene) deleteTexture(textureId gl.Uint) {
    scene.textures.Remove(textureId)
}

func (scene *Scene) destroy() {
    gl.DeleteTextures(gl.Sizei(scene.textures.Len()), &scene.textures.Slice()[0])
}

