# Full guide

You may be thinking: "Why put all the commands in another file than the README"

Well the answer to that is that I don't want to extend the README so much, The commands in the README are the 3 ones that I made first so... that is why they get that place.

# Random Command

This command is here to help you make decisions (if you are willing to accept the result), also can generate an __pseudo-random__ number between the numbers that you want, Look it:

```bash
$ gocharm random -1 {<number 1>} -2 {<number 2>}
```
if you don't provide any flags, number 1 will default to zero and number 2 will default to 99.

## Decisions Subcommand

This subcommand is used... well to make decisions, See it i action:

```bash
$ gocharm random decisions -1 "{<decision 1>}" -2 "{<decision 2>}"
```

The... beutiful output:

```bash


Gocharm says:

   Go for {<decision 1>}.

Gocharm speaked. Gocharm out.

```

For decision 1, for decision 2:

```bash


Gocharm says:

   Your destiny is {<decision >}.

Gocharm speaked. Gocharm out.

```

Pretty cool, huh.

Well that's it for this __command.__

# Cron command

This command is here to help you measure time, Look it, it does not bite:

```shell
$ gocharm cron
```
As simple at that, Look at the output:

```shell
Cronometer running...
Hit enter to stop.

// 1 minute and 30 seconds later

Time elapsed: 1m30.274817462s
```

Well, another command documented.

# Timer Command

This command was born for your baking needs, Just look at your future with him:

```shell
$ gocharm timer -s 50 -m 30 -o 2
Timer fired!
Press enter to stop it.

// *Inhales deeply* TWO hours THIRTY minutes and FIFTY minutes later *Exales*

Timer expired!
Press enter to continue
```
The "-s" switch is for seconds, "-m" for... well... minutes (I think is obvious, the "m".), and "-o" for hours. Why "o" for hours, and why not a beutiful "h". Well it's because "h" already is for help flag.

And that's it, yet another command documented.

# Timer command

This command shined his way onto the command list to help you bake that beutifuo cake, Look at his SHINE:

```shell
$ gocharm timer -s 0 -m 30 -o 1
```

OOF, I'm blind now...... Oh I can see now, Look at his even more bright output:

```shell
Timer fired!
Press enter to stop it.

Seconds: 0
Minutes: 30
Hours: 1

Timer expired!
Press enter to continue!
```

AAAAAAAAHHH, MY EYES......IT'S SO BRIGHT....
I will explain the switches:

The "-s" switch means seconds, the "-m" one means minutes and "-o" means climate chan..ehem hours.

Why not "-h".

Because "-h" means HELP ME THE BRIGHTNESS OF THIS COMMAND BLINDED ME, The only true thing here is "HELP".

So that's it for this VERY BRIGHT command, now if you don't mind I will see the eyes doctor.

# Config command

This command rushed his way onto the command list with help of his ICP (INFINITE CONFIG POWER) to help you config the CLI (I think that is obvious), Look it:

```shell
$ gocharm config
```

Aha, It is indeed very simple but with INFINITE CONFIG POWER, Look at hist ICP output:

```shell
Creating config file...
? Choose a config file type:  [Use arrows to move, type to filter]
> YAML
  TOML
  JSON
  HCL                                                                    PROPERTIES
? Choose a config file type: YAML
? Set default timezone now? (y/N) Yes
? Default Timezone: America/New_York
✓ Succesfully saved default timezone
Config file created succesfully at PATH/TO/HOME/.gocharm.yaml
```

NOTE: All of the command output comes with an AICP (AWESOME and INFINITE COLOR POWER), And is colored with that AICP ;)

Then it will create a file name ".gocharm" with type "YAML" on your HOMEDIR.

That's it for this command, I will go recover from this rush.
