package bluebook

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// Stuff for Exercise01

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`

	pageTop02 = `<!DOCTYPE HTML><html><head>
	<style>.error{color:#FF0000;}</style></head>
	<title>Quadratic Equation Solver</title><body>
	<h3>Quadratic Equation Solver</h3><p>Solves equations of the form
	a<i>x</i>² + b<i>x</i> + c</p>`
	form02 = `<form action="/" method="POST">
	<input type="text" name="a" size="1"><label for="a"><i>x</i>²</label> +
	<input type="text" name="b" size="1"><label for="b"><i>x</i></label> +
	<input type="text" name="c" size="1"><label for="c"> →</label>
	<input type="submit" name="calculate" value="Calculate">
	</form>`
	pageBottom02 = "</body></html>"
	error02      = `<p class="error">%s</p>`
	solution     = "<p>%s → %s</p>"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
	sigma   float64 // Standard deviation
	mode    []float64
}

func getStatistics(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	stats.sigma = sigma(stats.numbers, stats.mean, len(numbers))
	stats.mode = mode(stats.numbers)
	return stats
}

func sum(numbers []float64) (theSum float64) {
	for _, n := range numbers {
		theSum += n
	}
	return theSum
}

func median(numbers []float64) (result float64) {
	middle := len(numbers) / 2
	result = numbers[middle]

	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}

	return result

}

func sigma(numbers []float64, mean float64, samples int) (stdDeviation float64) {
	partialSum := 0.0
	for _, number := range numbers {
		partialSum += math.Pow(number-mean, 2)
	}

	stdDeviation = math.Sqrt(partialSum / float64((len(numbers) - 1)))
	return stdDeviation
}

func mode(numbers []float64) (modes []float64) {
	frequencies := make(map[float64]int, len(numbers))
	highestFreq := 0

	for _, number := range numbers {
		frequencies[number]++

		if frequencies[number] > highestFreq {
			highestFreq = frequencies[number]
		}
	}

	for number, frequency := range frequencies {
		if frequency == highestFreq {
			modes = append(modes, number)
		}
	}

	if len(modes) == len(numbers) {
		return []float64{}
	}

	sort.Float64s(modes)
	return modes
}

func startServer01() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprint(rw, pageTop, form)
	err := request.ParseForm()

	if err != nil {
		fmt.Fprintf(rw, anError, err)

	} else {
		if numbers, message, ok := processRequest(request); ok {

			stats := getStatistics(numbers)
			fmt.Fprint(rw, formatStats(stats))

		} else if message != "" {
			fmt.Fprintf(rw, anError, message)
		}
	}
	fmt.Fprint(rw, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)

		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false

			} else {
				numbers = append(numbers, x)
			}
		}
	}

	if len(numbers) == 0 {
		return numbers, "", false
	}

	return numbers, "", true
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
<tr><td>Sigma</td><td>%f</td></tr>
<tr><td>Mode</td><td>%f</td></tr>
</table>`,
		stats.numbers,
		len(
			stats.numbers),
		stats.mean,
		stats.median,
		stats.sigma,
		stats.mode,
	)
}

func Exercise01C02(args ...interface{}) error {
	fmt.Println("Executing exercise 01 blue book, chapter 01")
	startServer01()
	return nil
}

// Stuff for exercise 02

func Exercise02C02(args ...interface{}) error {
	fmt.Println("Executing exercise 02 blue book, chapter 01")
	startServer02()
	return nil
}

func startServer02() {
	http.HandleFunc("/", homePage02)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage02(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, pageTop02, form02)
	err := request.ParseForm()

	if err != nil {
		fmt.Fprintf(writer, error02, err)

	} else {
		// There are no numbers (first time I visit page)
		// There are valid numbers
		// There are not valid numbers
	}

	fmt.Fprint(writer, pageBottom02)
}
