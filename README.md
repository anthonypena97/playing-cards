# playing-cards
![MIT License badge](https://img.shields.io/badge/license-MIT_License-green)

![playing-cards-demo](https://user-images.githubusercontent.com/79285555/167043136-1781f0ba-3613-4fde-bf4f-58c6f7b4b423.gif)

<p> ---------------------------------------- </p>
Program for simulating playing with a deck of cards. Written in Go.

The following Go Standard Library packages are imported: fmt, io/ioutil, math/rand, os, strings, and time

## Table of Contents

- [Usage](#usage)
- [Installation](#installation)
- [Notes](#notes)
- [Testing](#testing)
- [License](#license)
- [Version History](#version)
- [Contributing](#contributing)
- [Questions](#questions)
- [Acknowledgmenets](#acknowledgments)

## Usage

In order to use application, you must download this repository and run program in terminal.

## Installation

To install:

- download latest version of [Go](https://go.dev/dl/) on your machine
- download, fork, or clone this repository
- open terminal in application root directory
- enter,`go run main.go deck.go`

* a deck of 16 cards randomly shuffled shall be printed to the console *

## Notes

The application as written, uses the newDeck(), print() and shuffle() methods. 

In order the use the deal(), saveToFile() and newDeckFromFile() methods, one must modify the main.go file and call the methods as desired.

## Testing

To run tests:

- after installation, open terminal in application root directory
- enter, `go test`

* a pass or fail message will be printed to terminal console*

## License

MIT License

Copyright (c) 2022 Anthony Pena

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation fil (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge,publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so,subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

## Version History

- v1.0.0 is latest version
  - - See [commit change](https://github.com/anthonypena97/playing-cards/commits/main) or See [release history](https://github.com/anthonypena97/playing-cards/releases)

## Contributing

Please refer to the [Contributor Covenenant](https://www.contributor-covenant.org/) for guidelines on contributing on this project.

## Questions

For any inquiries or questions, please contact Anthony Pena via:

- GitHub: [anthonypena97](https://github.com/anthonypena97)
- Email: <anthony.e.p3na@gmail.com>

## Acknowledgments

Inspiration, code snippets, etc.

- [Go: The Complete Developer's Guide(Goalang)](https://www.googleadservices.com/pagead/aclk?sa=L&ai=DChcSEwjo7b_HxMn3AhWK4bMKHfz2B2oYABAAGgJxbg&ae=2&ohost=www.google.com&cid=CAESbeD2vywAvEPOK1nuVHlm8maYKSepwZTzG1txyvM5da19gslLTgYJDe1EXna4UveyHNqtoIaHVT81PuzC3Fm_Z2KIrQYws-sh20oLooafXWZrtE3MedtjEqen1ps6Ia01ISxKHiJq_H6VLT7WbKg&sig=AOD64_23Y8EHs-FJ2l8vuKFKXwE-w83J-Q&q&adurl&ved=2ahUKEwjV-rLHxMn3AhU2j4kEHezAAMwQ0Qx6BAgCEAE)
