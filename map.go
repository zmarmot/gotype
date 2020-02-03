package gotype
import(
    "sync"
)

type Set struct{
    m map[interface{}]bool
    sync.RWMutex
}

func NewSet()*Set{
    return &Set{m: map[interface{}]bool{},}
}

func(s *Set)Insert(data interface{}){
    s.Lock()
    defer s.Unlock()
    s.m[data] = true
}

func(s *Set)Remove(data interface{}){
    s.Lock()
    defer s.Unlock()
    delete(s.m,data)
}

func(s *Set)Contains(data interface{})bool{
    s.Lock()
    defer s.Unlock()
    _,ok := s.m[data]
    return ok
}

func(s *Set)List()[]interface{}{
    s.Lock()
    defer s.Unlock()
    list := []interface{}{}
    for k,_ := range s.m{
        list = append(list,k)
    }
    return list
}

func(s *Set)Length()int{
    return len(s.List())
}

func(s *Set)IsEmpty()bool{
    return s.Length() == 0
}

func(s *Set)Clear(){
    s.Lock()
    defer s.Unlock()
    s.m = map[interface{}]bool{}
}
