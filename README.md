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

Also check example directory for an example with markdown file

### Usage Help
use ``.\executable.exe -h``

#### Avaiable command line flags :

   **-i** string  **NOT OPTIONAL**
   
        Path to the input filename (must be defined)
        EX : ./executable.exe -i path/to/input.txt
        
        Any ascii readable file (tested with .txt and .md)
        
  **-ext** string
  
        output file's extension without dot (default "randify")
        EX : ./executable.exe -ext myOutputFileName
        
        you will get myOutputFileName0, myOutputFileName1, etc... for each file 
        
  **-id** string
  
        Identifier to put in brackets before expressions (default "RAND")
        EX : ./executable.exe -id myOwnIdentifier
        
        it will be the the identifier that you ll have to put in bracket before expressions
        EX  in your input file : "this is a random integer between -10 and 10 : [myOwnIdentifier int -10 10]"
        
  **-n** uint
  
        number of file to be created (all with randomly different values) (default 1)
        EX : ./executable.exe -n 5
        
        n must be > 0 to output a file
        
  **-o** string
  
        output file's name without extension (default "output")
        EX : ./executable.exe -o path/to/myFileWithoutextension
        
        only the filename without extension if you want to change the extension use -ext
        
  **-seed** int
  
        seed for random values DEFAULT: current time (default time.Now().UnixNano())
        EX : ./executable.exe -seed 544
        
        set the seed if you want 

## known bug
~~None 😎😎 yet...~~
output same file sometimes *(problem with coroutines and math/rand ?)*
