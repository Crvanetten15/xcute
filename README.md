# xcute

## Overview

`xcute` is a Command Line Interface (CLI) tool designed to manage, find, and run your executables from one location. It simplifies the process of handling various scripts and binaries, providing a centralized way to execute them efficiently.

## Features

- **Discoverability:** Quickly list all executables in the specified directories.
- **Management:** Add, remove, or update the executables in your watch list.
- **Execution:** Run executables directly through `xcute` with command-line arguments support.
- **Portability:** Easy to set up and use across different environments.

## Getting Started

### Prerequisites

- Go 1.15 or higher

### Installation

To install `xcute`, clone the repository and build the binary using Go:

```bash
git clone https://github.com/yourusername/xcute.git
cd xcute
go build -o xcute .
```

Usage
To start using xcute, you can run it directly from the command line:

```bash
./xcute [command] [options]
```
For example, to list all executables in the current directory:

```bash
./xcute list
```
### Commands
- list: Lists all executables xcute is tracking.
- add [path]: Adds a new executable to xcute.
- remove [name]: Removes an executable from xcute.
- run [name] [args]: Executes an executable managed by xcute.

### Configuration
Explain how users can configure xcute, including where the configuration files are located and how they can be modified.

### Contributing
Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.

### Fork the Project
Create your Feature Branch (git checkout -b feature/AmazingFeature)
Commit your Changes (git commit -m 'Add some AmazingFeature')
Push to the Branch (git push origin feature/AmazingFeature)
Open a Pull Request

### License
Distributed under the MIT License. See LICENSE for more information.

### Contact
Your Name – @ettensoftware

Project Link: https://github.com/Crvanetten15/xcute


Ensure you replace placeholder texts like `yourusername`, `Your Name`, and contact details with your actual information. This template provides a basic structure that includes an overview, features, installation and usage instructions, contribution guidelines, and licensing info.