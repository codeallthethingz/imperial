# Imperial conversion [![CircleCI](https://circleci.com/gh/codeallthethingz/imperial.svg?style=svg)](https://circleci.com/gh/codeallthethingz/imperial)

Library to convert strings of imperial measurements specified in feet and inches into metric meter lengths.

## Install

`go get github.com/codeallthethingz/imperial`

## Usage

```go
convertMe := `1' 2 1/2\"`
parsed, err := imperial.Parse(convertMe)
if err != nil {
  panic(err)
}
fmt.Println(parsed)
```

Outputs

```bash
0.3683
```

## Example formats

- `1'` = 1 foot
- `1/2"` = 0.5 inches
- `2 1/2"` = 2.5 inches
- `1' 2 1/2"` = 1 foot 2.5 inches
- `1 1/2' 2 1/2"` = 1.5 feet 2.5 inches

Leaving out the `'` or `"` results in an error.
