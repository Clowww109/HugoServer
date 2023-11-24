package main

import "testing"

func TestNewNode(t *testing.T) {
	key := 10
	node := NewNode(key)

	if node.Key != key {
		t.Errorf("Expected node key to be %d, got %d", key, node.Key)
	}
}

func TestHeight(t *testing.T) {
	// Тестирование для нулевого узла
	var node *Node2
	result := height(node)
	if result != 0 {
		t.Errorf("Ожидалось, что высота нулевого узла будет 0, получено %d", result)
	}

	// Тестирование для узла с высотой
	nodeWithHeight := &Node2{Key: 10, Height: 3}
	result = height(nodeWithHeight)
	if result != 3 {
		t.Errorf("Ожидалось, что высота узла будет 3, получено %d", result)
	}
}

func TestMax(t *testing.T) {
	// Тестирование, когда a больше b
	result := max(5, 3)
	if result != 5 {
		t.Errorf("Ожидалось, что максимальное значение из 5 и 3 будет 5, получено %d", result)
	}

	// Тестирование, когда b больше a
	result = max(2, 8)
	if result != 8 {
		t.Errorf("Ожидалось, что максимальное значение из 2 и 8 будет 8, получено %d", result)
	}

	// Тестирование, когда a и b равны
	result = max(4, 4)
	if result != 4 {
		t.Errorf("Ожидалось, что максимальное значение из 4 и 4 будет 4, получено %d", result)
	}
}

func TestUpdateHeight(t *testing.T) {
	// Создание узла с левым и правым поддеревьями
	node := &Node2{Key: 10}
	node.Left = &Node2{Key: 5}
	node.Right = &Node2{Key: 15}

	// Обновление высоты узла
	updateHeight(node)

	expectedHeight := 2 // 1 (текущий узел) + 1 (максимальная высота левого и правого поддеревьев)
	if node.Height != expectedHeight {
		t.Errorf("Ожидалось, что высота узла будет %d, получено %d", expectedHeight, node.Height)
	}
}

func TestGetBalance(t *testing.T) {
	// Создание узла с левым и правым поддеревьями
	node := &Node2{Key: 10}
	node.Left = &Node2{Key: 5}
	node.Right = &Node2{Key: 15}

	// Обновление высоты узла
	updateHeight(node.Left)
	updateHeight(node.Right)

	balance := getBalance(node)

	expectedBalance := 1 // высота левого поддерева (1) - высота правого поддерева (0)
	if balance != expectedBalance {
		t.Errorf("Ожидалось, что баланс узла будет %d, получено %d", expectedBalance, balance)
	}
}

func TestLeftRotate(t *testing.T) {
	// Создание узла x и его правого поддерева y
	x := &Node2{Key: 10}
	x.Right = &Node2{Key: 20}
	x.Right.Left = &Node2{Key: 15}

	// Поворот влево
	result := leftRotate(x)

	// Проверка правильности поворота
	if result != x.Right {
		t.Errorf("Ожидалось, что функция leftRotate вернет правое поддерево x, получено %v", result)
	}
	if x.Right.Left != x {
		t.Error("Ожидалось, что левое поддерево узла x станет правым поддеревом узла y")
	}
}

func TestRightRotate(t *testing.T) {
	// Создание узла y и его левого поддерева x
	y := &Node2{Key: 20}
	y.Left = &Node2{Key: 10}
	y.Left.Right = &Node2{Key: 15}

	// Поворот вправо
	result := rightRotate(y)

	// Проверка правильности поворота
	if result != y.Left {
		t.Errorf("Ожидалось, что функция rightRotate вернет левое поддерево y, получено %v", result)
	}
	if y.Left.Right != y {
		t.Error("Ожидалось, что правое поддерево узла y станет левым поддеревом узла x")
	}

}
