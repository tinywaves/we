const KoaRouter = require('koa-router');

const { verifyAuth } = require('../middlewares/user.middlewares');
const {
  createMoment,
  getMomentDetail
} = require('../controllers/moment.controllers');

const momentRouters = new KoaRouter({ prefix: '/moment' });

momentRouters.post('/', verifyAuth, createMoment);

momentRouters.get('/:momentId', getMomentDetail);

module.exports = momentRouters;
