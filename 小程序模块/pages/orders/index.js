Page({
  data: {
    orderTabs: ['全部', '待支付', '进行中', '待审核'],
    activeTab: 0,
    orders: [
      {
        title: '极速上分 1小时',
        status: '待审核',
        amount: '¥88',
        time: '今天 14:20'
      },
      {
        title: '稳健冲分 3小时',
        status: '进行中',
        amount: '¥238',
        time: '今天 12:05'
      }
    ]
  },

  switchTab(e) {
    this.setData({
      activeTab: Number(e.currentTarget.dataset.index)
    })
  }
})
