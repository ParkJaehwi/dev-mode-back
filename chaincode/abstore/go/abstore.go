package main

import (
    "encoding/json"
    "fmt"
    
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// IdentityContract 체인코드의 스마트 계약 구조체
type IdentityContract struct {
    contractapi.Contract
}

// Identity 신원 정보를 나타내는 구조체
type Identity struct {
    Name     string `json:"name"`
    Gender   string `json:"gender"`
    DOB      string `json:"dob"` // Date of Birth
    Contact  string `json:"contact"`
    IDNumber string `json:"idNumber"`
}

// CreateIdentity 신원 정보를 생성하는 함수
func (s *IdentityContract) CreateIdentity(ctx contractapi.TransactionContextInterface, name string, gender string, dob string, contact string, idNumber string) error {
    // 신원 정보가 이미 존재하는지 확인
    identityJSON, err := ctx.GetStub().GetState(idNumber)
    if err != nil {
        return fmt.Errorf("failed to read from world state: %v", err)
    }
    if identityJSON != nil {
        return fmt.Errorf("identity with ID %s already exists", idNumber)
    }

    // 새로운 신원 정보 생성
    identity := Identity{
        Name:     name,
        Gender:   gender,
        DOB:      dob,
        Contact:  contact,
        IDNumber: idNumber,
    }
 
    // 신원 정보를 JSON 형식으로 변환하여 월드 스테이트에 저장
    identityJSON, err = json.Marshal(identity)
    if err != nil {
        return err
    }
 
    return ctx.GetStub().PutState(idNumber, identityJSON)
}

// QueryIdentity 신원 정보를 조회하는 함수
func (s *IdentityContract) QueryIdentity(ctx contractapi.TransactionContextInterface, idNumber string) (*Identity, error) {
    // 월드 스테이트에서 신원 정보를 가져와서 구조체로 언마샬링
    identityJSON, err := ctx.GetStub().GetState(idNumber)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if identityJSON == nil {
        return nil, fmt.Errorf("identity with ID %s does not exist", idNumber)
    }

    var identity Identity
    err = json.Unmarshal(identityJSON, &identity)
    if err != nil {
        return nil, err
    }

    return &identity, nil
}


func main() {
    identityChaincode, err := contractapi.NewChaincode(&IdentityContract{})
    if err != nil {
        fmt.Printf("Error creating identity chaincode: %s", err)
        return
    }

    if err := identityChaincode.Start(); err != nil {
        fmt.Printf("Error starting identity chaincode: %s", err)
    }
}
