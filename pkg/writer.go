package pkg

import (
	"fmt"
	"github.com/secsy/goftp"
	"os"
)

func Write(c *goftp.Client, list []string,  path, output string)  {
	fmt.Println(list)
	for _, item := range list{
		fmt.Println(item)
		fmt.Println(list)
		out, err := os.Create(item)
		if err != nil {
			fmt.Printf("err: %s", err)
		}
		_ = c.Retrieve(path+item, out)
	}

}
