package design

type SaveUnvalidatedApply = func(p UnvalidatedApplyForCVRegistration)

type ApplyForCVRegistrationWorkflow = func(SaveUnvalidatedApply)
