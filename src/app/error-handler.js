const {
  USERNAME_OR_PASSWORD_IS_REQUIRED,
  USER_ALREADY_EXIST
} = require('../constants/error-types');

module.exports = (error, ctx) => {
  let status, errorMessage;

  switch (error.message) {
    case USERNAME_OR_PASSWORD_IS_REQUIRED:
      {
        status = 404;
        errorMessage = 'username or password should not be empty';
        break;
      }
    case USER_ALREADY_EXIST:
      {
        status = 409;
        errorMessage = 'user already exits, please reset your username input';
        break;
      }
    default:
      {
        status = 500;
        errorMessage = 'system error';
        break;
      }
  }

  ctx.status = status;
  ctx.body = {
    errorCode: status,
    errorMessage
  };
};
