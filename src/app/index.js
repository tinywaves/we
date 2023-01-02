const Koa = require('koa');
const koaBodyParser = require('koa-bodyparser');

const { userRouters } = require('../routers');
const errorHandler = require('./error-handler');

const app = new Koa();

app.use(koaBodyParser());
app.use(userRouters.routes());
app.use(userRouters.allowedMethods());

app.on('error', errorHandler);

module.exports = app;
