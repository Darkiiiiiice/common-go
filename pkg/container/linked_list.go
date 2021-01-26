package container

type LinkedList struct {
    length int
    root   *Element
}

type Element struct {
    Val  interface{}
    next *Element
}

func InitLinkedList() *LinkedList {
    var l = new(LinkedList)
    l.root = nil
    l.length = 0
    return l
}

func (l *LinkedList) InsertOfHead(elem interface{}) {
    var n = &Element{
        Val:  elem,
        next: nil,
    }

    n.next = l.root
    l.root = n
    l.length++
}

func (l *LinkedList) InsertOfTail(elem interface{}) {
    var n = &Element{
        Val:  elem,
        next: nil,
    }

    if l.root == nil {
        l.root = n
        l.length++
        return
    }

    var cur = l.root

    for cur.next != nil {
        cur = cur.next
    }

    cur.next = n
    l.length++
}

func (l *LinkedList) InsertOfIndex(elem interface{}, index int) {
    var n = &Element{
        Val:  elem,
        next: nil,
    }

    if index <= 0 {
        l.InsertOfHead(elem)
        return
    } else if index >= l.length {
        l.InsertOfTail(elem)
        return
    }

    var cur = l.root

    for index-1 > 0 && cur.next != nil {
        cur = cur.next
        index--
    }

    n.next = cur.next
    cur.next = n
    l.length++
}

func (l *LinkedList) Get(index int) interface{} {
    var cur = l.root

    for index > 0 && cur.next != nil {
        cur = cur.next
        index--
    }
    return cur.Val
}


func (l *LinkedList) Search(elem interface{}) int {
    var cur = l.root

    for i:=0; cur.next != nil; i++ {
        if cur.Val == elem {
            return i
        }
        cur = cur.next
    }

    return -1
}

func (l *LinkedList) Remove(index int) interface{} {
    var cur = l.root

    if index < 1 {
        var tmp = cur
        l.root = cur.next
        return tmp.Val
    }

    if index >= l.length {
        index = l.length - 1
    }

    for index-1 > 0 && cur.next != nil {
        cur = cur.next
        index--
    }

    var tmp = cur.next
    cur.next = cur.next.next
    l.length--

    return tmp.Val
}

func (l *LinkedList) ToSlice() []interface{} {
    var slice = make([]interface{}, l.length)

    var i = 0
    var cur = l.root

    for i < l.length {
        slice[i] = cur.Val
        cur = cur.next
        i++
    }

    return slice
}
