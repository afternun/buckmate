![buckmate-logo](docs/docs/assets/logo.png)

Visit [buckmate.org](https://buckmate.org) for more information!

## buckmate - made primarily to deploy static websites to AWS S3, however it can be used to:
* transfer files between **buckets**,
* transfer files between **servers** and **buckets**,
* **replace content** in transfered files according to **yaml configuration**.

1. Define your configuration using `yaml` files
2. Configure your environment with AWS credentials and region details
3. Run `buckmate apply` to swap placeholders and upload your files to desired location

## Testing

Currently only testing in this project is a funny shell test suite. To run them you will need your own AWS buckets. If you want to do it, you will figure it out...

1. Modify `example` directory to contain your bucket names
1. `cd e2e`
2. `sh e2h.sh`

## Contributing / Code of conduct

Feel free to open issues, or PRs. Be nice.