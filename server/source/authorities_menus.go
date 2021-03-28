package source

import (
	"gin-vue-admin/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

type AuthorityMenus struct {
	AuthorityId string `gorm:"column:sys_authority_authority_id"`
	BaseMenuId  uint   `gorm:"column:sys_base_menu_id"`
}

var authorityMenus = []AuthorityMenus{
	{"888", 1},
	{"888", 2},
	{"888", 3},
	{"888", 4},
	{"888", 5},
	{"888", 6},
	{"888", 7},
	{"888", 8},
	{"888", 9},
	{"888", 10},
	{"888", 11},
	{"888", 12},
	{"888", 13},
	{"888", 14},
	{"888", 15},
	{"888", 16},
	{"888", 17},
	{"888", 18},
	{"888", 19},
	{"888", 20},
	{"888", 21},
	{"888", 22},
	{"888", 23},
	{"8881", 1},
	{"8881", 2},
	{"8881", 8},
	{"9528", 1},
	{"9528", 2},
	{"9528", 3},
	{"9528", 4},
	{"9528", 5},
	{"9528", 6},
	{"9528", 7},
	{"9528", 8},
	{"9528", 9},
	{"9528", 10},
	{"9528", 11},
	{"9528", 12},
	{"9528", 14},
	{"9528", 15},
	{"9528", 16},
	{"9528", 17},

	{"10001", 1},
	{"10001", 2},
	{"10001", 3},
	{"10001", 4},
	{"10001", 5},
	{"10001", 6},
	{"10001", 7},
	{"10001", 8},
	{"10001", 9},
	{"10001", 10},
	{"10001", 11},
	{"10001", 12},
	{"10001", 13},
	{"10001", 14},
	{"10001", 15},
	{"10001", 16},
	{"10001", 17},
	{"10001", 18},
	{"10001", 19},
	{"10001", 20},
	{"10001", 21},
	{"10001", 22},
	{"10001", 23},
	{"10001", 10001},
	{"10001", 10002},
	{"10001", 10003},
	{"10001", 10004},
	{"10001", 10005},
	{"10001", 10006},
	{"10001", 10007},
	{"10001", 10008},
	{"10001", 10009},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_authority_menus 表数据初始化
func (a *authoritiesMenus) Init() error {
	return global.GVA_DB.Table("sys_authority_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ('888', '8881', '9528')").Find(&[]AuthorityMenus{}).RowsAffected == 48 {
			color.Danger.Println("\n[Mysql] --> sys_authority_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authority_menus 表初始数据成功!")
		return nil
	})
}
