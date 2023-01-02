const { userServices } = require('../services');

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
    ctx.body = result;
  }
  /**
   * login
   * @param {*} ctx 
   * @param {*} next 
   */
  async login(ctx, next) {
    const { username, password } = ctx.request.body;

    ctx.body = `login success${username}`;
  }
}

module.exports = new UserControllers();
