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