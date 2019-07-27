package imperial

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var number = "\\s*[0-9]+\\s*"
var numberAndFraction = number + "\\s*[0-9]+\\s*/\\s*[0-9]+\\s*"
var matchers = []func(string) bool{
	regexp.MustCompile("^" + number + "'$").MatchString,
	regexp.MustCompile("^" + numberAndFraction + "'$").MatchString,
	regexp.MustCompile("^" + number + "\"$").MatchString,
	regexp.MustCompile("^" + numberAndFraction + "\"$").MatchString,
	regexp.MustCompile("^" + number + "' +" + number + "\"$").MatchString,
	regexp.MustCompile("^" + number + "' +" + numberAndFraction + "\"$").MatchString,
	regexp.MustCompile("^" + numberAndFraction + "' +" + number + "\"$").MatchString,
	regexp.MustCompile("^" + numberAndFraction + "' +" + numberAndFraction + "\"$").MatchString,
}

// Parse takes an input string and converts it to metric in the following formats
// 1 1/2' 2 1/2" = 1.5 feet, 2.5 inches
// 1' 2 1/2" = 1 foot, 2.5 inches
// 2 1/2" = 2.5 inches
// 1/2" = 0.5 inches
// 1' = 1 foot
// leaving out the ' or " results in an error
func Parse(input string) (float64, error) {
	if !isWellFormed(input) {
		return 0, fmt.Errorf("could not parse %s, must conform to pattern 0 0/0' 0 0/0\"", input)
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
	splits := strings.Split(input, " ")
	whole, err := strconv.ParseFloat(splits[0], 64)
	if err != nil {
		return 0, err
	}
	fraction := float64(0)
	if len(splits) == 2 {
		fraction, err = convertFraction(splits[1])
		if err != nil {
			return 0, err
		}
	}

	return whole*multiplier + fraction*multiplier, nil
}

func convertFraction(input string) (float64, error) {
	if !strings.Contains(input, "/") {
		return 0, fmt.Errorf("could not parse %s, expecting / in fractional string", input)
	}
	splits := strings.Split(input, "/")
	numerator, err := strconv.ParseFloat(splits[0], 64)
	if err != nil {
		return 0, err
	}
	denominator, err := strconv.ParseFloat(splits[1], 64)
	if err != nil {
		return 0, err
	}
	return numerator / denominator, nil
}

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
