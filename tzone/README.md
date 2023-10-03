### A CLI app that can show current time in any timezone and show the date in certain format

**DEMO**

```shell
gitpod /workspace/cli-gallery/tzone (main) $ go build && go install
gitpod /workspace/cli-gallery/tzone (main) $ tzone
Get the current date and time in diffferent formats as well as timezones

Usage:
  tzone [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  tzone       Get the current time in a given timezone

Flags:
  -h, --help     help for tzone
  -t, --toggle   Help message for toggle

Use "tzone [command] --help" for more information about a command.
gitpod /workspace/cli-gallery/tzone (main) $ tzone --help
Get the current date and time in diffferent formats as well as timezones

Usage:
  tzone [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  tzone       Get the current time in a given timezone

Flags:
  -h, --help     help for tzone
  -t, --toggle   Help message for toggle

Use "tzone [command] --help" for more information about a command.
gitpod /workspace/cli-gallery/tzone (main) $ tzone tzone --help
Get the current time in a given timezone.
        This command takes one argument, the timezone you want to get the current time in.
        eg - tzone tzone EST --date 2006-01-02

Usage:
  tzone tzone [flags]

Flags:
      --date string   returns the date in a time zone in a specified format
  -h, --help          help for tzone
gitpod /workspace/cli-gallery/tzone (main) $ tzone tzone EST
Tue, 03 Oct 2023 06:30:08 EST
gitpod /workspace/cli-gallery/tzone (main) $ tzone tzone EST --date 2006-01-02
Current date in EST: 2023-10-03
```