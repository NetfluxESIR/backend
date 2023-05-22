package api

import (
	"fmt"
	"github.com/NetfluxESIR/backend/pkg/api/gen"
)

func ValidateVideo(video gen.Video) error {
	if video.Title == "" {
		return fmt.Errorf("missing title")
	}
	return nil
}

func ValidateAccount(account gen.Account) error {
	if account.Email == "" || account.Password == "" {
		return fmt.Errorf("missing email or password")
	}
	if account.Role == nil {
		return fmt.Errorf("missing role")
	}
	switch *account.Role {
	case gen.ADMIN:
		return nil
	case gen.USER:
		return nil
	case gen.ROBOT:
		return nil
	default:
		return fmt.Errorf("invalid role")
	}
}

func ValidateStatus(status gen.ProcessingStatus) error {
	switch status {
	case gen.ProcessingStatusPENDING:
		return nil
	case gen.ProcessingStatusERROR:
		return nil
	case gen.ProcessingStatusFINISHED:
		return nil
	case gen.ProcessingStatusSTARTED:
		return nil
	default:
		return fmt.Errorf("invalid status")
	}
}

func ValidateStep(step gen.ProcessingStep) error {
	if step.Step == nil {
		return fmt.Errorf("missing step")
	}
	if step.Status == nil {
		return fmt.Errorf("missing status")
	}
	if err := ValidateStepStatus(*step.Status); err != nil {
		return err
	}
	if err := ValidateStepType(*step.Step); err != nil {
		return err
	}
	return nil
}

func ValidateStepStatus(status gen.ProcessingStepStatus) error {
	switch status {
	case gen.ProcessingStepStatusERROR:
		return nil
	case gen.ProcessingStepStatusFINISHED:
		return nil
	case gen.ProcessingStepStatusSTARTED:
		return nil
	default:
		return fmt.Errorf("invalid status")
	}
}

func ValidateStepType(step gen.ProcessingStepStep) error {
	switch step {
	case gen.ProcessingStepStepANIMALDETECTION:
		return nil
	case gen.ProcessingStepStepDOWNSCALE:
		return nil
	case gen.ProcessingStepStepNONE:
		return nil
	case gen.ProcessingStepStepLANGDETECTION:
		return nil
	case gen.ProcessingStepStepTRANSCRIPTION:
		return nil
	case gen.ProcessingStepStepFINISHED:
		return nil
	default:
		return fmt.Errorf("invalid step")
	}
}
