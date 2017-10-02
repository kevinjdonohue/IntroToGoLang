# Intro To Go

A repo for storing notes and code examples while taking the PluralSight course The Go Programming Language by John Sonmez

## Module 1:  Overview

3 Tenants of Go:

1. Compiled
2. Garbage Collected
3. Concurrent

### Compiled

Go can target multiple OSs:

* Linux
* MacOS
* FreeBSD
* OpenBSD
* Plan 9
* Windows

Processors:

* i386
* amd64
* ARM

### Garbage Collected

Fairly unusual feature of a compiled language - as compared to languages like C and C++.  Contrasted with C# or Java which are compiled languages but are compiled to an intermediate language.

Go's garbage collection does not introduce latency - so it is arguably better than other garbage collection implementations.

### Concurrent

Concurrent (def):  Literally doing more than one thing at the same time!

Unlike many other languages, concurrency is built into the language instead of being an add-on to it (as it is for many other languages).

The combination of these three key features or tenants of the language allow you to create fast, powerful applications.

Examples of things written in Go:

* Systems-level programs

Go Routines - constructs in the language.

Passing data instead of sharing data.

### Language

When reading Go code it (at least at first) appears similar to C.

NOT a derivative of C - it is it's own language.

### Origins of Go

Started out as a project at Google by three engineers.

#### Purpose:  

They wanted a modern Systems-level language.

Simple and fast language.

Fast compile times.

#### History

* september 2007 - Go dreamed up by Robert Griesmer, Rob Pike and Ken Thompson
* november 2009 - Go officially announced
* may 2010 - Go is used internally at Google
* march 2012 - Go 1 is released

#### What Makes It Different?

* Efficient like static language
* Ease of use like a dynamic language
* Type safe
* Memory safe
* Latency free garbage collection
* Fast compile times
* No inheritance!
* No generics? (possibly in the future?)
* No assertions
* No method overloading
* Different style of Interfaces
* No classes; user defined types
* Systems-level language - modern language replacement for C or C++

## Module 2:  Go Development

### Tools

