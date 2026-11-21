package spa

import (
	"errors"
	"sync"
)

type SpaManager struct {
	mu      sync.RWMutex
	members map[string]*Customer
}

func NewSpaManager() *SpaManager {
	m := &SpaManager{members: make(map[string]*Customer)}
	m.members["0909123456"] = &Customer{ID: 1, Name: "Thu Trang", Phone: "0909123456", Points: 350, Tier: TierSilver, Sessions: 5}
	return m
}

func (m *SpaManager) CheckInSession(phone string) (*Customer, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	c, ok := m.members[phone]
	if !ok {
		return nil, errors.New("khách hàng chưa đăng ký thành viên")
	}
	if c.Sessions <= 0 {
		return nil, errors.New("khách hàng đã sử dụng hết số buổi liệu trình")
	}

	c.Sessions--
	c.Points += 50 // Tích 50 điểm mỗi buổi chăm sóc

	if c.Points >= 500 {
		c.Tier = TierGold
	}
	return c, nil
}
