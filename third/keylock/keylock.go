package keylock

import "sync"

// KeyLock 数据锁配置
type KeyLock struct {
	mu    sync.Mutex               // 加解锁的大锁
	locks map[string]*sync.RWMutex // 每个key的小锁
}

// NewKeyLock 新增key的锁
func NewKeyLock() *KeyLock {
	return &KeyLock{
		mu:    sync.Mutex{},
		locks: map[string]*sync.RWMutex{},
	}
}

// Lock 对key持互斥锁锁
func (l *KeyLock) Lock(key string) {
	l.mu.Lock()
	kmu, ok := l.locks[key]
	if !ok { // 锁不存在，初始化一个设置
		kmu = &sync.RWMutex{}
		l.locks[key] = kmu
	}
	kmu.Lock()
	l.mu.Unlock()
}

// Unlock 对加了写锁的key解锁
func (l *KeyLock) Unlock(key string) {
	l.mu.Lock()
	if kmu, ok := l.locks[key]; ok {
		kmu.Unlock()
	}
	l.mu.Unlock()
}

// RLocker 对key加读锁
func (l *KeyLock) RLocker(key string) {
	l.mu.Lock()
	if kmu, ok := l.locks[key]; !ok {
		kmu = &sync.RWMutex{}
		l.locks[key] = kmu
	} else {
		kmu.RLocker()
	}
	l.mu.Unlock()
}

// RUnlock 对加了读锁的key解锁
func (l *KeyLock) RUnlock(key string) {
	l.mu.Lock()
	if kmu, ok := l.locks[key]; ok {
		kmu.RUnlock()
	}
	l.mu.Unlock()
}
