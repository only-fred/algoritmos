package main

func main() {
	var j, sum, media int

	for i := 50; i <= 70; i++ {
		j = j + 1

		sum = sum + i
	}
	media = sum / j
	print(media)
}
