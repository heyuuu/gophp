package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

type RsrcDtorFuncT func(res *types.ZendResource)

var ListDestructors ListDestructorsType

type ListDestructorsType struct {
	nextResourceId int
	hash           map[int]*ZendRsrcListDtorsEntry
}

func (lds *ListDestructorsType) Init() {
	lds.nextResourceId = 1
	lds.hash = make(map[int]*ZendRsrcListDtorsEntry)
}

func (lds *ListDestructorsType) Append(lde *ZendRsrcListDtorsEntry) int {
	lde.resourceId = lds.nextResourceId
	lds.nextResourceId++
	lds.hash[lde.resourceId] = lde
	return lde.resourceId
}

func (lds *ListDestructorsType) Find(res *types.ZendResource) *ZendRsrcListDtorsEntry {
	return lds.hash[res.GetType()]
}

func (lds *ListDestructorsType) GetResourceIdByTypeName(typeName string) int {
	for resourceId, lde := range lds.hash {
		if lde.TypeName() == typeName {
			return resourceId
		}
	}

	return 0
}

func (lds *ListDestructorsType) CleanByModule(moduleNumber int) {
	for resourceId, lde := range lds.hash {
		if lde.moduleNumber != moduleNumber {
			continue
		}
		delete(lds.hash, resourceId)

		// CleanModuleResource
		EG__().PersistentList().Filter(func(_ string, res *types.ZendResource) bool {
			return res.GetType() != resourceId
		})
	}
}

func (lds *ListDestructorsType) Destroy() {
	lds.Init()
}

/**
 * ZendRsrcListDtorsEntry
 */
type ZendRsrcListDtorsEntry struct {
	listDtorEx   RsrcDtorFuncT
	plistDtorEx  RsrcDtorFuncT
	typeName     string
	moduleNumber int
	resourceId   int
}

func NewZendRsrcListDtorsEntry(listDtorEx RsrcDtorFuncT, plistDtorEx RsrcDtorFuncT, typeName string, moduleNumber int) *ZendRsrcListDtorsEntry {
	return &ZendRsrcListDtorsEntry{
		listDtorEx:   listDtorEx,
		plistDtorEx:  plistDtorEx,
		typeName:     typeName,
		moduleNumber: moduleNumber,
	}
}

func (this *ZendRsrcListDtorsEntry) TypeName() string              { return this.typeName }
func (this *ZendRsrcListDtorsEntry) GetListDtorEx() RsrcDtorFuncT  { return this.listDtorEx }
func (this *ZendRsrcListDtorsEntry) GetPlistDtorEx() RsrcDtorFuncT { return this.plistDtorEx }
func (this *ZendRsrcListDtorsEntry) GetTypeName() *byte            { return this.typeName }
func (this *ZendRsrcListDtorsEntry) GetModuleNumber() int          { return this.moduleNumber }
func (this *ZendRsrcListDtorsEntry) GetResourceId() int            { return this.resourceId }

/**
 * functions
 */
