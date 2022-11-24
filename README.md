# golang 
### Collections of exercises to learn and become fluent in Go programming language taken from the WASA (Web and Software architecture) courses

## Exercises part 1

### [Help the train company to arrive on time!](https://github.com/albertoscala/golang/blob/main/train_company.go)

Write a program that reads a JSON structure from a file representing a train track and writes an updated
JSON file with calculated values.

You should calculate the following:
- The total distance
- The total travel time in minutes
- The arrival time

Input file example:
```json
{
    "line": "FL7",
    "stations": [
        {
            "name": "Latina",
            "distance": 0,
            "ETA-from-previous-station": 0,
            "platform": "1T"
        },
        {
            "name": "Cisterna di Latina",
            "distance": 18,
            "ETA-from-previous-station": 7,
            "platform": "3"
        },
        {
            "name": "Campoleone",
            "distance": 26,
            "ETA-from-previous-station": 10,
            "platform": "4"
        },
        {
            "name": "Pomezia - S. Palomba",
            "distance": 19,
            "ETA-from-previous-station": 7,
            "platform": "2"
        },
        {
            "name": "Torricola",
            "distance": 21,
            "ETA-from-previous-station": 8,
            "platform": "2"
        },
        {
            "name": "Roma Termini",
            "distance": 32,
            "ETA-from-previous-station": 12,
            "platform": "16"
        }
    ],
    "departingTime": "2022-11-02T09:00:00Z"
}
```

The distance is expressed in km, and the ETA-from-previous-station is the number of
minutes to reach the station from the previous one.

> Hint: use time.Time for departingTime, so that Go will parse the timestamp for you!

An example of the output is:
```json
{
    "totalDistance": 116,
    "travelTime": 44,
    "arrivalTime": "2022-11-02T09:44:00Z"
}
```

> Hint: use struct for reading/writing JSON. See encoding/json for details.

### [Calculate your GPA](https://github.com/albertoscala/golang/blob/main/gpa_calculator.go)

Write a program to calculate the Grade point average (GPA) of a given set of scores. The GPA for
Computer Science courses in Sapienza is calculated as follows:
- The average algorithm is the “Weighted arithmetic mean”. For each grade to sum, the weight is
the ratio between the CFU/ECTS for that grade and the total number of CFU/ECTS.
- The program should ask the user for grades. The number of grades is not defined. You can
assume that no one will take more than 40 exams, but you should not require the user to input
40 values if the user did fewer exams.
- Grades in Italy are a natural number within the range [0, 30], but you should ignore grades
from zero to seventeen as they don’t count in the GPA.
- CFU/ECTS for each grade are a natural number in the range [1, 40].
- A score over 30 is referred to as “lode” or “cum laude” (Latin). For the GPA, scores over 30 are
counted as 30.
- You can use some values for special meaning – e.g., zero score means “done inserting scores”

Examples (grade/CFU):
- Alice: 30/6 30/9 25/6 22/12 26/6 28/6 = 26.4
- Bob: 18/6 20/6 30/12 30/9 24/6 18/6 29/12 28/9 = 25.91

> Hint: see bufio.NewReader() and its ReadString. You need to pass them the reference to
> the standard input: in Go, operating system-related functions/variables are inside the os package.

> Hint: strconv package might be useful.

### [Word count](https://github.com/albertoscala/golang/blob/main/count_words.go)

Write a program that opens and read a file (the name can be hard-coded inside the source code) and
counts all occurrences for every different word. The output should be the list of words and the number
of times the word is present in the text (one per line).

To test with a long text, you can download The Project Gutenberg eBook of Alice’s Adventures in
Wonderland, by Lewis Carroll (in this case I have used The Arabian Nights, you can find it in the txt folder).

### [Web-train](https://github.com/albertoscala/golang/blob/main/train_company_web.go)

In this exercise, you need to create a web server that implements the exercise “Help the train company
to arrive in time!” via HTTP. JSON should be received via an HTTP POST request (in the body), and the
response JSON should be sent in the body of the HTTP response.

### [Web-GPA](https://github.com/albertoscala/golang/blob/main/gpa_calculator_web.go)

This exercise is identical to the “Calculate your GPA”, except that input data must be sent as JSON via
an HTTP POST request to the server, and the response should be provided as plain text.

### [Web-Count](https://github.com/albertoscala/golang/blob/main/count_words_web.go)

This exercise is identical to the “Word count”, except that input data must be sent as plain text via an
HTTP POST request to the server, and the response should be provided as plain text.

## Exercises part 2

> You should be able to do these exercises after the second Go theory lecture.

