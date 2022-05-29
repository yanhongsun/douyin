use douyin;

CREATE TABLE users
(
  id            BIGINT        NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  u_name        VARCHAR(30)   NOT NULL COMMENT '用户名字',
  passwd        VARCHAR(18)   NOT NULL COMMENT '用户密码',
  follow_count  BIGINT         unsigned DEFAULT '0' COMMENT '关注数', 
  fans_count    BIGINT         unsigned DEFAULT '0' COMMENT '粉丝数',
  PRIMARY KEY (id),
  index index_name(u_name(11))
) ENGINE=InnoDB;

CREATE TABLE videos
(
  id              BIGINT        NOT NULL  AUTO_INCREMENT  ,
  u_id            BIGINT        NOT NULL  COMMENT '用户ID',
  play_url        VARCHAR(255)  NOT NULL  COMMENT '视频地址', 
  cover_url       VARCHAR(255)  NOT NULL  COMMENT '封面地址',
  title           VARCHAR(255)  NOT NULL  COMMENT '视频标题',
  favorite_count  BIGINT        unsigned DEFAULT '0'  COMMENT '点赞数',
  comment_count   BIGINT        unsigned DEFAULT '0'  COMMENT '评论数',
  create_time     BIGINT        NOT NULL  COMMENT '创建时间',
  PRIMARY KEY (id),
  index index_uid(u_id),
  index index_createTime(create_time)
) ENGINE=InnoDB;


CREATE TABLE comment
(
  id                BIGINT  NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  u_id              BIGINT  NOT NULL COMMENT '用户ID',
  v_id              BIGINT  NOT NULL COMMENT '视频ID',
  content           TEXT    NOT NULL COMMENT '评论内容',  
  create_time       BIGINT  NOT NULL COMMENT '创建时间',
  PRIMARY KEY (id),
  index undex_vid(v_id),
  index index_createTime(create_time)
) ENGINE=InnoDB;

CREATE TABLE relation
(
  follower1       BIGINT  NOT NULL   COMMENT '用户A',
  follower2       BIGINT  NOT NULL   COMMENT '用户B',
  tag   ENUM('1', '2', '3') NOT NULL COMMENT '1:A关注B 2:B关注A 3:相互关注',  
  PRIMARY KEY (follower1,follower2)
)ENGINE=InnoDB;

CREATE TABLE favorite
(
  u_id      BIGINT  NOT NULL COMMENT '用户ID',
  v_id      BIGINT  NOT NULL COMMENT '视频ID',
  PRIMARY KEY (u_id,v_id)
) ENGINE=InnoDB;


