---
layout: page
title: Getting Started with the Hedwig CLI
subtitle: The easiest way to use Hedwig!
---

First, make sure that you have the Hedwig app, and the Hedwig CLI installed. You can find out how to install the Hedwig app and the Hedwig CLI on the [installation instructions](../installation.html) page.

You can check to make sure that Hedwig is fully installed by running `hedwig -version`. You should get an output that looks something like this.

```shell
$ hedwig -version
🦉 Hedwig v0.1.0 (darwin/amd64)
```

Then, take note of your Push Token. It should start with `ExponentPushToken`. It can be found by opening the Hedwig app. You can then configure the Hedwig CLI to use your Push Token.

> **Note!** Replace `yourPushToken` with your actual Push Token! Also, make sure to put your Push token in quotes, otherwise, shells tend to get confused.

```shell
$ hedwig -add-recipient-global "yourPushToken"
Added recipient to global config
```

You are now ready to send your first push notification!
```shell
$ hedwig "Hello, world\!"
```

Hedwig will send a push notification to your phone!

{% include push.html content="Hello, world!" %}

You can also include a title in your push notification.
```shell
$ hedwig -title "Upcoming class\!" "You have \"Introduction to Charms\" in 10 minutes."
```
{% include push.html title="Upcoming class!" content="You have \"Introduction to Charms\" in 10 minutes." %}

You can use this along with `cron` to make a makeshift calendar system! Say you have Defence Against the Dark Arts II Monday through Friday at 11am. You could add a Hedwig command to your `crontab` to notify you every day.

```cron
55 10 * * 1-5	hedwig -title "Upcoming class\!" "You have \"Defence Against the Dark Arts II\" in 5 minutes."
```

{% include push.html title="Upcoming class!" content="You have \"Defence Against the Dark Arts II\" in 5 minutes." %}