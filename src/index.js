const Koa = require('koa');

const app = new Koa();

app.listen(9999, () => {
  console.log('tinyRipple-hub server...');
});
