-- =====================================================
-- TC Gaming - 投注详情存储架构
-- =====================================================
-- 此架构存储Live和RNG/FISH游戏的投注记录详情
-- 基于TC Gaming API投注详情接口的分析

-- Live游戏投注详情主表
CREATE TABLE game_live_bet_details (
    -- 主键和基本信息
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    
    -- 玩家信息 (来自BetDetail结构体)
    username VARCHAR(100) NOT NULL COMMENT '游戏账号的登录名',
    
    -- 投注金额信息
    bet_amount DECIMAL(15,2) NOT NULL COMMENT '投注金额',
    valid_bet_amount DECIMAL(15,2) NOT NULL COMMENT '有效投注金额',
    win_amount DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT '赢金额',
    net_pnl DECIMAL(15,2) NOT NULL COMMENT '净输赢 (正数为赢，负数为输)',
    
    -- 货币和游戏信息
    currency VARCHAR(10) NOT NULL COMMENT '币别 (例如: CNY, USD)',
    game_code VARCHAR(50) NOT NULL COMMENT '游戏代码',
    product_type INT NOT NULL COMMENT '产品类别 (191=Live游戏)',
    game_category VARCHAR(20) NOT NULL DEFAULT 'LIVE' COMMENT '游戏类别',
    
    -- 订单和会话信息
    bet_order_no VARCHAR(200) NOT NULL UNIQUE COMMENT '投注订单编号',
    session_id VARCHAR(100) NOT NULL COMMENT '会话标识',
    
    -- 时间信息
    bet_time TIMESTAMP NOT NULL COMMENT '投注时间',
    transaction_time TIMESTAMP NOT NULL COMMENT '交易时间',
    
    -- 附加详情 (JSON格式存储)
    additional_details JSON NULL COMMENT '产品追加投注详细信息',
    
    -- API响应元数据
    api_status INT NOT NULL DEFAULT 0 COMMENT 'API响应状态码',
    api_error_desc TEXT NULL COMMENT 'API错误描述',
    
    -- 数据管理
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    
    -- 性能索引
    INDEX idx_username (username),
    INDEX idx_bet_order_no (bet_order_no),
    INDEX idx_session_id (session_id),
    INDEX idx_game_code (game_code),
    INDEX idx_product_type (product_type),
    INDEX idx_bet_time (bet_time),
    INDEX idx_transaction_time (transaction_time),
    INDEX idx_created_at (created_at),
    INDEX idx_username_bet_time (username, bet_time),
    INDEX idx_game_code_bet_time (game_code, bet_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='存储Live游戏的投注详情记录';

-- RNG/FISH游戏投注详情主表
CREATE TABLE game_rng_bet_details (
    -- 主键和基本信息
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    
    -- 玩家信息 (来自RNGBetDetail结构体)
    username VARCHAR(100) NOT NULL COMMENT '游戏账号的登录名',
    
    -- 投注金额信息
    bet_amount DECIMAL(15,2) NOT NULL COMMENT '投注金额',
    valid_bet_amount DECIMAL(15,2) NOT NULL COMMENT '有效投注金额',
    win_amount DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT '赢金额',
    net_pnl DECIMAL(15,2) NOT NULL COMMENT '净输赢 (正数为赢，负数为输)',
    
    -- 货币和游戏信息
    currency VARCHAR(10) NOT NULL COMMENT '币别 (例如: CNY, USD)',
    game_code VARCHAR(50) NOT NULL COMMENT '游戏代码',
    game_name VARCHAR(200) NOT NULL COMMENT '游戏名称 (RNG/FISH特有字段)',
    product_type INT NOT NULL COMMENT '产品类别 (16=RNG, 其他=FISH)',
    game_category VARCHAR(20) NOT NULL COMMENT '游戏类别 (RNG, FISH)',
    
    -- 订单和会话信息
    bet_order_no VARCHAR(200) NOT NULL UNIQUE COMMENT '投注订单编号',
    session_id VARCHAR(100) NOT NULL COMMENT '会话标识',
    
    -- 时间信息
    bet_time TIMESTAMP NOT NULL COMMENT '投注时间',
    transaction_time TIMESTAMP NOT NULL COMMENT '交易时间',
    
    -- 附加详情 (JSON格式存储，包含gamehall, gamePlat, status等)
    additional_details JSON NULL COMMENT '产品追加投注详细信息 (gamehall, gamePlat, status, createTime, endRoundTime, balance)',
    
    -- API响应元数据
    api_status INT NOT NULL DEFAULT 0 COMMENT 'API响应状态码',
    api_error_desc TEXT NULL COMMENT 'API错误描述',
    
    -- 数据管理
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    
    -- 性能索引
    INDEX idx_username (username),
    INDEX idx_bet_order_no (bet_order_no),
    INDEX idx_session_id (session_id),
    INDEX idx_game_code (game_code),
    INDEX idx_game_name (game_name),
    INDEX idx_product_type (product_type),
    INDEX idx_game_category (game_category),
    INDEX idx_bet_time (bet_time),
    INDEX idx_transaction_time (transaction_time),
    INDEX idx_created_at (created_at),
    INDEX idx_username_bet_time (username, bet_time),
    INDEX idx_game_code_bet_time (game_code, bet_time),
    INDEX idx_game_category_bet_time (game_category, bet_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='存储RNG/FISH游戏的投注详情记录';

-- 投注详情API调用记录表
CREATE TABLE game_bet_details_api_calls (
    -- 主键和基本信息
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,

    -- API调用信息
    api_method VARCHAR(20) NOT NULL COMMENT 'API方法 (lbdm=Live, bdm=RNG)',
    username VARCHAR(100) NOT NULL COMMENT '查询的玩家用户名',
    start_date TIMESTAMP NOT NULL COMMENT '查询开始日期',
    end_date TIMESTAMP NOT NULL COMMENT '查询结束日期',
    page_number INT NOT NULL DEFAULT 1 COMMENT '请求的页码',

    -- API响应分页信息 (来自PageInfo结构体)
    total_pages INT NOT NULL DEFAULT 0 COMMENT '总页数',
    current_page INT NOT NULL DEFAULT 1 COMMENT '当前页码',
    total_count INT NOT NULL DEFAULT 0 COMMENT '总记录数',

    -- API响应状态
    api_status INT NOT NULL COMMENT 'API响应状态码 (0=成功)',
    api_error_desc TEXT NULL COMMENT 'API错误描述',
    response_time_ms INT NULL COMMENT 'API响应时间(毫秒)',

    -- 数据处理状态
    records_processed INT NOT NULL DEFAULT 0 COMMENT '成功处理的记录数',
    processing_status ENUM('PENDING', 'SUCCESS', 'PARTIAL', 'FAILED') NOT NULL DEFAULT 'PENDING' COMMENT '处理状态',
    processing_error TEXT NULL COMMENT '处理错误信息',

    -- 时间戳
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '调用时间',
    completed_at TIMESTAMP NULL COMMENT '处理完成时间',

    -- 性能索引
    INDEX idx_api_method (api_method),
    INDEX idx_username (username),
    INDEX idx_start_date (start_date),
    INDEX idx_end_date (end_date),
    INDEX idx_api_status (api_status),
    INDEX idx_processing_status (processing_status),
    INDEX idx_created_at (created_at),
    INDEX idx_username_date_range (username, start_date, end_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='投注详情API调用和分页信息记录';

-- 玩家投注汇总统计表
CREATE TABLE game_player_bet_summary (
    -- 主键和基本信息
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,

    -- 玩家和时间维度
    username VARCHAR(100) NOT NULL COMMENT '玩家用户名',
    summary_date DATE NOT NULL COMMENT '汇总日期',
    game_category VARCHAR(20) NOT NULL COMMENT '游戏类别 (LIVE, RNG, FISH)',

    -- Live游戏汇总
    live_bet_count INT NOT NULL DEFAULT 0 COMMENT 'Live游戏投注次数',
    live_total_bet_amount DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT 'Live游戏总投注金额',
    live_total_valid_bet DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT 'Live游戏总有效投注',
    live_total_win_amount DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT 'Live游戏总赢金额',
    live_total_net_pnl DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT 'Live游戏总净输赢',

    -- RNG游戏汇总
    rng_bet_count INT NOT NULL DEFAULT 0 COMMENT 'RNG游戏投注次数',
    rng_total_bet_amount DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT 'RNG游戏总投注金额',
    rng_total_valid_bet DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT 'RNG游戏总有效投注',
    rng_total_win_amount DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT 'RNG游戏总赢金额',
    rng_total_net_pnl DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT 'RNG游戏总净输赢',

    -- 整体汇总
    total_bet_count INT NOT NULL DEFAULT 0 COMMENT '总投注次数',
    total_bet_amount DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT '总投注金额',
    total_valid_bet DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT '总有效投注',
    total_win_amount DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT '总赢金额',
    total_net_pnl DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT '总净输赢',

    -- 数据管理
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',

    -- 唯一约束和索引
    UNIQUE KEY uk_username_date_category (username, summary_date, game_category),
    INDEX idx_username (username),
    INDEX idx_summary_date (summary_date),
    INDEX idx_game_category (game_category),
    INDEX idx_username_date (username, summary_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='玩家每日投注汇总统计表';

-- =====================================================
-- 视图定义
-- =====================================================

-- 统一投注详情视图 (合并Live和RNG数据)
CREATE VIEW unified_bet_details AS
SELECT
    'LIVE' as bet_type,
    id,
    username,
    bet_amount,
    valid_bet_amount,
    win_amount,
    net_pnl,
    currency,
    game_code,
    NULL as game_name,  -- Live游戏没有game_name字段
    product_type,
    game_category,
    bet_order_no,
    session_id,
    bet_time,
    transaction_time,
    additional_details,
    created_at
FROM game_live_bet_details
UNION ALL
SELECT
    'RNG' as bet_type,
    id,
    username,
    bet_amount,
    valid_bet_amount,
    win_amount,
    net_pnl,
    currency,
    game_code,
    game_name,
    product_type,
    game_category,
    bet_order_no,
    session_id,
    bet_time,
    transaction_time,
    additional_details,
    created_at
FROM game_rng_bet_details;

-- 玩家投注统计视图
CREATE VIEW player_bet_stats AS
SELECT
    username,
    bet_type,
    COUNT(*) as bet_count,
    SUM(bet_amount) as total_bet_amount,
    SUM(valid_bet_amount) as total_valid_bet,
    SUM(win_amount) as total_win_amount,
    SUM(net_pnl) as total_net_pnl,
    AVG(bet_amount) as avg_bet_amount,
    MIN(bet_time) as first_bet_time,
    MAX(bet_time) as last_bet_time,
    COUNT(DISTINCT game_code) as unique_games_played
FROM unified_bet_details
GROUP BY username, bet_type;

-- 游戏热度统计视图
CREATE VIEW game_popularity_stats AS
SELECT
    bet_type,
    game_code,
    game_name,
    game_category,
    COUNT(*) as total_bets,
    COUNT(DISTINCT username) as unique_players,
    SUM(bet_amount) as total_bet_amount,
    SUM(win_amount) as total_win_amount,
    SUM(net_pnl) as house_edge,
    AVG(bet_amount) as avg_bet_amount
FROM unified_bet_details
GROUP BY bet_type, game_code, game_name, game_category
ORDER BY total_bets DESC;

-- =====================================================
-- 存储过程
-- =====================================================

-- 更新玩家每日汇总统计
DELIMITER //
CREATE PROCEDURE UpdatePlayerDailySummary(
    IN p_username VARCHAR(100),
    IN p_summary_date DATE
)
BEGIN
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;
        RESIGNAL;
    END;

    START TRANSACTION;

    -- 删除现有的汇总记录
    DELETE FROM player_bet_summary
    WHERE username = p_username AND summary_date = p_summary_date;

    -- 插入Live游戏汇总
    INSERT INTO player_bet_summary (
        username, summary_date, game_category,
        live_bet_count, live_total_bet_amount, live_total_valid_bet,
        live_total_win_amount, live_total_net_pnl,
        total_bet_count, total_bet_amount, total_valid_bet,
        total_win_amount, total_net_pnl
    )
    SELECT
        p_username,
        p_summary_date,
        'LIVE',
        COUNT(*),
        COALESCE(SUM(bet_amount), 0),
        COALESCE(SUM(valid_bet_amount), 0),
        COALESCE(SUM(win_amount), 0),
        COALESCE(SUM(net_pnl), 0),
        COUNT(*),
        COALESCE(SUM(bet_amount), 0),
        COALESCE(SUM(valid_bet_amount), 0),
        COALESCE(SUM(win_amount), 0),
        COALESCE(SUM(net_pnl), 0)
    FROM game_live_bet_details
    WHERE username = p_username
      AND DATE(bet_time) = p_summary_date
    HAVING COUNT(*) > 0;

    -- 插入RNG游戏汇总
    INSERT INTO player_bet_summary (
        username, summary_date, game_category,
        rng_bet_count, rng_total_bet_amount, rng_total_valid_bet,
        rng_total_win_amount, rng_total_net_pnl,
        total_bet_count, total_bet_amount, total_valid_bet,
        total_win_amount, total_net_pnl
    )
    SELECT
        p_username,
        p_summary_date,
        'RNG',
        COUNT(*),
        COALESCE(SUM(bet_amount), 0),
        COALESCE(SUM(valid_bet_amount), 0),
        COALESCE(SUM(win_amount), 0),
        COALESCE(SUM(net_pnl), 0),
        COUNT(*),
        COALESCE(SUM(bet_amount), 0),
        COALESCE(SUM(valid_bet_amount), 0),
        COALESCE(SUM(win_amount), 0),
        COALESCE(SUM(net_pnl), 0)
    FROM game_rng_bet_details
    WHERE username = p_username
      AND DATE(bet_time) = p_summary_date
    HAVING COUNT(*) > 0;

    COMMIT;
END //
DELIMITER ;

-- 获取玩家指定日期范围的投注统计
DELIMITER //
CREATE PROCEDURE GetPlayerBetStats(
    IN p_username VARCHAR(100),
    IN p_start_date DATE,
    IN p_end_date DATE
)
BEGIN
    SELECT
        'LIVE' as game_type,
        COUNT(*) as bet_count,
        SUM(bet_amount) as total_bet_amount,
        SUM(valid_bet_amount) as total_valid_bet,
        SUM(win_amount) as total_win_amount,
        SUM(net_pnl) as total_net_pnl,
        AVG(bet_amount) as avg_bet_amount
    FROM game_live_bet_details
    WHERE username = p_username
      AND DATE(bet_time) BETWEEN p_start_date AND p_end_date

    UNION ALL

    SELECT
        'RNG' as game_type,
        COUNT(*) as bet_count,
        SUM(bet_amount) as total_bet_amount,
        SUM(valid_bet_amount) as total_valid_bet,
        SUM(win_amount) as total_win_amount,
        SUM(net_pnl) as total_net_pnl,
        AVG(bet_amount) as avg_bet_amount
    FROM game_rng_bet_details
    WHERE username = p_username
      AND DATE(bet_time) BETWEEN p_start_date AND p_end_date;
END //
DELIMITER ;

-- =====================================================
-- 示例查询
-- =====================================================

/*
-- 查询指定玩家的Live游戏投注记录
SELECT * FROM game_live_bet_details
WHERE username = 'phoenixgo'
  AND bet_time >= '2025-07-15 00:00:00'
ORDER BY bet_time DESC;

-- 查询指定玩家的RNG游戏投注记录
SELECT * FROM game_rng_bet_details
WHERE username = 'phoenixgo'
  AND bet_time >= '2025-07-15 00:00:00'
ORDER BY bet_time DESC;

-- 查询玩家统一投注记录
SELECT * FROM unified_bet_details
WHERE username = 'phoenixgo'
  AND bet_time >= '2025-07-15 00:00:00'
ORDER BY bet_time DESC;

-- 查询玩家投注统计
SELECT * FROM player_bet_stats
WHERE username = 'phoenixgo';

-- 查询游戏热度排行
SELECT * FROM game_popularity_stats
ORDER BY total_bets DESC
LIMIT 10;

-- 调用存储过程更新玩家每日汇总
CALL UpdatePlayerDailySummary('phoenixgo', '2025-07-15');

-- 调用存储过程获取玩家投注统计
CALL GetPlayerBetStats('phoenixgo', '2025-07-15', '2025-07-15');

-- 查询API调用记录
SELECT * FROM bet_details_api_calls
WHERE username = 'phoenixgo'
ORDER BY created_at DESC;
*/

-- =====================================================
-- 数据库维护建议
-- =====================================================

/*
建议的维护任务:

1. 定期清理旧数据 (保留最近90天):
   DELETE FROM game_live_bet_details WHERE created_at < DATE_SUB(NOW(), INTERVAL 90 DAY);
   DELETE FROM game_rng_bet_details WHERE created_at < DATE_SUB(NOW(), INTERVAL 90 DAY);

2. 定期更新统计表:
   每日运行 UpdatePlayerDailySummary 存储过程

3. 定期优化表:
   OPTIMIZE TABLE game_live_bet_details;
   OPTIMIZE TABLE game_rng_bet_details;
   OPTIMIZE TABLE player_bet_summary;

4. 监控表大小:
   SELECT
       table_name,
       ROUND(((data_length + index_length) / 1024 / 1024), 2) AS 'Size (MB)'
   FROM information_schema.tables
   WHERE table_schema = DATABASE()
     AND table_name IN ('game_live_bet_details', 'game_rng_bet_details', 'player_bet_summary');
*/
