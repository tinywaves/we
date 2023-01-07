const KoaRouter = require('koa-router');

const {
  verifyUser,
  handlePassword,
  verifyLogin,
  verifyAuth
} = require('../middlewares/user.middlewares');
const { createNewUser, login, authSuccess } = require('../controllers/user.controllers');

const userRouters = new KoaRouter({ prefix: '/user' });

userRouters.post('/register', verifyUser, handlePassword, createNewUser);

userRouters.post('/login', verifyLogin, login);

userRouters.get('/authorization', verifyAuth, authSuccess);

module.exports = userRouters;
