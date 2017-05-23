package core

import (
	"encoding/json"
	"log"
	"strconv"
	"testing"
)

func TestListDelegates(t *testing.T) {
	arkapi := NewArkClient(nil)

	//params := TransactionQueryParams{Limit: 10, SenderID: "AQLUKKKyKq5wZX7rCh4HJ4YFQ8bpTpPJgK"}

	deleResponse, _, err := arkapi.ListDelegates()
	if deleResponse.Success {
		log.Println(t.Name(), "Success, returned ", deleResponse.TotalCount, "delegates")
		/*for _, element := range deleResponse.Delegates {
			log.Println(element.Username)
		}*/
	} else {
		t.Error(err.Error())
	}
}

func TestGetDelegateUsername(t *testing.T) {
	arkapi := NewArkClient(nil)

	params := DelegateQueryParams{UserName: "acf"}

	deleResponse, _, err := arkapi.GetDelegate(params)
	if deleResponse.Success {

		out, _ := json.Marshal(deleResponse.Delegate)
		log.Println(t.Name(), "Success, returned", string(out))

	} else {
		t.Error(err.Error())
	}
}

func TestGetDelegatePubKey(t *testing.T) {
	arkapi := NewArkClient(nil)

	params := DelegateQueryParams{PublicKey: "03e6397071866c994c519f114a9e7957d8e6f06abc2ca34dc9a96b82f7166c2bf9"}

	deleResponse, _, err := arkapi.GetDelegate(params)
	if deleResponse.Success {

		out, _ := json.Marshal(deleResponse.Delegate)
		log.Println(t.Name(), "Success, returned", string(out))

	} else {
		t.Error(err.Error())
	}
}

func TestGetDelegateVoters(t *testing.T) {
	arkapi := NewArkClient(nil)

	params := DelegateQueryParams{PublicKey: "03e6397071866c994c519f114a9e7957d8e6f06abc2ca34dc9a96b82f7166c2bf9"}

	deleResponse, _, err := arkapi.GetDelegateVoters(params)
	if deleResponse.Success {

		//calculating vote weight
		balance := 0
		for _, element := range deleResponse.Accounts {
			intBalance, _ := strconv.Atoi(element.Balance)
			balance += intBalance
		}

		log.Println(t.Name(), "Success, returned", len(deleResponse.Accounts), "voters for delegate with weight", balance)

	} else {
		t.Error(err.Error())
	}
}

func TestGetDelegateVoteWeight(t *testing.T) {
	arkapi := NewArkClient(nil)

	params := DelegateQueryParams{PublicKey: "03e6397071866c994c519f114a9e7957d8e6f06abc2ca34dc9a96b82f7166c2bf9"}

	voteWeight, _, _ := arkapi.GetDelegateVoteWeight(params)

	log.Println(t.Name(), "Success, returned delegate vote weight is", voteWeight)

}