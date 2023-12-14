package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

func FileAsString(day int) string {
	b := CheckErr(os.ReadFile(fmt.Sprintf("input/%d", day)))
	return string(b)
}

func FileAsStringArr(day int, separator string) []string {
	b := CheckErr(os.ReadFile(fmt.Sprintf("input/%d", day)))
	return strings.Split(string(b), separator)
}

func FileAsBytes(day int) []byte {
	b := CheckErr(os.ReadFile(fmt.Sprintf("input/%d", day)))
	return b
}

func FileAsScanner(day int) *bufio.Scanner {
	f := CheckErr(os.Open(fmt.Sprintf("input/%d", day)))
	return bufio.NewScanner(f)
}

func Run[T any](part int, fn func() T) {
	start := time.Now()
	res := fn()
	elapsed := time.Since(start).Seconds()
	fmt.Printf("[%d] Duration: %f seconds | Result: %v\n", part, elapsed, res)
}

func RunBoth[T any, V any](fn func() (T, V)) {
	start := time.Now()
	p1, p2 := fn()
	elapsed := time.Since(start).Seconds()
	fmt.Printf("[1] Duration: %f seconds | Result: %v\n", elapsed, p1)
	fmt.Printf("[2] Duration: %f seconds | Result: %v\n", elapsed, p2)
}

func StrsToIntSlice(nums ...string) (ints []int) {
	for _, num := range nums {
		ints = append(ints, Atoi(num))
	}
	return
}

func ValuesToNum[T string | rune](strs ...T) (n int) {
	var s string
	for _, num := range strs {
		s = fmt.Sprintf("%s%s", s, string(num))
	}
	return Atoi(s)
}

// Atoi is strconv.Atoi but without returning errors.
func Atoi(s string) (n int) {
	return CheckErr(strconv.Atoi(s))
}

// LCM (lcm or least common multiple) returns the least common multiple of the provided integers.
// Uses GCD in each iteration.
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// https://en.wikipedia.org/wiki/Least_common_multiple
func LCM(numbers ...int) (result int) {
	result = numbers[0]
	for _, num := range numbers[1:] {
		result = (result * num) / GCD(result, num)
	}
	return result
}

// GCD (greatest common divisor) Euclidean algorithm.
// Returns the greatest divisor that is common for a and b.
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// https://en.wikipedia.org/wiki/Euclidean_algorithm
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// North +Y direction
var North = Point{0, 1}

// South -Y direction
var South = Point{0, -1}

// East +X direction
var East = Point{1, 0}

// West -X direction
var West = Point{-1, 0}

// NorthWest -X+Y direction
var NorthWest = Point{-1, 1}

// NorthEast +X+Y direction
var NorthEast = Point{1, 1}

// SouthWest -X-Y direction
var SouthWest = Point{-1, -1}

// SouthEast +X-Y direction
var SouthEast = Point{1, -1}

var AdjacentWithDiagonals = []Point{North, South, East, West, NorthWest, NorthEast, SouthWest, SouthEast}

var Adjacent = []Point{North, South, East, West}

type Point struct{ X, Y int }

// Adjacent returns true if the coordinate is adjacent to the other coordinate.
// The offset allows the adjacent coordinate to be offset on the x-axis with the given size.
func (p *Point) Adjacent(other Point, adj []Point, offset ...int) bool {
	if len(offset) != 0 {
		return p.adjacentWithOffset(offset[0], other, adj)
	}
	for _, coord := range adj {
		if other.X == coord.X+p.X && other.Y == coord.Y+p.Y {
			return true
		}
	}
	return false
}

func (p *Point) adjacentWithOffset(offset int, other Point, adj []Point) bool {
	var digits int
	for offset != 0 {
		offset /= 10
		digits++
	}
	for _, coord := range adj {
		for i := 0; i < digits; i++ {
			if other.X == coord.X+p.X+i && other.Y == coord.Y+p.Y {
				return true
			}
		}
	}
	return false
}

// Add returns the sum of two points
func (p *Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

// Scale multiplies the point by a scalar
func (p *Point) Scale(factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// Right turn point right 90 degrees
func (p *Point) Right() {
	p.X -= p.X
}

// Left turns the point left 90 degrees
func (p *Point) Left() {
	p.Y -= p.Y
}

// Manhattan returns the manhattan magnitude |x|+|y|
func (p *Point) Manhattan() int {
	return Abs(p.X) + Abs(p.Y)
}

func (p *Point) ManhattanDistance(other Point) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

func (p *Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CheckErr[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

type Cache[K comparable, V any] struct {
	mut   sync.RWMutex
	cache map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		mut:   sync.RWMutex{},
		cache: make(map[K]V),
	}
}

func (c *Cache[K, V]) Get(key K) (val V, ok bool) {
	c.mut.RLock()
	defer c.mut.RUnlock()
	val, ok = c.cache[key]
	return val, ok
}

func (c *Cache[K, V]) Set(key K, val V) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.cache[key] = val
}

func ReverseStr(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func FlipGrid[T any](a [][]T, times ...int) (b [][]T) {
	rows, cols := len(a), len(a[0])
	b = make([][]T, cols)
	for i := range b {
		b[i] = make([]T, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			b[j][rows-i-1] = a[i][j]
		}
	}
	if times == nil {
		return b
	}
	t := times[0]
	if t > 1 {
		return FlipGrid(b, t-1)
	}
	return b
}

func SliceCount[T comparable](s []T, sub T) (sum int) {
	for _, v := range s {
		if v == sub {
			sum++
		}
	}
	return sum
}
