<h1 align="center">
  <br>
  <img src="assets/logo.svg" alt="aos" width="400">
  <br>
  <br>
</h1>

<h5 align="center">aos - Alias on Steroids</h5>
<h4 align="center">A simple command line tool that helps you run complex command lines.</h4>


<p align="center">
  <img src="https://github.com/todoesverso/aos/actions/workflows/main.yaml/badge.svg">
  <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square">
</p>

<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#usage">Usage</a> •
  <a href="#installation">Installation</a> •
  <a href="#download">Download</a> •
  <a href="#roadmap">Roadmap</a> •
  <a href="#license">License</a>
</p>

## Key Features

* Save complex command lines in a human readable yaml file and run it
* Add positional arguments
* Get detailed information on the command line 
* Render the final command

## Usage

  ```sh
  $ aos
  AliasOnSteroids

  Usage:
        aos <alias.yaml> [positional arguments of the alias]
  Options:
        In order to keep CLI arguments as straightforward as possible,
        options are passed thru the AOS environment variables.

        AOS=h   ./builds/aos <alias.yaml>       # Prints this usage
        AOS=H   ./builds/aos <alias.yaml>       # Prints a helper short description of the alias
        AOS=E   ./builds/aos <alias.yaml>       # Prints a helper long description of the alias
        AOS=R   ./builds/aos <alias.yaml>       # Renders the command and prints it to stdout
        [AOS=X] ./builds/aos <alias.yaml>       # Runs the command in a shell.

  $ cat shakespeare.txt | logss -c to -c be -c or -c 'in.*of'
  $ # 
  $ cat real_curl_example.yaml
    command:
      - curl
      - -s
      - https://raw.githubusercontent.com/linuxacademy/content-elastic-log-samples/master/access.log
    render: 75
    containers:
      - GET
      - "404"
      - ".*ERROR|error.*"
  $ logss -f real_curl_example.yaml 
  ```

## Installation

## Download

Pre compiled binaries for several platforms can be downloaded from the [release](https://github.com/todoesverso/aos/releases) section.

## Roadmap

This is just a personal project intended to learn some Go but I'll keep adding features that I see useful or requested by the users. 

## License

MIT
