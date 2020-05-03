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
$ gocharm tasks {<subcommand>}
```

Here's a list of __all__ subcommands:

1. __add__
2. __list__
3. __view__
4. __complete__

## Add subcommand

This subcommand is used to add a new task and save it, have a look at it:

```bash
$ gocharm tasks add -t "{<Task Title>}"
```

This will save the task on the .gocharm.yml file that has been created when you executed 'make'.

NOTE:The file is not always yaml, you will be prompted for the file type.

## List subcommand

This subcommand is used to list __every single one__ of your tasks. Look it:

```bash
$ gocharm tasks list
```

For the sake of demonstration, I will show you the output:

```bash
✓. {<Task Title>}  // Completed
2. {<Other Task Title>}  // Pending
```

If completed will show a __✓__ sign and that comment telling you that the task is completed, if the task is not completed it will show the task number.

## View subcommand

This subcommand is used to view a __single__ task, Behold the subcommand:

```bash
$ gocharm tasks view {<Task Number>}
```

It will show you the __title__ and the __status__.

## Complete subcommand

This command is used to mark tasks as completed, Look at his terminal eyes:

```bash
$ gocharm tasks complete {<Task Number>}
```

And that's it, the rest of commands will be documented in [this file.](/docs/full_guide.md)