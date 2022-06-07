use douyin;

INSERT INTO users(
u_name,
passwd,
follow_count,
fans_count
) VALUES('11111111111@qq.com','douyin',0,1),('22222222222@qq.com','douyin',1,0);
INSERT INTO users(
u_name,
passwd,
follow_count,
fans_count
) VALUES('333333333@qq.com','douyin',0,0),('444444444444@qq.com','douyin',0,0);

INSERT INTO relation(
  follower1 ,
  follower2,
  tag
) VALUES(6,20,2);
insert into videos(
  u_id           ,
  play_url        , 
  cover_url      ,
  title         ,
  favorite_count  ,
  comment_count   ,
  create_time  
)VALUES(2,'http://219.216.86.30:8086/resource/videos/1.mp4',' http://219.216.86.30:8086/resource/cover/1.png','this is a dog',0,0,11),
(4,'http://219.216.86.30:8086/resource/videos/2.mp4',' http://219.216.86.30:8086/resource/cover/2.png','this is a cat',0,0,15);

