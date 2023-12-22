package util

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/constraints"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// BEGIN: Input

func FileAsString(day int) string {
	b := CheckErr(os.ReadFile(fmt.Sprintf("%d/input.txt", day)))
	return string(b)
}

func FileAsStringArr(day int, separator string) []string {
	b := CheckErr(os.ReadFile(fmt.Sprintf("%d/input.txt", day)))
	return strings.Split(string(b), separator)
}

func FileAsBytes(day int) []byte {
	return CheckErr(os.ReadFile(fmt.Sprintf("%d/input.txt", day)))
}

func FileAsScanner(day int) *bufio.Scanner {
	f := CheckErr(os.Open(fmt.Sprintf("%d/input.txt", day)))
	return bufio.NewScanner(f)
}

// END: Input

// BEGIN: Run

func Run[T any](part int, fn func() (sum T)) {
	start := time.Now()
	res := fn()
	elapsed := time.Since(start).Seconds()
	fmt.Printf("[%d] Duration: %f seconds | Result: %v\n", part, elapsed, res)
}

func RunBoth[T any, V any](fn func() (p1 T, p2 V)) {
	start := time.Now()
	p1, p2 := fn()
	elapsed := time.Since(start).Seconds()
	fmt.Printf("[1] Duration: %f seconds | Result: %v\n", elapsed, p1)
	fmt.Printf("[2] Duration: %f seconds | Result: %v\n", elapsed, p2)
}

// END: Run

// BEGIN: Slices

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

func StrsToGrid[T constraints.Signed | constraints.Unsigned](strs ...string) (grid [][]T) {
	grid = make([][]T, len(strs))
	for i, s := range strs {
		grid[i] = make([]T, 0)
		for _, r := range s {
			grid[i] = append(grid[i], T(r))
		}
	}
	return grid
}

func StrsToPointIntGrid(strs ...string) (grid map[Point]int) {
	grid = make(map[Point]int)
	for y, line := range strs {
		for x, r := range line {
			grid[Point{X: x, Y: y}] = int(r - '0')
		}
	}
	return grid
}

func SliceSum[T constraints.Ordered](vals ...T) (sum T) {
	for _, v := range vals {
		sum += v
	}
	return sum
}

// END: Slices

// BEGIN: Math

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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// End: Math

// BEGIN: Point

type Point struct{ X, Y int }

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

var bound = Point{}

func SetBounds(b Point) {
	bound = b
}

func (p *Point) InBounds() bool {
	return p.X > 0 && p.X <= bound.X && p.Y > 0 && p.Y <= bound.Y
}

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
func (p *Point) Right() Point {
	switch *p {
	case North:
		return East
	case West:
		return North
	case South:
		return West
	case East:
		return South
	}
	return *p
}

// Left turns the point left 90 degrees
func (p *Point) Left() Point {
	switch *p {
	case North:
		return West
	case West:
		return South
	case South:
		return East
	case East:
		return North
	}
	return *p
}

func (p *Point) MoveLeft(cur Point) Point {
	return p.Move(cur.Left())
}

func (p *Point) MoveRight(cur Point) Point {
	return p.Move(cur.Right())
}

