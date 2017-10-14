package main

type Field struct {
	field  [][]string
	width  int
	height int
}

func (self *Field) Initialize(width, height int) {
	if height > 0 && width > 0 && len(self.field) != height {
		self.width = width
		self.height = height
		for i := 0; i < height; i++ {
			boardRow := make([]string, width)
			self.field = append(self.field, boardRow)
		}
	}
}
