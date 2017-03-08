/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"errors"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"

)

// Test comment
// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type Business 	struct {
	ID			string   `json:"ID"`
	Name   		string   `json:"Name"`
	Street 		string   `json:"NumberOfTransactions"`
	Zip      	string 	 `json:"Status"`
}


// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
    // Create a bank
	var bank Business
	bank.ID 	= "B1";
	bank.Name 	= "Local Bank"
	bank.Street = "123 Elm Street"
	bank.Zip  	= "2345"

	// Convert the bank struct to bytes
	jsonAsBytes, _  := json.Marshal(bank)
	
	// Add the bank to the blockchain
	err 			:= stub.PutState(bank.ID, jsonAsBytes)		
	
	// Check for errors 
	if err != nil {
		fmt.Println("Error Creating bank during initialization")
		return nil, err
	}
	
	return nil, nil
}

// ============================================================================================================================
// Run - Our entry point for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
func (t *SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	
	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} 
	
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation")
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	
	if function == "getBusiness" { return t.getBusiness(stub, args[1]) }

	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}

// ============================================================================================================================
// Get Open Points member account from the blockchain
// ============================================================================================================================
func (t *SimpleChaincode) getBusiness(stub shim.ChaincodeStubInterface, businessID string)([]byte, error){
	
	fmt.Println("Start getBusiness")
	fmt.Println("Looking for business with ID " + businessID);

	//get the User index
	businessAsBytes, err := stub.GetState(businessID)
	if err != nil {
		return nil, errors.New("Failed to get business from blockchain")
	}

	return businessAsBytes, nil
}