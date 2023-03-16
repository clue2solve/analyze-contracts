package main

import "github.com/username/analyze-contracts/contractparser"


func main() {
	sourceCode := `
        pragma solidity ^0.8.0;

        contract MyContract {
            uint public myNumber;
            uint[] public myArray;
            mapping(address => uint) public myMap;

            constructor(uint _myNumber) {
                myNumber = _myNumber;
            }

            function setNumber(uint _myNumber) public {
                myNumber = _myNumber;
            }

            function setArray(uint[] memory _myArray) public {
                myArray = _myArray;
            }

            function setMap(address _address, uint _value) public {
                myMap[_address] = _value;
            }
        }
    `

	parser, err := NewContractParser(sourceCode)
	if err != nil {
		panic(err)
	}

	parser.ParseVariables()
}
