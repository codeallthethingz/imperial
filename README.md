# Imperial conversion

Library to convert strings of imperial measurements in feet and inches into metric meter lengths.

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
