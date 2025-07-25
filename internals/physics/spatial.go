package physics

import "scroller_game/internals/config"

// SpatialGrid provides efficient spatial partitioning for collision detection
type SpatialGrid struct {
	CellSize int
	Cols     int
	Rows     int
	Grid     [][]interface{}
}

// NewSpatialGrid creates a new spatial grid for collision optimization
func NewSpatialGrid(cellSize int) *SpatialGrid {
	cols := (config.ScreenWidth + cellSize - 1) / cellSize
	rows := (config.ScreenHeight + cellSize - 1) / cellSize

	grid := make([][]interface{}, cols*rows)
	for i := range grid {
		grid[i] = make([]interface{}, 0, 4) // Pre-allocate capacity for typical use
	}

	return &SpatialGrid{
		CellSize: cellSize,
		Cols:     cols,
		Rows:     rows,
		Grid:     grid,
	}
}

// Clear resets the grid for the next frame
func (sg *SpatialGrid) Clear() {
	for i := range sg.Grid {
		sg.Grid[i] = sg.Grid[i][:0] // Reuse underlying arrays
	}
}

// Insert adds an object to the appropriate grid cells
func (sg *SpatialGrid) Insert(x, y, width, height float32, object interface{}) {
	minCol := int(x) / sg.CellSize
	maxCol := int(x+width) / sg.CellSize
	minRow := int(y) / sg.CellSize
	maxRow := int(y+height) / sg.CellSize

	// Clamp to grid bounds
	if minCol < 0 {
		minCol = 0
	}
	if maxCol >= sg.Cols {
		maxCol = sg.Cols - 1
	}
	if minRow < 0 {
		minRow = 0
	}
	if maxRow >= sg.Rows {
		maxRow = sg.Rows - 1
	}

	for col := minCol; col <= maxCol; col++ {
		for row := minRow; row <= maxRow; row++ {
			idx := row*sg.Cols + col
			sg.Grid[idx] = append(sg.Grid[idx], object)
		}
	}
}

// Query returns all objects in cells that overlap with the given rectangle
func (sg *SpatialGrid) Query(x, y, width, height float32) []interface{} {
	minCol := int(x) / sg.CellSize
	maxCol := int(x+width) / sg.CellSize
	minRow := int(y) / sg.CellSize
	maxRow := int(y+height) / sg.CellSize

	// Clamp to grid bounds
	if minCol < 0 {
		minCol = 0
	}
	if maxCol >= sg.Cols {
		maxCol = sg.Cols - 1
	}
	if minRow < 0 {
		minRow = 0
	}
	if maxRow >= sg.Rows {
		maxRow = sg.Rows - 1
	}

	var results []interface{}
	seen := make(map[interface{}]bool) // Prevent duplicates

	for col := minCol; col <= maxCol; col++ {
		for row := minRow; row <= maxRow; row++ {
			idx := row*sg.Cols + col
			for _, obj := range sg.Grid[idx] {
				if !seen[obj] {
					seen[obj] = true
					results = append(results, obj)
				}
			}
		}
	}

	return results
}
