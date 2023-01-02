const Koa = require('koa');
const koaBodyParser = require('koa-bodyparser');

const useRouters = require('../routers');
const errorHandler = require('./error-handler');

const app = new Koa();

app.use(koaBodyParser());

useRouters(app);

app.on('error', errorHandler);

module.exports = app;
