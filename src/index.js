const app = require('./app');
const { APP_PORT } = require('./app/config');
require('./app/database');

app.listen(APP_PORT, () => {
  console.log(`tinyRipple-hub server in ${APP_PORT}...`);
});
