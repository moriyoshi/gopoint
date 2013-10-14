package main

import (
    gl "github.com/chsc/gogl/gl21"
)

type UintVec struct {
    array []gl.Uint
}

func (vec *UintVec) Len() int {
    return len(vec.array)
}

func (vec *UintVec) Add(item gl.Uint) {
    vec.array = append(vec.array, item)
}

func (vec *UintVec) IndexOf(item gl.Uint) int {
    for i, _item := range(vec.array) {
        if _item == item { return i }
    }
    return -1
}

func (vec *UintVec) Remove(item gl.Uint) bool {
    idx := vec.IndexOf(item)
    if idx < 0 { return false }
    copy(vec.array[idx:], vec.array[idx + 1:])
    return true
}

func (vec *UintVec) Slice() []gl.Uint {
    return vec.array
}
