const app = require('./app');
const { APP_PORT } = require('./app/config');

app.listen(APP_PORT, () => {
  console.log(`tinyRipple-hub server in ${APP_PORT}...`);
});
