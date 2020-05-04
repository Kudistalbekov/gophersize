package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func lAdd(cmd *cobra.Command, args []string, flagval string) error {
	usercommand := strings.Join(args, " ")
	if list, ok := Data[flagval]; ok == true {
		Data[flagval] = Yaml{Commands: append(list.Commands, usercommand)}
		Infolog.Println("added!")
		return nil
	}
	return errors.New(flagval + " does not exist")
}

func cAdd(cmd *cobra.Command, args []string, flagval string) error {
	if _, ok := Data[flagval]; ok == true {
		return errors.New(flagval + " is already exists")
	}
	Data[flagval] = Yaml{}
	Infolog.Println("created!")
	return nil
}

func llist(cmd *cobra.Command, args []string, flagval string) {
	printfunc := func(s []string) {
		for i, v := range s {
			fmt.Printf("%v---%v\n", i, v)
		}
	}

	if len(flagval) != 0 {
		if val, ok := Data[flagval]; ok == true {
			printfunc(val.Commands)
		} else {
			Errorlog.Println("not found!")
		}
	} else {
		for key, val := range Data {
			fmt.Println(key)
			printfunc(val.Commands)
			fmt.Println()
		}
	}
}

func lremove(cmd *cobra.Command, args []string, flagval string) error {
	var rem []int
	var str []string
	check := func(i int) bool {
		for _, v := range rem {
			if v == i {
				return false
			}
		}
		return true
	}

	for _, uservalue := range args {
		if n, err := strconv.Atoi(uservalue); err != nil {
			Errorlog.Printf("%v can't be converted to int\n", uservalue)
		} else {
			rem = append(rem, n)
		}
	}

	for i, val := range Data[flagval].Commands {
		if i := check(i); i == true {
			str = append(str, val)
		}
	}
	Data[flagval] = Yaml{Commands: str}
	return nil
}
