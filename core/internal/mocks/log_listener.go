// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	log "github.com/smartcontractkit/chainlink/core/services/log"
	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"
)

// LogListener is an autogenerated mock type for the Listener type
type LogListener struct {
	mock.Mock
}

// HandleLog provides a mock function with given fields: lb, err
func (_m *LogListener) HandleLog(lb log.Broadcast, err error) {
	_m.Called(lb, err)
}

// IsV2Job provides a mock function with given fields:
func (_m *LogListener) IsV2Job() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// JobID provides a mock function with given fields:
func (_m *LogListener) JobID() *models.ID {
	ret := _m.Called()

	var r0 *models.ID
	if rf, ok := ret.Get(0).(func() *models.ID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ID)
		}
	}

	return r0
}

// JobIDV2 provides a mock function with given fields:
func (_m *LogListener) JobIDV2() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// OnConnect provides a mock function with given fields:
func (_m *LogListener) OnConnect() {
	_m.Called()
}

// OnDisconnect provides a mock function with given fields:
func (_m *LogListener) OnDisconnect() {
	_m.Called()
}
