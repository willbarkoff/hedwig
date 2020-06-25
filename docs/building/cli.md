---
title: Building the Hedwig CLI
subtitle: The Hedwig app is written in Go and can easily be built from source.
layout: page
---

The Hedwig app is written in Go and can easily be built from source. First, if you haven't already, clone the Hedwig repository. If you don't know how to use Git, Atlassian has some nice tutorial [here](https://www.atlassian.com/git). Next, install Go. You can get the latest version at the [Go website](https://golang.org).


Then, `cd` into the `hedwig` directory. You can then install dependencies with Go:

```bash
$ go get
```

To build and install the program, simply run

```bash
$ go install
```
