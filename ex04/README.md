# Assignment for Day 4

Goal: Count the number of different integers in a String.
Outputs: Number of different integers
Given that a string word consists of digits and lowercase English
letters, 2 integers are considered different if their decimal
representation without any leading zeros are different.
E.g: “a123bc34d8ef34” => 3 (123, 34, 8)
"A1b01c001" => 1 (1)

## Rectangle counter
- Goal: Count the number of different integers in a String.
- Inputs: A string word consists of digits and lowercase English
  letters, 2 integers are considered different if their decimal
  representation without any leading zeros are different.
- Outputs: Number of different integers

## Usage
```console
func main() {
    word := "a123bc34d8ef34"
    counter := counter2.NewCounter(word)

    fmt.Println(counter.GetDifferentIntegers())
}
```
