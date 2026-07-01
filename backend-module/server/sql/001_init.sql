CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  openid TEXT UNIQUE NOT NULL,
  unionid TEXT,
  username TEXT,
  nickname TEXT NOT NULL DEFAULT '',
  avatar_url TEXT NOT NULL DEFAULT '',
  phone TEXT NOT NULL DEFAULT '',
  status TEXT NOT NULL DEFAULT 'normal',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_roles (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  role_type TEXT NOT NULL,
  role_status TEXT NOT NULL DEFAULT 'active',
  is_default BOOLEAN NOT NULL DEFAULT FALSE,
  apply_status TEXT NOT NULL DEFAULT 'none',
  apply_remark TEXT NOT NULL DEFAULT '',
  reviewed_by BIGINT,
  reviewed_at TIMESTAMPTZ,
  game_main TEXT NOT NULL DEFAULT '',
  rank_info TEXT NOT NULL DEFAULT '',
  service_tags TEXT[] NOT NULL DEFAULT '{}',
  service_desc TEXT NOT NULL DEFAULT '',
  accept_order_status BOOLEAN NOT NULL DEFAULT TRUE,
  online_status BOOLEAN NOT NULL DEFAULT FALSE,
  service_score NUMERIC(4,2) NOT NULL DEFAULT 0,
  completed_count INT NOT NULL DEFAULT 0,
  balance_available NUMERIC(12,2) NOT NULL DEFAULT 0,
  balance_frozen NUMERIC(12,2) NOT NULL DEFAULT 0,
  balance_total NUMERIC(12,2) NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS banners (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  subtitle TEXT NOT NULL DEFAULT '',
  image_url TEXT NOT NULL DEFAULT '',
  jump_type TEXT NOT NULL DEFAULT 'tab',
  jump_target TEXT NOT NULL DEFAULT '',
  status TEXT NOT NULL DEFAULT 'on',
  sort INT NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS notices (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  notice_type TEXT NOT NULL DEFAULT 'system',
  status TEXT NOT NULL DEFAULT 'on',
  sort INT NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS packages (
  id BIGSERIAL PRIMARY KEY,
  game_code TEXT NOT NULL DEFAULT 'delta_force',
  title TEXT NOT NULL,
  price NUMERIC(12,2) NOT NULL,
  original_price NUMERIC(12,2) NOT NULL DEFAULT 0,
  duration_minutes INT NOT NULL DEFAULT 60,
  description TEXT NOT NULL DEFAULT '',
  cover_image TEXT NOT NULL DEFAULT '',
  is_hot BOOLEAN NOT NULL DEFAULT FALSE,
  status TEXT NOT NULL DEFAULT 'on',
  sort INT NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS discounts (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  discount_type TEXT NOT NULL DEFAULT 'activity',
  discount_mode TEXT NOT NULL DEFAULT 'amount',
  discount_value NUMERIC(12,2) NOT NULL DEFAULT 0,
  condition_min_amount NUMERIC(12,2) NOT NULL DEFAULT 0,
  scope_type TEXT NOT NULL DEFAULT 'all',
  scope_ids BIGINT[] NOT NULL DEFAULT '{}',
  valid_start_at TIMESTAMPTZ,
  valid_end_at TIMESTAMPTZ,
  status TEXT NOT NULL DEFAULT 'on',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS orders (
  id BIGSERIAL PRIMARY KEY,
  order_no TEXT UNIQUE NOT NULL,
  boss_user_id BIGINT NOT NULL REFERENCES users(id),
  package_id BIGINT NOT NULL REFERENCES packages(id),
  package_snapshot JSONB NOT NULL DEFAULT '{}'::jsonb,
  discount_id BIGINT,
  discount_snapshot JSONB NOT NULL DEFAULT '{}'::jsonb,
  specified_player_id BIGINT,
  assigned_player_id BIGINT,
  order_amount NUMERIC(12,2) NOT NULL DEFAULT 0,
  discount_amount NUMERIC(12,2) NOT NULL DEFAULT 0,
  pay_amount NUMERIC(12,2) NOT NULL DEFAULT 0,
  pay_status TEXT NOT NULL DEFAULT 'unpaid',
  order_status TEXT NOT NULL DEFAULT 'pending_pay',
  dispatch_type TEXT NOT NULL DEFAULT 'public',
  paid_at TIMESTAMPTZ,
  dispatch_at TIMESTAMPTZ,
  accept_at TIMESTAMPTZ,
  start_at TIMESTAMPTZ,
  finish_at TIMESTAMPTZ,
  boss_confirm_at TIMESTAMPTZ,
  review_at TIMESTAMPTZ,
  cancel_reason TEXT NOT NULL DEFAULT '',
  abnormal_reason TEXT NOT NULL DEFAULT '',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS order_assignments (
  id BIGSERIAL PRIMARY KEY,
  order_id BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
  player_id BIGINT NOT NULL REFERENCES users(id),
  assign_type TEXT NOT NULL DEFAULT 'public',
  assign_status TEXT NOT NULL DEFAULT 'pending',
  read_at TIMESTAMPTZ,
  accepted_at TIMESTAMPTZ,
  rejected_at TIMESTAMPTZ,
  expire_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS order_logs (
  id BIGSERIAL PRIMARY KEY,
  order_id BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
  action_type TEXT NOT NULL,
  operator_user_id BIGINT,
  operator_role_type TEXT NOT NULL DEFAULT '',
  before_status TEXT NOT NULL DEFAULT '',
  after_status TEXT NOT NULL DEFAULT '',
  note TEXT NOT NULL DEFAULT '',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS withdrawals (
  id BIGSERIAL PRIMARY KEY,
  player_user_id BIGINT NOT NULL REFERENCES users(id),
  amount NUMERIC(12,2) NOT NULL,
  withdraw_method TEXT NOT NULL DEFAULT 'wechat',
  account_info JSONB NOT NULL DEFAULT '{}'::jsonb,
  withdraw_status TEXT NOT NULL DEFAULT 'pending',
  audit_admin_id BIGINT,
  audit_remark TEXT NOT NULL DEFAULT '',
  applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  audited_at TIMESTAMPTZ,
  paid_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS reviews (
  id BIGSERIAL PRIMARY KEY,
  order_id BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
  boss_user_id BIGINT NOT NULL REFERENCES users(id),
  player_user_id BIGINT NOT NULL REFERENCES users(id),
  score INT NOT NULL DEFAULT 5,
  content TEXT NOT NULL DEFAULT '',
  status TEXT NOT NULL DEFAULT 'visible',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS audit_logs (
  id BIGSERIAL PRIMARY KEY,
  operator_user_id BIGINT,
  operator_role_type TEXT NOT NULL DEFAULT '',
  action_type TEXT NOT NULL,
  target_type TEXT NOT NULL DEFAULT '',
  target_id BIGINT,
  note TEXT NOT NULL DEFAULT '',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_users_openid ON users(openid);
CREATE INDEX IF NOT EXISTS idx_user_roles_user_id ON user_roles(user_id);
CREATE INDEX IF NOT EXISTS idx_packages_hot ON packages(is_hot, sort DESC);
CREATE INDEX IF NOT EXISTS idx_banners_status_sort ON banners(status, sort DESC);
CREATE INDEX IF NOT EXISTS idx_notices_status_sort ON notices(status, sort DESC);
CREATE INDEX IF NOT EXISTS idx_orders_boss_user_id ON orders(boss_user_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(order_status);
CREATE INDEX IF NOT EXISTS idx_order_assignments_order_id ON order_assignments(order_id);
CREATE INDEX IF NOT EXISTS idx_withdrawals_player_user_id ON withdrawals(player_user_id);
