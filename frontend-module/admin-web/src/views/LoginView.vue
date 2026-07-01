<template>
  <div class="login">
    <div class="card">
      <div class="title">后台登录</div>
      <div class="subtitle">Ovo Gaming 管理控制台</div>
      <a-input placeholder="管理员账号" style="margin-top: 24px" />
      <a-input-password placeholder="密码" style="margin-top: 12px" />
      <a-button type="primary" block style="margin-top: 20px">登录</a-button>
    </div>
  </div>
</template>

<style scoped>
.login {
  min-height: 100vh;
  display: grid;
  place-items: center;
  background: linear-gradient(135deg, #0f172a 0%, #1d4ed8 100%);
}

.card {
  width: 420px;
  padding: 32px;
  border-radius: 24px;
  background: rgba(255, 255, 255, 0.96);
  box-shadow: 0 24px 60px rgba(15, 23, 42, 0.25);
}

.title {
  font-size: 28px;
  font-weight: 800;
  color: #0f172a;
}

.subtitle {
  margin-top: 8px;
  color: #64748b;
}
</style>
