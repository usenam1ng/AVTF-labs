package main

import (
	"errors"
	"fmt"
	"sync"
)

// stack

type NodeS struct {
	data string
	next *NodeS
}

type Stack struct {
	Name string
	head *NodeS
}

// queue

type NodeQ struct {
	data string
	next *NodeQ
}

type Queue struct {
	Name string
	head *NodeQ
	tail *NodeQ
}

// list

type NodeL struct {
	data int
	next *NodeL
}

type List struct {
	head *NodeL
}

// double list

type NodeDL struct {
	data int
	next *NodeDL
	prev *NodeDL
}

type DoublyList struct {
	head *NodeDL
	tail *NodeDL
}

// array

type Array struct {
	data   []int
	length int
}

// tree

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

//hash-table

type NodeH struct {
	Key   string
	Value string
}

type HashTable struct {
	Table    []*NodeH
	Capacity int
}

func (ht *HashTable) hashFunc(a int, key string) int {
	sum1 := 0
	sum2 := 0
	for _, c := range key {
		sum1 += int(c)
	}

	for _, c := range key {
		sum2 += (int(c) % 2)
	}

	ans := (sum1 + a*sum2) % ht.Capacity
	return ans
}

func NewHashTable(capacity int, stname string) *HashTable {
	return &HashTable{
		Table:    make([]*NodeH, capacity),
		Capacity: capacity,
	}
}

//set

type NodeSH struct {
	Key   string
	Value string
}

type Set struct {
	Table    []*NodeSH
	Capacity int
}

func (st Set) hashfuncSet(a int, key string) int {
	sum1 := 0
	sum2 := 0
	for _, c := range key {
		sum1 += int(c)
	}

	for _, c := range key {
		sum2 += (int(c) % 2)
	}

	ans := (sum1 + a*sum2) % st.Capacity
	return ans
}

func newSet(capacity int, stname string) *Set {
	return &Set{
		Table:    make([]*NodeSH, capacity),
		Capacity: capacity,
	}
}

type DataStructure struct {
	name       string
	hashTables []HashTable //`json:"hashTables"`
	stacks     []Stack     //`json:"stacks"`
	queues     []Queue     //`json:"queues"`
	sets       []Set       //`json:"sets"`
}

type MainDataStructure struct {
	datastructures []DataStructure
	mutex          sync.Mutex
}

//---------------------------------

func (stack *Stack) pop() (string, error) {
	if stack.head == nil {
		return "", errors.New("пустой стек")
	} else {
		x := stack.head.data
		stack.head = stack.head.next
		return x, nil
	}
}

func (stack *Stack) push(val string) {
	newnode := &NodeS{data: val}
	if stack.head == nil {
		stack.head = newnode
		stack.head.data = val
	} else {
		newnode.next = stack.head
		stack.head = newnode
		stack.head.data = val
	}
}

//queue

func (queue *Queue) pushQ(val string) {
	newnode := &NodeQ{data: val}
	if queue.head == nil {
		queue.head = newnode
		queue.tail = newnode
	} else {
		queue.tail.next = newnode
		queue.tail = newnode
	}
}

func (queue *Queue) popQ() (string, error) {
	if queue.head == nil {
		return "", errors.New("")
	} else {
		data := queue.head.data
		queue.head = queue.head.next
		return data, nil
	}
}

//hash-table

func (ht *HashTable) Add(key, value string) {
	index := ht.hashFunc(1, key)
	node := &NodeH{Key: key, Value: value}
	if ht.Table[index] == nil {
		ht.Table[index] = node
	} else {
		for i := 2; i < 32; i++ {
			index := ht.hashFunc(i, key)
			if ht.Table[index] == nil {
				ht.Table[index] = node
				break
			}
		}
	}
}

func (ht *HashTable) Get(key string) (string, bool) {
	index := ht.hashFunc(1, key)
	if ht.Table[index] != nil && ht.Table[index].Key == key {
		return ht.Table[index].Value, true
	}
	return "", false
}

func (ht *HashTable) Delete(key string) bool {
	index := ht.hashFunc(1, key)
	if ht.Table[index] != nil && ht.Table[index].Key == key {
		ht.Table[index].Key = "0"
		ht.Table[index].Value = "0"
		return true
	}
	return false
}

//set

func (st *Set) AddS(key, value string) {
	index := st.hashfuncSet(1, key)
	node := &NodeSH{Key: key, Value: value}
	if st.Table[index] == nil {
		st.Table[index] = node
	} else {
		for i := 2; i < 32; i++ {
			index := st.hashfuncSet(i, key)
			if st.Table[index] == nil {
				st.Table[index] = node
				break
			}
		}
	}
}

func (st *Set) GetS(key string) (string, bool) {
	index := st.hashfuncSet(1, key)
	if st.Table[index] != nil && st.Table[index].Key == key {
		return st.Table[index].Value, true
	}
	return "", false
}

