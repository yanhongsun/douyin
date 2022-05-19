use douyin;

CREATE TABLE users
(
  id      int       NOT NULL AUTO_INCREMENT,
  u_name    VARCHAR(30)  NOT NULL ,
  passwd VARCHAR(18),
  follow_count BIGINT, 
  fans_count  BIGINT ,
  PRIMARY KEY (id),
  index index_name(u_name(11))
) ENGINE=InnoDB;

CREATE TABLE vedios
(
  id     BIGINT  NOT NULL AUTO_INCREMENT,
  u_id    BIGINT  NOT NULL,
  play_url  VARCHAR(255), 
  cover_url   VARCHAR(255) ,
  title VARCHAR(255) ,
  favorite_count BIGINT,
  comment_count BIGINT,
  create_time BIGINT,
  PRIMARY KEY (id),
  index index_uid(u_id),
  index index_createTime(create_time)
) ENGINE=InnoDB;

CREATE TABLE comment
(
  id     BIGINT  NOT NULL AUTO_INCREMENT,
  u_id     BIGINT  NOT NULL,
  vid BIGINT,
  content   TEXT,  
  create_time  BIGINT,
  PRIMARY KEY (id),
  index undex_vid(vid),
  index index_createTime(create_time)
) ENGINE=InnoDB;

CREATE TABLE relation
(
  follower1     BIGINT ,
  follower2 BIGINT,
  tag   ENUM('1', '2', '3') NOT NULL,  
  PRIMARY KEY (follower1,follower2)
)ENGINE=InnoDB;

CREATE TABLE favorite
(
  u_id      BIGINT ,
  v_id  BIGINT,
  PRIMARY KEY (u_id,v_id)
) ENGINE=InnoDB;


