/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"

	"github.com/azeezlala/cobra-cli/cmd"
	customerModel "github.com/azeezlala/cobra-cli/model/customer"
)

func main() {
	example := make(map[string][]customerModel.CustomerDetail)
	example["customer"] = []customerModel.CustomerDetail{
		{
			Id:        1,
			FirstName: "lala",
			LastName:  "me",
		},
		{
			Id:        2,
			FirstName: "Azeez",
			LastName:  "lala",
		},
		{
			Id:        3,
			FirstName: "Azeez",
			LastName:  "lala",
		},
		{
			Id:        4,
			FirstName: "Azeez",
			LastName:  "lala",
		},
		{
			Id:        5,
			FirstName: "Azeez",
			LastName:  "lala",
		},
	}
	example["user"] = []customerModel.CustomerDetail{
		{
			Id:        1,
			FirstName: "lala",
			LastName:  "me",
		},
	}
	var customer map[string][]customerModel.CustomerDetail
	err := ReadFromDB(&customer)
	customer["customer"] = append(customer["customer"], customerModel.CustomerDetail{
		Id:        6,
		FirstName: "lolo",
		LastName:  "me",
	})
	example["customer"] = customer["customer"]
	err = WriteToDB(&example)
	fmt.Println("err", err)
	fmt.Println("error", err)
	cmd.Execute()
}
