// index.js
Page({
  data: {
    bossName: '凌晨电竞',
    announcement: '系统公告：当前仅开放三角洲行动，热门套餐支持指定打手与优惠选择。',
    currentBannerIndex: 0,
    banners: [
      {
        title: '三角洲行动热门开黑',
        subTitle: '后台可配置轮播图内容与跳转',
        background: 'linear-gradient(135deg, #18213f 0%, #26345f 100%)',
        jumpTarget: '/pages/packages/index'
      },
      {
        title: '指定打手更高效',
        subTitle: '支持老板指定熟悉的打手接单',
        background: 'linear-gradient(135deg, #3b1f32 0%, #6b2d4d 100%)',
        jumpTarget: '/pages/orders/index'
      },
      {
        title: '优惠活动实时更新',
        subTitle: '后台上架活动后首页自动同步展示',
        background: 'linear-gradient(135deg, #16363b 0%, #1f5a5f 100%)',
        jumpTarget: '/pages/profile/index'
      }
    ],
    quickEntries: [
      {
        title: '套餐页',
        desc: '查看三角洲行动热门套餐',
        key: 'packages'
      },
      {
        title: '订单页',
        desc: '跟踪待支付、进行中和待审核订单',
        key: 'orders'
      },
      {
        title: '我的页',
        desc: '查看账号、优惠与身份切换',
        key: 'profile'
      }
    ],
    stats: [
      { label: '今日订单', value: '12' },
      { label: '待处理', value: '4' },
      { label: '累计优惠', value: '¥268' }
    ],
    allPackages: [
      {
        title: '极速上分 1小时',
        price: '¥88',
        originalPrice: '¥108',
        badge: '热门',
        isHot: true,
        tags: ['支持指定打手', '优先派单', '结单审核'],
        highlight: '适合想快速开局的老板'
      },
      {
        title: '稳健冲分 3小时',
        price: '¥238',
        originalPrice: '¥268',
        badge: '常规',
        isHot: false,
        tags: ['稳定陪打', '可选优惠', '高性价比'],
        highlight: '适合连续上分和长时间陪玩'
      },
      {
        title: '周末包场 5小时',
        price: '¥358',
        originalPrice: '¥399',
        badge: '热门',
        isHot: true,
        tags: ['更强优惠', '优先响应', '适合组队'],
        highlight: '适合约好时间集中开黑'
      }
    ],
    hotPackages: []
  },

  onLoad() {
    this.setData({
      hotPackages: this.data.allPackages.filter((item) => item.isHot)
    })
  },

  onBannerChange(e) {
    this.setData({
      currentBannerIndex: e.detail.current
    })
  },

  onBannerTap(e) {
    const { target } = e.currentTarget.dataset
    if (!target) return
    wx.switchTab({
      url: target
    })
  },

  onQuickTap(e) {
    const { key } = e.currentTarget.dataset
    const urlMap = {
      packages: '/pages/packages/index',
      orders: '/pages/orders/index',
      profile: '/pages/profile/index'
    }
    wx.switchTab({
      url: urlMap[key]
    })
  },

  onPackageTap() {
    wx.switchTab({
      url: '/pages/packages/index'
    })
  }
})
