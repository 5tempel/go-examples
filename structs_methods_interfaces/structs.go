package structs
/*
	Perimeter takes a width and a height of a rectangle and returns its perimeter (perimeter = 2 * (width + height))
*/
func Perimeter (width float64, height float64) float64 {
	return 2 * (width + height)
}

/*
	Area takes a wifht and a height of a rectangle and returns its area (area = width * height )
*/
func Area (width, height float64) float64 {
	return width * height
}