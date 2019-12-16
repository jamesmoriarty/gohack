package util

func NeverExit(f func()) {
	defer func() {
		if v := recover(); v != nil {
			go NeverExit(f)
		}
	}()
	f()
}
