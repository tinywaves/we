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
}

module.exports = new UserControllers();
