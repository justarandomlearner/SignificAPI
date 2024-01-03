package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	sword "github.com/justarandomlearner/SignificAPI/internal/service"
)

func main() {
	word := os.Args[1]

	service := sword.NewService()

	ctx := context.Background()
	payload, err := service.FindWord(ctx, word)
	if err != nil {
		log.Fatalf("couldn't find word; %v", err)
	}

	fmt.Printf("Verbete: %q\n", payload.Word)

	for i, m := range payload.Meanings {
		if m.GramGrp != "" {
			fmt.Println(m.GramGrp)
		}

		fmt.Printf("%d. %s\n", i+1, strings.TrimFunc(m.Def, func(c rune) bool { return c == '\n' || c == '\r' }))

		usage := ""
		for _, u := range m.UsageInfo {
			usage += u.Text + " "
		}

		if usage != "" {
			fmt.Println(usage)
		}

	}

}
