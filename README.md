# Crackle

Crackle is a package manager for docker applications [crackle.sh](https://crackle.sh)

_Crackle is currently Alpha software with an ever changing core API._

## Install
`cr` is the CLI tool that powers Crackle.
```
go get github.com/sunshinekitty/cr
```

## Configure
`/etc/crackle` holds server configs `$HOME/.cr` holds client configs and can also hold server configs.

See [config/](config/) for other examples of config files.

## Running

Crackle is still alpha software.  To run it will require a Postgres database.  You can initialize the schemas by running the migrations in [db/migrations/](db/migrations/) with a tool such as [mattes/migrate](https://github.com/mattes/migrate).

## Examples

Upload a crackle application config using Crackle:
```
$ cr upload config/package-example.toml
Created package testing
```

Pull down a config for a Crackle application that exists on the server:
```
$ cr get testing
Downloaded config for testing
```

At this point you can execute the Crackle package with the executable located in `$HOME/.cr/bin` or calling `cr execute [package]` directly.

```
$ ~/.cr/bin/testing  
Go executable executed with Crackle!

$ cr execute testing
Go executable executed with Crackle!
```

### What happened when we executed?

Crackle looked in our `$HOME/.cr/packages` directory for a config for that package.  After it was found the config was read and used to construct a `docker run` command for that application.  Crackle then executed that `docker run` command and printed the output to stdout.