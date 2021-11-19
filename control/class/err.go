package class

func CheckError(err error) {
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}
}
