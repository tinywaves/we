const mysql = require('mysql2');

const {
  MYSQL_HOST,
  MYSQL_PORT,
  MYSQL_DATABASE,
  MYSQL_USER,
  MYSQL_PASSWORD
} = require('./config');

const connections = mysql.createPool({
  host: MYSQL_HOST,
  port: MYSQL_PORT,
  database: MYSQL_DATABASE,
  user: MYSQL_USER,
  password: MYSQL_PASSWORD
});

connections.getConnection((error, connection) => {
  connection.connect((errors) => {
    if (errors) {
      console.log('get connection fails', errors);
    } else {
      console.log('get connection succeed');
    }
  });
});

module.exports = connections.promise();
