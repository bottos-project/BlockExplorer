CREATE TABLE IF NOT EXISTS `users`(
    user_id VARCHAR(32) NOT NULL,
    username VARCHAR(32) NOT NULL DEFAULT '',
    nick_name VARCHAR(32) NOT NULL DEFAULT '',
    gender INT,
    avatar_url TEXT,
    phone VARCHAR(32),
    openid VARCHAR(32),
    session_key VARCHAR(32),
    is_vip BOOL
);
CREATE TABLE IF NOT EXISTS `products`(
    product_id VARCHAR(32) NOT NULL,
    product_name VARCHAR(32) NOT NULL,
    respons_admin_id VARCHAR(32),
    product_img TEXT,
    price DOUBLE NOT NULL,
    active_price DOUBLE,
    vip_price DOUBLE,
    content TEXT,
    product_describe TEXT,
    loop_imgs TEXT
);

CREATE TABLE IF NOT EXISTS `appoint`(
    appoint_id VARCHAR(32) NOT NULL,
    product_id VARCHAR(32) NOT NULL,
    user_id VARCHAR(32) NOT NULL,
    openid VARCHAR(32) NOT NULL,
    appoint_date VARCHAR(32),
    appoint_time VARCHAR(32),
    complate INT,
    remark TEXT
);

CREATE TABLE IF NOT EXISTS `appoint_date`(
    appoint_id VARCHAR(32) NOT NULL,
    appoint_date VARCHAR(32) UNIQUE
);
CREATE TABLE IF NOT EXISTS `admins`(
    admin_id VARCHAR(32) NOT NULL,
    admin_name VARCHAR(32),
    admin_password VARCHAR(32),
    admin_phone VARCHAR(32)
);

CREATE TABLE IF NOT EXISTS `test`(
    username VARCHAR(32),
    pwd VARCHAR(32)
)