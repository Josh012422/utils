# Description

A personal project to learn how to make a **_cli_** with **_cobra_** and go

Make your terminal life easier.

# Installation

To install, clone the repo then simply run:

```bash
$ make
```

Let it finish, then:

```bash
$ echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bash_profile
```

If you don't have **_make_** installed in your machine, then search Google.

**Hint:** If you're on '**_Windows_**' then install choco, then run:

```bash
choco install make
```
**_Hint:_** If you are using termux, just run:

```bash
$ pkg install make
```

# Usage

The '**_main_**' command is `gocharm`, there are currently two commands (There will be more), a list:

1. time
2. tasks

Each one with his own functionality.

## Time Command

This command displays the REAL TIME of any city you want, for example:

```bash
$ gocharm time -t "America/New_York" -1
```

The "-t" flag tells gocharm what city to look the time of, like the timezone.

The "-1" flag tells gocharm to display the time in 12-Hour Format

### Compare Subcommand

There's a subcommand of time that lets you display the time of two different timezones. The Compare Command.

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

NOTE: There's also other subcommands like '**_list_**', '**_view_**', '**_complete_**' .
