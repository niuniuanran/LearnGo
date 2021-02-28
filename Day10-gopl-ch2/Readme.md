# Comparability and Assignability
## [Value of a named type] comparable with [Value of the same type] 
``` 
    type Cat uint8
	var c1 Cat = 1
	var c2 Cat = 2
	fmt.Println(c1 < c2) // true
```

## [Value of an unnamed type] comparable with [Value of the same underlying type]

### value literal is of unnamed type
``` 
    type Cat uint8
	var c Cat = 1
	fmt.Println(1 == c) // true
```

### const is of unnamed type
``` 
	const b = 1
    type Cat int
	var c Cat = 2
	fmt.Println(b < c)
	fmt.Printf("%T %v\n", b, b) // int 1
```

## Two values of different types cannot be compared directly

### Value declaration gives a type
``` 
    type Cat int
	b := 1
	var c Cat = 2
	fmt.Println(b < c) // invalid operation: b < c (mismatched types int and Cat)
	fmt.Printf("%T %v\n", b, b) // int 1
```

### I can do a conversion

``` 
	type Cat uint8
	var c Cat = 2
	b := 1
	fmt.Println(Cat(b) < c)  // true
	fmt.Printf("%T %v\n", b, b) // int 1
```

#### Follow the assignability rule for type conversion

``` 
	b := 1
	type Cat uint8
	var c Cat = 2
	cat := Cat(b)
	fmt.Println(cat < c) // true
	fmt.Printf("%T %v\n", cat, cat) //main.Cat 1
```

``` 
	b := 1
	type Cat uint8
	var c Cat = 2
	b = Cat(b) // cannot use Cat(b) (type Cat) as type int in assignment
	fmt.Println(b < c) // invalid operation: b < c (mismatched types int and Cat)
```