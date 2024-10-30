package main

import (
	"fmt"
	"log"
	"os/exec"
)

func download(link, dir string) {
	cmd := exec.Command(
		"yt-dlp",
		"--embed-thumbnail",
		link,
		"-o",
		"%(playlist_index)s - %(title)s.%(ext)s",
	)

	cmd.Dir = dir

	if err := cmd.Run(); err != nil {
		log.Fatal("Erro ao executar o download:", err)
	} else {
		fmt.Println("Download concluído com sucesso!")
	}
}

func main() {

	var catchLink string

	fmt.Println("Link: ")
	fmt.Scan(&catchLink)

	var catchdir string

	fmt.Println("Dir: ")
	fmt.Scan(&catchdir)

	download(catchLink, catchdir)
}
