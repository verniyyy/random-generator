# Random Generator (rgen)

## Overview:
Random Generator (rgen) is a simple CLI tool designed to generate random values such as UUIDs and ULIDs effortlessly. This utility provides a quick and reliable way to create unique identifiers for various use cases.

## Features:
* UUID Generation: Create universally unique identifiers (UUIDs) with ease.
* ULID Generation: Generate sortable and highly entropic ULIDs for distributed systems.
* Straightforward CLI Interface: Simple commands for seamless integration into scripts or workflows.
* Customizable Count: Use the -n option to specify the number of values to generate.

## Installation:
Ensure you have Go installed, then run:

```sh
make install
```

## Usage:
```sh
rgen uuid [-n count]
rgen ulid [-n count]
```

## Options:
* `-n count``: Specifies the number of values to generate. Default is 1.

## Examples:

### Generate 5 UUIDs:
```sh
rgen uuid -n 5
```

### Generate 3 ULIDs:
```sh
rgen ulid -n 3
```

## Contributing:
Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request on the GitHub repository.

## License:
This project is licensed under the MIT License.