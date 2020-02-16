package gotype

import(
    "fmt"
)

//二叉树定义
type BNode struct{
    Data interface{}
    Left *BNode
    Right *BNode
}

func NewBNode()*BNode{
    return &BNode{}
}

func ArrayToTree(data []int,start,end int)*BNode{
    var root *BNode
    if end >= start{
        root = NewBNode()
        mid := (start+end+1)/2
        root.Data = data[mid]
        root.Left = ArrayToTree(data,start,mid-1)
        root.Right = ArrayToTree(data,mid+1,end)
    }
    return root
}
//用中序遍历的方式打印出二叉树结点的内容
func PrintTreeMidOrder(root *BNode){
    if root == nil{
        return
    }
    //遍历左结点树
    if root.Left != nil{
        PrintTreeMidOrder(root.Left)
    }
    //打印root结点
    fmt.Print(root.Data," ")
    if root.Right != nil{
        PrintTreeMidOrder(root.Right)
    }
}
//层序打印二叉树
func PrintTreeLayer(root *BNode){
    if root == nil{
        return
    }
    var p *BNode
    queue := NewSliceQueue()
    //根结点入队列
    queue.EnQueue(root)
    for queue.Size() > 0{
        p = queue.DeQueue().(*BNode)
        //打印当前结点
        fmt.Print(p.Data," ")
        if p.Left != nil{
            queue.EnQueue(p.Left)
        }
        if p.Right != nil{
            queue.EnQueue(p.Right)
        }
    }
}

//Trie树

type TrieNode struct{
    IsLeaf bool
    Url string
    Child []*TrieNode
}

func NewTrieNode(count int) *TrieNode{
    return &TrieNode{IsLeaf: false, Url: "", Child: make([]*TrieNode, count)}
}

//AVL树
type AVLNode struct{
    Data int
    Height int
    Count int
    Left *AVLNode
    Right *AVLNode
}

func NewAVLNode() *AVLNode{
    return &AVLNode{}
}