func (st *Set) DeleteS(key string) bool {
	index := st.hashfuncSet(1, key)
	if st.Table[index] != nil && st.Table[index].Key == key {
		st.Table[index].Key = "0"
		st.Table[index].Value = "0"
		return true
	}
	return false
}

// list

func (ll *List) addlist(data int) { //O(1), если список пуст; O(n), если есть значения
	newNode := &NodeL{data: data}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *List) istrue() bool { //O(1), проверка наличия головы списка
	return ll.head == nil
}

func (ll *List) dellist(data int) error { //O(1), если один элемент; O(n), если более одного элемента
	if ll.head == nil {
		return errors.New("list is empty")
	}

	if ll.head.data == data {
		ll.head = ll.head.next
		return nil
	}

	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return nil
		}
		current = current.next
	}

	return errors.New("no element")
}

func (ll *List) outlist() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

// double list

func (dll *DoublyList) adddoblist(data int) {
	newNode := &NodeDL{data: data}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
}

func (dll *DoublyList) deldlist(data int) error {
	if dll.head == nil {
		return errors.New("список пуст")
	}

	if dll.head.data == data {
		dll.head = dll.head.next
		if dll.head != nil {
			dll.head.prev = nil
		}
		return nil
	}

	if dll.tail.data == data {
		dll.tail = dll.tail.prev
		if dll.tail != nil {
			dll.tail.next = nil
		}
		return nil
	}

	current := dll.head
	for current != nil {
		if current.data == data {
			current.prev.next = current.next
			current.next.prev = current.prev
			return nil
		}
		current = current.next
	}

	return errors.New("элемент не найден")
}

func (dll *DoublyList) dlout() { //O(n), т.к. перебор всех значений
	current := dll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

// array

func (arr *Array) arrst(index, value int) error {
	if index < 0 || index >= arr.length {
		return errors.New("index out of range")
	}
	arr.data[index] = value
	return nil
}

func (arr *Array) arrget(index int) (int, error) {
	if index < 0 || index >= arr.length {
		return 0, errors.New("index out of range")
	}
	return arr.data[index], nil
}

func (arr *Array) arrdel(index int) error {
	if index < 0 || index >= arr.length {
		return errors.New("index out of range")
	}
	copy(arr.data[index:], arr.data[index+1:])
	arr.length--
	return nil
}

func (arr *Array) arrapp(value int) {
	arr.data = append(arr.data, value)
	arr.length++
}

// tree

func (bst *BinarySearchTree) Insert(value int) {
	newNode := &TreeNode{data: value}

	if bst.root == nil {
		bst.root = newNode
		return
	} else {
		insertNode(bst.root, newNode)
	}
}

func insertNode(node, newNode *TreeNode) {
	if newNode.data < node.data {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else if newNode.data > node.data {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (bst *BinarySearchTree) Delete(value int) {
	bst.root = deleteNode(bst.root, value)
}

func deleteNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return root
	}

	if value < root.data {
		root.left = deleteNode(root.left, value)
	} else if value > root.data {
		root.right = deleteNode(root.right, value)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		minRight := findMinNode(root.right)
		root.data = minRight.data
		root.right = deleteNode(root.right, minRight.data)
	}

	return root
}

func findMinNode(node *TreeNode) *TreeNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (bst *BinarySearchTree) Search(value int) bool {
	return search(bst.root, value)
}

func search(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	if value < node.data {
		return search(node.left, value)
	} else if value > node.data {
		return search(node.right, value)
	}
	return true
}

func (bst *BinarySearchTree) printtree() {
	prnt(bst.root, "", true)
}

func prnt(node *TreeNode, indent string, last bool) {
	if node != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("--")
			indent += "   "
		} else {
			fmt.Print("|--")
			indent += "|  "
		}
		fmt.Println(node.data)
		prnt(node.left, indent, false)
		prnt(node.right, indent, true)
	}
}

func (bst *BinarySearchTree) OutElements() {
	var elements []int
	allelement(bst.root, &elements)
	for _, element := range elements {
		fmt.Println(element)
	}
}

func allelement(node *TreeNode, elements *[]int) {
	if node != nil {
		*elements = append(*elements, node.data)
		allelement(node.left, elements)
		allelement(node.right, elements)
	}
}

// menu

