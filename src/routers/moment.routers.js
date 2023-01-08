const KoaRouter = require('koa-router');

const { verifyAuth } = require('../middlewares/user.middlewares');
const {
  createMoment,
  getMomentDetail,
  getMomentList
} = require('../controllers/moment.controllers');

const momentRouters = new KoaRouter({ prefix: '/moment' });

momentRouters.post('/', verifyAuth, createMoment);

momentRouters.get('/:momentId', getMomentDetail);

momentRouters.get('/', getMomentList);

module.exports = momentRouters;
