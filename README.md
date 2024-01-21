# buckmate

![buckmate-logo](docs/docs/assets/logo.png)

Visit [buckmate.org](https://buckmate.org) for more information!

## Made primarily to deploy static websites to AWS S3, however it can be used to:
* transfer files between **buckets**,
* transfer files between **servers** and **buckets**,
* **replace content** in transfered files according to **yaml configuration**.

1. Define your configuration using `yaml` files
2. Configure your environment with AWS credentials and region details
3. Run `buckmate apply` to swap placeholders and upload your files to desired location