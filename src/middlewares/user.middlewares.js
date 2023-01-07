const jwt = require('jsonwebtoken');

const {
  USERNAME_OR_PASSWORD_IS_REQUIRED,
  USER_ALREADY_EXIST,
  USER_IS_NOT_EXIST,
  PASSWORD_IS_INCORRECT,
  UNAUTHORIZED_TOKEN,
  AUTHORIZATION_MISSING
} = require('../constants/error-types');
const { getUserByUsername } = require('../services/user.services');
const { md5Transform } = require('../utils');
const { PUBLIC_KEY } = require('../app/config');

const verifyUser = async (ctx, next) => {
  const { username, password } = ctx.request.body;
  // username or password should not be empty
  if (
    !username ||
    !password ||
    username.length === 0 ||
    password.length === 0
  ) {
    return ctx.app.emit('error', new Error(USERNAME_OR_PASSWORD_IS_REQUIRED), ctx);
  }

  // username should be unique
  const result = await getUserByUsername(username);

  if (result.length) {
    return ctx.app.emit('error', new Error(USER_ALREADY_EXIST), ctx);
  }

  await next();
};

const handlePassword = async (ctx, next) => {
  const { password } = ctx.request.body;

  ctx.request.body.password = md5Transform(password);

  await next();
};

const verifyLogin = async (ctx, next) => {
  const { username, password } = ctx.request.body;
  // username or password should not be empty
  if (
    !username ||
    !password ||
    username.length === 0 ||
    password.length === 0
  ) {
    return ctx.app.emit('error', new Error(USERNAME_OR_PASSWORD_IS_REQUIRED), ctx);
  }

  // judge if user is existing
  const result = await getUserByUsername(username);
  const user = result[0];

  if (!user) {
    return ctx.app.emit('error', new Error(USER_IS_NOT_EXIST), ctx);
  }

  // judge if password is correct
  const md5TransformPassword = md5Transform(password);

  if (md5TransformPassword !== user?.password) {
    return ctx.app.emit('error', new Error(PASSWORD_IS_INCORRECT), ctx);
  }

  ctx.user = user;

  await next();
};

const verifyAuth = async (ctx, next) => {
  // get bearer token
  const authorization = ctx.headers.authorization;

  if (authorization) {
    // get token
    const token = authorization.replace('Bearer ', '');

    // verify token
    try {
      const user = jwt.verify(token, PUBLIC_KEY, {
        algorithms: ['RS256']
      });

      ctx.user = user;
    } catch (error) {
      return ctx.app.emit('error', new Error(UNAUTHORIZED_TOKEN), ctx);
    }

    await next();
  } else {
    return ctx.app.emit('error', new Error(AUTHORIZATION_MISSING), ctx);
  }
};

module.exports = { verifyUser, handlePassword, verifyLogin, verifyAuth };
