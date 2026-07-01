Page({
  data: {
    roles: [
      { name: '老板', active: true },
      { name: '打手', active: false },
      { name: '管理', active: false }
    ]
  },

  onRoleTap(e) {
    const { index } = e.currentTarget.dataset
    const roles = this.data.roles.map((role, i) => ({
      ...role,
      active: i === Number(index)
    }))
    this.setData({ roles })
  }
})
