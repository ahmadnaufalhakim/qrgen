package tables

var MaskPatterns = [8]func(r, c int) bool{
	func(r, c int) bool { return (r+c)%2 == 0 },
	func(r, c int) bool { return r%2 == 0 },
	func(r, c int) bool { return c%3 == 0 },
	func(r, c int) bool { return (r+c)%3 == 0 },
	func(r, c int) bool { return (r/2+c/3)%2 == 0 },
	func(r, c int) bool { return ((r*c)%2)+((r*c)%3) == 0 },
	func(r, c int) bool { return (((r*c)%2)+((r*c)%3))%2 == 0 },
	func(r, c int) bool { return (((r+c)%2)+((r*c)%3))%2 == 0 },
}
