---
title: Building the Hedwig documentation
subtitle: The Hedwig app is written using the Jekyll templating engine.
layout: page
---

First, if you haven't already, clone the Hedwig repository. If you don't know how to use Git, Atlassian has some nice tutorial [here](https://www.atlassian.com/git).

Next, install Ruby and Bundler. I won't go into installing Ruby here, because it varies greatly by operating system, but macOS and many Linux distributions ship with it. To install bundler, you can use RubyGems.

```bash
$ sudo gem install bundler
```

You can then use Bundler to install dependencies.

```bash
$ bundle install
```

You can then start a development server with Jekyll.

```bash
$ bundle exec jekyll serve
```