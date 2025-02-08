# Deployment.yaml
### source.address
* ##### Type: `string`
* ##### Required: `true`
* ##### Default: `none`

Location from which files should be copied (use `s3://` prefix for s3 buckets, absolute path for files on disk, or path relative to location of this file).

### source.prefix
* ##### Type: `string`
* ##### Required: `false`
* ##### Default: `none`

Prefix to use when downloading files from bucket. Useful if your bucket holds many versions or various types of data.

### target.address
* ##### Type: `string`
* ##### Required: `true`
* ##### Default: `none`

Location to which files should be copied (use `s3://` prefix for s3 buckets, absolute path for files on disk, or path relative to location of this file).

### configBoundary
* ##### Type: `string`
* ##### Required: `false`
* ##### Default: `%%%`

String that acts as prefix and suffix for config map values.

### keepPrevious
* ##### Type: `bool`
* ##### Required: `false`
* ##### Default: `false`

Whether Buckmate should clear files belonging to previous versions in the source destination.

### configMap
* ##### Type: `Record<string, string>`
* ##### Required: `false`
* ##### Default: `none`

Map of keys and values used to find and replace placeholders in deployment files.

### fileOptions
* ##### Type: `Record<string, { metadata: Record<string, string>, cacheControl: string }>`
* ##### Required: `false`
* ##### Default: `none`

Metadata and cache-control header settings for all files. Path should be relative to the root of the source directory. For cache-control values see https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9.