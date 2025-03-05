package daos

import (
	"demo/config"
	"fmt"
	"time"

	"github.com/ma-guo/niuhe/db"
	cache "github.com/patrickmn/go-cache"
	"xorm.io/xorm"
)

type Dao struct {
	_db *db.DB
}

func NewDao() *Dao {
	return &Dao{
		_db: db.NewDB(config.MainDB),
	}
}

func (dao *Dao) Close() {
	if dao._db != nil {
		dao._db.Close()
	}
}

func (dao *Dao) db() *xorm.Session {
	if dao._db != nil {
		return dao._db.GetDB()
	}
	// 理论上执行不到这里
	return nil
}
func (dao *Dao) Atom(fn func() error) error {
	if dao._db != nil {
		return dao._db.Atom(fn)
	}
	return nil
}
func (dao *Dao) getCache(prefix string, args ...interface{}) (interface{}, bool) {
	key := prefix
	for _, arg := range args {
		key += fmt.Sprintf(":%v", arg)
	}
	return localCache.Get(key)
}

func (dao *Dao) setCache(val interface{}, duration time.Duration, prefix string, args ...interface{}) {
	key := prefix
	for _, arg := range args {
		key += fmt.Sprintf(":%v", arg)
	}
	localCache.Set(key, val, duration)
}

// 查找记录
// @param row 要查找结构
func (dao *Dao) GetBy(row interface{}) (bool, error) {
	has, err := dao.db().Get(row)
	if err != nil {
		return false, err
	}
	return has, nil
}

// 插入记录
// @param row 要插入的字段
func (dao *Dao) Insert(row interface{}) (bool, error) {
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

// 更新记录, row 结构中需要包含 id 字段
// @param row 结构中需要包含 id 字段
func (dao *Dao) Update(id int64, row interface{}) (bool, error) {
	affected, err := dao.db().Where("`id`=?", id).AllCols().Update(row)
	return affected > 0, err
}

// 删除记录
// @param row 要删除的记录, 需为指针
func (dao *Dao) Delete(rows interface{}) (bool, error) {
	affected, err := dao.db().Delete(rows)
	return affected > 0, err
}

// 分页查找
// @param page 页码
// @param size 每页大小
func (dao *Dao) Limit(session *xorm.Session, page, size int) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	session.Limit(size, (page-1)*size)
}

// like 查找
func (dao *Dao) Like(session *xorm.Session, col string, val any) {
	value := fmt.Sprintf("%v", val)
	if len(col) > 0 && value != "" && value != "0" {
		session.Where(fmt.Sprintf("%v LIKE ?", col), "%"+value+"%")
	}
}

var localCache *cache.Cache

func init() {
	localCache = cache.New(5*time.Minute, 30*time.Second)
}
