# Description

A personal project to learn how to make a cli with cobra and go

Make to do your terminal life easier

# Installation

To install, clone the repo the simply run:

```bash
$ make
```

Let it finish, then:

```bash
$ echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bash_profile
```

If you don't have `make` install search Google.

Hint: If your on windows installl choco, then run:

```bash
choco install make
```
Hint: If you are using termux, just run:

```bash
$ pkg install make
```

# Usage

The "main" command is gocharm, there are currently 2 commands (There will be more), a list:

1. time
2. tasks

Each one with his own functionality.

## Time Command

This command displays the REAL TIME of any city you want, For example:

```bash
$ gocharm time -t "America/New_York" -1
```

The "-t" flag tells gocharm what city to look the time of, like the timezone.

The "-1" flag tells gocharm to display the time in 12-Hour Format

### Compare Subcommand

There's a subcommand of time that lets you display the time of 2 different timezones, The Compare Command.

```bash
$ gocharm time compare
```

If don't provide any flag, gocharm will prompt for Timezone 1 and Timezone 2.

# Tasks Command

The tasks command is a small and simple task manager, it lets you add and save tasks to a file in your systems, so if you prompted title wrong you can go and modify it.

```bash
$ gocharm tasks add
```

The above command will add a new task with a number a save it in your system.

Note: There's also other subcommands like: list, view, complete (Not finished yet) but for this README the above one is enough.
