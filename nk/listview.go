package nk

func (lv *ListView) Begin() int {
	return int(lv.begin)
}

func (lv *ListView) End() int {
	return int(lv.end)
}

func (lv *ListView) Count() int {
	return int(lv.count)
}
