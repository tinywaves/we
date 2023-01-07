const jwt = require('jsonwebtoken');

const { createNewUser } = require('../services/user.services');
const { PRIVATE_KEY } = require('../app/config');
const { REGISTER_FAILED } = require('../constants/error-types');

class UserControllers {
  // register
  async createNewUser(ctx) {
    // get user form data
    const userFormData = ctx.request.body;
    // control database
    const result = await createNewUser(userFormData);
    // return a response
    if (result.fieldCount === 0) {
      ctx.body = { username: userFormData.username };
    } else {
      ctx.app.emit('error', new Error(REGISTER_FAILED), ctx);
    }
  }

  // login
  async login(ctx) {
    const { id, username } = ctx.user;
    const token = jwt.sign({ id, username }, PRIVATE_KEY, {
      expiresIn: 30 * 24 * 60 * 60,
      algorithm: 'RS256'
    });

    ctx.body = {
      username,
      token
    };
  }

  // authenticate success
  async authSuccess(ctx) {
    ctx.body = {
      authorized: true
    };
  }
}

module.exports = new UserControllers();
