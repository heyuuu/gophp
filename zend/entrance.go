package zend

// todo 初始化入口
var currEntrance *Entrance = &Entrance{uid: -1, gid: -1, inode: -1, mtime: -1}

func CurrEntrance() *Entrance {
	return currEntrance
}

type Entrance struct {
	path  string
	uid   int
	uName string
	gid   int
	inode int
	mtime int64
}

func (e Entrance) Uid() int     { return e.uid }
func (e Entrance) Gid() int     { return e.gid }
func (e Entrance) Inode() int   { return e.inode }
func (e Entrance) Mtime() int64 { return e.mtime }

func (e Entrance) UserName() string {
	// todo uid 换用户名
	return e.uName
}
