${{ content_synopsis }} This image will run a [rootless](https://github.com/11notes/RTFM/blob/main/linux/container/image/rootless.md) sftp server based on openssh. Unlike other images, this image does not provide chroot jails, but it's intended to run as a single user with all the shares needed mounted in ```/home/%u```. Use an SSH proxy in front of this image when you need to expose multiple endpoints via a single entry point. You must provide secrets to use this image.

${{ content_uvp }} Good question! Because ...

${{ github:> [!IMPORTANT] }}
${{ github:> }}* ... this image runs [rootless](https://github.com/11notes/RTFM/blob/main/linux/container/image/rootless.md) as 1000:1000
${{ github:> }}* ... this image is auto updated to the latest version via CI/CD
${{ github:> }}* ... this image has a health check
${{ github:> }}* ... this image runs read-only
${{ github:> }}* ... this image is automatically scanned for CVEs before and after publishing
${{ github:> }}* ... this image is created via a secure and pinned CI/CD process
${{ github:> }}* ... this image is very small

If you value security, simplicity and optimizations to the extreme, then this image might be for you.

${{ content_compose }}

${{ content_defaults }}

${{ content_environment }}
| `SSH_USER` | username to access SSH server |  |
| `SSH_PASSWORD` | pasword to access SSH server | |
| `SSH_PASSWORD_FILE` *(optional)* | pasword file to access SSH server | |

${{ content_source }}

${{ content_parent }}

${{ content_built }}

${{ content_tips }}