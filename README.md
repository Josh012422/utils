# Gocharm

[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=Josh012422_utils)](https://sonarcloud.io/dashboard?id=Josh012422_utils) [![CircleCI Build](https://circleci.com/gh/Josh012422/gocharm.svg?style=svg)](https://app.circleci.com/pipelines/github/Josh012422/gocharm)

A personal project to learn how to make a **_cli_** with **_cobra_** and go.

Make your terminal life easier.

# Feedback

If you want any feature or command in particular you can open a issue or write me: djblueslime11@gmail.com

Also if you want you can fill this [form.](https://forms.gle/8fAoibD6MuwoNphBA)

# Contributing

If you spotted a bug, I encourage you to open a issue or if you want you can fork the repo, fix it, then create a pull request.

Also if you added a new feature or something you can do the same proccess.

Another great way to contribute is creating tests files, I don't know how to do that so that will help me A LOT.

# Supporting

This project is COMPLETELY FREE and open-source so if you want to make me a generous donation, you can do it [here.](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=57L8EZDJATSKL&source=url)

Or you can click the button that says SPONSOR above. (More options for sponsor will be available in the future)

# Installation

To install, clone the repo: 

### HTTPS
```
https://github.com/Josh012422/gocharm.git
```

### SSH
```
git@github.com:Josh012422/gocharm.git
```

That will help you if you are in mobile using TERMUX, then simply run:

```bash
$ cd path/to/repo && make
```

Let it finish, then:

```bash
$ echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bash_profile
```

If you don't have **_make_** installed, search Google how to install it. (It is IMPORTANT that is GNU Make)

**Hint:** If you're on '**_Windows_**' install choco (Search Google how), afterwards run:

```bash
choco install make
```
**_Hint:_** If you are using **__TERMUX__**, just run:

```bash
$ pkg install make
```

# Usage

The '**_main_**' command is `gocharm`, there are currently eight commands (There will be more), a list:

- __time__
- __tasks__
- __config__
- __cron__
- __timer__
- __welcome__
- __random__
- __date__

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
