const connections = require('../app/database');

class UserServices {
  async createNewUser({ username, password }) {
    const statement = `insert into users (username, password) values (?, ?);`;
    const result = await connections.execute(statement, [username, password]);

    return result[0];
  }

  async getUserByUsername(username) {
    const statement = `select * from users where username = ?;`;
    const result = await connections.execute(statement, [username]);

    return result[0];
  }
}

module.exports = new UserServices();
