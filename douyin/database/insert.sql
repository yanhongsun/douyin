use douyin;

INSERT INTO users(
u_name,
passwd,
follow_count,
fans_count
) VALUES('11111111111@qq.com','douyin',0,1),('22222222222@qq.com','douyin',1,0);

INSERT INTO relation(
  follower1 ,
  follower2,
  tag
) VALUES(1,2,1);

