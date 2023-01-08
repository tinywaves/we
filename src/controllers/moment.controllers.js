const {
  insertMoment,
  getMomentDetailById,
  getMomentListByLimited
} = require('../services/moment.services');
const { CREATE_MOMENT_ERROR } = require('../constants/error-types');

class MomentControllers {
  // create a new moment
  async createMoment(ctx) {
    const userId = ctx.user.id;
    const content = ctx.request.body.content;
    const result = await insertMoment(userId, content);

    if (result.fieldCount === 0) {
      ctx.body = { content };
    } else {
      ctx.app.emit('error', new Error(CREATE_MOMENT_ERROR), ctx);
    }
  }

  // get a moment detail
  async getMomentDetail(ctx) {
    const momentId = ctx.params.momentId;
    const [momentDetail] = await getMomentDetailById(momentId);

    ctx.body = momentDetail;
  }

  // get moment list
  async getMomentList(ctx) {
    const { offset, size } = ctx.query;
    const momentList = await getMomentListByLimited(offset, size);

    ctx.body = momentList;
  }
}

module.exports = new MomentControllers();
