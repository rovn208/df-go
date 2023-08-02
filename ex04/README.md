# Assignment for Day 4

## Different integers in a String counter
- Goal: Count the number of different integers in a String.
- Inputs: A string word consists of digits and lowercase English
  letters, 2 integers are considered different if their decimal
  representation without any leading zeros are different.
- Outputs: Number of different integers

## Usage
```console
func main() {
    word := "a123bc34d8ef34"
    simpleCounter := counter.NewCounter(word)
    fmt.Println(simpleCounter.GetDifferentIntegers())
}
```
