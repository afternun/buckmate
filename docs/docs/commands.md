# Commands

Common flags for every command:

| Flag   | Short | Description                                          | Default |
|--------|-------|------------------------------------------------------|---------|
| `--env`  | `-e`    | Specifies environment specific configuration         | -       |
| `--path` | `-p`    | Specifies directory that contains buckmate directory | -       |
| `--log` | `-l`    | Specifies log level, available options: panic, fatal, error, warn, info, debug, trace | info       |

## apply

Applies deployment to the infrastructure.

1. Copy files from source address
2. Copy files from common configuration
3. Copy files from environment specific configuration (if `--env` flag provided)
4. Scan and replace any config map values
5. Upload complete package to target address

## config

Applies config to local files, without downloading files from `source`

1. Copy files from common configuration
2. Copy files from environment specific configuration (if `--env` flag provided)
3. Scan and replace any config map values