package db

import (
	"context"
	"douyin/pkg/constants"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/// TODO  消除record not found警告
type Relation struct {
	Follower1 int64 `gorm:"column:follower1"`
	Follower2 int64 `gorm:"column:follower2"`
	Tag       int   `gorm:"column:tag"`
}

// User user model
type User struct {
	ID            int64  `gorm:"column:id"`
	Username      string `gorm:"column:u_name"`
	Password      string `gorm:"column:passwd"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:fans_count"`
}

type UserList struct {
	ID             int64
	Name           string
	Follow_count   int64
	Follower_count int64
	Is_follow      bool
}

func (v Relation) TableName() string {
	return constants.RelationTableName
}

// 比自己id小的关注列表
func GetFollows1ID(ctx context.Context, follower1 int64, str int) ([]Relation, error) {
	var tmp []Relation
	if err := DB.WithContext(ctx).Where("follower1 = ? and tag = ?", follower1, str).Take(&tmp).Error; err != nil {
		//fmt.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//fmt.Println("没有此类数据")
		} else if err != nil {
			return nil, err
		}
	}
	return tmp, nil
}

// 比自己id大的关注列表
func GetFollows2ID(ctx context.Context, follower1 int64, str int) ([]Relation, error) {
	var tmp []Relation
	if err := DB.WithContext(ctx).Where("follower2 = ? and tag = ?", follower1, str).Take(&tmp).Error; err != nil {
		//fmt.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//fmt.Println("没有此类数据")
		} else if err != nil {
			return nil, err
		}
	}
	return tmp, nil
}

// 根据单向关注的用户id查询出关注的具体信息
func GetFollowsInfo(ctx context.Context, res []Relation, tag int) ([]User, error) {
	len := len(res)
	var users []User
	if tag == 1 {
		for i := 0; i < len; i++ {
			var user User
			if err := DB.WithContext(ctx).Where("id=?", res[i].Follower1).Take(&user).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					//fmt.Println("没有此类数据")
				} else if err != nil {
					return nil, err
				}
			} else {
				users = append(users, user)
			}
		}
	} else if tag == 2 {
		for i := 0; i < len; i++ {
			var user User
			if err := DB.WithContext(ctx).Where("id=?", res[i].Follower2).Take(&user).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					//fmt.Println("没有此类数据")
				} else if err != nil {
					return nil, err
				}
			} else {
				users = append(users, user)
			}
		}
	}

	return users, nil
}

// 返回单向关注的respone
func ReturnFalseUserList(users []User) []UserList {
	var ans []UserList
	for _, v := range users {
		//TODO
		ans = append(ans, UserList{
			ID:             v.ID,
			Follow_count:   v.FollowCount,
			Follower_count: v.FollowerCount,
			Is_follow:      false,
			Name:           v.Username,
		})
	}
	return ans
}

// 返回双向关注的respone
func ReturnTureUserList(users []User) []UserList {
	var ans []UserList
	for _, v := range users {
		//TODO
		ans = append(ans, UserList{
			ID:             v.ID,
			Follow_count:   v.FollowCount,
			Follower_count: v.FollowerCount,
			Is_follow:      true,
			Name:           v.Username,
		})
	}
	return ans
}

//查询单向关注列表
func GetOneWayFollows(ctx context.Context, follower1 int64, str1 int, str2 int) ([]User, error) {
	// 中间变量  存储单向关系的relation表
	var tmp []Relation
	//根据单向关注的用户id查询出关注的具体信息
	var users []User
	// 比自己id大的关注列表
	tmp, err := GetFollows2ID(ctx, follower1, str1)
	if err != nil {
		return nil, err
	}
	if len(tmp) != 0 {
		users, err = GetFollowsInfo(ctx, tmp, 1)
		if err != nil {
			return nil, err
		}
	}

	// 比自己id小的关注列表
	tmp, err = GetFollows1ID(ctx, follower1, str2)
	if err != nil {
		return nil, err
	}
	var users1 []User
	if len(tmp) != 0 {
		users1, err = GetFollowsInfo(ctx, tmp, 2)
		if err != nil {
			return nil, err
		}
	}

	users = append(users, users1...)
	//fmt.Println(users1)
	return users, nil
}