func (p *Point) Move(dir Point) Point {
	return Point{
		X: p.X + dir.X,
		Y: p.Y + dir.Y,
	}
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

func (p *Point) Multiply(n int) Point {
	return Point{
		X: p.X * n,
		Y: p.Y * n,
	}
}

// SwitchPointPoles switches North and South directions.
// Remember to use with care ;)
func SwitchPointPoles() {
	North, South = South, North
}

// END: Point

// BEGIN: Cube

type Cube struct {
	X, Y, Z int
}

func NewCube(x, y, z int) Cube {
	return Cube{
		X: x,
		Y: y,
		Z: z,
	}
}

func CubeFromNums(n [3]int) Cube {
	return Cube{
		X: n[0],
		Y: n[1],
		Z: n[2],
	}
}

// END: Cube

// BEGIN: Cache

// Cache is a simple and safe map[K]V
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

// END: Cache

// BEGIN: String

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

func Palindrome(s string) bool {
	return s == ReverseStr(s)
}

// Atoi is strconv.Atoi but without returning errors.
func Atoi(s string) (n int) {
	return CheckErr(strconv.Atoi(s))
}

func ParseInt(s string, base int) (n int) {
	i, err := strconv.ParseInt(s, base, strconv.IntSize)
	if err != nil {
		panic(err)
	}
	return int(i)
}

// END: String

// BEGIN: PriorityQueue

type priorityQueueItem[T any, P constraints.Ordered] struct {
	value    T
	priority P
}

type PriorityQueue[T any, P constraints.Ordered] struct {
	mut         sync.RWMutex
	items       []*priorityQueueItem[T, P]
	count       uint
	compareFunc func(a, b P) bool
}

func NewPriorityQueue[T any, P constraints.Ordered](compareFunc func(a, b P) bool) *PriorityQueue[T, P] {
	items := make([]*priorityQueueItem[T, P], 1)
	items[0] = nil
	return &PriorityQueue[T, P]{
		mut:         sync.RWMutex{},
		items:       items,
		count:       0,
		compareFunc: compareFunc,
	}
}

func (pq *PriorityQueue[T, P]) Push(value T, priority P) {
	pq.mut.Lock()
	defer pq.mut.Unlock()
	pq.count++
	pq.items = append(pq.items, newPriorityQueueItem[T, P](value, priority))
	pq.swim(pq.size())
}

func (pq *PriorityQueue[T, P]) Pop() (value T, priority P, ok bool) {
	pq.mut.Lock()
	defer pq.mut.Unlock()
	if pq.size() < 1 {
		return value, priority, false
	}
	m := pq.items[1]
	pq.exchange(1, pq.size())
	pq.items = pq.items[:pq.size()]
	pq.count--
	pq.sink(1)

	return m.value, m.priority, true
}

func (pq *PriorityQueue[T, P]) Head() (value T, priority P, ok bool) {
	pq.mut.RLock()
	defer pq.mut.RUnlock()
	if pq.size() < 1 {
		return value, priority, false
	}
	m := pq.items[1]
	return m.value, m.priority, true
}

func (pq *PriorityQueue[T, P]) swim(k uint) {
	for k > 1 && pq.less(k/2, k) {
		pq.exchange(k, k/2)
		k /= 2
	}
}

func (pq *PriorityQueue[T, P]) sink(k uint) {
	for 2*k <= pq.count {
		j := 2 * k
		if j < pq.size() && pq.less(j, j+1) {
			j++
		}
		if !pq.less(k, j) {
			break
		}
		pq.exchange(k, j)
		k = j
	}
}

func (pq *PriorityQueue[T, P]) size() uint {
	return pq.count
}

func (pq *PriorityQueue[T, P]) less(i, j uint) bool {
	return pq.compareFunc(pq.items[j].priority, pq.items[i].priority)
}

func (pq *PriorityQueue[T, P]) exchange(i, j uint) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue[T, P]) IsEmpty() bool {
	pq.mut.RLock()
	defer pq.mut.RUnlock()
	return pq.size() == 0
}

func PriorityMax[T constraints.Ordered](a, b T) bool {
	return a > b
}

func PriorityMin[T constraints.Ordered](a, b T) bool {
	return a < b
}

func NewMaxPriorityQueue[T any, P constraints.Ordered]() *PriorityQueue[T, P] {
	return NewPriorityQueue[T](PriorityMax[P])
}

func NewMinPriorityQueue[T any, P constraints.Ordered]() *PriorityQueue[T, P] {
	return NewPriorityQueue[T](PriorityMin[P])
}

func newPriorityQueueItem[T any, P constraints.Ordered](value T, priority P) *priorityQueueItem[T, P] {
	return &priorityQueueItem[T, P]{value: value, priority: priority}
}

// END: PriorityQueue

// BEGIN: LinkedList

type LinkedList[T any] struct {
	mut   sync.RWMutex
	first *T
	last  *T
	size  int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{mut: sync.RWMutex{}, size: 0}
}

func (l *LinkedList[T]) AddFirst(value T) {
	l.mut.Lock()
	defer l.mut.Unlock()
	if l.size == 0 {
		l.first = &value
		l.last = &value
	} else {
		l.first = &value
	}
	l.size++
}

func (l *LinkedList[T]) AddLast(value T) {
	l.mut.Lock()
	defer l.mut.Unlock()
	if l.size == 0 {
		l.first = &value
		l.last = &value
	} else {
		l.last = &value
	}
	l.size++
}

func (l *LinkedList[T]) RemoveFirst() (value T, ok bool) {
	l.mut.Lock()
	defer l.mut.Unlock()
	if l.size == 0 {
		return value, false
	}
	value = *l.first
	l.first = l.last
	l.size--
	return value, true
}

func (l *LinkedList[T]) RemoveLast() (value T, ok bool) {
	l.mut.Lock()
	defer l.mut.Unlock()
	if l.size == 0 {
		return value, false
	}
	value = *l.last
	l.last = l.first
	l.size--
	return value, true
}

func (l *LinkedList[T]) First() (value T, ok bool) {
	l.mut.RLock()
	defer l.mut.RUnlock()
	if l.size == 0 {
		return value, false
	}
	return *l.first, true
}

func (l *LinkedList[T]) Last() (value T, ok bool) {
	l.mut.RLock()
	defer l.mut.RUnlock()
	if l.size == 0 {
		return value, false
	}
	return *l.last, true
}

func (l *LinkedList[T]) Size() int {
	l.mut.RLock()
	defer l.mut.RUnlock()
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	l.mut.RLock()
	defer l.mut.RUnlock()
	return l.size == 0
}

// END: LinkedList

func CheckErr[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
