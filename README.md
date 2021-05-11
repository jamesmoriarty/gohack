# Gohack 

[![Latest Tag][6]][5] [![CI][3]][4] [![Go Report][1]][2] ![GitHub Releases][8]

Experimental Go language CSGO exploit proof-of-concept.

## Features

- Trigger Bot (hold shift).
- Bunny Hop (hold space).
- Offsets Autoupdate From [Hazedumper][9].
- Automated tests utilizing stubbed external processes to avoid needing binary compatability.

## Screenshots

![Screenshot](docs/screenshot.png)

## Usage

```
.\gohack.exe
```

## Download

You can download [here][5].

## Install

```
go get -v github.com/jamesmoriarty/gohack
go install github.com/jamesmoriarty/gohack
```

## What

- Injected from external process.
- Tested against stub process and dlls.
- Tested against local bots.
- Assume you will be banned on VAC secured servers.
- Credit to [1jz][7] which this was originally based on.

[1]: https://goreportcard.com/badge/github.com/jamesmoriarty/gohack
[2]: https://goreportcard.com/report/github.com/jamesmoriarty/gohack
[3]: https://github.com/jamesmoriarty/gohack/workflows/Continuous%20Integration/badge.svg
[4]: https://github.com/jamesmoriarty/gohack/actions
[5]: https://github.com/jamesmoriarty/gohack/releases
[6]: https://img.shields.io/github/v/tag/jamesmoriarty/gohack.svg?logo=github&label=latest
[7]: https://github.com/1jz/csgo-go-bhop
[8]: https://img.shields.io/github/downloads/jamesmoriarty/gohack/total
[9]: https://github.com/frk1/hazedumper
