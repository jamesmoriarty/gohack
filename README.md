# Gohack 

[![CI][3]][4] [![Latest Tag][6]][5] [![Go Report][1]][2] ![GitHub Releases][8]

Experimental Go language proof-of-concept CSGO exploit. Automated tests use stubbed external processes to avoid needing binary compatability. Inspired [github.com/jamesmoriarty/gomem](https://github.com/jamesmoriarty/gomem).

## Features

- Trigger Bot (hold shift).
- Bunny Hop (hold space).
- Offsets autoupdate from [Hazedumper][9].

## Slides

[Talk - Writing My First Exploit](https://speakerdeck.com/jamesmoriarty/talk-writing-my-first-exploit)


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

[1]: https://goreportcard.com/badge/github.com/jamesmoriarty/gohack
[2]: https://goreportcard.com/report/github.com/jamesmoriarty/gohack
[3]: https://github.com/jamesmoriarty/gohack/workflows/Continuous%20Integration/badge.svg
[4]: https://github.com/jamesmoriarty/gohack/actions
[5]: https://github.com/jamesmoriarty/gohack/releases
[6]: https://img.shields.io/github/v/tag/jamesmoriarty/gohack.svg?logo=github&label=latest
[8]: https://img.shields.io/github/downloads/jamesmoriarty/gohack/total
[9]: https://github.com/frk1/hazedumper