Welcome commander! We have a new mission for you. The Republic is in danger! A rogue general
wants to impose his dictatorship. You will be sent back in time to block the war and maintain the
integrity of the Republic and its Senate! Here is the profile of your primary opponent:

> [Gaius Julius Caesar](https://en.wikipedia.org/wiki/Julius_Caesar) was an important Roman general and dictator. He led an army that conquered
> many territories and ultimately started a civil war inside the Roman Republic, transforming it into
> the Roman Empire.

Your mission is to block Caesar before he [crosses the Rubicon](https://en.wikipedia.org/wiki/Crossing_the_Rubicon) to stop the imminent civil war!

You will be accompanied by your trusted *Go Compiler*

### [Caesar’s cipher](https://github.com/albertoscala/golang/blob/main/ceaser_cypher.go)

*[Gallia Cisalpina](https://en.wikipedia.org/wiki/Crossing_the_Rubicon)*, January 5, 49 BC. You landed in what will be the north part of Italy in 2000 years, while
now it’s a Roman province managed by Celts. Based on your history courses in High School, you know
that Caesar must be here, waiting to cross the Rubicon. In the past days, Caesar prepared an army,
and it’s coordinating his attack with other generals. You don’t have weapons or people, so you can’t
compete. You must use another strategy.

To avoid being intercepted, Caesar used *cryptography* (yes, nearly 2070 years ago!) to protect messages
to commanders and other generals. *Caesar’s cipher* is a simple substitution cipher - see the Wikipedia
page for the algorithm: https://en.wikipedia.org/wiki/Caesar_cipher.

Thanks to your trusted companion (the Go Compiler), you can create false messages to send to other
generals. You can put generals against each other!

> In this exercise, you are asked to implement a structure that encrypts messages and prints them
> to the screen. The `struct` should conform to `io.Writer` so that you can use this structure to
> encrypt anything!
>
> Note: Caesar’s cipher is defined only for alphabet letters, so everything else can be left as it is. It
> means that if you print non-letters (numbers, symbols, etc.), they should be printed to the screen
> unchanged. Lowercase letters should be replaced with uppercase letters.

In the end, you should be able to use it as:

```go
caesarCipher := MyCaesarCipher{Key: 23}
fmt.Fprintln(&caesarCipher, "Marcus Tullius is stealing your wife!")
```

(the output depends on the key).
> Note: do NOT use external libraries.

### [Breaking the cipher!](https://github.com/albertoscala/golang/blob/main/crack_cypher.go)

*Gallia Cisalpina*, January 9, 49 BC. In less than 48 hours, the Republic will be doomed! Your tentative of subverting the attack failed because you need to know which key to use, and most importantly, you don’t know generals at all!

You need to intercept them first.

While getting the message itself is straightforward, they’re encrypted. However, this is not a problem for the Go Compiler!

You can break Caesar’s cipher in different ways:
- *Brute-force*: slow but easier to implement;
- *Frequency analysis*: fast but slightly more complex.

Unfortunately, you don’t have a dictionary with word frequency, so you must use *brute-force*. So, given that your Go Compiler can use multiple CPUs, you decide to use Goroutines!

> The goal of this exercise is to implement Caesar’s cipher cracker using goroutines for generating
> the potential plain text (the original text) for each possible key, given the cipher text (the encrypted
> text).
> 
> Output of all combinations is ok - you can assume that a human will look at all messages to spot
> the right one.

> Note: do NOT use external libraries.

> Trivia: “Brute-force” term comes from “Brute”, which in turn comes from “Brutus”. Caesar was
> killed by a senator named “Brutus” in 44 BC.

### Back to the future!

The mission is complete! Generals are fighting each other - you saved the Republic! Now, it’s about time to get back to the future!

Your supervisor gave you the coordinates for the temporal wormhole, but they are encoded, and you need to derive them. To do that, they gave you the derivation function generator.

However, to avoid misuse, your supervisor requires you to provide (to the generator) a function that adds the year of your birth to a value and returns the resulting value.

The result of the generator will be a function that you can call with a time (e.g., time.Now()) to get temporal coordinates.

The generator is:

```go
func temporalCoordinatesDerivation(yourBirthYear func(int64) int64) func(time.Time) (int64, int64) {
    x := 16342 * yourBirthYear(52634)
    y := 88345 * yourBirthYear(35237)
    return func(t time.Time) (int64, int64) {
        return x + t.Unix(), y + t.Unix()
    }
}
```

> In this exercise, you are required to:
> - define a function that adds your birth date to a parameter, and return the results >>(function
> signature: func(int64) int64);
> - use the resulting function from the generator to obtain the coordinates (function >   signature
> func(time.Time) (int64, int64))