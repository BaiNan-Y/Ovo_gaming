Page({
  data: {
    packages: [
      {
        title: '极速上分 1小时',
        price: '¥88',
        desc: '适合快速开局，支持指定打手和优先派单。'
      },
      {
        title: '稳健冲分 3小时',
        price: '¥238',
        desc: '更适合连续陪打，兼顾优惠和性价比。'
      },
      {
        title: '周末包场 5小时',
        price: '¥358',
        desc: '适合提前约档期，集中连打更划算。'
      }
    ]
  },

  onOrderTap() {
    wx.switchTab({
      url: '/pages/index/index'
    })
  }
})
