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

## interfaces

- <samp><b>user register</b></samp> `POST /user/register`
```json
request:
{
  "username": "tinyRipple",
  "password": "1234567890"
}
```
