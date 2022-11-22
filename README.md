# golang 
### Collections of exercises to learn and become fluent in Go programming language taken from the WASA (Web and Software architecture) courses

## Exercises part 1

### Help the train company to arrive on time!

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

### Calculate your GPA

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