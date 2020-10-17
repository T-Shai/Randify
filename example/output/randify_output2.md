### TEST N°3339

    This is a randify input file example :
    txt file with randify markup that will generate a mardown file
---

Comment n59 :

in python i am used to write :

```py
def func9 :
    print("this is an integer -54")
    print("this is a float -62.304966")
    return 10.201262053772082
```


### command used :

launched from the project directory :

`./randify.exe -i example/test.md -id r -ext md -o example/output/randify_output -n 5`

### Explanation :

> -i example/test.md

Input file's path (this file)

> -id r
    
The identifier i am using in this text is r

> -ext md

I want the output to be a md file too

> -o example/output/randify_output

Output file path without the extension not that if the directory doesn't exist it ll panic

> -n 5

I want 5 different files
