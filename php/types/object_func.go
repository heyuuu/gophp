package types

// Object 对象自动析构方法
func ObjectAutoFree(object *ZendObject) {
	// todo 待重构

	/*    Make sure we hold a reference count during the destructor call
	      otherwise, when the destructor ends the storage might be freed
	      when the refcount reaches 0 a second time
	*/

	// 调用 Dtor 方法
	if (object.GetGcFlags() & IS_OBJ_DESTRUCTOR_CALLED) == 0 {
		object.AddGcFlags(IS_OBJ_DESTRUCTOR_CALLED)
		object.Dtor()
	}

	// 调用 Free 方法
	if (object.GetGcFlags() & IS_OBJ_FREE_CALLED) == 0 {
		object.AddGcFlags(IS_OBJ_FREE_CALLED)
		object.Free()
	}
}