func main() {
	tabl := NewHashTable(512, "")
	stk := Stack{}
	st := newSet(512, "")
	que := Queue{}
	lst := List{}
	dlst := DoublyList{}
	ar := Array{}
	tr := BinarySearchTree{}

	for true {
		var a int
		fmt.Println("Выберите структуру. 1 - стек, 2 - очередь, 3 - хеш-таблица, \n4 - множество, 5 - список, 6 - двусвязный список, \n7 - массив, 8 - дерево, 0 - выход")
		fmt.Scan(&a)
		if a == 1 {
			fmt.Println("1 - запись в стек \n2 - чтение из стека")
			var b int
			fmt.Scan(&b)
			if b == 1 {
				fmt.Println("Введите строку для записи в стек")
				var s string
				fmt.Scan(&s)
				stk.push(s)
			} else if b == 2 {
				fmt.Println(stk.pop())
			}
		} else if a == 2 {
			fmt.Println("1 - запись в очередь \n2 - чтение из очереди")
			var b int
			fmt.Scan(&b)
			if b == 1 {
				fmt.Println("Введите строку для записи в очередь")
				var s string
				fmt.Scan(&s)
				que.pushQ(s)
			} else if b == 2 {
				fmt.Println(que.popQ())
			}
		} else if a == 3 {
			fmt.Println("1 - запись в хеш-таблицу \n2 - чтение из хеш-таблицы \n3 - удаление из хеш-таблицы")
			var b int
			fmt.Scan(&b)
			if b == 1 {
				fmt.Println("Введите ключ для записи в хеш-таблицу")
				var s string
				fmt.Scan(&s)
				fmt.Println("Введите значение для записи в хеш-таблицу")
				var s1 string
				fmt.Scan(&s1)
				tabl.Add(s, s1)
			} else if b == 2 {
				fmt.Println("Введите ключ для чтения из хеш-таблицы")
				var s string
				fmt.Scan(&s)
				fmt.Println(tabl.Get(s))
			} else if b == 3 {
				fmt.Println("Введите ключ для удаления из хеш-таблицы")
				var s string
				fmt.Scan(&s)
				tabl.Delete(s)
			}
		} else if a == 4 {
			fmt.Println("1 - запись в множество \n2 - чтение из множества \n3 - вывести множество на экран")
			var b int
			fmt.Scan(&b)
			if b == 1 {
				fmt.Println("Введите значение для записи в множество")
				var s string
				fmt.Scan(&s)
				st.AddS(s, "")
			} else if b == 2 {
				fmt.Println("Введите значение для чтения из множества")
				var s string
				fmt.Scan(&s)
				fmt.Println(st.GetS(s))
			} else if b == 3 {
				fmt.Println("Введите значение для удаления из множества")
				var s string
				fmt.Scan(&s)
				st.DeleteS(s)
			}
		} else if a == 5 {
			fmt.Println("1 - запись в список \n2 - удаление из спика \n3 - вывод списка на экран")
			var b int
			fmt.Scan(&b)
			if b == 1 {
				fmt.Println("Введите число для записи в список")
				var s int
				fmt.Scan(&s)
				lst.addlist(s)
			} else if b == 2 {
				fmt.Println("Введите число для удаления из списка")
				var s int
				fmt.Scan(&s)
				lst.dellist(s)
			} else if b == 3 {
				lst.outlist()
			}
		} else if a == 6 {
			fmt.Println("1 - запись в двусвязный список \n2 - удаление из двусвязного спика \n3 - вывод списка на двусвязного экран")
			var b int
			fmt.Scan(&b)
			if b == 1 {
				fmt.Println("Введите число для записи в двусвязноый список")
				var s int
				fmt.Scan(&s)
				dlst.adddoblist(s)
			} else if b == 2 {
				fmt.Println("Введите число для удаления из двусвязного списка")
				var s int
				fmt.Scan(&s)
				dlst.deldlist(s)
			} else if b == 3 {
				dlst.dlout()
			}
		} else if a == 7 {
			fmt.Println("1 -запись по индексу \n2 - чтение из массива \n3 - добавление в массив \n4 - удаление из массива")
			var b int
			fmt.Scan(&b)
			if b == 1 {
				fmt.Println("Введите ключ для записи в массив")
				var s int
				fmt.Scan(&s)
				fmt.Println("Введите значение для записи в массив")
				var s1 int
				fmt.Scan(&s1)
				ar.arrst(s, s1)
			} else if b == 2 {
				fmt.Println("Введите ключ для чтения из массива")
				var s int
				fmt.Scan(&s)
				fmt.Println(ar.arrget(s))
			} else if b == 3 {
				fmt.Println("Введите значение для добавления в массив")
				var s int
				fmt.Scan(&s)
				ar.arrget(s)
			} else if b == 4 {
				fmt.Println("Введите индекс для удаления из массива")
				var s int
				fmt.Scan(&s)
				ar.arrdel(s)
			}
		} else if a == 8 {
			fmt.Println("1 - запись в дерево \n2 - удаление из дерева \n3 - печать дерева \n4 - поиск по дереву")
			var b int
			fmt.Scan(&b)
			if b == 1 {
				fmt.Println("Введите значение для записи в дерево")
				var s int
				fmt.Scan(&s)
				tr.Insert(s)
			} else if b == 2 {
				fmt.Println("Введите значение для удаления из дерева")
				var s int
				fmt.Scan(&s)
				tr.Delete(s)
			} else if b == 3 {
				tr.printtree()
			} else if b == 4 {
				fmt.Println("Введите значение для поиска по дереву")
				var s int
				fmt.Scan(&s)
				fmt.Println(tr.Search(s))
			}
		} else if a == 0 {
			break
		}
	}
}
