package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	res := make(map[string]string)

	looking := []string{}
	scanner.Scan()
	for scanner.Scan() {
		contact := scanner.Text()

		parts := strings.Split(contact, " ")
		if len(parts) == 1 {
			looking = append(looking, parts[0])
			continue
		}

		res[parts[0]] = parts[1]
	}

	for _, v := range looking {
		if look, ok := res[v]; ok {
			fmt.Printf("%s=%s\n", v, look)
			continue
		}

		fmt.Println("Not found")
	}
}
