package colors

import "image/color"

const max = 255

// const half = uint(max / 2)

var Black = color.Black
var White = color.White
var Green = color.NRGBA{R: 0, G: max, B: 0, A: max}
var Blue = color.NRGBA{R: 0, G: 0, B: max, A: max}
var Teal = color.NRGBA{R: 0, G: 127, B: 127, A: max}
var Cyan = color.NRGBA{R: 0, G: max, B: max, A: max}
var Orange = color.NRGBA{R: max, G: 165, B: 0, A: max}
var Yellow = color.NRGBA{R: max, G: max, B: 0, A: max}
var Amethyst = color.NRGBA{R: 153, G: 102, B: 204, A: max}
