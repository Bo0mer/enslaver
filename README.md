# enslaver
Master of all os-agents arount the world.

## Running it

### Configuration
Create configuration directory, which should contain `config.yml` file. You can preview example config file in the config/config.yml.

When you're done with the file, you should provide the config to the application by exporting a env variable `ENSLAVER_CONFIG_DIR` containing the full path to the directory, where your config.yml file is stored.

### Actual run
As simple as that:

```bash
go run main.go
```
