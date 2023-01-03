# tinyRipple-hub

A backend project about a forum developed using nodejs/koa.

## project directory

You can install tree tool by homebrew.

```shell
brew install tree
```

Display project directory.

```shell
tree -I '*svn|*node_module*'
```

## .env configuration

.env should not be managed by git, but since this project is a personal project, it is convenient to start development and replay repeatedly, so add this file to git management.

```shell
APP_PORT=?

MYSQL_HOST=?
MYSQL_PORT=?
MYSQL_DATABASE=?
MYSQL_USER=?
MYSQL_PASSWORD=?
```

## api documentation

- <samp><b><a href="https://tinyripple.notion.site/tinyRipple-hub-99cfe913728e49f9b272d40e048f711a">Click Here</a></b></samp>

## generate private-key and public-key

- step 1

```shell
mkdir key
cd key
```

- step 2

Go to OpenSSL shell.

```shell
openssl
```

- step 3

Generate private-key. RSA key size must be 2048 bits or greater in `jsonwebtoken v9`.

```shell
OpenSSL> genrsa -out private.key 2048
```

- step 4

Generate public-key.

```shell
OpenSSL> rsa -in private.key -pubout -out public.key
```
