INSERT INTO banners (title, subtitle, image_url, jump_type, jump_target, status, sort)
VALUES
  ('三角洲行动热门开黑', '后台可配置轮播图内容与跳转', '', 'tab', '/pages/packages/index', 'on', 3),
  ('指定打手更高效', '支持老板指定熟悉的打手接单', '', 'tab', '/pages/orders/index', 'on', 2),
  ('优惠活动实时更新', '后台上架活动后首页自动同步展示', '', 'tab', '/pages/profile/index', 'on', 1)
ON CONFLICT DO NOTHING;

INSERT INTO notices (title, content, notice_type, status, sort)
VALUES
  ('平台公告', '当前仅开放三角洲行动，热门套餐支持指定打手与优惠选择。', 'system', 'on', 1)
ON CONFLICT DO NOTHING;

INSERT INTO packages (game_code, title, price, original_price, duration_minutes, description, cover_image, is_hot, status, sort)
VALUES
  ('delta_force', '极速上分 1小时', 88, 108, 60, '适合想快速开局的老板', '', TRUE, 'on', 3),
  ('delta_force', '稳健冲分 3小时', 238, 268, 180, '适合连续上分和长时间陪玩', '', FALSE, 'on', 2),
  ('delta_force', '周末包场 5小时', 358, 399, 300, '适合约好时间集中开黑', '', TRUE, 'on', 1)
ON CONFLICT DO NOTHING;

INSERT INTO discounts (title, discount_type, discount_mode, discount_value, condition_min_amount, scope_type, scope_ids, status)
VALUES
  ('新客立减20', 'activity', 'amount', 20, 0, 'all', '{}', 'on')
ON CONFLICT DO NOTHING;

