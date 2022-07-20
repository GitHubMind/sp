package main

import (
	"container/ring"
	"log"
	"time"
)

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func singleChiane() {
	node := new(LinkNode)
	node.Data = 2

	node1 := new(LinkNode)
	node1.Data = 3
	node2 := new(LinkNode)
	node2.Data = 4

	node.NextNode = node1
	node1.NextNode = node2
	node2.NextNode = node
	//头是
	nowNode := node
	for {
		if nowNode != nil {
			log.Println(nowNode.Data)
			nowNode = nowNode.NextNode
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
}

//func ()

func main() {
	r := new(ring.Ring)

}
