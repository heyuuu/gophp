package types

func ArrayCopy(target *Array, source *Array) {
	source.Each(func(key ArrayKey, value Zval) {
		target.Update(key, value)
	})
}

func ArrayMerge(target *Array, source *Array, overwrite bool) {
	if overwrite {
		source.Each(func(key ArrayKey, value Zval) {
			target.Update(key, value)
		})
	} else {
		source.Each(func(key ArrayKey, value Zval) {
			target.Add(key, value)
		})
	}
}
