# Randify
#### A simple markup language to randomize values in text files 
---

### Example :
in `input.txt`:
```
Math exercice n[RAND uint 100]:

Calculate x for :

    [RAND uint 20] * x  - [RAND float 10] = [RAND int 100 999]
    [RAND uint 20] * x  - [RAND float 10] = [RAND int 100 999]
    [RAND uint 20] * x  - [RAND float 10] = [RAND int 100 999]
    [RAND uint 20] * x  - [RAND float 10] = [RAND int 100 999]
    [RAND uint 20] * x  - [RAND float 10] = [RAND int 100 999]
    [RAND uint 20] * x  - [RAND float 10] = [RAND int 100 999]

Hi, i am [RAND uint 20 80] years old.
I am pretty rich, i have exactly [RAND float 1000 10000]€ in my pocket.
Well not exactly... my back account shows [RAND int -1000 0]€

for my birthday, i ll get [RAND int 10] gifts !
```

command line used :
```sh
./executable.exe -i input.txt
```
in `output0.randify`:
```
Math exercice n11:

Calculate x for :

    11 * x  - 5.1754527 = 993
    12 * x  - 3.3593845 = 303
    17 * x  - 9.753878 = 653
    1 * x  - 8.178357 = 344
    16 * x  - 6.438222 = 979
    6 * x  - 0.5194068 = 243

Hi, i am 55 years old.
I am pretty rich, i have exactly 9717.011€ in my pocket.
Well not exactly... my back account shows -943€

for my birthday, i ll get 5 gifts !
```

### Usage
use ``.\executable.exe -h``

## known bug
None 😎😎