# Toolbox

Toolbox is a command-line interface (CLI) tool written in Go.It leverages the power of Cobra, a CLI library for Go, to offer a seamless user experience.

## Overview
<p align="center">
  <img width="612" alt="overview" src="https://github.com/YashPimple/toolbox-cli/assets/97302447/7abeb5d5-bad5-4a97-a250-0d544c881070">
</p>

## Usage

To use Toolbox, simply execute the following command:

```
./toolbox [command]
```


## Available Commands

- net: Perform network-related tasks using a comprehensive set of commands.
- info: Retrieve detailed information about all available commands.

## Flags
- -h, --help: Display the help information for Toolbox.
- -t, --toggle: Display a help message for the toggle feature.

## Getting Started
To get started with Toolbox, follow the steps below:

1. Clone the Toolbox repository from GitHub:
```
git clone https://github.com/YashPimple/toolbox-cli.git
```

2. Build the Toolbox executable:

```
go build -o toolbox
```

3. Run Toolbox with the desired command:
```
./toolbox [command]
```

For more detailed information about each command, use the following syntax:

```
./toolbox [command] --help
```

Contributing
We welcome contributions to Toolbox! If you would like to contribute, please follow these steps:

- Fork the Toolbox repository.
- Create a new branch for your feature or bug fix.
- Make the necessary changes and commit them.
- Push your changes to your forked repository.
- Submit a pull request to the main Toolbox repository.

## License
This project is licensed under the MIT License. Please see the LICENSE file for more details.
