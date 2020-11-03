package main

func main() {
	fat := 1
	for i := 0; i < 10; i++ {
		if (i%3) == 0 && i != 0 {
			for j := 0; j < i; j++ {
				fat = fat * (i - j)
			}
			print(fat, "\n")
			fat = 1
		}
	}
}
