# Assignment for Day 3

## Rectangle counter
- Goal: Find and count number of rectangles in a 2D array.
- Inputs: An array filled with 0s and 1s.
- Outputs: Number of rectangles filled with 1s

## Usage
```console
func main() {
    arr := [][]int {
        {1, 0, 0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0, 0, 0},
        {1, 0, 0, 1, 1, 1, 0},
        {0, 1, 0, 1, 1, 1, 0},
        {0, 1, 0, 0, 0, 0, 0},
        {0, 1, 0, 1, 1, 0, 0},
        {0, 0, 0, 1, 1, 0, 0},
        {0, 0, 0, 0, 0, 0, 1},
    }
    
    count := countRectangles(arr)
    fmt.Printf("%v", count) // 6
}
```
