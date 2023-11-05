CREATE DATABASE IF NOT EXISTS bilibee_db;

CREATE TABLE IF NOT EXISTS bilibee_db.cron_task_tab
(
    id          bigint(21) unsigned            NOT NULL AUTO_INCREMENT,
    task_id     varchar(64)         DEFAULT '' NOT NULL,
    task_type   int(11) unsigned    DEFAULT 0  NOT NULL, -- 任务类型
    total_num   int(11) unsigned    DEFAULT 0  NOT NULL,
    offset_num  int(11) unsigned    DEFAULT 0  NOT NULL,
    task_status int(11) unsigned    DEFAULT 0  NOT NULL,
    create_time bigint(21) unsigned DEFAULT 0  NOT NULL,
    update_time bigint(21) unsigned DEFAULT 0  NOT NULL,
    KEY idx_update_time (update_time),
    UNIQUE KEY uniq_task_id_task_type (task_id, task_type),
    KEY idx_task_type (task_type),
    PRIMARY KEY (id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;


CREATE TABLE IF NOT EXISTS bilibee_db.video_info_tab
(
    id               bigint(21) unsigned            NOT NULL AUTO_INCREMENT,
    mid              int(11) unsigned    DEFAULT 0  NOT NULL,
    aid              bigint(21) unsigned DEFAULT 0  NOT NULL,
    bvid             varchar(128)        DEFAULT '' NOT NULL,
    url              varchar(256)        DEFAULT '' NOT NULL,
    title            varchar(1024)       DEFAULT '' NOT NULL,
    desc_v2          text,
    pubdate          bigint(21) unsigned DEFAULT 0  NOT NULL,
    user_ctime       bigint(21) unsigned DEFAULT 0  NOT NULL,
    subtitle_content mediumtext,
    raw_str          mediumtext,
    op_status        int(11) unsigned    DEFAULT 0  NOT NULL,
    create_time      bigint(21) unsigned DEFAULT 0  NOT NULL,
    update_time      bigint(21) unsigned DEFAULT 0  NOT NULL,
    KEY idx_mid (mid),
    KEY idx_bvid (bvid),
    UNIQUE KEY uniq_aid (aid),
    PRIMARY KEY (id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;
