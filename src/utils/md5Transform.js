const crypto = require('crypto');

module.exports = (password) => {
  const md5 = crypto.createHash('md5');

  return md5.update(password).digest('hex');
};
