# 29

![latest version](https://img.shields.io/github/v/tag/29-FYI/29?label=latest%20version)
![last commit](https://img.shields.io/github/last-commit/29-FYI/29)
![top language](https://img.shields.io/github/languages/top/29-FYI/29)
![license](https://img.shields.io/github/license/29-FYI/29?label=license)

This is command line tool for printing the latest posts from [29.FYI](https://29.fyi). It outputs for humans, not machines.

## Installation

Run `go get -u github.com/29-FYI/29`.

## Usage

Run `29` to print the latest 29.FYI posts, assuming your `$PATH` contains your `$GOPATH/bin`.
You can then pipe the output for a better view.
Such as `29 | nl | tac`.

## Screenshot

![Output of `29 | nl | tac` command](29-nl-tac.png)

## License

This project is licensed under the GNU General Public License v3.0.
