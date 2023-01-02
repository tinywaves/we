const {
  USERNAME_OR_PASSWORD_IS_REQUIRED,
  USER_ALREADY_EXIST
} = require('../constants/error-types');
const { userServices } = require('../services');
const { md5Transform } = require('../utils');

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
  const result = await userServices.getUserByUsername(username);

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

module.exports = { verifyUser, handlePassword };
