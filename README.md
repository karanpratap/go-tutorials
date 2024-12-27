# Notes for my Go tutorials

### Go slices 
Much like vectors, automatically resizing arrays.\
Omitting the size of the slice in brackets when declaring it, tells go that the size of the array underlying the slice can be dynamically changed. e.g.: `[]string`.\
Like python `a[low:high]` gives you a slice, where `a` is an array (including `low`; excluding `high`).\
> __A slice does not store any data, it just describes a section of an underlying array.__\
> __Changing the elements of a slice modifies the corresponding elements of its underlying array.__\
> __Other slices that share the same underlying array will see those changes.__\

Keeping the above points in mind, a slice literal like `[]bool{true, false, true}` creates an underlying array, builds a slice and then references it.\
Out of bounds runtime error is caused if one tries to create a slice with one of the limits exceeding the underlying array's capacity.\
The zero value of a slice is nil. The `make` function allocates a zeroed array and returns a slice that refers to that array:
```go
a := make([]int, 4, 5)      // slice with length 4 and capacity 5
b := make([]int, 5)         // slice with length = capacity = 5

cap(a)                      // get capacity of a
len(b)                      // get length of b

// append element(s) to a slice, resizing as required 
// (works on nil slices) - if the backing array is too
// small, a new array is allocated, and the returned slice
// of append points to the newly allocated array
a = append(a, 0, 1, ...)
```

### Loops

Go has only 1 looping construct: `for`\
Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.\
`for` is Go's `while`: after removing the init and test expressions from the loop, the semicolons can be omitted. Thus, this is valid syntax:
```go
sum := 0
for sum <= 1000 {
    sum += 1
}

// Infinite loop
for {
    // always enters
}
```
Omit all the expressions and the semicolons and we have an infinite loop

### Conditional

`if` is similar to `for`: It can also have an initialization expression, and the variable is limited to the scope of the if statement. This variable is also available in any of the corresponding `else` blocks.
```go
if v := math.Pow(x, n); v < lim {
    return v
} else {
    return v + 10
}
```

### Switch

Apparently, switch is not a jump table in Go as it is in C.\
The cases can be conditionals, and other expressions, which are just compared to the expression passed in to the `switch` statement.\
A switch with no initial expression is equivalent to saying `switch true`, and can be used to emulate long if-then-else chains.
```go
switch {
    case conditional_expression1:   // if
    case conditional_expression2:   // else if
    ...
    case conditional_expressionN:   // else if
    default:                        // else
}
```

### Range

The `range` form of the `for` loop iterates over a slice or map.\
When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

### Maps

The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added.\
The make function returns a map of the given type, initialized and ready for use.\
Map literals are like struct literals but the keys are required, unlike struct literals, where keys can be omitted.
```go
m[key] = elem
elem = m[key]
delete(m, key)

// ok is true if key is found in m, else false
elem, ok = m[key]
elem, ok := m[key] // if elem and ok haven't been declared before
```

### Defer

A defer statement defers the execution of a function until the surrounding function returns.\
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
```go
func another_fn() string {
    fmt.Println("Ahhh...side effect")
    return "world"
}

func main() {
	defer fmt.Println(another_fn())

	fmt.Print("hello ")
}
```
The code above will result in the following output, as the call to another_fn() within the defer statement is evaluated first:
```
Ahhh...side effect
hello world
```
Deferred statements are stacked, and executed in LIFO order.

### Pointers

`*T` is a pointer of type `T`. For example, a declaration `x *int` means x is a pointer of type int.\
The `&` operator generates a pointer to its operand, much like C.\
The `*` prefix operator denotes the pointer's underlying value. Again, much like C!\
BUT, unlike C, Go has no pointer arithmetic :(

### Structs

```go
type <struct-name> struct {
    <field-name1> <field-type1>
    <field-name2> <field-type2>
    ...
}
```
Dot operator, like in C, can be used to access struct fields. Struct pointers can also use dot pointers for access of fields, without explicit dereference.

### Arrays

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.\
`var a [10]int` declares an array `a` of 10 elements of type `int`. Slices, on the other hand, are created by not specifying the length.\
An array variable denotes the entire array; it is not a pointer to the first array element (as would be the case in C). This means that when you assign or pass around an array value you will make a copy of its contents. (To avoid the copy you could pass a pointer to the array, but then that’s a pointer to an array, not an array.)

### Interfaces and methods

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).\
A pointer receiver is required to a method, if the underlying value of the receiver object is to be modified. Else the receiver is just treated as a copy.\
An _interface_ type is defined as a set of method signatures. A value of interface type can hold any value that implements those methods.\
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.\
Under the hood, interface values can be thought of as a tuple of a value and a concrete type, which makes sense since it has to "bottom out" somewhere.
> Note that an interface value that holds a nil concrete value is itself non-nil. 

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.\
The interface type that specifies zero methods is known as the empty interface: `interface{}`. It may hold values of any type. Used, for example, `fmt.Print` takes any number of arguments of type empty interface.\
`String()` method is Go's display trait for types.\
The `error` type is a built-in interface similar to `fmt.Stringer`

### Misc

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.\
All values are automatically initialized to a zero-value on declaration\