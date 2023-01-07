const connections = require('../app/database');

class MomentServices {
  async insertMoment(userId, content) {
    const statement = 'insert into moment (content, user_id) values (?, ?);';
    const [result] = await connections.execute(statement, [content, userId]);

    return result;
  }
}

module.exports = new MomentServices;
