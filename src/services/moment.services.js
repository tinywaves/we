const connections = require('../app/database');

class MomentServices {
  async insertMoment(userId, content) {
    const statement = 'insert into moment (content, user_id) values (?, ?);';
    const [result] = await connections.execute(statement, [content, userId]);

    return result;
  }

  async getMomentDetailById(momentId) {
    const statement = `
      select m.id as id,
             m.content as content,
             m.createAt createTime,
             m.updateAt updateTime,
             JSON_OBJECT('id', u.id, 'authName', u.username) author
      from moment m
      left join users u on u.id = m.user_id
      where m.id = ?;
    `;
    const [[result]] = await connections.execute(statement, [momentId]);

    return result;
  }
}

module.exports = new MomentServices;