//查询双向关注列表
func GetTwoWayFollows(ctx context.Context, follower1 int64) ([]User, error) {
	// 中间变量  存储双向关系的relation表
	var tmp []Relation
	//根据单向关注的用户id查询出关注的具体信息
	var users []User

	// 比自己id大的关注列表
	tmp, err := GetFollows2ID(ctx, follower1, 3)
	if err != nil {
		return nil, err
	}
	if len(tmp) != 0 {
		users, err = GetFollowsInfo(ctx, tmp, 1)
		if err != nil {
			return nil, err
		}
	}

	var users1 []User
	// 比自己id小的关注列表
	tmp, err = GetFollows1ID(ctx, follower1, 3)
	if err != nil {
		return nil, err
	}
	if len(tmp) != 0 {
		users1, err = GetFollowsInfo(ctx, tmp, 2)
		if err != nil {
			return nil, err
		}
	}
	users = append(users, users1...)
	return users, nil
}

func GetFollows(ctx context.Context, follower1 int64) ([]UserList, error) {
	var users []User
	users, err := GetOneWayFollows(ctx, follower1, 2, 1)
	// fmt.Println("1111")
	if err != nil {
		return nil, err
	}
	var tmp []User
	tmp, err = GetTwoWayFollows(ctx, follower1)
	if err != nil {
		return nil, err
	}
	users = append(users, tmp...)

	// 返回所有关注的用户的response
	userList := ReturnTureUserList(users)
	fmt.Println("db 的 GetFollows 查询结束")
	fmt.Println(userList)
	return userList, nil
}

//查询粉丝列表
func GetFans(ctx context.Context, follower1 int64) ([]UserList, error) {

	users, err := GetOneWayFollows(ctx, follower1, 1, 2)

	if err != nil {
		return nil, err
	}
	//  仅仅是单向的粉丝
	userList := ReturnFalseUserList(users)
	fmt.Println("单向粉丝：")
	fmt.Println(userList)
	users3, err := GetTwoWayFollows(ctx, follower1)
	if err != nil {
		return nil, err
	}
	fmt.Println("双向粉丝：")
	fmt.Println(users3)
	// 返回所有关注的用户的response
	tmp := ReturnTureUserList(users3)
	userList = append(userList, tmp...)
	return userList, nil
}

// 这个查询可能会慢
func IsFollowed(ctx context.Context, userId, otherId int64) (bool, error) {
	fmt.Println("=================relation.db.is_followed")
	if userId == otherId {
		return false, nil
	}
	var tmp []Relation
	if userId < otherId {

		if err := DB.WithContext(ctx).Where("follower1 = ? and follower2 = ?", userId, otherId).Take(&tmp).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				//fmt.Println("没有此类数据")
				return false, nil
			} else if err != nil {
				return false, err
			}
		} else {
			for _, v := range tmp {
				if v.Tag == 1 || v.Tag == 3 {
					return true, nil
				}
			}
			return false, nil
		}
	} else {
		if err := DB.WithContext(ctx).Where("follower1 = ? and follower2 = ?", otherId, userId).Take(&tmp).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				//fmt.Println("没有此类数据")
				return false, nil
			} else if err != nil {
				return false, err
			}
		} else {
			for _, v := range tmp {
				if v.Tag == 2 || v.Tag == 3 {
					return true, nil
				}
			}
			return false, nil
		}
	}
	return false, nil
}

func AddFans(ctx context.Context, userId int64) error {
	user := User{}
	if err := DB.WithContext(ctx).Model(&user).Where("id = ?", userId).Update("fans_count", gorm.Expr("fans_count + 1")).Error; err != nil {
		return err
	}
	return nil
}

func AddFollows(ctx context.Context, userId int64) error {
	user := User{}
	if err := DB.WithContext(ctx).Model(&user).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
		return err
	}
	return nil
}

