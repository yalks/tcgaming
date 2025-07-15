-- =====================================================
-- TC Gaming - 游戏启动信息存储架构
-- =====================================================
-- 此架构存储启动/开始游戏所需的所有必要信息
-- 基于TC Gaming API代码库和main.go示例的分析

-- 游戏启动会话主表
CREATE TABLE game_launch_sessions (
    -- 主键和基本会话信息
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    session_uuid VARCHAR(36) NOT NULL UNIQUE COMMENT '唯一会话标识符(UUID)',

    -- 用户/玩家信息 (来自main.go: username = "phoenixGO")
    username VARCHAR(100) NOT NULL COMMENT '游戏启动的玩家用户名',
    user_password_hash VARCHAR(255) NULL COMMENT '加密密码哈希值，用于会话验证',

    -- 游戏配置 (来自LaunchGameRNG和LaunchGameLottery函数)
    product_type INT NOT NULL COMMENT '游戏产品类型 (EG5=191, PG=98, PP=39, ProductTypeLottery=2, ProductTypeRNG=7)',
    game_mode INT NOT NULL DEFAULT 1 COMMENT '游戏模式: 0=测试, 1=正式 (来自GameModeTest/GameModeLive常量)',
    game_code VARCHAR(50) NOT NULL COMMENT '具体游戏代码 (例如: "EG5353", "Lobby")',
    platform VARCHAR(20) NOT NULL COMMENT '平台类型: "flash", "html5", "all" (来自Platform常量)',

    -- 客户端和游戏类型信息 (来自GetGameList参数)
    client_type VARCHAR(20) NULL COMMENT '客户端类型: "pc", "phone", "web", "html5"',
    game_type VARCHAR(20) NULL COMMENT '游戏类型: "RNG", "LIVE", "PVP"',

    -- 彩票特定配置 (来自LaunchGameLottery函数)
    lottery_view VARCHAR(50) NULL COMMENT '彩票视图参数 (例如: "Lobby")',
    lottery_bet_mode VARCHAR(50) NULL COMMENT '彩票投注模式: "Traditional", "Traditional_Mobile"',

    -- API配置 (来自Config结构体和main.go)
    merchant_code VARCHAR(50) NOT NULL COMMENT 'API认证的商户代码',
    api_url VARCHAR(255) NOT NULL COMMENT 'TC Gaming API端点URL',
    currency VARCHAR(10) NOT NULL DEFAULT 'CNY' COMMENT '货币代码 (例如: CNY, USD)',

    -- 会话管理和安全
    des_key_hash VARCHAR(255) NOT NULL COMMENT 'API通信的加密DES密钥',
    sha256_key_hash VARCHAR(255) NOT NULL COMMENT 'API签名的加密SHA256密钥',

    -- 游戏启动响应数据 (来自LaunchGameResponse结构体)
    game_url TEXT NULL COMMENT '成功启动后生成的游戏URL',
    game_token VARCHAR(500) NULL COMMENT 'API响应的游戏会话令牌',

    -- 财务信息 (来自余额操作)
    player_balance DECIMAL(15,2) NULL COMMENT '游戏启动时的玩家余额',

    -- 引用和交易跟踪
    reference_no VARCHAR(100) NULL COMMENT '交易的唯一引用号 (基于UUID)',

    -- 会话状态和元数据
    launch_status ENUM('PENDING', 'SUCCESS', 'FAILED', 'EXPIRED') NOT NULL DEFAULT 'PENDING' COMMENT '游戏启动状态',
    api_status_code INT NULL COMMENT 'API响应状态码 (0=成功)',
    error_message TEXT NULL COMMENT '启动失败时的错误消息',

    -- 时间戳
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '会话创建时间',
    launched_at TIMESTAMP NULL COMMENT '成功游戏启动时间',
    expires_at TIMESTAMP NULL COMMENT '会话过期时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',

    -- 性能索引
    INDEX idx_username (username),
    INDEX idx_session_uuid (session_uuid),
    INDEX idx_product_type (product_type),
    INDEX idx_game_code (game_code),
    INDEX idx_launch_status (launch_status),
    INDEX idx_created_at (created_at),
    INDEX idx_merchant_code (merchant_code),
    INDEX idx_reference_no (reference_no)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='存储TC Gaming游戏启动所需的所有信息';

-- 彩票系列配置表 (来自SeriesConfig结构体)
CREATE TABLE game_lottery_series (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    session_id BIGINT UNSIGNED NOT NULL COMMENT '引用game_launch_sessions.id',

    -- 系列配置 (来自game.go中的SeriesConfig结构体)
    game_group_code VARCHAR(20) NOT NULL COMMENT '游戏组代码 (例如: "SSC")',
    prize_mode_id INT NOT NULL COMMENT '奖金模式标识符',
    max_series INT NOT NULL COMMENT '最大系列号',
    min_series INT NOT NULL COMMENT '最小系列号',
    max_bet_series INT NOT NULL COMMENT '最大投注系列',
    default_series INT NOT NULL COMMENT '默认系列号',

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (session_id) REFERENCES game_launch_sessions(id) ON DELETE CASCADE,
    INDEX idx_session_id (session_id),
    INDEX idx_game_group_code (game_group_code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='游戏启动的彩票特定系列配置';

-- 游戏信息表 (来自GameInfo结构体)
CREATE TABLE game_catalog (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,

    -- 游戏信息 (来自game.go中的GameInfo结构体)
    tcg_game_code VARCHAR(50) NOT NULL UNIQUE COMMENT 'TC Gaming游戏代码',
    game_name VARCHAR(200) NOT NULL COMMENT '游戏显示名称',
    product_code VARCHAR(50) NOT NULL COMMENT '产品代码',
    product_type VARCHAR(20) NOT NULL COMMENT '产品类型',
    platform VARCHAR(20) NOT NULL COMMENT '支持的平台',
    game_type VARCHAR(20) NOT NULL COMMENT '游戏类型分类',
    game_sub_type VARCHAR(50) NULL COMMENT '游戏子类型',
    display_status INT NOT NULL DEFAULT 1 COMMENT '显示状态 (1=激活, 0=非激活)',
    trial_support BOOLEAN NOT NULL DEFAULT FALSE COMMENT '是否支持试玩模式',

    -- 元数据
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX idx_tcg_game_code (tcg_game_code),
    INDEX idx_product_type (product_type),
    INDEX idx_platform (platform),
    INDEX idx_game_type (game_type),
    INDEX idx_display_status (display_status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='TC Gaming API可用游戏目录';

-- 游戏启动尝试审计日志表
CREATE TABLE game_launch_audit (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    session_id BIGINT UNSIGNED NOT NULL COMMENT '引用game_launch_sessions.id',

    -- 请求详情
    api_method VARCHAR(10) NOT NULL COMMENT '使用的API方法 (lg表示启动游戏)',
    request_payload JSON NULL COMMENT '完整的API请求载荷，用于调试',
    response_payload JSON NULL COMMENT '完整的API响应，用于调试',

    -- 性能指标
    request_duration_ms INT NULL COMMENT 'API请求持续时间(毫秒)',

    -- 审计信息
    ip_address VARCHAR(45) NULL COMMENT '客户端IP地址',
    user_agent TEXT NULL COMMENT '客户端用户代理字符串',

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (session_id) REFERENCES game_launch_sessions(id) ON DELETE CASCADE,
    INDEX idx_session_id (session_id),
    INDEX idx_api_method (api_method),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='所有游戏启动尝试和API调用的审计跟踪';
