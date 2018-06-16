# Vanity-ETH cli

[![Build Status](https://img.shields.io/travis/bokub/vanity-eth-cli/master.svg?style=flat-square)](https://travis-ci.org/bokub/vanity-eth-cli)
[![License](https://img.shields.io/badge/license-MIT-ff9860.svg?style=flat-square)](https://raw.githubusercontent.com/bokub/vanity-eth-cli/master/LICENSE)
[![Codecov](https://img.shields.io/codecov/c/github/bokub/vanity-eth-cli.svg?style=flat-square)](https://codecov.io/gh/bokub/vanity-eth-cli)

Simple and efficient ETH vanity address generator.

Available for Linux, Windows and macOS.

## What's a vanity address?

A vanity address is an address which part of it is chosen by yourself, making it look less random.

Examples: `0xc0ffee254729296a45a3885639AC7E10F9d54979`, or `0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E`

## How to use

1. Find and download the [latest release](https://github.com/bokub/vanity-eth-cli/releases) for your system

2. If you want to run Vanity-ETH from your terminal, run `chmod +x vanity-eth-cli-*`

3. Run the binary with a double-click, or run `./vanity-eth-cli-*` from your terminal


## Build from source

If the binary release doesn't suit you, you can build Vanity-ETH cli from source

You'll need:

- **Go** - version `1.7` or later
- **Glide** (optional) - install with `curl https://glide.sh/get | sh` or check out the [glide documentation](https://github.com/Masterminds/glide#install)

```bash
# Get Vanity-ETH cli
go get github.com/bokub/vanity-eth-cli
cd $GOPATH/src/github.com/bokub/vanity-eth-cli

# Download dependencies (optional but recommended)
glide install

# Build Vanity-ETH cli
go build
```
