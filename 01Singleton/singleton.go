package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// 单例模式

var (
	once     sync.Once
	instance *shoppingCartManager
)

type shoppingCartManager struct {
	cart map[string]int
	keys []string // 保存插入顺序
	mu   sync.Mutex
}

// addToCart 添加购物车
func (m *shoppingCartManager) addToCart(name string, num int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.cart[name]; !ok {
		// 不在购物车就新添加
		m.keys = append(m.keys, name)
	}
	m.cart[name] += num
}

// showCartList 显示购物车列表
func (m *shoppingCartManager) showCartList() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, name := range m.keys {
		fmt.Printf("%s %d\n", name, m.cart[name])
	}
}

// GetInstance 获取购物车实例
func GetInstance() *shoppingCartManager {
	once.Do(func() {
		instance = &shoppingCartManager{
			cart: make(map[string]int),
		}
	})
	return instance
}

func main() {
	cart := GetInstance()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Text()
		if in == "" {
			break
		}

		info := strings.Fields(in)
		shopName := info[0]
		shopNum, _ := strconv.Atoi(info[1])
		cart.addToCart(shopName, shopNum)
	}
	cart.showCartList()
}
