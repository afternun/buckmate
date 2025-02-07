# Commands

Common flags for every command:

| Flag   | Short | Description                                          | Default |
|--------|-------|------------------------------------------------------|---------|
| `--env`  | `-e`    | Specifies environment specific configuration         | -       |
| `--path` | `-p`    | Specifies directory that contains Deployment.yaml file | -       |

## apply

Applies deployment to the infrastructure.

1. Copy files from source address
2. Copy files from common configuration
3. Copy files from environment specific configuration (if `--env` flag provided)
4. Scan and replace any config map values
5. Upload complete package to target address

| Flag   | Short | Description                                          | Default |
|--------|-------|------------------------------------------------------|---------|
| `--dry`  | `-d`    | Dry run, do not upload files. Will output location of prepared deployment. Those files won't be used in future command calls.         | -       |

## config

Applies config to local files, without downloading files from `source`

1. Copy files from common configuration
2. Copy files from environment specific configuration (if `--env` flag provided)
3. Scan and replace any config map values