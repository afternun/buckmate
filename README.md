# buckmate

buckmate is a CLI tool that allows you to easily deploy to S3 buckets. Tool was developed with frontend deployments in mind, but in reality this can be used to deploy anything to S3.

## How to

Start by creating file called `Deployment.yaml`

File should contain following entries:

```
source:
  path: NAME OF SOURCE S3 BUCKET
target:
  path: NAME OF TARGET S3 BUCKET
  tags:
    CUSTOM TAG KEY: CUSTOM TAG VALUE
```

Add another file called `Config.yaml` with following contents:

```
version: VERSION LABEL
configMap:
  CONFIG KEY: CONFIG VALUE
```

- if `version` provided, it will be appended to `source.path`
- buckmate goes over all files and replaces `%%%CONFIG KEY%%%` with `CONFIG VALUE`

Those files will be "global", but buckmate can work with environment specific configurations. You may add a directory called "dev" or "prod" and have there `Config.yaml` file that will be merged with the global one when applied.

You may also add "files" directory both in global and environment scope. All files in this directory will be copied over to the `target.path` on apply.

## Commands

- `buckmate apply --env ENVIRONMENT` - apply current deployment to the infrastructure, optional `--env` flag indicating environment 

## AWS configuration

You should configure your OS with AWS credentials, these should be picked up automatically by buckmate.

## Notes

- buckmate will add metadata with unique version (unique per deployment) to each object. This is used to determine what files to remove on subsequent deployments - next deployment removes previous deployment's files
- buckmate does not remove old versions (S3 versioning versions)