package main

import (
    "fmt"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common/compiler"
    "regexp"
    "strings"
)

type ContractParser struct {
    AbiParser abi.ABI
}

func NewContractParser(sourceCode string) (*ContractParser, error) {
    compiledCode, err := compiler.CompileSolidityString("MyContract.sol", sourceCode)
    if err != nil {
        return nil, err
    }

    abiJSON, ok := compiledCode["MyContract.sol:MyContract"]
    if !ok {
        return nil, fmt.Errorf("Contract not found")
    }

    abiParser, err := abi.JSON(strings.NewReader(abiJSON.Info.AbiDefinition))
    if err != nil {
        return nil, err
    }

    return &ContractParser{abiParser}, nil
}

func (p *ContractParser) ParseVariables() {
    for _, variable := range p.AbiParser.Variables {
        varType := variable.Type.String()
        if match, _ := regexp.MatchString(`^uint(\d+)?$`, varType); match {
            fmt.Printf("%s: simple variable\n", variable.Name)
        } else if match, _ := regexp.MatchString(`^\w+\[\]$`, varType); match {
            fmt.Printf("%s: dynamic array\n", variable.Name)
        } else if match, _ := regexp.MatchString(`^mapping\(.*\)$`, varType); match {
            fmt.Printf("%s: mapping\n", variable.Name)
        } else {
            fmt.Printf("%s: unknown type\n", variable.Name)
        }
    }
}
