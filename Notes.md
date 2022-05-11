## Notes

- Just made an empty `.db` file to "initialize the db"
- Entered sqlite shell with `sqlite3 air_quality.db`
- Built new table with

```
CREATE TABLE Observations(
  "time" INTEGER PRIMARY KEY,
  "co2" INTEGER,
  "temp" DOUBLE,
  "humidity" DOUBLE
);
```

```sqlite3
.mode csv
.import air_quality_int_time.csv Observations
.mode column
```

Check on status with

```
SELECT * FROM Observations LIMIT 10;
```

- Used the following R script to convert the original date format into unix timestamps

```r
library(readr)
library(dplyr)
library(lubridate)

raw_data <- read_csv(
  "~/dev/go_experiments/linereader/air_quality.csv",
  col_names = c("time", "co2", "temp", "humidity"),
  col_types = cols(
    time = col_datetime(format = "%m/%d/%y %H:%M:%S"),
    co2 = col_integer(),
    temp = col_number(),
    humidity = col_number()
  )
) %>% mutate( time = as.integer(time) ) %>%
  tail(5000)


write_csv(
  raw_data,
  "~/dev/go_experiments/linereader/air_quality_int_time.csv",
  col_names = FALSE
)

```

- run go program with `go run .`
