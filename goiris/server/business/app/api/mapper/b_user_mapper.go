package mapper

import (
	"goiris/business/app/model"
	"goiris/business/app/api/vo"
	"goiris/common/storage"
)

type BUserMapper struct {}

func (user *BUserMapper) Insert(vo *vo.AcceptBUserVO) error {
	return user.createOrUpdate(vo, true)
}

func (user *BUserMapper) FindOne(cond *model.BUser) (*model.BUser, error) {
	var (
		err   error
		resut = model.BUser{}
	)
	if err = storage.G_DB.Where(&cond).First(&resut).Error; err != nil {
		return nil, err
	}
	return &resut, nil
}

func (user *BUserMapper) UpdateOne(vo *vo.AcceptBUserVO) error {
	return user.createOrUpdate(vo, false)
}

func (user *BUserMapper) Delete(ids []int64) error {
	var (
		err error
		tx = storage.G_DB.Begin()
	)
	if err = tx.Delete(model.BUser{}, "id in (?)", ids).Error; err != nil {
		goto ERR
	}
	tx.Commit()
	return nil
ERR:
	tx.Rollback()
	return err
}

func (user *BUserMapper) FindList(vo *vo.BUserVO) (int64, []*model.BUser, error) {
	var (
		err    error
		total  int64
		result = make([]*model.BUser, 0)
	)
	if err = storage.G_DB.Model(&model.BUser{}).Count(&total).Error; err != nil {
		goto ERR
	}
	if err = storage.G_DB.
		Select("id, username, gender, name, age, phone, email, avatar, create_time").
		Offset(vo.Start).Limit(vo.Size).Find(&result).Error; err != nil {
		goto ERR
	}
	return total, result, nil
ERR:
	return 0, nil, err
}
// 更改用户表,角色表
func (user *BUserMapper) createOrUpdate(vo *vo.AcceptBUserVO, isCreate bool) error {
	var (
		err error
		tx = storage.G_DB.Begin()
	)
	// 用户表
	if isCreate {
		if err = tx.Create(&vo.BUser).Error; err != nil {
			goto ERR
		}
	} else {
		if err = tx.Model(&model.BUser{}).Save(&vo.BUser).Error; err != nil {
			goto ERR
		}
	}
	tx.Commit()
	return nil
ERR:
	tx.Rollback()
	return err
}