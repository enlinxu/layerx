package fakes

import (
	"github.com/mesos/mesos-go/mesosproto"
	"github.com/gogo/protobuf/proto"
	"github.com/mesos/mesos-go/mesosproto/scheduler"
)


func FakeSubscribeCall() *scheduler.Call {
	callType := scheduler.Call_SUBSCRIBE
	return &scheduler.Call{
		FrameworkId: &mesosproto.FrameworkID{
			Value: proto.String("fake_framework_id"),
		},
		Type: &callType,
		Subscribe: &scheduler.Call_Subscribe{
			FrameworkInfo: FakeFramework(),
		},
	}
}

func FakeDeclineOffersCall(frameworkId string, offerIds ...string) *scheduler.Call {
	callType := scheduler.Call_DECLINE
	mesosOfferIds := []*mesosproto.OfferID{}
	for _, offerId := range offerIds {
		mesosOfferIds = append(mesosOfferIds, &mesosproto.OfferID{
			Value: proto.String(offerId),
		})
	}
	return &scheduler.Call {
		FrameworkId: &mesosproto.FrameworkID{
			Value: proto.String(frameworkId),
		},
		Type: &callType,
		Decline: &scheduler.Call_Decline{
			OfferIds: mesosOfferIds,
		},
	}
}

func FakeReconcileTasksCall(frameworkId string, taskIds ...string) *scheduler.Call {
	callType := scheduler.Call_RECONCILE
	reconcileTasks := []*scheduler.Call_Reconcile_Task{}
	for _, taskId := range taskIds {
		reconcileTasks = append(reconcileTasks, &scheduler.Call_Reconcile_Task{
			TaskId: &mesosproto.TaskID{
				Value: proto.String(taskId),
			},
		})
	}
	return &scheduler.Call {
		FrameworkId: &mesosproto.FrameworkID{
			Value: proto.String(frameworkId),
		},
		Type: &callType,
		Reconcile: &scheduler.Call_Reconcile{
			Tasks: reconcileTasks,
		},
	}
}