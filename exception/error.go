package exception

// PanicIfNeeded expose global
func PanicIfNeeded(err interface{})  {
	if err != nil {
		panic(err)
	}
}