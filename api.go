package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ShoppingItem struct {
	Name     string
	Quantity int
	Notes    string
}

func addItem(list []ShoppingItem, item ShoppingItem) []ShoppingItem {
	return append(list, item)
}

func removeItem(list []ShoppingItem, index int) []ShoppingItem {
	if index < 0 || index >= len(list) {
		return list
	}
	return append(list[:index], list[index+1:]...)
}

func displayList(list []ShoppingItem) {
	for i, item := range list {
		fmt.Printf("%d. %s (%d) - %s\n", i+1, item.Name, item.Quantity, item.Notes)
	}
}

func saveListToFile(list []ShoppingItem, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, item := range list {
		line := fmt.Sprintf("%s,%d,%s\n", item.Name, item.Quantity, item.Notes)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func loadListFromFile(filename string) ([]ShoppingItem, error) {
	var list []ShoppingItem

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid line in file: %s", line)
		}

		name := parts[0]
		quantity, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		notes := parts[2]

		item := ShoppingItem{
			Name:     name,
			Quantity: quantity,
			Notes:    notes,
		}
		list = append(list, item)
	}

	if err != nil {
		return nil, err
	}

	return list, nil
}

func main() {
	var shoppingList []ShoppingItem

	for {
		fmt.Println("Shopping List Manager")
		fmt.Println("1. Add Item")
		fmt.Println("2. Remove Item")
		fmt.Println("3. Display List")
		fmt.Println("4. Save List")
		fmt.Println("5. Load List")
		fmt.Println("6. Quit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter item name: ")
			var name string
			fmt.Scanln(&name)

			fmt.Print("Enter quantity: ")
			var quantity int
			fmt.Scanln(&quantity)

			fmt.Print("Enter notes: ")
			var notes string
			fmt.Scanln(&notes)

			item := ShoppingItem{Name: name, Quantity: quantity, Notes: notes}
			shoppingList = addItem(shoppingList, item)
			fmt.Println("Item added to the list.")
		case 2:
			fmt.Print("Enter the index of the item to remove: ")
			var index int
			fmt.Scanln(&index)

			shoppingList = removeItem(shoppingList, index-1)
			fmt.Println("Item removed from the list.")
		case 3:
			displayList(shoppingList)
		case 4:
			fmt.Print("Enter the filename to save the list: ")
			var filename string
			fmt.Scanln(&filename)

			if err := saveListToFile(shoppingList, filename); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Shopping list saved to", filename)
			}
		case 5:
			fmt.Print("Enter the filename to load the list: ")
			var filename string
			fmt.Scanln(&filename)

			loadedList, err := loadListFromFile(filename)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				shoppingList = loadedList
				fmt.Println("Shopping list loaded from", filename)
			}
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
