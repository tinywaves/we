const {
  USERNAME_OR_PASSWORD_IS_REQUIRED,
  USER_ALREADY_EXIST,
  USER_IS_NOT_EXIST,
  PASSWORD_IS_INCORRECT,
  REGISTER_FAILED
} = require('../constants/error-types');

module.exports = (error, ctx) => {
  let status, errorMessage;

  switch (error.message) {
    case USERNAME_OR_PASSWORD_IS_REQUIRED:
      {
        status = 404; // bad request
        errorMessage = 'username or password should not be empty';
        break;
      }
    case USER_ALREADY_EXIST:
      {
        status = 409; // conflict
        errorMessage = 'user already exits, please reset your username input';
        break;
      }
    case USER_IS_NOT_EXIST:
      {
        status = 400;
        errorMessage = 'your request user is not exist, please register before login';
        break;
      }
    case PASSWORD_IS_INCORRECT:
      {
        status = 400;
        errorMessage = 'your input password is incorrect';
        break;
      }
    case REGISTER_FAILED:
      {
        status = 400;
        errorMessage = 'register failed, please retry later';
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
