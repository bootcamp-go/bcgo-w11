Denormalized Table of Names

| ID | Name       |
|----|------------|
| 1  | John Smith |     -> group 1
| 2  | John Smith |     -> .
| 3  | John Smith |     -> .
| 4  | Emily Doe  |     -> group 2
| 5  | Emily Doe  |     -> .
| 6  | Alex Brown |     -> group 3

Grouped Table of Names
```sql
SELECT Id, Name FROM Names GROUP BY Name;
```

| ID | Name       |
|----|------------|
| 1  | John Smith |          -> group 1
| 4  | Emily Doe  |          -> group 2
| 6  | Alex Brown |          -> group 3

Grouped Table of Names with Count
```sql
SELECT Id, Name, COUNT(id) FROM Names GROUP BY Name;
```

| ID | Name       |
|----|------------|                         count
| 1  | John Smith |     -> group 1      |   3
| 2  | John Smith |     -> .
| 3  | John Smith |     -> .

| 4  | Emily Doe  |     -> group 2      |   2
| 5  | Emily Doe  |     -> .

| 6  | Alex Brown |     -> group 3      |   1

final output:
| ID | Name       | count
|----|------------|-------
| 1  | John Smith | 3
| 4  | Emily Doe  | 2
| 6  | Alex Brown | 1