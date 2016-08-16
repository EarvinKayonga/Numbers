package main

func main() {
	var filePath string = getFilePath()
	var content string = read(filePath)
	process(content)
}
