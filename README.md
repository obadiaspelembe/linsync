# LynodeSyncronzier

LynSync is tool for syncronize Linode Object Storage files in your system directory.

## Usage:

### Pushing to bucket`

To upload  files to a bucket using lynsinc use the push command

`lynsync push [filename]`

### Pull to directory

To download files to a bucket using lynsinc use the pull command

`lynsync pull [filename]`


## Configuration

LynSync needs to be configured with the following environment variables:

| Name | Description |
| ----------- | ----------- |
| `LINODE_ACCESS_KEY_ID` | Object Storage Access Key |
| `LINODE_SECRET_ACCESS_KEY` | Object Storage Secret Acess Key |
| `LINODE_DEFAULT_REGION` | Object Storage Cluster Id or region |
| `LINODE_BUCKET_NAME` | Object Storage Access Key |

 