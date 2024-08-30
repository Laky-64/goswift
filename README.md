<!--suppress HtmlUnknownTarget -->
<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://www.swift.org/assets/images/swift~dark.svg">
  <img src="https://www.swift.org/assets/images/swift.svg" alt="Swift logo" height="70">
</picture>

# GoSwift
**Swift Demangler** is a Go port of Apple's official Swift libraries. 
This project aims to provide the functionality of the Swift language libraries, 
enabling Go developers to interact with Swift code, such as demangling obfuscated 
Swift function names and data structures (mangled names).

Based on the official Swift libraries, this project has been adapted to integrate with the
Go language syntax and ecosystem, offering developers a native solution for working with Swift 
symbols in Go environments.

Learn more about Swift and the official project here:
<a href="https://github.com/swiftlang/swift">Official Swift</a>.


# Swift Demangler CLI
![CLI](https://vhs.charm.sh/vhs-2y0mYj35QtJTdYXPESKWXf.gif)
The **Swift Demangler CLI** is a standalone executable written in Go, designed specifically
for demangling obfuscated Swift names (mangled names).
This tool decodes Swift symbols and function names generated
by the Swift compiler, converting them into a human-readable format.

# Library Usage
Use `go get` to download the dependency.

```bash
go get github.com/Laky-64/swift@latest
```

Then, `import` it in your Go files:

```go
import "github.com/Laky-64/swift/demangling"
```

The library is designed to be simple and easy to use.
Here's an example of a simple demangling:

```go
demangled, err := demangling.Demangle("Say12Smol.Animals3FoxCG5foxes_Su5countt")
if err != nil {
    panic(err)
}
fmt.Println(demangled)
```
