// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"github.com/cloudfoundry/cli/cf/actors"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeServicePlanActor struct {
	GetServiceWithSinglePlanStub        func(string, string) (models.ServiceOffering, error)
	getServiceWithSinglePlanMutex       sync.RWMutex
	getServiceWithSinglePlanArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceWithSinglePlanReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	UpdateServicePlanAvailabilityStub        func(models.ServiceOffering, bool) error
	updateServicePlanAvailabilityMutex       sync.RWMutex
	updateServicePlanAvailabilityArgsForCall []struct {
		arg1 models.ServiceOffering
		arg2 bool
	}
	updateServicePlanAvailabilityReturns struct {
		result1 error
	}
}

func (fake *FakeServicePlanActor) GetServiceWithSinglePlan(arg1 string, arg2 string) (models.ServiceOffering, error) {
	fake.getServiceWithSinglePlanMutex.Lock()
	defer fake.getServiceWithSinglePlanMutex.Unlock()
	fake.getServiceWithSinglePlanArgsForCall = append(fake.getServiceWithSinglePlanArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	if fake.GetServiceWithSinglePlanStub != nil {
		return fake.GetServiceWithSinglePlanStub(arg1, arg2)
	} else {
		return fake.getServiceWithSinglePlanReturns.result1, fake.getServiceWithSinglePlanReturns.result2
	}
}

func (fake *FakeServicePlanActor) GetServiceWithSinglePlanCallCount() int {
	fake.getServiceWithSinglePlanMutex.RLock()
	defer fake.getServiceWithSinglePlanMutex.RUnlock()
	return len(fake.getServiceWithSinglePlanArgsForCall)
}

func (fake *FakeServicePlanActor) GetServiceWithSinglePlanArgsForCall(i int) (string, string) {
	fake.getServiceWithSinglePlanMutex.RLock()
	defer fake.getServiceWithSinglePlanMutex.RUnlock()
	return fake.getServiceWithSinglePlanArgsForCall[i].arg1, fake.getServiceWithSinglePlanArgsForCall[i].arg2
}

func (fake *FakeServicePlanActor) GetServiceWithSinglePlanReturns(result1 models.ServiceOffering, result2 error) {
	fake.GetServiceWithSinglePlanStub = nil
	fake.getServiceWithSinglePlanReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServicePlanActor) UpdateServicePlanAvailability(arg1 models.ServiceOffering, arg2 bool) error {
	fake.updateServicePlanAvailabilityMutex.Lock()
	defer fake.updateServicePlanAvailabilityMutex.Unlock()
	fake.updateServicePlanAvailabilityArgsForCall = append(fake.updateServicePlanAvailabilityArgsForCall, struct {
		arg1 models.ServiceOffering
		arg2 bool
	}{arg1, arg2})
	if fake.UpdateServicePlanAvailabilityStub != nil {
		return fake.UpdateServicePlanAvailabilityStub(arg1, arg2)
	} else {
		return fake.updateServicePlanAvailabilityReturns.result1
	}
}

func (fake *FakeServicePlanActor) UpdateServicePlanAvailabilityCallCount() int {
	fake.updateServicePlanAvailabilityMutex.RLock()
	defer fake.updateServicePlanAvailabilityMutex.RUnlock()
	return len(fake.updateServicePlanAvailabilityArgsForCall)
}

func (fake *FakeServicePlanActor) UpdateServicePlanAvailabilityArgsForCall(i int) (models.ServiceOffering, bool) {
	fake.updateServicePlanAvailabilityMutex.RLock()
	defer fake.updateServicePlanAvailabilityMutex.RUnlock()
	return fake.updateServicePlanAvailabilityArgsForCall[i].arg1, fake.updateServicePlanAvailabilityArgsForCall[i].arg2
}

func (fake *FakeServicePlanActor) UpdateServicePlanAvailabilityReturns(result1 error) {
	fake.UpdateServicePlanAvailabilityStub = nil
	fake.updateServicePlanAvailabilityReturns = struct {
		result1 error
	}{result1}
}

var _ actors.ServicePlanActor = new(FakeServicePlanActor)
