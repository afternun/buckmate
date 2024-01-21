1. Install **buckmate**

=== "Build from source"

    ```
    git clone afternun/buckmate
    go build
    mv buckmate /usr/local/bin
    ```

=== "Download latest release from GitHub"

    ```
    curl link && mv buckmate /usr/local/bin
    ```

!!! note "Configure AWS credentials"

    If you are deploying to or from AWS S3 bucket configure AWS credentials according to their instructions.

2. Create **buckmate** directory

!!! note "Examples"

    Take a look at `example` directory in the code repository

   In the directory create:

* `Deployment.yaml` - here you define common configuration for your deployment, one that is shared across any environment that you work on
    
    ```
    source:
      address: location from which files should be copied 
      (use `s3://` prefix for s3 buckets,
       absolute path for files on disk,
       or path relative to root directory)
    target:
      address: location to which files should be copied
      (use `s3://` prefix for s3 buckets,
       absolute path for files on disk,
       or path relative to root directory)
    configBoundary: string that acts as prefix and suffix for config map values (Default %%%)
    configMap:
      string key: string value
    ```

!!! note "Config Map"

    **buckmate** will go over files downloaded from `source` and files defined in `files` directory and look for strings that are wrapped in `configBoundary`. If such string is found, it will be replaced with corresponding value from `configMap`. 
    
    Example: If a file would contain string `%%%header%%%` and `configMap` an entry `header: My Awesome Header`, string `%%%header%%%` would be replaced with `My Awesome Header`. 

    **Environment specific configuration takes precedence over common configuration**

* (Optional): `files` directory

    This can hold any files that will be copied alongside files downloaded from `source`

* (Optional): directory with name of your choosing with another `Deployment.yaml` and `files` directory

    This can hold environment specific configuration. To use it, run **buckmate** with `--env` flag

3. Run **buckmate**

```
  buckmate apply
```