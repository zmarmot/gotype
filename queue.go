package gotype
import(
    "sync"
)

type element interface{}

type SliceQueue struct{
    array []element
    sync.RWMutex
}

func NewSliceQueue() *SliceQueue{
    return &SliceQueue{array:make([]element,0)}
}

func(s *SliceQueue)Size()int{
    return len(s.array)
}

func(s *SliceQueue)IsEmpty()bool{
    return s.Size() == 0
}

func(s *SliceQueue)Front()element{
    s.Lock()
    defer s.Unlock()
    if s.IsEmpty(){
        return nil
    }
    return s.array[0]
}

func(s *SliceQueue)Rear()element{
    s.Lock()
    defer s.Unlock()
    if s.IsEmpty(){
        return nil
    }
    return s.array[s.Size()-1]
}
//弹出最后一个元素
func(s *SliceQueue)Pop()element{
    s.Lock()
    defer s.Unlock()
    if s.IsEmpty(){
        return nil
    }
    ret := s.array[s.Size()-1]
    s.array = s.array[:s.Size()-1]
    return ret
}

func(s *SliceQueue)DeQueue()element{
    s.Lock()
    defer s.Unlock()
    if s.IsEmpty(){
        return nil
    }
    ret := s.array[0]
    s.array = s.array[1:s.Size()]
    return ret
}

func(s *SliceQueue)EnQueue(data element){
    s.Lock()
    defer s.Unlock()
    s.array = append(s.array,data)
}

func(s *SliceQueue)EnQueueFirst(data element){
    s.Lock()
    defer s.Unlock()
    newQueue := []element{data}
    s.array = append(newQueue,s.array...)
}

func(s *SliceQueue)Remove(data element){
    s.Lock()
    defer s.Unlock()
    for k,v := range s.array{
        if v == data{
            s.array = append(s.array[:k],s.array[k+1:]...)
        }
    }
}

func(s *SliceQueue)List()[]element{
    return s.array
}

type LinkedQueue struct{
    head *LinkNode
    end *LinkNode
    sync.RWMutex
}

func NewLinkedQueue() *LinkedQueue{
    return &LinkedQueue{}
}

func(p *LinkedQueue)Size()int{
    size := 0
    node := p.head
    for node != nil{
        size++
        node = node.Next
    }
    return size
}

func(p *LinkedQueue)IsEmpty()bool{
    return p.head == nil
}

func(p *LinkedQueue)EnQueue(data element){
    p.Lock()
    defer p.Unlock()
    node := &LinkNode{Data:data}
    if p.head == nil{
        p.head = node
        p.end = node
    }else{
        p.end.Next = node
        p.end = node
    }
}

func(p *LinkedQueue)EnQueueFirst(data element){
    p.Lock()
    defer p.Unlock()
    node := &LinkNode{Data:data}
    node.Next = p.head
    p.head = node
}

func(p *LinkedQueue)DeQueue()element{
    p.Lock()
    defer p.Unlock()
    if p.IsEmpty(){
        return nil
    }
    node := p.head
    p.head = p.head.Next
    if p.IsEmpty(){
        p.end = nil
    }
    return node
}

func(p *LinkedQueue)Front()element{
    if p.IsEmpty(){
        return nil
    }
    return p.head.Data
}

func(p *LinkedQueue)Rear()element{
    if p.IsEmpty(){
        return nil
    }
    return p.end.Data
}

func(p *LinkedQueue)Remove(data element){
    cur := p.head
    if cur.Data == data{
        p.head = p.head.Next
    }else{
        for cur.Next != nil {
			if cur.Next.Data == data{
				cur.Next = cur.Next.Next
			}else{
				cur = cur.Next
			}
		}
    }
}

func(p *LinkedQueue)List()[]element{
    ret := make([]element,0)
    node := p.head
    for node != nil{
        ret = append(ret,node.Data)
        node = node.Next
    }
    return ret
}

