package main

type BinarySearchTree struct {
    length int
    root node
}

type node struct {
    parentNode *node
    leftNode *node
    rightNode *node
    element int
}

func (bst BinarySearchTree) IsExist(i int) bool{
    if &bst.root == nil {
        return false
    }

    var currentNode *node = &bst.root

    for currentNode != nil {
        if currentNode.element == i {
        return true
        }

        if i < currentNode.element {
            currentNode = currentNode.leftNode
        } else {
            currentNode = currentNode.rightNode
        }
    }

    return false
}

func (bst *BinarySearchTree) Add(i int) {
    var newNode node
    newNode.element = i

    // Если длина древа равна нулю, то добавляемый элемент станет корнем
    if bst.length == 0 {
        bst.root = newNode
        bst.length = bst.length + 1
    }

    // Если такой элемент есть в дереве, то происходит return
    if bst.IsExist(i) {
        return
    }

    var parentNode *node = &bst.root

    for true {
        if newNode.element < parentNode.element {
            if parentNode.leftNode == nil{
                parentNode.leftNode = &newNode
                newNode.parentNode = parentNode
                bst.length = bst.length + 1
                return
            } else {
                parentNode = parentNode.leftNode
            }
        } else {
            if parentNode.rightNode == nil{
                parentNode.rightNode = &newNode
                newNode.parentNode = parentNode
                bst.length = bst.length + 1
                return
            } else {
                parentNode = parentNode.rightNode
            }
        }
    }
}

func (bst *BinarySearchTree) Delete(i int) {
    if bst.IsExist(i) == false {
        return
    }

    var isNodeRight bool
    var parentNode *node = &bst.root
    var currentNode *node

    // Нахождение нода-родителя удаляемого нода
    for true {
        if parentNode.rightNode != nil && parentNode.rightNode.element == i {
            isNodeRight = true
            currentNode = parentNode.rightNode
            break
        } else if parentNode.leftNode != nil && parentNode.leftNode.element == i {
            isNodeRight = false
            currentNode = parentNode.leftNode
            break
        } else if bst.root.element == i {
            currentNode = &bst.root
            break
        }
        if i < parentNode.element {
            parentNode = parentNode.leftNode
        } else {
            parentNode = parentNode.rightNode
        }
    }

    // Удаление элемента в зависимости от количества его потомков и их расположения
    if currentNode.leftNode != nil && currentNode.rightNode != nil {
        // Ищим минимальный элемент и его родителя в правом поддереве на замену удаляемому
        var minElParent *node = currentNode
        var minEl *node = currentNode.rightNode
        var isMinElNodeRight bool = true

        for minEl.leftNode != nil {
            isMinElNodeRight = false
            minElParent = minEl
            minEl = minEl.leftNode
        }

        // Даём новому ноду ссылки на дочернии ноды удаляемого нода
        minEl.leftNode = currentNode.leftNode
        minEl.rightNode = currentNode.rightNode
        minEl.rightNode.parentNode = minEl
        minEl.leftNode.parentNode = minEl

        // Устанавливаем родительскому ноду ссылку на новый элемент, вместо удалённого
        if isNodeRight {
            parentNode.rightNode = minEl
        } else {
            parentNode.leftNode = minEl
        }

        // Удаление элемента из древа
        if isMinElNodeRight {
            minEl.rightNode = nil
        } else {
            minElParent.leftNode = nil
        }

        // Проверка, является ли удаляемый элемент корнем
        if bst.root.element != i {
            minEl.parentNode = parentNode
        } else {
            bst.root = *minEl
            bst.root.parentNode = nil
        }
    } else if currentNode.leftNode == nil && currentNode.rightNode == nil {
        // Если удаляем лист
        if isNodeRight {
            parentNode.rightNode = nil
        } else {
            parentNode.leftNode = nil
        }
        // Если удаляемый элемент - корень
        if bst.root.element == i {
            bst.root = node{}
        }
    } else if currentNode.leftNode != nil {
        // Если удаляем нод, у которого есть потомок слева
        if isNodeRight {
            parentNode.rightNode = currentNode.leftNode
        } else {
            parentNode.leftNode = currentNode.leftNode
        }
        // Проверка, является ли удаляемый элемент корнем
        if bst.root.element != i {
            currentNode.leftNode.parentNode = parentNode
        } else {
            bst.root = *currentNode.leftNode
            bst.root.parentNode = nil
        }
    } else {
        // Если удаляем нод, у которого есть потомок справа
        if isNodeRight {
            parentNode.rightNode = currentNode.rightNode
        } else {
            parentNode.leftNode = currentNode.rightNode
        }
        // Проверка, является ли удаляемый элемент корнем
        if bst.root.element != i {
            currentNode.rightNode.parentNode = parentNode
        } else {
            bst.root = *currentNode.rightNode
            bst.root.parentNode = nil
        }
    }

    bst.length = bst.length - 1
}
