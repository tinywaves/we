const jwt = require('jsonwebtoken');

const { userServices } = require('../services');
const { PRIVATE_KEY } = require('../app/config');
const { REGISTER_FAILED } = require('../constants/error-types');

class UserControllers {
  /**
   * register
   * @param {*} ctx 
   * @param {*} next 
   */
  async createNewUser(ctx, next) {
    // get user form data
    const userFormData = ctx.request.body;
    // control database
    const result = await userServices.createNewUser(userFormData);
    // return a response
    if (result.fieldCount === 0) {
      ctx.body = { username: userFormData.username };
    } else {
      ctx.app.emit('error', new Error(REGISTER_FAILED), ctx);
    }
  }

  /**
   * login
   * @param {*} ctx 
   * @param {*} next 
   */
  async login(ctx, next) {
    const { id, username } = ctx.user;
    const token = jwt.sign({ id, username }, PRIVATE_KEY, {
      expiresIn: 24 * 60 * 60,
      algorithm: 'RS256'
    });

    ctx.body = {
      username,
      token
    };
  }
}

module.exports = new UserControllers();
