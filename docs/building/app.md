---
title: Building the Hedwig app
subtitle: The Hedwig app is written in React Native and can easily be built from source.
layout: page
---

The Hedwig app is written in React Native and can easily be built from source. First, if you haven't already, clone the Hedwig repository. If you don't know how to use Git, Atlassian has some nice tutorial [here](https://www.atlassian.com/git). Next, install Node.js and NPM. They install together, and you can get the latest version at the [Node.js website](https://nodejs.org/en/).

You can then install the [Expo CLI](https://expo.io) with NPM:

```bash
$ npm install -g expo-cli
```

Then, `cd` into the `hedwig-app` directory. You can then install dependencies with NPM:

```bash
$ npm install
```

To start a development server, you can then run

```bash
$ npm start
```

To build a production version, first, use the Expo CLI to sign in. You can then trigger an iOS or Android build.

```bash
$ expo login # log into expo
$ expo build:ios # build an iOS version
$ expo build:android # build an Android version
```