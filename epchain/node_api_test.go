package epchain

import (
	"fmt"
	"testing"
)

const nodeurl  = "http://ip:port"

func Test_getBlockHeight(t *testing.T) {
	c := NewClient(nodeurl, false)

	r, err := c.getBlockHeight()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Test_getBlockHash(t *testing.T) {
	c := NewClient(nodeurl, false)

	height := uint64(184952)

	r, err := c.getBlockHash(height)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_getBalance(t *testing.T) {
	c := NewClient(nodeurl, true)

	address := "epc1t47a6yl9fpygvaxh485v5yd7rykcf6ygu4vkqv"

	r, err := c.getBalance(address, "uepc")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_getTransaction(t *testing.T) {
	c := NewClient(nodeurl, true)
	txid := "575DFAA744177AB9A8469411BBA78B33F1BFF7594A28D3A8BC6F7C443CCA8D96" //"9KBoALfTjvZLJ6CAuJCGyzRA1aWduiNFMvbqTchfBVpF"

	path := "/txs/" + txid
	r, err := c.Call(path, nil, "GET")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	trx := NewTransaction(r, "cosmos-sdk/StdTx", "cosmos-sdk/MsgSend", "uepc")

	fmt.Println(trx)
}

func Test_getBlockByHeight(t *testing.T) {
	height := uint64(429734)
	c := NewClient("http://127.0.0.1:1317", false)
	result, err := c.getBlockByHeight(height)
	if err != nil {
		t.Error("get block failed!")
	} else {
		fmt.Println(result)
	}
}

func Test_sequence(t *testing.T) {
	addr := "epc1ulpuamh6jey3rxqtjn4xd9gas6k2atztvds3cj"
	c := NewClient(nodeurl, false)
	accountnumber, sequence, err := c.getAccountNumberAndSequence(addr)
	fmt.Println(err)
	fmt.Println(accountnumber)
	fmt.Println(sequence)
}
