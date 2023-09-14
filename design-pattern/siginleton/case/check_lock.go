package _case

import "sync"

var checkLockConf *Config
var checkLockConfLock = new(sync.Mutex)

func GetCheckLockConf() *Config {
	if checkLockConf == nil {
		checkLockConfLock.Lock()
		defer checkLockConfLock.Unlock()
		if checkLockConf == nil {
			checkLockConf = &Config{
				name: "checkLock",
			}
		}
	}
	return checkLockConf
}
