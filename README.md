# hebiten

Utils for "ebiten" 2D graphics engine for Go programming language.

Update 2023: this library is now outdated. It was originally created for an old version of EbitenEngine. I am now archiving this repository

## Requirements:
* ebiten 2D graphics engine package https://github.com/hajimehoshi/ebiten
* my hinst/hgo package: https://github.com/hinst/hgo

## Overview:
* Math
    * Int2 {X Y}
    * BigFloat2 {X Y}
    * IntRect {X Y W H}
    * FloatRect {X Y W H}
    * TFloatColor {R G B A}
* Graphics
    * Draw
    * DrawTexts
    * TextureAtlas
* Reader for "Tiled" map editor data file
    * TiledData
    * TiledDataLoaders
* HealthBar class
    * BaseHealth, CurrentHealth, Texture, Color
