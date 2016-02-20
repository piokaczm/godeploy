###GoDeploy *work in progress*

*it's actually a spec for me while it's in development :)*
Simple deploy tool:
- build a binary using cross-compilation
- copy over to a remote server
- set additional commands
- integrate with slack

###Configuration

First of all you need to properly set your configuration file. Structure should look something like this:

```
appname: NewApp # used for slack integration - it'll be used in msgs
goos: linux # variables needed to properly crosscompile for your machine
goarch: amd64
test: true # only if you want to run all test before deploy, defaults to false
godep: true # only if you're using godep and want to run godep restore before building a binary, defaults to false

environments:
  staging: # this is cli argument you'll be using to deploy to choosen env
    host: pizda.net # no need for specyfing number when deploying to one host only
    user: pizdek
    path: binaries/
    restart_command: etc/dupa/daemon restart
  production:
    host_1: real-pizda.net # the tool matches host and user using the provided digit, so make sure to fill it properly
    user_1: pizdekmaster
    host_2: real-pizda2.net
    user_2: pizdekmaster2
    path: current/binaries/
    restart_command: etc/prod/dupa/daemon restart

slack: # optional
  webhook: https://hooks.slack.com/services/sth/more
  emoji: ":rocket:"
  botname: bot
```
