package main

func safeInt(ptr *int) int {
	if ptr == nil {
		return 0
	}
	return *ptr
}
