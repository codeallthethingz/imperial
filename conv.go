package imperial

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var number = "\\s*[0-9]+\\s*"
var fraction = "\\s*[0-9]+\\s*/\\s*[0-9]+\\s*"
var numberAndFraction = number + fraction
var matchers = []func(string) bool{
	regexp.MustCompile(`^` + fraction + `"$`).MatchString,
	regexp.MustCompile(`^` + fraction + `'$`).MatchString,
	regexp.MustCompile(`^` + number + `'$`).MatchString,
	regexp.MustCompile(`^` + numberAndFraction + `'$`).MatchString,
	regexp.MustCompile(`^` + number + `"$`).MatchString,
	regexp.MustCompile(`^` + numberAndFraction + `"$`).MatchString,
	regexp.MustCompile(`^` + number + `' +` + number + `"$`).MatchString,
	regexp.MustCompile(`^` + number + `' +` + numberAndFraction + `"$`).MatchString,
	regexp.MustCompile(`^` + numberAndFraction + `' +` + number + `"$`).MatchString,
	regexp.MustCompile(`^` + numberAndFraction + `' +` + numberAndFraction + `"$`).MatchString,
}

// Parse takes an input string and converts it to metric in the following formats
// 1 1/2' 2 1/2" = 1.5 feet, 2.5 inches
// 1' 2 1/2" = 1 foot, 2.5 inches
// 2 1/2" = 2.5 inches
// 1/2" = 0.5 inches
// 1' = 1 foot
// leaving out the ' or " results in an error
func Parse(input string) (string, error) {
	if !isWellFormed(input) {
		return "", fmt.Errorf("could not parse %s, must conform to pattern 0 0/0' 0 0/0\"", input)
	}

	result := float64(0)
	inchesIndex := -1
	if strings.Contains(input, "'") {
		feetString := strings.Split(input, "'")[0]
		feet, _ := convert(feetString, 0.3048)
		result += feet
		inchesIndex = len(feetString)
	}
	if strings.Contains(input, "\"") {
		inchesString := strings.Split(input[inchesIndex+1:], "\"")[0]
		inches, _ := convert(inchesString, 0.0254)
		result += inches
	}
	return round(result, 0.0001), nil
}

func isWellFormed(input string) bool {
	for _, wellFormed := range matchers {
		if wellFormed(input) {
			return true
		}
	}
	return false
}

func convert(input string, multiplier float64) (float64, error) {
	input = strings.Trim(input, " ")
	splits := strings.SplitN(input, " ", 2)
	whole := float64(0)
	fraction := float64(0)
	if len(splits) == 2 {
		whole, _ = strconv.ParseFloat(splits[0], 64)
		fraction = convertFraction(splits[1])
	} else {
		if strings.Contains(splits[0], "/") {
			fraction = convertFraction(splits[0])
		} else {
			whole, _ = strconv.ParseFloat(splits[0], 64)
		}
	}

	return whole*multiplier + fraction*multiplier, nil
}

func convertFraction(input string) float64 {
	splits := strings.Split(input, "/")
	numerator, _ := strconv.ParseFloat(splits[0], 64)
	denominator, _ := strconv.ParseFloat(splits[1], 64)
	return numerator / denominator
}

func round(x, unit float64) string {
	return fmt.Sprintf("%.4f", math.Round(x/unit)*unit)
}
