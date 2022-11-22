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