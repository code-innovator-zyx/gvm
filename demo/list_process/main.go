package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/26 上午11:19
* @Package:
 */
func main() {
	items := []list.Item{
		SimpleListItem("Ramen"),
		SimpleListItem("Tomato Soup"),
		SimpleListItem("Hamburgers"),
		SimpleListItem("Cheeseburgers"),
		SimpleListItem("Currywurst"),
		SimpleListItem("Okonomiyaki"),
		SimpleListItem("Pasta"),
		SimpleListItem("Fillet Mignon"),
		SimpleListItem("Caviar"),
		SimpleListItem("Just Wine"),
	}

	const defaultWidth = 20

	l := list.New(items, SimpleListItemDelegate{}, defaultWidth, SimpleListHeight)
	l.Title = "select a fixed version to install"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)

	program := NewListProgram(l)
	finalModel, err := program.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
	}
	fmt.Println("choice", finalModel.(*SimpleListModel).Index())
}
