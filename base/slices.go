package base

func In[E comparable](target E, vals ...E) bool {
	for _, val := range vals {
		if target == val {
			return true
		}
	}
	return false
}