// TODO
func DelFans(ctx context.Context, userId int64) error {
	user := User{}
	if err := DB.WithContext(ctx).Model(&user).Where("id = ?", userId).Update("fans_count", gorm.Expr("fans_count - 1")).Error; err != nil {
		return err
	}
	return nil
}

func DelFollows(ctx context.Context, userId int64) error {
	user := User{}
	if err := DB.WithContext(ctx).Model(&user).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func Follow(ctx context.Context, userId, otherId int64) error {
	if userId == otherId {
		// 自己关注自己，不进行处理。
		// return simple.NewErrorMsg("自己不能关注自己")
		return nil
	}
	tag, err := IsFollowed(ctx, userId, otherId)
	if err != nil {
		return err
	}
	if tag {
		return nil
	}

	// 如果是枚举类型  存在问题
	if userId < otherId {
		// 当 `id` 有冲突时，更新指定列为新值
		if err := DB.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "follower1"}, {Name: "follower2"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"tag": gorm.Expr("tag | ?", 1)}), //更新字段 配合Expr使用
			//创建字段,
		}).Create(&Relation{
			Follower1: userId,
			Follower2: otherId,
			Tag:       1,
		}).Error; err != nil {
			return err
		}

	} else {

		if err := DB.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "follower1"}, {Name: "follower2"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"tag": gorm.Expr("tag | ?", 2)}), //更新字段 配合Expr使用
			//创建字段,
		}).Create(&Relation{
			Follower1: otherId,
			Follower2: userId,
			Tag:       2,
		}).Error; err != nil {
			return err
		}
	}
	//更新粉丝个数以及关注个数
	err = AddFollows(ctx, userId)
	if err != nil {
		return err
	}
	err = AddFans(ctx, otherId)
	if err != nil {
		return err
	}
	return nil
}

func Unfollow(ctx context.Context, userId, otherId int64) error {
	fmt.Println("db 中的unfollow函数")
	fmt.Println(userId, otherId)
	if userId == otherId {
		// 自己取消关注自己，不进行处理。
		// return simple.NewErrorMsg("自己不能取消关注自己")
		return nil
	}
	tag, err := IsFollowed(ctx, userId, otherId)
	if err != nil {
		return err
	}
	if !tag {
		return nil
	}

	//TODO 取消关注的关键步骤
	if userId < otherId {
		relation := Relation{}
		if err := DB.WithContext(ctx).Model(&relation).Where("follower1 = ? and follower2 = ?", userId, otherId).Update("tag", gorm.Expr("tag & ?", 2)).Error; err != nil {
			return err
		}

		fmt.Println(relation)
		tmp := Relation{}
		if err := DB.WithContext(ctx).Where("follower1 = ? and follower2 = ?", userId, otherId).Take(&tmp).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				//fmt.Println("没有此类数据")
			} else if err != nil {
				return err
			}
		}
		fmt.Println(tmp)
		if tmp.Tag == 0 {
			if err = DB.WithContext(ctx).Where("follower1 = ? and follower2 = ?", userId, otherId).Delete(&relation).Error; err != nil {
				return err
			}
		}

	} else {
		relation := Relation{}
		if err := DB.WithContext(ctx).Model(&relation).Where("follower1 = ? and follower2 = ?", otherId, userId).Update("tag", gorm.Expr("tag & ?", 1)).Error; err != nil {
			return err
		}
		tmp := Relation{}
		if err := DB.WithContext(ctx).Where("follower1 = ? and follower2 = ?", otherId, userId).Take(&tmp).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				//fmt.Println("没有此类数据")
			} else if err != nil {
				return err
			}
		}
		if tmp.Tag == 0 {
			if err = DB.WithContext(ctx).Where("follower1 = ? and follower2 = ?", otherId, userId).Delete(&relation).Error; err != nil {
				return err
			}
		}
	}

	//更新粉丝个数以及关注个数
	err = DelFollows(ctx, userId)
	if err != nil {
		return nil
	}
	err = DelFans(ctx, otherId)
	if err != nil {
		return nil
	}
	return nil
}
