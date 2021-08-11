package auth0

import "gopkg.in/auth0.v5/management"

func expandActionTrigger(d ResourceData) (actionTrigger *management.ActionTrigger) {
	var t []management.ActionTrigger
	t := List(d, "action_trigger").Elem(func(d ResourceData) {
		actionTrigger = &management.ActionTrigger{
			Version: String(d, "name"),
			Status:  String(d, "status"),
		}
	})
	return
}

func expandActionDependency(d ResourceData) (actionDependency *management.ActionDependency) {
	List(d, "action_dependency").Elem(func(d ResourceData) {
		actionDependency = &management.ActionDependency{
			Name:    String(d, "name"),
			Version: String(d, "version"),
		}
	})
	return
}

func expandActionSecret(d ResourceData) (actionSecret *management.ActionSecret) {
	List(d, "action_secret").Elem(func(d ResourceData) {
		actionSecret = &management.ActionSecret{
			Name:  String(d, "name"),
			Value: String(d, "value"),
		}
	})
	return
}