* Eclipse
* IntelliJ (author's preference) - http://go-ide.com
* Gogland (Jetbrains) - current, attractive offering in 2017

You can also do Go development in common text editors.

### Packages

* Way to modularize code
* Similar to namespaces
* Types and functions

### Imports

Code.go --> Source Code, Local Packages
        |--> Remote Packages

Mechanism to utilize Packages

### Workspaces

As an alternative / improvement over other languages, Go uses convention over configuration -- no build files are necessary for example -- given that you setup your workspace in a partciular way.

* bin
* src
* packages

The only folder that need to create is the ```src``` folder

### Hello World

The following is the hello.go code from the course:

```go
package main

import "fmt"

func main() {
        fmt.Println("Hello, Go")
}
```

A few things to note about a .go file:

* Since we declared our file to be a package called main we need an entry point, a func, to also be called main.
* The import keyword is used to import a package - in this case the format (fmt) package that comes with Go is being used
* Note that the Println func from the fmt package is capitalized -- capitalization is the mechanism for controlling visibility in Go
* Capitalized func names == exported function
* Lowercase func names == private; non-exported function

### Installing a Go Program

In order to install a Go program, you need to make use of the Go Workspace concept.

1. Create an enviornment variable called GOPATH that points to one or more locations on your system.  (e.g. c\GoWorkspace\)
1. Create a folder structure that conforms to the Go convention
   * {workspace}\src\{domain_or_namespace}\{project_name}
   * e.g. GoWorkspace\src\MyDomainName\Hello
1. From the ```src``` folder, run: ```go install MyDomainName\Hello```
1. Note that there is now a ```bin``` folder at the same level as ```src``` that contains the .exe

## Module 3:  Variables, Types and Pointers

### Basic Types

* bool
* string
* int, int8, int16, int32, int64
* uint, uint8, uint16, uint32, uint64, uintptr
* byte (uint8)
* rune (int32), like a char
* float32, float64
* complex64, complex128

### Additional Types

* Array - not as common; Slice used instead
* Slice - vector or a list; often used instead of Array
* Struct - similar to Structs from other languages
* Pointer - memory address
* Function - you can store or pass a function around; similar to JavaScript
* Interface - Unlike other languages interfaces; no explicit implementation required to conform to an interface
* Map - Dictionary-like; Map like in Java
* Channel - communication mechanism between Go routines

### Variable Declaration and Initialization

Three different forms can be utilized to declare initialize variables in Go

1. long form
1. shorter form
1. short form

```Go
var message string = "long form"

var message2 = "shorter form - no type"

message3 := "short form - no var or type"
```

### Pointer Basics

Here is a basic pointer example:

```Go
message := "Hello Go World"
greeting *string = &message  //this obtains the memory address for the message string

fmt.Println(greeting)  //this prints the memory address
fmt.Prinln(*greeting) //this prints the actual string value
```

### Passing Values

By default when values are passed into a function (func) the value is copied.  As an alternative, you can pass in a pointer - which will allow you to essentially pass a reference to the value instead of a copy.

### No Classes; User Defined Types Instead

Instead Go has User Defined Types.

User Defined Types can leverage any type from Go, say string or more usefully can leverage struct.  Here are a couple examples.

```Go

//created a user defined type based on a string
type Name string

var myName Name = "Kevin"

//created a user defined type based on a struct
//more interesting here because we can access the properties of the struct similar to getters on a class

type Salutation struct {
        name string
        greeting string
}

var mySalutation = Salutation{"Joe", "Hello!"}
```

### Constants

Constants can be declared by themselves like a variable or as a group using the ```Const()```syntax.

```iota``` is a shorthand in Go for incrementing numbers.

```Go

const (
        PI = 3.14
        Language = "Go"
)

const (
        A = iota
        B
        C
)

//printing A, B, C would result in 0, 1, 2

const (
        Mute = iota
        Mono
        Stereo
        _
        _
        Surround
)

//printing Mute, Mono, Stereo, Surround would result in 0, 1, 2, 5 because the _ allows us to skip values in the sequence

```

## Module 4:  Functions

Functions in Go are Types unlike in many other languages.

### What is a function in Go

* Can have multiple return values
* Use them like any other type (pass them around, return them, etc. -- think functions in JavaScript)
* Function literals - you can declare a function inside another function; allows for closures

There are a few different ways to setup a function.  Here are a few examles of those different ways.

Note:  In go if your function parameters are of the same type, you simply list the parameters and then the type after the last one - you can see this in the example function ```CreateMessage``` example below.

```Go

//simple function with a returna type; string
func CreateMessage(name, greeting string) string {
        return greeting + " " + name
}

//function with multiple (two) return values; strings
func CreateMessageWithTwoReturnValues(name, greeting string) (string, string) {
        return greeting + " " + name, "HEY! " + name
}

//function with multiple, named return values
func CreateMessageWithNamedReturnValues(name, greeting string) (message string, alternate string) {
        message = greeting + " " + name
        alternate = "HEY! " + name
        return
}

//variadic function - variable number of parameters indicated by the ...
func CreateMessageVariadic(name string, greeting ...string) (message string, alternate string) {

        message = greeting[0] + " " + name
        alternate = greeting[len(greeting)-1] + " " + name
        return
}

```

### Function Types & Function Literals

In Go you can declare function types.

Here is an example:

```Go

type Printer func(string) ()

//now we can create a function that returns this type + makes use of a closure
function CreatePrinterFunction(custom string) Printer {
        return func(s string) {
                fmt.Println(s + custom)
        }
}

```

## Module 5:  Branching

* If statement -  similar but not the same as in other languages
* Switch statement - completely different than other languages

### If Statements

```Go

if someBool {

} else {

}

//a funky twist is that you can declare and use a variable in the scope of the if statement - called embedded statements:

if prefix := "Foo "; someBool {
        doSomething(prefix)
} else {
        doSomethingElse()
}

```

### Switch Statements

* No default fall through
* Don't need an expression
* Cases can be expressions
* Can switch on types

Examples:

```Go

switch name {
        case "Bob":
                prefix = "Dr"
        case "Mary":
                prefix = "Mrs"
        case "Joe":
                prefix = "Mr"
        default:
                prefix = "Sir"
}

//alternatively, you can switch on a boolean value
switch {
        case name == "Bob":
                prefix = "Mr"
        case name == "Mary":
                prefix = "Mrs"
        case name == "Joe", name == "Amy", len(name) == 10:
                prefix = "Dr"
        default:
                prefix = "Sir"
}

//also, you can switch on types
switch t := x.(type) {
        case int:
                fmt.Println("int")
        case string:
                fmt.Println("string)
        default:
                fmt.Println("unknown")
}

```