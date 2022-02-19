package gfx

type TileMap struct {
	tileWidth  int
	tileHeight int
	width      int
	height     int
	tiles      [][]*Sprite
}

func NewTileMap(tileWidth, tileHeight int) *TileMap {
	return &TileMap{
		tileWidth:  tileWidth,
		tileHeight: tileHeight,
	}
}

func (tm *TileMap) SetTile(x, y int, sprite *Sprite) {
	if tm.width <= x {
		tm.width = x + 1
	}

	if tm.height <= y {
		tm.height = y + 1
	}

	for len(tm.tiles) < tm.height {
		tm.tiles = append(tm.tiles, make([]*Sprite, tm.width))
	}

	for y := range tm.tiles {
		for len(tm.tiles[y]) < tm.width {
			tm.tiles[y] = append(tm.tiles[y], nil)
		}
	}

	tm.tiles[y][x] = sprite
}

func (tm *TileMap) Draw() {
	for y, row := range tm.tiles {
		for x, tile := range row {
			if tile == nil {
				continue
			}

			tile.Draw(float64(x*tm.tileWidth), float64(y*tm.tileHeight))
		}
	}
}
